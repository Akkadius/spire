import util from "util";
import {SpireApi} from "@/app/api/spire-api";
import UserContext from "@/app/user/UserContext";

export class SpireWebsocket {
  private static _ws: any = null
  private static _reconnect: boolean = false;
  private static _listeners: any[] = [];

  static connect(): any {
    // singleton, we should have connected
    if (this._ws && !this._reconnect) {
      return this._ws;
    }

    const protocol = window.location.protocol === "https:" ? "wss" : "ws"
    let basePath = SpireApi.getBasePath()
    if (basePath.startsWith("http://")) {
      basePath = basePath.replace("http://", "")
    } else if (basePath.startsWith("https://")) {
      basePath = basePath.replace("https://", "")
    }

    const uri = util.format(
      "%s:%s/api/v1/websocket?jwt=%s",
      protocol,
      basePath,
      UserContext.getAccessToken()
    )

    this._ws = new WebSocket(uri)

    this._ws.onopen = () => {
      console.log('[SpireWebsocket] Connected at ' + new Date().toLocaleString())
      this._reconnect = false
    }

    this._ws.onerror = (e: any) => {
      console.log("[SpireWebsocket] error", e)
      this._ws.close();
    }
    this._ws.onclose = (e: any) => {
      console.log("[SpireWebsocket] Websocket connection closed at " + new Date().toLocaleString())
      console.log(e)
      this._reconnect = true
      setTimeout(() => {
        console.log("[SpireWebsocket] Reconnecting...")
        this.connect()
      }, 1000);
    }

    // Re-apply stored listeners
    this._listeners.forEach(listener => {
      this._ws.addEventListener(listener.type, listener.listener);
    });


    return this._ws
  }

  // addEventListener and removeEventListener are used to store listeners
  static addEventListener(type: any, listener: any) {
    this._listeners.push({ type, listener });
    this._ws.addEventListener(type, listener);
  }

  // removeEventListener is used to remove listeners
  static removeEventListener(type: any, listener: any) {
    this._listeners = this._listeners.filter(l => l.type !== type || l.listener !== listener);
    this._ws.removeEventListener(type, listener);
  }

  static close() {
    if (this._ws) {
      this._ws.close();
      this._ws = null;
    }
  }
}
