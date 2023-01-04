import {SpireApi} from "@/app/api/spire-api";
import {AaRankApi} from "@/app/api/api/aa-rank-api";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class AA {
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
    request.id = rankId

    const api = (new AaRankApi(...SpireApi.cfg()))
    // @ts-ignore
    const r   = await api.getAaRank(request)

    if (r.status === 200) {
      return r.data
    }

    return {}
  }
}
