<template>
  <div>
    <div class="row">
      <div class="col-12 col-xl-4">

        <!-- Card -->
        <div class="card">
          <div class="card-header">
            Logs
          </div>
          <div class="card-body p-3">

            <input
              type="text"
              class="form-control mb-4"
              placeholder="Filter"
              @keyup="filterLogs"
              v-model="filter"
            >

            <div style="max-height: 67vh; overflow-y: scroll">
              <p
                @click="viewLogTrigger(log)"
                v-for="log in filteredLogs"
                :key="log"
                style="font-size: 14px; color: rgb(35, 100, 210)"
                class="truncate"
              ><i class="fa fa-file"></i> {{ log.split('logs/')[1] }}</p>
            </div>

          </div>
        </div>

      </div>

      <div class="col-12 col-xl-8">
        <app-loader :is-loading="!loaded"/>

        <div class="card" v-if="loaded">
          <div class="card-header" v-if="currentLogName">
            {{ currentLogName }}
          </div>
          <div class="card-body" v-if="!logOutput">
            Select a log to view its contents
          </div>

          <div class="card-footer bg-dark pb-0 p-0" v-if="logOutput">
            <pre
              id="log-output" class="highlight html bg-dark hljs xml mb-0"
              style="color: #569CD6; max-height: 76vh; overflow-y: scroll"
            >{{ logOutput }}
            </pre>
          </div>

        </div>

      </div>
    </div>

  </div>
</template>

<script>
import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";
import {ROUTE}         from "@/routes";

export default {
  name: 'FileLogs',
  data() {
    return {
      loaded: false,
      logs: [],
      filteredLogs: [],
      filter: "",
      currentLogName: "",
      logOutput: ""
    }
  },
  watch: {
    '$route.query.log': function (id) {
      this.viewLog()
    }
  },
  async created() {
    let r = await OcculusClient.getServerLogs()
    if (r && r.status === 200) {
      this.logs = r.data
    }

    this.filterLogs()
    this.viewLog()
  },
  methods: {
    filterLogs() {
      let filteredResults = []
      this.logs.files.forEach((file) => {
        if (this.filter && this.filter !== '' && !file.includes(this.filter.toLowerCase())) {
          return;
        }

        filteredResults.push(file);
      })

      this.filteredLogs = filteredResults;
    },
    viewLogTrigger(log) {
      this.$router.push(
        {
          path: ROUTE.ADMIN_FILE_LOGS,
          query: {
            log: log
          }
        }
      )
    },
    async viewLog() {
      this.loaded = false
      this.logOutput = ""

      setTimeout(async () => {
        if (!this.$route.query.log) {
          this.loaded = true
          return false
        }

        const log = this.$route.query.log

        let r = await OcculusClient.getServerLog(log)
        if (r && r.status === 200) {
          this.logOutput = r.data.fileContents
        }

        this.currentLogName = log.split('logs/')[1].trim();

        this.loaded = true
      }, 100)
    }
  }
}
</script>

<style scoped>
.truncate {
  width: 450px !important;
  white-space: nowrap;
  overflow: hidden;
  font-size: 14px;
  margin-bottom: 0;


  text-overflow: ellipsis;
  cursor: pointer;
}


</style>
