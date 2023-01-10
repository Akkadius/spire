<template>
  <div>
    <eq-window title="Player Event Log Explorer">
      <div
        style="max-height: 80vh; overflow-y: scroll;  overflow-x: hidden"
      >
        <app-loader :is-loading="loading"/>

        <div v-if="!loading">

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

        <div v-if="events.length === 0" class="font-weight-bold text-center mt-3">
          No events found
        </div>

        <table
          class="eq-table eq-highlight-rows bordered player-events"
          v-if="events.length > 0"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th>ID</th>
            <th style="width: 150px" class="text-center">Player</th>
            <th style="width: 200px">Zone</th>

            <th class="text-right" style="width: 175px">Event</th>
            <th>Event</th>
            <th>Time</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="e in events" :key="e.id">
            <td>{{ commify(e.id) }}</td>
            <td>

              <div class="avatar-list avatar-list-stacked">
                <img class="avatar-img rounded-circle" style="width:20px" :src="getClassImage(e.character_datum.class)">
                <img class="avatar-img rounded-circle" style="width:20px" :src="getRaceImage(e.character_datum.race)">

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

            <td class="text-right">
              <a
                class="ml-1"
                @click="eventType = e.event_type_id; updateQueryState()"
              >{{ e.event_type_name }}</a> ({{ e.event_type_id }})
            </td>

            <td style="vertical-align: middle; text-align: left">
              <player-event-display-component
                :e="e"
              />
              <!--              <pre style="width: 100%">{{ e.event_data }}</pre>-->
            </td>
            <td>{{ fromNow(e.created_at) }}</td>
          </tr>
          </tbody>
        </table>
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
  components: { PlayerEventDisplayComponent, EqWindow },
  data() {
    return {
      search: "",
      eventType: 0,
      zoneId: null,
      characterId: null,

      loading: false,

      settings: [],

      events: [],
    }
  },
  watch: {
    $route(to, from) {
      this.loadQueryState()
      setTimeout(() => {
        this.loadEvents();
      }, 1)
    }
  },
  methods: {
    reset() {
      this.search      = "";
      this.eventType   = 0;
      this.zoneId      = null
      this.characterId = null
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

      builder.orderBy(["created_at"])
      builder.orderDirection("desc")

      // @ts-ignore
      const r    = await (new PlayerEventLogApi(...SpireApi.cfg())).listPlayerEventLogs(builder.get())
      let events = []
      if (r.status === 200) {
        events = r.data
      }

      let npcIds  = []
      let itemIds = []
      for (let e of events) {
        let d = JSON.parse(e.event_data)
        if (d && d.npc_id) {
          npcIds.push(d.npc_id)
        }
        if (d && d.item_id) {
          itemIds.push(d.item_id)
        }
      }

      await Promise.all(
        [
          AA.preLoad(),
          Zones.getZones(),
          Npcs.getNpcsBulk(npcIds),
          Items.loadItemsBulk(itemIds)
        ]
      ).then(async (r) => {
        console.log("Preloading done")
        this.events  = events
        this.loading = false
      });
    }
  },
  async mounted() {
    this.loadQueryState()
    const r = await (new PlayerEventLogSettingApi(...SpireApi.cfg())).listPlayerEventLogSettings()
    if (r.status === 200) {
      this.settings = r.data
    }

    this.loading = true
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
