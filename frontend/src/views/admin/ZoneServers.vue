<template>
  <div :style="loaded ? 'opacity:1' : 'opacity:.3'">

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

    <eq-window
      title="Zone Servers"
      class="mb-3 pb-3"
      v-if="zoneList.length > 0 && loaded"
      style=""
    >
      <div>
        <div class="row">
          <div class="col-10">
            <b-form-input
              type="text"
              class="form-control list-search mt-1"
              @keyup="debouncedUpdateQueryState()"
              v-model="search"
              placeholder="Search zone servers by zone name or player name..."
            />
          </div>

          <div class="text-center col-1 font-weight-bold">
            Show Players
            <eq-checkbox
              class="mt-1"
              :fade-when-not-true="true"
              :true-value="true"
              :false-value="false"
              v-model="showPlayers"
              @input="updateQueryState()"
            />
          </div>

          <div class="col-1 mt-1">
            <button
              title="Reset"
              class="eq-button m-0"
              @click="search = ''; updateQueryState()"
            ><i class="fa fa-refresh"></i> Reset
            </button>
          </div>

        </div>
      </div>
    </eq-window>

    <eq-window
      class="text-center"
      title="Zoneserver Stats"
      v-if="zoneList.length > 0"
    >

      <span class="font-weight-bold">• Zones</span> {{ zoneList.length }}
      <span class="font-weight-bold">• Statics</span> {{ zoneList.filter(z => z.is_static_zone === true).length }}
      <span class="font-weight-bold">• Dynamics</span> {{ zoneList.filter(z => z.is_static_zone === false).length }}
      <span class="font-weight-bold">• Sleeping</span> {{ zoneList.filter(z => z.number_players === 0).length }}
      <span class="font-weight-bold">• Highest Zone CPU</span>
      {{ zoneList.length > 0 ? parseFloat(Math.max.apply(Math, zoneList.map(z => z.cpu))).toFixed(2) : 0 }}%
      <span class="font-weight-bold">• Highest Zone Memory</span> {{
        zoneList.length > 0 ? parseFloat(Math.max.apply(Math, zoneList.map(z => z.memory / 1024 / 1024))).toFixed(2) : 0
      }}MB
      <span class="font-weight-bold">• Highest Zone Uptime</span> {{
        zoneList.length > 0 ? parseFloat(Math.max.apply(Math, zoneList.map(z => z.elapsed / 60 / 60))).toFixed(2) : 0
      }}h

      <div v-if="zoneList.length > 1" class="mt-2">

      <span v-for="(count, ip) in zoneCountByIP" :key="ip" class="mt-2">
        <span class="font-weight-bold">• {{ ip }}</span> - {{ count }} Zones
      </span>
      </div>

    </eq-window>

    <eq-window
      class="p-0"
    >
      <app-loader
        style="opacity: .3; top: 300px; left: 20%; position: absolute;"
        :is-loading="!loaded"
      />

      <div
        style="max-height: 60vh; overflow-y: scroll; overflow-x: hidden; border: 1px solid #ffffff1c !important"
        class="p-0"
      >

        <div v-if="filterZoneList(zoneList).length === 0" class="p-3 text-center">
          No zones found with specified criteria or server is offline.
        </div>

        <table
          :style="(!loaded ? 'opacity:.3; pointer-events: none;' : 'opacity: 1; pointer-events: all;') + 'table-layout: fixed !important; width: 100% '"
          class="eq-table bordered eq-highlight-rows"
          v-if="filterZoneList(zoneList).length > 0"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th style="width: 100px; text-align: center">Tools</th>
            <th>Zone</th>
            <th style="width: 70px; text-align: center"><i class="fe fe-users"></i></th>
            <th style="width: 100px; text-align: center"><i class="fe fe-cpu"></i> PID</th>
            <th style="width: 100px">Zone Type</th>
            <th style="width: 85px">Zone ID</th>
            <th>Instance ID</th>
            <th style="text-align: center"><i class="fe fe-cloud"></i> IP</th>
            <th style="text-align: center"><i class="fe fe-cpu"></i> CPU</th>
            <th style="text-align: center"><i class="fe fe-hard-drive"></i> Memory</th>
            <th style="text-align: center"><i class="fe fe-clock"></i> Uptime</th>
          </tr>
          </thead>
          <tbody
            v-for="zone in filterZoneList(zoneList)"
            :key="zone.zone_id + '_' + zone.zone_os_pid + '_' + zone.zone_name + '_' + parseFloat(zone.cpu).toFixed(2) + '_' + zone.number_players + '_' + showPlayersForZone[zone.id]"
          >
          <tr
            :style="formatZoneRow(zone)"
          >
            <td style="text-align: center">
              <router-link
                style="font-size:12px; color: white"
                class="btn btn-sm btn-primary mr-1"
                :to="'zoneservers/' +zone.client_port + '/logs'"
                title="Logs"
              >
                <i class="fa fa-eye"></i>
              </router-link>

              <button
                class="btn btn-sm btn-danger"
                @click="killZone(zone)"
                style="font-size:12px"
                title="Kill Zone"
              >
                <i class="fa fa-power-off"></i>
              </button>
            </td>
            <td>{{ formatZoneName(zone.zone_name) }}</td>
            <td style="text-align: center">
              <a
                href="javascript:;"
                @click="toggleShowPlayers(zone.id)"
                v-if="zone.number_players > 0"
              >
                {{ (zone.number_players).toLocaleString('en-US') }}
              </a>
              <span v-if="zone.number_players === 0">
                    {{ (zone.number_players).toLocaleString('en-US') }}
                  </span>
            </td>
            <td style="text-align: center">{{ zone.zone_os_pid }}</td>
            <td>{{ zone.is_static_zone === true ? 'Static' : 'Dynamic' }}</td>
            <td>{{ zone.zone_id }}</td>
            <td>{{ zone.instance_id ? zone.instance_id : 0 }}</td>
            <td style="text-align: center">{{ zone.zone_server_address }}:{{ zone.client_port }}</td>
            <td :style="'min-width: 70px; text-align: center'">
              <span :style="'color: ' + getCpuUsageColor(zone) + ' !important'">
                {{ zone.cpu ? parseFloat(zone.cpu).toFixed(0) : "N/A" }} %
              </span>
              <eq-progress-bar
                :percent="parseFloat(zone.cpu)"
                :show-percent="false"
                :color="getCpuLoadColor(parseFloat(zone.cpu))"
              />
            </td>
            <td :style="'color: ' + getMemUsageColor(zone) + ' !important; min-width: 100px; text-align: center'">
              {{ zone.memory ? parseFloat(zone.memory / 1024 / 1024).toFixed(2) + "MB" : "N/A" }}
            </td>
            <td style="text-align: center">
              {{ zone.elapsed > 0 ? parseFloat(zone.elapsed / 60 / 60).toFixed(2) + "h" : "N/A" }}
            </td>
          </tr>

          <tr
            class="fade-in"
            :style="formatZoneRow(zone)"
            v-if="(showPlayers || showPlayersForZone[zone.id]) && zone.number_players > 0"
          >
            <td
              colspan="11"
              style="box-shadow: inset 3px 3px 6px rgba(0, 0, 0, 0.2), inset -3px -3px 30px rgba(0,0,0,0.7), 3px 3px 6px rgba(0, 0, 0, 0.1);"
            >
              <div class="d-flex flex-wrap p-3">
                <div
                  v-for="(c, index) in filterClients(zone.clients)"
                  :key="c.id"
                >
                  <div
                    class="avatar-group"
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
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </eq-window>

  </div>
