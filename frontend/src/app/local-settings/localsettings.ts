export class Setting {
  static DEBUG_MODE                  = "debug-mode";
  static TAB_HOVER                   = "tab-hover";
  static SPELL_LEGACY_ICONS          = "spell-legacy-icons";
  static DEFAULT_LANGUAGE_PREFERENCE = "default-language-preference";

  // spire update
  static LAST_CHECKED_UPDATE_TIME    = "last-checked-update-time";
  static LATEST_UPDATE_VERSION       = "latest-update-version";
  static LATEST_RELEASE_PAYLOAD      = "latest-release-payload";
  static IGNORED_UPDATE_VERSION      = "ignored-update-version";
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

  static isSpellLegacyIconsEnabled() {
    return this.get(Setting.SPELL_LEGACY_ICONS) === "true" ? this.get(Setting.SPELL_LEGACY_ICONS) : false
  }

  static setLastCheckedUpdateTime(time: number) {
    this.set(Setting.LAST_CHECKED_UPDATE_TIME, time)
  }

  static getLastCheckedUpdateTime() {
    return this.get(Setting.LAST_CHECKED_UPDATE_TIME) ? this.get(Setting.LAST_CHECKED_UPDATE_TIME) : 0
  }

  static setLatestUpdateVersion(version: string) {
    this.set(Setting.LATEST_UPDATE_VERSION, version)
  }

  static getLatestUpdateVersion() {
    return this.get(Setting.LATEST_UPDATE_VERSION) ? this.get(Setting.LATEST_UPDATE_VERSION) : ""
  }

  static getIgnoredUpdateVersion() {
    return this.get(Setting.IGNORED_UPDATE_VERSION) ? this.get(Setting.IGNORED_UPDATE_VERSION) : ""
  }

  static setIgnoredUpdateVersion(version: string) {
    this.set(Setting.IGNORED_UPDATE_VERSION, version)
  }

  static setLatestReleasePayload(payload: string) {
    this.set(Setting.LATEST_RELEASE_PAYLOAD, payload)
  }

  static getLatestReleasePayload() {
    return this.get(Setting.LATEST_RELEASE_PAYLOAD) ? this.get(Setting.LATEST_RELEASE_PAYLOAD) : ""
  }

  static clearUpdateVariables() {
    localStorage.removeItem(Setting.LAST_CHECKED_UPDATE_TIME)
    localStorage.removeItem(Setting.LATEST_UPDATE_VERSION)
    localStorage.removeItem(Setting.LATEST_RELEASE_PAYLOAD)
    localStorage.removeItem(Setting.IGNORED_UPDATE_VERSION)
  }
}
