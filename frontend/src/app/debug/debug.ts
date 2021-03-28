import {App} from "@/constants/app";

export default class Debug {
  static blankLine() {
    if (App.DEBUG) {
      console.log("")
    }
  }

  static log(string, ...optionalParams) {
    if (App.DEBUG) {
      console.log("[debug] " + Debug.time() + " " + string, ...optionalParams)
    }
  }

  static addZero(x, n) {
    while (x.toString().length < n) {
      x = "0" + x;
    }
    return x;
  }

  static time() {
    let d  = new Date();
    let h  = Debug.addZero(d.getHours(), 2);
    let m  = Debug.addZero(d.getMinutes(), 2);
    let s  = Debug.addZero(d.getSeconds(), 2);
    let ms = d.getMilliseconds()
    return h + ":" + m + ":" + s + ":" + ms;
  }
}
