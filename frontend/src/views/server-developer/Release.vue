<template>
  <div>
    <app-loader :is-loading="loading"/>

    <eq-window
      :title="`Release Version (${release})`"
      class="p-0"
      v-if="!loading"
    >
      <div class="p-4 text-center" v-if="crashes.length === 0">
        <span class="font-weight-bold ">
          No crashes found
        </span>
      </div>

      <div
        :style="'max-height: ' + (highlightedId ? 35 : 95) + 'vh; overflow-y: scroll; '"
        v-if="crashes.length > 0"
      >
        <table
          class="eq-table bordered eq-highlight-rows mb-0"
          style="overflow-x: scroll; min-width: 80vw"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th></th>
            <th>ID</th>
            <th>Fingerprint</th>
            <th>Compile Time</th>
            <th>Server Name</th>
            <th>Process</th>
            <th>PID</th>
            <th>OS</th>
            <th>Lines</th>
            <th>Created</th>
          </tr>
          </thead>
          <tbody>
          <tr
            @click="viewCrash(c)"
            v-for="(c, index) in crashes"
            :key="c.id"
            :class="(highlightedId && c.id === highlightedId ? 'pulsate-highlight-white' : '')"
          >
            <td class="text-center">
              <b-button
                class="btn-dark btn-sm btn-outline-warning"
                style="padding: 0px 6px;"
                title="View Crash"
                @click="viewCrash(c)"
              >
                <i class="fa fa-eye"></i>
              </b-button>
            </td>
            <td>{{ c.id }}</td>
            <td style="text-align: center">{{ c.fingerprint.substr(0, 5) }}</td>

            <td>{{ c.compile_date }} {{ c.compile_time }}</td>
            <td>{{ c.server_name }} ({{ c.server_short_name }})</td>
            <td>{{ c.platform_name }}</td>
            <td>{{ c.process_id }}</td>
            <td>{{ c.os_sysname }} ({{ c.os_machine }})</td>
            <td>{{ c.crash_report.split("\n").length }}</td>
            <td>{{ c.created_at ? c.created_at : "" }}</td>
          </tr>
          </tbody>
        </table>
      </div>
    </eq-window>

    <eq-window
      class="mt-4"
      :title="`Crash Stack (${highlightedId})`"
      v-if="crash && crash.id"
    >
      <div class="pb-3">
        <b>Fingerprint</b> {{ crash.fingerprint }}
        <b>OS</b> {{ crash.os_version }} ({{ crash.os_machine }}) {{ crash.os_release }}
        <b v-if="crash.origination_info">Origination</b> {{ crash.origination_info }}
      </div>

      <pre
        style="width: 100%; height: 50vh; overflow-y: scroll"
      >{{ crash.crash_report }}</pre>
    </eq-window>
  </div>
</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "@/app/api/spire-api";
import {ROUTE}    from "@/routes";
import util       from "util";

export default {
  name: "Release",
  components: { EqWindow },
  data() {
    return {
      release: "",

      loading: false,

      crashes: [],

      crash: {},
      highlightedId: 0
    }
  },

  watch: {
    $route(to, from) {
      this.reset()
      this.loadQueryState()
      this.load()
    }
  },

  async mounted() {
    this.loadQueryState()

    this.loading = true

    const r = await SpireApi.v1().get(`analytics/server-crash-reports`, {
      params: {
        version: this.$route.params.version
      }
    })
    if (r.status === 200) {
      this.crashes = r.data
    }

    this.release = this.$route.params.version

    this.load()

    this.loading = false
  },
  methods: {
    reset() {
      this.crash         = {}
      this.highlightedId = 0
    },

    async load() {
      for (let r of this.crashes) {
        if (r.id === this.highlightedId) {
          this.crash = r
        }
      }
    },
    viewCrash(e) {
      this.crash         = e
      this.highlightedId = e.id
      this.updateQueryState()
    },
    updateQueryState: function () {
      let queryState = {};

      if (typeof this.highlightedId !== "undefined") {
        queryState.id = this.highlightedId
      }

      this.$router.push(
        {
          path: util.format(ROUTE.RELEASE, this.$route.params.version),
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState: function () {
      if (parseInt(this.$route.query.id) !== 0) {
        this.highlightedId = parseInt(this.$route.query.id)
      }
    },
  }
}
</script>

<style scoped>

</style>
