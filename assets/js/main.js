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

// FlyonUI Component Manager
class FlyonUIManager {
  constructor() {
    this.initialized = false;
    this.components = new Set();
  }

  init() {
    if (this.initialized) return;
    
    // Initialize FlyonUI if available
    if (typeof window.HSStaticMethods !== 'undefined') {
      window.HSStaticMethods.autoInit();
      this.initialized = true;
      console.log('FlyonUI components initialized via HSStaticMethods');
    } else {
      console.warn('HSStaticMethods not found - FlyonUI components may not be interactive');
    }
    
    this.setupComponentObserver();
  }

  reinitialize() {
    if (typeof window.HSStaticMethods !== 'undefined') {
      window.HSStaticMethods.autoInit();
      console.log('FlyonUI components reinitialized via HSStaticMethods');
    }
  }

  setupComponentObserver() {
    // Watch for dynamically added components
    const observer = new MutationObserver((mutations) => {
      let shouldReinit = false;
      
      mutations.forEach((mutation) => {
        if (mutation.type === 'childList') {
          mutation.addedNodes.forEach((node) => {
            if (node.nodeType === Node.ELEMENT_NODE) {
              // Check if added node has FlyonUI components
              if (this.hasFlyonUIComponents(node)) {
                shouldReinit = true;
              }
            }
          });
        }
      });
      
      if (shouldReinit) {
        // Debounce reinitalization
        clearTimeout(this.reinitTimeout);
        this.reinitTimeout = setTimeout(() => {
          this.reinitialize();
        }, 100);
      }
    });
    
    observer.observe(document.body, {
      childList: true,
      subtree: true
    });
  }

  hasFlyonUIComponents(element) {
    // Check for common FlyonUI component classes
    const flyonUIClasses = [
      'dropdown', 'modal', 'alert', 'accordion', 'collapse',
      'drawer', 'tabs', 'tooltip', 'datepicker', 'combobox'
    ];
    
    return flyonUIClasses.some(className => 
      element.classList?.contains(className) || 
      element.querySelector?.(`.${className}`)
    );
  }
}

// Theme management utilities
class ThemeManager {
  constructor() {
    this.storageKey = 'flyonui-theme';
    this.defaultTheme = 'light';
    this.init();
  }

  init() {
    // Load saved theme or detect system preference
    const savedTheme = localStorage.getItem(this.storageKey);
    const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    const theme = savedTheme || systemTheme;
    
    this.setTheme(theme);
    this.setupSystemThemeListener();
  }

  setTheme(theme) {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem(this.storageKey, theme);
    
    // Dispatch custom event for Go-WASM to listen to
    window.dispatchEvent(new CustomEvent('themeChanged', {
      detail: { theme }
    }));
  }

  getTheme() {
    return document.documentElement.getAttribute('data-theme') || this.defaultTheme;
  }

  toggleTheme() {
    const currentTheme = this.getTheme();
    const newTheme = currentTheme === 'light' ? 'dark' : 'light';
    this.setTheme(newTheme);
    return newTheme;
  }

  setupSystemThemeListener() {
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
      if (!localStorage.getItem(this.storageKey)) {
        this.setTheme(e.matches ? 'dark' : 'light');
      }
    });
  }
}

// WASM loading utilities
class WASMLoader {
  constructor() {
    this.loadingClass = 'wasm-loading';
    this.readyClass = 'wasm-ready';
    this.init();
  }

  init() {
    // Add loading state to body
    document.body.classList.add(this.loadingClass);
    
    // Listen for WASM ready event
    window.addEventListener('wasmReady', () => {
      this.onWASMReady();
    });
    
    // Fallback timeout
    setTimeout(() => {
      if (document.body.classList.contains(this.loadingClass)) {
        console.warn('WASM loading timeout - removing loading state');
        this.onWASMReady();
      }
    }, 10000);
  }

  onWASMReady() {
    document.body.classList.remove(this.loadingClass);
    document.body.classList.add(this.readyClass);
    
    // Dispatch ready event for components
    window.dispatchEvent(new CustomEvent('appReady'));
  }

  showLoadingSpinner(element) {
    if (element) {
      element.classList.add(this.loadingClass);
    }
  }

  hideLoadingSpinner(element) {
    if (element) {
      element.classList.remove(this.loadingClass);
      element.classList.add(this.readyClass);
    }
  }
}

// Accessibility enhancements
class AccessibilityManager {
  constructor() {
    this.init();
  }

