package devserver

import (
	"bufio"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

// Live-reload SSE hub
type sseHub struct {
	clients map[chan string]struct{}
	mu      sync.Mutex
}

func newSSEHub() *sseHub { return &sseHub{clients: make(map[chan string]struct{})} }

func (h *sseHub) addClient(ch chan string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[ch] = struct{}{}
}

func (h *sseHub) removeClient(ch chan string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, ch)
	close(ch)
}

func (h *sseHub) broadcast(msg string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for ch := range h.clients {
		select {
		case ch <- msg:
		default:
		}
	}
}

//go:embed wasm_exec-tinygo.js
var wasmExecJSTinyGo []byte

//go:embed wasm_exec.js
var wasmExecJS []byte

// getWasmExecJS returns the wasm_exec.js content from file
func getWasmExecJS() []byte {
	return wasmExecJS
}

// BuildViteAssets builds CSS and JS assets using Vite at repo root (legacy)
func BuildViteAssets(mode string) error {
	return BuildViteAssetsWithConfig(ViteConfig{Root: ".", BuildMode: mode})
}

// BuildViteAssetsWithConfig builds CSS and JS assets using Vite with the given config
func BuildViteAssetsWithConfig(cfg ViteConfig) error {
	mode := cfg.BuildMode
	if mode == "" {
		mode = "development"
	}
	logutil.Logf("==> Building Vite assets in %s mode at %s...\n", mode, cfg.Root)

	root := cfg.Root
	if strings.TrimSpace(root) == "" {
		root = "."
	}

	// Check if package.json exists
	if _, err := os.Stat(filepath.Join(root, "package.json")); err != nil {
		return fmt.Errorf("package.json not found at %s - run 'npm install' first", root)
	}

	// Check if node_modules exists
	if _, err := os.Stat(filepath.Join(root, "node_modules")); err != nil {
		logutil.Logf("==> Installing npm dependencies in %s...\n", root)
		installCmd := exec.Command("npm", "install")
		installCmd.Dir = root
		if out, err := installCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("npm install failed in %s: %w\nOutput: %s", root, err, string(out))
		}
	}

	var cmd *exec.Cmd
	// To avoid hanging the server, we run a one-off build even in development mode for now
	cmd = exec.Command("npm", "run", "build")
	cmd.Dir = root

	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		s := string(out)
		scanner := bufio.NewScanner(strings.NewReader(s))
		for scanner.Scan() {
			logutil.Log(scanner.Text())
		}
	}
	return err
}

// LoadViteManifest loads the Vite build manifest for asset resolution (legacy, repo root dist)
func LoadViteManifest() (ViteManifest, error) {
	return LoadViteManifestWithConfig(ViteConfig{Root: "."})
}

// LoadViteManifestWithConfig loads manifest using provided vite config
func LoadViteManifestWithConfig(cfg ViteConfig) (ViteManifest, error) {
	manifestPath := cfg.ManifestPath()
	if _, err := os.Stat(manifestPath); err != nil {
		return nil, fmt.Errorf("Vite manifest not found at %s - run 'npm run build' first", manifestPath)
	}
	file, err := os.Open(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open manifest: %w", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}
	var manifest ViteManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}
	return manifest, nil
}

