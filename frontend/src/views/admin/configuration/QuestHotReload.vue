<template>
  <div>

    <div
      class="row justify-content-center"
      style="position: absolute; top: 10%; z-index: 9999999; width: 100%"
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

    <eq-window title="Quest Hot Reload Settings" v-if="configExists()">

      <div class="row">
        <div class="col-3 text-right mt-3">
          <eq-checkbox
            label="Quest Hot Reloading"
            class="d-inline-block"
            :true-value="true"
            :false-value="false"
            v-model="config['web-admin'].quests.hotReload"
            @input="toggleHRM"
          />
        </div>
        <div class="col-9">
          <small class="text-muted">
            This enables Quest Hot reloading within Occulus.
          </small>
          <div class="mt-3">
            When you make changes to your quest files, Occulus will immediately signal your game server to reload the
            appropriate zone's quests for you to help speed up development feedback loops
          </div>
        </div>
      </div>

    </eq-window>

    <eq-window
      title="Hot Reload Rule Settings"
      class="mt-5"
      v-if="configExists()"
    >
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
        <tr v-for="(rule, index) in rules" :key="index">
          <td style="text-align: center">{{ rule.ruleset_id }}</td>
          <td>{{ rule.rule_name }}</td>

          <td class="text-center">

            <label
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
    </eq-window>

  </div>
</template>

<script>
import EqWindow            from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}          from "@/app/api/spire-api";
import EqCheckbox          from "@/components/eq-ui/EQCheckbox.vue";
import {RuleValueApi}      from "@/app/api/api/rule-value-api";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import InfoErrorBanner     from "@/components/InfoErrorBanner.vue";

export default {
  name: "QuestHotReload",
  components: { InfoErrorBanner, EqCheckbox, EqWindow },
  data() {
    return {
      config: {},
      rules: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async created() {
    this.loadRules()
    this.loadConfig()
  },
  methods: {
    configExists() {
      return this.config &&
        Object.keys(this.config).length > 0 &&
        this.config['web-admin'] &&
        this.config['web-admin'].quests &&
        typeof this.config['web-admin'].quests.hotReload !== "undefined"
    },

    async toggleHRM() {
      try {
        const r = await SpireApi.v1().post("admin/serverconfig", this.config)
        if (r.status === 200) {
          this.config = r.data
          this.notification = "Server configuration updated with hot reload settings!"
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },

    async loadConfig() {
      const r = await SpireApi.v1().get("admin/serverconfig")
      if (r.status === 200) {
        this.config = r.data
      }
    },

    async loadRules() {
      let r = await (new RuleValueApi(...SpireApi.cfg())).listRuleValues(
        (new SpireQueryBuilder())
          .where("rule_name", "like", "HotReload")
          .get()
      )
      if (r.status === 200) {
        this.rules = r.data
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
  }
}
</script>

<style scoped>

</style>
