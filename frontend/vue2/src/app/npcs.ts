import {SPECIAL_ATTACKS} from "@/app/constants/eq-special-attacks";
import {NpcTypeApi, Spawn2Api} from "@/app/api";
import {SpireApi} from "./api/spire-api";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

type NpcByZoneQueryRequest = {
  relations?: string[];
  uniqueEntries?: Boolean;
}

export class Npcs {

  /**
   * @param name
   */
  public static getCleanName(name: String) {
    name = name.replaceAll("_", " ")
    name = name.replaceAll("#", " ")
    return name
  }

  public static getRaceImage(npc: any) {
    let texture     = npc.texture
    let helmTexture = npc.helmtexture

    if (this.isPlayableRace(npc.race)) {
      if (helmTexture > 3) {
        helmTexture = 0
      }
      if (texture > 16 || (texture > 3 && texture < 10)) {
        texture = 0
      }
    }


    return npc.race + '-' + npc.gender + '-' + texture + '-' + helmTexture
  }

  public static isPlayableRace(raceId: any) {
    return [
      1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 128, 130, 330, 522
    ].includes(raceId)
  }

  /**
   * @param abilities
   */
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

  /**
   * @param id
   * @param npcType
   */
  static async updateNpc(id: number, npcType: any) {
    const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
    return await npcTypeApi.updateNpcType({
      id: id,
      npcType: npcType
    })
  }

  /**
   * @param name
   * @param relations
   */
  static async listNpcsByName(name: string, relations: any[] = []) {
    const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
    let builder      = (new SpireQueryBuilder())

    let includes: any[] = [];
    if (relations.includes("all")) {
      includes = [...includes, ...[
        "NpcSpell.NpcSpellsEntries.SpellsNew",
        "NpcFactions.NpcFactionEntries.FactionList",
        "NpcFactions",
        "NpcEmotes",
        "Merchantlists.Items",
        "Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item"
      ]]
    } else if (relations.length > 0) {
      includes = [...includes, ...relations]
    }

    builder.where("name", "like", name)
    builder.whereOr("lastname", "like", name)
    builder.includes(includes)

    // @ts-ignore
    const r = await npcTypeApi.listNpcTypes(builder.get())
    if (r.status === 200) {
      return r.data
    }
  }

  static getBaseNpcRelationships() {
    return [
      "NpcSpell.NpcSpellsEntries.SpellsNew",
      "NpcFactions.NpcFactionEntries.FactionList",
      "NpcFactions",
      "NpcEmotes",
      "Merchantlists.Items",
      "Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item"
    ]
  }

  /**
   * @param emoteId
   * @param relations
   */
  static async listNpcsByEmoteId(emoteId: number, relations: any[] = []) {
    const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
    let builder      = (new SpireQueryBuilder())

    let includes: any[] = [];
    if (relations.includes("all")) {
      includes = [...includes, ...this.getBaseNpcRelationships()]
    } else if (relations.length > 0) {
      includes = [...includes, ...relations]
    }

    builder.where("emoteid", "=", emoteId)
    builder.includes(includes)

    // @ts-ignore
    const r = await npcTypeApi.listNpcTypes(builder.get())
    if (r.status === 200) {
      return r.data
    }
  }

  /**
   * @param npcSpellsId
   * @param relations
   */
  static async listNpcsByNpcSpellsId(npcSpellsId: number, relations: any[] = []) {
    const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
    let builder      = (new SpireQueryBuilder())

    let includes: any[] = [];
    if (relations.includes("all")) {
      includes = [...includes, ...this.getBaseNpcRelationships()]
    } else if (relations.length > 0) {
      includes = [...includes, ...relations]
    }

    builder.where("npc_spells_id", "=", npcSpellsId)
    builder.includes(includes)
    builder.limit(100)

    // @ts-ignore
    const r = await npcTypeApi.listNpcTypes(builder.get())
    if (r.status === 200) {
      return r.data
    }
  }

  public static npc = <any>{}

  /**
   * @param id
   * @param relations
   */
  static async getNpc(id: number, relations: any[] = []) {
    // return cache if exist
    if (this.npc[id]) {
      return this.npc[id]
    }

    const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
    let builder      = (new SpireQueryBuilder())

    let includes: any[] = [];
    if (relations.includes("all")) {
      includes = [...includes, ...[
        "NpcSpell.NpcSpellsEntries.SpellsNew",
        "NpcFactions.NpcFactionEntries.FactionList",
        "NpcFactions",
        "NpcEmotes",
        "Merchantlists.Items",
        "Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item"
      ]]
    } else if (relations.length > 0) {
      includes = [...includes, ...relations]
    }

    // console.log(includes)

    builder.includes(includes)

    const r = await npcTypeApi.getNpcType({
        id: id,
      },
      {query: builder.get()})
    if (r.status === 200) {
      // cache
      this.npc[id] = r.data

      return r.data
    }
  }

