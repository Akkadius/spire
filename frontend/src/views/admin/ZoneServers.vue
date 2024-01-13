<template>
  <div :style="loaded ? 'opacity:1' : 'opacity:.3'">
    <div v-if="zoneList.length === 0">
      <div class="card">
        <div class="card-body">
          Zoneservers are offline
        </div>
      </div>
    </div>

    <div
      class="row justify-content-center"
      style="position: absolute; top: 100px; z-index: 9999999; width: 100%"
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

    <div class="card mb-3 mr-4" v-if="zoneList.length > 0 && loaded">
      <div class="card-body pt-0 pb-0 pl-3 pr-3">
        <div class="input-group input-group-flush">
          <div class="input-group-prepend">
            <span class="input-group-text"><i class="fe fe-search"></i></span>
          </div>
          <input
            class="list-search form-control"
            type="search"
            @keyup="updateQueryState"
            v-model="search"
            placeholder="Filter zones..."
            autofocus
          >
        </div>
      </div>
    </div>

    <div style="max-height: 80vh; overflow-y: scroll">
      <div
        :class="['card', 'mb-3']"
        v-if="getProcessStats(zone.zone_os_pid)"
        v-for="zone in filterZoneList(zoneList)"
      >
        <div
          :class="['card-body', 'lift', 'btn-default', 'card-slim']"
          style="box-shadow: 0 2px 4px 0 rgba(30,55,90,.1);"
        >
          <div class="row align-items-center">

            <div class="col ml-n1">
              <h4 class="mb-1">

                <span class="h2 fe text-muted mb-0 fe-layers mr-3"></span>

                <a href="#" style="color:#2364d2;font-weight:200;font-size: 18px;" v-if="zone.zone_long_name">
                  {{ formatZoneName(zone.zone_long_name ? zone.zone_long_name : 'Standby Zone') }}
                </a>

                <span
                  style="color:#666;font-weight:200;font-size: 18px;"
                  class="text-muted"
                  v-if="!zone.zone_long_name"
                >
                Ready for players...
              </span>

              </h4>
            </div>

            <div class="col-auto" v-if="zone.zone_name">
              <p class="card-text small text-muted mb-1 mt-2">
              <span class="text-muted ml-2" style="font-size:12px">
                  {{ formatZoneName(zone.zone_name) }}
                  ({{ zone.zone_id }})
                  <span v-if="zone.instance_id">Instance: {{ zone.instance_id }}</span>
                </span>
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
              <span class="text-muted ml-2" style="font-size:12px">
                  {{ zone.is_static_zone === true ? 'Static' : 'Dynamic' }}
              </span>
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
                <i class="fe fe-proc"></i> PID {{ zone.zone_os_pid }}
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
                <i class="fe fe-cloud"></i> Port {{ zone.client_port }}
              </p>
            </div>


            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
                <i class="fe fe-users"></i> {{ zone.number_players }}
              </p>
            </div>

            <div class="col-auto">
              <p
                class="card-text small text-muted mb-1 mt-2 fade-in"
                :style="'color: ' + getCpuUsageColor(zone.zone_os_pid) + ' !important'"
                v-if="getProcessStats(zone.zone_os_pid)"
              >
                <i class="fe fe-cpu"></i> {{ parseFloat(getProcessStats(zone.zone_os_pid).cpu).toFixed(2) }} %
              </p>
            </div>

            <div class="col-auto">
              <p
                :style="'color: ' + getMemUsageColor(zone.zone_os_pid) + ' !important'"
                class="card-text small text-muted mb-1 mt-2" v-if="getProcessStats(zone.zone_os_pid)">
                {{ parseFloat(getProcessStats(zone.zone_os_pid).memory / 1024 / 1024).toFixed(2) }}MB
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2" v-if="getProcessStats(zone.zone_os_pid)">
                <i class="fe fe-clock"></i>
                {{ parseFloat(getProcessStats(zone.zone_os_pid).elapsed / 1000 / 60 / 60).toFixed(2) }}h
              </p>
            </div>

            <div class="col-auto">
              <router-link
                class="text-muted"
                style="font-size:12px"
                :to="'zoneservers/' +zone.client_port + '/logs'"
              >
                <i class="fa fa-eye"></i> Logs
              </router-link>
            </div>

            <div class="col-auto">
              <a
                class="text-muted"
                href="javascript:void(0)"
                @click="killZone(zone.zone_os_pid)"
                style="font-size:12px"
              >
                <i class="fa fa-power-off"></i> Kill Zone
              </a>
            </div>

          </div>
        </div>
      </div>
    </div>


  </div>
