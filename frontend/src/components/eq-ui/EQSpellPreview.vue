<template>
  <div class="pb-4" style="min-width: 400px; max-width: 400px; padding: 5px" v-if="spellData">

    <div class="row">
      <div class="col-1" v-if="spellData.new_icon > 0">
        <img :src="spellCdnUrl + spellData.new_icon + '.gif'" style="width:40px;height:auto;border-radius: 10px">
      </div>
      <div class="col-11 pl-5">
        <h6 class="eq-header" style="margin: 0px; margin-bottom: 10px">
          {{ spellData.name }}
        </h6>

      </div>
    </div>

    <!-- Info -->

    <table style="width: 100%" class="mt-3">
      <tbody>
      <tr style="vertical-align:middle !important">
      </tr>
      <tr v-if="getClasses() !== ''">
        <td class="spell-field-label">Classes</td>
        <td>{{ getClasses() }}</td>
      </tr>
      <tr v-if="spellData['you_cast'] !== ''">
        <td class="spell-field-label">When you cast</td>
        <td> {{ spellData["you_cast"] }}</td>
      </tr>
      <tr v-if="spellData['other_casts'] !== ''">
        <td class="spell-field-label">When others cast</td>
        <td> {{ spellData["other_casts"] }}</td>
      </tr>
      <tr v-if="spellData['cast_on_you'] !== ''">
        <td class="spell-field-label">When cast on you</td>
        <td> {{ spellData["cast_on_you"] }}</td>
      </tr>
      <tr v-if="spellData['cast_on_other'] !== ''">
        <td class="spell-field-label">When cast on other</td>
        <td> {{ spellData["cast_on_other"] }}</td>
      </tr>
      <tr v-if="spellData['spell_fades'] !== ''">
        <td class="spell-field-label">When fading</td>
        <td> {{ spellData["spell_fades"] }}</td>
      </tr>
      <tr v-if="spellData['mana'] > 0">
        <td class="spell-field-label">Mana</td>
        <td> {{ spellData["mana"] }}</td>
      </tr>
      <tr v-if="spellData['skill'] < 52 && getDatabaseSkillName(spellData['skill']) !== ''">
        <td class="spell-field-label">Skill</td>
        <td> {{ getDatabaseSkillName(spellData["skill"]) }}</td>
      </tr>
      <tr>
        <td class="spell-field-label">Casting Time</td>
        <td> {{ (spellData["cast_time"] / 1000) }} sec</td>
      </tr>
      <tr>
        <td class="spell-field-label">Recovery Time</td>
        <td> {{ (spellData["recovery_time"] / 1000) }} sec</td>
      </tr>
      <tr>
        <td class="spell-field-label">Recast Time</td>
        <td> {{ (spellData["recast_time"] / 1000) }} sec</td>
      </tr>
      <tr>
        <td class="spell-field-label">Range</td>
        <td> {{ spellData["range"] }}</td>
      </tr>
      <tr v-if="spellData['targettype'] > 0 && getTargetTypeName(spellData['targettype']) !== ''">
        <td class="spell-field-label">Target</td>
        <td> {{ getTargetTypeName(spellData["targettype"]) }}</td>
      </tr>
      <tr v-if="spellData['resisttype'] > 0 && getSpellResistTypeName(spellData['resisttype']) !== ''">
        <td class="spell-field-label">Resist Type</td>
        <td> {{ getSpellResistTypeName(spellData["resisttype"]) }}
          <span v-if="spellData['resist_diff'] > 0">(adjust: {{ spellData["resist_diff"] }} )</span>
        </td>
      </tr>
      <tr v-if="spellData['time_of_day'] === 2">
        <td class="spell-field-label">Casting Restrictions</td>
        <td> Night Time</td>
      </tr>
      <tr v-if="getBuffDuration()">
        <td class="spell-field-label">Duration</td>
        <td> {{ humanTime(getBuffDuration() * 6) }} - {{ getBuffDuration() }} tic(s)</td>
      </tr>

      <!-- TODO: Display Reagents - the data should be passed in? -->

      <!-- TODO: Optionally Display Items that have this spell as a scroll / tome -->

      <!--      <eq-item-preview/>-->

      </tbody>
    </table>

    <!-- Spell Effects -->
    <h6 class="eq-header mt-3 mb-3" v-if="spellEffectInfo.length > 0">Effects</h6>
    <div v-if="spellEffectInfo.length > 0">
      <div v-for="effect in spellEffectInfo" style="width:100%">
        <v-runtime-template :template="'<span>' + effect + '</span>'" class="pb-6 mt-3 doc"/>
      </div>
    </div>

    <!-- Reagents -->
    <h6 class="eq-header mt-3 mb-3" v-if="reagents.length > 0">Reagents</h6>
    <div v-if="reagents.length > 0">
      <div v-for="reagent in reagents" style="width:100%">
        <div :id="reagent.id + '-' + reagent.item.id + '-' + componentId" style="display:inline-block" class="ml-2">

          <div style="display: inline-block">
            <img
              :src="itemCdnUrl + 'item_' + reagent.item.icon + '.png'"
              style="height:15px; border-radius: 25px; width:auto;"
              class="mr-2">
            <span class="mr-1">{{ reagent.item.name }}</span>
          </div>

        </div>

        <b-popover
          :target="reagent.id + '-' + reagent.item.id + '-' + componentId"
          placement="auto"
          custom-class="no-bg"
          delay="1"
          triggers="hover focus"
          style="width: 500px !important"
        >
          <eq-window style="margin-right: 10px; width: auto; height: 90%">
            <eq-item-preview :item-data="reagent.item"/>
          </eq-window>
        </b-popover>
      </div>
    </div>

    <eq-debug :data="spellData"/>
  </div>
