import {Npcs} from "@/app/npcs";
import {ItemApi, MerchantlistApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {chunk} from "@/app/utility/chunk";

export class Merchants {
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

    if (parseInt(name) > 0) {
      builder.where("id", "=", name)

    } else {
      builder.where("name", "=", name)
    }

    builder.includes(["Merchantlists.NpcTypes"])

    const items = await (new ItemApi(SpireApiClient.getOpenApiConfig()))
      .listItems(
        // @ts-ignore
        builder.get()
      )

    // console.log(items.data)

    let npcIds: any[] = []
    for (let i of items.data) {
      // console.log(i)
      if (i.merchantlists) {
        // console.log(i)
        for (let m of i.merchantlists) {
          if (m.npc_types) {
            for (let n of m.npc_types) {
              npcIds.push(n.id)
            }
          }
        }
      }
    }

    if (npcIds.length === 0) {
      return []
    }

    // @ts-ignore
    let r = (await Npcs.getNpcsBulk(
      npcIds,
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
      {
        relations: ["Spawnentries.NpcType.Merchantlists.Items"]
      }
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

  static async deleteMerchantEntry(merchantId: number, slotId: number) {
    let request = (new SpireQueryBuilder())
      .where("merchantid", "=", merchantId)
      .where("slot", "=", slotId)
      .get()

    // @ts-ignore
    return await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .deleteMerchantlist(
        {
          id: merchantId,
        },
        {query: request}
      )
  }

  static async getMerchantsBulk(ids: number[], relations: any[] = []) {
    const r = await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .getMerchantlistsBulk(
        {
          body: {
            ids: ids
          }
        },
        {
          query: // @ts-ignore
            (new SpireQueryBuilder())
              .groupBy(["merchantid"])
              .includes(relations)
              .orderBy(["merchantid"])
              .orderDirection("desc")
              .get()
        }
      )

    if (r.status === 200) {
      return r.data
    }
  }

  static async fallBackChunkLoad(inNpcs) {
    // edge case, if we loaded too much data and failed to load items, load each npc
    let npcs = []
    // chunk requests
    for (let c of chunk(inNpcs, 10)) {
      // @ts-ignore
      let npcIds = c.map((e) => {
        return e.id
      })

      const b = await Npcs.getNpcsBulk(npcIds, ["Merchantlists.Items"])
      // @ts-ignore
      npcs    = [...npcs, ...b]
    }
    return npcs
  }

  static async getById(id: number, relations: any[] = []) {
    const r = await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .listMerchantlists(
        // @ts-ignore
        (new SpireQueryBuilder())
          .where("merchantid", "=", id)
          .includes(relations)
          .get()
      )

    if (r.status === 200) {
      return r.data
    }

    return []
  }

  // @TODO: Swap this out with an actual count endpoint later
  static async getTotalMerchants() {
    const r = await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .listMerchantlists(
        // @ts-ignore
        (new SpireQueryBuilder())
          .groupBy(["merchantid"])
          .orderBy(["merchantid"])
          .orderDirection("desc")
          .limit(100000000)
          .get()
      )

    return r.data.length
  }

  static async updateSlotForEntry(merchantId: number, currentSlot: number, destinationEntry: any) {
    let request = (new SpireQueryBuilder())
      .where("merchantid", "=", merchantId)
      .where("slot", "=", currentSlot)
      .get()

    // @ts-ignore
    return await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .updateMerchantlist(
        {
          id: merchantId,
          merchantlist: destinationEntry
        },
        {query: request}
      )
  }

  static async addMerchantListEntry(merchantlistEntry: any){
    // @ts-ignore
    return await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .createMerchantlist(
        {merchantlist: merchantlistEntry}
      )
  }

  static async addItemToMerchant(merchantId: number, newSlot: any, itemId: any) {
    const newMerchantEntry = {
      merchantid: merchantId,
      slot: newSlot,
      item: itemId,
      faction_required: -1100,
      classes_required: 65535,
      probability: 100,
      min_expansion: -1,
      max_expansion: -1,
      bucket_comparison: 0,
    }

    // @ts-ignore
    return await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .createMerchantlist(
        {merchantlist: newMerchantEntry}
      )
  }

  static async create() {
    const r = await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .listMerchantlists(
        // @ts-ignore
        (new SpireQueryBuilder())
          .orderBy(["merchantid"])
          .orderDirection("desc")
          .limit(1)
          .get()
      )

    const newMerchantId    = r.status === 200 && r.data && r.data[0] && r.data[0].merchantid ? r.data[0].merchantid + 1 : 0
    const newMerchantEntry = {
      merchantid: newMerchantId,
      slot: 1,
      item: 1001,
    }

    // @ts-ignore
    return await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .createMerchantlist(
        {merchantlist: newMerchantEntry}
      )
  }

  static async deleteMerchant(merchantid: number) {

    // @ts-ignore
    return await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
      .deleteMerchantlist(
        {id: merchantid}
      )
  }
}
