<template>
  <content-area>
    <eq-window
      v-if="zoneData"
      :title="zoneData.long_name + ` (${zoneData.short_name}) (${zoneData.version})`">
      {{npcTypes.length}} NPC(s)
    </eq-window>
    <eq-window
      style="overflow-x: scroll; height: 88vh"
      v-if="npcTypes"
    >
      <table
        id="npcs-table"
        class="eq-table eq-highlight-rows"
        style="font-size: 14px; "
        v-if="npcTypes && npcTypes.length > 0"
      >
        <thead class="eq-table-floating-header" style="top: -20px">
        <tr>
          <th
            v-for="(header, index) in Object.keys(npcTypes[0])"
            :style="'text-align: center; ' + getColumnHeaderWidth(header) + '' + ([0, 1].includes(index) ? ' position: sticky; z-index: 9999; background-color: rgba(25,31,41, 1); ' + getLeftOffsetFromIndex(index) : '')"
          >{{ header }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(row, index) in npcTypes" :key="index">
          <td
            :style="' text-align: center; ' + ([0, 1].includes(colIndex) ? ' position: sticky; z-index: 999; background-color: rgba(25,31,41, .6);' + getLeftOffsetFromIndex(colIndex): '')"
            v-for="(key, colIndex) in Object.keys(row)"
            v-if="doesRowColumnHaveObjects(row, key)"
          >
            {{ row[key] }}
          </td>
        </tr>
        </tbody>
      </table>

    </eq-window>
  </content-area>

</template>

<script>
import EqWindow            from "../../components/eq-ui/EQWindow";
import ContentArea         from "../../components/layout/ContentArea";
import {Navbar}            from "../../app/navbar";
import {Zones}             from "../../app/zones";
import {Spawn2Api}         from "../../app/api";
import {SpireApiClient}    from "../../app/api/spire-api-client";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import Tablesort           from "@/app/utility/tablesort.js";

export default {
  name: "NPCs",
  components: { ContentArea, EqWindow },
  data() {
    return {
      zone: "",
      version: "",

      zoneData: {},
    }
  },
  beforeDestroy() {
    Navbar.expand()
  },

  mounted() {
    this.init()
  },

  created() {

    // data
    this.npcTypes = []
  },

  methods: {
    getColumnHeaderWidth(header) {
      if (header.includes("lastname")) {
        return 'min-width: 200px; '
      }

      return ''
    },

    getLeftOffsetFromIndex(index) {
      if (index === 1) {
        return 'left: 77px;';
      }

      return 'left: 0px;';
    },

    doesColumnHaveObjects(data, column) {
      if (typeof column === 'object') {
        return true
      }

      return data.find((row) => {
        return typeof row[column] === 'object' && row[column] !== null && Object.keys(row[column])
      })
    },
    doesRowColumnHaveObjects(r, key) {
      return (typeof r[key] !== 'undefined') && !(typeof r[key] === 'object' && r[key] !== null && Object.keys(r[key]))
    },

    async init() {
      // pull from router
      this.zone    = this.$route.params.zone
      this.version = this.$route.query.v

      // get zone data
      this.zoneData = (await Zones.getZoneByShortName(this.zone))

      Navbar.collapse()

      this.loadNpcTypes().then((r) => {
        if (this.npcTypes.length > 0) {
          if (document.getElementById('npcs-table')) {
            new Tablesort(document.getElementById('npcs-table'));
          }
        }
      })


    },

    async loadNpcTypes() {
      const api   = (new Spawn2Api(SpireApiClient.getOpenApiConfig()))
      let builder = (new SpireQueryBuilder())
      builder.where("zone", "=", this.zoneData.short_name)
      builder.where("version", "=", this.zoneData.version)
      builder.includes([
        "Spawnentries.NpcType",
        // "Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
        // "Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList",
        // "Spawnentries.NpcType.NpcFactions",
        // "Spawnentries.NpcType.NpcEmotes",
        // "Spawnentries.NpcType.Merchantlists.Items",
        // "Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item"
      ])

      let npcTypes = [];
      const r      = await api.listSpawn2s(builder.get())
      if (r.status === 200 && r.data) {
        for (let spawn2 of r.data) {
          if (spawn2.spawnentries) {
            for (let spawnentry of spawn2.spawnentries) {
              if (spawnentry.npc_type) {

                // make sure we only add unique NPC IDs since spawns can use multiple
                // of the same NPC ID
                if (npcTypes.filter(f => f.id === spawnentry.npc_type.id).length === 0) {
                  npcTypes.push(
                    spawnentry.npc_type
                  )
                }

              }
            }
          }
        }

        // sort alpha, upper case first
        npcTypes = npcTypes.sort((a, b) => {
          if (this.startsWithUppercase(a.name) && !this.startsWithUppercase(b.name)) {
            return -1;
          } else if (this.startsWithUppercase(b.name) && !this.startsWithUppercase(a.name)) {
            return 1;
          }
          return a.name.localeCompare(b.name);
        });

        this.npcTypes = npcTypes

        this.$forceUpdate()
      }
    },

    startsWithUppercase(str) {
      return str.substr(0, 1).match(/[A-Z\u00C0-\u00DC]/);
    },
  }
}
</script>

<style scoped>

</style>
