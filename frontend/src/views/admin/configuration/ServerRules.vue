<template>
  <eq-window>

    <div
      class="row justify-content-center"
      style="position: absolute; top: 0%; z-index: 9999999; width: 100%"
    >
      <div class="col-6">
        <info-error-banner
          style="width: 100%"
          :slim="true"
          :notification="notification"
          :error="error"
          @dismiss-error="error = ''"
          @dismiss-notification="notification = ''"
          class="mt-3"
        />
      </div>
    </div>

    <div style="max-height: 86vh; overflow-y: scroll; overflow-x: hidden">

      <div class="row">
        <div class="col-11">
          <b-form-input
            type="text"
            class="form-control list-search"
            autofocus
            @keyup="updateQueryState()"
            v-model="search"
            placeholder="Search rules..."
          />
        </div>

        <div class="col-1">
          <button
            title="Reset"
            class="btn m-0"
            @click="search = ''; updateQueryState()"
          ><i class="fa fa-refresh"></i> Reset
          </button>
        </div>
      </div>

      <table
        class="eq-table eq-highlight-rows bordered log-settings mt-3"
      >
        <thead class="eq-table-floating-header">
        <tr>
          <th tabindex="0" style="width: 60px">RID</th>
          <th tabindex="0" style="width: 320px !important">Rule Name</th>
          <th tabindex="0" style="width: 250px" class="text-center">Value</th>
          <th tabindex="0">Description</th>
        </tr>
        </thead>
        <tbody>

        <!-- Loop through rules -->
        <tr v-for="(rule, index) in filteredRules" :key="index">
          <td style="text-align: center">{{ rule.ruleset_id }}</td>
          <td>{{ rule.rule_name }}</td>

          <td class="text-center">

            <label
              class="mb-0"
              v-if="rule.rule_value === 'true' || rule.rule_value === 'false'"
            >
              <eq-checkbox
                :fade-when-not-true="true"
                class="d-inline-block mt-2"
                true-value="true"
                false-value="false"
                v-model="rule.rule_value"
                @change="updateRule(rule)"
              />
            </label>

            <span
              v-if="!(rule.rule_value === 'true' || rule.rule_value === 'false')"
            >
                <input
                  type="text" class="form-control"
                  v-model="rule.rule_value"
                  @change="updateRule(rule)"
                >
            </span>

            <!-- {{rule.rule_value}} -->
          </td>

          <td style="overflow: auto; white-space: normal;">
            {{ rule.notes }}
          </td>

        </tr>
        </tbody>
      </table>
    </div>

  </eq-window>

</template>

<script>
import {ROUTE}             from "@/routes";
import EqWindow            from "@/components/eq-ui/EQWindow.vue";
import EqCheckbox          from "@/components/eq-ui/EQCheckbox.vue";
import {SpireApi}          from "@/app/api/spire-api";
import {RuleValueApi}      from "@/app/api/api/rule-value-api";
import InfoErrorBanner     from "@/components/InfoErrorBanner.vue";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export default {
  components: { InfoErrorBanner, EqCheckbox, EqWindow },
  data() {
    return {
      search: "",

      rules: [],
      filteredRules: [],
      loaded: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },

  watch: {
    '$route'() {
      this.loadQueryState()
      this.filterRules()
    },
  },

  async created() {
    await this.loadRules()

    this.filterRules();
    this.loaded = true;
    // this.initTable()

    this.loadQueryState()
    this.filterRules()
  },
  methods: {

    async loadRules() {
      let r = await (new RuleValueApi(...SpireApi.cfg())).listRuleValues({limit: 10000})
      if (r.status === 200) {
        this.rules = r.data
      }
    },

    updateQueryState() {
      let q = {};

      if (this.search !== "") {
        q.search = this.search
      }

      this.$router.push(
        {
          path: ROUTE.ADMIN_CONFIG_SERVER_RULES,
          query: q
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (this.$route.query.search && this.$route.query.search.length > 0) {
        this.search = this.$route.query.search
      }
    },

    async updateRule(e) {
      try {
        const r = await (new RuleValueApi(...SpireApi.cfg())).updateRuleValue(
          {
            id: e.ruleset_id,
            ruleValue: e
          },
          {
            query: (new SpireQueryBuilder())
              .where("ruleset_id", "=", e.ruleset_id)
              .where("rule_name", "=", e.rule_name)
              .get()
          }
        )
        if (r.status === 200) {
          this.notification = `Updated rule (${e.ruleset_id}) [${e.rule_name}] to value (${e.rule_value})!`
          this.loadRules()

          const r = await SpireApi.v1().post("eqemuserver/reload/rules")
          if (r.status === 200) {
            setTimeout(() => {
              this.notification = "Server rules reloaded in-game!"
            }, 1000)
          }
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }

    },
    filterRules() {
      if (!this.search) {
        this.filteredRules = this.rules;
        return;
      }

      let filteredRules = [];
      this.rules.forEach(row => {
        if (
          row.rule_name.toLowerCase().includes(this.search.toLowerCase()) ||
          row.notes.toLowerCase().includes(this.search.toLowerCase())
        ) {
          filteredRules.push(row)
        }
      });

      this.filteredRules = filteredRules
    }
  }
}
</script>

<style>
.rule_table table {
  position: relative;
}

.rule_table thead {
  position: sticky;
  top: 0;
  z-index: 9999;
}
</style>
