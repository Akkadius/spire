import axios, {AxiosRequestConfig} from 'axios'
import util from "util";

export class OcculusClient {

  /**
   * Base URL for admin api
   */
  private static _baseUrl = (
    process.env.VUE_APP_BACKEND_BASE_URL
      ? process.env.VUE_APP_BACKEND_BASE_URL
      : window.location.origin
  );

  static getBaseUrl(): string {
    return this._baseUrl
  }

  static client(config: AxiosRequestConfig = {}) {
    config.baseURL = (this._baseUrl + '/api/v1/admin/occulus')
    if (this.getAccessToken() !== '') {
      config.headers = {'Authorization': 'Bearer ' + this.getAccessToken()}
    }

    return axios.create(config)
  }

  /**
   * @param accessToken
   */
  static storeAccessToken(accessToken: string) {
    if (typeof (Storage) !== 'undefined') {
      localStorage.setItem('access_token', accessToken)
    }
  }

  /**
   * Get access token
   */
  static getAccessToken() {
    if (typeof (Storage) !== 'undefined') {
      const accessToken = localStorage.getItem('access_token') as string

      return (typeof accessToken === 'undefined' || accessToken === null ? '' : accessToken.toString().trim())
    }

    return ''
  }

  static deleteAccessToken() {
    if (typeof (Storage) !== 'undefined') {
      localStorage.removeItem('access_token')
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

  /**
   * @param url
   * @param returnResponse
   */
  static async get(url: string, returnResponse = false) {
    try {
      if (returnResponse) {
        return await this.client().get(url)
      }

      const response = await this.client().get(url)
      return response.data
    } catch (e) {
      this.errorHandler(e)
    }
  }

  /**
   * @param url
   */
  static async getFile(url: string) {
    try {
      const response = await this.client({responseType: 'blob'}).get(url,
        {
          onDownloadProgress: (progressEvent) => {
            let percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total)

            window.dispatchEvent(
              new CustomEvent(
                'file-download-progress',
                {
                  detail:
                    {
                      percent: percentCompleted,
                      loaded: progressEvent.loaded,
                      total: progressEvent.total
                    }
                }
              )
            )
          }
        }
      )

      return response
    } catch (e) {
      this.errorHandler(e)
    }
  }

  /**
   * @param url
   * @param data
   * @param returnResponse
   */
  static async post(url: string, data: any, returnResponse = false) {
    try {
      if (returnResponse) {
        return await this.client().post(url, data)
      }

      const response = await this.client().post(url, data)

      return response.data
    } catch (e) {
      this.errorHandler(e)
    }
  }

  private static errorHandler(e: any) {
    if (!e.status) {
      if (!e.response) {
        return false
      } else {
        if (e.response.status === 401) {
          window.location.assign('/login')
        }
        return e.response
      }
    }
  }

  static async getDashboardStats() {
    return this.get('/dashboard/stats')
  }

  static async getProcessCounts() {
    return this.get('/server/process_counts')
  }

  static async getSysInfo() {
    return this.get('/server/sysinfo')
  }

  static async getZoneList() {
    return this.get('/zoneserver/list')
  }

  static async stopServer(data: any) {
    return this.post('/server/stop', data)
  }

  static async startServer() {
    return this.get('/server/start')
  }

  static async restartServer(data: any) {
    return this.post('/server/restart', data)
  }

  static async cancelRestartServer(data: any) {
    return this.post('/server/restart/cancel', data)
  }

  static async getZoneAttributes(port: number) {
    return this.get('/zoneserver/' + port + '/attributes')
  }

  static async getZoneClientList(port: number) {
    return this.get('/zoneserver/' + port + '/client/list/detail')
  }

  static async getZoneNpcList(port: number) {
    return this.get('/zoneserver/' + port + '/npc/list/detail')
  }

  static async getZoneClientNetstats(port: number) {
    return this.get('/zoneserver/' + port + '/netstats')
  }

  static async getZoneNetstatChartData(port: number) {
    return this.get('/zoneserver/netstat/' + port + '/chart')
  }

  static async logout() {
    return this.get('/auth/logout')
  }

  static async login(data: any) {
    return this.post('/auth/login', data)
  }

  static async getServerConfig() {
    return this.get('/admin/config')
  }

  static async postServerConfig(data: any) {
    return this.post('/admin/config', data)
  }

  static async getServerMotd() {
    return this.get('/admin/motd')
  }

  static async postServerMotd(data: any) {
    return this.post('/admin/motd', data)
  }

  static async getServerRules() {
    return this.get('/admin/rule_values')
  }

  static async getWorldClientList() {
    return this.get('/world/clients')
  }

  static async getWebsocketAuthorization() {
    return this.get('/server/websocket-authorization')
  }

  static async postServerRule(data: any) {
    return this.post('/admin/rule_values', data)
  }

  static async downloadBackup(backupType: string) {
    return this.getFile('/backup/' + backupType)
  }

  static async getLogCategories() {
    return this.get('/server/log_categories')
  }

  static async postLauncherConfig(data: any) {
    return this.post('/server/launcher/config', data)
  }

  static async getLauncherConfig() {
    return this.get('/server/launcher/config')
  }

  static async killProcessByPid(pid) {
    return this.post('/admin/process/kill/' + pid, {})
  }

  static async getServerLogs() {
    return (await this.get('/admin/logs/list', true))
  }

  static async getServerLog(fileName: string) {
    return (await this.get(
      util.format('/admin/logs/view/%s', encodeURIComponent(fileName)), true
    ))
  }

  static async getServerCodeBranches() {
    return (await this.get('/admin/code/git/branches', true))
  }

  static async getServerCodeBranch() {
    return (await this.get('/admin/code/git/branch', true))
  }

  static async setServerCodeBranch(branch: string) {
    return this.post(
      util.format(
        '/admin/code/git/branch/%s',
        encodeURIComponent(branch)
      ),
      {},
      true
    )
  }

  static async getBuildStatus() {
    return (await this.get('/admin/code/build/status', true))
  }

  static async buildServerCode() {
    return this.post('/admin/code/build', {}, true)
  }

  static async updateServerCode() {
    return this.post('/admin/code/git/update', {}, true)
  }

  static async cancelServerBuild() {
    return this.post('/admin/code/build/cancel', {}, true)
  }


  static async getServerQuestsBranches() {
    return (await this.get('/admin/code/git/quests/branches', true))
  }

  static async getServerQuestsBranch() {
    return (await this.get('/admin/code/git/quests/branch', true))
  }

  static async setServerQuestsBranch(branch: string) {
    return this.post(
      util.format(
        '/admin/code/git/quests/branch/%s',
        encodeURIComponent(branch)
      ),
      {},
      true
    )
  }
}
