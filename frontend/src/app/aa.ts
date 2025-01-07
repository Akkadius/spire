import {SpireApi} from "@/app/api/spire-api";
import {AaRankApi} from "@/app/api/api/aa-rank-api";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {AaAbilityApi} from "@/app/api/api/aa-ability-api";
import {DbStrApi} from "@/app/api";

export class AA {
  public static _aaRanks   = <any>[]
  public static _aaAbility = <any>[]
  public static _dbStrs    = <any>[]

  static isPreloaded() {
    return this._aaRanks.length > 0
  }

  static async preLoad() {
    if (this.isPreloaded()) {
      return
    }

    let builder = (new SpireQueryBuilder())
    builder.limit(100000)

    await Promise.all(
      [
        // @ts-ignore
        (new AaRankApi(...SpireApi.cfg())).listAaRanks(
          // @ts-ignore
          (new SpireQueryBuilder())
            .includes(["AaAbility", "SpellsNew"])
            .limit(100000000)
            .get()
        ),
        // @ts-ignore
        (new AaAbilityApi(...SpireApi.cfg())).listAaAbilities(builder.get()),
        // @ts-ignore
        (new DbStrApi(...SpireApi.cfg())).listDbStrs(builder.get())
      ]
    ).then(async (r) => {
        const aaRanks = await r[0].data
        const aaAbils = await r[1].data
        const dbStrs  = await r[2].data.filter((e) => {
          return e.type === 1;
        })

        if (aaRanks.length > 0) {
          this._aaRanks = aaRanks
        }
        if (aaAbils.length > 0) {
          this._aaAbility = aaAbils
        }
        if (dbStrs.length > 0) {
          this._dbStrs = dbStrs
        }
      }
    ).catch((e) => {
      console.error(e)
    });
  }

  static async getAARankById(rankId) {
    let builder = (new SpireQueryBuilder())
    builder.includes(
      [
        "AaAbility",
        "SpellsNew",
      ]
    )
    let request = builder.get()
    // @ts-ignore
    request.id  = rankId

    const api = (new AaRankApi(...SpireApi.cfg()))
    // @ts-ignore
    const r   = await api.getAaRank(request)

    if (r.status === 200) {
      return r.data
    }

    return {}
  }

  static getAARankByRankId(rankId) {
    return this._aaRanks.find((e) => {
      return parseInt(e.id) === parseInt(rankId)
    })
  }

  static getAANameDbString(id) {
    return this._dbStrs.find((e) => {
      return parseInt(e.id) === parseInt(id);
    })
  }
}
