<template>
  <content-area style="padding-left: 0 !important; padding-right: 0 !important;">
    <eq-window title="Audit log" class="p-0 pt-3">

      <info-error-banner
        :slim="true"
        :notification="notification"
        :error="error"
        @dismiss-error="error = ''"
        @dismiss-notification="notification = ''"
        class="mt-0"
      />

      <app-loader :is-loading="loading"/>

      <div v-if="rows && rows.length === 0 && !loading" class="font-weight-bold pl-0 p-3">
        No entries found
      </div>

      <div style="max-height: 90vh; overflow-y: scroll" v-if="!loading">
        <table
          class="eq-table bordered eq-highlight-rows"
          style="font-size: 14px; "
          v-if="rows && rows.length > 0"
          id="audit-log"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th>ID</th>
            <th style="min-width: 150px" class="text-center">User</th>
            <th>Action</th>
            <th>Description</th>
            <th style="min-width: 150px">Time</th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="e in rows"
            :id="'row-' + e.id"
            :key="'row-' + e.id"
          >
            <td>{{ e.id }}</td>
            <td class="text-center">
              <b-avatar
                style="background-color: transparent;"
                :src="e.user.avatar"
                variant="info"
                v-b-tooltip.hover.v-dark.top
                :title="e.user.user_name"
                size="30"
                class="mr-1"
              />
              {{ e.user.user_name }}
            </td>
            <td>{{ e.event_name }}</td>
            <td>{{ e.data }}</td>
            <td>
              <div
                v-b-tooltip.hover.v-dark.top
                :title="e.created_at"
              >
                {{ formatTime(e.created_at) }}
              </div>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <div class="text-center">
        <b-pagination
          v-if="!loading"
          class="mb-1 mt-1"
          v-model="currentPage"
          :total-rows="totalRows"
          :hide-ellipsis="true"
          per-page="20"
          @change="paginate"
        />
      </div>
    </eq-window>
  </content-area>
</template>

<script>
import ContentArea     from "../../components/layout/ContentArea";
import EqWindow        from "../../components/eq-ui/EQWindow";
import InfoErrorBanner from "../../components/InfoErrorBanner";
import {SpireApi}      from "../../app/api/spire-api";
import moment          from "moment/moment";
import Tablesort       from "../../app/utility/tablesort";
import {ROUTE}         from "../../routes";
import util            from "util";

export default {
  name: "AuditLog",
  components: { InfoErrorBanner, EqWindow, ContentArea },
  data() {
    return {
      rows: [],

      loading: false,

      // notification / errors
      notification: "",
      error: "",

      // pagination (all)
      currentPage: 1,
      totalRows: 0,
    }
  },
  watch: {
    '$route'() {
      this.init()
    },
  },
  mounted() {
    this.init()
  },
  methods: {

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.currentPage > 0) {
        queryState.page = this.currentPage
      }

      const connectionId = this.$route.params.connection
      this.$router.push(
        {
          path: util.format(ROUTE.DATABASE_CONNECTION_AUDIT_LOG, connectionId),
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
      }
    },

    paginate() {
      // models aren't quite updated when we trigger this so queue the pagination
      setTimeout(() => {
        console.log("We're paginating")
        console.log(this.currentPage)
        console.log(this.totalRows)
        this.updateQueryState()
      }, 100)
    },

    formatTime(time) {
      return moment(time).fromNow()
    },

    init() {
      this.totalRows   = 0
      this.currentPage = 1
      this.loadQueryState()
      this.loading = true
      this.rows    = []

      // need to queue to get the loader to show properly
      setTimeout(() => {
        const connectionId = this.$route.params.connection
        SpireApi.v1().get(`/connection/${connectionId}/audit-log`, {
          params: {
            page: this.currentPage
          }
        }).then((r) => {
          if (r.data) {
            this.rows      = r.data.data
            this.totalRows = r.data.total_rows

            if (this.rows.length > 0) {
              setTimeout(() => {
                const target = document.getElementById('audit-log')
                if (target) {
                  new Tablesort(target);
                }
              }, 100)
            }

          }
          this.loading = false
        })
      }, 100)
    }
  }
}
</script>

<style scoped>

</style>
