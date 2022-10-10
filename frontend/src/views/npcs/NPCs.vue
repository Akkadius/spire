<template>
  <div>
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
          id="npcs-table-window"
          v-if="npcTypes"
          class="p-0"
        >
          <div
            id="npcs-table-container"
            style="overflow-x: scroll; height: 85vh"
          >
            <table
              id="npcs-table"
              class="eq-table eq-highlight-rows bordered"
              style="font-size: 14px; "
              v-if="npcTypes && npcTypes.length > 0"
            >
              <thead class="eq-table-floating-header">
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
                v-for="(row, index) in npcTypes"
                :id="'npc-' + row.id"
                :key="index"
              >
                <td
                  :style="' text-align: center; ' + ([0, 1].includes(colIndex) ? ' position: sticky; z-index: 999; background-color: rgba(25,31,41, .95);' + getColumnStylingFromIndex(colIndex): '')"
                  v-for="(key, colIndex) in Object.keys(row)"
                  v-if="doesRowColumnHaveObjects(row, key)"
                >
                  <npc-popover
                    v-if="key === 'name'"
                    :show-image="false"
                    :show-label="false"
                    :npc="row"
                  >
                    {{ row[key] }}
                  </npc-popover>

                  <span v-if="key !== 'name'">{{ row[key] }}</span>

                  <!-- Set all values preview -->
                  <span
                    v-if="isPreviewValueChangeable(row[key], previewValue) && previewField === key && row[key] !== getTypedField(previewValue)"
                    style="color: yellow"
                    class="ml-1"
                  >-> {{
                      previewValue && isNumeric(previewValue) ? previewValue.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") : previewValue
                    }}</span>

                  <!-- Min / Max -->
                  <span
                    v-if="isPreviewValueChangeable(row[key], previewMinMaxData[row.id]) && previewField === key && row[key] !== getTypedField(previewMinMaxData[row.id]) && previewMinMaxData[row.id]"
                    style="color: yellow"
                    class="ml-1"
                  >-> {{
                      previewMinMaxData[row.id] && isNumeric(previewMinMaxData[row.id]) ? previewMinMaxData[row.id].toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") : previewMinMaxData[row.id]
                    }}</span>

                  <!-- Percentage -->
                  <span
                    v-if="isPreviewValueChangeable(row[key], previewPercentageData[row.id]) && previewField === key && row[key] !== getTypedField(previewPercentageData[row.id]) && previewPercentageData[row.id]"
                    style="color: yellow"
                    class="ml-1"
                  >-> {{
                      previewPercentageData[row.id] && isNumeric(previewPercentageData[row.id]) ? previewPercentageData[row.id].toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") : previewPercentageData[row.id]
                    }}</span>

                </td>
              </tr>
              </tbody>
            </table>
          </div>

        </eq-window>
      </div>

      <div class="col-5" v-if="isAnySelectorActive()">
        <npcs-bulk-editor
          @field-selected="scrollToColumn($event)"
          @set-values-preview="handleSetValuesPreview($event)"
          @set-values-commit="handleSetValuesCommit($event)"
          @set-min-max-values-preview="handleMinMaxSetValuesPreview($event)"
          @set-min-max-values-commit="handleSetValuesCommit($event)"
          @set-percentage-values-preview="handlePercentageSetValuesPreview($event)"
          @set-percentage-values-commit="handleSetValuesCommit($event)"
          :edit-feedback="bulkEditFeedback"
          v-if="selectorActive['bulk-editor']"
        />

        <!--        <eq-window title="Test!" v-if="selectorActive['bulk-editor']">-->
        <!--          Test!-->
        <!--        </eq-window>-->
      </div>
    </div>

  </div>

</template>

<script>
import EqWindow                from "../../components/eq-ui/EQWindow";
import ContentArea             from "../../components/layout/ContentArea";
import {Navbar}                from "../../app/navbar";
import {Zones}                 from "../../app/zones";
import {NpcTypeApi, Spawn2Api} from "../../app/api";
import {SpireApi}              from "../../app/api/spire-api";
import {SpireQueryBuilder}     from "../../app/api/spire-query-builder";
import Tablesort               from "@/app/utility/tablesort.js";
import DbColumnFilter          from "../../components/DbColumnFilter";
import {DbSchema}              from "../../app/db-schema";
import {ROUTE}                 from "../../routes";
import {EditFormFieldUtil} from "../../app/forms/edit-form-field-util";
import NpcsBulkEditor      from "./components/NpcsBulkEditor";
import util                from "util";
import NpcPopover              from "../../components/NpcPopover";
import {Npcs}                  from "../../app/npcs";

