//   static DEBUG_MODE = "debug-mode";

import {SpireApiClient} from "@/app/api/spire-api-client";

// these should be mirrored with env/app.go
const AppEnvTesting    = "testing"
const AppEnvLocal      = "local"
const AppEnvDev        = "dev"
const AppEnvDesktop    = "desktop"
const AppEnvStaging    = "staging"
const AppEnvProduction = "production"

export class AppEnv {
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

  private static _env;
  private static _version;

  static init() {
    SpireApiClient.v1().get("/app/env").then((response) => {
      if (response.data && response.data.data) {
        const data = response.data.data
        this.setEnv(data.env)
        this.setVersion(data.version)
      }
    })
  }
}
