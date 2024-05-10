// import api from '@/app/core/api-client';

import {MeApi} from "@/app/api";
import {SpireApi} from "../api/spire-api";

const TOKEN_KEY = 'spire-web-access-token-' + location.host;

export default class UserContext {
  private static user: any          = {};
  private static loaded: boolean;
  private static permissions: any   = {};
  private static _is_admin: boolean = false;

  static isAdmin() {
    return this._is_admin
  }

  static setIsAdmin(value) {
    this._is_admin = value
  }

  /**
   * Get user data
   */
  static async getUser() {
    if (this.loaded) {
      return this.user;
    }

    try {
      // @ts-ignore
      const result = await ((new MeApi({
          accessToken: undefined,
          apiKey: undefined,
          basePath: SpireApi.getBaseV1Path(),
          // @ts-ignore
          formDataCtor: undefined,
          password: "",
          username: "",
          isJsonMime(mime: string): boolean {
            return false;
          },
          baseOptions: SpireApi.getAxiosConfig()
        })
      ).me())

      // @ts-ignore
      if (result.status === 200 && !result.data.error) {
        // @ts-ignore
        this.user   = result.data
        this.loaded = true

        if (this.user.is_admin) {
          this.setIsAdmin(true)
        }

        return this.user
      }
    } catch (e) {
      // squash
    }

    return null
  }

  /**
   * Reset user context
   */
  static reset() {
    this.user        = {};
    this.loaded      = false;
    this.permissions = {};
    this._is_admin   = false
  }

  /**
   * @param accessToken
   */
  static storeAccessToken(accessToken: string) {
    if (typeof (Storage) !== 'undefined') {
      localStorage.setItem(TOKEN_KEY, accessToken)
    }
  }

  /**
   * Get access token
   */
  static getAccessToken() {
    if (typeof (Storage) !== 'undefined') {
      const accessToken = localStorage.getItem(TOKEN_KEY) as string

      return (typeof accessToken === 'undefined' || accessToken === null ? '' : accessToken.toString().trim())
    }

    return ''
  }

  static deleteAccessToken() {
    if (typeof (Storage) !== 'undefined') {
      localStorage.removeItem(TOKEN_KEY)
    }
  }

  static getAccessTokenUserInfo() {
    let accessToken = this.getAccessToken()
    if (accessToken === '' || !accessToken.includes('.')) {
      return {}
    }

    const decodedToken = Buffer.from(accessToken.split('.')[1].trim(), 'base64').toString()

    return JSON.parse(decodedToken)
  }
}
