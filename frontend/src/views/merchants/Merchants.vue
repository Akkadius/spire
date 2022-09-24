<template>
  <div>
    <!-- Browse View -->
    <div class="row">
      <div :class="(Object.keys(activeMerchantNpc).length > 0 || Object.keys(activeMerchantList).length > 0 ? 'col-7' : 'col-12')">

        <!-- form inputs -->
        <eq-window title="Merchant Editor">
          <div class="row">
            <div class="col-lg-3">
              <input
                type="text"
                class="form-control ml-2"
                placeholder="Merchants by Name"
                v-model="search"
                @keyup.enter="zoneSelection = 0; searchItemName = ''; updateQueryState()"
              >
            </div>

            <div class="col-lg-2">
              <input
                type="text"
                class="form-control ml-2"
                placeholder="Item Name or ID"
                v-model="searchItemName"
                @keyup.enter="zoneSelection = 0; search = ''; updateQueryState()"
              >
            </div>

            <div class="col-lg-2">
              <select
                class="form-control"
                v-model="zoneSelection"
                @change="search = ''; searchItemName = ''; updateQueryState()"
              >
                <option value="0">-- Select --</option>
                <option v-for="z in zones" v-bind:value="{z: z.short_name, v: z.version}">
                  {{ z.short_name }} ({{ z.version }}) ({{ z.zoneidnumber }}) {{ z.long_name }}
                </option>
              </select>
            </div>

            <div class="col-lg-5 text-center p-0 mt-1">
              <div class="btn-group" role="group" aria-label="Basic example">
                <b-button title="Search" @click="updateQueryState()" size="sm" variant="outline-warning">
                  <i class="fa fa-search"></i> Search
                </b-button>
                <b-button
                  title="Show all Merchant Tables"
                  @click="reset(); showAll = true; updateQueryState()"
                  size="sm"
                  variant="outline-warning"
                >
                  <i class="ra ra-emerald"></i> All
                </b-button>
                <b-button title="Reset" @click="reset(); updateQueryState()" size="sm" variant="outline-danger">
                  <i class="fa fa-eraser"></i> Reset
                </b-button>
                <b-button
                  title="Create New Merchant Table"
                  @click="createNewMerchant();"
                  size="sm"
                  variant="outline-success"
                >
                  <i class="ra ra-emerald"></i> New Merchant
                </b-button>
              </div>
            </div>
          </div>
        </eq-window>

        <!-- Loading -->
        <eq-window v-if="loading">
          <div class="text-center">
            Loading
            <loader-fake-progress class="mt-3"/>
          </div>
        </eq-window>

        <!-- Not found -->
        <eq-window v-if="isNavigated() && !loading && ((ml && ml.length === 0) && (merchantLists && merchantLists.length === 0))">
          No merchants found...
        </eq-window>

        <!-- List Merchants by NPC -->
        <eq-window
          v-if="!loading && ml && ml.length > 0"
          class="p-2 mt-5"
          :title="'Merchants (' + ml.length + ')'"
        >
          <div style="overflow-y: scroll; max-height: 83vh">
            <table
              class="eq-table bordered eq-highlight-rows"
              style="font-size: 14px; "
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th style="width: 80px"></th>
                <th class="text-center" style="width: 80px">Merchant ID</th>
                <th>Merchant (NPC)</th>
              </tr>
              </thead>
              <tbody>
              <tr
                v-for="(n, index) in ml"
                :id="'ml-' + n.id"
                :key="'ml-' + n.id + '-' + n.merchant_id"
                @mouseover="activeMerchantNpc = n"
              >
                <td
                  class="text-center"
                >
                  <b-button
                    v-if="isSelector"
                    class="btn-dark btn-sm btn-outline-warning mr-3"
                    title="Select Merchant List"
                    @click="selectMerchantList(n.merchant_id);"
                  >
                    <i class="fa fa-arrow-left"></i>
                  </b-button>

                  <b-button
                    @click="editMerchantList(n.merchant_id)"
                    size="sm"
                    title="Edit Merchant List"
                    class="btn btn-dark btn-outline-success mr-2"
                  >
                    <i class="fa fa-pencil-square"></i>
                  </b-button>

                </td>
                <td
                  class="text-center"
                  style="vertical-align: middle"
                >{{ n.id }}
                </td>
                <td class="text-left">
                  <npc-popover
                    :popover-enabled="false"
                    :limit-entries="25"
                    :no-stats="true"
                    :npc="n"
                  />
                </td>
              </tr>
              </tbody>
            </table>
          </div>
        </eq-window>

        <!-- Merchantlist (Raw) top level data -->
        <eq-window
          v-if="!loading && merchantLists && merchantLists.length > 0"
          class="p-2 mt-5"
          :title="'Merchants (' + totalRows + ')'"
        >
          <div style="overflow-y: scroll; max-height: 77vh">

            <table
              class="eq-table bordered eq-highlight-rows"
              style="font-size: 14px; "
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th class="text-center"></th>
                <th class="text-center">Merchant ID</th>
                <th class="text-center">Item Count</th>
                <th class="text-center">NPC(s) Linked to Merchant Table</th>
              </tr>
              </thead>
              <tbody>
              <tr
                v-for="(m, index) in merchantLists"
                :id="'ml-' + m.merchantid"
                :key="'ml-' + m.merchantid"
                :class="(isActiveMerchant(m) ? 'pulsate-highlight-white' : '')"
                @click="showMerchantList(m)"
              >
                <td
                  class="text-center"
                  style="width: 120px"
                >
                  <b-button
                    v-if="isSelector"
                    class="btn-dark btn-sm btn-outline-warning mr-3"
                    title="Select Merchant List"
                    @click="selectMerchantList(m.merchantid);"
                  >
                    <i class="fa fa-arrow-left"></i>
                  </b-button>

                  <b-button
                    class="btn-dark btn-sm btn-outline-danger mr-3"
                    @click="deleteMerchantList(m)"
                    title="Delete Merchant List"
                  >
                    <i class="fa fa-trash-o"></i>
                  </b-button>

                  <b-button
                    class="btn-dark btn-sm btn-outline-warning"
                    @click="editMerchantList(m.merchantid)"
                    title="Edit Merchant List"
                  >
                    <i class="fa fa-edit"></i>
                  </b-button>
                </td>

                <td class="text-center" style="width: 50px">{{ m.merchantid }}</td>
                <td class="text-center" style="width: 50px">{{ m.slot }}</td>
                <td class="text-left">
                  <div
                    v-for="d in associatedNpcs[m.merchantid]"
                  >
                    ({{ d.npc.id }}) {{ getNpcCleanName(d.npc.name) }} {{ d.npc.lastname ? `(${d.npc.lastname})` : '' }}
                    ({{ d.zone }})

                  </div>
                </td>
              </tr>
              </tbody>
            </table>

          </div>

          <div class="row text-center justify-content-center">
            <div class="col-12 text-center mt-3">
              <b-pagination
                class="mb-1"
                v-model="currentPage"
                :total-rows="totalRows"
                :hide-ellipsis="true"
                :per-page="100"
                @change="paginate"
              />
            </div>
          </div>

        </eq-window>
      </div>

      <!-- Active Merchant (raw) pane -->
      <div v-if="activeMerchantList && Object.keys(activeMerchantList).length > 0" class="col-5">
        <eq-window
          :title="`Merchant List Preview (${activeMerchantList[0].merchantid})`"
        >
          <div style="max-height: 92vh; overflow-y: scroll; overflow-x: hidden">
            <table
              class="eq-table eq-highlight-rows minified-inputs"
              style="width: 95%"
            >
              <thead>
              <tr>
                <th>Item</th>
              </tr>
              </thead>
              <tbody>
              <tr
                v-for="(e, i) in activeMerchantList"
                :key="e.slot + '-' + e.item"
                v-if="itemData[e.item]"
              >
                <td>
                  <!--                  <input type="text" v-model="e.item" class="mr-3 m-0" style="width: 120px">-->
                  <item-popover
                    class="d-inline-block"
                    :item="itemData[e.item]"
                    v-if="itemData[e.item] && Object.keys(itemData[e.item]).length > 0"
                    size="sm"
                  />

                  {{ itemData[e.item] && itemData[e.item].stacksize > 0 ? `(${itemData[e.item].stacksize})` : '' }}
                  <eq-cash-display
                    class="ml-1"
                    :price="parseInt(itemData[e.item].price)"
                    v-if="itemData[e.item]"
                  />
                </td>
              </tr>
              </tbody>
            </table>
          </div>
        </eq-window>
      </div>


      <!-- Active Merchant (NPC) pane -->
      <div :class="(Object.keys(activeMerchantNpc).length > 0 ? 'col-5' : 'col-12')">
        <eq-window
          v-if="Object.keys(activeMerchantNpc).length > 0"
          style="max-height: 95vh; overflow-y: scroll; overflow-x: hidden"
        >
          <eq-npc-card-preview
            :npc="activeMerchantNpc"
            :no-stats="true"
            :limit-entries="1000"
          />
        </eq-window>
      </div>
    </div>
  </div>
