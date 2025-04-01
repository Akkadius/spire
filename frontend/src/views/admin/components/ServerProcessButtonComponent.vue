<template>
  <div class="dropdown">
    <a
      href="#"
      data-toggle="dropdown"
      aria-haspopup="true"
      aria-expanded="false"
      :class="serverStatus === 'Online' ? 'text-success' : 'text-danger'"
    >
      <i class="fe fe-power"></i> {{ serverStatus }}
    </a>

    <eq-window class="dropdown-menu dropdown-menu-left p-0">
      <a href="#" @click="startServerModal" class="dropdown-item pl-3">
        <i class="fa fa-keyboard-o" aria-hidden="true"></i> (p)
        Power On
      </a>
      <a href="#" @click="stopServerModal" class="dropdown-item pl-3">
        <i class="fa fa-keyboard-o" aria-hidden="true"></i> (s)
        Power Off
      </a>
      <a href="#" @click="restartServerModal" class="dropdown-item pl-3">
        <i class="fa fa-keyboard-o" aria-hidden="true"></i> (r)
        Restart [r]
      </a>
      <a href="#" @click="cancelServerRestartModal" class="dropdown-item pl-3">
        <i class="fa fa-keyboard-o" aria-hidden="true"></i> (c)
        Cancel Restart [c]
      </a>
    </eq-window>

    <!-- Start Server -->
    <EqModal
      v-show="showStartServerModal"
      title="Start Server"
      @close="showStartServerModal = false; resetPreflight()"
    >
      <template v-slot:header></template>
      <template v-slot:body>

        <div style="width: 500px">
          <LauncherOptions :launcherConfig="launcher" v-if="!preflight"/>
        </div>

        <!-- Preflight Checks -->
        <div style="width: 1000px" v-if="preflight">
          <table
            class="eq-table eq-highlight-rows bordered mt-3 fade-in"
          >
            <thead class="eq-table-floating-header">
            <tr>
              <th class="text-center" style="width: 120px">Status</th>
              <th class="text-center" style="width: 140px">Process</th>
              <th class="text-left">Messages</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="p in processTypes" :key="p.name">
              <td class="text-center">
                <check-mark-animated style="height: 20px; width: 20px" v-if="p.checkSuccess"/>
                <b-spinner small v-if="!p.checkSuccess && (!p.message || p.message.includes('Running'))"/>
                <error-mark-animated
                  v-if="!p.checkSuccess && (p.message && !p.message.includes('Running'))"
                  style="height: 30px; width: 30px"
                  class="ml-1"
                />
              </td>
              <td class="text-center">{{ p.desc }} {{ (!p.required ? '*' : '') }}</td>
              <td><span v-if="p.message" v-html="p.message"></span></td>
            </tr>
            </tbody>
          </table>

          <div class="">
            * denotes a process that is not required to start the server.
          </div>

          <eq-tabs @on-selected="scrollPreflight($event)">
            <eq-tab :name="p.desc" :selected="p.name === 'world'" v-for="p in processTypes" :key="p.name">
            <pre
              :id="'preflight-output-' + p.name"
              style="width: 100%; height: 40vh; word-wrap: break-word; white-space: pre-wrap; overflow-x: hidden; text-align: left;"
              v-html="p.console"
            ></pre>
            </eq-tab>
          </eq-tabs>

        </div>

      </template>

      <template v-slot:footer>
        <b-button size="sm" variant="dark" @click="runPreflightChecks()" v-if="!preflight" class="mr-3">
          <i class="fa fa-check-circle"></i>
          Start Server Pre-Flight Checks
        </b-button>
        <b-button
          size="sm"
          variant="primary"
          @click="startServer(); showStartServerModal = false; resetPreflight()"
          :disabled="!havePreflightChecksRan()"
          class="mr-3"
        >
          <i class="fa fa-rocket"></i> Start Server
        </b-button>
        <b-button size="sm" variant="outline-secondary" @click="showStartServerModal = false; resetPreflight()">
          <i class="fa fa-remove"></i> Close
        </b-button>
      </template>
    </EqModal>

    <!-- Stop Server -->
    <EqModal
      v-show="showStopServerModal"
      title="Stop Server"
      @close="showStopServerModal = false"
    >
      <template v-slot:body>

        <p class="pt-3 pb-3">Are you sure you want to stop your server?</p>


        <table class="eq-table bordered mb-3">
          <thead class="eq-table-floating-header">
          <tr>
            <td class="font-weight-bold p-3">Stop Delay Options</td>
          </tr>
          </thead>

          <tbody>
          <tr v-for="opt in stopDelays" :key="opt.value">
            <td>
              <eq-checkbox
                :fade-when-not-true="true"
                class="d-inline-block mr-3"
                :true-value="opt.value"
                :false-value="null"
                v-model="delayedStop"
                @change="updateStopDelay(opt.value)"
              />
              {{ opt.label }}
            </td>
          </tr>
          </tbody>
        </table>
      </template>

      <template v-slot:footer>
        <b-button size="sm" variant="danger" @click="stopServer(); showStopServerModal = false" class="mr-3">
          <i class="fa fa-power-off"></i> Stop Server
        </b-button>
        <b-button size="sm" variant="outline-secondary" @click="showStopServerModal = false">
          <i class="fa fa-remove"></i> Cancel
        </b-button>
      </template>
    </EqModal>


    <!-- Restart Server -->
    <EqModal
      v-show="showRestartServerModal"
      title="Restart Server"
      @close="showRestartServerModal = false"
    >
      <template v-slot:body>

        <div style="width: 500px">
          <p class="pt-3 pb-3">Are you sure you want to restart your server?</p>

          <LauncherOptions :launcherConfig="launcher"/>

          <table class="eq-table bordered mb-0">
            <thead class="eq-table-floating-header">
            <tr>
              <td class="font-weight-bold p-3">Restart Delay Options</td>
            </tr>
            </thead>

            <tbody>
            <tr v-for="opt in stopDelays" :key="opt.value">
              <td>
                <eq-checkbox
                  :fade-when-not-true="true"
                  class="d-inline-block mr-3"
                  :true-value="opt.value"
                  :false-value="null"
                  v-model="delayedRestart"
                  @change="updateStopDelay(opt.value)"
                />
                {{ opt.label }}
              </td>
            </tr>
            </tbody>
          </table>

        </div>
      </template>

      <template v-slot:footer>
        <div class="mt-3">
          <b-button size="sm" variant="primary" @click="restartServer(); showRestartServerModal = false" class="mr-3">
            <i class="fa fa-refresh"></i> Restart Server
          </b-button>
          <b-button size="sm" variant="outline-secondary" @click="showRestartServerModal = false">
            <i class="fa fa-remove"></i> Cancel
          </b-button>
        </div>
      </template>
    </EqModal>

    <!-- Cancel Restart Server -->
    <EqModal
      v-show="showCancelRestartModal"
      title="Cancel Server Restart"
      @close="showCancelRestartModal = false"
    >
      <template v-slot:body>
        <div class="p-3">
        Are you sure you want to cancel your timed restart?
        </div>
      </template>

      <template v-slot:footer>
        <button class="btn btn-danger btn-sm mr-3" @click="cancelRestartServer(); showCancelRestartModal = false">
          <i class="fa fa-times-circle"></i> Cancel Restart
        </button>
        <button class="btn btn-outline-secondary btn-sm" @click="showCancelRestartModal = false">
          <i class="fa fa-remove"></i> Close
        </button>
      </template>
    </EqModal>

  </div>
