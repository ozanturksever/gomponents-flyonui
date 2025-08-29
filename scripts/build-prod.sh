#!/bin/bash

# Production Build Script for Go-WASM + Vite Assets
# This script creates an optimized production build with asset bundling

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BUILD_DIR="build"
DIST_DIR="dist"
EXAMPLE="${1:-counter}"
OUTPUT_DIR="${BUILD_DIR}/${EXAMPLE}"

# Helper functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Cleanup function
cleanup() {
    log_info "Cleaning up temporary files..."
    # Remove any temporary build artifacts if needed
}

# Set up cleanup trap
trap cleanup EXIT

# Main build function
main() {
    log_info "Starting production build for example: ${EXAMPLE}"
    
    # Validate example exists
    if [ ! -d "examples/${EXAMPLE}" ]; then
        log_error "Example '${EXAMPLE}' not found in examples/ directory"
        exit 1
    fi
    
    # Create build directory
    log_info "Creating build directory: ${OUTPUT_DIR}"
    rm -rf "${OUTPUT_DIR}"
    mkdir -p "${OUTPUT_DIR}"
    
    # Step 1: Build Vite assets for production
    log_info "Building optimized CSS and JS assets..."
    if ! NODE_ENV=production npm run build; then
        log_error "Failed to build Vite assets"
        exit 1
    fi
    
    # Verify dist directory was created
    if [ ! -d "${DIST_DIR}" ]; then
        log_error "Vite build did not create dist directory"
        exit 1
    fi
    
    # Step 2: Build WASM binary
    log_info "Compiling Go to WebAssembly..."
    if ! GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o "${OUTPUT_DIR}/main.wasm" "examples/${EXAMPLE}/main.go"; then
        log_error "Failed to build WASM binary"
        exit 1
    fi
    
    # Step 3: Copy wasm_exec.js
    log_info "Copying wasm_exec.js..."
    if [ -f "internal/devserver/wasm_exec.js" ]; then
        cp "internal/devserver/wasm_exec.js" "${OUTPUT_DIR}/"
    else
        # Fallback to Go installation
        WASM_EXEC_JS="$(go env GOROOT)/misc/wasm/wasm_exec.js"
        if [ -f "${WASM_EXEC_JS}" ]; then
            cp "${WASM_EXEC_JS}" "${OUTPUT_DIR}/"
        else
            log_error "Could not find wasm_exec.js"
            exit 1
        fi
    fi
    
    # Step 4: Copy and optimize assets
    log_info "Copying optimized assets..."
    cp -r "${DIST_DIR}/" "${OUTPUT_DIR}/dist/"
    
    # Step 5: Generate optimized index.html
    log_info "Generating production index.html..."
    generate_production_html
    
    # Step 6: Create deployment package
    log_info "Creating deployment package..."
    create_deployment_package
    
    # Step 7: Generate build report
    generate_build_report
    
    log_success "Production build completed successfully!"
    log_info "Build output: ${OUTPUT_DIR}"
    log_info "Deployment package: ${OUTPUT_DIR}.tar.gz"
}

# Generate optimized production HTML
generate_production_html() {
    local html_file="${OUTPUT_DIR}/index.html"
    local manifest_file="${OUTPUT_DIR}/dist/manifest.json"
    
    # Read the original index.html if it exists
    local original_html="examples/${EXAMPLE}/index.html"
    if [ ! -f "${original_html}" ]; then
        log_warning "No index.html found for ${EXAMPLE}, creating minimal template"
        create_minimal_html_template
        return
    fi
    
    # Copy original and inject optimized assets
    cp "${original_html}" "${html_file}"
    
    # If manifest exists, use it to inject proper asset URLs
    if [ -f "${manifest_file}" ]; then
        log_info "Injecting optimized asset URLs from manifest"
        inject_assets_from_manifest "${html_file}" "${manifest_file}"
    else
        log_warning "No Vite manifest found, using fallback asset injection"
        inject_fallback_assets "${html_file}"
    fi
}

# Create minimal HTML template
create_minimal_html_template() {
    cat > "${OUTPUT_DIR}/index.html" << 'EOF'
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go-WASM Application</title>
    <link rel="stylesheet" href="dist/assets/main.css">
</head>
<body>
    <div id="app"></div>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
            .then((result) => {
                go.run(result.instance);
            })
            .catch((err) => {
                console.error("Failed to load WASM:", err);
            });
    </script>
    <script type="module" src="dist/assets/main.js"></script>
</body>
</html>
EOF
}

