<template>
  <eq-window
    :title="'Players Online (' + (clientList ? clientList.length : 0).toLocaleString() + ')'"
    class="p-0 mb-3"
  >
    <div
      v-if="tooManyOnlinet"
      class="m-3"
    >
      Too many online to display, for full list see
      <router-link class="ml-2" style="color: lightblue" :to="ROUTE.ADMIN_PLAYERS_ONLINE">
        <i class="fe fe-user"></i> Players
        Online
      </router-link>
    </div>

    <table
      v-if="clientList && clientList.length > 0 && loaded && Object.keys(filteredClientList).length > 0"
      class="eq-table bordered eq-highlight-rows mb-0"
    >
      <thead class="eq-table-floating-header">
      <tr>
        <th style="width: 100px"></th>
        <th>Player</th>
        <th style="width: 75px">Level</th>
        <th style="min-width: 100px">Zone</th>
        <th style="width: 150px">Client</th>
        <th v-if="fullList">IP</th>
      </tr>
      </thead>

      <tbody
        style="padding: 30px; overflow-y: scroll !important"
      >

      <tr
        v-for="(client, index) in filteredClientList.slice().reverse().slice(0, listLimitSize)"
        :key="client && client.name && client.name.length > 0 ? client.name : index"
      >
        <td style="text-align:center">
          <div class="avatar-list avatar-list-stacked">
            <img class="avatar-img rounded-circle" style="width:25px" :src="getClassImage(client.class)">
            <img class="avatar-img rounded-circle" style="width:25px" :src="getRaceImage(client.race)">
          </div>
        </td>
        <td style="align-content: center">{{ client.name }}</td>
        <td>{{ client.level }}</td>
        <td>
            <span v-if="client.server && client.server.zone_name">{{ client.server.zone_name }}
              <span class="badge badge-soft-primary">{{ client.server.zone_id }} ({{client.server.instance_id }})</span>
            </span>
          <span v-if="!client.server && client.online === 1">Character Select</span>
          <span v-if="!client.server && client.online > 0">Zoning</span>
        </td>
        <td>
          <span v-if="client.client_version">{{ eqClientVersionConstants[client.client_version] }}</span>
        </td>
        <td v-if="fullList">
          <span v-if="client.ip">{{ intToIP(client.ip) }}</span>
        </td>
      </tr>
      </tbody>
    </table>

    <div class="card-body" v-if="clientList && clientList.length === 0 && loaded">
      There are currently no players online...
    </div>

  </eq-window>
</template>

<script>
import Timer                       from "@/app/timer/timer";
import eqClassIntToStringConstants from "@/app/constants/eq-class-int-to-string-constants";
import {DB_CLASSES_ICONS}          from "@/app/constants/eq-class-icon-constants";
import eqClientVersionConstants    from "@/app/constants/eq-client-version-constants";
import {DB_RACES_ICONS}            from "@/app/constants/eq-race-icon-constants";
import {ROUTE}                     from "@/routes";
import {SpireApi}                  from "@/app/api/spire-api";
import EqWindow                    from "@/components/eq-ui/EQWindow.vue";

