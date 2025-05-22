import {SpireApi} from "./api/spire-api";
import {HttpStatus} from "@/app/api/http-status";

export class FreeIdFetcher {
  public static async get(table: string, idFieldName: string, reservedFieldName: string = "name") {
    let r = await SpireApi.v1().get(`query/free-ids-reserved/${table}/${idFieldName}/${reservedFieldName}`)
    if (r.status === HttpStatus.OK) {

      // grab first "reserved" entry available
      if (r.data.data.length > 0) {
        return parseInt(r.data.data[0].id)
      }

      // grab first free id in range entry available
      r = await SpireApi.v1().get(`query/free-id-ranges/${table}/${idFieldName}`)
      if (r.status === HttpStatus.OK) {
        if (r.data && r.data.data) {
          return parseInt(r.data.data[0].start_id)
        }
      }
    }

    return -1
  }
}
