<template>
  <content-area>

    <div class="row">
      <div :class="(isAnySelectorActive() ? 'col-7' : 'col-12')">
        <eq-window
          v-if="zoneData"
          :title="`${zoneData.long_name} Short Name (${zoneData.short_name}) Version (${zoneData.version}) NPC(s) (${npcTypes.length})`"
        >
          <div class="row">
            <div :class="(isAnySelectorActive() ? 'col-2' : 'col-1') + 'text-right'">
              <button
                class='btn btn-outline-warning btn-sm mt-1'
                @click="reset"
              >
                <i class="fa fa-refresh"></i> Reset
              </button>
              <button
                class='btn btn-outline-warning btn-sm mt-1 ml-3'
                @click="bulkEdit()"
              >
                <i class="fa fa-edit"></i> Bulk Edit
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
          id="npcs-table-container"
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
                :id="'column-' + header"
                :style="previewStyles(header) + 'text-align: center; ' + getColumnHeaderWidth(header) + '' + ([0, 1].includes(index) ? ' position: sticky; z-index: 9999; background-color: rgba(25,31,41, 1); ' + getColumnStylingFromIndex(index) : '')"
              >{{ header }}
              </th>
            </tr>
            </thead>
            <tbody>
            <tr
              v-for="(row, index) in npcTypes" :key="index"
            >
              <td
                :style="' text-align: center; ' + ([0, 1].includes(colIndex) ? ' position: sticky; z-index: 999; background-color: rgba(25,31,41, .6);' + getColumnStylingFromIndex(colIndex): '')"
                v-for="(key, colIndex) in Object.keys(row)"
                v-if="doesRowColumnHaveObjects(row, key)"
              >
                {{ row[key] }}

                <span v-if="previewField === key" style="color: yellow">{{previewValue}}</span>
              </td>
            </tr>
            </tbody>
          </table>

        </eq-window>
      </div>

      <div class="col-5" v-if="isAnySelectorActive()">
        <npcs-bulk-editor
          @field-selected="scrollToColumn($event)"
          @set-values-preview="handleSetValuesPreview($event)"
          @set-values-commit="handleSetValuesCommit($event)"
          v-if="selectorActive['bulk-editor']"
        />

        <!--        <eq-window title="Test!" v-if="selectorActive['bulk-editor']">-->
        <!--          Test!-->
        <!--        </eq-window>-->
      </div>
    </div>

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
import {EditFormFieldUtil}     from "../../app/forms/edit-form-field-util";
import NpcsBulkEditor          from "./components/NpcsBulkEditor";
import util                    from "util";

export default {
  name: "NPCs",
  components: { NpcsBulkEditor, DbColumnFilter, ContentArea, EqWindow },
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

      // preview / selectors
      selectorActive: {},

      // preview value
      previewField: "",
      previewValue: "",
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
    previewStyles(header) {
      if (this.previewField === header) {
        return 'padding-left: 30px !important; padding-right: 30px !important; '
      }

      return ''
    },

    isFloat(value) {
      return typeof value === 'number' &&
        !Number.isNaN(value) &&
        !Number.isInteger(value);
    },

    isNumeric(value) {
      return /^-?\d+$/.test(value);
    },

    async handleSetValuesCommit(e) {

      console.log("event", e)


      for (let n of this.npcTypes) {
        console.log(n)

        const npcTypeApi = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))

        n[e.field]       = e.value


        // float
        if (this.isFloat(e.value)){
          console.log("FLOAT")
          n[e.field] = parseFloat(e.value)
        }
        // integer
        else if (this.isNumeric(e.value)) {
          console.log("INTEGER")
          n[e.field] = parseInt(e.value)
        }

        await npcTypeApi.updateNpcType({
          id: n.id,
          npcType: n
        })

        this.$forceUpdate()
      }

      // reset
      this.previewField = ""
      this.previewValue = ""

      // strip existing columns with the header
      for (let e of document.getElementsByClassName("pulsate-highlight")) {
        e.classList.remove("pulsate-highlight")
      }

      this.reset()
      this.updateQueryState()
    },

    handleSetValuesPreview(e) {
      this.previewField = e.field
      this.previewValue = e.value
    },

    scrollToColumn(e) {
      const container = document.getElementById("npcs-table-container");
      const target    = document.getElementById(util.format("column-%s", e))

      if (container && target) {
        container.scrollLeft = container.offsetLeft + target.offsetLeft - 400;

        // strip existing columns with the header
        for (let e of document.getElementsByClassName("pulsate-highlight")) {
          e.classList.remove("pulsate-highlight")
        }

        // add to the target column
        target.classList.add("pulsate-highlight");
      }
    },

    bulkEdit() {
      this.setSelectorActive('bulk-editor')
    },

    reset() {
      this.npcNameSearch = ""
      this.filters       = []

      this.resetPreviewComponents()
      this.updateQueryState()

      // reset scroll to 0
      const container = document.getElementById("npcs-table-container");
      container.offsetLeft = 0
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

    isAnySelectorActive() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        if (this.selectorActive[k]) {
          return true;
        }
      }
    },

    resetPreviewComponents() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }

      EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
    },
    setSelectorActive(selector) {
      this.resetPreviewComponents()
      this.previewTaskActive        = false;
      this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()

      EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    }
  }
}
</script>

<style scoped>

</style>
