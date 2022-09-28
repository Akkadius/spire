<template>
  <div>
    <eq-window-simple
      style="height: 95vh; overflow-y: scroll; overflow-x: hidden" class="p-0"
      v-if="isItemIdMatchList()"
    >
      <div class="font-weight-bold text-center">
        Goal Match List Preview ({{ TASK_ACTIVITY_TYPES[activity.activitytype] }})
      </div>

      <div v-if="items && items.length > 0" class="text-center mt-1">
        Found ({{ items.length }}) matching item(s)
      </div>

      <table
        class="eq-table eq-highlight-rows bordered"
        style="display: table; overflow-x: scroll"
      >
        <thead>
        <tr>
          <th style="text-align: center; width: 50px" class="text-center">Id</th>
          <th style="text-align: center;">Name</th>
        </tr>
        </thead>
        <tbody>
        <tr
          v-for="(item, index) in items"
          :key="item.id"
          :id="'item-selection-row-' + item.id"
        >
          <td>
            {{ item.id }}
          </td>
          <td class="text-left" style="vertical-align: middle">
            <item-popover
              :item="item"
              v-if="Object.keys(item).length > 0 && item"
              size="regular"
            />
          </td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>

    <eq-window-simple
      style="height: 95vh; overflow-y: scroll; overflow-x: hidden" class="p-0"
      v-if="isNpcMatchList()"
    >
      <div class="font-weight-bold text-center">
        Goal Match List Preview ({{ TASK_ACTIVITY_TYPES[activity.activitytype] }})
      </div>

      <div v-if="npcs && npcs.length > 0" class="text-center mt-1 ">
        Found ({{ npcs.length }}) matching NPC(s). <br> NPC(s) can still be matched if they are quest spawned and within the filtered zone(s)
      </div>

      <table
        id="npctable"
        class="eq-table eq-highlight-rows bordered"
        v-if="npcs && npcs.length > 0"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 30px">ID</th>
          <th></th>
          <th>Search Type</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'npc-' + npc.npc.short_name"
          v-for="(npc, index) in npcs"
          :key="index + '-' + npc.npc.id"
        >
          <td style="text-align: center">{{ npc.npc.id }}</td>
          <td style="min-width: 250px">
            <npc-popover :npc="npc.npc" size="regular"/>
          </td>
          <td>{{ npc.search }}</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple                            from "@/components/eq-ui/EQWindowSimple";
import {ItemApi}                                 from "@/app/api";
import {SpireApiClient}                          from "@/app/api/spire-api-client";
import EqCheckbox                                from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder}                       from "@/app/api/spire-query-builder";
import {Zones}                                   from "@/app/zones";
import {TASK_ACTIVITY_TYPE, TASK_ACTIVITY_TYPES} from "@/app/constants/eq-task-constants";
import ItemPopover                               from "@/components/ItemPopover";
import NpcPopover                                from "@/components/NpcPopover";
import {Npcs}                                    from "@/app/npcs";

export default {
  name: "TaskGoalMatchListPreviewer",
  components: { NpcPopover, ItemPopover, EqCheckbox, EqWindowSimple },
  data() {
    return {
      // result sets
      npcs: {},
      items: {},

      // constants
      TASK_ACTIVITY_TYPE: TASK_ACTIVITY_TYPE,
      TASK_ACTIVITY_TYPES: TASK_ACTIVITY_TYPES,
    }
  },
  props: {
    activity: {
      type: Object,
      required: true,
    },
  },

  watch: {
    activity: {
      deep: true,
      handler() {
        this.load()
      }
    },
  },

  created() {
    this.zoneCache = {}
  },

  methods: {
    isItemIdMatchList() {
      return [
        TASK_ACTIVITY_TYPE.LOOT,
        TASK_ACTIVITY_TYPE.TRADESKILL,
        TASK_ACTIVITY_TYPE.DELIVER,
        TASK_ACTIVITY_TYPE.FISH,
        TASK_ACTIVITY_TYPE.FORAGE
      ].includes(
        parseInt(this.activity.activitytype)
      )
    },

    isNpcMatchList() {
      return [
        TASK_ACTIVITY_TYPE.KILL,
        TASK_ACTIVITY_TYPE.DELIVER,
        TASK_ACTIVITY_TYPE.SPEAK_WITH,
        // TASK_ACTIVITY_TYPE.GIVE
      ].includes(
        parseInt(this.activity.activitytype)
      )
    },

    async load() {
      if (this.isNpcMatchList()) {
        this.loadNpcMatches()
      }
      if (this.isItemIdMatchList()) {
        this.loadItemMatches()
      }
    },

    async loadItemMatches() {
      const api       = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      let builder     = (new SpireQueryBuilder())
      const matchList = this.goalMatchList ? this.goalMatchList : ""
      for (let m of matchList.split("|")) {
        m = m.toLowerCase()
        if (m.length === 0 && matchList.length !== 0) {
          continue;
        }

        builder.whereOr("id", "=", m)
      }

      const result = await api.listItems(builder.get())
      if (result.status === 200) {
        this.items = result.data
      }
    },

    async loadNpcMatches() {

      // 1) Load by zone
      // TODO: skip loading globally for now, maybe implement later
      let zones = []
      let npcs  = [];


      console.log("this.activity.zones", this.activity.zones)

      this.activity.zones = this.activity.zones.toString();
      if (this.activity.zones && this.activity.zones !== 0) {

        // build zone names
        if (this.activity.zones && this.activity.zones.length > 0) {
          for (let zoneId of this.activity.zones.split(",")) {
            zones.push(
              (await Zones.getZoneById(parseInt(zoneId))).short_name
            )
          }
        }

        for (let zone of zones) {

          // cache zone npcs so we don't hit it realtime
          let zoneNpcs = []
          if (this.zoneCache[this.activity.zones]) {
            zoneNpcs = this.zoneCache[this.activity.zones]
          } else {
            zoneNpcs                            = (await Npcs.getNpcsByZone(zone, parseInt(this.activity.zone_version)))
            this.zoneCache[this.activity.zones] = zoneNpcs
          }

          for (let n of zoneNpcs) {
            let found = true

            // if match list is not empty - lets filter
            if (this.activity.npc_match_list && this.activity.npc_match_list.length > 0) {
              found = false
              for (let m of this.activity.npc_match_list.split("|")) {

                console.log(m)

                if (m === "") {
                  continue;
                }


                if (
                  m.toString() === n.id.toString()
                  || n.name.toLowerCase().includes(m.toLowerCase())
                  || Npcs.getCleanName(n.name).toLowerCase().includes(m.toLowerCase())
                ) {
                  found = true;
                  break;
                }
              }
            }

            if (found) {
              npcs.push({
                npc: n,
                search: `Zone [${zone}]`
              })
            }

          }
        }

        this.npcs = npcs
      }


    }
  },
  mounted() {
    this.load()
  }
}
</script>

<style scoped>
#npctable td {
  vertical-align: middle !important;
}
</style>