</template>

<script>

import {App} from "@/constants/app";
import {DB_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_SKILLS} from "@/app/constants/eq-skill-constants";
import {DB_SPELL_EFFECTS, DB_SPELL_RESISTS, DB_SPELL_TARGETS} from "@/app/constants/eq-spell-constants";
import * as util from "util";
import {DB_RACE_NAMES} from "@/app/constants/eq-races-constants";
import {ItemApi, SpellsNewApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqWindow from "@/components/eq-ui/EQWindow";
import EqDebug from "@/components/eq-ui/EQDebug";

let unknowns = {}

export default {
  name: "EqSpellPreview",
  components: {
    EqWindow,
    EqDebug,
    "eq-item-preview": () => import("@/components/eq-ui/EQItemPreview.vue"),
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {
      debug: App.DEBUG,
      debugSpellEffects: false,
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      spellEffectInfo: [],
      itemData: {},
      componentId: "",
      reagents: [],
    }
  },
  created() {
    this.init()
  },
  methods: {
    async init() {
      this.spellEffectInfo = await this.getSpellEffectInfo()

      const uuidv4     = require("uuid/v4")
      this.componentId = uuidv4()

      // reagents
      let reagents = []
      for (let i = 0; i < 4; i++) {
        if (this.spellData["components_" + i] > 0) {
          let reagent  = {}
          reagent.id   = this.spellData["components_" + i]
          reagent.item = await this.getItem(this.spellData["components_" + i])
          reagents.push(reagent)
        }
      }

      this.reagents = reagents
    },
    getClasses() {
      let classData = []
      for (let i = 1; i <= 16; i++) {
        const classIndex = "classes_" + i
        if ((this.spellData[classIndex] > 0) && (this.spellData[classIndex] < 255)) {
          classData.push(DB_CLASSES[i] + " (" + this.spellData[classIndex] + ")")
        }
      }

      return classData.join(", ")
    },
    getDatabaseSkillName: function (skillId) {
      return DB_SKILLS[skillId] ? DB_SKILLS[skillId] : ""
    },
    getTargetTypeName: function (targetType) {
      return DB_SPELL_TARGETS[targetType] ? DB_SPELL_TARGETS[targetType] : "Unknown Target (" + targetType + ")"
    },
    getSpellResistTypeName: function (resist) {
      return DB_SPELL_RESISTS[resist] ? DB_SPELL_RESISTS[resist] : ""
    },
    getMinLevel: function () {
      let minLevel = 0
      for (let i = 1; i <= 16; i++) {
        const classIndex = "classes_" + i
        if ((this.spellData[classIndex] > 0) && (this.spellData[classIndex] < 255)) {
          if (this.spellData[classIndex] < minLevel) {
            minLevel = this.spellData[classIndex];
          }
        }
      }
      return minLevel
    },
    humanTime: function (sec) {
      let result = ""
      if (sec === 0) {
        result = "time";
      } else {
        let h  = Math.floor(sec / 3600);
        let m  = Math.floor((sec - h * 3600) / 60);
        let s  = sec - h * 3600 - m * 60;
        result = (h > 1 ? h + " hours " : "") + (h === 1 ? "1 hour " : "") + (m > 0 ? m + " min " : "") + (s > 0 ? s + " sec" : "");
      }

      return result;
    },
    getSpell: async function (spellId) {
      const api    = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
      const result = await api.getSpellsNew({ id: spellId })
      if (result.status === 200) {
        return result.data
      }

      return {}
    },
    getSpellName: async function (spellId) {
      const spell = await this.getSpell(spellId)
      return spell.name ? spell.name : "Unknown Spell Name"
    },
    getItem: async function (itemId) {
      const api    = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      const result = await api.getItem({ id: itemId })
      if (result.status === 200) {
        return result.data
      }

      return {}
    },
    getItemName: async function (itemId) {
      const item = await this.getItem(itemId)

      return item.name ? item.name : "Unknown Item Name"
    },
    getBuffDuration: function () {
      let i            = 0
      let minLevel     = this.getMinLevel()
      let buffDuration = this.spellData["buffduration"]

      switch (this.spellData["buffdurationformula"]) {
        case 0:
          return 0;
        case 1:
          i = Math.ceil(minLevel / 2);

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 2:
          i = Math.ceil(buffDuration / 5 * 3);

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 3:
          i = minLevel * 30;

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 4:
          return buffDuration;
        case 5:
          i = buffDuration;

          return (i < 3 ? (i < 1 ? 1 : i) : 3);
        case 6:
          i = Math.ceil(minLevel / 2);

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 7:
          i = minLevel;

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 8:
          i = minLevel + 10;

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 9:
          i = minLevel * 2 + 10;

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 10:
          i = minLevel * 3 + 10;

          return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
        case 11:
        case 12:
          return buffDuration;
        case 50:
          return 72000;
        case 3600:
          return (buffDuration ? buffDuration : 3600);
        default:
          return "???"
      }
    },
    getSpellEffectName(effectId) {
      if (!DB_SPELL_EFFECTS[effectId]) {
        if (typeof unknowns[effectId] === "undefined") {
          unknowns[effectId] = 0
        }
        unknowns[effectId]++
        // console.log(unknowns)
      }
      return DB_SPELL_EFFECTS[effectId] ? DB_SPELL_EFFECTS[effectId] : "??? (" + effectId + ")"
    },
    calcSpellEffectValue(form, base, max, lvl) {
      let sign   = 1;
      let ubase  = Math.abs(base);
      let result = 0;
      if ((max < base) && (max !== 0)) {
        sign = -1;
      }
      switch (form) {
        case 0:
        case 100:
          result = ubase;
          break;
        case 101:
          result = ubase + sign * (lvl / 2);
          break;
        case 102:
          result = ubase + sign * lvl;
          break;
        case 103:
          result = ubase + sign * lvl * 2;
          break;
        case 104:
          result = ubase + sign * lvl * 3;
          break;
        case 105:
        case 107:
          result = ubase + sign * lvl * 4;
          break;
        case 108:
          result = Math.floor(ubase + sign * lvl / 3);
          break;
        case 109:
          result = Math.floor(ubase + sign * lvl / 4);
          break;
        case 110:
          result = Math.floor(ubase + lvl / 5);
          break;
        case 111:
          result = ubase + 5 * (lvl - 16);
          break;
        case 112:
          result = ubase + 8 * (lvl - 24);
          break;
        case 113:
          result = ubase + 12 * (lvl - 34);
          break;
        case 114:
          result = ubase + 15 * (lvl - 44);
          break;
        case 115:
          result = ubase + 15 * (lvl - 54);
          break;
        case 116:
          result = Math.floor(ubase + 8 * (lvl - 24));
          break;
        case 117:
          result = ubase + 11 * (lvl - 34);
          break;
        case 118:
          result = ubase + 17 * (lvl - 44);
          break;
        case 119:
          result = Math.floor(ubase + lvl / 8);
          break;
        case 121:
          result = Math.floor(ubase + lvl / 3);
          break;

        default:
          if (form < 100) {
            result = ubase + (lvl * form);
          }
      } // end switch
      if (max !== 0) {
        if (sign === 1) {
          if (result > max) {
            result = max;
          }
        } else {
          if (result < max) {
            result = max;
          }
        }
      }
      if ((base < 0) && (result > 0)) {
        result *= -1;
      }

      return result;
    },
    getSpellEffectInfo: async function () {

      let effectsInfo = []

      // TODO: Remove
      let csv = false

      // TODO: Handle elsewhere
      let serverMaxLevel = 100;

      for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
        const spell     = this.spellData
        let printBuffer = "";
        let name        = ""
        let v           = ""

        if ((spell["effectid_" + effectIndex] !== 254) && (spell["effectid_" + effectIndex] !== 10)) {
          let maxlvl = spell["effect_base_value_" + effectIndex];
          let minlvl = serverMaxLevel;
          for (let classId = 1; classId <= 16; classId++) {
            if (spell["classes" + classId] < minlvl) {
              minlvl = spell["classes" + classId];
            }
          }

          let min = this.calcSpellEffectValue(
            spell["formula_" + effectIndex],
            spell["effect_base_value_" + effectIndex],
            spell["max_" + effectIndex],
            minlvl
          );

          let max = this.calcSpellEffectValue(
            spell["formula_" + effectIndex],
            spell["effect_base_value_" + effectIndex],
            spell["max_" + effectIndex],
            serverMaxLevel
          );

          let base_limit = spell["effect_limit_value_" + effectIndex];
          if ((min < max) && (max < 0)
          ) {
            tn  = min;
            min = max;
            max = tn;
          }

          switch (spell["effectid_" + effectIndex]) {
            case 3: // Increase Movement (% / 0)
              if (max < 0) { // Decrease
                printBuffer += "Decrease Movement";
                if (min !== max) {
                  printBuffer += " by " + Math.abs(min) + "% (L" + minlvl + ") to " + Math.abs(max) + "% (L" + maxlvl + ")";
                } else {
                  printBuffer += " by " + Math.abs(100) + "%";
                }
              } else {
                printBuffer += "Increase Movement";
                if (min !== max) {
                  printBuffer += " by " + min + "% (L" + minlvl + ") to " + (max) + "% (L" + maxlvl + ")";
                } else {
                  printBuffer += " by " + (max) + "%";
                }
              }
              break;
            case 11: // Decrease OR Inscrease AttackSpeed (max/min = percentage of speed / normal speed, IE, 70=>-30% 130=>+30%
              if (max < 100) { // Decrease
                printBuffer += "Decrease Attack Speed";
                if (min !== max) {
                  printBuffer += " by " + (100 - min) + "% (L" + minlvl + ") to " + (100 - max) + "% (L" + maxlvl + ")";
                } else {
                  printBuffer += " by " + (100 - max) + "%";
                }
              } else {
                printBuffer += "Increase Attack Speed";
                if (min !== max) {
                  printBuffer += " by " + (min - 100) + "% (L" + minlvl + ") to " + (max - 100) + "% (L" + maxlvl + ")";
                } else {
                  printBuffer += " by " + (max - 100) + "%";
                }
              }
              break;
            case 21: // stun
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (min !== max) {
                printBuffer += " (" + (min / 1000) + " sec (L" + minlvl + ") to " + (max / 1000) + " sec (L" + maxlvl + "))";
              } else {
                printBuffer += " (" + (max / 1000) + " sec)";
              }
              break;
            case 32: // summonitem
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);

              const item             = (await this.getItem(spell["effect_base_value_" + effectIndex]));
              this.itemData[item.id] = item

              if (item.name) {
                printBuffer += `
                <div :id="${effectIndex} + '-' + ${item.id} + '-' + componentId" style="display:inline-block" class="ml-2">

                  <div style="display: inline-block">
                    <img
                      :src="itemCdnUrl + 'item_' + ${item.icon} + '.png'"
                      style="height:15px; border-radius: 25px; width:auto;"
                      class="mr-1">
                    <span class="mr-1">${item.name}</span>
                  </div>

                </div>

                <b-popover
                  :target="${effectIndex} + '-' + ${item.id} + '-' + componentId"
                  placement="auto"
                  custom-class="no-bg"
                  delay="1"
                  triggers="hover focus"
                  style="width: 500px !important"
                >
                  <eq-window style="margin-right: 10px; width: auto; height: 90%">
                    <eq-item-preview :item-data="itemData[${item.id}]"/>
                  </eq-window>
                </b-popover>`
              }
              break;
            case 87: // Increase Magnification
            case 98: // Increase Haste v2
            case 114: // Increase Agro Multiplier
            case 119: // Increase Haste v3
            case 123: // Increase Spell Damage
            case 124: // Increase Spell Damage
            case 125: // Increase Spell Healing
            case 127: // Increase Spell Haste
            case 128: // Increase Spell Duration
            case 129: // Increase Spell Range
            case 130: // Decrease Spell/Bash Hate
            case 131: // Decrease Chance of Using Reagent
            case 132: // Decrease Spell Mana Cost
            case 158: // Increase Chance to Reflect Spell
            case 168: // Increase Melee Mitigation
            case 169: // Increase Chance to Critical Hit
            case 172: // Increase Chance to Avoid Melee
            case 173: // Increase Chance to Riposte
            case 174: // Increase Chance to Dodge
            case 175: // Increase Chance to Parry
            case 176: // Increase Chance to Dual Wield
            case 177: // Increase Chance to Double Attack
            case 180: // Increase Chance to Resist Spell
            case 181: // Increase Chance to Resist Fear Spell
            case 183: // Increase All Skills Skill Check
            case 184: // Increase Chance to Hit With all Skills
            case 185: // Increase All Skills Damage Modifier
            case 186: // Increase All Skills Minimum Damage Modifier
            case 188: // Increase Chance to Block
            case 200: // Increase Proc Modifier
            case 201: // Increase Range Proc Modifier
            case 216: // Increase Accuracy
            case 227: // Reduce Skill Timer
            case 266: // Add Attack Chance
            case 273: // Increase Critical Dot Chance
            case 294: // Increase Critical Spell Chance
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (min !== max) {
                printBuffer += " by " + min + "% (L" + minlvl + ") to " + max + "% (L" + maxlvl + ")";
              } else {
                printBuffer += " by " + max + "%";
              }
              break;
            case 15: // Increase Mana per tick
            case 100: // Increase Hitpoints v2 per tick
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              let duration = spell["buffduration"]
              if (min !== max) {
                printBuffer += " by " + Math.abs(min) + " (L" + minlvl + ") to " + Math.abs(
                  max
                ) + " (L" + maxlvl + ") per tick (total " + Math.abs(min * duration) + " to " + Math.abs(
                  max * duration
                ) + ")";
              } else {
                printBuffer += " by " + max + " per tick (total " + Math.abs(max * duration) + ")";
              }
              break;
            case 30: // Frenzy Radius
            case 86: // Reaction Radius
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + spell["effect_base_value_" + effectIndex] + "/" + spell["effect_limit_value_" + effectIndex] + ")";
              break;
            case 22: // Charm
            case 23: // Fear
            case 31: // Mesmerize
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " up to level " + spell["effect_limit_value_" + effectIndex];
              break;
            case 33: // Summon Pet:
            case 68: // Summon Skeleton Pet:
            case 106: // Summon Warder:
            case 108: // Summon Familiar:
            case 113: // Summon Horse:
            case 152: // Summon Pets:
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " <a href=?a=pet&name=" + spell["teleport_zone"] + ">" + spell["teleport_zone"] + "</a>";
              break;
            case 13: // See Invisible
            case 18: // Pacify
            case 20: // Blindness
            case 25: // Bind Affinity
            case 26: // Gate
            case 28: // Invisibility versus Undead
            case 29: // Invisibility versus Animals
            case 40: // Invunerability
            case 41: // Destroy Target
            case 42: // Shadowstep
            case 44: // Lycanthropy
            case 52: // Sense Undead
            case 53: // Sense Summoned
            case 54: // Sense Animals
            case 56: // True North
            case 57: // Levitate
            case 61: // Identify
            case 64: // SpinStun
            case 65: // Infravision
            case 66: // UltraVision
            case 67: // Eye of Zomm
            case 73: // Bind Sight
            case 74: // Feign Death
            case 75: // Voice Graft
            case 76: // Sentinel
            case 77: // Locate Corpse
            case 82: // Summon PC
            case 90: // Cloak
            case 93: // Stop Rain
            case 94: // Make Fragile (Delete if combat)
            case 95: // Sacrifice
            case 96: // Silence
            case 99: // Root
            case 101: // Complete Heal (with duration)
            case 103: // Call Pet
            case 104: // Translocate target to their bind point
            case 105: // Anti-Gate
            case 115: // Food/Water
            case 117: // Make Weapons Magical
            case 135: // Limit: Resist(Magic allowed)
            case 137: // Limit: Effect(Hitpoints allowed)
            case 138: // Limit: Spell Type(Detrimental only)
            case 141: // Limit: Instant spells only
            case 150: // Death Save - Restore Full Health
            case 151: // Suspend Pet - Lose Buffs and Equipment
            case 154: // Remove Detrimental
            case 156: // Illusion: Target
            case 178: // Lifetap from Weapon Damage
            case 179: // Instrument Modifier
            case 182: // Hundred Hands Effect
            case 194: // Fade
            case 195: // Stun Resist
            case 205: // Rampage
            case 206: // Area of Effect Taunt
            case 311: // Limit: Combat Skills Not Allowed
            case 314: // Fixed Duration Invisbility
            case 299: // Wake the Dead
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              break;
            case 58: // Illusion:
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += DB_RACE_NAMES[spell["effect_base_value_" + effectIndex]];
              break;
            case 63: // Memblur
            case 120: // Set Healing Effectiveness
            case 330: // Critical Damage Mob
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + max + "%)";
              break;
            case 81: // Resurrect
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " and restore " + spell["effect_base_value_" + effectIndex] + "% experience";
              break;
            case 83: // Teleport
            case 88: // Evacuate
            case 145: // Teleport v2
              //print_buffer += " (Need to add zone to spells table)";
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " <a href=?a=zone&name=" + spell["teleport_zone"] + ">" + spell["teleport_zone"] + "</a>";
              break;
            case 85: // Add Proc:
            case 289: // Improved Spell Effect:
            case 323: // Add Defensive Proc:
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);

              printBuffer += "<a href=?a=spell&id=" + spell["effect_base_value_" + effectIndex] + "> " + (await this.getSpellName(spell["effect_base_value_" + effectIndex])) + "</a>";
              break;
            case 89: // Increase Player Size
              name = this.getSpellEffectName(spell["effectid_" + effectIndex]);
              min -= 100;
              max -= 100;
              if (max < 0) {
                name = name.replace("Increase", "Decrease");
              }
              printBuffer += name;
              if (min !== max) {
                printBuffer += " by " + min + "% (L" + minlvl + ") to " + " + max + " + "% (L" + maxlvl + ")";
              } else {
                printBuffer += " by " + max + "%";
              }
              break;
            case 27: // Cancel Magic
            case 134: // Limit: Max Level
            case 157: // Spell-Damage Shield
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + max + ")";
              break;
            case 121: // Reverse Damage Shield
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (-" + max + ")";
              break;
            case 91: // Summon Corpse
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (max level " + max + ")";
              break;
            case 136: // Limit: Target
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (max < 0) {
                max = -max;
                v   = " excluded";
              } else {
                v = "";
              }
              printBuffer += " (" + DB_SPELL_TARGETS[max] + " " + v + ")";
              break;
            case 139: // Limit: Spell
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              max = spell["effect_base_value_" + effectIndex];
              if (max < 0) {
                max = -max;
                v   = " excluded";
              }

              name = (await this.getSpellName(max))
              if (csv === false) {
                printBuffer += "(" + name + ")";
              } else {
                printBuffer += " (<a href=?a=spell&id=" + spell["effect_base_value_" + effectIndex] + ">" + name + "</a>v)";
              }
              break;
            case 140: // Limit: Min Duration
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              min *= 6;
              max *= 6;
              if (min !== max) {
                printBuffer += " (" + min + " sec (L" + minlvl + ") to " + max + " sec (L" + maxlvl + "))";
              } else {
                printBuffer += " (max sec)";
              }
              break;
            case 143: // Limit: Min Casting Time
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              min *= 6;
              max *= 6;
              if (min !== max) {
                printBuffer += " (" + (min / 6000) + " sec (L" + minlvl + ") to " + (max / 6000) + " sec (L" + maxlvl + "))";
              } else {
                printBuffer += " (" + (max / 6000) + " sec)";
              }
              break;
            case 148: // Stacking: Overwrite existing spell
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " if slot " + (spell["formula_" + effectIndex] - 200) + " is effect '" + this.getSpellEffectName(spell["effect_base_value_" + effectIndex]) + "' and < " + spell["max_" + effectIndex];
              break;
            case 149: // Stacking: Overwrite existing spell
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " if slot " + (spell["formula_" + effectIndex] - 200) + " is effect '" + this.getSpellEffectName(spell["effect_base_value_" + effectIndex]) + "' and < " + spell["max_" + effectIndex];
              break;
            case 147: // Increase Hitpoints (%)
              name = this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (max < 0) {
                name = name.replace("Increase", "Decrease");
              }
              printBuffer += name + " by " + spell["effect_limit_value_" + effectIndex] + " (" + max + "% " + max + ")";
              break;
            case 153: // Balance Party Health
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + max + "% penalty)";
              break;
            case 0: // In/Decrease hitpoints
            case 1: // Increase AC
            case 2: // Increase ATK
            case 4: // Increase STR
            case 5: // Increase DEX
            case 6: // Increase AGI
            case 7: // Increase STA
            case 8: // Increase INT
            case 9: // Increase WIS
            case 19: // Increase Faction
            case 35: // Increase Disease Counter
            case 36: // Increase Poison Counter
            case 46: // Increase Magic Fire
            case 47: // Increase Magic Cold
            case 48: // Increase Magic Poison
            case 49: // Increase Magic Disease
            case 50: // Increase Magic Resist
            case 55: // Increase Absorb Damage
            case 59: // Increase Damage Shield
            case 69: // Increase Max Hitpoints
            case 78: // Increase Absorb Magic Damage
            case 79: // Increase HP when cast
            case 92: // Increase hate
            case 97: // Increase Mana Pool
            case 111: // Increase All Resists
            case 112: // Increase Effective Casting
            case 116: // Decrease Curse Counter
            case 118: // Increase Singing Skill
            case 159: // Decrease Stats
            case 167: // Pet Power Increase
            case 192: // Increase hate
            default:
              name = this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (max < 0) {
                name = name.replace("Increase", "Decrease");
              }
              printBuffer += name;
              if (min !== max) {
                printBuffer += " by " + min + " (L" + minlvl + ") to " + max + " (L" + maxlvl + ")";
              } else {
                if (max < 0) {
                  max = -max;
                }
                printBuffer += " by " + max + "";
              }

              break;
          }


          if (printBuffer !== "") {
            effectsInfo.push(effectIndex + ") " + printBuffer)
          }

          if (this.debugSpellEffects && printBuffer !== "") {
            const debug = util.format(
              "--- Debug: Effect ID (%s) (%s) Index (%s)",
              spell["effectid_" + effectIndex],
              DB_SPELL_EFFECTS[spell["effectid_" + effectIndex]] ? DB_SPELL_EFFECTS[spell["effectid_" + effectIndex]] : "UNKNOWN",
              effectIndex
            )
            effectsInfo.push(debug)
          }

        }

      }

      return effectsInfo;
    }
  },
  props: {
    spellData: Object
  }
}
</script>

<style>
.spell-field-label {
  text-align:    right;
  font-weight:   bold;
  width:         40%;
  padding-right: 10px;
}
</style>
