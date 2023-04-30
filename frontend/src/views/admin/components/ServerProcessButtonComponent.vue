<template>
  <div class="dropdown">

    <a
      href="#"
      style="color: white"
      class="dropdown-toggle btn btn-primary lift"
      role="button"
      data-toggle="dropdown"
      aria-haspopup="true"
      aria-expanded="false"
    >
      <i class="fe fe-power"></i> Power
    </a>

    <div class="dropdown-menu dropdown-menu-right">
      <a href="#" @click="startServerModal" class="dropdown-item">
        <small class="text-muted"><i class="fa fa-keyboard-o" aria-hidden="true"></i> (p)</small>
        Power On
      </a>
      <a href="#" @click="stopServerModal" class="dropdown-item">
        <small class="text-muted"><i class="fa fa-keyboard-o" aria-hidden="true"></i> (s)</small>
        Power Off
      </a>
      <a href="#" @click="restartServerModal" class="dropdown-item">
        <small class="text-muted"><i class="fa fa-keyboard-o" aria-hidden="true"></i> (r)</small>
        Restart [r]
      </a>
      <a href="#" @click="cancelServerRestartModal" class="dropdown-item">
        <small class="text-muted"><i class="fa fa-keyboard-o" aria-hidden="true"></i> (c)</small>
        Cancel Restart [c]
      </a>
    </div>

    <!-- Start Server -->
    <b-modal
      centered
      id="start-server-modal"
      :size="startModalSize"
      title="Start Server"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Start Server"
      @hidden="resetPreflight()"
      @ok="startServer"
    >
      <LauncherOptions :launcherConfig="launcher" v-if="!preflight"/>

      <eq-window title="Pre-Flight Checks" v-if="preflight" class="fade-in">

        <table
          class="eq-table eq-highlight-rows bordered mt-3"
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
              <error-mark-animated v-if="!p.checkSuccess && (p.message && !p.message.includes('Running'))" style="height: 30px; width: 30px" class="ml-1"/>
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
              style="width: 100%; height: 40vh; word-wrap: break-word; white-space: pre-wrap; overflow-x: hidden"
              v-html="p.console"
            ></pre>
          </eq-tab>
        </eq-tabs>

      </eq-window>


      <template #modal-footer="{ ok, cancel, hide }">
        <b-button size="sm" variant="outline-primary" @click="runPreflightChecks()" v-if="!preflight">
          <i class="fa fa-check-circle"></i>
          Start Server Pre-Flight Checks
        </b-button>
        <b-button size="sm" variant="primary" @click="ok()" :disabled="!havePreflightChecksRan()">
          <i class="fa fa-rocket"></i> Start Server
        </b-button>
        <b-button size="sm" variant="outline-secondary" @click="hide('forget')">
          <i class="fa fa-remove"></i> Close
        </b-button>
      </template>

    </b-modal>

    <!-- Stop Server -->
    <b-modal
      centered
      no-fade
      id="stop-server-modal"
      title="Stop Server"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Stop Server"
      @ok="stopServer"
    >
      <p class="pt-3 pb-3">Are you sure you want to stop your server?</p>

      <b-card
        header="Stop Announcement Warning"
      >
        <b-card-text>
          <b-form-radio v-model="delayedStop" name="some-radios" value="0">None</b-form-radio>
          <b-form-radio v-model="delayedStop" name="some-radios" :value="5 * 60">5 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedStop" name="some-radios" :value="10 * 60">10 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedStop" name="some-radios" :value="15 * 60">15 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedStop" name="some-radios" :value="30 * 60">30 Minute(s)</b-form-radio>
        </b-card-text>
      </b-card>
    </b-modal>

    <!-- Restart Server -->
    <b-modal
      centered
      no-fade
      id="restart-server-modal"
      title="Restart Server"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Restart Server"
      @ok="restartServer"
    >

      <p class="pt-3 pb-3">Are you sure you want to restart your server?</p>

      <LauncherOptions :launcherConfig="launcher"></LauncherOptions>

      <b-card
        header="Restart Announcement Warning"
      >
        <b-card-text>
          <b-form-radio v-model="delayedRestart" name="some-radios" value="0">None</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="5 * 60">5 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="10 * 60">10 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="15 * 60">15 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="30 * 60">30 Minute(s)</b-form-radio>
        </b-card-text>
      </b-card>

    </b-modal>

    <!-- Cancel Restart Server -->
    <b-modal
      centered
      no-fade
      id="cancel-restart-server-modal"
      title="Cancel Server Restart"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Cancel Restart"
      @ok="cancelRestartServer"
    >
      Are you sure you want to cancel your timed restart?

    </b-modal>

  </div>
