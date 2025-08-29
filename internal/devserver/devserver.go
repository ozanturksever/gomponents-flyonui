package devserver

import (
	"bufio"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//go:embed wasm_exec-tinygo.js
var wasmExecJSTinyGo []byte

//go:embed wasm_exec.js
var wasmExecJS []byte

// getWasmExecJS returns the wasm_exec.js content from file
func getWasmExecJS() []byte {
	return wasmExecJS
}

// BuildViteAssets builds CSS and JS assets using Vite
func BuildViteAssets(mode string) error {
	log.Printf("==> Building Vite assets in %s mode...\n", mode)

	// Check if package.json exists
	if _, err := os.Stat("package.json"); err != nil {
		return fmt.Errorf("package.json not found - run 'npm install' first")
	}

	// Check if node_modules exists
	if _, err := os.Stat("node_modules"); err != nil {
		log.Println("==> Installing npm dependencies...")
		installCmd := exec.Command("npm", "install")
		if out, err := installCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("npm install failed: %w\nOutput: %s", err, string(out))
		}
	}

	var cmd *exec.Cmd
	if mode == "development" {
		// For development, we'll use Vite's build --watch mode
		cmd = exec.Command("npm", "run", "build:watch")
	} else {
		// For production builds
		cmd = exec.Command("npm", "run", "build")
	}

	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		s := string(out)
		scanner := bufio.NewScanner(strings.NewReader(s))
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}
	return err
}

// LoadViteManifest loads the Vite build manifest for asset resolution
func LoadViteManifest() (ViteManifest, error) {
	manifestPath := "dist/manifest.json"
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
	log.Printf("==> Building WASM binary for '%s' example...\n", example)

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
			log.Println(scanner.Text())
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

// Server represents a development server instance
type Server struct {
	server       *http.Server
	mux          *http.ServeMux
	example      string
	addr         string
	listener     net.Listener
	viteEnabled  bool
	viteBuild    bool
	viteManifest ViteManifest
}

// NewServer creates a new development server for the given example
// If addr is empty or "localhost:0", it will use a random available port
func NewServer(example, addr string) *Server {
	if addr == "" || addr == "localhost:0" {
		addr = "localhost:0"
	}
	return &Server{
		example: example,
		addr:    addr,
	}
}

// NewServerWithVite creates a new development server with Vite asset integration
func NewServerWithVite(example, addr string) *Server {
	server := NewServer(example, addr)
	server.viteEnabled = true
	return server
}

// NewServerWithViteBuild creates a new development server with Vite asset integration
func NewServerWithViteBuild(example, addr string) *Server {
	server := NewServer(example, addr)
	server.viteEnabled = true
	server.viteBuild = true
	return server
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
	if s.viteEnabled && s.viteBuild {
		if err := BuildViteAssets("development"); err != nil {
			log.Printf("Warning: Vite build failed: %v", err)
			// Continue without Vite assets for development
			s.viteEnabled = false
		} else {
			// Load manifest for asset resolution
			if manifest, err := LoadViteManifest(); err != nil {
				log.Printf("Warning: Failed to load Vite manifest: %v", err)
			} else {
				s.viteManifest = manifest
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

	// Vite asset serving
	if s.viteEnabled {
		// Serve built assets from dist directory
		distFS := http.FileServer(http.Dir("dist"))
		mux.Handle("/assets/", http.StripPrefix("/assets/", distFS))

		// Serve CSS and JS files with proper MIME types
		mux.HandleFunc("/dist/", func(w http.ResponseWriter, r *http.Request) {
			filePath := strings.TrimPrefix(r.URL.Path, "/dist/")
			fullPath := filepath.Join("dist", filePath)

			// Set appropriate content type
			if strings.HasSuffix(filePath, ".css") {
				w.Header().Set("Content-Type", "text/css")
			} else if strings.HasSuffix(filePath, ".js") {
				w.Header().Set("Content-Type", "application/javascript")
			} else if strings.HasSuffix(filePath, ".map") {
				w.Header().Set("Content-Type", "application/json")
			}

			http.ServeFile(w, r, fullPath)
		})

		// Asset helper endpoint for Go code to get asset URLs
		mux.HandleFunc("/__vite/assets", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if len(s.viteManifest) > 0 {
				json.NewEncoder(w).Encode(s.viteManifest)
			} else {
				w.Write([]byte("{}"))
			}
		})
	}

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

	go func() {
		log.Printf("==> Serving http://%s (example: %s)\n", s.addr, s.example)
		if err := s.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)
	return nil
}

// Stop stops the development server
func (s *Server) Stop() error {
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
