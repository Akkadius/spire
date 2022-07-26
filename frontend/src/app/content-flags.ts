import {ContentFlagApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";

export class ContentFlags {
  public static flags = <any>[]

  public static async get() {
    if (this.flags && this.flags.length > 0) {
      return this.flags
    }

    const r = await (new ContentFlagApi(SpireApiClient.getOpenApiConfig())).listContentFlags()
    if (r.status === 200) {
      this.flags = r.data
      return r.data
    }

    return []

  }
}
