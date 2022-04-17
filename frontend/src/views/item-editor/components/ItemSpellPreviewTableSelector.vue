<template>
  <div>
    <div class='eq-window-nested-blue text-center' v-if="spells.length === 0">
      No spells were found
    </div>

    <div class='spell-table' v-if="spells.length > 0">
      <table class="eq-table eq-highlight-rows" style="display: table;">
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
        >
          <td style="vertical-align: middle">
            <b-button-group variant="outline-warning">
              <b-dropdown
                block right variant="outline-warning" size="sm"
                title="Select"
              >
                <b-dropdown-item @click="selectAsEffect('scrolleffect', spell)">As Scroll Effect</b-dropdown-item>
                <b-dropdown-item @click="selectAsEffect('clickeffect', spell)">As Click Effect</b-dropdown-item>
                <b-dropdown-item @click="selectAsEffect('proceffect', spell)">As Proc Effect</b-dropdown-item>
                <b-dropdown-item @click="selectAsEffect('focuseffect', spell)">As Focus Effect</b-dropdown-item>
                <b-dropdown-item @click="selectAsEffect('worneffect', spell)">As Worn Effect</b-dropdown-item>
                <b-dropdown-item @click="selectAsEffect('bardeffect', spell)">As Bard Effect</b-dropdown-item>
              </b-dropdown>
            </b-button-group>
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
            <spell-popover
              :spell="spell"
              :size="30"
              :spell-name-length="25"
              v-if="spell && Object.keys(spell).length > 0"
              class="mt-2"
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
import EqSpellEffects     from "@/components/eq-ui/EQSpellEffects";
import EqSpellPreview     from "@/components/eq-ui/EQSpellCardPreview.vue";
import {App}              from "@/constants/app";
import EqSpellDescription from "@/components/eq-ui/EQSpellDescription";
import {DB_SPELL_TARGETS} from "@/app/constants/eq-spell-constants";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT} from "@/app/constants/eq-classes-constants";
import {ROUTE}            from "@/routes";
import * as util          from "util";
import SpellPopover       from "@/components/SpellPopover";

export default {
  name: "ItemSpellPreviewTableSelector",
  components: {
    SpellPopover,
    EqSpellDescription,
    EqSpellEffects,
    EqSpellPreview,
    EqWindow,
  },
  data() {
    return {
      title: "",
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,
    }
  },
  async created() {
    this.title = "Spells (" + this.spells.length + ")";
  },
  props: {
    spells: Array
  },
  methods: {
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
.eq-table tr {
  border-bottom: .4px solid #ffffff1c;
}

.eq-table td {
  padding-top: 5px;
  padding-bottom: 5px;
  border-right: .1px solid #ffffff1c;
  border-left: .1px solid #ffffff1c;
}

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
