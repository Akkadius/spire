<template>
  <div
    style="min-width: 400px; max-width: 500px; padding: 5px"
    v-if="spellData && spellData['targettype']"
  >

    <div class="row">
      <div class="col-1">

        <span
          :style="'width: 40px; height: 40px; border: 1px solid ' + getTargetTypeColor(this.spellData['targettype']) + '; border-radius: 7px; '"
          :class="'spell-' + spellData.new_icon + '-40'"
        />

      </div>
      <div class="col-11 pl-5">
        <h6 class="eq-header" style="margin: 0px; margin-bottom: 10px">
          {{ spellData.name }}
        </h6>

      </div>
    </div>

    <!-- Info -->

    <table class="mt-3 spell-preview-table">
      <tbody>
      <tr v-if="spellData['id'] !== ''">
        <td class="spell-field-label">Spell ID</td>
        <td> {{ spellData["id"] }}
          <span v-if="spellData['spellgroup'] !== 0">(Group: {{ spellData["spellgroup"] }})</span>
          <!-- <span v-if="spellData['rank'] !== 0">, Rank: {{spellData['rank']}})</span> -->
        </td>
      </tr>
      <tr v-if="getClasses(spellData) !== ''">
        <td class="spell-field-label">Classes</td>
        <td style="width: 250px">
          <div v-for="(icon, index) in dbClassIcons" style="display: inline-block">
            <div
              v-if="spellData['classes_' + index] > 0 && spellData['classes_' + index] < 255"
              class="mr-2"
            >
               <span
                 style="border-radius: 4px"
                 :class="'item-' + icon + '-sm'"
                 :title="dbClassesShort[index]"
               />
              {{ dbClassesShort[index] }}
              ({{ spellData["classes_" + index] }})
            </div>
          </div>
        </td>

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

      <tr
        v-if="spellData['typedescnum'] !== '' && getSpellTypeDescNumName(spellData['typedescnum']) !== '' && getSpellTypeDescNumName(spellData['effectdescnum']) !== ''"
      >
        <td class="spell-field-label">Book Category</td>
        <td> {{ getSpellTypeDescNumName(spellData["typedescnum"]) }}
          <span v-if="spellData['effectdescnum'] !== ''"> / {{
              getSpellTypeDescNumName(spellData["effectdescnum"])
            }} </span>
        </td>
      </tr>

      <tr v-if="spellData['skill'] < 116 && getDatabaseSkillName(spellData['skill']) !== ''">
        <td class="spell-field-label">Skill</td>
        <td> {{ getDatabaseSkillName(spellData["skill"]) }}
          <span v-if="spellData['is_discipline'] !== 0">(Combat Skill)</span>
        </td>
      </tr>
      <tr v-if="spellData['good_effect'] >= 0">
        <td class="spell-field-label">Type</td>
        <td>
          <span v-if="spellData['good_effect'] === 0">Detrimental</span>
          <span v-if="spellData['good_effect'] === 1">Beneficial</span>
          <span v-if="spellData['good_effect'] === 2">Beneficial Group</span>
          <span v-if="spellData['good_effect'] === 3">Unknown</span>
        </td>
      </tr>

      <!-- Resources -->

      <tr v-if="spellData['mana'] > 0">
        <td class="spell-field-label">Mana</td>
        <td> {{ spellData["mana"] }}</td>
      </tr>

      <tr v-if="spellData['endur_cost'] !== 0 ">
        <td class="spell-field-label">Endurance</td>
        <td> {{ spellData["endur_cost"] }}</td>
      </tr>
      <tr v-if="spellData['endur_upkeep'] !== 0 ">
        <td class="spell-field-label">Endurance Upkeep</td>
        <td> {{ spellData["endur_upkeep"] }} per second</td>
      </tr>

      <!-- Restrictions -->

      <tr v-if="spellData['cast_restriction'] > 2">
        <td class="spell-field-label">Target Restriction</td>
        <td> {{ getSpellTargetRestrictionTypeName(spellData["cast_restriction"]) }}</td>
      </tr>

      <tr v-if="spellData['field_220'] > 2">
        <td class="spell-field-label">Caster Restriction</td>
        <td> {{ getSpellTargetRestrictionTypeName(spellData["field_220"]) }}</td>
      </tr>

      <tr v-if="spellData['in_combat'] === 0 && spellData['outof_combat'] !== 0">
        <td class="spell-field-label">Restriction</td>
        <td> Out of Combat Only</td>
      </tr>

      <tr v-if="spellData['in_combat'] !== 0 && spellData['outof_combat'] == 0">
        <td class="spell-field-label">Restriction</td>
        <td> In Combat Only</td>
      </tr>

      <tr v-if="spellData['field_234'] !== 0">
        <td class="spell-field-label">Restriction</td>
        <td> Only During Fast Regen</td>
      </tr>

      <tr v-if="spellData['disallow_sit'] !== 0">
        <td class="spell-field-label">Restriction</td>
        <td> Cancel on Sit</td>
      </tr>

      <tr v-if="spellData['sneaking'] !== 0">
        <td class="spell-field-label">Restriction</td>
        <td> Must Be Sneaking</td>
      </tr>

      <tr v-if="spellData['zonetype'] > 0">
        <td class="spell-field-label">Restriction</td>
        <td>
          <span v-if="spellData['zonetype'] === 1"> Outdoor Only </span>
          <span v-if="spellData['zonetype'] === 2"> Indoor Only </span>
        </td>
      </tr>

      <!-- Casting -->

      <tr
        v-if="(spellData['cast_time'] > 0 || spellData['recovery_time'] > 0 || spellData['recast_time'] > 0) && spellData['is_discipline'] === 0"
      >
        <td class="spell-field-label">Casting Time</td>
        <td> {{ (spellData["cast_time"] / 1000) }} sec
          <span v-if="spellData['uninterruptable'] !== 0">(Uninterruptable)</span>
        </td>
      </tr>
      <tr
        v-if="(spellData['cast_time'] > 0 || spellData['recovery_time'] > 0 || spellData['recast_time'] > 0) && spellData['is_discipline'] === 0"
      >
        <td class="spell-field-label">Recovery Time</td>
        <td> {{ (spellData["recovery_time"] / 1000) }} sec</td>
      </tr>
      <tr v-if="spellData['cast_time'] > 0 || spellData['recovery_time'] > 0 || spellData['recast_time'] > 0">
        <td class="spell-field-label">Recast Time</td>
        <td> {{ (spellData["recast_time"] / 1000) }} sec</td>
      </tr>
      <tr v-if="spellData['endur_timer_index'] > 0 ">
        <td class="spell-field-label">Timer</td>
        <td> {{ spellData["endur_timer_index"] }}</td>
      </tr>

      <!-- Duration / Buffs -->

      <tr v-if="getBuffDuration(spellData)">
        <td class="spell-field-label">Duration</td>
        <td> {{ humanTime(getBuffDuration(spellData) * 6) }} - {{ getBuffDuration(spellData) }} tic(s)</td>
      </tr>

      <tr v-if="getBuffDuration(spellData)">
        <td class="spell-field-label">Dispelable</td>
        <td>
          <span v-if="spellData['dispel_flag'] !== 0">No</span>
          <span v-if="spellData['dispel_flag'] === 0">Yes</span>
          <span v-if="spellData['field_232'] !== 0">, Can Not Remove</span>
        </td>
      </tr>

      <tr v-if="spellData['persistdeath'] !== 0">
        <td class="spell-field-label">Persist After Death</td>
        <td> Yes</td>
      </tr>

      <!-- ToDO
      <span v-if="spellData['dispel_flag'] !== 0">,  Dispelable: Yes </span>
      <span v-if="spellData['short_buff_box'] !== 0">, Song </span>
      <span v-if="spellData['no_remove'] !== 0">, Can Not Remove </span>
      -->

      <tr v-if="spellData['ae_duration'] > 0">
        <td class="spell-field-label">AE Waves</td>
        <td> {{ spellData["ae_duration"] / 2500 }} waves</td>
      </tr>


      <tr v-if="spellData['range'] > 0">
        <td class="spell-field-label">Range</td>
        <td>
          <span v-if="spellData['min_range'] > 0 && spellData['aoerange'] === 0 ">  {{
              spellData["min_range"]
            }}' to </span>
          {{ spellData["range"] }}'
        </td>
      </tr>
      <tr v-if="spellData['aoerange'] > 0">
        <td class="spell-field-label">AOE Range</td>
        <td>
          <span v-if="spellData['min_range'] > 0">  {{ spellData["min_range"] }}' to </span>
          {{ spellData["aoerange"] }}'
        </td>
      </tr>
      <tr
        v-if="(spellData['max_dist'] !== 0 || spellData['min_dist'] !== 0) && (spellData['max_dist_mod'] !== 0 || spellData['min_dist_mod'] !== 0) "
      >
        <td class="spell-field-label">Range Based Mod</td>
        <td> ({{ spellData["min_dist_mod"] * 100 }}% at {{ spellData["min_dist"] }}') to
          ({{ spellData["max_dist_mod"] * 100 }}% at {{ spellData["max_dist"] }}')
        </td>
      </tr>

      <tr v-if="spellData['viral_range'] > 0">
        <td class="spell-field-label">Viral Range</td>
        <td> {{ spellData["viral_range"] }}, Recast: {{ spellData["viral_targets"] }}s to
          {{ spellData["viral_timer"] }}s
        </td>
      </tr>

      <tr v-if="spellData['targettype'] > 0 && getTargetTypeName(spellData['targettype']) !== ''">
        <td class="spell-field-label">Target</td>
        <td> {{ getTargetTypeName(spellData["targettype"]) }}
          <span
            v-if="spellData['can_mgb'] == 0 && (spellData['buffduration'] > 0) && (spellData['targettype'] === 3  || spellData['targettype'] === 40 || spellData['targettype'] === 41)"
          > &nbsp; (No MGB)</span>
          <span
            v-if="spellData['can_mgb'] === 1 && (spellData['buffduration'] === 0) && (spellData['targettype'] === 3  || spellData['targettype'] === 40 || spellData['targettype'] === 41)"
          > &nbsp; (Can MGB)</span>
        </td>
      </tr>
      <tr v-if="spellData['aemaxtargets'] > 0 ">
        <td class="spell-field-label">Max Targets</td>
        <td> {{ spellData["aemaxtargets"] }}</td>
      </tr>

      <tr v-if="spellData['cone_start_angle'] !== 0 || spellData['cone_stop_angle'] !== 0">
        <td class="spell-field-label">Cone Angle</td>
        <td> {{ getConeAngleDescription(spellData["cone_start_angle"], spellData["cone_stop_angle"]) }}</td>
      </tr>

      <tr v-if="spellData['resisttype'] > 0 && getSpellResistTypeName(spellData['resisttype']) !== ''">
        <td class="spell-field-label">Resist Type</td>
        <td> {{ getSpellResistTypeName(spellData["resisttype"]) }}
          <span
            v-if="spellData['resist_diff'] !== 0 && spellData['field_209'] === 0  && spellData['no_partial_resist'] === 0"
          > &nbsp;({{
              spellData["resist_diff"]
            }})</span>
          <span
            v-if="spellData['resist_diff'] !== 0 && spellData['field_209'] === 0  && spellData['no_partial_resist'] !== 0"
          > &nbsp;({{
              spellData["resist_diff"]
            }}) &nbsp; (No Partial Resist)</span>
          <span v-if="spellData['field_209'] !== 0">(Unresistable)</span>
        </td>
      </tr>
      <tr v-if="spellData['max_resist'] > 0 || spellData['min_resist'] > 0">
        <td class="spell-field-label">Resist Chance Limits</td>
        <td>
          <span v-if="spellData['max_resist'] !== 0">Max: {{ spellData["max_resist"] / 2 }}% </span>
          <span v-if="spellData['min_resist'] !== 0">Min: {{ spellData["min_resist"] / 2 }}% </span>
        </td>
      </tr>

      <tr v-if="spellData['not_extendable'] !== 0 ">
        <td class="spell-field-label">Focusable</td>
        <td>No</td>
      </tr>

      <!-- Added effects -->

      <tr v-if="spellData['pushback'] !== 0 || spellData['pushup'] !== 0">
        <td class="spell-field-label">Knockback</td>
        <td>
          <span v-if="spellData['pushback'] !== 0">Push Back: {{ spellData["pushback"] }}' </span>
          <span v-if="spellData['pushup'] !== 0">Push Up: {{ spellData["pushup"] }}' </span>
        </td>
      </tr>

      <tr v-if="spellData['bonushate'] !== 0">
        <td class="spell-field-label">Hate Mod</td>
        <td>
          <span v-if="spellData['bonushate'] > 0">+{{ spellData["bonushate"] }} </span>
          <span v-if="spellData['bonushate'] < 0"> {{ spellData["bonushate"] }} </span>
        </td>
      </tr>

      <tr v-if="spellData['hate_added'] !== 0">
        <td class="spell-field-label">Hate</td>
        <td>
          <span v-if="spellData['hate_added'] > 0">+{{ spellData["hate_added"] }} </span>
          <span v-if="spellData['hate_added'] < 0"> {{ spellData["hate_added"] }} </span>
        </td>
      </tr>

      <tr v-if="spellData['field_198'] !== 0 ">
        <td class="spell-field-label">Spell Hate</td>
        <td>No Detrimental Spell Aggro</td>
      </tr>

      <tr v-if="spellData['field_217'] > 0 ">
        <td class="spell-field-label">Max Critical Chance</td>
        <td> {{ spellData["field_217"] }}%</td>
      </tr>

      <!-- other -->


      <tr v-if="spellData['nimbuseffect'] > 0 ">
        <td class="spell-field-label">Nimbus Type</td>
        <td> {{ spellData["nimbuseffect"] }}</td>
      </tr>

      <tr v-if="spellData['numhitstype'] > 0 ">
        <td class="spell-field-label">Max Hits</td>
        <td> {{ spellData["numhits"] }} {{ getSpellNumHitsTypeName(spellData["numhitstype"]) }}</td>
      </tr>

      <tr v-if="spellData['recourse_link'] > 0 && recourseLink !== ''">
        <td class="spell-field-label">Recourse</td>
        <v-runtime-template :template="'<td>' + recourseLink + '</td>'"/>
      </tr>


      <!-- TODO: Optionally Display Items that have this spell as a scroll / tome -->

      <!--      <eq-item-card-preview/>-->

      </tbody>
    </table>

    <!-- Spell Effects -->
    <h6 class="eq-header mt-3" v-if="spellEffectInfo.length > 0">Effects</h6>
    <div v-if="spellEffectInfo.length > 0">
      <div v-for="effect in spellEffectInfo">
        <v-runtime-template
          :template="'<span>' + effect + '</span>'"
          v-if="typeof effect !== 'undefined'"
          class="pb-6 mt-3 doc"
        />
      </div>
    </div>

    <!-- Reagents -->
    <h6 class="eq-header mt-3" v-if="reagents.length > 0">Reagents</h6>
    <div v-if="reagents.length > 0">
      <div v-for="reagent in reagents" style="width:100%">
        <div :id="reagent.id + '-' + reagent.item.id + '-' + componentId" style="display:inline-block">

          <div style="display: inline-block">

            <div :class="'ml-1 mr-1 item-' + reagent.item.icon + '-sm'"/>

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
            <eq-item-card-preview :item-data="reagent.item"/>
          </eq-window>
        </b-popover>
      </div>
    </div>

    <h6 class="eq-header mt-3" v-if="spellData['descnum'] > 0 && effectDescription !== ''">Description</h6>
    <div class="mt-3 mb-3" v-if="spellData['descnum'] > 0 && effectDescription !== ''" style="width: 70%">
      {{ effectDescription }}
    </div>

    <eq-debug :data="spellData"/>
  </div>
</template>

<script>

import {App}              from "@/constants/app";
import {DB_SKILLS}        from "@/app/constants/eq-skill-constants";
import {
  DB_SPELL_EFFECTS,
  DB_SPELL_NUMHITSTYPE,
  DB_SPELL_RESISTS,
  DB_SPELL_TARGET_RESTRICTION,
  DB_SPELL_TARGETS,
  DB_SPELL_TYPEDESCNUM
}                         from "@/app/constants/eq-spell-constants";
import EqWindow           from "@/components/eq-ui/EQWindow";
import EqDebug            from "@/components/eq-ui/EQDebug";
import {Spells}           from "@/app/spells";
import {Items}            from "@/app/items";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT} from "@/app/constants/eq-classes-constants";

let unknowns = {}

export default {
  name: "EqSpellPreview",
  components: {
    EqWindow,
    EqDebug,
    "eq-item-card-preview": () => import("@/components/preview/EQItemCardPreview.vue"),
    "v-runtime-template": () => import("v-runtime-template")
  },
  watch: {
    spellData: {
      handler: function (val, oldVal) {
        this.init()
      },
      deep: true
    },
  },
  data() {
    return {
      debug: App.DEBUG,
      debugSpellEffects: false,
      spellEffectInfo: [],
      sideLoadedSpellData: {},
      componentId: "",
      reagents: [],
      itemData: {},
      effectDescription: "",
      recourseLink: "",
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT
    }
  },
  created() {
    this.itemData = Items.items
  },

  mounted() {
    this.init()
  },
  methods: {
    async init() {

      if (!this.spellData || !this.spellData["targettype"]) {
        return
      }

      // async each effect index if it exists
      // this is so loading spell effects and any subsequent ajax requests
      // do not block the card from loading
      for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
        if (this.spellEffectInfo[effectIndex]) {
          this.spellEffectInfo[effectIndex] = "";
          this.$forceUpdate()
        }

        if (this.spellData["effectid_" + effectIndex] !== 254) {
          this.getSpellEffectInfo(this.spellData, effectIndex).then((result) => {
            this.spellEffectInfo[result.index] = result.info;
            this.$forceUpdate()
          })
        }
      }

      const uuidv4     = require("uuid/v4")
      this.componentId = uuidv4()

      // reagents
      let reagents = []
      for (let i = 0; i <= 4; i++) {
        if (this.spellData["components_" + i] > 0) {
          let reagent  = {}
          reagent.id   = this.spellData["components_" + i]
          reagent.item = await Items.getItem(this.spellData["components_" + i])
          reagents.push(reagent)
        }
      }

      // recourse
      if (this.spellData["recourse_link"] > 0) {
        // async
        Spells.renderSpellMini(this.spellData.id, this.spellData["recourse_link"]).then((result) => {
          this.recourseLink = result;
        })
      }

      this.reagents = reagents
      this.loadSpellDescription();

      this.sideLoadedSpellData = Spells.data;
    },
    getTargetTypeColor(targetType) {
      return Spells.getTargetTypeColor(targetType)
    },

    loadSpellDescription() {
      this.effectDescription = ""
      Spells.getSpellDescription(this.spellData).then((result) => {
        if (result && result.trim() !== "") {
          this.effectDescription = result;
        }
      })
    },

    getClasses(spell) {
      return Spells.getClasses(spell)
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
    getSpellTargetRestrictionTypeName: function (id) {
      return DB_SPELL_TARGET_RESTRICTION[id] ? DB_SPELL_TARGET_RESTRICTION[id] : ""
    },
    getSpellNumHitsTypeName: function (id) {
      return DB_SPELL_NUMHITSTYPE[id] ? DB_SPELL_NUMHITSTYPE[id] : ""
    },
    getSpellTypeDescNumName: function (id) {
      return DB_SPELL_TYPEDESCNUM[id] ? DB_SPELL_TYPEDESCNUM[id] : ""
    },

    getConeAngleDescription: function (start, stop) {

      if (start >= 1 && start <= 180) {
        return "Right 180 degree Arc"
      }

      if ((start >= 270 && stop <= 90) && ((360 - start) === stop)) {
        return "Frontal " + ((360 - start) + stop) + " degree Arc"
      }
      if ((start >= 90 && start <= 180) && (stop >= 180 && stop <= 270) && ((360 - start) === stop)) {
        return "Rear " + Math.abs((start) - stop) + " degree Arc"
      }
      if ((start >= 180 && start <= 270) && (stop >= 270 && stop <= 360) && (Math.abs(270 - start) === Math.abs(270 - stop))) {
        return "Left " + Math.abs(start - stop) + " degree Arc"
      }
      if ((start >= 0 && start <= 90) && (stop >= 90 && stop <= 180) && (Math.abs(90 - start) === Math.abs(90 - stop))) {
        return "Right " + Math.abs(start - stop) + " degree Arc"
      }

      return start + " degrees to " + stop + " degrees"
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
    getItem: async function (itemId) {
      return await Items.getItem(itemId)
    },
    getItemName: async function (itemId) {
      const item = await this.getItem(itemId)

      return item.name ? item.name : "Unknown Item Name"
    },
    getBuffDuration: function (spell) {
      return Spells.getBuffDuration(spell)
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

    getSpellEffectInfo: async function (spell, effectIndex) {
      return await Spells.getSpellEffectInfo(spell, effectIndex)
    }
  },
  props: {
    spellData: Object
  }
}
</script>

<style>
.spell-field-label {
  text-align: right;
  font-weight: bold;
  padding-right: 10px;
  width: 35%;
}

.spell-preview-table {
  word-wrap: break-word;
  width: 100%;
}

.spell-preview-table th, td {
  word-wrap: break-word;
}

@media only screen and (max-width: 600px) {
  .spell-preview-table {
    width: 90%;
  }
}

</style>
