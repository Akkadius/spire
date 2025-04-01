<template>
  <div>
    <div class="row">
      <div class="col-lg-3 col-sm-12">
        <eq-window title="Current Version" style="height: 100%" class="p-0">
          <table class="eq-table eq-highlight-rows bordered m-0 mt-3" v-if="version">
            <tbody>
            <tr class="fade-in" v-if="version.os">
              <td class="text-right font-weight-bold">Operating System</td>
              <td>{{ version.os }}</td>
            </tr>
            <tr class="fade-in" v-if="version.server_version">
              <td class="text-right font-weight-bold">Server Version</td>
              <td>{{ version.server_version }}</td>
            </tr>
            <tr class="fade-in" v-if="version.compile_date">
              <td class="text-right font-weight-bold">Compile Date</td>
              <td>{{ version.compile_date }}</td>
            </tr>
            <tr class="fade-in" v-if="version.compile_time">
              <td class="text-right font-weight-bold">Compile Time</td>
              <td>{{ version.compile_time }}</td>
            </tr>
            <tr class="fade-in" v-if="version.database_version">
              <td class="text-right font-weight-bold">Database Version</td>
              <td>{{ version.database_version }}</td>
            </tr>
            <tr class="fade-in" v-if="version.bots_database_version">
              <td class="text-right font-weight-bold">Bots Database Version</td>
              <td>{{ version.bots_database_version }}</td>
            </tr>
            </tbody>
          </table>

        </eq-window>
      </div>
      <div class="col-lg-9 col-sm-12">
        <eq-window title="Update Settings" style="height: 100%">
          <div class="row">
            <div class="col-3 text-right mt-3">
              <eq-checkbox
                label="Use Release Binaries"
                class="d-inline-block"
                :true-value="BUILD_TYPE.RELEASE"
                :false-value="BUILD_TYPE.SELF_COMPILED"
                v-model="updateType"
                @input="setUpdateOption(BUILD_TYPE.RELEASE)"
              />
            </div>
            <div class="col-9">
              <small class="text-muted">
                Uses official release binaries from
                <a href="https://github.com/EQEmu/Server/releases" target="releases">EverQuest Emulator Server</a>
              </small>
              <div class="mt-1">
                Use this when you are not developing server or making code changes to server code
              </div>
            </div>
          </div>

          <div class="row mt-3" v-if="os === 'linux'">
            <div class="col-3 text-right mt-3">
              <eq-checkbox
                label="Self-Compiled Binaries"
                class="d-inline-block"
                :false-value="BUILD_TYPE.RELEASE"
                :true-value="BUILD_TYPE.SELF_COMPILED"
                v-model="updateType"
                @input="setUpdateOption(BUILD_TYPE.SELF_COMPILED)"
              />
            </div>
            <div class="col-9">
              <small class="text-muted">
                Compiles binaries locally <i class="fa fa-linux"></i> (Linux only)
              </small>
              <div class="mt-1">
                Use this if you are a developer or intend on making modifications to your server
              </div>
            </div>
          </div>

          <div
            class="row mt-4"
            v-if="updateType === 'self-compiled'"
          >
            <div class="col-3 text-center">
              <button class="btn btn-dark mt-4 btn-sm" @click="buildSource()" :disabled="buildRunning">
                <i class="fa fa-wrench"></i> Build
              </button>
              <button class="btn btn-outline-primary mt-4 btn-sm ml-3" @click="buildClean()" :disabled="buildRunning">
                <i class="fa fa-refresh"></i> Clean
              </button>
              <button class="btn btn-dark mt-4 btn-sm ml-3" @click="buildCancel()">
                <i class="fa fa-remove"></i> Cancel
              </button>
            </div>
            <div
              :style="(buildRunning ? 'opacity: .5' : 'opacity: 1')"
              class="col-3 text-center"
            >
              <span class="font-weight-bold">Branch</span>
              <b-input-group>
                <b-select v-model="currentBranch" :options="branches" :disabled="buildRunning"/>
                <b-input-group-append>

                  <b-button variant="white" class="btn-sm" @click="setBranch" :disabled="buildRunning">
                    <i class="fa fa-dot-circle-o mr-2"></i>
                    Set
                  </b-button>

                </b-input-group-append>
              </b-input-group>
            </div>
            <div
              :style="(buildRunning ? 'opacity: .5' : 'opacity: 1')"
              class="col-3 text-center"
            >
              <span class="font-weight-bold">Source Location</span>
              <input
                type="text"
                class="form-control"
                v-model="sourceLocation"
                @change="updateSourceLocation()"
                :disabled="buildRunning"
              >
            </div>
            <div class="col-2 text-center">
              <span class="font-weight-bold">Make Tool</span>
              <input
                type="text"
                class="form-control"
                disabled
                v-model="makeTool"
                style="opacity: .5"
                :disabled="buildRunning"
              >
            </div>
            <div
              :style="(buildRunning ? 'opacity: .5' : 'opacity: 1')"
              class="col-1 text-center"
            >
              <span class="font-weight-bold">Cores</span>
              <input
                type="text"
                class="form-control"
                v-model.number="cores"
                @change="updateBuildCores()"
                :disabled="buildRunning"
              >
              <span class="text-muted">(Jobs)</span>
            </div>
          </div>

          <div
            class="row justify-content-center"
            style="position: absolute; bottom: -5%; z-index: 9999999; width: 100%"
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
        </eq-window>
      </div>
    </div>

    <update-releases
      class="mt-5"
      :version="version.server_version"
      @refresh-version="getVersions()"
      v-if="updateType === BUILD_TYPE.RELEASE && version"
    />

    <eq-window title="Build Output" class="mt-5 p-1" v-show="output && updateType === BUILD_TYPE.SELF_COMPILED">
      <pre
        class="mt-3 fade-in mb-1"
        id="output"
        style="width: 100%; height: 50vh; word-wrap: break-word; white-space: pre-wrap; overflow-x: hidden"
        v-html="output"
      ></pre>

      <eq-progress-bar :percent="buildPercent" v-if="buildRunning"/>

      <div
        class="row justify-content-center"
        style="position: absolute; bottom: -5%; z-index: 9999999; width: 100%"
      >
        <div class="col-6">
          <info-error-banner
            style="width: 100%"
            :slim="true"
            :notification="buildNotification"
            :error="buildError"
            @dismiss-error="buildError = ''"
            @dismiss-notification="buildNotification = ''"
            class="mt-3"
          />
        </div>
      </div>
    </eq-window>

  </div>
