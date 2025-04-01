<template>
  <div>
    <div
      class="row justify-content-center"
      style="position: absolute; top: -2%; z-index: 9999999; width: 100%"
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

    <!-- make window opacity .5 when loading -->
    <eq-window
      :title="`Release Version (${release})`"
      class="p-0"
      :style="{'opacity': loading ? '.5' : '1'}"
    >

      <div class="p-4 text-center" v-if="crashes.length === 0">
        <span class="font-weight-bold ">
          No crashes found
        </span>
      </div>

      <div
        :style="'max-height: ' + (highlightedId ? 35 : 95) + 'vh; overflow-y: scroll; '"
        v-if="crashes.length > 0"
        id="crash-list-viewport"
      >
        <table
          class="eq-table bordered eq-highlight-rows mb-0"
          style="overflow-x: scroll; min-width: 80vw"

        >
          <thead class="eq-table-floating-header">
          <tr>
            <th style="min-width: 100px"></th>
            <th>ID</th>
            <th>Fingerprint</th>
            <th style="min-width: 100px">Compile Time</th>
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
            :style="{'opacity': c.resolved ? '.5' : '1'}"
            :id="'crash-' + c.id"
          >
            <td class="text-center">
              <b-button
                class="btn-dark btn-sm btn-dark d-inline-block"
                style="padding: 0px 6px;"
                title="View Crash"
                @click="viewCrash(c)"
              >
                <i class="fa fa-eye"></i>
              </b-button>

              <div class="d-inline-block" v-if="user && user.is_server_developer">
                <b-button
                  class="btn-dark btn-sm btn-outline-success d-inline-block ml-3"
                  style="padding: 0px 6px;"
                  title="Mark Crash as Resolved"
                  @click="markCrashResolved(c)"
                  v-if="!c.resolved"
                >
                  <i class="fa fa-check"></i>
                </b-button>

                <b-button
                  class="btn-dark btn-sm btn-outline-danger d-inline-block ml-3"
                  style="padding: 0px 6px;"
                  title="Mark Crash as Unresolved"
                  @click="markCrashUnresolved(c)"
                  v-if="c.resolved"
                >
                  <i class="fa fa-remove"></i>
                </b-button>
              </div>

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
        <div v-if="crash.resolved" class="d-block mb-3">
          <div style="width: 20px;" class="d-inline-block">
            <check-mark-animated style="height: 15px; width: 15px"/>
          </div>
          Resolved by <b>{{ crash.user.user_name }}</b> at <b>{{ crash.resolved_at }}</b>
        </div>

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
import EqWindow          from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}        from "@/app/api/spire-api";
import {ROUTE}           from "@/routes";
import util              from "util";
import UserContext       from "@/app/user/UserContext";
import InfoErrorBanner   from "@/components/InfoErrorBanner.vue";
import CheckMarkAnimated from "@/components/CheckMarkAnimated.vue";

export default {
  name: "Release",
  components: { CheckMarkAnimated, InfoErrorBanner, EqWindow },
  data() {
    return {
      release: "",

      loading: false,

      crashes: [],

      crash: {},
      highlightedId: 0,

      user: {},

      // notification / errors
      notification: "",
      error: "",
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

    await this.load()

    this.loading = false
  },
  methods: {

    async markCrashResolved() {
      try {
        const r = await SpireApi.v1().post(`analytics/server-crash-reports/${this.crash.id}/mark-resolved`)
        if (r.status === 200) {
          this.notification = "Crash marked as resolved!"
          this.loading      = true
          setTimeout(async () => {
            await this.load()
            this.loading = false
          }, 10)
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },

    async markCrashUnresolved() {
      try {
        const r = await SpireApi.v1().post(`analytics/server-crash-reports/${this.crash.id}/mark-unresolved`)
        if (r.status === 200) {
          this.notification = "Crash marked as unresolved!"
          this.loading      = true
          setTimeout(async () => {
            await this.load()
            this.loading = false
          }, 10)
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },

    reset() {
      this.crash         = {}
      this.highlightedId = 0
    },

    async load() {
      const r = await SpireApi.v1().get(`analytics/server-crash-reports`, {
        params: {
          version: this.$route.params.version
        }
      })
      if (r.status === 200) {
        this.crashes = r.data
      }

      this.release = this.$route.params.version
      this.user    = await UserContext.getUser()

      for (let r of this.crashes) {
        if (r.id === this.highlightedId) {
          this.crash = r
        }
      }

      setTimeout(() => {
        const container = document.getElementById("crash-list-viewport");
        const target    = document.getElementById("crash-" + this.crash.id);
        if (container && target) {
          container.scrollTop = target.offsetTop - 100;
        }
      }, 100)
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
