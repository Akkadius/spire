export class Setting {
  static DEBUG_MODE = "debug-mode";
  static TAB_HOVER  = "tab-hover";
  static ALPHA_ENABLED  = "alpha-enabled";
  static SPELL_LEGACY_ICONS  = "spell-legacy-icons";
  static DEFAULT_LANGUAGE_PREFERENCE  = "default-language-preference";
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

  static isAlphaEnabled() {
    return this.get(Setting.ALPHA_ENABLED) === "true" ? this.get(Setting.ALPHA_ENABLED) : false
  }

  static isSpellLegacyIconsEnabled() {
    return this.get(Setting.SPELL_LEGACY_ICONS) === "true" ? this.get(Setting.SPELL_LEGACY_ICONS) : false
  }
}
