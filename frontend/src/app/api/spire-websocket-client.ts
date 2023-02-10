import util from "util";
import {SpireApi} from "@/app/api/spire-api";
import UserContext from "@/app/user/UserContext";

export class SpireWebsocketClient {
  static websocket(): any {
    return this._ws;
  }

  private static _ws: any = null

  static connect() {
    // singleton, we should have connected
    if (this._ws) {
      return;
    }

    const uri = util.format(
      "ws:%s/api/v1/websocket?jwt=" + UserContext.getAccessToken(),
      SpireApi.getBasePath().replaceAll("http://", "")
    )

    this._ws = new WebSocket(uri)

    this._ws.onopen = () => {
      console.log('Connected')
    }

    this._ws.onerror = (e) => {
      console.log("error", e)
      this._ws.close();
    }
    this._ws.onclose = (e) => {
      console.log("Websocket connection closed")
      console.log(e)
    }

    return this._ws
  }
}
