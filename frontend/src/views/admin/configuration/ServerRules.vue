<template>
  <div>

    <div class="card">
      <div class="card-body">
        <div class="row justify-content-between align-items-center">
          <div class="col-12 col-md-9 col-xl-7">
            <h2 class="">
              Server Rules
            </h2>
            <p class="text-muted mb-md-0">
              Server rules can change the behavior of your server
            </p>
          </div>
        </div>
      </div>
    </div>

    <div class="card">
      <div class="card-body p-3">
        <input
          type="text"
          class="form-control form-control-prepended list-search mb-3"
          @keyup="updateQueryState()"
          v-model="search"
          placeholder="Search rules..."
        >

        <div class="row">
          <div class="col-12">
            <div
              class="card mb-0"
            >
              <div
                class="table-responsive mb-0"
                style="max-height: 61vh; overflow-y: scroll"
              >
                <table
                  class="table table-sm table-nowrap card-table rule_table"
                >
                  <thead>
                  <tr>
                    <th class="sorting" tabindex="0" style="width: 60px">ID</th>
                    <th class="sorting" tabindex="0" style="width: 320px !important">Rule Name</th>
                    <th class="sorting" tabindex="0" style="width: 250px">Value</th>
                    <th class="sorting" tabindex="0">Description</th>
                  </tr>
                  </thead>
                  <tbody>

                  <!-- Loop through rules -->
                  <tr v-for="(rule, index) in filteredRules" :key="index">
                    <td class="">{{ rule.ruleset_id }}</td>
                    <td>{{ rule.rule_name }}</td>

                    <td>

                      <label
                        class="pb-2 pt-2"
                        v-if="rule.rule_value === 'true' || rule.rule_value === 'false'"
                      >
                        <b-form-checkbox
                          value="true"
                          unchecked-value="false"
                          :checked="rule.rule_value"
                          v-model="rule.rule_value"
                          @change="updateRule(rule)"
                          switch
                        />
                      </label>

                      <span v-if="!(rule.rule_value === 'true' || rule.rule_value === 'false')">
                      <input
                        type="text" class="form-control"
                        v-model="rule.rule_value"
                        @change="updateRule(rule)"
                      >
                  </span>

                      <!-- {{rule.rule_value}} -->
                    </td>

                    <td class="text-muted" style="overflow: auto; white-space: normal;">
                      {{ rule.notes }}
                    </td>

                  </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {EqemuAdminClient} from "@/app/api/eqemu-admin-client-occulus";
import {ROUTE}            from "@/routes";

export default {
  data() {
    return {
      search: "",

      rules: [],
      filteredRules: [],
      loaded: false,
    }
  },

  watch: {
    '$route'() {
      this.loadQueryState()
      this.filterRules()
    },
  },

  async created() {
    this.rules = await EqemuAdminClient.getServerRules();
    this.filterRules();
    this.loaded = true;
    // this.initTable()

    this.loadQueryState()
    this.filterRules()
  },
  methods: {

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

    async updateRule(rule) {
      const response = await EqemuAdminClient.postServerRule(rule);
      if (response.success) {
        this.$bvToast.toast(
          response.success,
          {
            title: "Rule updated!",
            toaster: 'b-toaster-bottom-right',
            autoHideDelay: 3000,
            solid: true,
            appendToast: false
          }
        )
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
