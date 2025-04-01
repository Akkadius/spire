<template>
  <div>

    <eq-window v-if="logCategories.length === 0">
        Zone not loaded
    </eq-window>

    <div class="row" v-if="logCategories.length > 0">
      <div class="col-lg-2 order-lg-0 mb-4 pr-0">
        <eq-window-simple
          class="pt-0 pb-0 mt-0 pl-3 pr-3"
        >
          <div class="row">
            <div class="col-2 p-0" style="margin-top: 4px">
              <button
                type="submit"
                title="Clear all categories"
                class="btn btn-sm btn-white ml-auto"
                @click="clearCategories()"
              >
                <i class="fa fa-eraser pr-1"></i>
              </button>
            </div>
            <div class="col-10 p-0">
              <input
                type="text"
                class="form-control"
                placeholder="Filter categories..."
                v-model="categoryFilter"
                @keyup="updateQueryState()"
                autofocus=""
                value=""
              >
            </div>
          </div>
        </eq-window-simple>

        <eq-window-simple
          class="p-0"
        >
          <div style="height: 76vh; overflow-y: scroll; overflow-x: hidden">
            <label
              class="custom-control custom-checkbox pl-0"
              v-for="(category, index) in filteredLogCategories"
              :key="index"
            >
              <eq-checkbox
                :label-right="category.log_category_description + ' (' + category.log_category_id + ')'"
                :fade-when-not-true="true"
                :true-value="1"
                :false-value="0"
                v-model="category.log_to_console"
                @change="updateCategoryLevel(category)"
              />

            </label>
          </div>
        </eq-window-simple>
      </div>
      <div class="col-lg-10">
        <eq-window-simple
          class="pt-0 pb-0 mt-0"
        >
          <div class="row">
            <div class="col-1 p-0 mt-1 text-center">
              <button
                type="submit"
                title="Clear all logs"
                class="btn btn-sm btn-white ml-auto"
                @click="clearLogs()"
              >
                <i class="fa fa-history pr-1"></i>
              </button>
              <button
                type="submit"
                class="btn btn-sm btn-warning ml-1"
                @click="isPaused = true"
                v-if="!isPaused"
                title="Pause"
              >
                <i class="fa fa-pause pr-1"></i>
              </button>
              <button
                type="submit"
                class="btn btn-sm btn-primary ml-1"
                @click="isPaused = false"
                title="Resume"
                v-if="isPaused"
              >
                <i class="fa fa-play pr-1"></i>
              </button>
              <b-button
                class="btn-white btn-sm ml-1"
                title="Copy to clipboard"
                @click="copyFileContentsToClipboard()"
              >
                <i class="fa fa-copy"></i>
              </b-button>
            </div>

            <div class="col-11">
              <input
                type="text"
                class="form-control"
                v-model="textFilter"
                placeholder="Text filter..."
                @change="updateQueryState()"
                autofocus=""
                value=""
              >
            </div>
          </div>
        </eq-window-simple>

        <eq-window-simple
          class="p-0"
        >
          <pre
            class="highlight html bg-dark hljs xml mb-0"
            id="log-contents"
            style="color: rgb(235 235 235); height: 76vh; overflow-y: scroll; width:100%"
          >{{ logData }}</pre>
        </eq-window-simple>
      </div>
    </div>

  </div>
</template>


<script>

import {EqemuWebsocketClient} from "@/app/api/eqemu-websocket-client";
import EqWindowSimple         from "@/components/eq-ui/EQWindowSimple.vue";
import EqCheckbox             from "@/components/eq-ui/EQCheckbox.vue";
import {Navbar}               from "@/app/navbar";
import {ROUTE}                from "@/routes";
import ClipBoard              from "@/app/clipboard/clipboard";
import EqWindow               from "@/components/eq-ui/EQWindow.vue";
import {Notify}               from "@/app/Notify";

const LOG_STREAM_TRUNCATE_CHARACTER_LENGTH = 300000;

