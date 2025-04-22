<template>
  <div>
    <div class='eq-window-nested-blue text-center' v-if="spells.length === 0">
      No spells were found
    </div>

    <div
      class='spell-preview-table fill-screen'
      v-if="spells.length > 0"
    >
      <table
        id="spell-preview-table"
        class="eq-table bordered eq-highlight-rows p-0" style="display: table;">
        <thead class="eq-table-floating-header">
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
            <div class="dropdown d-inline-block">
              <a
                href="#"
                data-toggle="dropdown"
                aria-haspopup="true"
                aria-expanded="false"
                class="btn btn-dark btn-sm dropdown-toggle"
              >
                <i class="ra ra-sapphire"></i> Select
              </a>

              <eq-window class="dropdown-menu dropdown-menu-left p-0">
                <a href="#" @click="selectAsEffect('scrolleffect', spell)" class="dropdown-item pl-3">
                  As Scroll Effect
                </a>
                <a href="#" @click="selectAsEffect('clickeffect', spell)" class="dropdown-item pl-3">
                  As Click Effect
                </a>
                <a href="#" @click="selectAsEffect('proceffect', spell)" class="dropdown-item pl-3">
                  As Proc Effect
                </a>
                <a href="#" @click="selectAsEffect('focuseffect', spell)" class="dropdown-item pl-3">
                  As Focus Effect
                </a>
                <a href="#" @click="selectAsEffect('worneffect', spell)" class="dropdown-item pl-3">
                  As Worn Effect
                </a>
                <a href="#" @click="selectAsEffect('bardeffect', spell)" class="dropdown-item pl-3">
                  As Bard Effect
                </a>
              </eq-window>
            </div>
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
          <td style="vertical-align: middle" class="text-left">
            <eq-spell-effects :spell="spell"/>
          </td>
          <td style="vertical-align: middle" class="text-left">
            <spell-popover
              :spell="spell"
              :size="30"
              placement="auto"
              :offset="-100"
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
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import EqSpellEffects from "@/components/preview/EQSpellEffects";
import EqSpellPreview from "@/components/preview/EQSpellCardPreview.vue";
import EqSpellDescription from "@/components/preview/EQSpellDescription";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT} from "@/app/constants/eq-classes-constants";
import SpellPopover from "@/components/SpellPopover";
import {WindowManager} from "@/app/window";

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
  mounted() {
    setTimeout(() => {
      WindowManager.resizeFillScreenElements()
    }, 100);
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
  }
}
</script>

<style scoped>

/* For Mobile */
@media screen and (max-width: 540px) {
  .spell-preview-table {
    overflow-x: visible;
    overflow-y: scroll !important
  }
}

/* For Tablets */
@media screen and (min-width: 540px) and (max-width: 780px) {
  .spell-preview-table {
    overflow-x: visible;
    overflow-y: scroll !important
  }
}
</style>
