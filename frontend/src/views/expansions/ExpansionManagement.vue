<template>
  <content-area style="padding: 0 !important">
    <eq-window title="Expansion Management">

      <div class="row">
        <div class="col-2">
          <small class="text-muted">
            You can use up and down arrow keys to navigate through expansions
          </small>
          <content-expansion-selector
            class="mt-3"
            v-if="selectedExpansion !== -100"
            @input="selectedExpansion = $event; updateQueryState()"
            :show-names="true"
            :value="selectedExpansion"
          />
        </div>
        <div class="col-10" style="border-left: 1px solid gray">
          <div v-if="selectedExpansion >= 0" :key="selectedExpansion">
            <div>
              <div class="font-weight-bold d-inline-block mr-1">Expansion</div>
              {{ selectedExpansionData.expansion_name }} ({{ selectedExpansionData.expansion_number }})
            </div>
            <div>
              <div class="font-weight-bold d-inline-block mr-1">Max Level</div>
              {{ selectedExpansionData.max_level }}
            </div>
          </div>

          <table class="eq-table eq-highlight-rows bordered log-settings minified-inputs mt-3">
            <thead class="eq-table-floating-header">
            <tr>
              <th style="width: 200px">Rule</th>
              <th style="width: 100px">Value</th>
              <th>Comment</th>
            </tr>

            </thead>
            <tbody>
            <tr
              v-for="r in selectedExpansionData.rules"
              :key="r.name"
            >
              <td class="font-weight-bold">{{ r.name }}</td>
              <td>{{ r.value }}</td>
              <td>{{ r.comment }}</td>
            </tr>
            </tbody>
          </table>

          <eq-debug :data="selectedExpansionData"></eq-debug>
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
import {ROUTE}                    from "@/routes";
import EqDebug                    from "@/components/eq-ui/EQDebug.vue";
import EqCheckbox                 from "@/components/eq-ui/EQCheckbox.vue";

export default {
  name: "Expansion",
  components: {
    EqCheckbox,
    EqDebug,
    ExpansionBitmaskCalculator,
    ContentExpansionSelector,
    ExpansionIcon,
    EqWindow,
    ContentArea
  },
  data() {
    return {
      tables: [],
      expansions: EXPANSIONS_FULL,

      loadedExpansion: -100,
      selectedExpansion: -100, // can be updated through selection
      selectedExpansionData: {},

      rules: [],
      expansionData: [],
    }
  },
  watch: {
    '$route'() {
      this.loadQueryState()
      this.init()
    },
  },
  async mounted() {
    this.loadQueryState()
    this.init()

    document.addEventListener('keydown', this.arrowKeyHandler);
  },
  beforeDestroy() {
    document.removeEventListener('keydown', this.arrowKeyHandler);
  },
  methods: {
    async init() {
      // load rules if not loaded
      if (this.rules && this.rules.length === 0) {
        await this.loadRules()
      }

      const currentExpansion = this.rules.filter(
        r => r.rule_name === 'Expansion:CurrentExpansion'
      )[0].rule_value

      this.loadedExpansion = parseInt(currentExpansion)

      if (this.selectedExpansion === -100) {
        this.selectedExpansion = currentExpansion
        this.updateQueryState()
      }

      // load if not loaded
      if (this.expansionData.length === 0) {
        const r = await SpireApi.v1().get("/expansions")
        if (r.data) {
          this.expansionData = r.data
        }
      }

      this.selectedExpansionData = this.expansionData.filter(
        e => e.expansion_number === this.selectedExpansion
      )[0]

      // as we progress through each expansion, rules are stacked and the
      // same named rule can be overridden
      let rules = []
      for (let e of this.expansionData) {
        if (e.expansion_number > this.selectedExpansion) {
          break
        }

        if (!e.rules) {
          continue
        }
        for (let rule of e.rules) {
          let existing = rules.filter(r => r.name === rule.name)
          if (existing.length === 0) {
            rules.push(rule)
          } else {
            // replace the existing rule with the new one
            rules = rules.map(r => {
              if (r.name === rule.name) {
                return rule
              }
              return r
            })
          }
        }
      }

      // sort this.selectedExpansionData.rules by this.selectedExpansionData.rules.name
      rules.sort((a, b) => {
        if (a.name < b.name) {
          return -1
        }
        if (a.name > b.name) {
          return 1
        }
        return 0
      })

      if (rules && rules.length > 0) {
        this.selectedExpansionData.rules = rules
      }

    },

    async loadRules() {
      let r = await (new RuleValueApi(...SpireApi.cfg())).listRuleValues()
      if (r.status === 200) {
        this.rules = r.data
      }
    },

    // state
    updateQueryState() {
      let q = {};
      if (this.selectedExpansion !== -100) {
        q.expansion = this.selectedExpansion
      }

      this.$router.push(
        {
          path: ROUTE.EXPANSIONS_MANAGEMENT,
          query: q
        }
      ).catch(() => {
      })
    },
    loadQueryState() {
      if (this.$route.query.expansion && this.$route.query.expansion.length > 0) {
        this.selectedExpansion = parseInt(this.$route.query.expansion)
      }
    },

    arrowKeyHandler(event) {
      switch (event.key) {
        case "ArrowUp":
          event.preventDefault()

          this.selectedExpansion = this.selectedExpansion - 1
          if (this.selectedExpansion < 0) {
            this.selectedExpansion = 0
          }
          this.updateQueryState()
          break;
        case "ArrowDown":
          event.preventDefault()

          this.selectedExpansion = this.selectedExpansion + 1

          const max = Object.keys(this.expansions).length
          // cap by max expansion
          console.log("expansion length", )

          if (this.selectedExpansion >= max) {
            this.selectedExpansion = max - 1
          }

          this.updateQueryState()
          break;
        case "ArrowLeft":
          console.log("Left arrow pressed");
          break;
        case "ArrowRight":
          console.log("Right arrow pressed");
          break;
        default:
          // Do nothing for other keys
          break;
      }
    }
  }
}
</script>

<style>
.expansion-data-table th, .expansion-data-table td {
  text-align: center !important;
  padding: 2px !important;
}
</style>
