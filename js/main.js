/**
 * Main JavaScript Entry Point for gomponents-flyonui
 * 
 * This file provides client-side utilities and enhancements for Go WebAssembly applications
 * using FlyonUI and TailwindCSS. It handles theme management, accessibility, and performance optimizations.
 */

// Import CSS for Vite processing
import '../css/main.css';

// Import FlyonUI JavaScript components
import 'flyonui/flyonui.js';
import { FlyonUIManager, ThemeManager, WASMLoader, AccessibilityManager, PerformanceMonitor, GoWASMUtils } from './lib.js';

// Initialize everything when DOM is ready
function initializeApp() {
    // Initialize managers
    window.flyonUIManager = new FlyonUIManager();
    window.themeManager = new ThemeManager();
    window.wasmLoader = new WASMLoader();
    window.accessibilityManager = new AccessibilityManager();
    window.performanceMonitor = new PerformanceMonitor();

    // Initialize FlyonUI components
    window.flyonUIManager.init();

    // Make utilities available globally
    window.GoWASMUtils = GoWASMUtils;

    // Register common callbacks for Go
    GoWASMUtils.registerForGo('toggleTheme', () => {
        return window.themeManager.toggleTheme();
    });

    GoWASMUtils.registerForGo('getTheme', () => {
        return window.themeManager.getTheme();
    });

    GoWASMUtils.registerForGo('setTheme', (theme) => {
        window.themeManager.setTheme(theme);
    });

    GoWASMUtils.registerForGo('getPerformanceMetrics', () => {
        return window.performanceMonitor.getMetrics();
    });

    // Register FlyonUI callbacks for Go
    GoWASMUtils.registerForGo('reinitializeFlyonUI', () => {
        window.flyonUIManager.reinitialize();
    });

    GoWASMUtils.registerForGo('isFlyonUIInitialized', () => {
        return window.flyonUIManager.initialized;
    });

    // Dispatch initialization complete event
    window.dispatchEvent(new CustomEvent('jsReady'));

    console.log('gomponents-flyonui JavaScript initialized');
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', initializeApp);
} else {
  initializeApp();
}

