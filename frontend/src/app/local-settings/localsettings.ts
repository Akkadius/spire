export class Setting {
  static DEBUG_MODE = "debug-mode";
  static TAB_HOVER  = "tab-hover";
}

export class LocalSettings {
  static get = name => localStorage.getItem(name);

  static set(name, key) {
    localStorage.setItem(name, key);
  }

  static isDebugEnabled() {
    return this.get(Setting.DEBUG_MODE) === "true" ? this.get(Setting.DEBUG_MODE) : false
  }

  static isTabHoverEnabled() {
    return this.get(Setting.TAB_HOVER) === "true" ? this.get(Setting.TAB_HOVER) : false
  }
}
