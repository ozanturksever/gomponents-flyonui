package vite

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// AssetEntry represents a Vite manifest entry
type AssetEntry struct {
	File    string   `json:"file"`
	Src     string   `json:"src,omitempty"`
	IsEntry bool     `json:"isEntry,omitempty"`
	CSS     []string `json:"css,omitempty"`
	Assets  []string `json:"assets,omitempty"`
}

// Manifest represents the Vite build manifest
type Manifest map[string]AssetEntry

// AssetResolver handles Vite asset resolution
type AssetResolver struct {
	manifest    Manifest
	baseURL     string
	development bool
}

// NewAssetResolver creates a new asset resolver
func NewAssetResolver(baseURL string, development bool) *AssetResolver {
	return &AssetResolver{
		baseURL:     strings.TrimSuffix(baseURL, "/"),
		development: development,
	}
}

// LoadManifest loads the Vite manifest from the server
func (ar *AssetResolver) LoadManifest() error {
	if ar.development {
		// In development, try to fetch manifest from dev server
		resp, err := http.Get(ar.baseURL + "/__vite/assets")
		if err != nil {
			return fmt.Errorf("failed to fetch manifest: %w", err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&ar.manifest); err != nil {
			return fmt.Errorf("failed to decode manifest: %w", err)
		}
	}
	return nil
}

// GetAssetURL returns the URL for a given asset
func (ar *AssetResolver) GetAssetURL(entryPoint string) string {
	if ar.development {
		// In development, return direct paths
		switch entryPoint {
		case "main.css":
			return ar.baseURL + "/dist/assets/main.css"
		case "main.js":
			return ar.baseURL + "/dist/assets/main.js"
		default:
			return ar.baseURL + "/dist/assets/" + entryPoint
		}
	}

	// In production, use manifest
	if entry, exists := ar.manifest[entryPoint]; exists {
		return ar.baseURL + "/" + entry.File
	}
	return ar.baseURL + "/assets/" + entryPoint
}

// GetCSSLinks returns link elements for CSS assets
func (ar *AssetResolver) GetCSSLinks(entryPoint string) []g.Node {
	var links []g.Node

	if ar.development {
		// In development, include main CSS
		links = append(links, h.Link(
			h.Rel("stylesheet"),
			h.Href(ar.GetAssetURL("main.css")),
		))
	} else {
		// In production, use manifest to get all CSS files
		if entry, exists := ar.manifest[entryPoint]; exists {
			for _, cssFile := range entry.CSS {
				links = append(links, h.Link(
					h.Rel("stylesheet"),
					h.Href(ar.baseURL+"/"+cssFile),
				))
			}
		}
	}

	return links
}

// GetJSScript returns script element for JS assets
func (ar *AssetResolver) GetJSScript(entryPoint string) g.Node {
	if ar.development {
		// In development, include main JS
		return h.Script(
			h.Type("module"),
			h.Src(ar.GetAssetURL("main.js")),
		)
	}

	// In production, use manifest
	if entry, exists := ar.manifest[entryPoint]; exists {
		return h.Script(
			h.Type("module"),
			h.Src(ar.baseURL+"/"+entry.File),
		)
	}

	return h.Script(
		h.Type("module"),
		h.Src(ar.GetAssetURL(entryPoint)),
	)
}

// GetAssetTags returns both CSS and JS tags for an entry point
func (ar *AssetResolver) GetAssetTags(entryPoint string) []g.Node {
	var tags []g.Node

	// Add CSS links
	tags = append(tags, ar.GetCSSLinks(entryPoint)...)

	// Add JS script
	tags = append(tags, ar.GetJSScript(entryPoint))

	return tags
}

// DefaultAssetTags returns the standard asset tags for the main entry point
func (ar *AssetResolver) DefaultAssetTags() []g.Node {
	return ar.GetAssetTags("assets/main.js")
}