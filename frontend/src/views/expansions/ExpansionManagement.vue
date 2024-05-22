<template>
  <content-area style="padding: 0 !important">
    <eq-window title="Expansion Management">

      <div class="row">
        <div class="col-2">
          <content-expansion-selector
            v-if="currentExpansion !== -100"
            @input="currentExpansion = $event"
            :show-names="true"
            :value="currentExpansion"
          />
        </div>
      </div>


    </eq-window>
  </content-area>
</template>

<script>
import ContentArea                from "@/components/layout/ContentArea";
import EqWindow                   from "@/components/eq-ui/EQWindow";
import {SpireApi}                 from "@/app/api/spire-api";
import {EXPANSIONS_FULL}          from "@/app/constants/eq-expansions";
import ExpansionIcon              from "@/components/preview/ExpansionIcon";
import ContentExpansionSelector   from "@/components/selectors/ContentExpansionSelector.vue";
import ExpansionBitmaskCalculator from "@/components/tools/ExpansionsCalculator.vue";
import {RuleValueApi}             from "@/app/api/api/rule-value-api";

export default {
  name: "Expansion",
  components: { ExpansionBitmaskCalculator, ContentExpansionSelector, ExpansionIcon, EqWindow, ContentArea },
  data() {
    return {
      tables: [],
      expansionData: [],
      expansions: EXPANSIONS_FULL,

      loadedExpansion: -100, // initial state
      currentExpansion: -100, // can be updated through selection

      rules: [],
    }
  },
  mounted() {
    this.loadRules().then(() => {
      this.currentExpansion = parseInt(this.rules.filter(
        r => r.rule_name === 'Expansion:CurrentExpansion'
      )[0].rule_value)

      this.loadedExpansion = parseInt(this.currentExpansion)

      console.log(this.currentExpansion)
    })
  },
  methods: {
    async loadRules() {
      let r = await (new RuleValueApi(...SpireApi.cfg())).listRuleValues()
      if (r.status === 200) {
        this.rules = r.data
      }
    },
  }
}
</script>

<style>
.sticky-first-column {
  text-align: center;
  position: sticky;
  z-index: 9999;
  background-color: rgb(25, 31, 41);
  left: 0px;
  font-weight: bold;
}

.expansion-data-table th, .expansion-data-table td {
  text-align: center !important;
  padding: 2px !important;
}
</style>
