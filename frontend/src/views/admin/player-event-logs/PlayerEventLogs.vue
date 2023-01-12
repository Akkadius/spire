<template>
  <div>

    <app-loader :is-loading="initialLoading"/>

    <eq-window
      v-if="!initialLoading"
      class="pb-1"
      title="Player Event Log Explorer"
    >
      <div v-if="!initialLoading">
        <div class="row mb-3">
          <div class="col-2 text-center font-weight-bold">
            Event Type
            <select
              class="form-control form-control-prepended list-search"
              v-model="eventType"
              @change="updateQueryState()"
            >
              <option value="0">-- Select Event Type --</option>
              <option v-for="s in settings" v-bind:value="s.id">
                {{ s.event_name }} ({{ s.id }})
              </option>

            </select>
          </div>

          <div class="col-2 text-center font-weight-bold">
            Zone ID
            <b-form-input
              type="text"
              class="form-control list-search"
              @keyup="updateQueryState()"
              v-model="zoneId"
              placeholder="Zone ID"
            />
          </div>

          <div class="col-2 text-center font-weight-bold">
            Character ID
            <b-form-input
              type="text"
              class="form-control list-search"
              @keyup="updateQueryState()"
              v-model="characterId"
              placeholder="Character ID"
            />
          </div>

          <div class="col-1 text-center font-weight-bold">
            Auto Refresh
            <select
              class="form-control form-control-prepended list-search"
              v-model="refreshInterval"
              @change="updateQueryState()"
            >
              <option v-for="s in timerOptions" v-bind:value="s.timer">
                {{ secondsToHumanTime(s.timer / 1000) }}
              </option>

            </select>
          </div>

          <div class="col-1">
            <button
              title="Reset"
              class="eq-button m-0"
              style="margin-top: 20px !important"
              @click="reset(); updateQueryState()"
            ><i class="fa fa-refresh"></i> Reset
            </button>
          </div>
        </div>
      </div>
    </eq-window>

    <eq-window
      v-if="!initialLoading"
      class="p-1"
    >
      <app-loader
        style="opacity: .3; top: 300px; left: 20%; position: absolute;"
        :is-loading="loading"
      />

      <div v-if="events.length === 0" class="font-weight-bold text-center p-3">
        No events found with specified criteria.
      </div>

      <div
        style="height: 69vh; overflow-y: scroll; overflow-x: hidden; border: 1px solid #ffffff1c !important"
      >
        <table
          :style="(loading ? 'opacity:.3; pointer-events: none;' : 'opacity: 1; pointer-events: all;') + 'table-layout: fixed !important; '"
          class="eq-table eq-highlight-rows bordered player-events"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th style="width: 150px">Event ID</th>
            <th style="width: 150px" class="text-center">Player</th>
            <th style="width: 200px">Zone</th>

            <th class="text-center" style="width: 175px">Event Type</th>
            <th class="text-center">Event</th>
            <th style="width: 140px">Time</th>
          </tr>
          </thead>
          <tbody>
          <tr
            class="fade-in"
            v-for="e in events"
            :key="e.id"
          >
            <td>{{ commify(e.id) }}</td>
            <td>

              <div class="avatar-list avatar-list-stacked">
                <img
                  class="avatar-img rounded-circle"
                  style="width:20px; border: 1px solid gray"
                  :src="getClassImage(e.character_datum.class)"
                >

                <img
                  class="avatar-img rounded-circle"
                  style="width:20px; border: 1px solid gray"
                  :src="getRaceImage(e.character_datum.race)"
                >

                <a
                  class="ml-1"
                  @click="characterId = e.character_datum.id; updateQueryState()"
                >
                  {{ e.character_datum.name }}
                </a>

              </div>
            </td>
            <td>
              <a
                class="ml-1"
                @click="zoneId = e.zone.zoneidnumber; updateQueryState()"
              >
                {{ e.zone.long_name }}
              </a>({{ e.zone.zoneidnumber }})
            </td>

            <td class="text-center">
              <a
                class="ml-1"
                @click="eventType = e.event_type_id; updateQueryState()"
              >{{ e.event_type_name }}</a> ({{ e.event_type_id }})
            </td>

            <td
              style="vertical-align: middle; text-align: left;"
            >
              <button
                title="View raw"
                @click="showRawEvent(e)"
                v-if="!showRaw[e.id] && Object.keys(JSON.parse(e.event_data)).length > 0"
                class="mr-3 btn btn-sm btn-warning"
                style="font-size: 10px; min-width: 50px"
              >
                <i class="fa fa-search"></i> ({{ Object.keys(JSON.parse(e.event_data)).length }})
              </button>

              <player-event-display-component
                :e="e"
              />

              <pre
                v-if="showRaw[e.id]"
                class="text-left code fade-in mt-2"
                style="width: 100%; padding: 0 !important; margin-bottom: 0 !important"
              ><code class="language-json">{{ e.event_data.replaceAll("    ", "  ") }}</code></pre>

            </td>
            <td>{{ fromNow(e.created_at) }}</td>
          </tr>
          </tbody>
        </table>
      </div>

      <div class="text-center">

        <div class="mr-3 d-inline-block">
          Pages ({{ commify(Math.round(totalRows / pageLimit)) }})
        </div>

        <b-pagination
          :disabled="loading"
          class="mb-1 mt-1"
          v-model="currentPage"
          :total-rows="totalRows"
          :hide-ellipsis="true"
          :per-page="pageLimit"
          @change="paginate"
        />

        <div class="ml-3 d-inline-block">
          Rows ({{ commify(events.length) }})
          Total ({{ commify(totalRows) }})
        </div>
      </div>

    </eq-window>
  </div>
