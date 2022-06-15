import {LoottableApi, ZoneApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class Loot {
  public static async getLoot() {
    const result = await (new LoottableApi(SpireApiClient.getOpenApiConfig()))
      .listLoottables(
        // @ts-ignore
        (new SpireQueryBuilder())
          .includes([
            "LoottableEntries",
            "LoottableEntries.Lootdrop",
            "LoottableEntries.Lootdrop.LootdropEntries",
            "LoottableEntries.Lootdrop.LootdropEntries.Item",
          ])
          .limit(100)
          .get()
      )
    if (result.status === 200) {
      return result.data
    }

    return {}
  }

}
