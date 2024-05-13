<template>
  <div class="row" :style="loaded ? 'opacity:1' : 'opacity:.3'">
    <div class="col-12">
      <div class="card mb-3" v-if="clientList.length > 0 && filteredClientList.length > 0">
        <div class="card-body pt-0 pb-0 pl-3 pr-3">
          <div class="input-group input-group-flush">
            <div class="input-group-prepend">
              <span class="input-group-text"><i class="fe fe-search"></i></span>
            </div>
            <input
              class="list-search form-control"
              type="search"
              placeholder="Search for player(s) or zone..."
              @keyup="filterPlayers"
              v-model="playerSearchFilter"
            >
          </div>
        </div>

      </div>

      <div class="card">
        <div
          style="max-height: 78vh; overflow-y: scroll"
          class="table-responsive mb-0"
        >
          <div
            v-if="Object.keys(filteredClientList).length > listLimitSize && !fullList"
            class="m-3"
          >
            Too many online ({{ clientList.length }}) to display, for full list see
            <router-link class="ml-2" style="color: lightblue" :to="ROUTE.ADMIN_PLAYERS_ONLINE">
              <i class="fe fe-user"></i> Players
              Online
            </router-link>
          </div>

          <div
            v-if="(Object.keys(filteredClientList).length < listLimitSize && !fullList) || fullList"
            class="card-header"
          >
            <h4 class="card-header-title" v-if="loaded">Players Online ({{ clientList ? clientList.length : 0 }})</h4>
            <h4 class="card-header-title" v-if="!loaded">Loading...</span></h4>
          </div>

          <table
            v-if="clientList && clientList.length > 0 && loaded && Object.keys(filteredClientList).length > 0"
            class="table table-sm table-nowrap players-online sticky"
          >
            <thead class="sticky">
            <tr>
              <th style="width: 100px"></th>
              <th>Player</th>
              <th style="width: 75px">Level</th>
              <th style="min-width: 100px">Zone</th>
              <th style="width: 150px">Client</th>
              <th>IP</th>
            </tr>
            </thead>

            <tbody
              style="padding: 30px; overflow-y: scroll !important"
            >

            <tr
              v-for="(client, index) in filteredClientList.slice().reverse().slice(0, listLimitSize)"
              :key="client.name"
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
            <span v-if="client.server && client.server.zone_name">
              {{ client.server.zone_name }}
              <span class="badge badge-soft-primary">{{ client.server.zone_id }} ({{
                  client.server.instance_id
                }})</span>
            </span>
                <span v-if="!client.server && client.online === 1">
              Character Select
            </span>
                <span v-if="!client.server && client.online > 0">
              Zoning
            </span>
              </td>
              <td>
            <span v-if="client.client_version">
              {{ eqClientVersionConstants[client.client_version] }}
            </span>
              </td>
              <td>
            <span v-if="client.ip">
              {{ intToIP(client.ip) }}
            </span>
              </td>
            </tr>
            </tbody>
          </table>

        </div>

        <div class="card-body" v-if="clientList && clientList.length === 0 && loaded">
          There are currently no players online...
        </div>

      </div>
    </div>
  </div>

</template>

<script>
import Timer                       from "@/app/timer/timer";
import eqClassIntToStringConstants from "@/app/constants/eq-class-int-to-string-constants";
import {DB_CLASSES_ICONS}          from "@/app/constants/eq-class-icon-constants";
import eqClientVersionConstants    from "@/app/constants/eq-client-version-constants";
import {DB_RACES_ICONS}            from "@/app/constants/eq-race-icon-constants";
import {ROUTE}                     from "@/routes";
import {SpireApi}                  from "@/app/api/spire-api";

export default {
  name: 'PlayersOnlineComponent',
  computed: {
    ROUTE() {
      return ROUTE
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
      listLimitSize: null,
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

    filterPlayers() {
      this.filteredClientList = []

      if (!this.playerSearchFilter) {
        this.filteredClientList = this.clientList;
        return;
      }

      let clients = [];
      this.clientList.forEach(c => {
        const filter = this.playerSearchFilter.toLowerCase();

        if (
          c.name.toLowerCase().includes(filter) ||
          (c.server && c.server.zone_long_name.toLowerCase().includes(filter)) ||
          (c.server && c.server.zone_name.toLowerCase().includes(filter)) ||
          (c.ip && this.intToIP(c.ip).includes(filter))
        ) {
          clients.push(c);
        }
      })

      this.filteredClientList = clients;
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
      }
    }
  },

  /**
   * Destroy
   */
  beforeDestroy() {
    clearInterval(Timer.timer['players-online'])
  },

  /**
   * Create
   *
   * @returns {Promise<void>}
   */
  async created() {

    try {
      await this.buildPlayersOnlineList()
    } catch (e) {
    }

    this.loaded = true

    if (Timer.timer['players-online']) {
      clearInterval(Timer.timer['players-online'])
    }

    Timer.timer['players-online'] = setInterval(() => {
      if (!document.hidden) {
        this.buildPlayersOnlineList()
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