</template>

<script>

import {EventBus}        from "@/app/event-bus/event-bus";
import LauncherOptions   from "@/views/admin/components/LauncherOptions.vue";
import {HttpStream}      from "@/app/httpstream/http-stream";
import EqWindow          from "@/components/eq-ui/EQWindow.vue";
import {debounce}        from "@/app/utility/debounce";
import CheckMarkAnimated from "@/components/CheckMarkAnimated.vue";
import EqTabs            from "@/components/eq-ui/EQTabs.vue";
import EqTab             from "@/components/eq-ui/EQTab.vue";
import ErrorMarkAnimated from "@/components/ErrorMarkAnimated.vue";
import {SpireApi}        from "@/app/api/spire-api";
import {Notify}          from "@/app/Notify";
import EqModal           from "@/components/eq-ui/EQModal.vue";
import EqCheckbox        from "@/components/eq-ui/EQCheckbox.vue";

const Convert = require('ansi-to-html');
const convert = new Convert();

export default {
  name: 'ServerProcessButtonComponent',
  components: {
    EqCheckbox,
    EqModal,
    ErrorMarkAnimated,
    EqTab,
    EqTabs,
    CheckMarkAnimated,
    EqWindow,
    LauncherOptions
  },
  data() {
    return {
      delayedRestart: 0,
      delayedStop: 0,
      launcher: {},

      cancelledRestart: false,

      showStartServerModal: false,
      showRestartServerModal: false,
      showStopServerModal: false,
      showCancelRestartModal: false,

      stopDelays: [
        { label: "None", value: 0 },
        { label: "1 Minute", value: 60 },
        { label: "5 Minute(s)", value: 5 * 60 },
        { label: "10 Minute(s)", value: 10 * 60 },
        { label: "15 Minute(s)", value: 15 * 60 },
        { label: "30 Minute(s)", value: 30 * 60 },
      ],

      startModalSize: "md",

      scrollTimer: null,

      preflight: false,
      processTypes: [
        { name: "world", desc: "World", checkSuccess: false, console: "", required: true },
        { name: "zone", desc: "Zone", checkSuccess: false, console: "", required: true },
        { name: "shared_memory", desc: "Shared Memory", checkSuccess: false, console: "", required: true },
        { name: "ucs", desc: "UCS (Chat)", checkSuccess: false, console: "", required: false },
        { name: "loginserver", desc: "Loginserver", checkSuccess: false, console: "", required: false },
      ],
    }
  },

  props: {
    serverStatus: {
      default: 'Online',
      required: false
    },
  },

  created() {
    const pattern = [
      '[\\u001B\\u009B][[\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]+)*|[a-zA-Z\\d]+(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\\u0007)',
      '(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-nq-uy=><~]))'
    ].join('|');

    this.ansiRegex = new RegExp(pattern);

    window.addEventListener('keypress', this.keypressHandler)
  },

  beforeDestroy() {
    window.removeEventListener('keypress', this.keypressHandler);
  },

  async mounted() {
    try {
      const r = await SpireApi.v1().get('admin/launcherconfig')
      if (r.status === 200) {
        if (r.data) {
          this.launcher = r.data
        }
      }
    } catch (e) {
    }
  },
  methods: {
    updateStopDelay(value) {
      this.delayedStop = value;
    },

    startServerModal() {
      this.$root.$emit('bv::show::modal', 'start-server-modal')
      this.showStartServerModal = true;
    },

    scrollPreflight: debounce(function (e) {
      setTimeout(() => {
        const t = document.getElementById('preflight-output-' + e)
        if (t) {
          t.scrollTop = t.scrollHeight + 100;
        }
      }, 1)
    }, 10),

    resetPreflight() {
      this.startModalSize = "md"
      this.preflight      = false
      for (let [i, p] of this.processTypes.entries()) {
        this.processTypes[i].checkSuccess = false
        delete this.processTypes[i].message
      }
    },

    havePreflightChecksRan() {
      for (let p of this.processTypes) {
        if (typeof p.message === "undefined" || p.message.length === 0) {
          return false;
        }
      }
      return true;
    },

    async runPreflightChecks() {
      this.startModalSize = "xl"
      this.preflight      = true

      // zero out the console and message
      for (let [i, p] of this.processTypes.entries()) {
        this.processTypes[i].checkSuccess = false
        this.processTypes[i].console      = ""
        delete this.processTypes[i].message
      }

      // run the preflight checks
      for (let [i, p] of this.processTypes.entries()) {
        this.processTypes[i].message = "Running preflight checks..."

        try {
          HttpStream.get("/api/v1/eqemuserver/pre-flight/" + p.name.toLowerCase()).then(async (r) => {
            for await (const m of HttpStream.read(r)) {
              this.processTypes[i].console += this.getStreamFormatted(m)
              if (m.includes("QueryErr") || m.includes("Error")) {
                this.processTypes[i].message = this.getStreamFormatted(m).split("\n").filter((e) => {
                  return e.includes("QueryErr") || e.includes("Error")
                }).join("\n")
              } else if (m.includes("no such file or directory")) {
                this.processTypes[i].message = "Could not find the " + p.name + " executable. Please check your configuration."
              }

              this.$forceUpdate()
              this.scrollPreflight(p.name)
            }
          })
            .finally(() => {
              // if we have no message, then we succeeded
              if (!this.processTypes[i].message || this.processTypes[i].message.includes("Running")) {
                this.processTypes[i].checkSuccess = true
                this.processTypes[i].message      = "Checks succeeded"
              }
              this.$forceUpdate()
            });
        } catch (e) {
          // failed for some reason
          if (!this.processTypes[i].message) {
            this.processTypes[i].checkSuccess = false
            this.processTypes[i].message      = "Failed to run preflight checks"

            if (e.response && e.response.data && e.response.data.error) {
              this.processTypes[i].message += e.response.data.error
            }
          }
          this.$forceUpdate()
        }
      }

    },

    async startServer(e) {
      try {
        await SpireApi.v1().post('eqemuserver/server/start')
        this.notify("Server Start", "Server is starting!");
        this.notifyProcessChange()
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.notify("Launcher Error", e.response.data.error);
        }
      }
    },

    getStreamFormatted(m) {
      if (this.ansiRegex.test(m)) {
        return convert.toHtml(m)
      } else {
        return m
      }
    },

    /**
     * Stop
     */
    stopServerModal() {
      this.showStopServerModal = true
      // this.$root.$emit('bv::show::modal', 'stop-server-modal')
    },
    async stopServer() {
      try {
        await SpireApi.v1().post('eqemuserver/server/stop', { timer: this.delayedStop })
        this.delayedStop = 0;

        if (this.delayedStop > 0) {
          this.notify("Server Stopped", "Server delayed stop timer started!");
        } else if (!this.cancelledRestart) {
          this.notify("Server Stopped", "Server has been stopped!");
        }

      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.notify("Launcher Error", e.response.data.error);
        }
      }

      this.notifyProcessChange()
    },

    /**
     * Restart
     */
    restartServerModal() {
      this.showRestartServerModal = true
      this.delayedRestart         = 0;
    },
    async restartServer() {
      try {
        await SpireApi.v1().post('eqemuserver/server/restart', { timer: this.delayedRestart })
        this.delayedRestart = 0;
      } catch (e) {
        console.log(e)
      }

      if (this.delayedRestart > 0) {
        this.notify("Server Restarted", "Server restart warning timer has been started!");
      } else if (!this.cancelledRestart) {
        this.notify("Server Restarted", "Server has been restarted!");
      }
      this.notifyProcessChange()
    },

    /**
     * Cancel
     */
    cancelServerRestartModal() {
      this.showCancelRestartModal = true
      this.delayedRestart = 0;
    },
    async cancelRestartServer() {
      this.cancelledRestart = true;
      try {
        await SpireApi.v1().post('eqemuserver/server/stop-cancel')
      } catch (e) {
        console.log(e)
      }

      this.notify("Server Restart Cancelled", "Server stop or restart has been cancelled");

      // we do this because we have an open request to the server and it will falsely say the server
      // has been restarted or stopped when it hasn't
      setTimeout(() => {
        this.cancelledRestart = false;
      }, 3000)
    },

    notify(title, message) {
      Notify.toast(message);
    },

    notifyProcessChange() {
      setTimeout(function () {
        EventBus.$emit('process-change');
      }, 1000);

      setTimeout(function () {
        EventBus.$emit('process-change');
      }, 3000);
    },


    keypressHandler(e) {
      if (e.srcElement.tagName !== 'BODY' && e.srcElement.tagName !== 'A') {
        return
      }

      if (window.location.pathname === '/login') {
        return
      }

      switch (String.fromCharCode(e.keyCode)) {
        case 'p':
          // this.$root.$emit('bv::show::modal', 'start-server-modal')
          this.startServerModal()
          break
        case 'r':
          this.restartServerModal()
          break
        case 'c':
          this.cancelServerRestartModal();
          break
        case 's':
          this.stopServerModal();
          break
      }
    },

  },
}
</script>

<style scoped>

</style>
