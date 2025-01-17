<template>
  <div>
    <div class='eq-window-nested-blue text-center' v-if="spells.length === 0">
      No spells were found
    </div>

    <div
      style="overflow-y: scroll; overflow-x: hidden; height: 58vh"
      id="spell-effect-selector-view-port"
      class='spell-table p-0'
         v-if="spells.length > 0">
      <table
        id="spell-preview-table"
        class="eq-table bordered eq-highlight-rows spell-preview-table"
        style="display: table;">
        <thead>
        <tr>
          <th></th>
          <th style="width: auto; min-width: 100px; text-align: center">Level</th>
          <th style="width: 400px; text-align: center">Effects</th>
          <th style="width: auto; min-width: 200px; text-align: center">Spell</th>
        </tr>
        </thead>
        <tbody>
        <tr
          v-for="(spell, index) in spells"
          :key="spell.id"
          :class="(isSpellSelected(spell) ? 'pulsate-highlight-white' : '')"
        >
          <td style="vertical-align: middle">

            <b-button
              class="btn-dark btn-sm btn-dark mb-3"
              @click="selectSpell(spell.id)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>

          </td>
          <td
            style="text-align: center; padding: 5px; vertical-align: middle"
          >
                <span v-for="(icon, index) in dbClassIcons">
                  <div v-if="spell['classes_' + index] > 0 && spell['classes_' + index] < 255">

                     <span
                       style="border-radius: 4px"
                       :class="'item-' + icon + '-sm'"
                       :title="dbClassesShort[index]"
                     />
                    {{ dbClassesShort[index] }}
                    ({{ spell["classes_" + index] }})
                    </div>
                </span>
          </td>
          <td style="vertical-align: middle">
            <eq-spell-effects :spell="spell"/>
          </td>
          <td style="vertical-align: middle">
            <v-runtime-template
              v-if="spellMinis"
              :template="'<span>' + spellMinis[spell.id] + '</span>'"
            />
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import {Spells}           from "@/app/spells";
import EqWindow           from "@/components/eq-ui/EQWindow.vue";
import EqSpellEffects     from "@/components/preview/EQSpellEffects";
import EqSpellPreview     from "@/components/preview/EQSpellCardPreview.vue";
import {App}              from "@/constants/app";
import EqSpellDescription from "@/components/preview/EQSpellDescription";
import {DB_SPELL_TARGETS} from "@/app/constants/eq-spell-constants";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT} from "@/app/constants/eq-classes-constants";
import {ROUTE}            from "@/routes";
import * as util          from "util";

export default {
  name: "SpellPreviewTableSelector",
  components: {
    EqSpellDescription,
    EqSpellEffects,
    EqSpellPreview,
    EqWindow,
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {
      debug: App.DEBUG,
      debugSpellEffects: false,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      spellEffectInfo: [],
      itemData: {},
      sideLoadedSpellData: {},
      componentId: "",
      reagents: [],
      effectDescription: "",
      recourseLink: "",
      title: "",
      spellMinis: {},
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,

      selectedSpellId: 0,
    }
  },
  async created() {
    let spellMinis = []
    for (const spell of this.spells) {
      Spells.setSpell(spell["id"], spell)

      spellMinis[spell["id"]] = await Spells.renderSpellMini("0", spell["id"], 30)
    }
    this.spellMinis = spellMinis

    // this.$forceUpdate()

    // do this once so we're not triggering vue re-renders in the loop
    this.sideLoadedSpellData = Spells.data

    this.title = "Spells (" + this.spells.length + ")";
  },
  props: {
    spells: Array
  },
  methods: {

    isSpellSelected(spell) {
      return spell.id === this.selectedSpellId
    },

    selectSpell(spellId) {
      const event = {
        spellId: spellId,
      }

      this.$emit('input', event);

      this.selectedSpellId = spellId
    },

    selectAsEffect(field, spell) {
      const event = {
        field: field,
        spell: spell,
      }

      this.$emit('input', event);
    },
    getClasses: function (spell) {
      return Spells.getClasses(spell)
    },
    getTargetTypeColor: function (targetType) {
      return Spells.getTargetTypeColor(targetType)
    },
    getTargetTypeName: function (targetType) {
      return DB_SPELL_TARGETS[targetType] ? DB_SPELL_TARGETS[targetType] : "Unknown Target (" + targetType + ")"
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
    getBuffDuration: function (spell) {
      return Spells.getBuffDuration(spell)
    },
    editSpell(spellId) {
      this.$router.push(
        {
          path: util.format(ROUTE.SPELL_EDIT, spellId),
          query: {}
        }
      ).catch(() => {
      })
    }
  }
}
</script>

<style scoped>

/* For Mobile */
@media screen and (max-width: 540px) {
  .spell-table {
    overflow-x: visible;
    overflow-y: scroll !important
  }
}

/* For Tablets */
@media screen and (min-width: 540px) and (max-width: 780px) {
  .spell-table {
    overflow-x: visible;
    overflow-y: scroll !important
  }
}
</style>
