import axios from "axios";
import UserContext from "@/app/user/UserContext";
import Debug from "@/app/debug/debug";
import {ROUTE} from "@/routes";


export class SpireApi {
  static getBasePath() {
    return process.env.VUE_APP_BACKEND_BASE_URL && process.env.NODE_ENV !== 'production' ?
      process.env.VUE_APP_BACKEND_BASE_URL :
      window.location.origin
  }

  static getBaseV1Path() {
    return this.getBasePath() + "/api/v1"
  }

  static getAxiosConfig() {
    let spireAxiosConfig = <any>{
      baseURL: this.getBaseV1Path()
    }

    if (UserContext.getAccessToken() !== "") {
      spireAxiosConfig.headers = {'Authorization': 'Bearer ' + UserContext.getAccessToken()}
    }

    return spireAxiosConfig
  }

  static getAccessTokenQueryString() {
    const token = this.getAccessToken()
    return token !== '' ? {'jwt': token} : {}
  }

  private static getAccessToken() {
    return UserContext.getAccessToken()
  }

  static cfg() {
    return [
      this.getOpenApiConfig(),
      this.getBaseV1Path(),
      spireAxios
    ]
  }

  private static getOpenApiConfig() {
    let openApiConfig      = <any>{baseOptions: SpireApi.getAxiosConfig()}
    openApiConfig.basePath = this.getBaseV1Path()

    return openApiConfig
  }

  static newAxiosWithConfig() {
    // @ts-ignore
    let client = axios.create(this.getAxiosConfig())

    client.interceptors.request.use(x => {
      // @ts-ignore
      x.meta                  = x.meta || {}
      // @ts-ignore
      x.meta.requestStartedAt = Date.now()

      // ${x.baseURL}
      Debug.log(`[Request Start] [${x.url}]`)

      return x
    })

    client.interceptors.response.use(response => {
      return response;
    }, error => {
      console.log("401 error", error)
      if (error.response.status === 401) {
        console.log("401 error", error.response)

        // when we're not authorized (logged in)
        // redirect to login screen
        window.location.href = ROUTE.LOGIN
        return

        // blanket error for now
        if (error.response.data && error.response.data.error) {
          setTimeout(() => {
            alert(error.response.data.error)
          }, 500)
        }

      }

      return Promise.reject(error);
    });

    client.interceptors.response.use(x => {
      // @ts-ignore
      Debug.log(`[Request End] [${(Date.now() - x.config.meta.requestStartedAt)} ms] [${x.config.url}]`)
      // Debug.log('Response:', JSON.stringify(response, null, 2))
      return x
    })

    return client
  }

  static v1() {
    return spireAxios
  }

  static reloadAxios() {
    spireAxios = SpireApi.newAxiosWithConfig()
  }
}

// singleton - we don't want to create a new instance every single time
// helpers are invoked
let spireAxios = SpireApi.newAxiosWithConfig()
