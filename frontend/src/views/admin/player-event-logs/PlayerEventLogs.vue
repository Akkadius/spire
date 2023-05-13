<template>
  <div>


    <v-runtime-template/>

    <eq-window
      class="pb-1"
      title="Player Event Log Explorer"
    >
      <app-loader :is-loading="initialLoading"/>


      <div
        class="row justify-content-center"
        style="position: absolute; top: -10px; z-index: 9999999; width: 100%"
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

          <div class="col-1 text-center font-weight-bold">
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

          <div
            class="col-5 text-center font-weight-bold minified-inputs p-0"
            v-if="filters && filters.length > 0"
          >
            <div>Event Filter(s)</div>

            <div v-for="f in filters" :key="f.key">

              <input
                type="text"
                class="form-control"
                style="width: 240px"
                v-model="f.key"
                v-on:keyup.enter="updateQueryState()"
              >

              <select
                class="form-control form-control-prepended list-search"
                v-model="f.operator"
                style="width: 100px"
                @change="updateQueryState()"
              >
                <option
                  v-for="s in filterOperators"
                  v-bind:value="s.operator"
                >
                  {{ s.desc }}
                </option>
              </select>

              <input
                type="text"
                class="form-control"
                style="width: 200px"
                v-model="f.value"
                v-on:keyup.enter="updateQueryState()"
              >

              <b-button
                @click="deleteFilter(f)"
                size="sm"
                class="ml-3"
                variant="outline-danger"
                title="Remove Filter"
              >
                <i class="fa fa-remove"></i>
              </b-button>

            </div>

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
          class="eq-table bordered player-events"
          v-if="events.length > 0"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th style="width: 150px">Event ID</th>
            <th style="width: 150px" class="text-center">Player</th>
            <th style="width: 230px">Zone</th>

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
            <td :style="(characterId ? 'background-color: rgba(123, 113, 74, .1);' : '')">

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
            <td :style="(zoneId ? 'background-color: rgba(123, 113, 74, .1);' : '')">
              <a
                class="ml-1"
                @click="zoneId = e.zone_id; updateQueryState()"
              >
                {{ getZoneLongName(e.zone_id) }}
              </a>
              ({{ e.zone_id }})
            </td>

            <td class="text-center" :style="(eventType ? 'background-color: rgba(123, 113, 74, .1);' : '')">
              <a
                class="ml-1"
                @click="eventType = e.event_type_id; updateQueryState()"
              >{{ e.event_type_name }}</a> ({{ e.event_type_id }})
            </td>

            <td
              style="vertical-align: middle; text-align: left;"
            >
              <player-event-display-component
                :e="e"
              />

              <button
                title="View raw"
                @click="showRawEvent(e)"
                v-if="!showRaw[e.id] && Object.keys(JSON.parse(e.event_data)).length > 0"
                class="ml-3 btn btn-sm btn-warning"
                style="font-size: 10px; min-width: 50px"
              >
                <i class="fa fa-search"></i> ({{ Object.keys(JSON.parse(e.event_data)).length }})
              </button>

              <!--              <pre-->
              <!--                v-if="showRaw[e.id]"-->
              <!--                class="text-left code fade-in mt-2"-->
              <!--                style="width: 100%; padding: 0 !important; margin-bottom: 0 !important"-->
              <!--              ><code class="language-json">{{ formatPayload(e.event_data.replaceAll("    ", "  ")) }}</code></pre>-->

              <pre
                v-if="showRaw[e.id]"
                class="text-left code fade-in mt-2"
                style="width: 100%; padding: 0 !important; margin-bottom: 0 !important"
              ><code class="language-json"><v-runtime-template :template="'<div>' + formatPayload(e) + '</div>'"/></code></pre>

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
          v-if="!loading && currentPage > 0"
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
import {PlayerEventLogSettingApi}  from "@/app/api/api/player-event-log-setting-api";
import Timer                       from "@/app/timer/timer";
import EqProgressBar               from "@/components/eq-ui/EQProgressBar.vue";
import LoaderFakeProgress          from "@/components/LoaderFakeProgress.vue";
import hljs                        from "highlight.js";
import {Navbar}                    from "@/app/navbar";
import {Characters}                from "@/app/characters";
import util                        from "util";
import InfoErrorBanner             from "@/components/InfoErrorBanner.vue";

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
  components: {
    InfoErrorBanner,
    "v-runtime-template": () => import("v-runtime-template"),
    LoaderFakeProgress,
    EqProgressBar,
    PlayerEventDisplayComponent,
    EqWindow
  },
  data() {
    return {
      search: "",
      eventType: 0,
      zoneId: null,
      characterId: null,

      notification: "",
      error: "",

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


      // filters
      filterOperators: [
        { operator: "=", desc: "equals" },
        { operator: "like", desc: "contains" },
      ],

      // state filters
      filters: [],

      // pagination (all)
      currentPage: 1,
      totalRows: 0,
      pageLimit: 100,

      objCount: 0,
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

    getZoneLongName(zoneId) {
      const z = Zones.getZoneByIdSync(zoneId)

      return z ? z.long_name : ""
    },

    deleteFilter(f) {
      this.filters = this.filters.filter((e) => {
        return e.key !== f.key
      })
      this.updateQueryState()
    },

    handleClick(e) {
      e.preventDefault()

      if (e.target && e.target.getAttribute('filter-key')) {
        const filterValue = e.target.getAttribute('value')
        const filterKey   = e.target.getAttribute('filter-key')
        const anyFilter   = e.target.getAttribute('any-filter') ? e.target.getAttribute('any-filter') : 0
        const filterEvent = e.target.getAttribute('event')
        if (filterValue && filterKey) {
          const a = {
            key: "." + filterKey,
            operator: anyFilter ? "like" : "=",
            value: filterValue
          }

          // don't add the same filter
          const found = this.filters.find((f) => {
            return f.key === a.key && f.operator === a.operator && f.value === a.value
          })

          if (found) {
            return;
          }

          this.filters.push(a)

          if (filterEvent) {
            this.eventType = parseInt(filterEvent)
          }
        }
        this.updateQueryState()
      }
    },

    formatPayload(e) {
      const dot = require("dot-object")
      const f   = JSON.stringify(dot.dot(JSON.parse(e.event_data.replaceAll('    ', '  '))));

      let lines = []
      for (let line of f.split(",\"")) {
        line       = line.replaceAll("{", "")
        line       = line.replaceAll("}", "")
        line       = line.replaceAll("\"", "")
        const s    = line.split(":")
        const key  = s[0].trim()
        let values = []
        for (let i = 1; i < s.length; i++) {
          values.push(s[i])
        }

        let value    = values.join(":")
        let rawValue = values.join(":")
        if (!Number.isInteger(value)) {
          value = util.format('"%s"', value)
        }

        let equalLink = util.format(
          "<a href=\"#\" event=\"%s\" filter-key=\"%s\" value=\"%s\"><i class='fa fa-filter'></i> = [%s]</a>",
          e.event_type_id,
          key,
          rawValue,
          rawValue
        );

        let anyLink = ""
        if (key.includes("[")) {
          let label       = key.split("]")[1].trim()
          const filterKey = key.replace(/\[.*]/, '[*]')

          anyLink = util.format(
            "<a href=\"#\" event=\"%s\" filter-key=\"%s\" value=\"%s\"><i class='fa fa-filter'></i> Any [%s] = [%s]</a>",
            e.event_type_id,
            filterKey,
            rawValue,
            label,
            rawValue
          );
        }

        lines.push(
          util.format(
            '  "%s": %s, %s %s',
            key,
            value,
            equalLink,
            anyLink
          )
        )
      }

      return util.format("{\n%s\n}", lines.join("\n"));
    },

    paginate() {
      // models aren't quite updated when we trigger this so queue the pagination
      setTimeout(() => {
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
      this.filters     = []
      this.eventType   = 0;
      this.zoneId      = null
      this.characterId = null
      this.showRaw     = {}
      this.currentPage = 1
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
      if (this.filters && this.filters.length > 0) {
        q.filters = JSON.stringify(this.filters)
      }

      this.$router.push(
        {
          path: this.$route.path,
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
      if (typeof this.$route.query.filters !== 'undefined') {
        this.filters = JSON.parse(this.$route.query.filters)
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

      if (this.filters && this.filters.length > 0) {
        for (let f of this.filters) {
          builder.whereJson(
            "event_data",
            f.key,
            f.operator,
            f.value
          );
        }
      }

      // builder.whereJson("event_data", ".target", "=", 'Dargon_McPherson000');
      // builder.whereJson("event_data", ".target", "like", 'Dargon');

      builder.page(this.currentPage)
      builder.limit(this.pageLimit)
      builder.orderBy(["id"])
      builder.orderDirection("desc")

      let events = []
      try {
        // @ts-ignore
        this.requesting = true
        const r         = await (new PlayerEventLogApi(...SpireApi.cfg()))
          .listPlayerEventLogs(
            {},
            {
              params: builder.get()
            }
          )
        if (r.status === 200) {
          events          = r.data
          this.requesting = false
        }
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }

      // get total count
      builder.includes([])
      builder.select(["id"])
      builder.limit(10000000000000)
      builder.page(0)

      try {
        SpireApi.v1().get(`player_event_logs/count`, { params: builder.get() }).then((r) => {
          if (r.status === 200) {
            this.totalRows = r.data.count
          }
        })
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }

      let shouldPreload = false
      let npcIds        = []
      let itemIds       = []
      let characterIds  = []
      for (let e of events) {
        let d = JSON.parse(e.event_data)
        if (d && d.npc_id && !Npcs.cacheExists(d.npc_id) && !npcIds.includes(d.npc_id)) {
          npcIds.push(d.npc_id)
          shouldPreload = true
        }
        if (d && d.item_id && !Items.cacheExists(d.item_id) && !itemIds.includes(d.item_id)) {
          itemIds.push(d.item_id)
          shouldPreload = true
        }
        if (d && d.character_1_id && !Characters.cacheExists(d.character_1_id) && !characterIds.includes(d.character_1_id)) {
          characterIds.push(parseInt(d.character_1_id))
          shouldPreload = true
        }
        if (d && d.character_2_id && !Characters.cacheExists(d.character_2_id) && !characterIds.includes(d.character_2_id)) {
          characterIds.push(parseInt(d.character_2_id))
          shouldPreload = true
        }
        if (d && d.character_1_give_items && d.character_1_give_items.length > 0) {
          for (let i of d.character_1_give_items) {
            if (!Items.cacheExists(i.item_id) && !itemIds.includes(i.item_id)) {
              itemIds.push(i.item_id)
              shouldPreload = true
            }
          }
        }
        if (d && d.character_2_give_items && d.character_2_give_items.length > 0) {
          for (let i of d.character_2_give_items) {
            if (!Items.cacheExists(i.item_id) && !itemIds.includes(i.item_id)) {
              itemIds.push(i.item_id)
              shouldPreload = true
            }
          }
        }
        if (d && d.handin_items && d.handin_items.length > 0) {
          for (let i of d.handin_items) {
            if (!Items.cacheExists(i.item_id) && !itemIds.includes(i.item_id)) {
              itemIds.push(i.item_id)
              shouldPreload = true
            }
          }
        }
      }

      if (!AA.isPreloaded() || !Zones.isPreloaded()) {
        shouldPreload = true
      }

      if (shouldPreload) {
        await Promise.all(
          [
            Characters.bulkLoadCharacters(characterIds),
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
    window.removeEventListener("click", this.handleClick);
  },

  async mounted() {
    Navbar.collapse()

    window.addEventListener("click", this.handleClick);

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

.player-events TBODY TR:hover {
  /*background: linear-gradient(180deg, rgba(123, 113, 74, .1) 0%, rgba(123, 113, 74, 0.4) 50%, rgba(123, 113, 74, .1) 100%);*/
}

</style>