</template>

<script>
import {ROUTE}                     from "@/routes";
import {SpireApi}                  from "@/app/api/spire-api";
import InfoErrorBanner             from "@/components/InfoErrorBanner.vue";
import {DB_RACES_ICONS}            from "@/app/constants/eq-race-icon-constants";
import util                        from "util";
import {DB_PLAYER_CLASSES}         from "@/app/constants/eq-classes-constants";
import {DB_CLASSES_ICONS}          from "@/app/constants/eq-class-icon-constants";
import EqWindow                    from "@/components/eq-ui/EQWindow.vue";
import EqCheckbox                  from "@/components/eq-ui/EQCheckbox.vue";
import PlayerEventDisplayComponent from "@/views/admin/player-event-logs/components/PlayerEventDisplayComponent.vue";
import EqProgressBar               from "@/components/eq-ui/EQProgressBar.vue";

// Simple debounce function
function debounce(func, delay) {
  let timer;
  return function (...args) {
    clearTimeout(timer);
    timer = setTimeout(() => func.apply(this, args), delay);
  };
}

export default {
  name: "ZoneServers",
  components: { EqProgressBar, PlayerEventDisplayComponent, EqCheckbox, EqWindow, InfoErrorBanner },
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

      showPlayers: false,
      showPlayersForZone: {},

      highlightedZone: {},

      guilds: {},

      // api responses
      error: "",
      notification: "",
    }
  },

  computed: {
    zoneCountByIP() {
      const count = {};

      this.zoneList.forEach(zone => {
        const ip = zone.zone_server_address;

        // Initialize count for the IP if it doesn't exist
        if (!count[ip]) {
          count[ip] = 0;
        }

        // Increment the count for the current IP
        count[ip]++;
      });

      return count; // Return the object with IP -> count mapping
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
    getCpuLoadColor(load) {
      if (load > 80) {
        return 'red'
      }
      if (load > 50) {
        return 'orange'
      }

      return '#2c7be5'
    },

    debouncedUpdateQueryState: debounce(function () {
      this.updateQueryState();
    }, 300),

    toggleShowPlayers(zoneId) {
      this.$set(this.showPlayersForZone, zoneId, !this.showPlayersForZone[zoneId]);
    },

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
          opacity: '.5',
        }
      }

      if (zone.number_players === 0) {
        return {
          color: 'rgba(255, 255, 255, .9) !important',
          opacity: '.5',
        }
      }

      return {
        color: 'rgba(18,38,63,1) !important',
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
      if (typeof this.showPlayers !== "undefined") {
        q.showPlayers = this.showPlayers
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

      if (this.$route.query.showPlayers) {
        this.showPlayers = this.$route.query.showPlayers === "true"
      }
    },

    formatZoneName(name) {
      return name.replaceAll("UNKNOWN", "Idle, ready to boot")
    },

    async killZone(zone) {
      if (confirm("Are you sure that you want to kill this process pid (" + zone.zone_os_pid + ")?")) {
        this.zoneList = []

        try {
          const r = await SpireApi.v1().post(`eqemuserver/server/process-kill/${zone.zone_os_pid}`, zone)
          if (r.status === 200) {
            this.notification = r.data.message
            this.error        = ""
            setTimeout(() => {
              this.getZoneList()
            }, 100)
            setTimeout(() => {
              this.getZoneList()
            }, 1000)
          }
        } catch (e) {
          // error notify
          if (e.response && e.response.data && e.response.data.error) {
            this.error = e.response.data.error
          }
        }
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
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error

          if (this.error.includes("Failed to") || this.error.includes("connection refused")) {
            this.error    = ""
            this.zoneList = []
          }
        }
      }

      this.lastUpdateTime = Date.now()

      if (this.zoneList.length > 1000) {
        this.updateIntervalSeconds = 30
      }
      if (this.zoneList.length > 500) {
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
        let clients       = this.filterClients(z.clients)
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