# Inject assets from Vite manifest
inject_assets_from_manifest() {
    local html_file="$1"
    local manifest_file="$2"
    
    # This is a simplified version - in a real implementation,
    # you'd parse the JSON manifest and inject the correct asset URLs
    log_info "Asset injection from manifest not yet implemented - using fallback"
    inject_fallback_assets "${html_file}"
}

# Inject fallback assets
inject_fallback_assets() {
    local html_file="$1"
    
    # Remove CDN links and replace with local assets
    sed -i.bak 's|https://cdn.jsdelivr.net/npm/tailwindcss@[^"]*|dist/assets/main.css|g' "${html_file}"
    sed -i.bak 's|https://cdn.jsdelivr.net/npm/flyonui@[^"]*|dist/assets/main.css|g' "${html_file}"
    
    # Add main.js if not present
    if ! grep -q "dist/assets/main.js" "${html_file}"; then
        sed -i.bak 's|</body>|    <script type="module" src="dist/assets/main.js"></script>\n</body>|' "${html_file}"
    fi
    
    # Remove backup file
    rm -f "${html_file}.bak"
}

# Create deployment package
create_deployment_package() {
    log_info "Creating compressed deployment package..."
    
    # Create tarball
    tar -czf "${OUTPUT_DIR}.tar.gz" -C "${BUILD_DIR}" "${EXAMPLE}"
    
    # Create deployment instructions
    cat > "${OUTPUT_DIR}/DEPLOY.md" << EOF
# Deployment Instructions

## Files in this package:
- \`index.html\`: Main application entry point
- \`main.wasm\`: Compiled Go WebAssembly binary
- \`wasm_exec.js\`: Go WebAssembly runtime
- \`dist/\`: Optimized CSS and JS assets

## Deployment:
1. Upload all files to your web server
2. Ensure your server serves \`.wasm\` files with \`application/wasm\` MIME type
3. Ensure your server serves \`.js\` files with \`application/javascript\` MIME type
4. For SPA routing, configure your server to serve \`index.html\` for all routes

## Server Configuration Examples:

### Nginx:
\`\`\`nginx
location ~* \\.wasm$ {
    add_header Content-Type application/wasm;
}
\`\`\`

### Apache (.htaccess):
\`\`\`apache
AddType application/wasm .wasm
\`\`\`

### Go (net/http):
\`\`\`go
http.Handle("/", http.FileServer(http.Dir("./")))
\`\`\`
EOF
}

# Generate build report
generate_build_report() {
    local report_file="${OUTPUT_DIR}/BUILD_REPORT.md"
    
    log_info "Generating build report..."
    
    cat > "${report_file}" << EOF
# Build Report

**Example:** ${EXAMPLE}  
**Build Time:** $(date)  
**Build Directory:** ${OUTPUT_DIR}  

## File Sizes:
EOF
    
    # Add file sizes
    echo "" >> "${report_file}"
    echo "| File | Size |" >> "${report_file}"
    echo "|------|------|" >> "${report_file}"
    
    find "${OUTPUT_DIR}" -type f -exec ls -lh {} \; | while read -r line; do
        size=$(echo "$line" | awk '{print $5}')
        file=$(echo "$line" | awk '{print $9}' | sed "s|${OUTPUT_DIR}/||")
        echo "| \`${file}\` | ${size} |" >> "${report_file}"
    done
    
    # Add total size
    total_size=$(du -sh "${OUTPUT_DIR}" | cut -f1)
    echo "" >> "${report_file}"
    echo "**Total Size:** ${total_size}" >> "${report_file}"
    
    log_success "Build report generated: ${report_file}"
}

# Show usage
show_usage() {
    echo "Usage: $0 [example_name]"
    echo ""
    echo "Examples:"
    echo "  $0                    # Build 'counter' example (default)"
    echo "  $0 todo              # Build 'todo' example"
    echo "  $0 flyonui_demo      # Build 'flyonui_demo' example"
    echo ""
    echo "Available examples:"
    if [ -d "examples" ]; then
        find examples -mindepth 1 -maxdepth 1 -type d | sed 's|examples/|  - |' | sort
    else
        echo "  No examples directory found"
    fi
}

# Handle help flag
if [ "${1:-}" = "-h" ] || [ "${1:-}" = "--help" ]; then
    show_usage
    exit 0
fi

# Run main function
main "$@"