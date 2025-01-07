import {CharacterDatumApi} from "@/app/api/api/character-datum-api";
import {SpireApi} from "@/app/api/spire-api";

export class Characters {
  public static _cachedCharacters = {}

  static isPreloaded() {
    return Object.keys(this._cachedCharacters).length > 0
  }

  static cacheExists(id) {
    return this._cachedCharacters[id]
  }

  static getCacheById(id) {
    return this._cachedCharacters[id]
  }

  static async get(id) {
    if (this._cachedCharacters[id]) {
      return this._cachedCharacters[id]
    }

    const r = await (new CharacterDatumApi(...SpireApi.cfg()))
      .getCharacterDatum({id: id})

    if (r.status === 200) {
      // @ts-ignore
      if (!this._cachedCharacters[r.data.id]) {
        // @ts-ignore
        this._cachedCharacters[r.data.id] = r.data
      }

      return r.data
    }
  }

  static async bulkLoadCharacters(ids: any[]) {
    if (ids.length === 0) {
      return []
    }

    try {
      const r = await (new CharacterDatumApi(...SpireApi.cfg()))
        .getCharacterDataBulk({body: {ids: ids}})
      if (r.status === 200) {
        for (let c of r.data) {
          // @ts-ignore
          this._cachedCharacters[c.id] = c
        }

        return r.data
      }
    }
    catch (e) {
      console.error(e)
    }

    return []
  }
}
