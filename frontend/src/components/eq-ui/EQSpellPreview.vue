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
import {DB_SKILLS} from "@/app/constants/eq-skill-constants";
import {DB_SPELL_EFFECTS, DB_SPA, DB_SPELL_RESISTS, DB_SPELL_TARGETS, DB_SPELL_TARGET_RESTRICTION} from "@/app/constants/eq-spell-constants";
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
       return printBuffer += "Focus: " + modifier + effect_name + " by " + Math.abs(min) + "%";
     }
     return printBuffer += "Focus: " + modifier + effect_name + " by " + Math.abs(min) + "% to " + Math.abs(max) + "%";
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
              printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl)  + pertick + special_range + tmp
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
              printBuffer += "Summon Pet: " +spell["teleport_zone"]
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
              printBuffer += this.getFormatStandard("Damage Shield", "", value_min, value_max, minlvl, maxlvl);
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
              printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl)  + special_range + tmp
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

            case 100: //TODO add heal over time
              // heal over time
             // Spell.FormatCount("Current HP", value) + repeating + range + (base2 > 0 ? " (If " + Spell.FormatEnum((SpellTargetRestrict)base2) + ")" : "");
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
              if (value_min !== value_max) {
                printBuffer += " by " + value_min + "% (L" + minlvl + ") to " + value_max + "% (L" + maxlvl + ")";
              } else {
                printBuffer += " by " + value_max + "%";
              }
              break;
            case 15: // Increase Mana per tick
            case 100: // Increase Hitpoints v2 per tick
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              let duration = spell["buffduration"]
              if (value_min !== value_max) {
                printBuffer += " by " + Math.abs(value_min) + " (L" + minlvl + ") to " + Math.abs(
                  value_max
                ) + " (L" + maxlvl + ") per tick (total " + Math.abs(value_min * duration) + " to " + Math.abs(
                  value_max * duration
                ) + ")";
              } else {
                printBuffer += " by " + value_max + " per tick (total " + Math.abs(value_max * duration) + ")";
              }
              break;

            case 152: // Summon Pets:
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " <a href=?a=pet&name=" + spell["teleport_zone"] + ">" + spell["teleport_zone"] + "</a>";
              break;

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
              printBuffer += " (" + value_max + "%)";
              break;
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
              case 134: // Limit: Max Level
            case 157: // Spell-Damage Shield
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + value_max + ")";
              break;
            case 121: // Reverse Damage Shield
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (-" + value_max + ")";
              break;
            case 91: // Summon Corpse
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (value_max level " + value_max + ")";
              break;
            case 136: // Limit: Target
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (value_max < 0) {
                value_max = -value_max;
                v   = " excluded";
              } else {
                v = "";
              }
              printBuffer += " (" + DB_SPELL_TARGETS[value_max] + " " + v + ")";
              break;
            case 139: // Limit: Spell
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              value_max = spell["effect_base_value_" + effectIndex];
              if (value_max < 0) {
                value_max = -value_max;
                v   = " excluded";
              }

              name = (await this.getSpellName(value_max))
              if (csv === false) {
                printBuffer += "(" + name + ")";
              } else {
                printBuffer += " (<a href=?a=spell&id=" + spell["effect_base_value_" + effectIndex] + ">" + name + "</a>v)";
              }
              break;
            case 140: // Limit: Min Duration
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              value_min *= 6;
              value_max *= 6;
              if (value_min !== value_max) {
                printBuffer += " (" + value_min + " sec (L" + minlvl + ") to " + value_max + " sec (L" + maxlvl + "))";
              } else {
                printBuffer += " (value_max sec)";
              }
              break;
            case 143: // Limit: Min Casting Time
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              value_min *= 6;
              value_max *= 6;
              if (value_min !== value_max) {
                printBuffer += " (" + (value_min / 6000) + " sec (L" + minlvl + ") to " + (value_max / 6000) + " sec (L" + maxlvl + "))";
              } else {
                printBuffer += " (" + (value_max / 6000) + " sec)";
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
              if (value_max < 0) {
                name = name.replace("Increase", "Decrease");
              }
              printBuffer += name + " by " + spell["effect_limit_value_" + effectIndex] + " (" + value_max + "% " + value_max + ")";
              break;
            case 153: // Balance Party Health
              printBuffer += this.getSpellEffectName(spell["effectid_" + effectIndex]);
              printBuffer += " (" + value_max + "% penalty)";
              break;
            case 0: // In/Decrease hitpoints

            case 159: // Decrease Stats
            case 167: // Pet Power Increase
            case 192: // Increase hate
            default:
              name = this.getSpellEffectName(spell["effectid_" + effectIndex]);
              if (value_max < 0) {
                name = name.replace("Increase", "Decrease");
              }
              printBuffer += name;
              if (value_min !== value_max) {
                printBuffer += " by " + value_min + " (L" + minlvl + ") to " + value_max + " (L" + maxlvl + ")";
              } else {
                if (value_max < 0) {
                  value_max = -value_max;
                }
                printBuffer += " by " + value_max + "";
              }

              break;
          }


          if (printBuffer !== "") {

            //let test = this.getFormatStandard(" Movement Speed", "%", value_min, value_max, minlvl, maxlvl)
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
