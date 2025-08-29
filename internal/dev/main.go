package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/ozanturksever/gomponents-flyonui/internal/devserver"
	"github.com/ozanturksever/gomponents-flyonui/logutil"
)

func main() {
	example := flag.String("example", "counter", "example directory under ./examples to run")
	port := flag.Int("port", 8080, "port to serve the dev server on")
	flag.Parse()

	exampleDir := filepath.Join("examples", *example)
	if info, err := os.Stat(exampleDir); err != nil || !info.IsDir() {
		logutil.Logf("example '%s' not found at %s\n", *example, exampleDir)
		os.Exit(1)
	}

	if _, err := os.Stat(filepath.Join(exampleDir, "main.go")); err != nil {
		logutil.Logf("example '%s' missing main.go\n", *example)
		os.Exit(1)
	}

	// Create and start the development server
	server := devserver.NewServer(*example, fmt.Sprintf(":%d", *port))
	if err := server.Start(); err != nil {
		logutil.Logf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
	defer server.Stop()

	// Wait for interrupt
	sigCh := make(chan os.Signal, 1)
	signalNotify(sigCh)
	<-sigCh
	logutil.Log("Shutting down...")
	server.Stop()
}

// small wrapper to register for SIGINT/SIGTERM
func signalNotify(ch chan os.Signal) {
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}