export default {
  name: "NPCs",
  components: { NpcPopover, NpcsBulkEditor, DbColumnFilter, ContentArea, EqWindow },
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
      bulkEditFeedback: [],

      // preview min / max
      previewMinMaxData: {},

      // preview percentage
      previewPercentageData: {},
    }
  },

  watch: {
    $route(to, from) {
      this.init()
    }
  },

  beforeDestroy() {
    Navbar.expand()

    if (this.interval) {
      clearInterval(this.interval)
    }
  },

  mounted() {
    this.init()
  },

  created() {

    // data
    this.npcTypes = []

    // background
    this.backgroundImages  = []
    this.currentImageIndex = 0

    // cycle background images
    this.interval = setInterval(this.setBackgroundImage, 10 * 1000)
  },

  methods: {
    getTypedField(value) {
      if (this.isFloat(value)) {
        return parseFloat(value);
      } else if (this.isNumeric(value)) {
        return parseInt(value);
      }

      return value
    },

    isPreviewValueChangeable(fieldValue, previewValue) {
      if (this.isFloat(fieldValue) && previewValue !== '') {
        return true;
      } else if (this.isNumeric(fieldValue) && previewValue !== '') {
        return true;
      } else if ((!this.isNumeric(fieldValue) && !this.isFloat(fieldValue)) && previewValue !== '') {
        return true;
      }

      return false
    },

    previewStyles(header) {
      if (this.previewField === header) {
        return 'padding-left: 75px !important; padding-right: 75px !important; '
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
      let editFeedback = []
      this.scrollToColumn(e.field)

      for (let n of this.npcTypes) {
        let newValue = e.value

        // when min / max is passed
        if (e.min) {
          newValue = this.previewMinMaxData[n.id]
        }

        // when percentage is passed
        if (e.percentage) {
          newValue = this.previewPercentageData[n.id]
        }

        editFeedback.push(
          `NPC ID (${n.id}) field [${e.field}] has changed from [${n[e.field]}] to [${newValue}]`
        )

        n[e.field] = newValue

        // float
        if (this.isFloat(newValue)) {
          n[e.field] = parseFloat(newValue)
        }
        // integer
        else if (this.isNumeric(newValue)) {
          n[e.field] = parseInt(newValue)
        }

        await Npcs.updateNpc(n.id, n);

        // scroll to table row entry as we are editing entries
        const container = document.getElementById("npcs-table-container");
        const target    = document.getElementById(util.format("npc-%s", n.id))
        if (container && target) {
          container.scrollTop = container.scrollTop + target.getBoundingClientRect().top - 200;
        }

        this.$forceUpdate()
      }

      this.bulkEditFeedback = editFeedback

      // reset
      this.previewField = ""
      this.previewValue = ""

      // this.reset()

      // reset scroll to 0
      const container     = document.getElementById("npcs-table-container");
      container.scrollTop = 0

      this.updateQueryState()

      this.bulkEdit()
    },

    resetPulseHighlights() {
      // strip existing columns with the header
      for (let e of document.getElementsByClassName("pulsate-highlight")) {
        e.classList.remove("pulsate-highlight")
      }
    },

    handleSetValuesPreview(e) {
      // reset min max
      this.previewMinMaxData     = {}
      this.previewPercentageData = {}

      this.previewField = e.field
      this.previewValue = e.value
    },
    getRandomArbitrary(min, max) {
      return Math.random() * (max - min) + min;
    },

    handleMinMaxSetValuesPreview(e) {
      // reset other previews
      this.previewValue          = ''
      this.previewPercentageData = {}

      const field = e.field
      const max   = e.max
      const min   = e.min

      this.previewField     = field
      let previewMinMaxData = {}
      for (let n of this.npcTypes) {
        previewMinMaxData[n.id] = Math.round(this.getRandomArbitrary(min, max))
      }

      this.previewMinMaxData = previewMinMaxData
    },
    handlePercentageSetValuesPreview(e) {
      // reset other previews
      this.previewValue      = ''
      this.previewMinMaxData = {}

      const field      = e.field
      const percentage = e.percentage

      this.previewField         = field
      let previewPercentageData = {}
      for (let n of this.npcTypes) {
        previewPercentageData[n.id] = Math.round(n[field] * percentage)
      }

      this.previewPercentageData = previewPercentageData
    },
    scrollToColumn(e) {
      const container = document.getElementById("npcs-table-container");
      const target    = document.getElementById(util.format("column-%s", e))

      if (container && target) {
        container.scrollLeft = container.offsetLeft + target.offsetLeft - 400;

        this.resetPulseHighlights()

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

      // reset
      this.previewField = ""
      this.previewValue = ""

      this.resetPulseHighlights()

      this.resetPreviewComponents()
      this.updateQueryState()

      // reset scroll to 0
      const container      = document.getElementById("npcs-table-container");
      container.scrollTop  = 0
      container.scrollLeft = 0
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

      this.loadBackgroundImages().then(() => {
        this.setBackgroundImage()
      })
    },

    async loadNpcTypes() {

      // TODO: Clean this up later
      // First pass
      // We grab NPC IDs by spawn zone / version and then do a bulk call with
      // filters as a second pass
      const api   = (new Spawn2Api(...SpireApi.cfg()))
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
        const npcTypeApi = (new NpcTypeApi(...SpireApi.cfg()))
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

    /**
     * Selectors logic
     */
    isAnySelectorActive() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        if (this.selectorActive[k]) {
          return true;
        }
      }
    },
    shouldReset() {
      return (Date.now() - this.lastResetTime) > MILLISECONDS_BEFORE_WINDOW_RESET
    },
    previewNPC(force = false) {
      if ((this.shouldReset() && !this.previewMain) || force) {
        this.resetPreviewComponents()
        this.previewMain = true
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
      this.previewMain        = false;
      this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()

      EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    },


    /**
     * Image slider background
     */
    shuffle(array) {
      let currentIndex = array.length, randomIndex;

      // While there remain elements to shuffle.
      while (currentIndex !== 0) {

        // Pick a remaining element.
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex--;

        // And swap it with the current element.
        [array[currentIndex], array[randomIndex]] = [
          array[randomIndex], array[currentIndex]];
      }

      return array;
    },

    async loadBackgroundImages() {
      console.log("[EQZoneCardPreview] loadBackgroundImages")

      document.body.style.setProperty("--zone-background", "none");
      document.body.style.setProperty("--zone-background-size", "auto");

      // get zone wallpaper
      await SpireApi.v1().get('/assets/zone-images/' + encodeURIComponent(this.zoneData.long_name)).then((r) => {
        if (r.status === 200) {
          this.backgroundImages = this.shuffle(r.data.images)
        }
      })
    },
    setBackgroundImage() {
      if (this.backgroundImages && this.backgroundImages.length > 0) {
        const image = this.backgroundImages[this.currentImageIndex];
        // console.log("IMAGE ", image)

        // console.log(
        //   "[EQZoneCardPreview] loadBackgroundImages Playing index [%s] out of [%s]",
        //   this.currentImageIndex,
        //   this.backgroundImages.length
        // )

        if (image.length > 0) {
          let img     = new Image();
          img.src     = image;
          img.onload  = () => {
            // document.body.style.setProperty("--image", "url(" + image + ")");
            document.body.style.setProperty("--zone-background", "url(" + image + ")");
            document.body.style.setProperty("--zone-background-size", "cover");

            // increment
            this.currentImageIndex++;

            // reset if rollover
            if (this.currentImageIndex >= this.backgroundImages.length) {
              // console.log("[EQZoneCardPreview] loadBackgroundImages resetting")
              this.currentImageIndex = 0;
            }
          }
          img.onerror = () => {
            // console.log(
            //   "[EQZoneCardPreview] loadBackgroundImages Failed to load index [%s] out of [%s]",
            //   this.currentImageIndex,
            //   this.backgroundImages.length
            // )

            this.currentImageIndex++
            this.setBackgroundImage()
          }

        }
      }
    },
  }
}
</script>

<style>
:root {
  --zone-background-size: auto;
  --zone-background: none;
}

#npcs-table-window::before {
  content: "";

  background-size: var(--zone-background-size) !important;
  background-repeat: no-repeat !important;
  background-attachment: fixed !important;
  background-position: center !important;

  z-index: -99999;

  top: 0;
  right: 0;
  bottom: 0;
  left: 0;

  background: var(--zone-background);
  opacity: .2;

  --webkit-transition: background-image 1s ease-in-out;
  transition: background-image 1s ease-in-out;
}
</style>