  init() {
    this.setupKeyboardNavigation();
    this.setupFocusManagement();
    this.setupReducedMotion();
  }

  setupKeyboardNavigation() {
    // Enhanced keyboard navigation for dropdowns and modals
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape') {
        // Close open dropdowns and modals
        const openDropdowns = document.querySelectorAll('.dropdown.dropdown-open');
        openDropdowns.forEach(dropdown => {
          dropdown.classList.remove('dropdown-open');
        });
        
        const openModals = document.querySelectorAll('.modal.modal-open');
        openModals.forEach(modal => {
          modal.classList.remove('modal-open');
        });
      }
    });
  }

  setupFocusManagement() {
    // Trap focus within modals
    document.addEventListener('focusin', (e) => {
      const modal = e.target.closest('.modal.modal-open');
      if (modal) {
        const focusableElements = modal.querySelectorAll(
          'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
        );
        
        if (focusableElements.length > 0) {
          const firstElement = focusableElements[0];
          const lastElement = focusableElements[focusableElements.length - 1];
          
          if (e.target === modal && firstElement) {
            firstElement.focus();
          }
        }
      }
    });
  }

  setupReducedMotion() {
    // Respect user's motion preferences
    const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)');
    
    const handleMotionPreference = (mediaQuery) => {
      if (mediaQuery.matches) {
        document.documentElement.style.setProperty('--animation-duration', '0.01ms');
        document.documentElement.style.setProperty('--transition-duration', '0.01ms');
      } else {
        document.documentElement.style.removeProperty('--animation-duration');
        document.documentElement.style.removeProperty('--transition-duration');
      }
    };
    
    handleMotionPreference(prefersReducedMotion);
    prefersReducedMotion.addEventListener('change', handleMotionPreference);
  }
}

// Performance monitoring
class PerformanceMonitor {
  constructor() {
    this.metrics = {};
    this.init();
  }

  init() {
    // Monitor Core Web Vitals
    this.observeLCP();
    this.observeFID();
    this.observeCLS();
  }

  observeLCP() {
    // Largest Contentful Paint
    if ('PerformanceObserver' in window) {
      const observer = new PerformanceObserver((list) => {
        const entries = list.getEntries();
        const lastEntry = entries[entries.length - 1];
        this.metrics.lcp = lastEntry.startTime;
      });
      
      observer.observe({ entryTypes: ['largest-contentful-paint'] });
    }
  }

  observeFID() {
    // First Input Delay
    if ('PerformanceObserver' in window) {
      const observer = new PerformanceObserver((list) => {
        const entries = list.getEntries();
        entries.forEach((entry) => {
          this.metrics.fid = entry.processingStart - entry.startTime;
        });
      });
      
      observer.observe({ entryTypes: ['first-input'] });
    }
  }

  observeCLS() {
    // Cumulative Layout Shift
    if ('PerformanceObserver' in window) {
      let clsValue = 0;
      const observer = new PerformanceObserver((list) => {
        const entries = list.getEntries();
        entries.forEach((entry) => {
          if (!entry.hadRecentInput) {
            clsValue += entry.value;
            this.metrics.cls = clsValue;
          }
        });
      });
      
      observer.observe({ entryTypes: ['layout-shift'] });
    }
  }

  getMetrics() {
    return { ...this.metrics };
  }
}

// Utility functions for Go-WASM integration
const GoWASMUtils = {
  // Safe way to call Go functions from JavaScript
  callGoFunction(funcName, ...args) {
    if (window.go && window.go.exports && window.go.exports[funcName]) {
      try {
        return window.go.exports[funcName](...args);
      } catch (error) {
        console.error(`Error calling Go function ${funcName}:`, error);
        return null;
      }
    } else {
      console.warn(`Go function ${funcName} not available`);
      return null;
    }
  },

  // Register JavaScript functions for Go to call
  registerForGo(name, func) {
    if (!window.goCallbacks) {
      window.goCallbacks = {};
    }
    window.goCallbacks[name] = func;
  },

  // Emit events that Go can listen to
  emitToGo(eventName, data) {
    window.dispatchEvent(new CustomEvent(`go:${eventName}`, {
      detail: data
    }));
  }
};

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

// Export for module systems
if (typeof module !== 'undefined' && module.exports) {
  module.exports = {
    FlyonUIManager,
    ThemeManager,
    WASMLoader,
    AccessibilityManager,
    PerformanceMonitor,
    GoWASMUtils
  };
}