</template>

<script>
import EqWindow                    from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}                  from "@/app/api/spire-api";
import {PlayerEventLogApi}         from "@/app/api/api/player-event-log-api";
import {SpireQueryBuilder}         from "@/app/api/spire-query-builder";
import moment                      from "moment/moment";
import {DB_CLASSES_ICONS}          from "@/app/constants/eq-class-icon-constants";
import {DB_RACES_ICONS}            from "@/app/constants/eq-race-icon-constants";
import PlayerEventDisplayComponent from "@/views/admin/player-event-logs/components/PlayerEventDisplayComponent.vue";
import {AA}                        from "@/app/aa";
import {Zones}                     from "@/app/zones";
import {Npcs}                      from "@/app/npcs";
import {Items}                     from "@/app/items";
import {ROUTE}                     from "@/routes";
import {PlayerEventLogSettingApi}  from "@/app/api/api/player-event-log-setting-api";
import Timer                       from "@/app/timer/timer";
import EqProgressBar               from "@/components/eq-ui/EQProgressBar.vue";
import LoaderFakeProgress          from "@/components/LoaderFakeProgress.vue";
import hljs                        from "highlight.js";
import {Navbar}                    from "@/app/navbar";

// GM_COMMAND           | [x] Implemented Formatter
// ZONING               | [x] Implemented Formatter
// AA_GAIN              | [x] Implemented Formatter
// AA_PURCHASE          | [x] Implemented Formatter
// FORAGE_SUCCESS       | [x] Implemented Formatter
// FORAGE_FAILURE       | [x] Implemented Formatter
// FISH_SUCCESS         | [x] Implemented Formatter
// FISH_FAILURE         | [x] Implemented Formatter
// ITEM_DESTROY         | [x] Implemented Formatter
// WENT_ONLINE          | [x] Implemented Formatter
// WENT_OFFLINE         | [x] Implemented Formatter
// LEVEL_GAIN           | [x] Implemented Formatter
// LEVEL_LOSS           | [x] Implemented Formatter
// LOOT_ITEM            | [x] Implemented Formatter
// MERCHANT_PURCHASE    | [x] Implemented Formatter
// MERCHANT_SELL        | [x] Implemented Formatter
// GROUP_JOIN           | [] Implemented Formatter - not implemented in server
// GROUP_LEAVE          | [] Implemented Formatter - not implemented in server
// RAID_JOIN            | [] Implemented Formatter - not implemented in server
// RAID_LEAVE           | [] Implemented Formatter - not implemented in server
// GROUNDSPAWN_PICKUP   | [x] Implemented Formatter
// NPC_HANDIN           | [] Implemented Formatter
// SKILL_UP             | [x] Implemented Formatter
// TASK_ACCEPT          | [x] Implemented Formatter
// TASK_UPDATE          | [] Implemented Formatter
// TASK_COMPLETE        | [] Implemented Formatter
// TRADE                | [] Implemented Formatter
// GIVE_ITEM            | [] Implemented Formatter
// SAY                  | [x] Implemented Formatter
// REZ_ACCEPTED         | [x] Implemented Formatter
// DEATH                | [x] Implemented Formatter
// COMBINE_FAILURE      | [x] Implemented Formatter
// COMBINE_SUCCESS      | [x] Implemented Formatter
// DROPPED_ITEM         | [x] Implemented Formatter
// SPLIT_MONEY          | [] Implemented Formatter
// DZ_JOIN              | [] Implemented Formatter - not implemented in server
// DZ_LEAVE             | [] Implemented Formatter - not implemented in server
// TRADER_PURCHASE      | [] Implemented Formatter
// TRADER_SELL          | [] Implemented Formatter
// BANDOLIER_CREATE     | [] Implemented Formatter - not implemented in server
// BANDOLIER_SWAP       | [] Implemented Formatter - not implemented in server
// DISCOVER_ITEM        | [x] Implemented Formatter

