import {Npcs} from "@/app/npcs";
import {ItemApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class Merchants {
  static spliceIntoChunks(arr, chunkSize) {
    const res = [];
    while (arr.length > 0) {
      const chunk = arr.splice(0, chunkSize);
      // @ts-ignore
      res.push(chunk);
    }
    return res;
  }

  static async getMerchantsByName(name: string) {
    // @ts-ignore
    let r = (await Npcs.listNpcsByName(
      name,
      ["Merchantlists.Items"]
    )).filter((e) => {
      // @ts-ignore
      return e.merchant_id > 0
    })

    const withItems = r.filter((e) => {
      let hasItems = false
      if (e.merchantlists) {
        for (const i of e.merchantlists) {
          // @ts-ignore
          if (i.items && i.items.length > 0) {
            hasItems = true;
          }
        }
      }

      return hasItems
    })

    if (withItems.length === 0) {
      r = (await this.fallBackChunkLoad(r))
    }

    return r
  }

  static async getMerchantsByItemName(name: string) {
    let builder = (new SpireQueryBuilder())
    builder.where("name", "like", name)
    builder.includes(["Merchantlists"])

    const items = await (new ItemApi(SpireApiClient.getOpenApiConfig()))
      .listItems(
        // @ts-ignore
        builder.get()
      )

    // console.log(items.data)

    for (let i of items.data) {
      if (i.merchantlists) {
        console.log(i)
      }
    }

    // @ts-ignore
    let r = (await Npcs.listNpcsByName(
      name,
      ["Merchantlists.Items"]
    )).filter((e) => {
      // @ts-ignore
      return e.merchant_id > 0
    })

    const withItems = r.filter((e) => {
      let hasItems = false
      if (e.merchantlists) {
        for (const i of e.merchantlists) {
          // @ts-ignore
          if (i.items && i.items.length > 0) {
            hasItems = true;
          }
        }
      }

      return hasItems
    })

    if (withItems.length === 0) {
      r = (await this.fallBackChunkLoad(r))
    }

    return r
  }

  static async getMerchantsByZone(zone: string, version: number) {
    let r = (await Npcs.getNpcsByZone(
      zone,
      version,
      ["Spawnentries.NpcType.Merchantlists.Items"]
    )).filter((e) => {
      return e.merchant_id > 0
    })

    const withItems = r.filter((e) => {
      let hasItems = false
      if (e.merchantlists) {
        for (const i of e.merchantlists) {
          if (i.items && i.items.length > 0) {
            hasItems = true;
          }
        }
      }

      return hasItems
    })

    if (withItems.length === 0) {
      r = (await this.fallBackChunkLoad(r))
    }

    return r
  }

  static async fallBackChunkLoad(inNpcs) {
    // edge case, if we loaded too much data and failed to load items, load each npc
    let npcs = []
    // chunk requests
    for (let chunk of this.spliceIntoChunks(inNpcs, 10)) {
      // @ts-ignore
      let npcIds = chunk.map((e) => {
        return e.id
      })

      const b = await Npcs.getNpcsBulk(npcIds, ["Merchantlists.Items"])
      // @ts-ignore
      npcs    = [...npcs, ...b]
    }
    return npcs
  }
}
