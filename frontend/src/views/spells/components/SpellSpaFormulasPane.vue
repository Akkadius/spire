<template>
  <div>
    <table class="mt-3 spell-preview-table">
      <tbody>
	      <tr>
	        <td class="spell-field-label">Formula ID</td>
	        <td>Description</td>
	      </tr>
        <tr v-for="description, formula_id in SPA_EFFECT_FORMULAS">
          <td class="spell-field-label">{{ formula_id }}</td>
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
  </div>
</template>

<script>
import {SPA_EFFECT_FORMULAS}   from "@/app/constants/eq-spell-constants";
import EqDebug                 from "@/components/eq-ui/EQDebug";
import EqWindow                from "@/components/eq-ui/EQWindow";

export default {
  name: "SpellSpaFormulaPane",
  components: {
    EqDebug,
    "v-runtime-template": () => import("v-runtime-template"),
    EqWindow,
  },
  data() {
    return {
      SPA_EFFECT_FORMULAS: SPA_EFFECT_FORMULAS,
    }
  },
  methods: {
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