export default {
  components: { EqWindow, EqCheckbox, EqWindowSimple },
  data() {
    return {
      logCategories: [],
      filteredLogCategories: [],
      logData: '',
      zonePort: 0,
      zoneAttributes: null,
      wsEqemuClient: null,
      ws: null,
      panelHeight: 0,

      textFilter: "",
      categoryFilter: "",
      isPaused: false,
    }
  },
  beforeDestroy() {
    Navbar.expand()
  },
  watch: {
    // reset state vars when we navigate
    '$route'() {
      this.loadQueryState()

      this.init()
    },
  },
  async created() {
    this.loadQueryState()

    setTimeout(() => {
      Navbar.collapse()
    }, 100)

    this.$route.meta.title = "Zone Log Streaming"
    if (this.$route.query.zone) {
      this.$route.meta.title = `Zone Log Streaming (${this.$route.query.zone})`
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
        let categories = response.data
        categories.sort(function (a, b) {
          const textA = a.log_category_description.toUpperCase()
          const textB = b.log_category_description.toUpperCase()
          return (textA < textB) ? -1 : (textA > textB) ? 1 : 0
        })
        // set log_to_console to 1 if value is > 0
        categories.forEach((category) => {
          category.log_to_console = category.log_to_console > 0 ? 1 : 0
        })

        this.logCategories = categories

        if (this.categoryFilter && this.categoryFilter !== "") {
          const filter               = this.categoryFilter
          this.filteredLogCategories = categories.filter((category) => {
            return category.log_category_description.toLowerCase().includes(filter)
          })
        } else {
          this.filteredLogCategories = categories
        }
      }

      if (response.method === this.wsEqemuClient.methods.GET_ZONE_ATTRIBUTES) {
        this.zoneAttributes = response.data
      }

      /**
       * Logging
       */
      if (response.event === this.wsEqemuClient.subscriptions.LOGGING && !this.isPaused) {
        const logString = this.logTimeStamp() + ' | ' + response.data.msg + '\n' + this.logData
        let log         = logString.substring(0, LOG_STREAM_TRUNCATE_CHARACTER_LENGTH)

        if (this.textFilter) {
          log = log.split('\n').filter((line) => {
            return line.toLowerCase().includes(this.textFilter.toLowerCase())
          }).join('\n')
        }

        this.logData = log
      }
    }
  },
  destroyed() {
    if (this.ws) {
      this.ws.close()
    }
  },
  methods: {
    clearCategories() {
      for (let category of this.logCategories) {
        if (category.log_category_description.toLowerCase().includes("info")) {
          continue;
        }
        category.log_to_console = 0
        this.updateCategoryLevel(category)
      }
    },
    init() {
      if (this.logCategories.length > 0) {
        if (this.categoryFilter && this.categoryFilter !== "") {
          const filter               = this.categoryFilter
          this.filteredLogCategories = this.logCategories.filter((category) => {
            return category.log_category_description.toLowerCase().includes(filter)
          })
        } else {
          this.filteredLogCategories = this.logCategories
        }
      }
    },

    logTimeStamp: function () {
      const now        = new Date()
      const date       = [now.getMonth() + 1, now.getDate(), now.getFullYear()]
      const time       = [now.getHours(), now.getMinutes(), now.getSeconds()]
      const suffix     = (time[0] < 12) ? 'AM' : 'PM'
      time[0]          = (time[0] < 12) ? time[0] : time[0] - 12
      time[0]          = time[0] || 12
      let milliseconds = now.getMilliseconds()
      // milliseconds always take up same amount of space
      if (milliseconds < 10) {
        milliseconds = '00' + milliseconds
      } else if (milliseconds < 100) {
        milliseconds = '0' + milliseconds
      }

      for (let i = 1; i < 3; i++) {
        if (time[i] < 10) {
          time[i] = '0' + time[i]
        }
      }

      return date.join('/') + ' ' + time.join(':') + milliseconds + ' ' + suffix
    },
    clearLogs: function () {
      this.logData = ''
    },
    updateCategoryLevel: function (category) {
      this.wsEqemuClient.setLoggingLevel(category.log_category_id, (category.log_to_console ? 3 : 0))
    },

    updateQueryState: function () {
      let queryState = {};

      if (this.npcNameSearch !== "") {
        queryState.q = this.textFilter
      }
      if (this.categoryFilter !== "") {
        queryState.category = this.categoryFilter
      }

      this.$router.push(
        {
          path: ROUTE.ADMIN_ZONESERVERS_LOGS.replaceAll(":port", this.$route.params.port),
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState: function () {
      if (this.$route.query.q !== "") {
        this.textFilter = this.$route.query.q
      }
      if (this.$route.query.category !== "") {
        this.categoryFilter = this.$route.query.category
      }
    },

    copyFileContentsToClipboard() {
      ClipBoard.copyFromElement("log-contents");
      Notify.toast("Copied logs to clipboard");
    },
  }
}
</script>
