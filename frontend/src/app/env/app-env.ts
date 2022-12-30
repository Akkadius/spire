//   static DEBUG_MODE = "debug-mode";

import {SpireApi} from "../api/spire-api";

// these should be mirrored with env/app.go
const AppEnvTesting    = "testing"
const AppEnvLocal      = "local"
const AppEnvDev        = "dev"
const AppEnvDesktop    = "desktop"
const AppEnvStaging    = "staging"
const AppEnvProduction = "production"

export class AppEnv {
  static isSpireInitialized() {
    return this._is_spire_initialized;
  }

  static setIsSpireInitialized(value) {
    this._is_spire_initialized = value;
  }

  static getFeatures() {
    return this._features;
  }

  static setFeatures(value) {
    this._features = value;
  }

  static getSettings() {
    return this._settings;
  }

  static setSettings(value) {
    this._settings = value;
  }

  static getEnv() {
    return this._env;
  }

  static setEnv(value) {
    this._env = value;
  }

  static getVersion() {
    return this._version;
  }

  static setVersion(value) {
    this._version = value;
  }

  static isAppLocal() {
    return this.getEnv() !== AppEnvProduction &&
      ([AppEnvLocal, AppEnvDesktop, AppEnvDev].includes(this.getEnv()))
  }

  static isAppProduction() {
    return this.getEnv() === AppEnvProduction
  }

  static isGithubAuthEnabled() {
    return this.getFeatures() ? this.getFeatures().github_auth_enabled && this.isAppProduction() : false
  }

  static isLocalAuthEnabled() {
    for (let s of this.getSettings()) {
      if (s.setting === "AUTH_ENABLED" && s.value === "true") {
        return true
      }
    }
    return false
  }

  private static _env;
  private static _version;
  private static _features;
  private static _settings;
  private static _is_spire_initialized;

  static async init() {
    const r = await SpireApi.v1().get("/app/env")
    if (r.data && r.data.data) {
      const data = r.data.data
      this.setEnv(data.env)
      this.setVersion(data.version)
      this.setFeatures(data.features)
      this.setSettings(data.settings)
      this.setIsSpireInitialized(data.is_spire_initialized)

      return true;
    }
    return false;
  }
}
