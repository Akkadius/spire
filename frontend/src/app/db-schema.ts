import {SpireApiClient} from "@/app/api/spire-api-client";
import util from "util";

export class DbSchema {
  public static async getTableSchema(tableName: string) {
    const r = await SpireApiClient.v1().get(util.format("query/schema/table/%s", tableName))
    if (r.data && r.data.data) {
      return r.data.data
    }

    return {}
  }

  public static async getTableColumns(tableName: string) {
    const schema = await this.getTableSchema(tableName)
    let fields   = <any>[]
    if (schema) {
      schema.forEach((e) => {
        fields.push(e.Column)
      })
    }

    return fields
  }
}
