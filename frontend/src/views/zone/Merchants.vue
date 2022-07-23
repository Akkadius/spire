<template>
  <div>
    <!-- Browse View -->
    <div class="row" v-if="editMerchantId === 0">
      <div :class="(Object.keys(activeMerchantNpc).length > 0 || Object.keys(activeMerchantList).length > 0 ? 'col-7' : 'col-12')">
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

        <eq-window v-if="loading">
          <div class="text-center">
            Loading
            <loader-fake-progress class="mt-3"/>
          </div>
        </eq-window>

        <eq-window v-if="isNavigated() && !loading && ml && ml.length === 0">
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
                <th></th>
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
                  style="width: 100px"
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
                    class="btn-dark btn-sm btn-outline-warning"
                    @click="editMerchantList(n.merchant_id)"
                    title="Edit Merchant List"
                  >
                    <i class="fa fa-edit"></i>
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
          :title="'Merchants (' + merchantLists.length + ')'"
        >
          <div style="overflow-y: scroll; max-height: 83vh">
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
                :key="index"
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

                    <!--                    <npc-popover-->
                    <!--                      :limit-entries="25"-->
                    <!--                      :additional-label="`(${d.zone})`"-->
                    <!--                      :no-stats="true"-->
                    <!--                      :npc="d.npc"-->
                    <!--                    />-->
                  </div>
                </td>
              </tr>
              </tbody>
            </table>
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
                v-if="editItems[e.item]"
              >
                <td>
                  <!--                  <input type="text" v-model="e.item" class="mr-3 m-0" style="width: 120px">-->
                  <item-popover
                    class="d-inline-block"
                    :item="editItems[e.item]"
                    v-if="editItems[e.item] && Object.keys(editItems[e.item]).length > 0"
                    size="sm"
                  />

                  {{ editItems[e.item] && editItems[e.item].stacksize > 0 ? `(${editItems[e.item].stacksize})` : '' }}
                  <eq-cash-display
                    class="ml-1"
                    :price="parseInt(editItems[e.item].price)"
                    v-if="editItems[e.item]"
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

    <!-- Edit View -->
    <div class="row" v-if="editMerchantId > 0">
      <div :class="(Object.keys(editMerchantEntry).length > 0 || addItem ? 'col-7' : 'col-12')">
        <eq-window :title='`Edit Merchant (${editMerchantId})`'>
          <!--          <pre style="width: 100%">{{editList}}</pre>-->

          <div v-if="editList && editList.length === 0" class="font-weight-bold mb-3">
            There are no items on this Merchant, perhaps you should add some?
          </div>

          <div class="row">
            <div class="col-6">
              <div class="btn-group d-inline-block" role="group">
                <b-button
                  size="sm"
                  variant="warning"
                  @click="$router.go(-1)"
                >
                  <i class="fa fa-arrow-left mr-1"></i>
                  Go back to Merchants
                </b-button>

                <b-button
                  size="sm"
                  variant="outline-warning"
                  @click="addItemToMerchantListQueue()"
                >
                  <i class="fa fa-plus mr-1"></i>
                  Add Item
                </b-button>

                <b-button
                  class="btn-dark btn-sm btn-outline-danger mr-3"
                  @click="deleteMerchantListFromEdit(editList[0])"
                  title="Delete Merchant List"
                  v-if="editList && editList[0]"
                >
                  <i class="fa fa-trash-o"></i> Delete Merchant
                </b-button>

              </div>
            </div>
            <div class="col-2 text-center" v-if="applyingChanges">
              <div class="ml-3 p-0 m-0">
                Applying changes... Please wait...
                <loader-fake-progress class="mt-2"/>
              </div>
            </div>
          </div>

          <div
            v-if="editList && editList.length > 0"
            class=""
          >
            <table
              class="eq-table eq-highlight-rows minified-inputs"
              style="width: 95%"
            >
              <thead>
              <tr>
                <th class="text-center" style="width: 100px">Actions</th>
                <th class="text-center" style="width: 50px">Slot</th>
                <th>Item</th>
              </tr>
              </thead>
              <tbody>
              <tr
                v-for="(e, i) in editList"
                :key="e.slot + '-' + e.item"
                :class="(isMerchantEntrySelected(e) ? 'pulsate-highlight-white' : '')"
              >
                <td class="text-left p-0">
                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-danger ml-1"
                    style="padding: 0px 4px;"
                    title="Edit"
                    @click="deleteMerchantRow(e)"
                  >
                    <i class="fa fa-trash"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-success ml-1"
                    style="padding: 0px 4px;"
                    title="Edit"
                    @click="editMerchantRow(e)"
                  >
                    <i class="fa fa-pencil"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-light ml-1"
                    style="padding: 0px 4px;"
                    title="Move slot up"
                    @click="moveSlotUp(e)"
                    v-if="editList[i - 1]"
                  >
                    <i class="fa fa-arrow-up"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-light ml-1"
                    style="padding: 0px 4px;"
                    title="Move slot down"
                    @click="moveSlotDown(e)"
                    v-if="editList[i + 1]"
                  >
                    <i class="fa fa-arrow-down"></i>
                  </b-button>
                </td>
                <td class="text-center p-0">
                  {{ e.slot }}
                </td>
                <td>
                  <!--                  <input type="text" v-model="e.item" class="mr-3 m-0" style="width: 120px">-->
                  <item-popover
                    class="d-inline-block"
                    :item="editItems[e.item]"
                    v-if="editItems[e.item] && Object.keys(editItems[e.item]).length > 0"
                    size="sm"
                  />

                  {{ editItems[e.item] && editItems[e.item].stacksize > 0 ? `(${editItems[e.item].stacksize})` : '' }}
                  <eq-cash-display
                    class="ml-1"
                    :price="parseInt(editItems[e.item].price)"
                    v-if="editItems[e.item]"
                  />
                </td>
              </tr>
              </tbody>
            </table>

          </div>

        </eq-window>
      </div>

      <div class="col-5" v-if="(Object.keys(editMerchantEntry).length > 0)">
        <eq-window :title="`Edit Merchant List Entry (${editMerchantEntrySlot})`">
          <div
            v-for="field in editMerchantEntryFields"
            :key="field.field"
            :class="'row'"
          >
            <div
              class="col-4 text-right p-0 m-0 mr-1 mt-3"
              style="position: relative; bottom: 6px;"
              v-if="field.fType === 'checkbox'"
            >
              <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
              {{ field.desc }}
            </div>
            <div
              class="col-4 text-right p-0 m-0 mr-3"
              v-if="field.fType !== 'checkbox'"
              style="margin-top: 10px !important"
            >
              <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
              {{ field.desc }}
            </div>

            <!--                  <div class="text-center" v-if="field.fType !== 'checkbox'">-->
            <!--                    <span-->
            <!--                      v-if="field.itemIcon"-->
            <!--                      :class="'item-' + field.itemIcon + '-sm'"-->
            <!--                      style="display: inline-block"-->
            <!--                    />-->
            <!--                    {{ field.desc }}-->
            <!--                  </div>-->

            <div class="col-7 text-left p-0 mt-2">

              <!-- checkbox -->
              <div :class="'text-left ml-2 mt-1'" v-if="field.fType === 'checkbox'">
                <!--                        <div class="d-inline-block" style="bottom: 2px; position: relative; margin-right: 1px">-->
                <!--                          {{ field.desc }}-->
                <!--                        </div>-->
                <eq-checkbox
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  class="d-inline-block text-center"
                  :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                  :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                  v-model.number="editMerchantEntry[field.field]"
                  @input="editMerchantEntry[field.field] = $event"

                />
              </div>

              <!-- input number -->
              <b-form-input
                v-if="field.fType === 'number'"
                :id="field.field"
                v-model.number="editMerchantEntry[field.field]"
                class="m-0 mt-1"
                v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                :style="(editMerchantEntry[field.field] === 0 ? 'opacity: .5' : '')"
              />

              <!-- input text -->
              <b-form-input
                v-if="field.fType === 'text'"
                :id="field.field"
                v-model.number="editMerchantEntry[field.field]"
                class="m-0 mt-1"
                v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                :style="(editMerchantEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
              />

              <!-- textarea -->
              <b-textarea
                v-if="field.fType === 'textarea'"
                :id="field.field"
                v-model="editMerchantEntry[field.field]"
                class="m-0 mt-1"
                rows="2"
                max-rows="6"
                v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                :style="(editMerchantEntry[field.field] === '' ? 'opacity: .5' : '') + ';'"
              ></b-textarea>

              <!-- select -->
              <select
                v-model.number="editMerchantEntry[field.field]"
                :id="field.field"
                class="form-control m-0 mt-1"
                v-if="field.selectData"
                v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                :style="(editMerchantEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
              >
                <option
                  v-for="(desc, index) in field.selectData"
                  :key="index"
                  :value="parseInt(index)"
                >
                  {{ index }}) {{ desc }}
                </option>
              </select>

            </div>
          </div>

          <eq-debug :data="editMerchantEntry"/>
        </eq-window>
      </div>

      <div class="col-5" v-if="addItem">
        <item-selector
          @input="addItemToMerchantList($event)"
        />
      </div>
    </div>

  </div>
