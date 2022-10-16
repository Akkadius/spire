<template>
  <content-area style="padding: 0px !important">
    <div class="row">
      <div :class="(isSubEditActive() ? 'col-6' : 'col-12')">
        <eq-window
          title="NPC Spells"
          class="p-0"
        >
          <div class="row minified-inputs mt-4 text-center">
            <div class="col-3 p-0 text-right">
              <div class="d-inline-block btn-group ml-4 text-right" role="group" style="margin-top: 26px">
                <b-button
                  size="sm"
                  variant="outline-warning"
                  @click="zeroState();"
                >
                  <i class="fa fa-refresh mr-1"></i>
                  Reset
                </b-button>
              </div>
            </div>

            <div class="col-4">
              Search
              <b-form-input
                v-model="search"
                v-on:keyup="doSearch()"
                placeholder="Search by text or id"
              />
            </div>

            <div class="col-2" v-if="loading">
              <div class="text-center" style="margin-top: 17px">
                Loading...
                <loader-fake-progress/>
              </div>
            </div>
          </div>

          <div
            id="npc-spell-viewport"
            style="max-height: 80vh; overflow-y: scroll"
          >
            <table
              class="eq-table bordered eq-highlight-rows row-table npc-spell-sets-table"
              style="display: table; font-size: 14px; overflow-x: scroll"
              v-if="rows && rows.length > 0"
              id="npc-spell-sets-table"
            >
              <thead
                class="eq-table-floating-header"
              >
              <tr>
                <th style="width: 100px"></th>
                <th>ID</th>
                <th>Name</th>
                <th>Parent List ID</th>
                <th>Spell Count</th>
              </tr>
              </thead>
              <tbody>
              <tr
                :id="'row-' + e.id"
                v-for="(e, index) in rows"
                :key="e.id"
                :class="isRowSelected(e) ? 'pulsate-highlight-white' : ''"
              >
                <td class="text-center pl-0 pr-0">
                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-success"
                    style="padding: 0px 6px;"
                    title="Edit Spell Set"
                    @click="editNpcSpellSet(e.id)"
                  >
                    <i class="fa fa-pencil-square"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn btn-dark btn-sm btn-outline-danger ml-1 btn-primary"
                    style="padding: 0px 6px;"
                    title="Delete spell entry"
                    @click="deleteNpcSpellSet(e)"
                  >
                    <i class="fa fa-trash"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-white ml-1"
                    style="padding: 0px 6px;"
                    title="View Spell Set"
                    @click="selectSpellSet(e)"
                  >
                    <i class="fa fa-eye"></i>
                  </b-button>
                </td>
                <td class="text-center">{{ e.id }}</td>
                <td>{{ e.name }}</td>
                <td>
                  {{ e.parent_list }}

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-success ml-1"
                    style="padding: 0px 6px;"
                    title="Edit Parent Spell Set"
                    @click="editNpcSpellSet(e.parent_list)"
                    v-if="e.parent_list > 0"
                  >
                    <i class="fa fa-pencil-square"></i>
                  </b-button>
                </td>
                <td>{{ getSpellCount(e) }}</td>
              </tr>
              </tbody>
            </table>


          </div>

          <!-- Pagination -->
          <div class="row text-center justify-content-center">
            <div class="col-12 text-center mt-3">
              <b-pagination
                :key="currentPage"
                class="mb-3"
                v-model="currentPage"
                :total-rows="totalRows"
                :hide-ellipsis="true"
                :per-page="pageSize"
                @change="paginate"
              />
            </div>
          </div>

        </eq-window>
      </div>

      <!-- Preview Pane -->
      <div class="col-6 fade-in" v-if="isSubEditActive()">
        <eq-window
          v-if="selectedSpellSet && selectedSpellSet.npc_spells_entries && selectedSpellSet.npc_spells_entries.length"
          :title="`NPC Spells ID (${selectedSpellSet.id}) [${selectedSpellSet.name}] Count (${selectedSpellSet.npc_spells_entries.length})`"
          class="p-2"
        >
          <div style="max-height: 44vh; overflow-y: scroll; overflow-x: hidden">
            <npc-spell-preview
              :spells="selectedSpellSet"
            />
          </div>
        </eq-window>

        <!-- Show NPC(s) that use this spell set -->
        <eq-window
          class="mt-5 p-0"
          :title="`NPC Spells Set ID (${selectedSpellSet.id}) (${selectedSpellSet.name}) NPC(s) (${npcs.length}) ` + (npcs.length === 100 ? '(Max 100)' : '')"
          v-if="selectedSpellSet && selectedSpellSet.id && npcs && npcs.length > 0"
        >
          <div style="max-height: 45vh; overflow-y: scroll; overflow-x: hidden">
            <table
              id="npctable"
              class="eq-table eq-highlight-rows"
              style="display: table; font-size: 14px; "
            >
              <thead
                class="eq-table-floating-header"
              >
              <tr>
                <th class="text-center">
                  NPC
                </th>
                <th>
                  Zone(s)
                </th>
              </tr>
              </thead>
              <tbody>
              <tr
                :id="'npc-' + n.short_name"
                v-for="(n, index) in npcs"
                :key="n.id"
              >
                <td style="position: relative">
                  <npc-popover
                    :npc="n"
                  />
                </td>
                <td style="vertical-align: middle">
                  {{ getZonesByNpc(n).join(",") }}
                </td>

              </tr>
              </tbody>
            </table>
          </div>
        </eq-window>

      </div>
    </div>

  </content-area>
