<template>
  <content-area style="padding: 0px !important">
    <div class="row">
      <div :class="(isSubEditActive() ? 'col-4' : 'col-12')">
        <eq-window
          title="Bot Spells"
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
                autofocus
              />
            </div>

            <div class="col-2" v-if="loading">
              <div class="text-center" style="margin-top: 17px">
                Loading...
                <loader-fake-progress/>
              </div>
            </div>

          </div>

          <!-- Notification / Error -->
          <info-error-banner
            class="mr-3 ml-3"
            style="margin-top: 20px"
            :notification="notification"
            :error="error"
            :slim="true"
            @dismiss-error="error = ''"
            @dismiss-notification="notification = ''"
          />

          <div
            id="bot-spell-viewport"
            style="max-height: 80vh; overflow-y: scroll"
          >
            <table
              class="eq-table bordered eq-highlight-rows row-table bot-spell-sets-table"
              style="display: table; font-size: 14px; overflow-x: scroll"
              v-if="rows && rows.length > 0"
              id="bot-spell-sets-table"
            >
              <thead
                class="eq-table-floating-header"
              >
              <tr>
                <th style="width: 100px"></th>
                <th>ID</th>
                <th>Name</th>
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
                    class="btn-dark btn-sm"
                    style="padding: 0px 6px;"
                    title="Edit Spell Set"
                    @click="editBotSpellSet(e.id)"
                  >
                    <i class="fa fa-pencil-square"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm ml-1"
                    style="padding: 0px 6px;"
                    title="View Spell Set"
                    @click="selectSpellSet(e)"
                  >
                    <i class="fa fa-eye"></i>
                  </b-button>
                </td>
                <td class="text-center">{{ e.id }}</td>
                <td>{{ e.name }}</td>

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
      <div class="col-8 fade-in" v-if="isSubEditActive()">
        <eq-window
          v-if="selectedSpellSet && selectedSpellSet.bot_spells_entries && selectedSpellSet.bot_spells_entries.length"
          :title="`Bot Spells ID (${selectedSpellSet.id}) [${selectedSpellSet.name}] Count (${selectedSpellSet.bot_spells_entries.length})`"
          class="p-2"
        >
          <div style="max-height: 88vh; overflow-y: scroll; overflow-x: hidden">
            <bot-spell-preview
              :spells="selectedSpellSet"
            />
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
import {NpcSpellApi}       from "../../app/api";
import {SpireApi}          from "../../app/api/spire-api";
import LoaderFakeProgress  from "../../components/LoaderFakeProgress";
import {ROUTE}             from "../../routes";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import InfoErrorBanner     from "../../components/InfoErrorBanner";
import EqWindow            from "../../components/eq-ui/EQWindow";
import Tablesort           from "../../app/utility/tablesort";
import {scrollToTarget}    from "../../app/utility/scrollToTarget";
import {debounce}          from "../../app/utility/debounce";
import BotSpellPreview     from "../../components/preview/BotSpellPreview";
import EqDebug             from "../../components/eq-ui/EQDebug";
import util                from "util";

const BotSpellsClient = (new NpcSpellApi(...SpireApi.cfg()))

export default {
  name: "BotSpellsEditor",
  components: {
    EqDebug,
    BotSpellPreview,
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
      search: "",

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

    editBotSpellSet(id) {
      this.$router.push(
        {
          path: util.format(ROUTE.BOT_SPELL_EDIT, id)
        }
      ).catch(() => {
      })
    },

    getSpellCount(e) {
      return e.bot_spells_entries ? e.bot_spells_entries.length : 0
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

      console.log("[bot-spells-editor] Updating query state", queryState)

      this.$router.push(
        {
          path: ROUTE.BOT_SPELLS_EDIT,
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
        this.totalRows = await this.getBotSpellCount()
      }

      // console.log("last page", this.lastPage)
      // console.log("current page", this.currentPage)

      if (this.lastPage !== this.currentPage || reset) {
        console.log("reloading content")
        this.loading  = true
        this.lastPage = parseInt(this.currentPage)

        try {
          this.rows = await this.getBotSpells()
        } catch (err) {
          if (err.response.data.error) {
            this.error = err.response.data.error
          }
        }

        this.originalRows = JSON.parse(JSON.stringify(this.rows))
      }

      setTimeout(() => {
        if (document.getElementById('bot-spell-sets-table')) {
          new Tablesort(document.getElementById('bot-spell-sets-table'));
        }
      }, 100)

      if (this.subSelectedId > 0) {
        for (const e of this.rows) {
          if (e.id === this.subSelectedId) {
            this.selectedSpellSet = JSON.parse(JSON.stringify(e))
          }
        }
      }

      this.loading = false
    },

    isNumeric(value) {
      return /^-?\d+$/.test(value);
    },

    async getBotSpellCount() {
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

      builder.where("id", ">=", 3001)
      builder.where("id", "<=", 3016)

      // console.log("SEARCH IS (count) ", this.search)

      const response = await BotSpellsClient.listNpcSpells(
        builder.get()
      )
      if (response.status === 200 && response.data) {
        return response.data.length
      }

      return 0
    },

    async getBotSpells() {
      console.log("getbotspells", this.currentPage)

      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
        console.log("current page", this.currentPage)
      }

      let builder = (new SpireQueryBuilder())
        .page(this.currentPage)
        .includes([
          "BotSpellsEntries.SpellsNew",
        ])
        .where("id", ">=", 3001)
        .where("id", "<=", 3016)
        .limit(this.pageSize)

      if (this.search && this.search.length > 0) {
        if (this.isNumeric(this.search)) {
          builder.where("id", "=", this.search)
        } else {
          builder.where("name", "like", this.search)
        }
      }

      // console.log("SEARCH IS (data) ", this.search)

      const response = await BotSpellsClient.listNpcSpells(
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
        "bot-spell-viewport",
        'row-' + this.subSelectedId
      )
    }
  },
}
</script>

<style>
.bot-spell-sets-table td {
  vertical-align: middle !important;
}
</style>
