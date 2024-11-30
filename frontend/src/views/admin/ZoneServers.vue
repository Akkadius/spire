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
        :class="['card', 'mb-3', 'zone-card', 'zone-background-' + zone.zone_name]"
        v-for="zone in filterZoneList(zoneList)"
        :style="formatZoneRow(zone)"
        :key="zone.zone_id + '_' + zone.zone_os_pid + '_' + zone.zone_name + '_' + parseFloat(zone.cpu).toFixed(2) + '_' + zone.number_players"
      >
        <div
          :class="['card-body', 'btn-default', 'card-slim', 'zone-card-body']"
          :style="'box-shadow: 0 2px 4px 0 rgba(30,55,90,.1);'"
        >
          <div class="row align-items-center">

            <div class="col-2 ml-n1">
              <h4 class="mb-1">

                <span
                  :class="'h2 fe mb-0 ' + formatRowIcon(zone) + ' mr-3'"
                />

                <h3 v-if="zone.zone_long_name" class="d-inline-block header-title">
                  {{ formatZoneName(zone.zone_long_name ? zone.zone_long_name : 'Standby Zone') }}
                </h3>

                <span
                  style="color:#666;font-weight:200;font-size: 18px;"
                  class="text-muted"
                  v-if="!zone.zone_long_name"
                >
                Ready for players...
              </span>

              </h4>
            </div>

            <div class="col-10 text-right">
              <div class="ml-2 d-inline-block" style="min-width: 200px">
                <div class="ml-3 d-inline-block">
                  {{ formatZoneName(zone.zone_name) }}
                  ({{ zone.zone_id }})
                  <span v-if="zone.instance_id">Instance: {{ zone.instance_id }}</span>
                </div>

                <div class="d-inline-block ml-3">{{ zone.is_static_zone === true ? 'Static' : 'Dynamic' }}</div>
                <i class="fe fe-proc ml-3"></i> PID {{ zone.zone_os_pid }}
                <i class="fe fe-cloud ml-3"></i> IP {{ zone.zone_server_address }}:{{ zone.client_port }}
                <i class="fe fe-users ml-3"></i> {{ zone.number_players }}
              </div>

              <div
                class="text-left d-inline-block ml-3"
                :style="'color: ' + getCpuUsageColor(zone) + ' !important; min-width: 70px;'"
              >
                <i class="fe fe-cpu"></i> {{ zone.cpu ? parseFloat(zone.cpu).toFixed(2) : "N/A" }} %
              </div>

              <div
                :style="'color: ' + getMemUsageColor(zone) + ' !important; min-width: 100px'"
                class="ml-3 d-inline-block text-left"
              >
                <i class="fe fe-hard-drive"></i>
                {{ zone.memory ? parseFloat(zone.memory / 1024 / 1024).toFixed(2) + "MB" : "N/A" }}
              </div>

              <div class="ml-3 d-inline-block text-left" style="min-width: 70px">
                <i class="fe fe-clock"></i>
                {{ zone.elapsed > 0 ? parseFloat(zone.elapsed / 60 / 60).toFixed(2) + "h" : "N/A" }}
              </div>

              <div class="d-inline-block btn-group ml-4 text-right" role="group">
                <router-link
                  style="font-size:12px"
                  class="btn btn-sm btn-primary"
                  :to="'zoneservers/' +zone.client_port + '/logs'"
                >
                  <i class="fa fa-eye"></i> Logs
                </router-link>

                <button
                  class="btn btn-sm btn-danger"
                  @click="killZone(zone.zone_os_pid)"
                  style="font-size:12px"
                >
                  <i class="fa fa-power-off"></i> Kill Zone
                </button>
              </div>
            </div>

          </div>

          <div
            class="row pl-0 mt-3"
            v-if="zone.number_players > 0"
            style="border-top: 1px solid rgba(0, 0, 0, .2); padding-top: 10px; margin-top: 10px;"
          >
            <div class="col-auto">
              <div class="d-flex flex-wrap">
                <div
                  v-for="(c, index) in filterClients(zone.clients)"
                  :key="c.id"
                >

                  <div
                    class="avatar-group mr-3"
                    :title="formatPlayerTooltip(c)"
                  >
                    <img
                      class="avatar-img rounded-circle avatar-xs"
                      :src="getClassImage(c.class)"
                      style="border: 1px solid #666;"
                    >

                    <img
                      class="avatar-img rounded-circle avatar-xs"
                      :src="getRaceImage(c.race)"
                      style="margin-left: -.80625rem; border: 1px solid #666;"
                    >

                    <small
                      style="line-height: 24px; margin-left: 5px;"
                    >
                      {{ c.name }}
                      {{ c.guild ? '<' + c.guild + '>' : '' }}
                    </small>
                  </div>
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
import {ROUTE}             from "@/routes";
import {SpireApi}          from "@/app/api/spire-api";
import InfoErrorBanner     from "@/components/InfoErrorBanner.vue";
import {DB_RACES_ICONS}    from "@/app/constants/eq-race-icon-constants";
import util                from "util";
import {DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_CLASSES_ICONS}  from "@/app/constants/eq-class-icon-constants";

