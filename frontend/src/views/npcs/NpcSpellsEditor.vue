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
                  @click="reset(); rows = []; loading = true; updateQueryState(); init()"
                >
                  <i class="fa fa-refresh mr-1"></i>
                  Reset
                </b-button>
                <b-button
                  size="sm"
                  variant="outline-warning"
                  @click="newEmote()"
                >
                  <i class="fa fa-plus mr-1"></i>
                  New
                </b-button>
              </div>
            </div>
            <!--            <div class="col-2">-->
            <!--              Event-->
            <!--              <b-form-select-->
            <!--                v-model="eventSelection"-->
            <!--                @change="subSelectedId = -1; rows = []; updateQueryState()"-->
            <!--                :options="NPC_EMOTE_EVENTS"-->
            <!--              />-->
            <!--            </div>-->
            <!--            <div class="col-2">-->
            <!--              Type-->
            <!--              <b-form-select-->
            <!--                v-model="typeSelection"-->
            <!--                @change="subSelectedId = -1; rows = []; updateQueryState()"-->
            <!--                :options="NPC_EMOTE_TYPES"-->
            <!--              />-->
            <!--            </div>-->
            <div class="col-2">
              Search
              <b-form-input
                v-model="search"
                v-on:keyup="subSelectedId = -1; rows = []; loading = true; updateQueryStateDebounce()"
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
              class="eq-table bordered eq-highlight-rows row-table emotes-table"
              style="display: table; font-size: 14px; overflow-x: scroll"
              v-if="rows && rows.length > 0"
              id="emotes-table"
            >
              <thead
                class="eq-table-floating-header"
              >
              <tr>
                <th style="width: 50px"></th>
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
                @click="selectSpellSet(e)"
              >
                <td class="text-center">
                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-success"
                    style="padding: 0px 6px;"
                    title="Edit Spell set"
                    @click="editNpcSpellSet(e)"
                  >
                    <i class="fa fa-pencil-square"></i>
                  </b-button>
                </td>
                <td class="text-center">{{ e.id }}</td>
                <td>{{ e.name }}</td>
                <td>{{ e.parent_list }}</td>
                <td>{{ getSpellCount(e) }}</td>
              </tr>
              </tbody>
            </table>


          </div>

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
      <div class="col-6 fade-in" v-if="isSubEditActive()">
        <eq-window :title="selectedSpellSet.name" class="p-2">
          <div style="max-height: 93vh; overflow-y: scroll; overflow-x: hidden">
            <npc-spell-preview
              :spells="selectedSpellSet"
            />
          </div>
        </eq-window>

        <eq-window
          class="mt-5"
          :title="'Emote ID (' + selectedSpellSet.emoteid + ') NPC(s) (' + npcs.length + ')'"
          v-if="selectedSpellSet && selectedSpellSet.emoteid && npcs && npcs.length > 0"
        >
          <div style="max-height: 46vh; overflow-y: scroll; overflow-x: hidden">
            <table
              id="npctable"
              class="eq-table eq-highlight-rows"
              style="display: table; font-size: 14px; "
            >
              <thead
                class="eq-table-floating-header"
              >
              <tr>
                <th>
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
import EqWindowSimple                      from "../../components/eq-ui/EQWindowSimple";
import EqAutoTable                         from "../../components/eq-ui/EQAutoTable";
import ContentArea                         from "../../components/layout/ContentArea";
import {NpcSpellApi}                       from "../../app/api";
import {SpireApiClient}                    from "../../app/api/spire-api-client";
import LoaderFakeProgress                  from "../../components/LoaderFakeProgress";
import {ROUTE}                             from "../../routes";
import {SpireQueryBuilder}                 from "../../app/api/spire-query-builder";
import InfoErrorBanner                     from "../../components/InfoErrorBanner";
import EqWindow                            from "../../components/eq-ui/EQWindow";
import {NPC_EMOTE_EVENTS, NPC_EMOTE_TYPES} from "../../app/constants/eq-npc-emotes";
import Tablesort                           from "../../app/utility/tablesort";
import {scrollToTarget}                    from "../../app/utility/scrollToTarget";
import {debounce}                          from "../../app/utility/debounce";
import {Npcs}                              from "../../app/npcs";
import NpcPopover                          from "../../components/NpcPopover";
import NpcSpellPreview                     from "../../components/preview/NpcSpellPreview";
import EqDebug                             from "../../components/eq-ui/EQDebug";

