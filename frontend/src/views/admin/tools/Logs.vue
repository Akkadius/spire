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
                v-for="log in filteredLogs" :key="log"
                class="truncate"
              >{{ log.split('logs/')[1] }}</p>
            </div>

          </div>
        </div>

      </div>

      <div class="col-12 col-xl-8">
        <div class="card">
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
import util               from 'util'
import {EqemuAdminClient} from "@/app/api/eqemu-admin-client-occulus";

export default {
  name: 'Logs',
  data() {
    return {
      logs: null,
      filteredLogs: null,
      filter: null,
      currentLogName: null,
      logOutput: null
    }
  },
  watch: {
    '$route.query.log': function (id) {
      this.viewLog()
    }
  },
  async created() {
    setTimeout(() => {
      if (document.getElementsByClassName("content-area").length > 0) {
        const container = document.getElementsByClassName("content-area")[0];
        container.setAttribute('style', 'max-width: 95% !important');
      }
    }, 10)

    let r = await EqemuAdminClient.getServerLogs()
    if (r && r.status === 200) {
      this.logs = r.data
    }

    this.filterLogs()
    this.viewLog()
  },
  beforeDestroy() {
    if (document.getElementsByClassName("content-area").length > 0) {
      const container          = document.getElementsByClassName("content-area")[0];
      container.style.maxWidth = null;
    }
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
          path: '/admin/tools/logs',
          query: {
            log: log
          }
        }
      )
    },
    async viewLog() {
      if (!this.$route.query.log) {
        return false
      }

      const log = this.$route.query.log

      let r = await EqemuAdminClient.getServerLog(log)
      if (r && r.status === 200) {
        this.logOutput = r.data.fileContents
      }

      this.currentLogName = log.split('logs/')[1].trim();
    }
  }
}
</script>

<style scoped>
.truncate {
  width: 450px !important;
  white-space: nowrap;
  overflow: hidden;
  color: blue;
  font-size: 14px;
  margin-bottom: 0;


  text-overflow: ellipsis;
  cursor: pointer;
}


</style>
