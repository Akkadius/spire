import * as util from 'util'
import {SpireApi} from "@/app/api/spire-api";

export class EqemuWebsocketClient {
  private uuid: string = ''
  private ws: any      = null
  private port: number = 0

  public methods = {
    SUBSCRIBE: 'subscribe',
    GET_LOGSYS_CATEGORIES: 'get_logsys_categories',
    SET_LOGGING_LEVEL: 'set_logging_level',
    GET_ZONE_ATTRIBUTES: 'get_zone_attributes',
  }

  public subscriptions = {
    LOGGING: 1
  }

  /**
   * @param port
   */
  public async initClient(port: number) {
    this.ws = new Promise(async (resolve, reject) => {
      const r = await SpireApi.v1().get("eqemuserver/get-websocket-auth")

      // @ts-ignore
      let auth = null
      if (r.status === 200) {
        auth = r.data
      }

      this.port = port
      this.ws   = new WebSocket('ws://' + location.hostname + ':' + port)
      this.uuid = this.generateUUID()

      console.log(auth)

      this.ws.onerror = (event: any) => {
        this.log(util.format('Error %s', JSON.stringify(event)))
        reject(event)
      }

      this.ws.onclose = (event: any) => {
        this.log(util.format('Closed %s', JSON.stringify(event)))
        reject(event)
      }

      try {
        this.ws.onopen = (event: any) => {
          console.log("open")
          this.log(util.format('Opened %s', JSON.stringify(event)))
          // @ts-ignore
          this.send('login', [auth.account_name, auth.password])
          this.ws.onmessage = (event: any) => {
            const response = JSON.parse(event.data)
            const loggedIn =
                    (
                      response
                      && response.data
                      && response.data.status === 'Ok'
                      && response.method == 'login'
                    )

            if (loggedIn) {
              resolve(this.ws)
            }
          }
        }
      } catch (e) {
        console.log('Open' + e)
      }
    })

    return this.ws
  }

  /**
   * @param message
   */
  private log(message: string) {
    console.log('[EqemuWebsocketClient] Port: %s | %s', this.port, message)
  }

  /**
   * @param method
   * @param parameters
   */
  public send(method: string, parameters: any[] = []) {
    this.log(util.format('Method: \'%s\' Params %s', method, JSON.stringify(parameters)))

    this.ws.send(
      JSON.stringify(
        {
          id: this.uuid,
          method: method,
          params: parameters
        }, null, 2)
    )
  }

  public subscribeToLogging() {
    this.send(this.methods.SUBSCRIBE, [this.subscriptions.LOGGING])
  }

  public getLogCategories() {
    this.send(this.methods.GET_LOGSYS_CATEGORIES)
  }

  public getZoneAttributes() {
    this.send(this.methods.GET_ZONE_ATTRIBUTES)
  }

  public setLoggingLevel(category_id: number, log_level: number) {
    this.send(this.methods.SET_LOGGING_LEVEL, [category_id, log_level])
  }

  private generateUUID() {
    let dateTime = new Date().getTime()
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
      let random = (dateTime + Math.random() * 16) % 16 | 0
      dateTime   = Math.floor(dateTime / 16)
      return (c === 'x' ? random : (random & 0x3 | 0x8)).toString(16)
    })
  }
}