// BuildWASM compiles the Go code to WebAssembly for the given example
func BuildWASM(example string) error {
	logutil.Logf("==> Building WASM binary for '%s' example...\n", example)

	// Determine the correct paths based on current working directory
	var outPath, srcPath, workDir string

	// Check if we're already in the example directory
	if _, err := os.Stat("main.go"); err == nil {
		// We're in the example directory
		outPath = "main.wasm"
		srcPath = "main.go"
		workDir = "."
	} else {
		// We're in the project root or elsewhere
		outPath = filepath.Join("examples", example, "main.wasm")
		srcPath = filepath.Join("examples", example, "main.go")
		workDir = "."

		// Check if examples directory exists relative to current dir; if not, adjust working dir up to repo root
		if _, err := os.Stat(filepath.Join("examples", example)); err != nil {
			// Set working directory to repo root (two levels up from internal/* packages)
			workDir = filepath.Join("..", "..")
			// Build paths relative to the working directory
			outPath = filepath.Join("examples", example, "main.wasm")
			srcPath = filepath.Join("examples", example, "main.go")
		}
	}

	cmd := exec.Command("go", "build", "-o", outPath, srcPath)
	//cmd := exec.Command("tinygo", "build", "-o", outPath, srcPath)
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Dir = workDir
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		s := string(out)
		// Print only interesting lines
		scanner := bufio.NewScanner(strings.NewReader(s))
		for scanner.Scan() {
			logutil.Log(scanner.Text())
		}
	}
	return err
}

// ViteManifest represents the Vite build manifest structure
type ViteManifest map[string]struct {
	File    string   `json:"file"`
	Src     string   `json:"src,omitempty"`
	IsEntry bool     `json:"isEntry,omitempty"`
	CSS     []string `json:"css,omitempty"`
	Assets  []string `json:"assets,omitempty"`
}

// ViteConfig holds configuration for integrating Vite-built assets
type ViteConfig struct {
	Enabled      bool
	Root         string
	OutDir       string
	BuildMode    string // "development" | "production"
	ManifestFile string // optional; default: manifest.json under OutDir
}

func (v ViteConfig) OutDirOrDefault() string {
	if strings.TrimSpace(v.OutDir) != "" {
		return v.OutDir
	}
	return "dist"
}

func (v ViteConfig) ManifestPath() string {
	mf := v.ManifestFile
	if strings.TrimSpace(mf) == "" {
		// Vite v6 places manifest under .vite/manifest.json by default
		mf = filepath.Join(".vite", "manifest.json")
	}
	return filepath.Join(v.Root, v.OutDirOrDefault(), mf)
}

// Options to construct a dev server
type Options struct {
	Example   string
	Addr      string
	Vite      *ViteConfig
	WatchDirs []string
}

// Server represents a development server instance
type Server struct {
	server       *http.Server
	mux          *http.ServeMux
	example      string
	addr         string
	listener     net.Listener
	viteCfg      *ViteConfig
	viteManifest ViteManifest
	viteProc     *exec.Cmd

	// Live reload
	hub         *sseHub
	watchDirs   []string
	watcher     *fsnotify.Watcher
	watchCtx    context.Context
	watchCancel context.CancelFunc
}

// NewServerWithOptions creates a new development server with options
func NewServerWithOptions(opts Options) *Server {
	addr := opts.Addr
	if addr == "" || addr == "localhost:0" {
		addr = "localhost:0"
	}
	return &Server{
		example:   opts.Example,
		addr:      addr,
		viteCfg:   opts.Vite,
		watchDirs: opts.WatchDirs,
	}
}

// NewServer creates a new development server for the given example (backward compatible)
func NewServer(example, addr string) *Server {
	return NewServerWithOptions(Options{Example: example, Addr: addr})
}

// NewServerWithVite creates a new development server with Vite asset integration (backward compatible)
func NewServerWithVite(example, addr string) *Server {
	cfg := &ViteConfig{Enabled: true, Root: ".", BuildMode: "development"}
	return NewServerWithOptions(Options{Example: example, Addr: addr, Vite: cfg})
}

// NewServerWithViteBuild creates a new development server with Vite asset integration (backward compatible)
func NewServerWithViteBuild(example, addr string) *Server {
	cfg := &ViteConfig{Enabled: true, Root: ".", BuildMode: "production"}
	return NewServerWithOptions(Options{Example: example, Addr: addr, Vite: cfg})
}

// ViteEnabled reports whether Vite integration is enabled for this server
func (s *Server) ViteEnabled() bool {
	return s.viteCfg != nil && s.viteCfg.Enabled
}

