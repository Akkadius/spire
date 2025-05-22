//   static DEBUG_MODE = "debug-mode";

import {SpireApi} from "../api/spire-api";
import VueRouter, {Route} from "vue-router";
import {ROUTE} from "@/routes";
import UserContext from "@/app/user/UserContext";

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

  static getSetting(name) {
    for (let s of this.getSettings()) {
      if (s.setting === name) {
        return s
      }
    }
  }

  static getSettingValue(name, fallback = "") {
    for (let s of this.getSettings()) {
      if (s.setting === name) {
        return s.value
      }
    }

    return fallback
  }

  static async setSetting(name, value) {
    const r = await SpireApi.v1().post(`spire/setting/${name}/value/${value}`)
    if (r.status === 200) {
      await this.reload() // reload settings into memory
    }
  }

  static getSettings() {
    return this._settings;
  }

  static setSettings(value) {
    this._settings = value;
  }

  static setOS(value) {
    this._os = value;
  }

  static getOS() {
    return this._os;
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
    if (!this.getSettings()) {
      return false;
    }

    for (let s of this.getSettings()) {
      if (s.setting === "AUTH_ENABLED" && s.value === "true") {
        return true
      }
    }
    return false
  }

  private static _os;
  private static _env;
  private static _version;
  private static _features;
  private static _settings;
  private static _is_spire_initialized;

  static async init() {
    const r = await SpireApi.v1().get("/app/env")
    if (r.data && r.data.data) {
      const data = r.data.data
      this.setOS(data.os)
      this.setEnv(data.env)
      this.setVersion(data.version)
      this.setFeatures(data.features)
      this.setSettings(data.settings)
      this.setIsSpireInitialized(data.is_spire_initialized)

      // if a local install is not using auth but somehow the user still
      // has a JWT stored locally, let's purge it
      if (this.isAppLocal() && !this.isLocalAuthEnabled()) {
        UserContext.deleteAccessToken()
      }

      return true;
    }
    return false;
  }

  static async reload() {
    await this.init()
  }

  static routeCheckSpireInitialized(to: Route, router: VueRouter) {
    // @ts-ignore
    if (!AppEnv.isSpireInitialized() && to.fullPath !== ROUTE.SPIRE_INITIALIZE) {
      console.log("hello")
      // re-route to spire setup if not setup yet
      router.push(ROUTE.SPIRE_INITIALIZE).catch((e) => {
      })
    } else if (AppEnv.isSpireInitialized() && to.fullPath === ROUTE.SPIRE_INITIALIZE) {
      router.push(ROUTE.HOME).catch((e) => {
      })
    }
  }
}
