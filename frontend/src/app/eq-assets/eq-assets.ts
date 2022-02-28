import PlayerAnimations from '@/app/eq-assets/player-animations.json';
import Emitters from "@/app/eq-assets/emitters.json";
import SpellAnimations from "@/app/eq-assets/spell-animations-map.json";
import SpellIcons from "@/app/eq-assets/spell-icons-map.json";

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

  public static spellIcons = []

  public static getSpellIcons() {
    // return cached set
    if (this.spellIcons.length > 0) {
      return this.spellIcons
    }

    let ids = <any>[];
    if (SpellIcons[0].contents) {
      SpellIcons[0].contents.forEach((row) => {
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
