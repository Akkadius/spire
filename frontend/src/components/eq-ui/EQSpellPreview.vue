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
      <tr v-if="spellData['id'] !== ''">
        <td class="spell-field-label">Spell ID: </td>
        <td> {{ spellData["id"] }}</td>
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
import {DB_SKILLS, DB_BARD_SKILLS} from "@/app/constants/eq-skill-constants";
import {DB_SPELL_EFFECTS, DB_SPA, DB_SPELL_RESISTS, DB_SPELL_TARGETS, DB_SPELL_TARGET_RESTRICTION, DB_SPELL_WORN_ATTRIBUTE_CAP, DB_SPELL_PETCMDS} from "@/app/constants/eq-spell-constants";
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

 // (int calc, int base1, int max, int tick, int level = MAX_LEVEL)
    calcSpellEffectValue2(calc, base1, max, tick, level) {

      if (calc == 0) {
        return base1;
      }

      if (calc == 100)
      {
        if (max > 0 && base1 > max) {
          return max;
        }
        return base1;
      }

      let change = 0;

      switch (calc)
      {
        case 100:
          break;

        case 101:
          change = level / 2;
          break;
        case 102:
          change = level;
          break;
        case 103:
          change = level * 2;
          break;
        case 104:
          change = level * 3;
          break;
        case 105:
          change = level * 4;
          break;
        case 107:
          change = -1 * tick;
          break;
        case 108:
          change = -2 * tick;
          break;
        case 109:
          change = level / 4;
          break;
        case 110:
          change = level / 6;
          break;
        case 111:
          if (level > 16) change = (level - 16) * 6;
          break;
        case 112:
          if (level > 24) change = (level - 24) * 8;
          break;
        case 113:
          if (level > 34) change = (level - 34) * 10;
          break;
        case 114:
          if (level > 44) change = (level - 44) * 15;
          break;
        case 115:
          if (level > 15) change = (level - 15) * 7;
          break;
        case 116:
          if (level > 24) change = (level - 24) * 10;
          break;
        case 117:
          if (level > 34) change = (level - 34) * 13;
          break;
        case 118:
          if (level > 44) change = (level - 44) * 20;
          break;
        case 119:
          change = level / 8;
          break;
        case 120:
          change = -5 * tick;
          break;
        case 121:
          change = level / 3;
          break;
        case 122:
          change = -12 * tick;
          break;
        case 123:
          change = (Math.abs(max) - Math.abs(base1)) / 2;
          break;
        case 124:
          if (level > 50) change = (level - 50);
          break;
        case 125:
          if (level > 50) change = (level - 50) * 2;
          break;
        case 126:
          if (level > 50) change = (level - 50) * 3;
          break;
        case 127:
          if (level > 50) change = (level - 50) * 4;
          break;
        case 128:
          if (level > 50) change = (level - 50) * 5;
          break;
        case 129:
          if (level > 50) change = (level - 50) * 10;
          break;
        case 130:
          if (level > 50) change = (level - 50) * 15;
          break;
        case 131:
          if (level > 50) change = (level - 50) * 20;
          break;
        case 132:
          if (level > 50) change = (level - 50) * 25;
          break;
        case 139:
          if (level > 30) change = (level - 30) / 2;
          break;
        case 140:
          if (level > 30) change = (level - 30);
          break;
        case 141:
          if (level > 30) change = 3 * (level - 30) / 2;
          break;
        case 142:
          if (level > 30) change = 2 * (level - 60);
          break;
        case 143:
          change = 3 * level / 4;
          break;

        case 3000:
          // todo: this appears to be scaled by the targets level
          // base1 value how it affects a level 100 target
          return base1;

        default:
          if (calc > 0 && calc < 1000) {
            change = level * calc;
          }

          if (calc >= 1000 && calc < 2000) {
            change = tick * (calc - 1000) * -1;
          }

          if (calc >= 2000 && calc < 3000) {
            change = level * (calc - 2000);
          }

          if (calc >= 4000 && calc < 5000) {
            change = -tick * (calc - 4000);
          }

          break;
      }

      let value = Math.abs(base1) + change;

      if (max != 0 && value > Math.abs(max)) {
        value = Math.abs(max);
      }

      if (base1 < 0) {
        value = -value;
      }

      return value;
    },

    CalcValueRange(calc, base1, max, spa, duration, level)
    {
      let printBuffer = ""
      let start = this.calcSpellEffectValue2(calc, base1, max, 1, level);
      let finish = Math.abs(this.calcSpellEffectValue2(calc, base1, max, duration, level));

      let type = Math.abs(start) < Math.abs(finish) ? "Growing" : "Decaying";

      if (calc == 123){
        if (base1 < 0){
          max = max * -1;
        }
        printBuffer = " (Random: " + Math.abs(base1) + " to " + Math.abs(max) +  ")"
      }

      if (calc == 107) {
        printBuffer = " (" + type + " to " + finish + " @ 1/tick)"
      }

      if (calc == 108) {
        printBuffer = " (" + type + " to " + finish + " @ 2/tick)"
      }

      if (calc == 120) {
        printBuffer = " (" + type + " to " + finish + " @ 5/tick)"
      }

      if (calc == 122) {
        printBuffer = " (" + type + " to " + finish + " @ 12/tick)"
      }

      if (calc > 1000 && calc < 2000) {
        printBuffer = " (" + type + " to " + finish + " @ " + (calc - 1000) + "/tick)"
      }

      if (calc >= 3000 && calc < 4000)
      {
        if (calc - 3000 == spa)
        {
          printBuffer = " (Scales, Base Level: 100)";
        }
        if (calc - 3500 == spa)
        {
          printBuffer = " (Scales, Base Level: 105)";
        }
      }

      if (calc > 4000 && calc < 5000) {
        printBuffer = " (" + type + " to " + finish + " @ " + (calc - 4000) + "/tick)"
      }

      return printBuffer;
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

    getFormatStandard(effect_name, type, value_min, value_max, minlvl, maxlvl){

      let modifier = ""

      if (value_max < 0){
        modifier = "Decrease "
      }
      else{
        modifier = "Increase "
      }

      let printBuffer = modifier + effect_name

      if (value_min !== value_max) {
        printBuffer += " by " + Math.trunc(Math.abs(value_min))+ type + " (L" + minlvl + ") to " + Math.trunc(Math.abs(value_max) )+ type + " (L" + maxlvl + ")";
      } else {
        printBuffer += " by " + Math.trunc(Math.abs(value_max)) + type
      }

      return printBuffer;
    },

    getUpToMaxLvl(max){

      let printBuffer = ""
      if (max > 0){
        printBuffer = " up to level " + max
      }
      return printBuffer;
    },

    getFocusPercentRange(effect_name, min, max, negate) {

      let printBuffer = ""
      let modifier = ""

      if (min < 0)  {

        if (min < max) {
          let temp = min;
          min = max;
          max = temp;
        }
      }
      else {
        if (min > max)
          max = min;
      }

      if (negate) {
        min = -min;
        max = -max;
      }

      if (max < 0){
        modifier = "Decrease "
      }
      else{
        modifier = "Increase "
      }

     if (min == max || max == 0) {
       return printBuffer += modifier + effect_name + " by " + Math.abs(min) + "%";
     }
     return printBuffer += modifier + effect_name + " by " + Math.abs(min) + "% to " + Math.abs(max) + "%";
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
        let tmp         = ""
        let tmp2        = ""
        let pertick = spell["buffduration"] ? " per tick " : ""

        let base = spell["effect_base_value_" + effectIndex]
        let limit = spell["effect_limit_value_" + effectIndex]
        let max = spell["max_" + effectIndex]


        if (spell["effectid_" + effectIndex] !== 254) {

          //let maxlvl = spell["effect_base_value_" + effectIndex];
          let maxlvl = serverMaxLevel;
          let minlvl = 1; // make this 255; FIX THIS

          for (let classId = 1; classId <= 16; classId++) {
            if (spell["classes" + classId] < minlvl) {
              minlvl = spell["classes" + classId];
            }
          }

          /* OLD
          let value_min = this.calcSpellEffectValue(
            spell["formula_" + effectIndex],
            spell["effect_base_value_" + effectIndex],
            spell["max_" + effectIndex],
            minlvl
          );

          let value_max = this.calcSpellEffectValue(
            spell["formula_" + effectIndex],
            spell["effect_base_value_" + effectIndex],
            spell["max_" + effectIndex],
            serverMaxLevel
          );
          */

          let value_min = this.calcSpellEffectValue2( spell["formula_" + effectIndex], base, max, 1,minlvl);
          let value_max = this.calcSpellEffectValue2( spell["formula_" + effectIndex], base, max, 1,serverMaxLevel);


          if ((value_min < value_max) && (value_max < 0)
          ) {
            let tn  = value_min;
            value_min = value_max;
            value_max = tn;
          }

          let special_range = this.CalcValueRange(spell["formula_" + effectIndex], base, max, spell["effectid_" + effectIndex],spell["buffduration"],serverMaxLevel)

          switch (spell["effectid_" + effectIndex]) {

            case 0:
              tmp += limit ? " (" + DB_SPELL_TARGET_RESTRICTION[Math.abs(limit)] + ")" : ""
              printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range + tmp
              break;

            case 1:
              printBuffer += this.getFormatStandard("AC", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 2:
              printBuffer += this.getFormatStandard("ATK", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 3:
              printBuffer += this.getFormatStandard("Movement Speed", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 4:
              printBuffer += this.getFormatStandard("STR", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 5:
              printBuffer += this.getFormatStandard("DEX", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 6:
              printBuffer += this.getFormatStandard("AGI", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 7:
              printBuffer += this.getFormatStandard("STA", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 8:
              printBuffer += this.getFormatStandard("INT", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 9:
              printBuffer += this.getFormatStandard("WIS", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 10:
              if (base == 0 && (spell["formula_" + effectIndex] == 100)){ //This is used as a placeholder do not calculate
                printBuffer = ""
              }
              else {
                printBuffer += this.getFormatStandard("CHA", "", value_min, value_max, minlvl, maxlvl);
              }
              break;

            case 11: // Slow 70=>-30%, Haste 130=>+30%
              if (base < 100){
                value_max = (100 - value_max) * -1;
                value_min = (100 - value_min) * -1;
              }
              else {
                value_max = value_max - 100;
                value_min = value_min - 100;

              }
              printBuffer += this.getFormatStandard("Attack Speed", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 12: //note: eqemu does not support base1 "enhanced invisibility" value
              printBuffer += "Invisibility (Unstable)"
              break;

            case 13: //note: eqemu does not support base1 "enhanced see invisibility" value
              printBuffer += "See Invisible"
              break;

            case 14:
              printBuffer += "Enduring Breath"
              break;

            case 15:
              printBuffer += this.getFormatStandard("Current Mana", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range
              break;

            case 16:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 17:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 18:
              printBuffer += "Pacify"
              break;

            case 19:
              printBuffer += this.getFormatStandard("Faction", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 20:
              printBuffer += "Blind"
              break;

            case 21:
              let pvpstun = ""
              if (base !== limit && limit !== 0){
                pvpstun += " ( " + (limit/1000) + " in PvP)"
              }

              printBuffer += "Stun for " + (base / 1000) + " sec" + pvpstun + this.getUpToMaxLvl(max)
              break;

            case 22:
              printBuffer += "Charm" + this.getUpToMaxLvl(max)
              break;

            case 23:
              printBuffer += "Fear" + this.getUpToMaxLvl(max)
              break;

            case 24:
              printBuffer += this.getFormatStandard("Stamina Loss", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 25:
              let bindtype = ""
              if (base == 2){
                bindtype += " (Secondary Bind Point)"
              }
              else if (base == 3){
                bindtype += " (Tertiary Bind Point)"
              }
              printBuffer += "Bind" + bindtype
              break;

            case 26:
              let gatetype = ""
               if (limit == 2){
                gatetype += " to Secondary Bind Point "
              }
              else if (limit == 3){
                gatetype += " to Tertiary Bind Point "
              }

              printBuffer += "Gate" + gatetype + " (" + (100 - base) + "% chance to fail)"
              break;

            case 27:
              printBuffer += "Dispel" + " with level modifier of " + " (" + base + ")"
              break;

            case 28:
              printBuffer += "Invisibility to Undead (Unstable)"
              break;

            case 29:
              printBuffer += "Invisibility to Animals (Unstable)"
              break;

            case 30:
              printBuffer += "Decrease Aggro Radius to " + base + this.getUpToMaxLvl(max)
              break;

            case 31:
              printBuffer += "Mesmerize" + this.getUpToMaxLvl(max)
              break;

            case 32: //Review this
              printBuffer += "Summon Item: "

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

            case 33:
              printBuffer += "Summon Pet: " + spell["teleport_zone"]
              break;

            case 34:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 35:
              printBuffer += this.getFormatStandard("Disease Counter", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 36:
              printBuffer += this.getFormatStandard("Poison Counter", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 37:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 38:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 39:
              printBuffer += "Can not be Twincast"
              break;

            case 40:
              printBuffer += "Invulnerability"
              break;

            case 41:
              printBuffer += "Destroy"
              break;

            case 42: //handled by client, unknown what base value represents
              printBuffer += "Shadowstep"
              break;

            case 43: //custom on eqemu, not used on live. Any client with this effect now has chance to crippling blow
              printBuffer += "Berserk: Allows chance to crippling blow"
              break;

            case 44: //TODO This is a type of buff stacker that I need to figure out
              printBuffer += "Lycanthropy: Need to implement on Eqemu"
              break;

            case 45: //custom on eqemu, not used on live. Stackable melee lifesteal effect.
              printBuffer += "Lifetap from Weapon Damage: " + base + "%"
              break;

            case 46:
              printBuffer += this.getFormatStandard("Fire Resist", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 47:
              printBuffer += this.getFormatStandard("Cold Resist", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 48:
              printBuffer += this.getFormatStandard("Poison Resist", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 49:
              printBuffer += this.getFormatStandard("Disease Resist", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 50:
              printBuffer += this.getFormatStandard("Magic Resist", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 51:
              printBuffer += "Sense Undead"
              break;

            case 52:
              printBuffer += "Sense Summoned"
              break;

            case 53:
              printBuffer += "Sense Animal"
              break;

            case 54:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 55: //rune
              printBuffer += "Absorb Damage: 100%, Total: " + base
              break;

            case 56:
              printBuffer += "True North"
              break;

            case 57:
              printBuffer += "Levitate"
              break;

            case 58: //TODO NEED FINISH THIS
              printBuffer += "Illusion: "
              break;

            case 59:
              printBuffer += this.getFormatStandard("Damage Shield", "", -value_min, -value_max, minlvl, maxlvl);
              break;

            case 60:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 61:
              printBuffer += "Identify Item"
              break;

            case 62:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 63:
              printBuffer += "Memory Blur" + " (" + base + "% chance)"
              break;

            case 63:
              printBuffer += "Memory Blur" + " (" + base + "% chance)"
              break;

            case 64:
              if (base !== limit && limit !== 0){
                printBuffer += "Stun and Spin NPC for " + (base / 1000) + " sec (PC for "  + (limit/ 1000) + " sec " + this.getUpToMaxLvl(max)
              }
              else {
                printBuffer += "Stun and Spin for " + (base / 1000) + " sec "+ this.getUpToMaxLvl(max)
              }
              break;

            case 65:
              printBuffer += "Infravision"
              break;

            case 66:
              printBuffer += "Ultravision"
              break;

            case 67:
              printBuffer += "Eye of Zomm"
              break;

            case 68:
              printBuffer += "Reclaim Pet Mana"
              break;

            case 69:
              printBuffer += this.getFormatStandard("Max HP", "", value_min, value_max, minlvl, maxlvl) + special_range
              break;

            case 70:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 71:
              printBuffer += "Summon Pet: " + spell["teleport_zone"]
              break;

            case 72:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 73:
              printBuffer += "Bind Sight"
              break;

            case 74:
              printBuffer += "Feign Death (" + base + "% chance)"
              break;

            case 75:
              printBuffer += "Project Voice"
              break;

            case 76:
              printBuffer += "Sentinel"
              break;

            case 77:
              printBuffer += "Locate Corpse"
              break;

            case 78: //spell rune
              printBuffer += "Absorb Spell Damage: 100%, Total: " + base
              break;

            case 79:
              tmp += limit ? " (" + DB_SPELL_TARGET_RESTRICTION[Math.abs(limit)] + ")" : ""
              printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl) + special_range + tmp
              break;

            case 80:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 81:
              printBuffer += "Resurrect with " + base + "% XP"
              break;

            case 82:
              printBuffer += "Summon Player"
              break;

            case 83: //TODO teleport
              //return String.Format("Teleport to {0}", Extra);
              break;

            case 84: // base on emu does damage. Correct?
              printBuffer += "Gravity Flux"
              break;

            case 85: //TODO Spell proc LINK
              if (limit != 0){
                tmp += " with " + limit + " % Rate Mod"
              }
              printBuffer += "Add Melee Proc: " + (await this.getSpellName(spell["effect_base_value_" + effectIndex])) + tmp
              break;

            case 86:
              printBuffer += "Decrease Social Radius to " + base +  + this.getUpToMaxLvl(max)
              break;

            case 87:
              printBuffer += this.getFormatStandard("Magnification", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 88: //TODO clean up, enum for zones
              if (spell["teleport_zone"] != "same"){
                tmp += " (" + spell["effect_base_value_" + (effectIndex + 1)]
                + ", " + spell["effect_base_value_" + effectIndex] + ", "
                + spell["effect_base_value_" + (effectIndex + 2)] + ", "
                + spell["effect_base_value_" + (effectIndex + 3)] + ")"
              }
              printBuffer += "Evacuate to " + spell["teleport_zone"] + tmp
              break;

            case 89:
              if (base < 100){
                value_max = (100 - value_max) * -1;
                value_min = (100 - value_min) * -1;
              }
              else {
                value_max = value_max - 100;
                value_min = value_min - 100;

              }
              printBuffer += this.getFormatStandard("Player Size", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 90: // pet invisible - This is not implemented on eqemu
              printBuffer += "Ignore Pet (not implemented)"
              break;

            case 91:
              printBuffer += "Summon Corpse up to level " + base
              break;

            case 92:
              printBuffer += this.getFormatStandard("Hate", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 93:
              printBuffer += "Stop Rain"
              break;

             case 94:
              printBuffer += "Cancel if Combat Initiated"
              break;

            case 95:
              printBuffer += "Sacrifice"
              break;

            case 96:
              printBuffer += "Inhibit Spell Casting (Silence)"
              break;

            case 97:
              printBuffer += this.getFormatStandard("Max Mana", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 98:
              if (base < 100){
                value_max = (100 - value_max) * -1;
                value_min = (100 - value_min) * -1;
              }
              else {
                value_max = value_max - 100;
                value_min = value_min - 100;

              }
              printBuffer += this.getFormatStandard("Attack Speed (v2 capped)", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 99:
              printBuffer += "Root"
              break;

            case 100:
              tmp += limit ? " (" + DB_SPELL_TARGET_RESTRICTION[Math.abs(limit)] + ")" : ""
              printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl)  + pertick + special_range + tmp
              break;

            case 101:
              printBuffer += "Increase Current HP by " + (base * 7500) + " with recast blocking buff"
              break;

            case 102:
              printBuffer += "Fear Immunity"
              break;

            case 103:
              printBuffer += "Summon Pet to Player"
              break;

            case 104: //TODO clean up, enum for zones
              if (spell["teleport_zone"] != ""){
                tmp += spell["teleport_zone"] +
                  " (" + spell["effect_base_value_" + (effectIndex + 1)]
                  + ", " + spell["effect_base_value_" + effectIndex] + ", "
                  + spell["effect_base_value_" + (effectIndex + 2)] + ", "
                  + spell["effect_base_value_" + (effectIndex + 3)] + ")"
              }
              else{
                tmp += "bind"
              }
              printBuffer += "Translocate to " + tmp
              break;

            case 105:
              printBuffer += "Inhibit Gate";
              return;

            case 106:
              printBuffer += "Summon Warder: " + spell["teleport_zone"]
              break;

            case 107: //not on live
              printBuffer += this.getFormatStandard("NPC Level", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 108: //limit "Ignore Auto Leave" not supported, not in current spell file
              printBuffer += "Summon Familiar:"  + spell["teleport_zone"]
              break;

            case 109: //later expansions allow stacks to put into bags using limit value.
              printBuffer += "Summon into Bag: "

              const item2             = (await this.getItem(spell["effect_base_value_" + effectIndex]));
              this.itemData[item2.id] = item2

              if (item2.name) {
                printBuffer += `
                <div :id="${effectIndex} + '-' + ${item2.id} + '-' + componentId" style="display:inline-block" class="ml-2">

                  <div style="display: inline-block">
                    <img
                      :src="itemCdnUrl + 'item_' + ${item2.icon} + '.png'"
                      style="height:15px; border-radius: 25px; width:auto;"
                      class="mr-1">
                    <span class="mr-1">${item2.name}</span>
                  </div>

                </div>

                <b-popover
                  :target="${effectIndex} + '-' + ${item2.id} + '-' + componentId"
                  placement="auto"
                  custom-class="no-bg"
                  delay="1"
                  triggers="hover focus"
                  style="width: 500px !important"
                >
                  <eq-window style="margin-right: 10px; width: auto; height: 90%">
                    <eq-item-preview :item-data="itemData[${item2.id}]"/>
                  </eq-window>
                </b-popover>`
              }
              break;

            case 110:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 111:
              printBuffer += this.getFormatStandard("All Resists", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 112:
              printBuffer += this.getFormatStandard("Effective Casting Level", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 113:
              printBuffer += "Summon Mount: " + spell["teleport_zone"]
              break;

            case 114:
              printBuffer += this.getFormatStandard("Hate Generated", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 115:
              printBuffer += "Reset Hunger Counter"
              break;

            case 116:
                 printBuffer += this.getFormatStandard("Curse Counter", "", value_min, value_max, minlvl, maxlvl);
                 break;

            case 117:
              return "Make Weapon Magical";
              break;

            case 118:
              printBuffer += this.getFormatStandard("Singing Amplification", "%", (value_min * 10), (value_max * 10), minlvl, maxlvl);
              break;

            case 119:
              if (base < 100){
                value_max = (100 - value_max) * -1;
                value_min = (100 - value_min) * -1;
              }
              else {
                value_max = value_max - 100;
                value_min = value_min - 100;

              }
              printBuffer += this.getFormatStandard("Attack Speed (v3 over cap)", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 120:
              printBuffer += this.getFormatStandard("Healing Taken", "%", value_min, value_max, minlvl, maxlvl);
              break;

            case 121:
              printBuffer += this.getFormatStandard("Reverse Damage Shield", "", value_min, value_max, minlvl, maxlvl);
              break;

            case 122: //TODO implement this
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not implemented"
                break;

            case 123:
              printBuffer += "Buff Blocker: Screech (" + base + ")"
              break;

            case 124:
              printBuffer += this.getFocusPercentRange("Spell Damage", base, limit, false);
              break;

            case 125:
              printBuffer += this.getFocusPercentRange("Healing", base, limit, false);
              break;

            case 126:
              printBuffer += this.getFocusPercentRange("Spell Resist Rate", base, limit, true);
              break;

            case 127:
              printBuffer += this.getFocusPercentRange("Spell Haste", base, limit, false);
              break;

            case 128:
              printBuffer += this.getFocusPercentRange("Spell Duration", base, limit, false);
              break;
;
            case 129:
              printBuffer += this.getFocusPercentRange("Spell Range", base, limit, false);
              break;

            case 130:
              printBuffer += this.getFocusPercentRange("Spell and Bash Hate", base, limit, false);
              break;

            case 131:
              printBuffer += this.getFocusPercentRange("Chance of Using Reagent", base, limit, true);
              break;

            case 132:
              printBuffer += this.getFocusPercentRange("Spell Mana Cost", base, limit, true);
              break;

            case 133: //Not on live
              printBuffer += this.getFocusPercentRange("Stun Time", base, limit, false);
              break;

            case 134:
              limit = limit ? limit : 100
              printBuffer += "Limit Max Level: " + base + " (lose " + limit + "% per level)"
              break;

            case 135:
              if (base < 0){
                tmp += "Exclude "
              }
              printBuffer += "Limit Resist: " + tmp + DB_SPELL_RESISTS[Math.abs(base)]
              break;

            case 136:
              if (base < 0){
                tmp += "Exclude "
              }
              printBuffer += "Limit Target: " + tmp + DB_SPELL_TARGETS[Math.abs(base)]
              break;

            case 137:
              if (base < 0){
                tmp += "Exclude "
              }
              printBuffer += "Limit Effect: " + tmp + DB_SPA[Math.abs(base)]
              break;

            case 138:
              tmp += base ? "Beneficial" : "Detrimental"
              printBuffer += "Limit Type: " + tmp
              break;

            case 139://TODO need spell links
              if (base < 0){
                tmp += "Exclude "
              }
              printBuffer += "Limit Spell: " + tmp + (await this.getSpellName(Math.abs(base)))
              break;

            case 140:
              printBuffer += "Limit Min Duration: " + (base * 6) + "s"
              break;

            case 141:
              tmp += base ? "Non-Duration Spells" : "Duration Spells"
              printBuffer += "Limit Duration Type: " + tmp
              break;

            case 142:
              printBuffer += "Limit Min Level: " + base
              break;

            case 143:
              printBuffer += "Limit Min Casting Time: " + (base/1000) + "s"
              break;

            case 144:
              printBuffer += "Limit Max Casting Time: " + (base/1000) + "s"
              break;

            case 145:
              printBuffer += "Teleport to " + spell["teleport_zone"]
              break;

            case 146: //todo data location for port xyz for 45 , Set position to
              break;

            case 147:
              printBuffer += this.getFormatStandard("Current HP", "%", value_min, value_max, minlvl, maxlvl) + " up to " + max
              break;

            case 148:
              tmp += limit ? limit : spell["formula_" + effectIndex] % 100
              printBuffer += "Stacking: Block new spell if slot " + tmp + " is " + DB_SPA[Math.abs(base)] + " and Less Than " + max
              break;

            case 149:
              tmp += limit ? limit : spell["formula_" + effectIndex] % 100
              printBuffer += "Stacking: Overwrite existing spell if slot " + tmp + " is " + DB_SPA[Math.abs(base)] + " and Less Than " + max
              break;

            case 150:
              tmp += max ? " (Increase heal by " + max + " if affected is above lv " + limit + ")" : ""
              printBuffer += (base == 1) ? "Divine Intervention with 300 Heal" + tmp : "Divine Intervention with 8000 Heal" + tmp
              break;

            case 151:
              tmp += base ? " with Buffs" : ""
              printBuffer += "Suspend Pet" + tmp
              break;

            case 152:
              printBuffer += "Summon Temp Pet: " + spell["teleport_zone"] + " x " + base + " for " + max + "s"
              break;

            case 153:
              printBuffer += "Balance Group HP with " + base + "% Penalty (Max HP taken: " + limit + ")"
              break;

            case 154: //TODO need to update emulator code to use percent based (+0.5% per level difference) and confirm duration change mechanic
              if (limit != 0) {
                printBuffer += "Decrease Detrimental Duration by 50% " + (base/10) + "% Chance)" + this.getUpToMaxLvl(max)
              }
              else{
                printBuffer += "Dispel Detrimental " + (base/10) + "% Chance" + this.getUpToMaxLvl(max)
              }
              break;

            case 156:
              printBuffer += "Illusion: Target"
              break;

            case 157:
              printBuffer += this.getFormatStandard("Spell Damage Shield", "", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 158:
              tmp  += max ? " with up to " +max + "% Base Damage" : ""
              if (limit > 0) {
                tmp += " and " + limit + " Improved Resist Mod"
              }
              else if (limit < 0){
                tmp += " and " + limit + " Reduced Resist Mod"
              }
              printBuffer += this.getFormatStandard("Chance to Reflect Spell", "%", value_min, value_max, minlvl, maxlvl) + tmp
              break;

            case 159:
              printBuffer += this.getFormatStandard("Base Stats", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 160:
              printBuffer += "Intoxicate if Tolerance under " + base
              break;

            case 161:
              tmp += limit ? "Max Per Hit: " + limit : ""
              tmp += max ? ", Total: " + max : ""
              printBuffer +=  "Absorb Spell Damage: " + base + "%," + tmp
              break;

            case 162:
              tmp += limit ? "Max Per Hit: " + limit : ""
              tmp += max ? ", Total: " + max : ""
              printBuffer +=  "Absorb Melee Damage: " + base + "%," + tmp
              break;

            case 163:
              tmp += max ? ", Max Per Hit: " + max : ""
              printBuffer +=  "Absorb " + base + " Hits or Spells " + base + "%," + tmp
              break;

            case 164:
              printBuffer += "Appraise Chest " + value_max
              break;

            case 165:
              printBuffer += "Disarm Chest "  + value_max
              break;

            case 166:
              printBuffer += "Unlock Chest " + value_max
              break;

            case 167:
              printBuffer += "Increase Pet Power " + value_max
              break;

            case 168:
              printBuffer += this.getFormatStandard("Melee Mitigation", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 169:
              tmp += limit >= 0 ? " with " + DB_SKILLS[limit] : ""
              printBuffer += this.getFormatStandard("Chance to Critical Hit", "%", value_min, value_max, minlvl, maxlvl) + tmp
              break;

            case 170:
              printBuffer += this.getFormatStandard("Critical Nuke Damage", "%", value_min, value_max, minlvl, maxlvl) + " of Base Damage"
              break;

            case 171:
              printBuffer += this.getFormatStandard("Chance to Crippling Blow", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 172:
              printBuffer += this.getFormatStandard("Chance to Avoid Melee", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 173:
              printBuffer += this.getFormatStandard("Chance to Riposte", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 174:
              printBuffer += this.getFormatStandard("Chance to Dodge", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 175:
              printBuffer += this.getFormatStandard("Chance to Parry", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 176:
              printBuffer += this.getFormatStandard("Chance to Dual Wield", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 177:
              printBuffer += this.getFormatStandard("Chance to Double Attack", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 178:
              printBuffer += "Lifetap from Weapon Damage: " + base + "%"
              break;

            case 179:
              printBuffer += "Instrument Modifier: " + DB_SKILLS[spell["skill"]] + " " + value_max
              break;

            case 180:
              printBuffer += this.getFormatStandard("Chance to Resist Spell", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 181:
              printBuffer += this.getFormatStandard("Chance to Resist Fear Spell", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 182:
              printBuffer += this.getFormatStandard("Weapon Delay", "%", value_min/10, value_max/10, minlvl, maxlvl)
              break;

            case 183:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 184:
              tmp += (limit >= 0) ? " with " + DB_SKILLS[limit] : ""
              printBuffer += this.getFormatStandard("Chance to Hit", "%", value_min, value_max, minlvl, maxlvl) + tmp
              break;

            case 185:
              tmp += (limit >= 0) ? " with " + DB_SKILLS[limit] : " with all skills"
              printBuffer += this.getFormatStandard("Hit Damage", "%", value_min, value_max, minlvl, maxlvl) + tmp
              break;

            case 186:
              tmp += (limit >= 0) ? " with " + DB_SKILLS[limit] : " with all skills"
              printBuffer += this.getFormatStandard("Min Hit Damage", "%", value_min, value_max, minlvl, maxlvl) + tmp
              break;

            case 187:
              printBuffer += "Balance Group Mana with " + base + "% Penalty (Max HP taken: " + limit + ")"
              break;

            case 188:
              printBuffer += this.getFormatStandard("Chance to Block", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 189:
              printBuffer += this.getFormatStandard("Current Endurance", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range
              break;

            case 190:
              printBuffer += this.getFormatStandard("Max Endurance", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 191:
              printBuffer += "Inhibit Combat"
              break;

            case 192:
              printBuffer += this.getFormatStandard("Hate", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range
              break;

            case 193:
              printBuffer += DB_SKILLS[spell["skill"]] + " Attack for " + base + " with " + limit + " % Accuracy Mod"
              break;

            case 194:
              printBuffer += "Cancel Aggro (" + base + " % Chance)"
              break;

            case 195:
              printBuffer += this.getFormatStandard("Chance to Resist Any Stun", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 196:
              printBuffer += this.getFormatStandard("Srikethrough", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 197:
              tmp += (limit >= 0) ? DB_SKILLS[limit] : " Hit "
              printBuffer += this.getFormatStandard(tmp + " Damage Taken", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 198:
              printBuffer += this.getFormatStandard("Current Endurance", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 199:
              printBuffer += "Taunt with " + limit + " added Hate (Chance " + base + "%)"
              break;

            case 200:
              printBuffer += this.getFormatStandard("Worn Proc Rate", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 201:
              printBuffer += "Add Range Proc: "  + (await this.getSpellName(base)) + " with " + limit + "% Rate Mod"
              break;

            case 202:
              printBuffer += "Project Illusion on Next Spell"
              break;

            case 203:
              printBuffer += "Mass Group Buff on Next Spell"
              break;

            case 204:
              printBuffer += "Group Fear Immunity for " + (base * 10) + "s"
              break;

            case 205:
              printBuffer += "Rampage"
              break;

            case 206:
              printBuffer += "AE Taunt with " + base + " added Hate"
              break;

            case 207:
              printBuffer += "Flesh to Bone Chips"
              break;

            case 208:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 209: //TODO need to update emulator code to use percent based (+0.5% per level difference) and confirm duration change mechanic
              if (limit != 0) {
                printBuffer += "Decrease Beneficial Duration by 50% " + (base/10) + "% Chance)" + this.getUpToMaxLvl(max)
              }
              else{
                printBuffer += "Dispel Beneficial " + (base/10) + "% Chance" + this.getUpToMaxLvl(max)
              }
              break;

            case 210:
              printBuffer += "Pet Shielding for " + base * 12 + "s"
              break;

            case 211: //eqemu uses this, Live uses different formula now, base=chance, limit=damage mod
              printBuffer += "AE Melee for " + base * 12 + "s"
              break;

            case 212: //eqemu uses this, Live uses different formula now, base=chance, limit, mana cost mod
              printBuffer += this.getFormatStandard("Frenzied Devastation: Chance to Critical Nuke", "%", value_min, value_max, minlvl, maxlvl) + " and Increase Spell Mana Cost 100%"
              break;

            case 213:
              printBuffer += this.getFormatStandard("Pet Max HP", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 214:
              printBuffer += this.getFormatStandard("Max HP", "%", value_min/100, value_max/100, minlvl, maxlvl)
              break;

            case 215:
              printBuffer += this.getFormatStandard("Pet Chance to Avoid Melee", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 216:
              tmp += (limit >= 0) ? " with " + DB_SKILLS[limit] : ""
              printBuffer += this.getFormatStandard("Accuracy", "", value_min, value_max, minlvl, maxlvl) + tmp
              break;

            case 217:
              printBuffer += "Add Headshot Proc with up to " + limit + " Damage"
              break;

            case 218:
              printBuffer += this.getFormatStandard("Pet Chance to Critical Hit", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 219:
              printBuffer += this.getFormatStandard("Chance to Slay Undead", "%", value_min/100, value_max/100, minlvl, maxlvl) + " with " + limit + " Damage Mod"
              break;

            case 220:
              printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Damage Bonus", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 221:
              printBuffer += this.getFormatStandard("Weight", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 222:
              printBuffer += this.getFormatStandard("Chance to Block from Back", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 223:
              printBuffer += this.getFormatStandard("Chance to Double Riposte", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 224:
              if (limit > 0) {
                printBuffer += this.getFormatStandard("Chance of Additional Riposte", "%", value_min, value_max, minlvl, maxlvl) + " with " + DB_SKILLS[limit]
              }
              else {
                printBuffer += this.getFormatStandard("Chance of Additional Riposte", "%", value_min, value_max, minlvl, maxlvl)
              }
              break;

            case 225:
              printBuffer += this.getFormatStandard("Chance to Double Attack ", "%", value_min, value_max, minlvl, maxlvl) + " (Additive)";
              break;

            case 226:
              printBuffer += "Add Two-Handed Bash Ability";
              break;

            case 227:
              printBuffer += "Decrease " + DB_SKILLS[limit] + " Timer by " + this.humanTime(base) + " (Before Haste)"
              break;

            case 228:
              printBuffer += this.getFormatStandard("Falling Damage", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 229:
              printBuffer += this.getFormatStandard("Chance to Cast Through Stun", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 230:
              printBuffer += this.getFormatStandard("Shield Ability Range", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 231:
              printBuffer += this.getFormatStandard("Chance to Stun Bash", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 232:
              tmp += limit ? (await this.getSpellName(limit)) : (await this.getSpellName(4789))
              printBuffer += "Cast: " + tmp + " on Death (" + base + "% Chance)"
              break;

            case 233:
              printBuffer += this.getFormatStandard("Food Consumption", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 234:
              printBuffer += "Decrease Poison Application Time by " + (10 - base / 1000) + + "s"
              break;

            case 235:
              printBuffer += this.getFormatStandard("Chance to Channel Spells", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 236:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 237:
              printBuffer += "Enable Pet Ability: Receive Group Buffs"
              break;

            case 238:
              printBuffer += (base == 3) ? "Permanent Illusion (Persist After Death)" : " Permanent Illusion"
              break;

            case 239:
              printBuffer += this.getFormatStandard("Chance to Feign Death Through Spell Hit", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 240:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 241:
              printBuffer += "Reclaim Pet Mana (Return " + base + "%)"
              break;

            case 242:
              printBuffer += this.getFormatStandard("Chance to Memory Blur", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 243:
              printBuffer += this.getFormatStandard("Chance of Charm Breaking", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 244:
              printBuffer += this.getFormatStandard("Chance of Root Breaking", "%",  -value_min, -value_max, minlvl, maxlvl)
              break;

            case 245:
              printBuffer += this.getFormatStandard("Chance of Trap Circumvention", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 246:
              printBuffer += this.getFormatStandard("Lung Capacity", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 247:
              printBuffer += this.getFormatStandard( DB_SKILLS[limit] + " Skill Cap", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 248:
              printBuffer += "Train Second Magic Specialization Ability (Secondary Forte)"
              break;

            case 249:
              printBuffer += this.getFormatStandard( "Offhand Weapon Damage Bonus", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 250:
              printBuffer += this.getFormatStandard( "Melee Proc Rate (from buffs, abilities and skills", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 251:
              printBuffer += this.getFormatStandard( "Chance of Using Ammo", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 252:
              printBuffer += this.getFormatStandard( "Chance to Backstab From Front", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 253:
              printBuffer +=  "Allow Frontal Backstab for Minimum Damage"
              break;

            case 255:
              printBuffer += "Increase Shield Ability Duration by " + base + "s"
              break;

            case 256:
              printBuffer += "Shroud of Stealth (" + base + ")"
              break;

            case 257:
              printBuffer += "Enable Pet Ability: Hold"
              break;

            case 258:
              printBuffer += this.getFormatStandard( "Chance to Triple Backstab", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 259:
              printBuffer += this.getFormatStandard( "AC Soft Cap", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 260:
              printBuffer += this.getFormatStandard( DB_BARD_SKILLS[limit] + " Bonus", "%", (value_min * 10), (value_max * 10), minlvl, maxlvl)
              break;

            case 261:
              printBuffer += this.getFormatStandard( "Song Cap", "", (value_min * 10), (value_max * 10), minlvl, maxlvl)
              break;

            case 262:
              printBuffer += this.getFormatStandard( DB_SPELL_WORN_ATTRIBUTE_CAP[limit] + " Cap", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 263:
              printBuffer += this.getFormatStandard( "Ability to Specialize Tradeskills", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 264:
              printBuffer += "Reduce [AA " + limit + "] Timer by " + this.humanTime(base)
              break;

            case 265:
              printBuffer += "No Fizzle up to level " + base
              break;

            case 266:
              printBuffer += this.getFormatStandard( "Chance of " + limit + " Additional 2H Attacks", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 267:
              printBuffer += "Enable Pet Ability: (" + DB_SPELL_PETCMDS[limit] + ")"
              break;

            case 268:
              printBuffer += this.getFormatStandard( "Chance to Fail " + DB_SKILLS[limit] + " Combine", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 269:
              printBuffer += this.getFormatStandard( "Bandage HP Cap", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 270:
              printBuffer += this.getFormatStandard( "Beneficial Song Range", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 271:
              printBuffer += this.getFormatStandard( "Innate Movement Speed", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 272:
              printBuffer += this.getFormatStandard( "Song effective casting level", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 273: //Does live effect now have decay component?
              printBuffer += this.getFormatStandard( "Chance to Critical DoT", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 274: //Does live effect now have decay component?
              printBuffer += this.getFormatStandard( "Chance to Critical Heal", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 275:
              printBuffer += this.getFormatStandard( "Chance to Critical Mend", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 276:
              printBuffer += this.getFormatStandard( "Dual Wield Skill Amount", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 277:
              printBuffer += "Second Chance to Trigger Divine Intervention with a Heal for " + base + "% of baseline"
              break;

            case 278:
              printBuffer += " Add Finishing Blow Proc with up to " + (base / 10) + " Damage (" + limit + "% Chance)"
              break;

            case 279:
              printBuffer += this.getFormatStandard( "Chance to Flurry", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 280:
              printBuffer += this.getFormatStandard( "Pet Chance to Flurry", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 281:
              printBuffer += "Pet Chance to Feign Death (" + base + "%)"
              break;

            case 282:
              printBuffer += this.getFormatStandard( "Bandage Amount", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 283:
              printBuffer += this.getFormatStandard( "Chance to perform a Double Special Attack", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 284:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 285:
              printBuffer += "Chance Hide skill will succeed while moving (" + base + "%)"
              break;

            case 286:
              printBuffer += this.getFormatStandard( "Spell Damage", "", value_min, value_max, minlvl, maxlvl) + " (before crit)"
              break;

            case 287:
              printBuffer += this.getFormatStandard( "Spell Duration", "seconds", (value_min * 6), (value_max * 6), minlvl, maxlvl)
              break;

            case 288: //TODO finish this when AA tables are added, this procs the spell associated with the AA, rank.spell is what the spell id that procs is
              printBuffer += "Add [Insert AA spell] Proc to" + DB_SPA[limit] + "(" + (base/10) + "% Chance)"
              break;

            case 289:
              printBuffer += "Cast:" + (await this.getSpellName(base)) + "on Duration Fade"
              break;

            case 290:
              printBuffer += this.getFormatStandard( "Movement Speed Cap", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 291: //TODO Fix this in source code, not coded correct
              printBuffer += "Remove up to (" + base + ") detrimental effects"
              break;

            case 292:
              printBuffer += this.getFormatStandard( "Chance of Strikethrough", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 293:
              printBuffer += this.getFormatStandard( "Chance to Resist Melee Stun", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 294:
              printBuffer += (base) ? this.getFormatStandard( "Chance to Critical Nuke", "%", value_min, value_max, minlvl, maxlvl) : ""
              printBuffer += (base) ? " and " : ""
              printBuffer += (limit) ?  this.getFormatStandard( "Critical Nuke Damage", "%", limit, limit, minlvl, maxlvl) + " of Base Damage" : ""
              break;

            case 295:
              printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
              break;

            case 296:
              printBuffer += this.getFocusPercentRange("Spell Damage Taken", base, limit, false);
              break;

            case 297:
              printBuffer += this.getFormatStandard( "Spell Damage Taken", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 298:
              printBuffer += this.getFormatStandard( "Pet Size", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 299:
              printBuffer += "Wake the Dead (" + max + ")"
              break;

            case 300:
              printBuffer += "Summon Doppelganger: " + spell["teleport_zone"]
              break;

            case 301:
              printBuffer += this.getFormatStandard( "Archery Damage", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 302: //crit
              printBuffer += this.getFocusPercentRange("Spell Damage", base, limit, false);
              break;

            case 303: //crit
              printBuffer += this.getFormatStandard( "Spell Damage", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 304:
              printBuffer += this.getFormatStandard( "Chance to Avoid Offhand Riposte", "%", -value_min, -value_max, minlvl, maxlvl)
              break;

            case 305:
              if (max) {
                printBuffer += this.getFormatStandard("Offhand Damage Shield Taken", "", value_min, value_max, minlvl, maxlvl)
              }
              else {
                printBuffer += this.getFormatStandard("Offhand Damage Shield Taken", "%", -value_min, -value_max, minlvl, maxlvl)
              }
              break;

            case 306:
              printBuffer += "Wake the Dead: + " + spell["teleport_zone"] + " x " + base + " for " + max + "s"
              break;

            case 307:
              printBuffer += "Appraisal"
              break;

            case 308:
              printBuffer += "Suspend Minion to remain after Zoning";
              break;

            case 309:
              printBuffer += "Teleport to Caster's Bind";
              break;

            case 310:
              printBuffer += "Reduce Timer by " + this.humanTime(base / 1000)
              break;

            case 311:
              printBuffer +=  "Limit Type: " + (base == 1 ? "Include" : "Exclude") + " Combat Skills"
              break;

            case 312:
              printBuffer +=  "Sanctuary: Place caster bottom hate list, fades if cast on other than self."
              break;

            case 313:
              printBuffer += this.getFormatStandard("Chance to Double Forage", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 314:
              printBuffer += "Invisibility" + (base > 1 ? " (Enhanced " + base + ")" : "")
              break;

            case 315:
              printBuffer += "Invisibility to Undead" + (base > 1 ? " (Enhanced " + base + ")" : "")
              break;

            case 316:
              printBuffer += "Invisibility to Animals" + (base > 1 ? " (Enhanced " + base + ")" : "")
              break;

            case 317:
              printBuffer += this.getFormatStandard("HP Regen Cap", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 318:
              printBuffer += this.getFormatStandard("Mana Regen Cap", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 319:
              printBuffer += this.getFormatStandard("Chance to Critical HoT", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 320:
              printBuffer += this.getFormatStandard("Shield Block Chance", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 321:
              printBuffer += this.getFormatStandard("Target's Target Hate", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 322:
              printBuffer +=  "Gate to Home City"
              break;

            case 323:
              if (limit != 0){
                tmp += " with " + limit + " % Rate Mod"
              }
              printBuffer += "Add Defensive Proc: " + (await this.getSpellName(spell["effect_base_value_" + effectIndex])) + tmp
              break;

            case 324:
              printBuffer += "Cast from HP with " + base + "% Penalty"
              break;

            case 325:
              printBuffer += this.getFormatStandard("Chance to Remain Hidden When Hit By AE", "%", value_min, value_max, minlvl, maxlvl)
              break;

            case 326:
              printBuffer += this.getFormatStandard("Spell Memorization Gems", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 327:
              printBuffer += this.getFormatStandard("Buff Slots", "", value_min, value_max, minlvl, maxlvl)
              break;

            case 328:
              printBuffer += this.getFormatStandard("Max Negative HP", "", value_min, value_max, minlvl, maxlvl)
              break;


     
            case 311: // Limit: Combat Skills Not Allowed
            case 314: // Fixed Duration Invisbility
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              break;
            case 58: // Illusion:
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += DB_RACE_NAMES[spell["effect_base_value_" + effectIndex]];
              break;
            case 330: // Critical Damage Mob
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + value_max + "%)";
              break;

            case 323: // Add Defensive Proc:
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);

              printBuffer += "<a href=?a=spell&id=" + spell["effect_base_value_" + effectIndex] + "> " + (await this.getSpellName(spell["effect_base_value_" + effectIndex])) + "</a>";
              break;
          }


          if (printBuffer !== "") {

            effectsInfo.push("Slot " + effectIndex + ": &nbsp " + printBuffer)
/*
            " &nbsp &nbsp &nbsp [ * DEBUG * " + "(ID: " + spell["effectid_" + effectIndex] + ") "
            +   "(Base:" +base + ") "+ "(Limit: " + limit + ") "+ "(max: " + max + ")]"
            +   " Min: " + value_min + " Max: " + value_max + " MinLv: " + minlvl + " MaxLv: " + maxlvl
            + "  [TEST]  " + 0 + " :End"
            )
*/

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