export default {
  name: "PlayerEventLogs",
  components: { LoaderFakeProgress, EqProgressBar, PlayerEventDisplayComponent, EqWindow },
  data() {
    return {
      search: "",
      eventType: 0,
      zoneId: null,
      characterId: null,

      refreshInterval: 5000,

      showRaw: {},

      timerOptions: [
        { timer: 1000 },
        { timer: 5000 },
        { timer: 10000 },
        { timer: 60000 },
      ],

      loading: false,
      initialLoading: false,

      settings: [],

      events: [],

      // pagination (all)
      currentPage: 1,
      totalRows: 0,
      pageLimit: 100,
    }
  },
  watch: {
    $route(to, from) {
      this.reset()
      this.loadQueryState()
      this.stopTimer()
      this.startTimer()

      this.loading = true
      setTimeout(() => {
        this.loadEvents();
      }, 1)
    }
  },
  methods: {

    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },

    paginate() {
      // models aren't quite updated when we trigger this so queue the pagination
      setTimeout(() => {
        console.log("We're paginating")
        console.log(this.currentPage)
        console.log(this.totalRows)
        this.updateQueryState()
      }, 100)
    },

    showRawEvent(e) {
      this.showRaw[e.id] = 1
      this.$forceUpdate()

      setTimeout(() => {
        // hljs.initHighlighting()
        for (let b of document.querySelectorAll('pre code')) {
          hljs.highlightBlock(b)
        }
      }, 10)
    },

    reset() {
      this.search      = "";
      this.eventType   = 0;
      this.zoneId      = null
      this.characterId = null
      this.showRaw     = {}
    },

    updateQueryState() {
      let q = {};

      if (this.search !== "") {
        q.search = this.search
      }
      if (parseInt(this.eventType) !== 0) {
        q.eventType = parseInt(this.eventType)
      }
      if (this.zoneId && parseInt(this.zoneId) !== 0) {
        q.zoneId = parseInt(this.zoneId)
      }
      if (this.characterId && parseInt(this.characterId) !== 0) {
        q.characterId = parseInt(this.characterId)
      }
      if (this.currentPage > 0) {
        q.page = this.currentPage
      }

      this.$router.push(
        {
          path: ROUTE.ADMIN_TOOL_PLAYER_EVENT_LOGS,
          query: q
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (this.$route.query.search && this.$route.query.search.length > 0) {
        this.search = this.$route.query.search
      }
      if (this.$route.query.eventType && parseInt(this.$route.query.eventType) > 0) {
        this.eventType = parseInt(this.$route.query.eventType)
      }
      if (this.$route.query.zoneId && parseInt(this.$route.query.zoneId) > 0) {
        this.zoneId = parseInt(this.$route.query.zoneId)
      }
      if (this.$route.query.characterId && parseInt(this.$route.query.characterId) > 0) {
        this.characterId = parseInt(this.$route.query.characterId)
      }
      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
      }
    },

    secondsToHumanTime(seconds) {
      let levels     = [
        [Math.floor(seconds / 31536000), 'y'],
        [Math.floor((seconds % 31536000) / 86400), 'd'],
        [Math.floor(((seconds % 31536000) % 86400) / 3600), 'h'],
        [Math.floor((((seconds % 31536000) % 86400) % 3600) / 60), 'm'],
        [(((seconds % 31536000) % 86400) % 3600) % 60, 's'],
      ];
      let returntext = '';
      for (let i = 0, max = levels.length; i < max; i++) {
        if (levels[i][0] === 0) continue;
        returntext += ' ' + levels[i][0] + '' + (levels[i][1]);
      }
      return returntext.trim();
    },

    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },

    fromNow(time) {
      return moment(time).fromNow()
    },

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
    async loadEvents() {
      if (this.requesting) {
        return;
      }

      let builder = (new SpireQueryBuilder())
      builder.includes(
        [
          "Account",
          "CharacterDatum",
          "Zone",
        ]
      )

      if (this.eventType) {
        builder.where("event_type_id", "=", this.eventType)
      }

      if (this.characterId > 0) {
        builder.where("character_id", "=", this.characterId)
      }

      if (this.zoneId > 0) {
        builder.where("zone_id", "=", this.zoneId)
      }

      builder.page(this.currentPage)
      builder.limit(this.pageLimit)
      builder.orderBy(["id"])
      builder.orderDirection("desc")

      // @ts-ignore
      this.requesting = true
      const r         = await (new PlayerEventLogApi(...SpireApi.cfg())).listPlayerEventLogs(builder.get())
      let events      = []
      if (r.status === 200) {
        events          = r.data
        this.requesting = false
      }

      // get total count
      builder.includes([])
      builder.select(["id"])
      builder.limit(10000000000000)
      builder.page(0)

      SpireApi.v1().get(`player_event_logs/count`, { params: builder.get() }).then((r) => {
        if (r.status === 200) {
          this.totalRows = r.data.count
        }
      })

      let shouldPreload = false
      let npcIds        = []
      let itemIds       = []
      for (let e of events) {
        let d = JSON.parse(e.event_data)
        if (d && d.npc_id && !Npcs.cacheExists(d.npc_id)) {
          npcIds.push(d.npc_id)
          shouldPreload = true
        }
        if (d && d.item_id && !Items.cacheExists(d.item_id)) {
          itemIds.push(d.item_id)
          shouldPreload = true
        }
      }

      if (!AA.isPreloaded() || !Zones.isPreloaded()) {
        shouldPreload = true
      }

      if (shouldPreload) {
        await Promise.all(
          [
            AA.preLoad(),
            Zones.getZones(),
            Npcs.getNpcsBulk(npcIds),
            Items.loadItemsBulk(itemIds)
          ]
        ).then(async (r) => {
          console.log("Preloading done")
          this.events         = events
          this.initialLoading = false
          this.loading        = false

        });
      } else {
        this.events         = events
        this.initialLoading = false
        this.loading        = false
      }
    },

    startTimer() {
      if (Timer.timer['player-event-refresh']) {
        clearInterval(Timer.timer['player-event-refresh'])
      }

      Timer.timer['player-event-refresh'] = setInterval(async () => {
        if (!document.hidden) {
          this.loadEvents();
        }
      }, this.refreshInterval)
    },

    stopTimer() {
      if (Timer.timer['player-event-refresh']) {
        clearInterval(Timer.timer['player-event-refresh'])
      }
    }
  },

  beforeDestroy() {
    Navbar.expand()
    this.stopTimer()
  },

  async mounted() {
    Navbar.collapse()

    // non-reactive
    this.requesting = false;


    this.loadQueryState()
    const r = await (new PlayerEventLogSettingApi(...SpireApi.cfg())).listPlayerEventLogSettings()
    if (r.status === 200) {
      this.settings = r.data
    }

    this.startTimer()

    this.initialLoading = true
    setTimeout(() => {
      this.loadEvents();
    }, 1)
  }
}
</script>

<style scoped>
.player-events td, .player-events th {
  text-align: center;
}

.player-events td A {
  color: lightblue !important;
}
</style>
