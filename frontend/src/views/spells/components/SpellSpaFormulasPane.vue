<template>
  <div>
    <eq-window
      v-if="SPA_EFFECT_FORMULAS"
      id="zone-view-container"
      style="height: 95vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="formulatable"
        class="eq-table eq-highlight-rows bordered"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead class="eq-table-floating-header">
          <tr>
            <th></th>
            <th>Formula ID</th>
            <th>Description</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="description, formula_id in SPA_EFFECT_FORMULAS">
            <td>
              <b-button
                class="btn-dark btn-sm btn-dark"
                @click="selectFormula(formula_id)"
                v-if="!formula_id.includes('-')"
              >
                <i class="fa fa-arrow-left"></i>
              </b-button>
            </td>
            <td>{{ formula_id }}</td>
            <td>{{ description }}</td>
          </tr>
        </tbody>
      </table>
      <div>
        <h3>Notes:</h3>
        <ul>
          <li>Level = Caster Level</li>
          <li>BuffCalc = CalcBuffDuration_formula(level, buff_duration_formula, buff_duration) - Tics - 1</li>
        </ul>
      </div>
    </eq-window>
  </div>
</template>

<script>
import {SPA_EFFECT_FORMULAS} from "@/app/constants/eq-spell-constants";
import EqDebug               from "@/components/eq-ui/EQDebug";
import EqWindow              from "@/components/eq-ui/EQWindow";

export default {
  name: "SpellSpaFormulaPane",
  components: {
    EqDebug,
    EqWindow,
  },
  data() {
    return {
      SPA_EFFECT_FORMULAS: SPA_EFFECT_FORMULAS,
    }
  },
  methods: {
    selectFormula(formula_id) {
      this.$emit('input', {
        formulaId: formula_id
      });
    },
    toTitleCase(str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    },
  }
}
</script>