</template>

<script>
import EqWindow           from "@/components/eq-ui/EQWindow.vue";
import EqCheckbox         from "@/components/eq-ui/EQCheckbox.vue";
import {SpireApi}         from "@/app/api/spire-api";
import InfoErrorBanner    from "@/components/InfoErrorBanner.vue";
import UpdateReleases     from "@/views/admin/server-update/UpdateReleases.vue";
import {AppEnv}           from "@/app/env/app-env";
import UserContext        from "@/app/user/UserContext";
import {debounce}         from "@/app/utility/debounce";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";
import EqProgressBar      from "@/components/eq-ui/EQProgressBar.vue";
import {Navbar}           from "@/app/navbar";

const Convert = require('ansi-to-html');
const convert = new Convert();

function readChunks(reader) {
  return {
    async* [Symbol.asyncIterator]() {
      let readResult = await reader.read();
      while (!readResult.done) {
        yield readResult.value;
        readResult = await reader.read();
      }
    },
  };
}


export default {
  name: "ServerUpdate",
  components: { EqProgressBar, LoaderFakeProgress, UpdateReleases, InfoErrorBanner, EqCheckbox, EqWindow },
  data() {
    return {
      updateType: "",

      version: {},

      releases: [],

      os: AppEnv.getOS(),

      // notification / errors
      notification: "",
      error: "",

      // self-compilation
      sourceLocation: "",
      makeTool: "",
      cores: 0,
      branches: [],
      currentBranch: "",
      buildRunning: false,
      buildPercent: 0,
      buildNotification: "",
      buildError: "",

      BUILD_TYPE: {
        RELEASE: 'release',
        SELF_COMPILED: 'self-compiled'
      },

    }
  },
  beforeDestroy() {
    Navbar.expand();
  },
  beforeRouteLeave(to, from, next) {
    if (this.buildRunning) {
      if (!window.confirm("You have a build running, are you sure you want to leave?")) {
        return;
      }
    }
    next();
  },
  created() {
    this.updateType = this.BUILD_TYPE.RELEASE

    this.init()

    this.output = ""

    // windows default
    if (this.os !== 'linux') {
      if (AppEnv.getSettingValue('SERVER_UPDATE_TYPE') !== this.BUILD_TYPE.RELEASE) {
        AppEnv.setSetting('SERVER_UPDATE_TYPE', this.BUILD_TYPE.RELEASE)
      }
    }

    setTimeout(() => {
      this.outputContainer = document.getElementById("output");
    }, 1000)

    const pattern = [
      '[\\u001B\\u009B][[\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]+)*|[a-zA-Z\\d]+(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\\u0007)',
      '(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-nq-uy=><~]))'
    ].join('|');

    this.ansiRegex = new RegExp(pattern);

    Navbar.collapse();
  },
  methods: {
    updateSourceLocation() {
      try {
        AppEnv.setSetting("BUILD_LOCATION", this.sourceLocation)
        this.notification = `Updated build location to [${this.sourceLocation}]`
        this.init()
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },
    updateBuildCores() {
      try {
        AppEnv.setSetting("BUILD_CORES", this.cores)
        this.notification = `Updated build cores to [${this.cores}]`
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },
    renderOutput: debounce(function () {
      this.$forceUpdate()

      setTimeout(() => {
        if (this.outputContainer) {
          this.outputContainer.scrollTop = this.outputContainer.scrollHeight + 100;
        }
      }, 1)

    }, 10),

    async buildSource() {

      this.buildRunning = true
      this.output       = "Sending job to build\n"
      this.$forceUpdate()

      fetch(SpireApi.getBasePath() + '/api/v1/eqemuserver/build', {
        method: 'post',
        body: JSON.stringify({
          "source_directory": this.sourceLocation,
          "build_tool": this.makeTool,
          "cores": this.cores,
        }),
        headers: {
          Authorization: `Bearer ` + UserContext.getAccessToken(),
          'Content-Type': "application/json",
        },
      }).then(async (response) => {
        const textDecoder = new TextDecoder();
        const reader      = response.body.getReader();
        for await (const chunk of readChunks(reader)) {
          const chunkText = textDecoder.decode(chunk);
          this.output += chunkText

          if (chunkText.includes("[")) {
            let s       = chunkText
            s           = s.replaceAll("[", "")
            s           = s.replaceAll("]", "")
            const split = s.split(" ")
            if (split.length > 0) {
              const first = split[0].trim()
              if (first.includes("%")) {
                this.buildPercent = first.replaceAll("%", "")
              } else if (first.includes("/")) {
                const ninjaSplit = first.split("/")
                if (ninjaSplit.length > 0) {
                  const progress    = ninjaSplit[0].trim()
                  const total       = ninjaSplit[1].trim()
                  this.buildPercent = parseInt((progress / total) * 100)
                }
              }
            }
          }

          if (this.outputContainer) {
            this.outputContainer.scrollTop = this.outputContainer.scrollHeight;
          }
          this.renderOutput()
        }
      }).finally(() => {
        this.buildNotification = "Build complete!"
        this.buildRunning      = false
      });
    },

    async buildClean() {
      this.output = "Sending clean to system\n"
      this.$forceUpdate()

      fetch(SpireApi.getBasePath() + '/api/v1/eqemuserver/build-clean', {
        method: 'post',
        body: JSON.stringify({
          "source_directory": this.sourceLocation,
          "build_tool": this.makeTool,
        }),
        headers: {
          Authorization: `Bearer ` + UserContext.getAccessToken(),
          'Content-Type': "application/json",
        },
      }).then(async (response) => {
        const textDecoder = new TextDecoder();
        const reader      = response.body.getReader();
        for await (const chunk of readChunks(reader)) {
          const chunkText = textDecoder.decode(chunk);
          this.output += chunkText + "\n"

          if (this.outputContainer) {
            this.outputContainer.scrollTop = this.outputContainer.scrollHeight;
          }
          this.renderOutput()
        }
      }).finally(() => {
        this.buildNotification = "Clean complete!"
      });
    },

    async buildCancel() {
      this.output = "Sending cancel to system\n"
      this.$forceUpdate()

      fetch(SpireApi.getBasePath() + '/api/v1/eqemuserver/build-cancel', {
        method: 'post',
        body: JSON.stringify({
          "source_directory": this.sourceLocation,
          "build_tool": this.makeTool,
        }),
        headers: {
          Authorization: `Bearer ` + UserContext.getAccessToken(),
          'Content-Type': "application/json",
        },
      }).then(async (response) => {
        const textDecoder = new TextDecoder();
        const reader      = response.body.getReader();
        for await (const chunk of readChunks(reader)) {
          const chunkText = textDecoder.decode(chunk);
          this.output += chunkText + "\n"

          if (this.outputContainer) {
            this.outputContainer.scrollTop = this.outputContainer.scrollHeight;
          }
          this.renderOutput()
        }
      }).finally(() => {
        this.buildNotification = "Cancel complete!"
      });
    },

    async init() {
      await this.getVersions()

      const r = await SpireApi.v1().get(`eqemuserver/update-type`)
      if (r.status === 200) {
        this.updateType = r.data.updateType
      }

      if (AppEnv.getOS().includes("linux")) {
        if (this.updateType === this.BUILD_TYPE.SELF_COMPILED) {
          await this.getBuildInfo()
          await this.getBranchInfo()
        }
      }
    },

    async getVersions() {
      try {
        const v = await SpireApi.v1().get(`eqemuserver/version`)
        if (v.status === 200) {
          this.version    = v.data
          this.version.os = AppEnv.getOS()
        }
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async setUpdateOption(option) {
      try {
        const r = await SpireApi.v1().post(`eqemuserver/update-type/${option}`)
        if (r.status === 200) {
          this.notification = r.data.message
          this.init()
        }
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async getBuildInfo() {
      const r = await SpireApi.v1().get(`eqemuserver/get-build-info`)
      if (r.status === 200) {
        // always pull the make tool from server
        this.makeTool = r.data.build_tool

        // we auto look for the build location regardless
        // we only use it if something else wasn't manually set
        const b = AppEnv.getSetting("BUILD_LOCATION")
        if (!b || (b && b.value === "")) {
          await AppEnv.setSetting("BUILD_LOCATION", r.data.source_directory)
          this.sourceLocation = r.data.source_directory
        } else {
          this.sourceLocation = b.value
        }

        this.cores = parseInt(AppEnv.getSettingValue("BUILD_CORES", 4))
      }
    },
    async getBranchInfo() {
      this.branches      = []
      this.currentBranch = ""

      SpireApi.v1().get(`eqemuserver/build/branches`).then((r) => {
        if (r.status === 200) {
          this.branches = r.data
        }
      }).catch((e) => {
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      })
      SpireApi.v1().get(`eqemuserver/build/current-branch`).then((r) => {
        if (r.status === 200) {
          this.currentBranch = r.data
        }
      }).catch((e) => {
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      })
    },
    setBranch() {
      if (confirm(`Are you sure you want to switch to this branch? Any pending changes on current branch will be lost`)) {
        SpireApi.v1().post(`eqemuserver/build/branch/${this.currentBranch}`).then((r) => {
          if (r.status === 200) {
            this.notification = `Branch has been set to [${this.currentBranch}]`
            this.init()
          }
        }).catch((e) => {
          if (e.response && e.response.data && e.response.data.error) {
            this.error = e.response.data.error
          }
        })
      }
    }
  }
}
</script>

<style scoped>

</style>