export default {
  name: 'PlayersOnlineComponent',
  components: { EqWindow },
  computed: {
    ROUTE() {
      return ROUTE
    },
    hasClients() {
      return Array.isArray(this.clientList) && this.clientList.length > 0;
    },
    hasFilteredClients() {
      return Array.isArray(this.filteredClientList) && this.filteredClientList.length > 0;
    }
  },

  props: {
    fullList: {
      default: false,
      required: false
    }
  },

  data() {
    return {
      lastUpdateTime: 0,
      updateIntervalSeconds: 1,
      tooManyOnline: false,

      listLimitSize: 10000,
      loaded: false,
      DB_CLASSES_ICONS: {},
      DB_RACES_ICONS: {},
      eqClientVersionConstants: {},
      clientList: [],
      filteredClientList: [],
      playersOnlineChart: null,
      playerSearchFilter: null
    }
  },

  mounted() {
    this.listLimitSize = (this.fullList ? 10000 : 50)
  },

  methods: {

    readyToPoll() {
      return Date.now() - this.lastUpdateTime > this.updateIntervalSeconds * 1000
    },

    filterPlayers() {
      this.filteredClientList = [];
      if (!this.playerSearchFilter) {
        this.filteredClientList = this.clientList || [];
        return;
      }

      const filter = this.playerSearchFilter.toLowerCase();
      this.clientList.forEach(client => {
        if (!client || !client.name) {
          return;
        }

        const nameMatch = client.name.toLowerCase().includes(filter);
        const zoneMatch = client.server && client.server.zone_name && client.server.zone_name.toLowerCase().includes(filter);
        const ipMatch   = client.ip && this.intToIP(client.ip).includes(filter);

        if (nameMatch || zoneMatch || ipMatch) {
          this.filteredClientList.push(client);
        }
      });
    },


    /**
     * @param num
     */
    intToIP: num => {
      var d = num % 256
      for (var i = 3; i > 0; i--) {
        num = Math.floor(num / 256)
        d   = num % 256 + '.' + d
      }

      var ip = d.split('.')

      return ip[3] + '.' + ip[2] + '.' + ip[1] + '.' + ip[0]
    },

    /**
     * @param classId
     * @returns {string|any}
     */
    getClassImage: classId => {
      if (DB_CLASSES_ICONS[classId]) {
        return require('@/assets/img/icons/classes-races/item_' + DB_CLASSES_ICONS[classId] + '.png')
      }

      return 'data:image/gif;base64,R0lGODlhAQABAIAAAMLCwgAAACH5BAAAAAAALAAAAAABAAEAAAICRAEAOw=='
    },

    /**
     * @param raceId
     * @returns {string|any}
     */
    getRaceImage: raceId => {
      if (DB_RACES_ICONS[raceId]) {
        return require('@/assets/img/icons/classes-races/item_' + DB_RACES_ICONS[raceId] + '.png')
      }

      return 'data:image/gif;base64,R0lGODlhAQABAIAAAMLCwgAAACH5BAAAAAAALAAAAAABAAEAAAICRAEAOw=='
    },

    /**
     * Builds online player list
     *
     * @returns {boolean}
     */
    async buildPlayersOnlineList() {

      try {
        const r = await SpireApi.v1().get("eqemuserver/client-list")
        if (r.status === 200) {
          const apiClientList = r.data.data
          if (!apiClientList) {
            this.filteredClientList = []
            this.clientList         = []
            return false
          }

          let clientList = []
          if (apiClientList.length > 0) {
            apiClientList.forEach(function (row) {
              if (row.character_id === 0) {
                return
              }
              clientList.push(row)
            })
          }

          if (clientList.length === 0) {
            this.clientList = []
            return false
          }

          this.clientList = clientList
          this.filterPlayers()

          const clientCount = this.clientList.length
          if (clientCount > 2000) {
            this.updateIntervalSeconds = 300
          } else if (clientCount > 1000) {
            this.updateIntervalSeconds = 60
          } else if (clientCount > 500) {
            this.updateIntervalSeconds = 10
          } else if (clientCount > 100) {
            this.updateIntervalSeconds = 5
          }

          this.lastUpdateTime = Date.now()
        }
      } catch (e) {
        this.clientList         = []
        this.filteredClientList = []
        console.log("Error fetching client list")
        return false
      }
    }
  },

  /**
   * Destroy
   */
  beforeDestroy() {
    if (Timer.timer['players-online']) {
      clearInterval(Timer.timer['players-online'])
    }
  },

  /**
   * Create
   *
   * @returns {Promise<void>}
   */
  async created() {

    await this.buildPlayersOnlineList()

    this.loaded = true

    if (Timer.timer['players-online']) {
      clearInterval(Timer.timer['players-online'])
    }

    Timer.timer['players-online'] = setInterval(() => {
      if (!document.hidden && this.readyToPoll()) {
        this.buildPlayersOnlineList()
      }

      if (Object.keys(this.filteredClientList).length > this.listLimitSize && !this.fullList) {
        clearInterval(Timer.timer['players-online'])
        this.filteredClientList = []
        this.clientList         = []
        this.tooManyOnline = true
      }
    }, 1000)

    /**
     * Classes Online breakdown
     */
    let classCounts = []
    if (this.clientList && this.clientList.length > 0) {
      this.clientList.forEach(function (row) {
        const classId = parseInt(row.class)
        if (typeof classCounts[classId] === 'undefined') {
          classCounts[classId] = 0
        }
        classCounts[classId]++
      })
    }

    let classCountsDataColumn  = []
    let classCountsColumnNames = {}
    for (let i = 1; i <= 16; i++) {
      const classCount = (typeof classCounts[i] !== 'undefined' ? classCounts[i] : 0)
      if (classCount > 0) {
        classCountsDataColumn.push([eqClassIntToStringConstants[i], classCount])
        classCountsColumnNames[eqClassIntToStringConstants[i]] = eqClassIntToStringConstants[i]
      }
    }

    this.DB_CLASSES_ICONS         = DB_CLASSES_ICONS
    this.DB_RACES_ICONS           = DB_RACES_ICONS
    this.eqClientVersionConstants = eqClientVersionConstants

  }
}

</script>

<style scoped>
.players-online td {
  border-radius: 10px;
  padding: 0.4rem;
  padding-top: 0.5rem;
  padding-right: 0.5rem;
  padding-bottom: 0.5rem;
  padding-left: 0.5rem;
}
</style>
