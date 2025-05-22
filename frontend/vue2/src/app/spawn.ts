import {Spawn2Api} from "@/app/api";
import {SpireApi} from "./api/spire-api";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class Spawn {
  public static async getByZone(zoneShortName: string, version: number, withRelations: boolean = false) {
    const spawn2Api = (new Spawn2Api(...SpireApi.cfg()))
    const builder   = (new SpireQueryBuilder())

    builder.where("zone", "=", zoneShortName)
    builder.where("version", "=", version)

    let includes = [
      "Spawnentries.NpcType"
    ]

    if (withRelations) {
      includes = [...includes, ...[
        "Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
        "Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList",
        "Spawnentries.NpcType.NpcFactions",
        "Spawnentries.NpcType.NpcEmotes",
        "Spawnentries.NpcType.Merchantlists.Items",
        "Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item"
      ]]
    }

    builder.includes(includes)

    // @ts-ignore
    const r = await spawn2Api.listSpawn2s(builder.get())
    if (r.status === 200 && r.data) {
      return r.data
    }

    return []
  }
}
