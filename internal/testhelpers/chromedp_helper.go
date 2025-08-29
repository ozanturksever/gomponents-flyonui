//go:build !js && !wasm

package testhelpers

import (
	"context"
	"errors"
	"time"

	"github.com/chromedp/chromedp"
)

// ChromedpConfig holds configuration options for chromedp browser setup
type ChromedpConfig struct {
	// Headless determines if the browser runs in headless mode
	Headless bool
	// Timeout sets the context timeout for the entire test
	Timeout time.Duration
	// DisableGPU disables GPU acceleration
	DisableGPU bool
	// NoSandbox disables the sandbox
	NoSandbox bool
	// DisableDevShmUsage disables /dev/shm usage
	DisableDevShmUsage bool
	// AdditionalFlags allows adding custom Chrome flags
	AdditionalFlags []chromedp.ExecAllocatorOption
}

// DefaultConfig returns a sensible default configuration for chromedp tests
func DefaultConfig() ChromedpConfig {
	return ChromedpConfig{
		Headless:           true,
		Timeout:            30 * time.Second,
		DisableGPU:         true,
		NoSandbox:          true,
		DisableDevShmUsage: true,
	}
}

// VisibleConfig returns a configuration for visible browser testing (useful for debugging)
func VisibleConfig() ChromedpConfig {
	return ChromedpConfig{
		Headless:           false,
		Timeout:            30 * time.Second,
		DisableGPU:         false,
		NoSandbox:          true,
		DisableDevShmUsage: true,
		AdditionalFlags: []chromedp.ExecAllocatorOption{
			chromedp.Flag("auto-open-devtools-for-tabs", true), // Dev tools'u otomatik aç
			chromedp.Flag("disable-extensions", true),          // Extension'ları devre dışı bırak
			chromedp.Flag("disable-default-apps", true),
		},
	}
}

// ExtendedTimeoutConfig returns a configuration with longer timeout for complex tests
func ExtendedTimeoutConfig() ChromedpConfig {
	config := DefaultConfig()
	config.Timeout = 60 * time.Second
	return config
}

// DevToolsConfig returns a configuration with developer tools open
func DevToolsConfig() ChromedpConfig {
	return ChromedpConfig{
		Headless:           false, // Dev tools headless modda çalışmaz
		Timeout:            30 * time.Second,
		DisableGPU:         false,
		NoSandbox:          true,
		DisableDevShmUsage: true,
		AdditionalFlags: []chromedp.ExecAllocatorOption{
			chromedp.Flag("auto-open-devtools-for-tabs", true), // Dev tools'u otomatik aç
		},
	}
}

// DevToolsWithConsoleConfig returns a configuration with developer tools open and console tab selected
func DevToolsWithConsoleConfig() ChromedpConfig {
	return ChromedpConfig{
		Headless:           false,
		Timeout:            30 * time.Second,
		DisableGPU:         false,
		NoSandbox:          true,
		DisableDevShmUsage: true,
		AdditionalFlags: []chromedp.ExecAllocatorOption{
			chromedp.Flag("auto-open-devtools-for-tabs", true),
			// Console panelini açık tutmak için ek flag'ler
			chromedp.Flag("disable-extensions", true), // Extension'ları devre dışı bırak
			chromedp.Flag("disable-default-apps", true),
		},
	}
}

// ChromedpTestContext holds the context and cancel function for a chromedp test
type ChromedpTestContext struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

// NewChromedpContext creates a new chromedp context with the given configuration
// Returns a ChromedpTestContext that should be cleaned up with defer ctx.Cancel()
func NewChromedpContext(config ChromedpConfig) (*ChromedpTestContext, error) {
	// Create base context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)

	// Build Chrome options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", config.Headless),
		chromedp.Flag("disable-gpu", config.DisableGPU),
		chromedp.Flag("no-sandbox", config.NoSandbox),
	)

	// Add disable-dev-shm-usage if requested
	if config.DisableDevShmUsage {
		opts = append(opts, chromedp.Flag("disable-dev-shm-usage", true))
	}

	// Add any additional flags
	opts = append(opts, config.AdditionalFlags...)

	// Create allocator context
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)

	// Create browser context
	browserCtx, browserCancel := chromedp.NewContext(allocCtx)

	// Create a combined cancel function that cleans up all contexts
	combinedCancel := func() {
		browserCancel()
		allocCancel()
		cancel()
	}

	return &ChromedpTestContext{
		Ctx:    browserCtx,
		Cancel: combinedCancel,
	}, nil
}

