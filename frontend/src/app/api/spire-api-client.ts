import axios       from "axios";
import UserContext from "@/app/user/UserContext";

export class SpireApiClient {
  static getAxiosConfig() {
    let spireAxiosConfig = <any>{
      baseURL: process.env.VUE_APP_BACKEND_BASE_URL + "/api/v1/",
    }

    if (UserContext.getAccessToken() !== "") {
      spireAxiosConfig.headers = {'Authorization': 'Bearer ' + UserContext.getAccessToken()}
    }

    return spireAxiosConfig
  }

  static newAxiosWithConfig() {
    // @ts-ignore
    return axios.create(this.getAxiosConfig())
  }

  static v1() {
    return this.newAxiosWithConfig()
  }
}