  /**
   * @param ids
   * @param relations
   */
  static async getNpcsBulk(ids: number[], relations: any[] = []) {
    if (ids.length === 0) {
      return []
    }

    const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
    let builder      = (new SpireQueryBuilder())

    let includes: any[] = [];
    if (relations.includes("all")) {
      includes = [...includes, ...[
        "NpcSpell.NpcSpellsEntries.SpellsNew",
        "NpcFactions.NpcFactionEntries.FactionList",
        "NpcFactions",
        "NpcEmotes",
        "Merchantlists.Items",
        "Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item"
      ]]
    } else if (relations.length > 0) {
      includes = [...includes, ...relations]
    }

    builder.includes(includes)

    const r = await npcTypeApi.getNpcTypesBulk({
        body: {
          ids: ids
        },
      },
      {
        query: builder.get()
      }
    )
    if (r.status === 200) {

      for (let n of r.data) {
        // @ts-ignore
        this.npc[n.id] = n
      }

      return r.data
    }
  }

  static async getNpcsByZone(
    zoneShortName: string,
    version: number                = 0,
    request: NpcByZoneQueryRequest = {
      relations: [],
      uniqueEntries: true
    }
  ) {
    const spawn2Api = (new Spawn2Api(...SpireApi.cfg()))
    const builder   = (new SpireQueryBuilder())

    builder.where("zone", "=", zoneShortName)

    if (version === -1) {
      builder.where("version", ">=", version)
    } else {
      builder.where("version", "=", version)
    }

    let includes = [
      "Spawnentries.NpcType"
    ]

    // @ts-ignore
    if (request.relations.includes("all")) {
      includes = [...includes, ...[
        "Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
        "Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList",
        "Spawnentries.NpcType.NpcFactions",
        "Spawnentries.NpcType.NpcEmotes",
        "Spawnentries.NpcType.Merchantlists.Items",
        "Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item"
      ]]
      // @ts-ignore
    } else if (request.relations.length > 0) {
      // @ts-ignore
      includes = [...includes, ...request.relations]
    }

    builder.limit(1000000)

    builder.includes(includes)

    // @ts-ignore
    const r = await spawn2Api.listSpawn2s(builder.get())
    if (r.status === 200 && r.data) {
      let npcs = <any>[]
      for (let spawn2 of r.data) {
        if (spawn2.spawnentries) {
          for (let spawnentry of spawn2.spawnentries) {
            if (spawnentry.npc_type) {

              // only add unique entries
              if (request.uniqueEntries) {
                let found = npcs.find((e) => {
                  // @ts-ignore
                  return e.id === spawnentry.npc_type.id
                })

                if (!found) {
                  npcs.push(spawnentry.npc_type)
                }
              }
              else {
                npcs.push(spawnentry.npc_type)
              }

            }
          }
        }
      }

      // filter by unique npc ids
      return npcs
    }

    return []
  }