const NpcSpellsClient = (new NpcSpellApi(SpireApiClient.getOpenApiConfig()))

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

      // api responses
      error: "",
      notification: "",

      // selection
      eventSelection: -1,
      typeSelection: -1,
      search: "",

      // npcs that use the emote
      npcs: [],

      // constants
      NPC_EMOTE_TYPES: NPC_EMOTE_TYPES,
      NPC_EMOTE_EVENTS: NPC_EMOTE_EVENTS,

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

    editNpcSpellSet(e) {

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
          e.spawngroup.spawn_2.zone = e.spawngroup.spawn_2.zone.toLowerCase()
          if (!zones.includes(e.spawngroup.spawn_2.zone)) {
            zones.push(e.spawngroup.spawn_2.zone)
          }
        }
      }

      return zones
    },

    async loadNpcsByEmote(e) {
      this.npcs = await Npcs.listNpcsByEmoteId(e.emoteid,
        [
          "NpcSpell.NpcSpellsEntries.SpellsNew",
          "NpcFactions.NpcFactionEntries.FactionList",
          "NpcFactions",
          "NpcEmotes",
          "Merchantlists.Items",
          "Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
          "Spawnentries.Spawngroup.Spawn2"
        ])
    },

    updateQueryStateDebounce: debounce(function () {
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
      this.currentPage    = 1
      this.search         = ""
      this.error          = ""
      this.notification   = ""
      this.eventSelection = -1
      this.typeSelection  = -1
      this.subSelectedId  = -1
      this.npcs           = []
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
      if (this.eventSelection !== -1) {
        queryState.event = this.eventSelection
      }
      if (this.typeSelection !== -1) {
        queryState.type = this.typeSelection
      }
      if (this.currentPage > 0) {
        queryState.page = this.currentPage
      }
      if (this.search !== "") {
        queryState.search = this.search
      }

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
      if (this.$route.query.event >= 0) {
        this.eventSelection = parseInt(this.$route.query.event);
      }
      if (this.$route.query.type >= 0) {
        this.typeSelection = parseInt(this.$route.query.type);
      }
      if (this.$route.query.search !== "") {
        this.search = this.$route.query.search;
      }
      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
        console.log("current page", this.currentPage)
      }
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

    async newEmote() {

      let nextEmoteId = 0
      for (const e of this.rows) {
        if (e.emoteid > nextEmoteId) {
          nextEmoteId = e.emoteid
        }
      }

      try {
        const r = await NpcSpellsClient.createNpcEmote(
          {
            npcEmote: {
              emoteid: nextEmoteId + 1,
              event_: 1,
              type: 0,
              text: "You will not evade me!",
            }
          }
        )
        if (r.status === 200) {
          if (r.data.id > 0) {
            this.notification  = "New emote created!"
            this.subSelectedId = r.data.id
            this.rows          = []
            this.updateQueryState()
          }
        }
      } catch (err) {
        if (err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async save() {
      try {
        const r = await NpcSpellsClient.updateNpcEmote(
          {
            id: this.selectedSpellSet.id,
            npcEmote: this.selectedSpellSet
          }
        )
        if (r.status === 200) {
          this.notification = "Emote data saved!"
        }
      } catch (err) {
        if (err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async deleteEmote() {
      if (confirm("Are you sure you want to delete this emote (" + this.subSelectedId + ")?")) {
        try {
          const r = await NpcSpellsClient.deleteNpcEmote(
            {
              id: this.selectedSpellSet.id,
            }
          )

          if (r.status === 200) {
            this.notification = "Emote deleted!"
            this.reset()

            // convenience - we grab the last element in the list if we deleted the last element in the table
            const deletedLastElement = this.selectedSpellSet.id === this.rows[this.rows.length - 1].id
            if (deletedLastElement) {
              this.subSelectedId = this.rows[this.rows.length - 2].id
            }

            this.rows = []
            this.updateQueryState()
          }
        } catch (err) {
          console.log(err)
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }
    },

    /**
     * Initialize
     */
    async init(reset = false) {
      if (this.totalRows === 0) {
        this.totalRows = await this.getNpcSpellCount()
      }
      this.loadQueryState()

      console.log("last page", this.lastPage)
      console.log("current page", this.currentPage)

      if (this.lastPage !== this.currentPage || reset) {
        console.log("reloading content")
        this.loading      = true
        this.lastPage     = this.currentPage
        this.rows         = await this.getNpcSpells()
        this.originalRows = JSON.parse(JSON.stringify(this.rows))
      }

      setTimeout(() => {
        if (document.getElementById('emotes-table')) {
          new Tablesort(document.getElementById('emotes-table'));
        }
      }, 100)

      if (this.subSelectedId > 0) {
        for (const e of this.rows) {
          if (e.id === this.subSelectedId) {
            this.selectedSpellSet = JSON.parse(JSON.stringify(e))
            // re-hook up later
            // this.loadNpcsByEmote(this.selectedSpellSet)
          }
        }
      }

      // filters
      let rows = []
      if (this.eventSelection > -1 || this.typeSelection > -1 || this.search !== "") {
        for (const r of this.originalRows) {
          if (this.eventSelection > -1 && r.event_ !== this.eventSelection) {
            continue
          }
          if (this.typeSelection > -1 && r.type !== this.typeSelection) {
            continue;
          }
          if (this.search && this.search !== "" && r.text &&
            !r.text.toLowerCase().includes(this.search.toLowerCase()) &&
            !r.id.toString().includes(this.search) &&
            !r.emoteid.toString().includes(this.search)
          ) {
            continue;
          }

          rows.push(r)
        }

        if (rows.length > 0) {
          this.rows = rows
        }
      }

      this.loading = false

      // this.scrollToHighlighted()
    },

    async getNpcSpellCount() {
      let builder = (new SpireQueryBuilder())
        .select(["id"])
        .limit(100000)

      const response = await NpcSpellsClient.listNpcSpells(
        builder.get()
      )
      if (response.status === 200 && response.data) {
        return response.data.length
      }

      return 0
    },

    async getNpcSpells() {
      let builder = (new SpireQueryBuilder())
        .page(this.currentPage)
        .includes([
          "NpcSpellsEntries.SpellsNew",
        ])
        .limit(this.pageSize)

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
.emotes-table td {
  vertical-align: middle !important;
}
</style>
