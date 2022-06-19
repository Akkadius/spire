import {Npcs} from "@/app/npcs";

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

  static async getMerchantsByZone(zone: string, version: number) {
    let r   = (await Npcs.getNpcsByZone(
      zone,
      version,
      ["Spawnentries.NpcType.Merchantlists.Items"]
    )).filter((e) => {
      let hasItems = false
      if (e.merchantlists) {
        for (const e of e.merchantlists) {
          if (e.items && e.items.length > 0) {
            console.log("has items")
            console.log(e.items)
            hasItems = true;
          }
          // console.log(e)
        }
      }

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

    // edge case, if we loaded too much data and failed to load items, load each npc
    let npcs = []
    if (withItems.length === 0) {
      // chunk requests
      for (let chunk of this.spliceIntoChunks(r, 10)) {
        // @ts-ignore
        let npcIds = chunk.map((e) => {
          return e.id
        })

        const b = await Npcs.getNpcsBulk(npcIds, ["Merchantlists.Items"])
        // @ts-ignore
        npcs    = [...npcs, ...b]
      }
      r = npcs
    }

    return r
  }
}
