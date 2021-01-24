import axios       from "axios";
import UserContext from "@/app/user/UserContext";

export class SpireApiClient {
  static getBasePath() {
    return process.env.VUE_APP_BACKEND_BASE_URL ?
      process.env.VUE_APP_BACKEND_BASE_URL :
      window.location.host
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

  static getOpenApiConfig() {
    let openApiConfig = <any>{baseOptions: SpireApiClient.getAxiosConfig()}
    openApiConfig.basePath = this.getBaseV1Path()

    return openApiConfig
  }

  static newAxiosWithConfig() {
    // @ts-ignore
    return axios.create(this.getAxiosConfig())
  }

  static v1() {
    return this.newAxiosWithConfig()
  }
}