// SetWatchDirs sets the directories/files to watch for live reload (overrides defaults)
func (s *Server) SetWatchDirs(dirs ...string) {
	s.watchDirs = dirs
}

func (s *Server) computeDefaultWatchDirs() []string {
	// Prefer example dir and common project directories
	candidates := []string{
		filepath.Join("examples", s.example),
		"src",
	}
	var out []string
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			// include both dirs and files
			out = append(out, c)
			// if we are running from internal/* package directory in tests, try repo root
			if strings.HasPrefix(c, "examples"+string(os.PathSeparator)) {
				alt := filepath.Join("..", "..", c)
				if altFi, err2 := os.Stat(alt); err2 == nil && altFi.IsDir() {
					out = append(out, alt)
				}
			}
		} else {
			// also check two-levels-up variant for running from internal pkgs
			alt := filepath.Join("..", "..", c)
			if _, err2 := os.Stat(alt); err2 == nil {
				out = append(out, alt)
			}
		}
	}
	// Always try watching index.html of example if exists
	idx := filepath.Join("examples", s.example, "index.html")
	if _, err := os.Stat(idx); err == nil {
		out = append(out, idx)
	}
	return out
}

// Handle registers a handler on the server's mux
func (s *Server) Handle(pattern string, handler http.Handler) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.Handle(pattern, handler)
}

// HandleFunc registers a handler function on the server's mux
func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.Handle(pattern, http.HandlerFunc(handler))
}

