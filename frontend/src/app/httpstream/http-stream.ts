import {SpireApi} from "@/app/api/spire-api";
import UserContext from "@/app/user/UserContext";

export class HttpStream {
  static get(url: string) {
    let headers             = {}
    headers['Content-Type'] = "application/json";
    if (UserContext.getAccessToken() !== "") {
      // @ts-ignore
      headers.Authorization = `Bearer ` + UserContext.getAccessToken()
    }

    return fetch(SpireApi.getBasePath() + url, {
      method: 'get',
      headers: headers,
    })
  }

  static read(r) {
    const textDecoder = new TextDecoder();
    const reader      = r.body.getReader();
    return {
      async* [Symbol.asyncIterator]() {
        for await (const chunk of HttpStream.readChunks(reader)) {
          yield textDecoder.decode(chunk)
        }
      }
    }
  }

  static readChunks(reader) {
    return {
      async* [Symbol.asyncIterator]() {
        let readResult = await reader.read();
        while (!readResult.done) {
          yield readResult.value;
          readResult = await reader.read();
        }
      },
    };
  }


}
