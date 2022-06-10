<template>
  <content-area>
    <eq-window
      v-if="zoneData"
      :title="`${zoneData.long_name} Short Name (${zoneData.short_name}) Version (${zoneData.version}) NPC(s) (${npcTypes.length})`"
    >
      <div class="row">
        <div class="col-1 text-right">
          <button
            class='btn btn-outline-warning btn-sm mt-1'
            @click="reset"
          >
            <i class="fa fa-refresh"></i> Reset
          </button>
        </div>

        <div class="col-3">
          <b-input
            placeholder="Search by NPC name"
            v-on:keyup.enter="updateQueryState"
            v-model="npcNameSearch"
          ></b-input>
        </div>

        <div class="col-6 p-0">
          <db-column-filter
            v-if="npcTypeFields && filters"
            :set-filters="filters"
            @input="handleDbColumnFilters($event);"
            :columns="npcTypeFields"
          />
        </div>


        <!--        <div class="col-2">-->
        <!--          {{ npcTypes.length }} NPC(s)-->
        <!--        </div>-->
      </div>
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
            :style="'text-align: center; ' + getColumnHeaderWidth(header) + '' + ([0, 1].includes(index) ? ' position: sticky; z-index: 9999; background-color: rgba(25,31,41, 1); ' + getColumnStylingFromIndex(index) : '')"
          >{{ header }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(row, index) in npcTypes" :key="index">
          <td
            :style="' text-align: center; ' + ([0, 1].includes(colIndex) ? ' position: sticky; z-index: 999; background-color: rgba(25,31,41, .6);' + getColumnStylingFromIndex(colIndex): '')"
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
import EqWindow                from "../../components/eq-ui/EQWindow";
import ContentArea             from "../../components/layout/ContentArea";
import {Navbar}                from "../../app/navbar";
import {Zones}                 from "../../app/zones";
import {NpcTypeApi, Spawn2Api} from "../../app/api";
import {SpireApiClient}        from "../../app/api/spire-api-client";
import {SpireQueryBuilder}     from "../../app/api/spire-query-builder";
import Tablesort               from "@/app/utility/tablesort.js";
import DbColumnFilter          from "../../components/DbColumnFilter";
import {DbSchema}              from "../../app/db-schema";
import {ROUTE}                 from "../../routes";

export default {
  name: "NPCs",
  components: { DbColumnFilter, ContentArea, EqWindow },
  data() {
    return {
      // route params
      zone: "",
      version: "",

      // zone data
      zoneData: {},

      // filtering
      npcTypeFields: [],
      filters: [],

      // v-models
      npcNameSearch: "",
    }
  },

  watch: {
    $route(to, from) {
      this.init()
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
    reset() {
      this.npcNameSearch = ""
      this.filters = []

      this.updateQueryState()
    },

    updateQueryState: function () {
      let queryState = {};

      if (typeof this.zoneData.version !== "undefined") {
        queryState.v = this.zoneData.version
      }
      if (this.npcNameSearch !== "") {
        queryState.q = this.npcNameSearch
      }
      if (this.filters && this.filters.length > 0) {
        queryState.filters = JSON.stringify(this.filters)
      }

      console.log(queryState)

      this.$router.push(
        {
          path: ROUTE.NPCS_EDIT.replaceAll(":zone", this.zoneData.short_name),
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState: function () {

      if (this.$route.query.q !== "") {
        this.npcNameSearch = this.$route.query.q
      }

      if (this.$route.query.filters) {
        this.filters = JSON.parse(this.$route.query.filters);
      } else {
        this.filters = [];
      }
    },

    handleDbColumnFilters(filters) {
      this.filters = filters
      this.updateQueryState()
    },

    getColumnHeaderWidth(header) {
      if (header.includes("lastname")) {
        return 'min-width: 200px; '
      }

      return ''
    },

    getColumnStylingFromIndex(index) {
      let styling = '';

      if (index === 1) {
        styling += 'left: 77px; font-weight: bold;';
      }

      if (index === 0) {
        styling += 'left: 0px; font-weight: bold;'
      }

      return styling;
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
      this.loadQueryState()

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

      DbSchema.getTableColumns("npc_types").then((r) => {
        this.npcTypeFields = r
      })
    },

    async loadNpcTypes() {

      // TODO: Clean this up later
      // First pass
      // We grab NPC IDs by spawn zone / version and then do a bulk call with
      // filters as a second pass
      const api   = (new Spawn2Api(SpireApiClient.getOpenApiConfig()))
      let builder = (new SpireQueryBuilder())
      builder.where("zone", "=", this.zoneData.short_name)
      builder.where("version", "=", this.zoneData.version)
      builder.includes([
        "Spawnentries.NpcType",
      ])

      let npcTypes = [];
      let npcIds   = []
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

                  npcIds.push(spawnentry.npc_type.id)
                }

              }
            }
          }
        }

        // second pass
        const npcTypeApi = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))
        builder          = (new SpireQueryBuilder())

        if (this.filters && this.filters.length > 0) {
          this.filters.forEach((f) => {
            builder.where(f.f, f.o, f.v)
          });
        }

        console.log(this.npcNameSearch)

        if (typeof this.npcNameSearch !== "undefined" && this.npcNameSearch !== "") {
          builder.where("name", "like", this.npcNameSearch)
        }

        const rn = await npcTypeApi.getNpcTypesBulk(
          {
            body: {
              ids: npcIds
            }
          },
          {
            query: builder.get()
          }
        )
        if (rn.status === 200) {

          // sort alpha, upper case first
          rn.data = rn.data.sort((a, b) => {
            if (this.startsWithUppercase(a.name) && !this.startsWithUppercase(b.name)) {
              return -1;
            } else if (this.startsWithUppercase(b.name) && !this.startsWithUppercase(a.name)) {
              return 1;
            }
            return a.name.localeCompare(b.name);
          });

          this.npcTypes = rn.data
        }

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