</template>

<script>

import {OcculusClient}   from "@/app/api/eqemu-admin-client-occulus";
import {EventBus}        from "@/app/event-bus/event-bus";
import LauncherOptions   from "@/views/admin/components/LauncherOptions.vue";
import {HttpStream}      from "@/app/httpstream/http-stream";
import EqWindow          from "@/components/eq-ui/EQWindow.vue";
import {debounce}        from "@/app/utility/debounce";
import CheckMarkAnimated from "@/components/CheckMarkAnimated.vue";
import EqTabs            from "@/components/eq-ui/EQTabs.vue";
import EqTab             from "@/components/eq-ui/EQTab.vue";
import ErrorMarkAnimated from "@/components/ErrorMarkAnimated.vue";

const Convert = require('ansi-to-html');
const convert = new Convert();

export default {
  name: 'ServerProcessButtonComponent',
  components: {
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
  created() {
    const pattern = [
      '[\\u001B\\u009B][[\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]+)*|[a-zA-Z\\d]+(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\\u0007)',
      '(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-nq-uy=><~]))'
    ].join('|');

    this.ansiRegex = new RegExp(pattern);
  },

  async mounted() {
    const result  = await OcculusClient.getLauncherConfig();
    this.launcher = result.data;
  },
  methods: {
    startServerModal() {
      this.$root.$emit('bv::show::modal', 'start-server-modal')
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
      // await OcculusClient.stopServer()

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
        }
        catch (e) {
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

    startServer(e) {
      OcculusClient.startServer();
      this.notify("Server Start", "Server is starting!");
      this.notifyProcessChange()
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
      this.$root.$emit('bv::show::modal', 'stop-server-modal')
    },
    stopServer() {
      OcculusClient.stopServer({ timer: this.delayedStop });
      if (this.delayedStop > 0) {
        this.notify("Server Stopped", "Server delayed stop timer started!");
      } else {
        this.notify("Server Stopped", "Server has been stopped!");
      }
      this.notifyProcessChange()
    },

    /**
     * Restart
     */
    restartServerModal() {
      this.$root.$emit('bv::show::modal', 'restart-server-modal')
      this.delayedRestart = 0;
    },
    restartServer() {
      OcculusClient.restartServer({ timer: this.delayedRestart });
      if (this.delayedRestart > 0) {
        this.notify("Server Restarted", "Server restart warning timer has been started!");
      } else {
        this.notify("Server Restarted", "Server has been restarted!");
      }
      this.notifyProcessChange()
    },

    /**
     * Cancel
     */
    cancelServerRestartModal() {
      this.$root.$emit('bv::show::modal', 'cancel-restart-server-modal')
      this.delayedRestart = 0;
    },
    cancelRestartServer() {
      OcculusClient.cancelRestartServer({ cancel: 1 });
      this.notify("Server Restart Cancelled", "Server restart has been cancelled");
    },

    notify(title, message) {
      this.$bvToast.toast(message, {
        title: title,
        solid: true,
        toaster: 'b-toaster-top-center',
      })
    },

    notifyProcessChange() {
      setTimeout(function () {
        EventBus.$emit('process-change');
      }, 1000);

      setTimeout(function () {
        EventBus.$emit('process-change');
      }, 3000);
    },

  },
}
</script>

<style scoped>

</style>