</template>

<script>
import {Zones}               from "../../app/zones";
import ContentArea           from "../../components/layout/ContentArea";
import EqWindow              from "../../components/eq-ui/EQWindow";
import LoaderFakeProgress    from "../../components/LoaderFakeProgress";
import NpcPopover            from "../../components/NpcPopover";
import {ROUTE}               from "../../routes";
import EqNpcCardPreview      from "../../components/preview/EQNpcCardPreview";
import {Merchants}           from "../../app/merchants";
import ItemPopover           from "../../components/ItemPopover";
import EqCashDisplay         from "../../components/eq-ui/EqCashDisplay";
import {Items}               from "../../app/items";
import EqCheckbox            from "../../components/eq-ui/EQCheckbox";
import EqDebug               from "../../components/eq-ui/EQDebug";
import ItemSelector          from "../../components/selectors/ItemSelector";
import {MerchantlistApi}     from "../../app/api";
import {SpireApiClient}      from "../../app/api/spire-api-client";
import {SpireQueryBuilder}   from "../../app/api/spire-query-builder";
import {chunk}               from "../../app/utility/chunk";
import {Npcs}                from "../../app/npcs";
import MerchantlistEntryEdit from "./components/MerchantlistEntryEdit";
import MerchantEdit          from "./MerchantEdit";
import util                  from "util";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "MerchantSubEditor",
  components: {
    MerchantEdit,
    MerchantlistEntryEdit,
    ItemSelector,
    EqDebug,
    EqCheckbox,
    EqCashDisplay,
    ItemPopover,
    EqNpcCardPreview,
    NpcPopover,
    LoaderFakeProgress,
    EqWindow,
    ContentArea
  },
  async created() {
    this.loadQueryState()
    this.ml             = []
    this.editList       = []
    this.itemData      = {}
    this.merchantLists  = []
    this.associatedNpcs = {}

    // editing - entry level
    this.editMerchantEntry = {}
    this.zones             = await Zones.getZones()

    this.init()
  },
  props: {
    // if this component is being used as a selector
    isSelector: {
      type: Boolean,
      required: false,
      default: false
    },
  },
  data() {
    return {

      // pagination (all)
      currentPage: 1,
      totalRows: 0,

      // selection
      search: "",
      searchItemName: "",
      zoneSelection: 0,

      loading: false,

      // selectors
      selectorActive: {},
      lastResetTime: Date.now(),

      // state
      showAll: false,

      // preview
      activeMerchantNpc: {},
      activeMerchantList: {},

      // feeds selection
      zones: [],
    }
  },
  watch: {
    // when ran standalone, route drives state
    '$route'() {
      this.reset()
      this.loadQueryState()
      this.init()
    },
  },
  methods: {

    paginate() {
      // models aren't quite updated when we trigger this so queue the pagination
      setTimeout(() => {
        console.log("We're paginating")
        console.log(this.currentPage)
        console.log(this.totalRows)
        this.updateQueryState()
      }, 100)

    },

    isNavigated() {
      return Object.keys(this.$route.query).length
    },

    async createNewMerchant() {
      const r = await Merchants.create()
      if (r.status === 200) {
        // router navigate here
        this.editMerchantList(r.data.merchantid)
      }
    },

    async deleteMerchantList(m) {
      if (confirm(`Are you sure you want to delete this Merchant? (${m.merchantid}) with (${m.slot}) items?`)) {
        await Merchants.deleteMerchant(m.merchantid)
        this.activeMerchantList = {}
        this.loading = true
        await this.showAllMerchants()
        this.$forceUpdate()
        this.loading = false
      }
    },

    isActiveMerchant(m) {
      return this.activeMerchantList && this.activeMerchantList[0] && m.merchantid === this.activeMerchantList[0].merchantid
    },

    getNpcCleanName(name) {
      return Npcs.getCleanName(name)
    },

    async showMerchantList(m) {

      // pluck the merchantlist off of related data
      if (m.npc_types && m.npc_types.length > 0) {
        this.activeMerchantList = m.npc_types[0].merchantlists

        let itemIds = []
        for (let e of this.activeMerchantList) {
          if (e.item > 0 && !this.itemData[e.item]) {
            itemIds.push(e.item)
          }
        }

        if (itemIds.length > 0) {
          setTimeout(() => {
            Items.loadItemsBulk(itemIds).then(async () => {
              for (let e of this.activeMerchantList) {
                if (e.item > 0 && !this.itemData[e.item]) {
                  this.itemData[e.item] = await Items.getItem(e.item)
                }
              }
              this.$forceUpdate()
            })
          }, 10)
        }
      }
    },

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.zoneSelection !== 0) {
        queryState.zone = JSON.stringify(this.zoneSelection)
      }
      if (this.showAll) {
        queryState.showAll = this.showAll
      }
      if (this.currentPage > 0) {
        queryState.page = this.currentPage
      }
      if (this.search !== '') {
        queryState.q = this.search
      }
      if (this.searchItemName !== '') {
        queryState.s = this.searchItemName
      }

      this.$router.push(
        {
          path: ROUTE.MERCHANTS,
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (typeof this.$route.query.zone !== 'undefined' && this.$route.query.zone !== 0) {
        this.zoneSelection = JSON.parse(this.$route.query.zone);
      }
      if (typeof this.$route.query.q !== 'undefined' && this.$route.query.q !== '') {
        this.search = this.$route.query.q;
      }
      if (typeof this.$route.query.s !== 'undefined' && this.$route.query.s !== '') {
        this.searchItemName = this.$route.query.s;
      }
      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
      }
      if (typeof this.$route.query.showAll !== 'undefined' && this.$route.query.showAll) {
        this.showAll = true;
      }
    },

    /**
     * Init
     */
    async init() {
      this.activeMerchantNpc = {}
      this.ml                = []
      this.loading           = true
      const z                = this.zoneSelection

      if (Object.keys(this.zoneSelection).length > 0) {
        this.ml = (await Merchants.getMerchantsByZone(z.z, z.v))
      }
      if (this.search.length > 0) {
        this.ml = (await Merchants.getMerchantsByName(this.search))
      }
      if (this.searchItemName.length > 0) {
        this.ml = (await Merchants.getMerchantsByItemName(this.searchItemName))
      }

      if (this.showAll) {
        await this.showAllMerchants()
      }

      console.log("triggering force update")
      this.$forceUpdate()
      this.loading = false
    },

    async showAllMerchants() {
      console.log("show all")

      if (this.totalRows === 0) {
        this.totalRows = await Merchants.getTotalMerchants()
      }

      // @ts-ignore
      const r = await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
        .listMerchantlists(
          // @ts-ignore
          (new SpireQueryBuilder())
            .groupBy(["merchantid"])
            .orderBy(["merchantid"])
            .orderDirection("desc")
            .limit(100)
            .page(this.currentPage)
            .get()
        )

      let merchantIds = []
      if (r.status === 200) {
        merchantIds = r.data.map((e) => {
          return e.merchantid
        })
      }

      // chunk requests
      let merchants = []
      for (let c of chunk(merchantIds, 500)) {
        const b = await Merchants.getMerchantsBulk(c, ["NpcTypes.Spawnentries.Spawngroup.Spawn2", "NpcTypes.Merchantlists"])

        // @ts-ignore
        merchants = [...merchants, ...b]
      }

      if (r.status === 200) {
        this.merchantLists  = merchants
        this.associatedNpcs = {}

        // get associated NPCs to the merchant lists
        for (let m of merchants) {
          if (m.npc_types && m.npc_types.length > 0) {
            // console.log(m)

            for (let n of m.npc_types) {
              if (n.spawnentries && n.spawnentries.length > 0) {
                for (let s of n.spawnentries) {
                  // console.log(s)
                  if (s.spawngroup && s.spawngroup.spawn_2) {
                    if (typeof this.associatedNpcs[m.merchantid] === 'undefined') {
                      this.associatedNpcs[m.merchantid] = []
                    }

                    this.associatedNpcs[m.merchantid].push(
                      {
                        npc: n,
                        zone: s.spawngroup.spawn_2.zone
                      }
                    )
                  }
                }
              }
            }
          }
        }

        // console.log("associated")
        // console.log(this.associatedNpcs)

      }
    },

    /**
     * Controls
     */
    editMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Editing [%s]", merchantId)

      console.log(this.$router.query)

      this.$router.push(
        {
          path: util.format(ROUTE.MERCHANT_EDIT, merchantId),
          query: {
            return: JSON.stringify(this.$route.query)
          }
        }
      ).catch(() => {
      })
    },
    selectMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Selecting [%s]", merchantId)

      this.$emit('input', merchantId);
    },

    reset() {
      this.showAll               = false;
      this.search                = ""
      this.searchItemName        = ""
      this.zoneSelection         = 0
      this.editMerchantEntry     = {}
      this.editMerchantEntrySlot = 0
      this.activeMerchantNpc     = {}
      this.ml                    = []
      this.merchantLists         = []
      this.activeMerchantList    = []
      this.associatedNpcs        = {}
    },
  }
}
</script>

<style>

#merchant-list-table td {
  vertical-align: middle;
  padding: 10px;
  height: 60px;
}
</style>
