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
        :class="['card', 'mb-3', 'fade-in', 'zone-card', 'zone-background-' + zone.zone_name]"
        v-if="getProcessStats(zone.zone_os_pid)"
        v-for="zone in filterZoneList(zoneList)"
        :style="formatZoneRow(zone)"
      >
        <div
          :class="['card-body', 'btn-default', 'card-slim', 'zone-card-body']"
          :style="'box-shadow: 0 2px 4px 0 rgba(30,55,90,.1);'"
        >
          <div class="row align-items-center">

            <div class="col ml-n1">
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

            <div class="col-auto" v-if="zone.zone_name">
              <p class="mb-1 mt-2">
              <span class="ml-2">
                  {{ formatZoneName(zone.zone_name) }}
                  ({{ zone.zone_id }})
                  <span v-if="zone.instance_id">Instance: {{ zone.instance_id }}</span>
                </span>
              </p>
            </div>

            <div class="col-auto">
              <p class="mb-1 mt-2">
              <span class="ml-2">
                  {{ zone.is_static_zone === true ? 'Static' : 'Dynamic' }}
              </span>
              </p>
            </div>

            <div class="col-auto">
              <p class="mb-1 mt-2">
                <i class="fe fe-proc"></i> PID {{ zone.zone_os_pid }}
              </p>
            </div>

            <div class="col-auto">
              <p class="mb-1 mt-2">
                <i class="fe fe-cloud"></i> Port {{ zone.client_port }}
              </p>
            </div>


            <div class="col-auto">
              <p class="mb-1 mt-2">
                <i class="fe fe-users"></i> {{ zone.number_players }}
              </p>
            </div>

            <div class="col-auto">
              <p
                class="mb-1 mt-2 fade-in"
                :style="'color: ' + getCpuUsageColor(zone.zone_os_pid) + ' !important'"
                v-if="getProcessStats(zone.zone_os_pid)"
              >
                <i class="fe fe-cpu"></i> {{ parseFloat(getProcessStats(zone.zone_os_pid).cpu).toFixed(2) }} %
              </p>
            </div>

            <div class="col-auto">
              <p
                :style="'color: ' + getMemUsageColor(zone.zone_os_pid) + ' !important'"
                class="mb-1 mt-2" v-if="getProcessStats(zone.zone_os_pid)"
              >
                {{ parseFloat(getProcessStats(zone.zone_os_pid).memory / 1024 / 1024).toFixed(2) }}MB
              </p>
            </div>

            <div class="col-auto">
              <p class="mb-1 mt-2" v-if="getProcessStats(zone.zone_os_pid)">
                <i class="fe fe-clock"></i>
                {{ parseFloat(getProcessStats(zone.zone_os_pid).elapsed / 1000 / 60 / 60).toFixed(2) }}h
              </p>
            </div>

<!--            <div class="col-auto">-->
<!--              <router-link-->
<!--                style="font-size:12px"-->
<!--                :to="'zoneservers/' +zone.client_port + '/logs'"-->
<!--              >-->
<!--                <i class="fa fa-eye"></i> Logs-->
<!--              </router-link>-->
<!--            </div>-->

            <div class="col-auto">
              <button
                class="btn btn-sm btn-danger"
                @click="killZone(zone.zone_os_pid)"
                style="font-size:12px"
              >
                <i class="fa fa-power-off"></i> Kill Zone
              </button>
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
                  class="fade-in"
                  v-for="(c, index) in getClientByInternalZoneId(zone.id)"
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
                      {{ c.guild ? '<' + c.guild + '>' : ''}}
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

      zoneList: [],
      processStats: {},
      loaded: false,
      zoneServerLoop: null,
      clientPollingLoop: null,
      chartLoaded: false,
      zoneBackgroundImages: {},

      clients: [],
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

    getClientByInternalZoneId(id) {
      if (!this.clients) {
        return []
      }

      let clients = this.clients.filter((c) => {
        return c.server && c.server.id === id
      })

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

        // let id = 100;
        // for (let x = 0; x < 22; x++) {
        //   for (let i = 0; i < 11; i++) {
        //     list.push({
        //       id: id,
        //       name: Math.random().toString(36),
        //       race: c.race + i,
        //       class: c.class + i,
        //       guild: guild,
        //     })
        //     id++
        //   }
        // }

      }

      return list
    },

    async getPlayers() {
      const r = await SpireApi.v1().get("eqemuserver/client-list")
      if (r.status === 200) {
        this.clients = r.data.data
      }
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

    getMemUsageColor(pid) {
      const p = parseFloat(this.getProcessStats(pid).memory / 1024 / 1024)
      if (p > 600) {
        return 'red';
      } else if (p > 300) {
        return 'orange';
      }

      return '';
    },

    getCpuUsageColor(pid) {
      const p = parseFloat(this.getProcessStats(pid).cpu)
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
          if (!document.hidden) {
            this.getZoneList()
          }
        }, 1000)


        this.getPlayers()
        this.clientPollingLoop = setInterval(() => {
          if (!document.hidden) {
            this.getPlayers()
          }
        }, 5000)

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
      return name.replaceAll("UNKNOWN", "Idle, ready to boot")
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

          // sort list
          // sort by number of players first
          // then zones with id of 0 always on bottom of list
          // then sort by zone name
          this.zoneList = r.data.zone_list.data.sort(
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
      const search = this.search.toLowerCase()

      let zoneList = []
      for (let z of list) {
        let clients = this.getClientByInternalZoneId(z.id)
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
    clearInterval(this.clientPollingLoop)
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