// MustNewChromedpContext is like NewChromedpContext but panics on error
// Useful for test setup where you want to fail fast
func MustNewChromedpContext(config ChromedpConfig) *ChromedpTestContext {
	ctx, err := NewChromedpContext(config)
	if err != nil {
		panic(err)
	}
	return ctx
}

// CommonTestActions provides common chromedp actions used across tests
type CommonTestActions struct{}

// WaitForWASMInit waits for WASM to initialize by waiting for the status element to show 'Ready'
func (CommonTestActions) WaitForWASMInit(selector string, timeout time.Duration) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		<-time.After(10 * time.Second)
		// First wait for the element to be visible
		if err := chromedp.WaitVisible(selector, chromedp.ByQuery).Do(ctx); err != nil {
			return err
		}

		// Create a context with timeout for polling
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// Poll for the text to change to 'Ready'
		for {
			select {
			case <-ctxWithTimeout.Done():
				return errors.New("timeout waiting for WASM to be ready")
			default:
				var text string
				if err := chromedp.Text(selector, &text, chromedp.ByQuery).Do(ctx); err != nil {
					return err
				}
				if text == "Ready" {
					return nil
				}
				// Wait a bit before checking again
				time.Sleep(100 * time.Millisecond)
			}
		}
	})
}

// NavigateAndWaitForLoad navigates to a URL and waits for the page to load
func (CommonTestActions) NavigateAndWaitForLoad(url, waitSelector string) chromedp.Action {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(waitSelector, chromedp.ByQuery),
		chromedp.Sleep(1 * time.Second), // Basic WASM init time
	}
}

// ClickAndWait clicks an element and waits for a specified duration
func (CommonTestActions) ClickAndWait(selector string, wait time.Duration) chromedp.Action {
	return chromedp.Tasks{
		chromedp.Click(selector, chromedp.ByQuery),
		chromedp.Sleep(wait),
	}
}

// SendKeysAndWait sends keys to an element and waits for a specified duration
func (CommonTestActions) SendKeysAndWait(selector, text string, wait time.Duration) chromedp.Action {
	return chromedp.Tasks{
		chromedp.SendKeys(selector, text, chromedp.ByQuery),
		chromedp.Sleep(wait),
	}
}

// OpenDevToolsConsole opens the developer tools and focuses on the console tab
func (CommonTestActions) OpenDevToolsConsole() chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		// F12 tuşuna basarak dev tools'u aç
		if err := chromedp.KeyEvent(`F12`).Do(ctx); err != nil {
			return err
		}

		// Biraz bekle dev tools'un açılması için
		time.Sleep(1 * time.Second)

		// Console tab'ına geçmek için kısayol tuşu (Ctrl+`)
		return chromedp.KeyEvent("ctrl+`").Do(ctx)
	})
}

// ExecuteConsoleCommand executes a command in the browser console
func (CommonTestActions) ExecuteConsoleCommand(command string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		// JavaScript komutunu çalıştır
		var result interface{}
		return chromedp.Evaluate(command, &result).Do(ctx)
	})
}

// WaitForConsoleLog waits for a specific console log message
func (CommonTestActions) WaitForConsoleLog(expectedMessage string, timeout time.Duration) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// Console logları dinlemek için JavaScript kodu
		script := `
			window.testConsoleMessages = window.testConsoleMessages || [];
			if (!window.testConsoleLogIntercepted) {
				const originalLog = console.log;
				console.log = function(...args) {
					window.testConsoleMessages.push(args.join(' '));
					originalLog.apply(console, arguments);
				};
				window.testConsoleLogIntercepted = true;
			}
		`

		if err := chromedp.Evaluate(script, nil).Do(ctx); err != nil {
			return err
		}

		// Beklenen mesajı kontrol et
		for {
			select {
			case <-ctxWithTimeout.Done():
				return errors.New("timeout waiting for console log: " + expectedMessage)
			default:
				var messages []string
				if err := chromedp.Evaluate(`window.testConsoleMessages || []`, &messages).Do(ctx); err != nil {
					return err
				}

				for _, msg := range messages {
					if msg == expectedMessage {
						return nil
					}
				}

				time.Sleep(100 * time.Millisecond)
			}
		}
	})
}

// Global instance for easy access to common actions
var Actions = CommonTestActions{}
