/**
 * Vite Entry Point for gomponents-flyonui
 * 
 * This is the main entry point that Vite uses to bundle all assets.
 * It imports CSS and JavaScript modules for processing and optimization.
 */

// Import main CSS file (includes TailwindCSS and FlyonUI)
import './css/main.css';

// Import main JavaScript utilities
import './js/main.js';

// Hot Module Replacement (HMR) support for development
if (import.meta.hot) {
  // Accept HMR updates for CSS
  import.meta.hot.accept('./css/main.css', () => {
    console.log('CSS updated via HMR');
  });
  
  // Accept HMR updates for JavaScript
  import.meta.hot.accept('./js/main.js', () => {
    console.log('JavaScript updated via HMR');
    // Re-initialize if needed
    if (window.initializeApp) {
      window.initializeApp();
    }
  });
  
  // Custom HMR handling for theme changes
  import.meta.hot.on('theme-update', (data) => {
    if (window.themeManager) {
      window.themeManager.setTheme(data.theme);
    }
  });
}

// Development mode utilities
if (import.meta.env.DEV) {
  console.log('gomponents-flyonui assets loaded in development mode');
  
  // Add development helpers to window for debugging
  window.__FLYONUI_DEV__ = {
    version: import.meta.env.VITE_APP_VERSION || 'dev',
    mode: import.meta.env.MODE,
    env: import.meta.env,
    
    // Utility to reload CSS
    reloadCSS() {
      const links = document.querySelectorAll('link[rel="stylesheet"]');
      links.forEach(link => {
        const href = link.href;
        link.href = href.split('?')[0] + '?t=' + Date.now();
      });
    },
    
    // Utility to check theme
    checkTheme() {
      return {
        current: window.themeManager?.getTheme(),
        available: ['light', 'dark', 'cupcake', 'bumblebee', 'emerald', 'corporate', 'synthwave', 'retro', 'cyberpunk', 'valentine', 'halloween', 'garden', 'forest', 'aqua', 'lofi', 'pastel', 'fantasy', 'wireframe', 'black', 'luxury', 'dracula', 'cmyk', 'autumn', 'business', 'acid', 'lemonade', 'night', 'coffee', 'winter', 'dim', 'nord', 'sunset']
      };
    },
    
    // Utility to get performance metrics
    getMetrics() {
      return window.performanceMonitor?.getMetrics() || {};
    }
  };
}

// Production mode optimizations
if (import.meta.env.PROD) {
  console.log('gomponents-flyonui assets loaded in production mode');
}

// Export for potential module usage
export default {
  version: import.meta.env.VITE_APP_VERSION || '1.0.0',
  mode: import.meta.env.MODE,
  isDev: import.meta.env.DEV,
  isProd: import.meta.env.PROD
};