</template>

<script>
import EqWindowSimple      from "../../components/eq-ui/EQWindowSimple";
import EqAutoTable         from "../../components/eq-ui/EQAutoTable";
import ContentArea         from "../../components/layout/ContentArea";
import {NpcSpellApi}      from "../../app/api";
import {SpireApi}         from "../../app/api/spire-api";
import LoaderFakeProgress from "../../components/LoaderFakeProgress";
import {ROUTE}             from "../../routes";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import InfoErrorBanner     from "../../components/InfoErrorBanner";
import EqWindow            from "../../components/eq-ui/EQWindow";
import Tablesort           from "../../app/utility/tablesort";
import {scrollToTarget}    from "../../app/utility/scrollToTarget";
import {debounce}          from "../../app/utility/debounce";
import {Npcs}              from "../../app/npcs";
import NpcPopover          from "../../components/NpcPopover";
import NpcSpellPreview     from "../../components/preview/NpcSpellPreview";
import EqDebug             from "../../components/eq-ui/EQDebug";
import util                from "util";
import {NpcSpellsEntryApi} from "../../app/api/api/npc-spells-entry-api";

const NpcSpellsClient = (new NpcSpellApi(...SpireApi.cfg()))

export default {
  name: "NpcSpellsEditor",
  components: {
    EqDebug,
    NpcSpellPreview,
    NpcPopover,
    EqWindow,
    InfoErrorBanner,
    LoaderFakeProgress: LoaderFakeProgress,
    ContentArea,
    EqAutoTable,
    EqWindowSimple,
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {

      // pagination (all)
      lastPage: 0, // this keeps track of last page in state
      currentPage: 1,
      totalRows: 0,
      pageSize: 25,

      // table
      rows: [],
      filterColumns: [],

      // for the sub selector pane on the right
      subSelectedId: -1,
      selectedSpellSet: {},
      parentList: [],

      // api responses
      error: "",
      notification: "",

      // selection
      search: "",

      // npcs that use the emote
      npcs: [],

      lastSelectedTime: Date.now(),

      loading: false, // are we loading or not
    }
  },

  watch: {
    '$route'() {
      console.log("route trigger")
      this.reset()
      this.init()
    },
  },

  methods: {

    zeroState() {
      this.reset()
      // lastPage: 0, // this keeps track of last page in state
      //   currentPage: 1,
      this.lastPage = 0;

      this.loading = true;
      this.updateQueryState();
      this.init()
    },

    editNpcSpellSet(id) {
      this.$router.push(
        {
          path: util.format(ROUTE.NPC_SPELL_EDIT, id)
        }
      ).catch(() => {
      })
    },

    async deleteNpcSpellSet(e) {
      const entriesCount = (e && e.npc_spells_entries && e.npc_spells_entries.length ? e.npc_spells_entries.length : 0)

      if (confirm(`Are you sure you want to delete this NPC spell set and all of its entries? \n\nEntries (${entriesCount})`)) {
        try {
          const r = await (new NpcSpellApi(...SpireApi.cfg()))
            .deleteNpcSpell({ id: e.id })

          if (r.status === 200 && e.npc_spells_entries) {
            // delete every entry individually (for now)
            for (const row of e.npc_spells_entries) {
              await (new NpcSpellsEntryApi(...SpireApi.cfg()))
                .deleteNpcSpellsEntry({ id: row.id })
            }
          }

          this.reset()
          // also resetting...
          this.lastPage  = 0;
          this.rows      = []
          this.totalRows = 0
          // reload
          this.init()

        } catch (err) {
          if (err && err.response && err.response.data.error) {
            if (err.response && err.response.data && err.response.data.error) {
              this.error = err.response.data.error
            }
          } else {
            this.error = err
          }
        }
      }
    },

    getSpellCount(e) {
      return e.npc_spells_entries ? e.npc_spells_entries.length : 0
    },

    paginate() {
      // models aren't quite updated when we trigger this so queue the pagination
      setTimeout(() => {
        console.log("We're paginating")
        console.log(this.currentPage)
        console.log(this.totalRows)
        this.updateQueryState()
      }, 100)
    },

    getZonesByNpc(n) {
      let zones = []
      if (n.spawnentries) {
        for (let e of n.spawnentries) {
          if (e.spawngroup && e.spawngroup.spawn_2 && e.spawngroup.spawn_2.zone) {
            e.spawngroup.spawn_2.zone = e.spawngroup.spawn_2.zone.toLowerCase()
            if (!zones.includes(e.spawngroup.spawn_2.zone)) {
              zones.push(e.spawngroup.spawn_2.zone)
            }
          }
        }
      }

      return zones
    },

    async loadNpcsBySpellSet(npcSpellsId) {
      this.npcs = await Npcs.listNpcsByNpcSpellsId(
        npcSpellsId,
        [
          "Spawnentries.Spawngroup.Spawn2",
        ]
      )

      // load extra data if we don't have too many NPCs
      if (this.npcs && this.npcs.length < 50) {
        this.npcs = await Npcs.listNpcsByNpcSpellsId(
          npcSpellsId,
          [...["Spawnentries.Spawngroup.Spawn2"], ...Npcs.getBaseNpcRelationships()]
        )
      }
    },

    doSearch: debounce(function () {
      this.subSelectedId = -1;
      this.rows          = [];
      this.loading       = true;
      this.lastPage      = 0;
      this.updateQueryState()
    }, 600),

    getColumnWidth(field) {
      if (field === 'event_') {
        return 'width: 130px;'
      }
      if (field === 'type') {
        return 'width: 140px;'
      }

      return ''
    },

    /**
     * Resets
     */
    reset() {
      this.currentPage   = 1
      this.search        = ""
      this.error         = ""
      this.notification  = ""
      this.subSelectedId = -1
      this.npcs          = []
    },
    resetSelections() {
      this.subSelectedId = -1
    },

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.subSelectedId !== -1) {
        queryState.selectedId = this.subSelectedId
      }
      if (this.currentPage > 0) {
        queryState.page = this.currentPage
      }
      if (this.search !== "") {
        queryState.search = this.search
      }

      console.log("[npc-spells-editor] Updating query state", queryState)

      this.$router.push(
        {
          path: ROUTE.NPC_SPELLS_EDIT,
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      console.log("loading query state")
      if (this.$route.query.selectedId >= 0) {
        this.subSelectedId = parseInt(this.$route.query.selectedId);
      }

      if (this.$route.query.search !== "") {
        this.search = this.$route.query.search;
      }
      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
        console.log("current page", this.currentPage)
      }

      console.log("query params", this.$route.query)
    },

    /**
     * Sub editor selection
     */
    selectSpellSet(e) {
      this.subSelectedId = e.id
      this.updateQueryState()
    },

    isSubEditActive() {
      return this.subSelectedId >= 0
    },

    isRowSelected(e) {
      return e.id === this.subSelectedId
    },

    /**
     * Initialize
     */
    async init(reset = false) {
      this.loadQueryState()

      if (this.totalRows === 0) {
        this.totalRows = await this.getNpcSpellCount()
      }

      // console.log("last page", this.lastPage)
      // console.log("current page", this.currentPage)

      if (this.lastPage !== this.currentPage || reset) {
        console.log("reloading content")
        this.loading      = true
        this.lastPage     = parseInt(this.currentPage)
        this.rows         = await this.getNpcSpells()
        this.originalRows = JSON.parse(JSON.stringify(this.rows))
      }

      setTimeout(() => {
        if (document.getElementById('npc-spell-sets-table')) {
          new Tablesort(document.getElementById('npc-spell-sets-table'));
        }
      }, 100)

      if (this.subSelectedId > 0) {
        for (const e of this.rows) {
          if (e.id === this.subSelectedId) {
            this.selectedSpellSet = JSON.parse(JSON.stringify(e))
            this.loadNpcsBySpellSet(this.subSelectedId)
            this.parentList = []
          }
        }
      }

      this.loading = false
    },

    isNumeric(value) {
      return /^-?\d+$/.test(value);
    },

    async getNpcSpellCount() {
      let builder = (new SpireQueryBuilder())
        .select(["id"])
        .limit(100000)

      if (this.search && this.search.length > 0) {
        if (this.isNumeric(this.search)) {
          builder.where("id", "=", this.search)
        } else {
          builder.where("name", "like", this.search)
        }
      }

      // console.log("SEARCH IS (count) ", this.search)

      const response = await NpcSpellsClient.listNpcSpells(
        builder.get()
      )
      if (response.status === 200 && response.data) {
        return response.data.length
      }

      return 0
    },

    async getNpcSpells() {
      console.log("getnpcspells", this.currentPage)

      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
        console.log("current page", this.currentPage)
      }

      let builder = (new SpireQueryBuilder())
        .page(this.currentPage)
        .includes([
          "NpcSpellsEntries.SpellsNew",
        ])
        .limit(this.pageSize)

      if (this.search && this.search.length > 0) {
        if (this.isNumeric(this.search)) {
          builder.where("id", "=", this.search)
        } else {
          builder.where("name", "like", this.search)
        }
      }

      // console.log("SEARCH IS (data) ", this.search)

      const response = await NpcSpellsClient.listNpcSpells(
        builder.get()
      )
      if (response.status === 200 && response.data) {
        return response.data
      }
    },

  },
  created() {
    this.originalRows = []
  },
  async mounted() {
    await this.init()

    if (this.subSelectedId > 0) {
      scrollToTarget(
        "npc-spell-viewport",
        'row-' + this.subSelectedId
      )
    }
  },
}
</script>

<style>
.npc-spell-sets-table td {
  vertical-align: middle !important;
}
</style>
