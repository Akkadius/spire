import {SpireApiClient} from "@/app/api/spire-api-client";
import {HttpStatus} from "@/app/api/http-status";

export default class EqAssets {

  public static playerAnimations = []

  public static async getPlayerAnimationFileIds() {
    // return cached set
    if (this.playerAnimations.length > 0) {
      return this.playerAnimations
    }

    let ids = <any>[];

    const r = await SpireApiClient.v1().get('/static-map/player-animations.json')
    if (r.status === HttpStatus.OK) {
      if (r.data[0].contents) {
        r.data[0].contents.forEach((row) => {
          const pieces      = row.name.split(/\//);
          const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
          const animationId = parseInt(fileName)
          ids.push(animationId)
        })
      }

      ids.sort(function (a, b) {
        return a - b;
      });
    }

    // cache for second retrieval
    this.playerAnimations = ids

    return ids
  }

  public static async getEmitterPreviewFileIds() {
    let ids = <any>[];
    const r = await SpireApiClient.v1().get('/static-map/emitters.json')
    if (r.status === HttpStatus.OK) {
      if (r.data[0].contents) {
        r.data[0].contents.forEach((row) => {
          const pieces      = row.name.split(/\//);
          const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
          const animationId = parseInt(fileName)
          ids.push(animationId)
        })
      }

      ids.sort(function (a, b) {
        return a - b;
      });
    }

    return ids
  }

  public static spellIcons = []

  public static async getSpellIcons() {
    // return cached set
    if (this.spellIcons.length > 0) {
      return this.spellIcons
    }

    let ids = <any>[];

    const r = await SpireApiClient.v1().get('/static-map/spell-icons-map.json')
    if (r.status === HttpStatus.OK) {
      r.data[0].contents.forEach((row) => {
        const pieces   = row.name.split(/\//);
        const fileName = pieces[pieces.length - 1];
        const iconId   = fileName.replace(".png", "")

        ids.push(iconId)
      })
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    ids = ids.filter((id) => {
      return parseInt(id) <= 216
    })

    // cache for second retrieval
    this.spellIcons = ids

    return ids
  }

  public static spellAnimationFileIds = []

  public static async getSpellAnimationFileIds() {
    // return cached set
    if (this.spellAnimationFileIds.length > 0) {
      return this.spellAnimationFileIds
    }

    let ids = <any>[];

    const r = await SpireApiClient.v1().get('/static-map/spell-animations-map.json')
    if (r.status === HttpStatus.OK) {
      if (r.data[0].contents) {
        r.data[0].contents.forEach((row) => {
          const pieces      = row.name.split(/\//);
          const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
          const animationId = parseInt(fileName)
          ids.push(animationId)
        })
      }
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    // cache for second retrieval
    this.spellAnimationFileIds = ids

    return ids
  }

  public static itemIcons = []

  public static async getItemIcons() {
    // return cached set
    if (this.itemIcons.length > 0) {
      return this.itemIcons
    }

    let ids = <any>[];
    const r = await SpireApiClient.v1().get('/static-map/item-icons-map.json')
    if (r.status === HttpStatus.OK) {
      if (r.data[0].contents) {
        r.data[0].contents.forEach((row) => {
          const pieces   = row.name.split(/\//);
          const fileName = pieces[pieces.length - 1];
          ids.push(fileName)
        })
      }
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    // cache for second retrieval
    this.itemIcons = ids

    return ids
  }

  public static npcModels = []

  public static async getNpcModels() {
    // return cached set
    if (this.npcModels.length > 0) {
      return this.npcModels
    }

    let ids = <any>[];
    const r = await SpireApiClient.v1().get('/static-map/npc-models-map.json')
    if (r.status === HttpStatus.OK) {
      if (r.data[0].contents) {
        r.data[0].contents.forEach((row) => {
          const pieces     = row.name.split(/\//);
          const fileName   = pieces[pieces.length - 1];
          const paramSplit = fileName.split("_")
          const raceId     = paramSplit[1].trim();
          ids.push({fileName: fileName, raceId: raceId})
        })
      }
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    // cache for second retrieval
    this.npcModels = ids

    return ids
  }

  public static itemModelFileNames = []

  public static async getItemModelFileNames() {
    // return cached set
    if (this.itemModelFileNames.length > 0) {
      return this.itemModelFileNames
    }

    let ids = <any>[];
    const r = await SpireApiClient.v1().get('/static-map/objects-map.json')
    if (r.status === HttpStatus.OK) {
      if (r.data[0].contents) {
        r.data[0].contents.forEach((row) => {
          const pieces   = row.name.split(/\//);
          const fileName = pieces[pieces.length - 1];
          ids.push(fileName)
        })
      }
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    // cache for second retrieval
    this.itemModelFileNames = ids

    return ids
  }

  public static spellAnimNameMappings = []

  public static async getSpellAnimNameMappings() {
    // return cached set
    if (this.spellAnimNameMappings.length > 0) {
      return this.spellAnimNameMappings
    }

    const r = await SpireApiClient.v1().get('/static-map/spell-icon-anim-name-map.json')
    if (r.status === HttpStatus.OK) {
      this.spellAnimNameMappings = r.data
      return r.data
    }

    return []
  }

}