// Start starts the development server
func (s *Server) Start() error {
	// Build Vite assets if enabled
	if s.ViteEnabled() {
		cfg := *s.viteCfg
		// Verify vite root
		root := cfg.Root
		if strings.TrimSpace(root) == "" {
			root = "."
			cfg.Root = root
		}
		if _, err := os.Stat(filepath.Join(root, "package.json")); err != nil {
			logutil.Logf("Warning: Vite enabled but package.json not found at %s; disabling Vite for this run\n", root)
			s.viteCfg.Enabled = false
		} else {
			// Ensure dependencies
			if _, err := os.Stat(filepath.Join(root, "node_modules")); err != nil {
				logutil.Logf("==> Installing npm dependencies in %s...\n", root)
				cmd := exec.Command("npm", "install")
				cmd.Dir = root
				if out, err := cmd.CombinedOutput(); err != nil {
					logutil.Logf("Warning: npm install failed in %s: %v\nOutput: %s\n", root, err, string(out))
				}
			}
			// Build once (non-watch) for both modes to avoid blocking
			if err := BuildViteAssetsWithConfig(cfg); err != nil {
				logutil.Logf("Warning: Vite build failed: %v\n", err)
				// Continue without Vite assets
				s.viteCfg.Enabled = false
			} else {
				if manifest, err := LoadViteManifestWithConfig(cfg); err != nil {
					logutil.Logf("Warning: Failed to load Vite manifest: %v\n", err)
				} else {
					s.viteManifest = manifest
				}
			}
		}
	}

	// Build WASM
	if err := BuildWASM(s.example); err != nil {
		return fmt.Errorf("failed to build WASM: %w", err)
	}

	// Setup HTTP handlers using a new ServeMux to avoid conflicts between test instances
	mux := http.NewServeMux()
	s.mux = mux

	// Static files from examples/<example> or current directory
	var dir string
	if _, err := os.Stat("index.html"); err == nil {
		// We're in the example directory
		dir = "."
	} else {
		// We're in the project root or elsewhere
		dir = filepath.Join("examples", s.example)
		// Check if examples directory exists
		if _, err := os.Stat(dir); err != nil {
			// Try going up levels
			dir = filepath.Join("..", "..", "examples", s.example)
		}
	}
	fs := http.FileServer(http.Dir(dir))

	// Root handler with live-reload injection for index.html and SPA fallback
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		serveIndex := false

		// If root, index.html, or ends with slash -> serve index
		if p == "/" || p == "/index.html" || strings.HasSuffix(p, "/") {
			serveIndex = true
		} else {
			// Check if requested path corresponds to an actual file; if not, fallback to index.html (SPA routing)
			// Trim leading slash and join with dir
			cleanPath := strings.TrimPrefix(p, "/")
			filePath := filepath.Join(dir, cleanPath)
			if _, err := os.Stat(filePath); err != nil {
				// If file doesn't exist, serve index.html so client-side router can handle it
				serveIndex = true
			}
		}

		if serveIndex {
			// Serve index.html with livereload injection
			indexPath := filepath.Join(dir, "index.html")
			data, err := os.ReadFile(indexPath)
			if err != nil {
				http.Error(w, "index.html not found", http.StatusNotFound)
				return
			}
			html := string(data)

			// Determine dist base path for asset resolution
			var distBase string
			if s.ViteEnabled() {
				distBase = filepath.Join(s.viteCfg.Root, s.viteCfg.OutDirOrDefault())
			} else {
				// Try common locations when Vite isn't configured
				candidates := []string{"dist", filepath.Join("..", "..", "dist")}
				for _, c := range candidates {
					if fi, err := os.Stat(c); err == nil && fi.IsDir() {
						distBase = c
						break
					}
				}
			}

			// Rewrite asset tags
			jsPlaceholder := "/assets/js/main.js"
			cssPlaceholder := "/assets/css/main.css"

			if s.ViteEnabled() && len(s.viteManifest) > 0 {
				// Replace via manifest
				jsKey := strings.TrimPrefix(jsPlaceholder, "/")
				if entry, ok := s.viteManifest[jsKey]; ok && strings.TrimSpace(entry.File) != "" {
					hashed := "/dist/" + entry.File
					html = strings.ReplaceAll(html, jsPlaceholder, hashed)
					// If there is CSS produced by this entry and the placeholder CSS isn't in manifest, map it
					cssKey := strings.TrimPrefix(cssPlaceholder, "/")
					if _, ok := s.viteManifest[cssKey]; !ok && len(entry.CSS) > 0 {
						cssHashed := "/dist/" + entry.CSS[0]
						html = strings.ReplaceAll(html, cssPlaceholder, cssHashed)
					}
				}
				// Replace CSS if it has its own manifest entry
				cssKey := strings.TrimPrefix(cssPlaceholder, "/")
				if cssEntry, ok := s.viteManifest[cssKey]; ok && strings.TrimSpace(cssEntry.File) != "" {
					html = strings.ReplaceAll(html, cssPlaceholder, "/dist/"+cssEntry.File)
				}
			} else if strings.TrimSpace(distBase) != "" {
				// Fallback: glob for hashed asset files if manifest is unavailable or Vite disabled
				if files, _ := filepath.Glob(filepath.Join(distBase, "js", "main-*.js")); len(files) > 0 {
					html = strings.ReplaceAll(html, jsPlaceholder, "/dist/js/"+filepath.Base(files[0]))
				}
				if files, _ := filepath.Glob(filepath.Join(distBase, "css", "main-*.css")); len(files) > 0 {
					html = strings.ReplaceAll(html, cssPlaceholder, "/dist/css/"+filepath.Base(files[0]))
				}
			}

			inject := "<script>(function(){try{var es=new EventSource('/__livereload');es.onmessage=function(e){if(e.data==='reload'){location.reload();}}}catch(e){console.warn('livereload disabled',e);}})();</script>"
			if strings.Contains(html, "</body>") {
				html = strings.Replace(html, "</body>", inject+"</body>", 1)
			} else {
				html = html + inject
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write([]byte(html))
			return
		}

		// Delegate other paths to static file server
		fs.ServeHTTP(w, r)
	})

	// wasm_exec.js served from embedded content
	mux.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		_, _ = w.Write(getWasmExecJS())
	})

	// main.wasm served with correct MIME type
	mux.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		wasmPath := filepath.Join(dir, "main.wasm")
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, wasmPath)
	})

	// Redirect placeholder asset paths to hashed files using Vite manifest or glob fallback
	resolveAndRedirect := func(kind string, w http.ResponseWriter, r *http.Request) {
		// Determine dist base directory
		var distBase string
		if s.ViteEnabled() {
			distBase = filepath.Join(s.viteCfg.Root, s.viteCfg.OutDirOrDefault())
		} else {
			candidates := []string{"dist", filepath.Join("..", "..", "dist")}
			for _, c := range candidates {
				if fi, err := os.Stat(c); err == nil && fi.IsDir() {
					distBase = c
					break
				}
			}
		}

		var targetURL string

		if kind == "css" {
			// First try manifest direct mapping or via JS entry's CSS list
			if s.ViteEnabled() && len(s.viteManifest) > 0 {
				if cssEntry, ok := s.viteManifest["assets/css/main.css"]; ok && strings.TrimSpace(cssEntry.File) != "" {
					targetURL = "/dist/" + strings.ReplaceAll(cssEntry.File, string(os.PathSeparator), "/")
				} else if jsEntry, ok := s.viteManifest["assets/js/main.js"]; ok && len(jsEntry.CSS) > 0 {
					targetURL = "/dist/" + strings.ReplaceAll(jsEntry.CSS[0], string(os.PathSeparator), "/")
				}
			}
			// Fallback to glob
			if targetURL == "" && strings.TrimSpace(distBase) != "" {
				if files, _ := filepath.Glob(filepath.Join(distBase, "css", "main-*.css")); len(files) > 0 {
					if rel, err := filepath.Rel(distBase, files[0]); err == nil {
						targetURL = "/dist/" + strings.ReplaceAll(rel, string(os.PathSeparator), "/")
					}
				}
			}
		} else if kind == "js" {
			if s.ViteEnabled() && len(s.viteManifest) > 0 {
				if jsEntry, ok := s.viteManifest["assets/js/main.js"]; ok && strings.TrimSpace(jsEntry.File) != "" {
					targetURL = "/dist/" + strings.ReplaceAll(jsEntry.File, string(os.PathSeparator), "/")
				}
			}
			if targetURL == "" && strings.TrimSpace(distBase) != "" {
				if files, _ := filepath.Glob(filepath.Join(distBase, "js", "main-*.js")); len(files) > 0 {
					if rel, err := filepath.Rel(distBase, files[0]); err == nil {
						targetURL = "/dist/" + strings.ReplaceAll(rel, string(os.PathSeparator), "/")
					}
				}
			}
		}

		if targetURL != "" {
			// Temporary redirect to the correct hashed asset
			http.Redirect(w, r, targetURL, http.StatusFound)
			return
		}
		// Not resolved; return 404 to avoid confusion
		http.NotFound(w, r)
	}

	mux.HandleFunc("/assets/css/main.css", func(w http.ResponseWriter, r *http.Request) {
		resolveAndRedirect("css", w, r)
	})
	mux.HandleFunc("/assets/js/main.js", func(w http.ResponseWriter, r *http.Request) {
		resolveAndRedirect("js", w, r)
	})

	// Vite asset serving (and fallback when Vite is disabled but dist exists)
	if s.ViteEnabled() {
		outDir := s.viteCfg.OutDirOrDefault()
		distDir := filepath.Join(s.viteCfg.Root, outDir)
		distFS := http.FileServer(http.Dir(distDir))
		// Mount common prefixes used by Vite outputs
		mux.Handle("/assets/", http.StripPrefix("/assets/", distFS))
		mux.Handle("/dist/", http.StripPrefix("/dist/", distFS))

		// Back-compat manifest endpoint
		mux.HandleFunc("/__vite/assets", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if len(s.viteManifest) > 0 {
				_ = json.NewEncoder(w).Encode(s.viteManifest)
			} else {
				_, _ = w.Write([]byte("{}"))
			}
		})
	} else {
		// If Vite is disabled, try serving from a discovered dist directory
		candidates := []string{"dist", filepath.Join("..", "..", "dist")}
		var distDir string
		for _, c := range candidates {
			if fi, err := os.Stat(c); err == nil && fi.IsDir() {
				distDir = c
				break
			}
		}
		if distDir != "" {
			distFS := http.FileServer(http.Dir(distDir))
			mux.Handle("/assets/", http.StripPrefix("/assets/", distFS))
			mux.Handle("/dist/", http.StripPrefix("/dist/", distFS))
		}
	}

	// Live reload SSE endpoint
	if s.hub == nil {
		s.hub = newSSEHub()
	}
	mux.HandleFunc("/__livereload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		ch := make(chan string, 8)
		s.hub.addClient(ch)
		defer s.hub.removeClient(ch)
		// initial ping
		fmt.Fprintf(w, "event: ping\n")
		fmt.Fprintf(w, "data: ok\n\n")
		flusher.Flush()
		for {
			select {
			case <-r.Context().Done():
				return
			case msg := <-ch:
				fmt.Fprintf(w, "data: %s\n\n", msg)
				flusher.Flush()
			}
		}
	})

	// Create listener to get actual port
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}
	s.listener = listener
	s.addr = listener.Addr().String()

	s.server = &http.Server{
		Handler: mux,
	}

	// Start the HTTP server
	go func() {
		logutil.Logf("==> Serving http://%s (example: %s)\n", s.addr, s.example)
		if err := s.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			logutil.Logf("Server error: %v", err)
		}
	}()

	// Start file watcher for live reload
	if err := s.startWatcher(); err != nil {
		logutil.Logf("Warning: live-reload watcher failed to start: %v\n", err)
	}

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (s *Server) startWatcher() error {
	// if no hub, nothing to do
	if s.hub == nil {
		s.hub = newSSEHub()
	}
	// Determine watch dirs
	dirs := s.watchDirs
	if len(dirs) == 0 {
		dirs = s.computeDefaultWatchDirs()
	}
	if len(dirs) == 0 {
		return nil
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	s.watcher = w

	ctx, cancel := context.WithCancel(context.Background())
	s.watchCtx = ctx
	s.watchCancel = cancel

	// helper: add dir recursively
	addRecursive := func(root string) {
		fi, err := os.Stat(root)
		if err != nil {
			return
		}
		if !fi.IsDir() {
			_ = w.Add(root)
			return
		}
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				// skip hidden dirs
				if strings.HasPrefix(info.Name(), ".") {
					return filepath.SkipDir
				}
				_ = w.Add(path)
			}
			return nil
		})
	}
	for _, d := range dirs {
		addRecursive(d)
	}

	debounce := time.NewTimer(0)
	if !debounce.Stop() {
		select {
		case <-debounce.C:
		default:
		}
	}

	rebuild := func() {
		if err := BuildWASM(s.example); err != nil {
			logutil.Logf("[dev] Build failed: %v\n", err)
			return
		}
		s.hub.broadcast("reload")
		logutil.Log("[dev] Reload signaled")
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ev := <-w.Events:
				name := strings.ToLower(ev.Name)
				if !(strings.HasSuffix(name, ".go") || strings.HasSuffix(name, ".html")) {
					continue
				}
				debounce.Reset(200 * time.Millisecond)
			case <-debounce.C:
				rebuild()
			case err := <-w.Errors:
				logutil.Logf("[dev] watcher error: %v\n", err)
			}
		}
	}()

	return nil
}

// Stop stops the development server
func (s *Server) Stop() error {
	// stop watcher first
	if s.watchCancel != nil {
		s.watchCancel()
		s.watchCancel = nil
	}
	if s.watcher != nil {
		_ = s.watcher.Close()
		s.watcher = nil
	}
	if s.server == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.server.Shutdown(ctx)
	if s.listener != nil {
		s.listener.Close()
	}
	return err
}

// URL returns the server's base URL
func (s *Server) URL() string {
	return fmt.Sprintf("http://%s", s.addr)
}