</template>

<script>
import {Zones}             from "../../app/zones";
import ContentArea         from "../../components/layout/ContentArea";
import EqWindow            from "../../components/eq-ui/EQWindow";
import LoaderFakeProgress  from "../../components/LoaderFakeProgress";
import NpcPopover          from "../../components/NpcPopover";
import {ROUTE}             from "../../routes";
import EqNpcCardPreview    from "../../components/preview/EQNpcCardPreview";
import {Merchants}         from "../../app/merchants";
import ItemPopover         from "../../components/ItemPopover";
import EqCashDisplay       from "../../components/eq-ui/EqCashDisplay";
import {Items}             from "../../app/items";
import EqCheckbox          from "../../components/eq-ui/EQCheckbox";
import EqDebug             from "../../components/eq-ui/EQDebug";
import ItemSelector        from "../../components/selectors/ItemSelector";
import {MerchantlistApi}   from "../../app/api";
import {SpireApiClient}    from "../../app/api/spire-api-client";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import {chunk}             from "../../app/utility/chunk";
import {Npcs}              from "../../app/npcs";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "MerchantSubEditor",
  components: {
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
    this.ml                = []
    this.editList          = []
    this.editItems         = {}
    this.merchantLists     = []
    this.associatedNpcs    = {}
    // editing - entry level
    this.editMerchantEntry = 0
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
      // selection
      search: "",
      searchItemName: "",
      zoneSelection: 0,

      loading: false,
      applyingChanges: false,

      // selectors
      selectorActive: {},
      lastResetTime: Date.now(),

      // editing
      editMerchantId: 0,

      // state
      addItem: false,
      showAll: false,

      editMerchantEntryFields: [
        { desc: "faction_required", field: "faction_required", fType: "text" },
        { desc: "level_required", field: "level_required", fType: "text" },
        { desc: "alt_currency_cost", field: "alt_currency_cost", fType: "text" },
        { desc: "classes_required", field: "classes_required", fType: "text" },
        { desc: "probability", field: "probability", fType: "text" },
        { desc: "min_expansion", field: "min_expansion", fType: "text" },
        { desc: "max_expansion", field: "max_expansion", fType: "text" },
        { desc: "content_flags", field: "content_flags", fType: "text" },
        { desc: "content_flags_disabled", field: "content_flags_disabled", fType: "text" },
      ],

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

    isNavigated() {
      return Object.keys(this.$route.query).length
    },

    async createNewMerchant() {
      const r                    = await Merchants.create()
      if (r.status === 200) {
        this.editMerchantId        = r.data.merchantid
        this.addItem               = false
        this.updateQueryState()
        this.$forceUpdate()
      }
    },

    async deleteMerchantList(m) {
      if (confirm(`Are you sure you want to delete this Merchant? (${m.merchantid}) with (${m.slot}) items?`)) {
        await Merchants.deleteMerchant(m.merchantid)
        this.loading = true
        await this.showAllMerchants()
        this.$forceUpdate()
        this.loading = false
      }
    },

    async deleteMerchantListFromEdit(m) {
      if (confirm(`Are you sure you want to delete this Merchant? (${m.merchantid}) with (${m.slot}) items?`)) {
        await Merchants.deleteMerchant(m.merchantid)
        this.refreshMerchantlistEntries()
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
          if (e.item > 0 && !this.editItems[e.item]) {
            itemIds.push(e.item)
          }
        }

        if (itemIds.length > 0) {
          setTimeout(() => {
            Items.loadItemsBulk(itemIds).then(async () => {
              for (let e of this.activeMerchantList) {
                if (e.item > 0 && !this.editItems[e.item]) {
                  this.editItems[e.item] = await Items.getItem(e.item)
                }
              }
              this.$forceUpdate()
            })
          }, 10)
        }
      }
    },

    /**
     * Merchant List entry editing functions
     */
    async reorderItems() {
      console.log("[MerchantSubEditor] Reordering items")

      this.applyingChanges = true

      let startingRewriteSlot = 0
      for (let [i, e] of this.editList.entries()) {
        if (this.editList[i + 1] && this.editList[i]) {
          const isGap = (this.editList[i + 1].slot - this.editList[i].slot) > 1
          // console.log(
          //   "[MerchantSubEditor] isGap current [%s] next [%s] gap [%s]",
          //   this.editList[i].slot,
          //   this.editList[i + 1].slot,
          //   isGap
          // )

          // on our first gap set the first rewrite slot to renumber the rest of the entries
          if (isGap && startingRewriteSlot === 0) {
            startingRewriteSlot = this.editList[i].slot + 1
          }

          if (startingRewriteSlot > 0) {
            let desiredEntry  = JSON.parse(JSON.stringify(this.editList[i + 1]))
            desiredEntry.slot = startingRewriteSlot

            await Merchants.updateSlotForEntry(
              e.merchantid,
              this.editList[i + 1].slot,
              desiredEntry
            )
          }

          if (startingRewriteSlot !== 0) {
            startingRewriteSlot++
          }
        }
      }

      this.applyingChanges = false

    },
    editMerchantRow(row) {
      console.log("[MerchantSubEditor] Editing merchant row slot [%s]", row.slot)

      this.addItem               = false
      this.editMerchantEntrySlot = row.slot
      this.updateQueryState()
      this.$forceUpdate()
    },
    async deleteMerchantRow(row) {
      console.log("[MerchantSubEditor] Deleting merchant row slot [%s]", row.slot)

      await Merchants.deleteMerchantEntry(row.merchantid, row.slot)
      this.editList = this.editList.filter((e) => {
        return e.slot !== row.slot
      })

      this.refreshMerchantlistEntries()
    },
    async refreshMerchantlistEntries() {
      await this.reorderItems()

      // reset so we can reload after deleting entry
      this.editList              = []
      this.editMerchantEntrySlot = 0

      this.init()
    },

    async addItemToMerchantList(e) {
      const itemId = e.id
      let newSlot  = this.editList[this.editList.length - 1] ? this.editList[this.editList.length - 1].slot + 1 : 1

      await Merchants.addItemToMerchant(
        parseInt(this.editMerchantId),
        newSlot,
        itemId
      )

      this.refreshMerchantlistEntries()
    },

    addItemToMerchantListQueue() {
      this.editMerchantEntry     = 0
      this.editMerchantEntrySlot = 0

      console.log("[MerchantSubEditor] Queue adding new item (sub editor)")
      this.addItem = true
      this.updateQueryState()
    },
    async moveEntry(e, direction) {
      for (let [i, entry] of this.editList.entries()) {
        if (this.editList[i] && this.editList[i].slot === e.slot) {
          let accessIndex = 0
          if (direction === "up") {
            accessIndex = i - 1
          }
          if (direction === "down") {
            accessIndex = i + 1
          }

          console.log("we found the slot to move! [%s]", e.slot)

          if (this.editList[accessIndex]) {
            const entryToMove    = JSON.parse(JSON.stringify(e))
            const entryToReplace = JSON.parse(JSON.stringify(this.editList[accessIndex]))
            const toMoveSlot     = entryToMove.slot
            const toReplaceSlot  = entryToReplace.slot
            const originalItemId = entryToMove.item

            // console.log("entry to move", entryToMove)
            // console.log("entry to replace", entryToReplace)

            // move the current one up (update slot), temporarily set the item to -1
            let firstEntryMove  = entryToMove
            firstEntryMove.slot = toReplaceSlot
            firstEntryMove.item = -1

            // console.log("#1")

            await Merchants.updateSlotForEntry(
              e.merchantid,
              toReplaceSlot,
              firstEntryMove
            )

            // move the entry to replace down
            let secondEntryMove  = entryToReplace
            secondEntryMove.slot = toMoveSlot

            // console.log("#2")

            await Merchants.updateSlotForEntry(
              e.merchantid,
              toMoveSlot,
              secondEntryMove
            )

            // console.log("#3")

            // move the previous current on back to original item
            let lastEntryMove   = entryToMove
            firstEntryMove.slot = toReplaceSlot
            firstEntryMove.item = originalItemId

            await Merchants.updateSlotForEntry(
              e.merchantid,
              toReplaceSlot,
              lastEntryMove
            )
          }
        }
      }

      await this.refreshMerchantlistEntries()
    },

    async moveSlotUp(e) {
      await this.moveEntry(e, "up")

      console.log("[MerchantSubEditor] moveSlotUp [%s]", e.slot)
    },
    async moveSlotDown(e) {
      await this.moveEntry(e, "down")

      console.log("[MerchantSubEditor] moveSlotDown [%s]", e.slot)
    },
    isMerchantEntrySelected(e) {
      return e.slot === this.editMerchantEntry.slot;
    },

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.zoneSelection !== 0) {
        queryState.zone = JSON.stringify(this.zoneSelection)
      }
      if (this.addItem) {
        queryState.addItem = this.addItem
      }
      if (this.showAll) {
        queryState.showAll = this.showAll
      }
      if (this.search !== '') {
        queryState.q = this.search
      }
      if (this.searchItemName !== '') {
        queryState.s = this.searchItemName
      }
      if (this.editMerchantId !== 0) {
        queryState.edit = this.editMerchantId
      }
      if (this.editMerchantEntrySlot !== 0) {
        queryState.e = this.editMerchantEntrySlot
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
      if (typeof this.$route.query.edit !== 'undefined' && this.$route.query.edit !== 0) {
        this.editMerchantId = this.$route.query.edit;
      }
      if (typeof this.$route.query.e !== 'undefined' && parseInt(this.$route.query.e) !== 0) {
        this.editMerchantEntrySlot = parseInt(this.$route.query.e);
      }
      if (typeof this.$route.query.addItem !== 'undefined' && this.$route.query.addItem) {
        this.addItem = true;
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

      if (this.editMerchantId > 0) {
        console.log("[MerchantSubEditor] Editing merchant [%s]", this.editMerchantId)

        // this is to keep the same list from being redrawn / reloaded repeatedly
        let sameListLoaded = this.editList.find((e) => {
          return e.merchantid === parseInt(this.editMerchantId)
        })

        if (typeof sameListLoaded === 'undefined') {
          await this.loadEditMerchant()
        }
      }
      console.log("this.editMerchantEntrySlot", this.editMerchantEntrySlot)

      if (this.editMerchantEntrySlot) {
        console.log("[MerchantSubEditor] Editing merchant entry slot [%s]", this.editMerchantEntrySlot)
        for (let e of this.editList) {
          if (e.slot === parseInt(this.editMerchantEntrySlot)) {
            this.editMerchantEntry = e
          }
        }
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
      // @ts-ignore
      const r = await (new MerchantlistApi(SpireApiClient.getOpenApiConfig()))
        .listMerchantlists(
          // @ts-ignore
          (new SpireQueryBuilder())
            .groupBy(["merchantid"])
            .orderBy(["merchantid"])
            .orderDirection("desc")
            .limit(100000)
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

    async loadEditMerchant() {
      this.editList = await Merchants.getById(this.editMerchantId)

      let itemIds = []
      for (let e of this.editList) {
        if (e.item > 0 && !this.editItems[e.item]) {
          itemIds.push(e.item)
        }
      }
      if (itemIds.length > 0) {
        Items.loadItemsBulk(itemIds).then(async () => {
          for (let e of this.editList) {
            if (e.item > 0 && !this.editItems[e.item]) {
              this.editItems[e.item] = await Items.getItem(e.item)
            }
          }
          this.$forceUpdate()
        })
      }
    },

    /**
     * Misc
     */
    getFieldDescription(field) {
      return ""
    },

    /**
     * Controls
     */
    editMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Editing [%s]", merchantId)
      this.reset()
      this.editMerchantId = merchantId
      this.updateQueryState()
    },
    selectMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Selecting [%s]", merchantId)

      this.$emit('input', merchantId);
    },


    reset() {
      this.showAll               = false;
      this.addItem               = false
      this.search                = ""
      this.searchItemName        = ""
      this.zoneSelection         = 0
      this.editMerchantId        = 0
      this.editMerchantEntry     = 0
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
