<template>
  <div class="col-12 p-0">

    <div class="card" v-if="logCategories.length === 0">
      <div class="card-body">
          Zone not loaded
      </div>
    </div>

    <div class="row" v-if="logCategories.length > 0">
      <div class="col-lg-2 order-lg-0 mb-4 pr-0">
        <div class="card">
          <div class="card-header">
            <h4>Logging Categories</h4>
          </div>
          <div class="card-body p-3" style="max-height: 75vh; overflow-y: scroll">
            <div class="form-group">
              <div class="custom-controls-stacked" style="font-size: 11px;">
                <label
                  class="custom-control custom-checkbox" v-for="(category, index) in logCategories"
                  :key="index"
                >
                  <input
                    type="checkbox"
                    class="custom-control-input"
                    @change="updateCategoryLevel(category)"
                    v-model="category.log_to_console"
                  >

                  <span class="custom-control-label">
              {{ category.log_category_description }}
              ({{ category.log_category_id }})
            </span>
                </label>
              </div>
            </div>

          </div>
        </div>

      </div>

      <div class="col-lg-10">
        <div class="card">
          <div class="card-header"><h3 class="card-title">Logs</h3>
            <button type="submit" class="btn btn-primary ml-auto" @click="clearLogs()">
              <i class="fa fa-history pr-1"></i>
              Clear Logs
            </button>
          </div>
          <div class="card-body" >
            <div class="row">
              <div class="col-lg-12">
                <div class="card-footer bg-dark pb-0">
                  <pre
                    class="highlight html bg-dark hljs xml mb-0"
                    style="color: rgb(235 235 235);; height: 68vh; overflow-y: scroll;"
                  >{{ logData }}</pre>

                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>

import {EqemuWebsocketClient} from "@/app/api/eqemu-websocket-client";

const LOG_STREAM_TRUNCATE_CHARACTER_LENGTH = 300000;

export default {
  data() {
    return {
      logCategories: [],
      logData: '',
      zonePort: 0,
      zoneAttributes: null,
      wsEqemuClient: null,
      ws: null,
      panelHeight: 0
    }
  },
  async created() {
    console.log(this.$route)

    this.$route.meta.title = "Zone Log Streaming"
    if (this.$route.query.zone) {
      console.log(this.$route)
      this.$route.meta.title = `Zone Log Streaming (${this.$route.query.zone})`
      // this.$forceUpdate()
      // console.log(this.$route.meta)

    }

    this.zonePort = this.$route.params.port

    /**
     * Setup client and fetch logging categories
     */
    this.wsEqemuClient = new EqemuWebsocketClient()
    this.ws = await this.wsEqemuClient.initClient(this.zonePort)
    this.wsEqemuClient.subscribeToLogging()
    this.wsEqemuClient.getLogCategories()

    /**
     * Message handling
     */
    this.ws.onmessage = (event) => {
      const response = JSON.parse(event.data)

      /**
       * Categories
       */
      if (response.method === this.wsEqemuClient.methods.GET_LOGSYS_CATEGORIES) {
        this.logCategories = response.data
        this.logCategories.sort(function (a, b) {
          const textA = a.log_category_description.toUpperCase()
          const textB = b.log_category_description.toUpperCase()
          return (textA < textB) ? -1 : (textA > textB) ? 1 : 0
        })
      }

      if (response.method === this.wsEqemuClient.methods.GET_ZONE_ATTRIBUTES) {
        this.zoneAttributes = response.data
        console.log(this.zoneAttributes)
      }

      /**
       * Logging
       */
      if (response.event === this.wsEqemuClient.subscriptions.LOGGING) {
        const logString = this.logTimeStamp() + ' ' + response.data.msg + '\n' + this.logData
        this.logData    = logString.substring(0, LOG_STREAM_TRUNCATE_CHARACTER_LENGTH)
      }
    }
  },
  destroyed() {
    this.ws.close()
  },
  methods: {
    logTimeStamp: function () {
      const now    = new Date()
      const date   = [now.getMonth() + 1, now.getDate(), now.getFullYear()]
      const time   = [now.getHours(), now.getMinutes(), now.getSeconds()]
      const suffix = (time[0] < 12) ? 'AM' : 'PM'
      time[0]      = (time[0] < 12) ? time[0] : time[0] - 12
      time[0]      = time[0] || 12

      for (var i = 1; i < 3; i++) {
        if (time[i] < 10) {
          time[i] = '0' + time[i]
        }
      }

      return date.join('/') + ' ' + time.join(':') + ' ' + suffix
    },
    clearLogs: function () {
      this.logData = ''
    },
    updateCategoryLevel: function (category) {
      this.wsEqemuClient.setLoggingLevel(category.log_category_id, (category.log_to_console ? 3 : 0))
    }
  }
}
</script>
