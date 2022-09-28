<template>
  <div>
    <eq-window-simple
      style="height: 95vh; overflow-y: scroll; overflow-x: hidden" class="p-0"
    >
      <div class="font-weight-bold text-center">
        Goal Match List Preview ({{ TASK_ACTIVITY_TYPES[activity.activitytype] }})
      </div>

      <div v-if="npcs && npcs.length > 0" class="text-center mt-1 ">
        Found ({{ npcs.length }}) matching NPC(s). <br> NPC(s) can still be matched if they are quest spawned and within
        the filtered zone(s)
      </div>

      <div v-if="npcs.length === 0" class="text-center mt-3 font-weight-bold">
        No NPC(s) were found
      </div>

      <!-- Fake Loader -->
      <div v-if="!loaded" class="mt-3 text-center">
        <!--        <loader-fake-progress class="mt-3"/>-->
        <app-loader :is-loading="!loaded"/>
      </div>

      <table
        id="npctable"
        class="eq-table eq-highlight-rows bordered"
        v-if="npcs && npcs.length > 0 && loaded"
        style="display: table; font-size: 14px; overflow-x: scroll; overflow-y: hidden"
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
            <npc-popover :npc="npc.npc" size="regular" :show-last-name="false"/>
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
import EqCheckbox                                from "@/components/eq-ui/EQCheckbox";
import {Zones}                                   from "@/app/zones";
import {TASK_ACTIVITY_TYPE, TASK_ACTIVITY_TYPES} from "@/app/constants/eq-task-constants";
import ItemPopover                               from "@/components/ItemPopover";
import NpcPopover                                from "@/components/NpcPopover";
import {Npcs}                                    from "@/app/npcs";
import LoaderFakeProgress                        from "@/components/LoaderFakeProgress";

export default {
  name: "TaskNpcMatchListPreviewer",
  components: { LoaderFakeProgress, NpcPopover, ItemPopover, EqCheckbox, EqWindowSimple },
  data() {
    return {
      // result sets
      npcs: {},

      loaded: false,

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

    async load() {
      setTimeout(() => {
        this.loadNpcMatches()
      }, 1)
    },

    async loadNpcMatches() {

      // 1) Load by zone
      // TODO: skip loading globally for now, maybe implement later
      let zones           = []
      let npcs            = [];
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

      // Global
      if (this.npcs.length === 0 && this.activity.npc_match_list.length > 0) {
        let npcIds = []
        for (let m of this.activity.npc_match_list.split("|")) {
          if (m === "") {
            continue;
          }
          npcIds.push(parseInt(m))
        }
        let npcs = []
        for (let n of await Npcs.getNpcsBulk(npcIds)) {
          npcs.push({
            npc: n,
            search: 'Global'
          })
        }
        this.npcs = npcs
      }

      this.loaded = true
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
