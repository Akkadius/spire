import {SPECIAL_ATTACKS} from "@/app/constants/eq-special-attacks";

export class Npcs {
  public static getCleanName(name) {
    name = name.replace(/[^a-z0-9 _]/gi, '')
    name = name.replaceAll("_", " ")
    return name
  }

  static specialAbilitiesToHuman(abilities) {
    let rAbilities = <any>[]
    for (let a of abilities.split("^")) {
      const d              = a.split(",")
      const ability        = d[0] ? parseInt(d[0]) : 0
      const abilityEnabled = d[1] ? parseInt(d[1]) : 0
      if (ability > 0 && abilityEnabled) {
        if (ability === SPECIAL_ATTACKS.SUMMON) {
          rAbilities.push("Summon" + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.ENRAGE) {
          rAbilities.push("Enrage " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.AREA_RAMPAGE) {
          rAbilities.push("Area Rampage " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.FLURRY) {
          rAbilities.push("Flurry " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.TRIPLE) {
          rAbilities.push("Triple " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.QUAD) {
          rAbilities.push("Quad " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.INNATE_DW) {
          rAbilities.push("Innate DW " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.BANE) {
          rAbilities.push("Bane " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.MAGICAL) {
          rAbilities.push("Magical " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.RANGED_ATK) {
          rAbilities.push("Ranged Attack " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNSLOWABLE) {
          rAbilities.push("Unslowable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNMEZABLE) {
          rAbilities.push("Unmezable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNCHARMABLE) {
          rAbilities.push("Uncharmable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNSTUNABLE) {
          rAbilities.push("Unstunable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNSNAREABLE) {
          rAbilities.push("Unsnareable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNFEARABLE) {
          rAbilities.push("Unfearable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.UNDISPELLABLE) {
          rAbilities.push("Undispellable " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_MELEE) {
          rAbilities.push("Immune to Melee " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_MAGIC) {
          rAbilities.push("Immune to Magic " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_FLEEING) {
          rAbilities.push("Immune to Fleeing " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_MELEE_EXCEPT_BANE) {
          rAbilities.push("Immune to Melee Except Bane " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_MELEE_NONMAGICAL) {
          rAbilities.push("Immune to Melee Non-Magical " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_AGGRO) {
          rAbilities.push("Immune to Aggro " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_AGGRO_ON) {
          rAbilities.push("Immune to Aggro On " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_CASTING_FROM_RANGE) {
          rAbilities.push("Immune to Casting from Range " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_FEIGN_DEATH) {
          rAbilities.push("Immune to Feign Death " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_TAUNT) {
          rAbilities.push("Immune to Taunt " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.NPC_TUNNELVISION) {
          rAbilities.push("NPC Tunnelvision " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.NPC_NO_BUFFHEAL_FRIENDS) {
          rAbilities.push("NPC doesn't buff or heal other NPC's " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_PACIFY) {
          rAbilities.push("Immune to Pacify " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.LEASH) {
          rAbilities.push("Immune to Leash " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.TETHER) {
          rAbilities.push("Immune to Tether " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.DESTRUCTIBLE_OBJECT) {
          rAbilities.push("Destructible Object " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.NO_HARM_FROM_CLIENT) {
          rAbilities.push("Immune to Harm from Clients " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.ALWAYS_FLEE) {
          rAbilities.push("Always flees " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.FLEE_PERCENT) {
          rAbilities.push("Flees at percent X " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.ALLOW_BENEFICIAL) {
          rAbilities.push("Allow Beneficial " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.DISABLE_MELEE) {
          rAbilities.push("Disable Melee " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.NPC_CHASE_DISTANCE) {
          rAbilities.push("NPC Chase Distances " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.ALLOW_TO_TANK) {
          rAbilities.push("Allow to tank " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IGNORE_ROOT_AGGRO_RULES) {
          rAbilities.push("Ignore root aggro rules " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.CASTING_RESIST_DIFF) {
          rAbilities.push("Casting resist diff " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.COUNTER_AVOID_DAMAGE) {
          rAbilities.push("Counter avoid damage " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.PROX_AGGRO) {
          rAbilities.push("Proximity Aggro " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_RANGED_ATTACKS) {
          rAbilities.push("Immune to Ranged Attacks " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_DAMAGE_CLIENT) {
          rAbilities.push("Immune to Damage from Clients " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_DAMAGE_NPC) {
          rAbilities.push("Immune to Damage from NPCs " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_AGGRO_CLIENT) {
          rAbilities.push("Immune to Aggro from Clients " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_AGGRO_NPC) {
          rAbilities.push("Immune to Aggro from NPCs " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.MODIFY_AVOID_DAMAGE) {
          rAbilities.push("Modify avoid damage " + ` (${ability})`)
        }
        if (ability === SPECIAL_ATTACKS.IMMUNE_FADING_MEMORIES) {
          rAbilities.push("Immune to fading memories " + ` (${ability})`)
        }
      }
    }

    return rAbilities
  }
}