  // these were pulled from EOC and should be refined
  public static getFieldDescriptions() {
    return {
      "id": "ID number of the NPC",
      "name": "First Name of the NPC",
      "lastname": "Last Name of the NPC",
      "level": "Level of the NPC",
      "race": "Race of the NPC",
      "class": "Class of the NPC",
      "bodytype": "Bodytype of the NPC",
      "hp": "Maximum Hitpoints of the NPC",
      "mana": "Maximum Mana of the NPC",
      "gender": "Gender of the NPC",
      "texture": "Texture of the NPC",
      "helmtexture": "Helmet Texture of the NPC",
      "herosforgemodel": "Heroes Forge Model used on the NPC",
      "size": "Numerical size of the NPC",
      "hp_regen_rate": "Hitpoint Regeneration Rate of the NPC",
      "hp_regen_per_second": "Hitpoint Regeneration Rate (per second) of the NPC",
      "mana_regen_rate": "Mana Regen Rate of the NPC",
      "loottable_id": "Loottable Id of the NPC",
      "merchant_id": "Merchant Id of the NPC",
      "alt_currency_id": "Alt Currency Id used with the NPC",
      "npc_spells_effects_id": "Npc Spells Effects Id used with the NPC",
      "npc_faction_id": "Npc Faction Id of the NPC",
      "adventure_template_id": "Adventure Template Id of the NPC",
      "trap_template": "Trap Template of the NPC",
      "mindmg": "Minimum damge of the NPC",
      "maxdmg": "Maximum damage of the NPC",
      "attack_count": "Attack Count of the NPC",
      "npcspecialattks": "Special Attacks used by the NPC",
      "special_abilities": "Special Abilities used by the NPC",
      "aggroradius": "Aggro radius of the NPC",
      "assistradius": "Assist radius of the NPC",
      "face": "Face of the NPC",
      "luclin_hairstyle": "Luclin Hairstyle used with the NPC",
      "luclin_haircolor": "Luclin Haircolor used with the NPC",
      "luclin_eyecolor": "Luclin Eyecolor used with the NPC",
      "luclin_eyecolor_2": "Luclin Eyecolor 2 used with the NPC",
      "luclin_beardcolor": "Luclin Beardcolor used with the NPC",
      "luclin_beard": "Luclin Beard used with the NPC",
      "drakkin_heritage": "Drakkin Heritage used with the NPC",
      "drakkin_tattoo": "Drakkin Tattoo used with the NPC",
      "drakkin_details": "Drakkin Details used with the NPC",
      "armortint_id": "Armor tint Id used with the NPC",
      "armortint_red": "Armor tint Red used with the NPC",
      "armortint_green": "Armor tint Green used with the NPC",
      "armortint_blue": "Armor tint Blue used with the NPC",
      "d_melee_texture_1": "D Melee Texture 1 used with the NPC",
      "d_melee_texture_2": "D Melee Texture 2 used with the NPC",
      "ammo_idfile": "Ammo Id file used for the NPC",
      "prim_melee_type": "Primary Melee Type of the NPC",
      "sec_melee_type": "Secondary Melee Type of the NPC",
      "ranged_type": "Ranged Type of the NPC",
      "runspeed": "Runspeed of the NPC",
      "mr": "Magic Resistance of the NPC",
      "cr": "Cold Resistance of the NPC",
      "dr": "Disease Resistance of the NPC",
      "fr": "Fire Resistance of the NPC",
      "pr": "Poison Resistance of the NPC",
      "corrup": "Corruption Resistanceance of the NPC",
      "ph_r": "",
      "see_invis": "NPC Sees through invisibility",
      "see_invis_undead": "NPC Sees through invisibility vs Undead",
      "qglobal": "Qglobal enabled for the NPC",
      "ac": "Armor Class of the NPC",
      "npc_aggro": "NPC will aggro other NPCs",
      "spawn_limit": "Spawn Limit of the NPC",
      "attack_speed": "Attack Speed of the NPC",
      "attack_delay": "Attack Delay of the NPC",
      "findable": "NPC can be located with Find",
      "str": "Strength of the NPC",
      "sta": "Stamina of the NPC",
      "dex": "Dexterity of the NPC",
      "agi": "Agility of the NPC",
      "_int": "Intelligence of the NPC",
      "wis": "Wisdom of the NPC",
      "cha": "Charisma of the NPC",
      "see_hide": "NPC can see through hide",
      "see_improved_hide": "NPC can see through improved hide",
      "trackable": "NPC is trackable",
      "isbot": "NPC is a bot",
      "exclude": "",
      "atk": "Attack rating of the NPC",
      "accuracy": "Accuracy rating of the NPC",
      "avoidance": "Avoidance rating of the NPC",
      "slow_mitigation": "Slow Mitigation of the NPC",
      "version": "Version of the NPC",
      "maxlevel": "Maximum level of the NPC",
      "scalerate": "Scale rate of the NPC",
      "private_corpse": "",
      "unique_spawn_by_name": "NPC is a Unique Spawn By Name",
      "underwater": "Underwater NPC",
      "isquest": "The NPC has a quest script",
      "emoteid": "Emote id of the NPC",
      "spellscale": "Spell scale of the NPC",
      "healscale": "Heal scale of the NPC",
      "no_target_hotkey": "",
      "raidtarget": "NPC is a Raid Target",
      "armtexture": "Arm texture of the NPC",
      "bracertexture": "Bracer texture of the NPC",
      "handtexture": "Hand texture of the NPC",
      "legtexture": "Leg texture of the NPC",
      "feettexture": "Feet texture of the NPC",
      "light": "Light value of the NPC",
      "walkspeed": "Walkspeed of the NPC",
      "peqid": "",
      "unique": "NPC is unique",
      "fixed": "",
      "ignore_despawn": "",
      "show_name": "NPC name is shown",
      "untargetable": "NPC is untargetable",
      "charm_ac": "Charmed Armor Class of the NPC",
      "charm_min_dmg": "Charmed Minimum Damage of the NPC",
      "charm_max_dmg": "Charmed Maximum Damage of the NPC",
      "charm_attack_delay": "Charmed Attack Delay of the NPC",
      "charm_accuracy_rating": "Charmed Accuracy Rating of the NPC",
      "charm_avoidance_rating": "Charmed Avoidance Rating of the NPC",
      "charm_atk": "Charmed Attack Rating of the NPC",
      "skip_global_loot": "Skip Global Loot on this NPC",
      "rare_spawn": "NPC is a rare spawn",
      "stuck_behavior": "Behavior followed if NPC is stuck",
      "model": "Model of the NPC",
      "flymode": "Flymode for the NPC",
      "always_aggro": "NPC will always aggro",
      "exp_mod": "Experience modifier for the NPC",
    }
  }

  public static getFieldDescription(field: string) {
    // we do this because the payload we get back from spire API is
    // formatted slightly different
    let fieldLookup = field.toLowerCase().replace("_", "")

    for (let key in this.getFieldDescriptions()) {
      let keyLookup = key.toLowerCase().replace("_", "")
      if (keyLookup === fieldLookup) {
        return this.getFieldDescriptions()[key]
      }
    }

    return ''
  }

  static cacheExists(npc_id: any) {
    return this.npc[npc_id];
  }
}

