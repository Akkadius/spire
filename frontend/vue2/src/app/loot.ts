import {LoottableApi, MerchantlistApi, ZoneApi} from "@/app/api";
import {SpireApi} from "./api/spire-api";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class Loot {


  public static async listLootTables() {
    const result = await (new LoottableApi(...SpireApi.cfg()))
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

  // @TODO: Swap this out with an actual count endpoint later
  static async getTotalLootTables() {
    const r = await (new LoottableApi(...SpireApi.cfg()))
      .listLoottables(
        // @ts-ignore
        (new SpireQueryBuilder())
          .select(["id"])
          .limit(100000000)
          .get()
      );

    return r.data.length
  }

}
