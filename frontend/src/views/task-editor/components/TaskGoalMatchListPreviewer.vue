<template>
  <div>
    <eq-window-simple
      style="height: 95vh; overflow-y: scroll;" class="p-0 eq-window-hybrid"
      v-if="isItemIdMatchList()"
    >
      <div class="font-weight-bold text-center">
        Goal Match List Preview ({{ TASK_ACTIVITY_TYPES[activityType] }})
      </div>

      <div v-if="items && items.length > 0" class="text-center mt-1">
        Found ({{ items.length }}) matching item(s)
      </div>

      <table
        class="eq-table eq-highlight-rows"
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
      style="height: 95vh; overflow-y: scroll;" class="p-0 eq-window-hybrid"
      v-if="isNpcMatchList()"
    >
      <div class="font-weight-bold text-center">
        Goal Match List Preview ({{ TASK_ACTIVITY_TYPES[activityType] }})
      </div>

      <div v-if="npcs && npcs.length > 0" class="text-center mt-1 ">
        Found ({{ npcs.length }}) matching NPC(s)
      </div>

      <table
        id="npctable"
        class="eq-table eq-highlight-rows"
        v-if="npcs && npcs.length > 0"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 30px">ID</th>
          <th style="width: 100px">Name</th>
          <th></th>
          <th>Search Type</th>
          <th>Zones</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'npc-' + npc.npc.short_name"
          v-for="(npc, index) in npcs"
          :key="npc.npc.id"
        >
          <td style="text-align: center" class="p-0">{{ npc.npc.id }}</td>
          <td>{{ npc.npc.name }} <span v-if="npc.lastname && npc.lastname.length > 0">({{ npc.lastname }})</span></td>
          <td class="text-center"><span
            :class="'race-models-ctn-' + npc.npc.race + '-' + npc.npc.gender + '-' + npc.npc.texture + '-' + npc.npc.helmtexture + ''"
            style="zoom: 75%;"
          ></span></td>
          <td>{{npc.search}}</td>
          <td><span v-if="npcZones[npc.npc.id]">{{ npcZones[npc.npc.id].join(",") }}</span></td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple                            from "@/components/eq-ui/EQWindowSimple";
import {ItemApi, NpcTypeApi, Spawn2Api}          from "@/app/api";
import {SpireApiClient}                          from "@/app/api/spire-api-client";
import EqCheckbox                                from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder}                       from "@/app/api/spire-query-builder";
import {Zones}                                   from "@/app/zones";
import {TASK_ACTIVITY_TYPE, TASK_ACTIVITY_TYPES} from "@/app/constants/eq-task-constants";
import ItemPopover                               from "@/components/ItemPopover";

export default {
  name: "TaskGoalMatchListPreviewer",
  components: { ItemPopover, EqCheckbox, EqWindowSimple },
  data() {
    return {
      // result sets
      npcs: {},
      items: {},

      // key by npcid
      npcZones: {},

      // constants
      TASK_ACTIVITY_TYPE: TASK_ACTIVITY_TYPE,
      TASK_ACTIVITY_TYPES: TASK_ACTIVITY_TYPES,
    }
  },
  props: {
    goalMatchList: {
      type: String,
      required: false,
    },
    activityType: {
      type: Number,
      required: true,
    },
    zoneIds: {
      type: String,
      required: false,
    }
  },

  watch: {
    goalMatchList: {
      handler() {
        this.load()
      }
    },
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
        parseInt(this.activityType)
      )
    },

    isNpcMatchList() {
      return [
        TASK_ACTIVITY_TYPE.KILL,
        TASK_ACTIVITY_TYPE.SPEAK_WITH,
        // TASK_ACTIVITY_TYPE.GIVE
      ].includes(
        parseInt(this.activityType)
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
      const api = (new Spawn2Api(SpireApiClient.getOpenApiConfig()))
      let zones = []
      if (this.zoneIds && this.zoneIds.length > 0) {
        for (let zoneId of this.zoneIds.split(",")) {
          zones.push(
            (await Zones.getZoneById(parseInt(zoneId))).short_name
          )
        }
      }

      console.log("[match list previewer] Zone ID", this.zoneIds)
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

      this.npcs     = {}
      this.npcZones = {}
      let npcs      = [];

      // this is used only as a way to show options at the global level if things were not
      // found at the local zone level
      // for example if were to search "orc" from the match list, we'd only care to see orcs
      // at the zone level if we have a zone filter
      // if we have no matches against anything in zone we'd want to see what the global npc table
      // could give us for matches
      let npcNameMatches = {};

      const result = await api.listSpawn2s(builder.get())
      if (result.status === 200 && result.data) {
        for (let spawn2 of result.data) {
          if (spawn2.spawnentries) {
            for (let spawnentry of spawn2.spawnentries) {
              if (spawnentry.npc_type) {
                const n         = spawnentry.npc_type.name.toLowerCase()
                const nId       = spawnentry.npc_type.id.toString()
                const matchList = this.goalMatchList ? this.goalMatchList : ""
                for (let m of matchList.split("|")) {
                  m = m.toLowerCase()
                  if (m.length === 0 && matchList.length !== 0) {
                    continue;
                  }

                  // name match
                  if (n.includes(m) || nId === m) {

                    npcNameMatches[m] = true

                    // make sure npc id isn't already added to array
                    if (npcs.filter(e => e.npc.id === spawnentry.npc_type.id).length === 0) {
                      npcs.push({
                        npc: spawnentry.npc_type,
                        search: 'Zone'
                      })

                    }

                    // create association of an NPC ID to zones
                    for (let zone of zones) {
                      if (typeof this.npcZones[spawnentry.npc_type.id] === "undefined") {
                        this.npcZones[spawnentry.npc_type.id] = []
                      }

                      // make sure we haven't added the same zone twice
                      if (this.npcZones[spawnentry.npc_type.id].filter(e => e === zone).length === 0) {
                        this.npcZones[spawnentry.npc_type.id].push(zone)
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }

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
