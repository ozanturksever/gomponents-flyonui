// TypeScript type definitions for gomponents-flyonui

export interface GoCallbackMap {
  [name: string]: (...args: any[]) => any;
}

export interface PerformanceMetrics {
  lcp?: number;
  fid?: number;
  cls?: number;
  [key: string]: number | undefined;
}

export class FlyonUIManager {
  initialized: boolean;
  components: Set<any>;
  constructor();
  init(): void;
  reinitialize(): void;
  protected setupComponentObserver(): void;
  protected hasFlyonUIComponents(element: Element): boolean;
}

export class ThemeManager {
  storageKey: string;
  defaultTheme: string;
  constructor();
  init(): void;
  setTheme(theme: string): void;
  getTheme(): string;
  toggleTheme(): string;
  protected setupSystemThemeListener(): void;
}

export class WASMLoader {
  loadingClass: string;
  readyClass: string;
  constructor();
  init(): void;
  onWASMReady(): void;
  showLoadingSpinner(element?: Element | null): void;
  hideLoadingSpinner(element?: Element | null): void;
}

export class AccessibilityManager {
  constructor();
  init(): void;
  protected setupKeyboardNavigation(): void;
  protected setupFocusManagement(): void;
  protected setupReducedMotion(): void;
}

export class PerformanceMonitor {
  metrics: PerformanceMetrics;
  constructor();
  init(): void;
  protected observeLCP(): void;
  protected observeFID(): void;
  protected observeCLS(): void;
  getMetrics(): PerformanceMetrics;
}

export interface GoWASMUtilsType {
  callGoFunction(funcName: string, ...args: any[]): any;
  registerForGo(name: string, func: (...args: any[]) => any): void;
  emitToGo(eventName: string, data?: any): void;
}

export const GoWASMUtils: GoWASMUtilsType;

export interface FlyonUIDevInfo {
  version: string;
  mode?: string;
  env?: any;
  reloadCSS(): void;
  checkTheme(): { current?: string; available: string[] };
  getMetrics(): PerformanceMetrics;
}

declare global {
  interface Window {
    flyonUIManager?: FlyonUIManager;
    themeManager?: ThemeManager;
    wasmLoader?: WASMLoader;
    accessibilityManager?: AccessibilityManager;
    performanceMonitor?: PerformanceMonitor;
    GoWASMUtils?: GoWASMUtilsType;
    __FLYONUI_DEV__?: FlyonUIDevInfo;
    go?: any;
    goCallbacks?: GoCallbackMap;
  }

  interface DocumentEventMap {
    wasmReady: CustomEvent<any>;
    jsReady: CustomEvent<any>;
    themeChanged: CustomEvent<{ theme: string }>;
    appReady: CustomEvent<any>;
    [key: `go:${string}`]: CustomEvent<any>;
  }
}

// UMD global
export as namespace GomponentsFlyonUI;
