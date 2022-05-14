<template>
  <div>
    <eq-window-simple
      id="npc-view-container"
      style="height: 95vh; overflow-y: scroll;" class="p-0 eq-window-hybrid"
    >
      <div v-if="npcs && npcs.length > 0" class="text-center">
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
          <th>Zones</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'npc-' + npc.short_name"
          v-for="(npc, index) in npcs"
          :key="npc.id"
        >
          <td style="text-align: center" class="p-0">{{ npc.id }}</td>
          <td>{{ npc.name }} <span v-if="npc.lastname && npc.lastname.length > 0">({{ npc.lastname }})</span></td>
          <td class="text-center"><span
            :class="'race-models-ctn-' + npc.race + '-' + npc.gender + '-' + npc.texture + '-' + npc.helmtexture + ''"
            style="zoom: 75%;"
          ></span></td>
          <td><span v-if="npcZones[npc.id]">{{ npcZones[npc.id].join(",") }}</span></td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple       from "@/components/eq-ui/EQWindowSimple";
import {Spawn2Api}          from "@/app/api";
import {SpireApiClient}     from "@/app/api/spire-api-client";
import EqCheckbox           from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder}  from "@/app/api/spire-query-builder";
import {Zones}              from "@/app/zones";
import {TASK_ACTIVITY_TYPE} from "@/app/constants/eq-task-constants";

export default {
  name: "TaskGoalMatchListPreviewer",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      // filtered content
      npcs: {},

      // key by npcid
      npcZones: {},
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

    async load() {
      if (this.activityType === TASK_ACTIVITY_TYPE.KILL) {
        this.loadNpcMatches()
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

      const result = await api.listSpawn2s(builder.get())
      if (result.status === 200) {
        let npcs = [];

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

                    // make sure npc id isn't already added to array
                    if (npcs.filter(e => e.id === spawnentry.npc_type.id).length === 0) {
                      npcs.push(spawnentry.npc_type)
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

        // 2) Load by global namespace
        // an NPC could be quest-spawned into a zone and is still filterable

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