</template>

<script>
import {ROUTE}         from "@/routes";
import {SpireApi}      from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";

export default {
  name: "ZoneServers",
  components: { InfoErrorBanner },
  data() {
    return {
      search: "",

      zoneList: [],
      processStats: {},
      loaded: false,
      zoneServerLoop: null,
      chartLoaded: false,

      // api responses
      error: "",
      notification: "",
    }
  },

  watch: {
    '$route'() {
      this.loadQueryState()
      this.init()
    },
  },

  async created() {
    this.loadQueryState()
    this.init()
  },
  methods: {

    getMemUsageColor(pid) {
      const p = parseFloat(this.getProcessStats(pid).memory / 1024 / 1024)
      if (p > 600) {
        return 'red';
      }
      else if (p > 300) {
        return 'orange';
      }

      return '#95aac9';
    },

    getCpuUsageColor(pid) {
      const p = parseFloat(this.getProcessStats(pid).cpu)
      if (p > 15) {
        return 'red';
      }
      else if (p > 5) {
        return 'orange';
      }

      return '#95aac9';
    },

    init() {
      this.getZoneList()

      if (!this.zoneServerLoop) {
        this.zoneServerLoop = setInterval(() => {
          if (!document.hidden) {
            this.getZoneList()
          }
        }, 1000)
      }
    },

    getProcessStats(pid) {
      return this.processStats.find((e) => {
        return e.pid === pid
      })
    },

    // state
    updateQueryState() {
      let q = {};
      if (this.search !== "") {
        q.search = this.search
      }

      this.$router.push(
        {
          path: ROUTE.ADMIN_ZONE_SERVERS,
          query: q
        }
      ).catch(() => {
      })
    },
    loadQueryState() {
      if (this.$route.query.search && this.$route.query.search.length > 0) {
        this.search = this.$route.query.search
      }
    },

    formatZoneName(name) {
      return name.replaceAll("UNKNOWN", "Idle")
    },

    async killZone(pid) {
      if (confirm("Are you sure that you want to kill this process pid (" + pid + ")?")) {
        // await OcculusClient.killProcessByPid(pid)

        this.zoneList = []

        try {
          const r = await SpireApi.v1().post(`admin/system/process-kill/${pid}`)
          if (r.status === 200) {
            this.notification = r.data.message
            this.error        = ""
          }
        } catch (e) {
          // error notify
          if (e.response && e.response.data && e.response.data.error) {
            this.error = e.response.data.error
          }
        }

        this.getZoneList()
      }
    },
    async getZoneList() {
      try {
        const r = await SpireApi.v1().get('eqemuserver/zone-list')
        if (r.status === 200) {
          this.zoneList     = r.data.zone_list.data
          this.processStats = r.data.process_info
          this.error        = ""
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.zoneList = []
          if (e.response.data.error.includes("Failed to connect to gameserver")) {
            this.error = ""
          } else {
            this.error = e.response.data.error
          }
        }
      }

      this.loaded = true
    },
    filterZoneList(list) {
      let zoneList = []
      for (let z of list) {
        if (
          z.zone_name.toLowerCase().includes(this.search) ||
          z.zone_long_name.toLowerCase().includes(this.search) ||
          z.client_port.toString().includes(this.search)
        ) {
          zoneList.push(z)
        }
      }

      return zoneList
    }
  },
  beforeDestroy() {
    clearInterval(this.zoneServerLoop)
  },
}
</script>

<style>
.avatar-image {
  background-image: url('~@/assets/img/eqemu-avatar.png');
}

.connector {
  width: 15px;
  height: 15px;
  opacity: .2;
  border-bottom-left-radius: 8px;
  border-left: 1px solid #1e375a;
  border-bottom: 1px solid #1e375a;
  margin-left: 9px;
  display: inline-flex;
}

.image-label {
  background-color: rgba(35, 100, 210, .07);
  color: #2364d2;
  padding: 0 4px;
  border-radius: 2px;
}
</style>
