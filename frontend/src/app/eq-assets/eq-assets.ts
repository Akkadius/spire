import PlayerAnimations from '@/app/eq-assets/player-animations.json';
import Emitters from "@/app/eq-assets/emitters.json";
import SpellAnimations   from "@/app/eq-assets/spell-animations-map.json";

export default class EqAssets {

  public static getPlayerAnimationFileIds() {
    let ids = <any>[];
    if (PlayerAnimations[0].contents) {
      PlayerAnimations[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)
        ids.push(animationId)
      })
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    return ids
  }

  public static getEmitterPreviewFileIds() {
    let ids = <any>[];
    if (Emitters[0].contents) {
      Emitters[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)
        ids.push(animationId)
      })
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    return ids
  }

  public static spellAnimationFileIds = []

  public static getSpellAnimationFileIds() {

    // return cached set
    if (this.spellAnimationFileIds.length > 0) {
      return this.spellAnimationFileIds
    }

    let ids = <any>[];
    if (SpellAnimations[0].contents) {
      SpellAnimations[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)
        ids.push(animationId)
      })
    }

    ids.sort(function (a, b) {
      return a - b;
    });

    // cache for second retrieval
    this.spellAnimationFileIds = ids

    return ids
  }
}
