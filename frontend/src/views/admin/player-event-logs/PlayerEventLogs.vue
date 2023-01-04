<template>
  <div>
    <eq-window title="Player Event Log Explorer">
      <div
        style="height: 80vh; overflow-y: scroll"
      >
        <table
          class="eq-table eq-highlight-rows bordered player-events"
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
          <tr v-for="e in events">
            <td>{{ commify(e.id) }}</td>
            <td>

              <div class="avatar-list avatar-list-stacked">
                <img class="avatar-img rounded-circle" style="width:20px" :src="getClassImage(e.character_datum.class)">
                <img class="avatar-img rounded-circle" style="width:20px" :src="getRaceImage(e.character_datum.race)">

                <span class="ml-1">
                  {{ e.character_datum.name }}
                </span>

              </div>
            </td>
            <td>{{ e.zone.long_name }} ({{ e.zone.zoneidnumber }})</td>

            <td class="text-right">{{ e.event_type_name }} ({{ e.event_type_id }})</td>

            <td style="vertical-align: middle; text-align: left">
              <player-event-display-component :e="e"/>
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
// REZ_ACCEPTED         | [] Implemented Formatter
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
      events: [],
    }
  },
  methods: {
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
  },
  async mounted() {
    await AA.preLoad()

    let builder = (new SpireQueryBuilder())
    builder.includes(
      [
        "Account",
        "CharacterDatum",
        "Zone",
      ]
    )
    builder.orderBy(["created_at"])
    builder.orderDirection("desc")

    // @ts-ignore
    const api = (new PlayerEventLogApi(...SpireApi.cfg()))
    const r   = await api.listPlayerEventLogs(builder.get())
    let events = []
    if (r.status === 200) {
      events = r.data
    }

    this.events = events
  }
}
</script>

<style scoped>
.player-events td, .player-events th {
  text-align: center;
}
</style>