export default {
  name: "ZoneServers",
  components: { InfoErrorBanner },
  data() {
    return {
      search: "",

      lastUpdateTime: 0,
      updateIntervalSeconds: 1,

      zoneList: [],
      loaded: false,
      zoneServerLoop: null,
      chartLoaded: false,
      zoneBackgroundImages: {},

      guilds: {},

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
    readyToPoll() {
      return Date.now() - this.lastUpdateTime > this.updateIntervalSeconds * 1000
    },

    loadGuilds() {
      SpireApi.v1().get("guilds?limit=1000000").then((r) => {
        if (r.status === 200) {
          this.guilds = r.data
        }
      })
    },

    getRaceImage: raceId => {
      if (DB_RACES_ICONS[raceId]) {
        return require('@/assets/img/icons/classes-races/item_' + DB_RACES_ICONS[raceId] + '.png')
      }

      return 'data:image/gif;base64,R0lGODlhAQABAIAAAMLCwgAAACH5BAAAAAAALAAAAAABAAEAAAICRAEAOw=='
    },

    getClassImage: classId => {
      if (DB_CLASSES_ICONS[classId]) {
        return require('@/assets/img/icons/classes-races/item_' + DB_CLASSES_ICONS[classId] + '.png')
      }

      return 'data:image/gif;base64,R0lGODlhAQABAIAAAMLCwgAAACH5BAAAAAAALAAAAAABAAEAAAICRAEAOw=='
    },

    formatPlayerTooltip(c) {
      return util.format(
        "%s (%s)",
        c.name,
        DB_PLAYER_CLASSES[c.class] ? DB_PLAYER_CLASSES[c.class] : c.class
      )
    },

    filterClients(clients) {
      if (!clients) {
        return []
      }

      let list = []
      for (let c of clients) {
        let guild = "";

        if (c.guild_id) {
          let g = this.guilds.find((g) => {
            return g.id === c.guild_id
          })

          if (g) {
            guild = g.name
          }
        }

        list.push({
          id: c.id,
          name: c.name,
          race: c.race,
          class: c.class,
          guild: guild,
        })
      }

      return list
    },

    formatRowIcon(zone) {
      if (zone.zone_id === 0) {
        return "fe-loader"
      }

      if (zone.number_players === 0) {
        return "fe-loader"
      }

      return "fe-chevron-right"
    },

    formatZoneRow(zone) {
      if (zone.zone_id === 0) {
        return {
          backgroundColor: 'rgba(0,0,0,.8)',
          color: 'rgba(255, 255, 255, .7) !important',
          border: '1px solid #e9ecef',
          borderRadius: '5px',
          marginBottom: '10px'
        }
      }

      if (zone.number_players === 0) {
        return {
          backgroundColor: 'rgba(0,0,0,1)',
          color: 'rgba(255, 255, 255, .9) !important',
          border: '1px solid white',
          // inner shadow white
          // boxShadow: 'inset 0 0 13px rgba(255,255,255,.1)',
          borderRadius: '5px',
          marginBottom: '10px'
        }
      }

      return {
        backgroundColor: '#f8f9fa',
        border: '1px solid #e9ecef',
        color: 'rgba(18,38,63,1) !important',
        borderRadius: '5px',
        marginBottom: '10px'
      }
    },

    getMemUsageColor(zone) {
      const p = parseFloat(zone.memory / 1024 / 1024)
      if (p > 600) {
        return 'red';
      } else if (p > 300) {
        return 'orange';
      }

      return '';
    },

    getCpuUsageColor(zone) {
      const p = parseFloat(zone.cpu)
      if (p > 15) {
        return 'red';
      } else if (p > 5) {
        return 'orange';
      }

      return '';
    },

    init() {
      this.getZoneList()
      this.loadGuilds()

      if (!this.zoneServerLoop) {
        this.zoneServerLoop = setInterval(() => {
          if (!document.hidden && this.readyToPoll()) {
            this.getZoneList()
          }
        }, 1000)
      }
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
      return name.replaceAll("UNKNOWN", "Idle, ready to boot")
    },

    async killZone(pid) {
      if (confirm("Are you sure that you want to kill this process pid (" + pid + ")?")) {
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
        const r = await SpireApi.v1().get('eqemuserver/zoneserver-list')
        if (r.status === 200) {

          // sort list
          // sort by number of players first
          // then zones with id of 0 always on bottom of list
          // then sort by zone name
          this.zoneList = r.data.sort(
            (a, b) => {
              if (a.number_players > b.number_players) {
                return -1
              }

              if (a.number_players < b.number_players) {
                return 1
              }

              if (a.zone_id === 0) {
                return 1
              }

              if (b.zone_id === 0) {
                return -1
              }

              if (a.zone_name < b.zone_name) {
                return -1
              }

              if (a.zone_name > b.zone_name) {
                return 1
              }

              return 0
            }
          )

          if (typeof this.zoneList === "undefined") {
            this.zoneList = []
          }

          this.error = ""
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          if (e.response.data.error.includes("Failed to connect to gameserver")) {
            this.error = ""
          } else {
            this.error = e.response.data.error
          }
        }
      }

      this.lastUpdateTime = Date.now()

      if (this.zoneList.length > 1000) {
        this.updateIntervalSeconds = 30
      } if (this.zoneList.length > 500) {
        this.updateIntervalSeconds = 10
      } else if (this.zoneList.length > 100) {
        this.updateIntervalSeconds = 5
      }

      this.loaded = true
    },

    filterZoneList(list) {
      const search = this.search.toLowerCase()

      let zoneList = []
      for (let z of list) {
        let clients       = this.filterClients(z.id)
        let matchesClient = false
        for (let c of clients) {
          if (c.name.toLowerCase().includes(search)) {
            matchesClient = true
            break
          }
        }

        if (
          z.zone_name.toLowerCase().includes(search) ||
          z.zone_long_name.toLowerCase().includes(search) ||
          z.client_port.toString().includes(search) ||
          matchesClient
        ) {
          zoneList.push(z)
        }
      }

      return zoneList
    }
  },
  beforeDestroy() {
    clearInterval(this.zoneServerLoop)
  }
}
</script>

<style>
.zone-card-body {
  padding-top: 10px !important;
  padding-bottom: 10px !important;
  padding-left: 18px !important;
  padding-right: 18px !important;
}
</style>
