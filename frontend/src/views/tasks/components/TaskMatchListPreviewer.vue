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
        Found ({{ npcs.length }}) matching NPC(s)
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
import {ItemApi, NpcTypeApi}                     from "@/app/api";
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

      console.log(this.activity.zones)
      console.log("loading npc matches")

      // 1) Load by zone
      let zones = []
      let npcs  = [];

      this.activity.zones = this.activity.zones.toString();
      if (this.activity.zones && this.activity.zones !== 0) {

        console.log("NPC zone search", this.activity.zones)

        // build zone names
        if (this.activity.zones && this.activity.zones.length > 0) {
          for (let zoneId of this.activity.zones.split(",")) {
            zones.push(
              (await Zones.getZoneById(parseInt(zoneId))).short_name
            )
          }
        }

        console.log("zones", zones)

        for (let zone of zones) {
          console.log("zone", zone)
          console.log("version", parseInt(this.activity.zone_version))

          // cache zone npcs so we don't hit it realtime
          let zoneNpcs = []
          if (this.zoneCache[this.activity.zones]) {
            zoneNpcs = this.zoneCache[this.activity.zones]
          } else {
            zoneNpcs = (await Npcs.getNpcsByZone(zone, parseInt(this.activity.zone_version)))
            this.zoneCache[this.activity.zones] = zoneNpcs
          }

          for (let n of zoneNpcs) {
            let found = false
            for (let m of this.activity.npc_match_list.split("|")) {
              if (m === "") {
                continue;
              }

              if (m.toString() === n.id.toString() || n.name.toLowerCase().includes(m.toLowerCase())) {
                found = true;
                break;
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

      return;

      console.log("[match list previewer] Zone ID", this.activity.zones)
      console.log("[match list previewer] Zones", zones)

      let builder = (new SpireQueryBuilder())
      for (let zone of zones) {
        // if (zone.length === 0) {
        //   continue;
        // }
        builder.where("zone", "=", zone)
      }

      if (zones.length > 0) {
        builder.where("version", "=", 0)
      }

      builder.includes([
        "Spawnentries.NpcType",
      ])

      this.npcs = {}

      // this is used only as a way to show options at the global level if things were not
      // found at the local zone level
      // for example if were to search "orc" from the match list, we'd only care to see orcs
      // at the zone level if we have a zone filter
      // if we have no matches against anything in zone we'd want to see what the global npc table
      // could give us for matches
      let npcNameMatches = {};

      // 2) Load by global namespace
      // an NPC could be quest-spawned into a zone and is still filterable
      const npcTypeApi = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))
      builder          = (new SpireQueryBuilder())
      const matchList  = this.goalMatchList ? this.goalMatchList : ""
      let filterCount  = 0
      for (let name of matchList.split("|")) {
        if (!npcNameMatches[name] && name.length > 0) {
          builder.where("name", "like", name)
          filterCount++;
        }
      }

      if (filterCount > 0) {
        const r = await npcTypeApi.listNpcTypes(builder.get())
        if (r.status === 200) {
          for (let npc of r.data) {
            if (npcs.filter(e => e.npc.id === npc.id).length === 0) {
              npcs.push({
                npc: npc,
                search: 'Global'
              })
            }
          }
        }
      }

      this.npcs = npcs
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
