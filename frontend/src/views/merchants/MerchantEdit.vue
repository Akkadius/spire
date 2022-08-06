<template>
  <content-area style="padding: 0px !important">
    <div class="row">
      <div :class="(Object.keys(editMerchantEntry).length > 0 || addItem ? 'col-7' : 'col-12')">
        <eq-window :title='`Edit Merchant (${editMerchantId})`'>

          <div v-if="editList && editList.length === 0" class="font-weight-bold mb-3">
            There are no items on this Merchant, perhaps you should add some?
          </div>

          <div class="row mb-1">
            <div class="col-6">
              <div class="btn-group d-inline-block" role="group">
                <b-button
                  size="sm"
                  variant="warning"
                  @click="goBack"
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

            <div class="col-6" v-if="!applyingChanges">
              <div>
                <info-error-banner
                  :slim="true"
                  :notification="notification"
                  :error="error"
                  @dismiss-error="error = ''"
                  @dismiss-notification="notification = ''"
                  class="mt-0"
                />
              </div>
            </div>

            <div class="col-6 text-right" v-if="applyingChanges">
              <div class="ml-3 p-0 m-0">
                Applying
                <loader-fake-progress class="ml-3 mt-2 d-inline-block"/>
              </div>
            </div>

          </div>

          <div
            v-if="editList && editList.length > 0"
            class=""
            style="height: 87vh; overflow-y: scroll"
          >
            <table
              class="eq-table eq-highlight-rows minified-inputs"
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th class="text-center" style="width: 120px">Actions</th>
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
                    style="padding: 0px 6px;"
                    title="Edit"
                    :disabled="applyingChanges"
                    @click="deleteMerchantRow(e)"
                  >
                    <i class="fa fa-trash"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-success ml-1"
                    style="padding: 0px 6px;"
                    title="Edit"
                    :disabled="applyingChanges"
                    @click="editMerchantRow(e)"
                  >
                    <i class="fa fa-pencil"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-light ml-1"
                    style="padding: 0px 6px;"
                    title="Move slot up"
                    @click="moveSlotUp(e)"
                    :disabled="applyingChanges"
                    v-if="editList[i - 1]"
                  >
                    <i class="fa fa-arrow-up"></i>
                  </b-button>

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-light ml-1"
                    style="padding: 0px 6px;"
                    title="Move slot down"
                    :disabled="applyingChanges"
                    @click="moveSlotDown(e)"
                    v-if="editList[i + 1]"
                  >
                    <i class="fa fa-arrow-down"></i>
                  </b-button>
                </td>
                <td class="text-center p-0">
                  {{ e.slot }}
                </td>
                <td :style="(applyingChanges ? 'opacity: .2' : '')">
                  <!--                  <input type="text" v-model="e.item" class="mr-3 m-0" style="width: 120px">-->

                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-outline-success mr-3"
                    style="padding: 0px 6px;"
                    title="Edit Item"
                    @click="editItem(e.item)"
                  >
                    <i class="fa fa-pencil-square"></i>
                  </b-button>

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

                  <content-filter-display-pills :filter-data="e"/>

                </td>
              </tr>
              </tbody>
            </table>

          </div>

        </eq-window>
      </div>

      <merchantlist-entry-edit
        :edit-merchant-entry="editMerchantEntry"
        :edit-merchant-entry-slot="editMerchantEntrySlot"
        :key="resetTime"
      />

      <div class="col-5" v-if="addItem">
        <item-selector
          @input="addItemToMerchantList($event)"
        />
      </div>
    </div>
  </content-area>
</template>

<script>
import EqWindow                  from "../../components/eq-ui/EQWindow";
import EqCashDisplay             from "../../components/eq-ui/EqCashDisplay";
import ItemPopover               from "../../components/ItemPopover";
import LoaderFakeProgress        from "../../components/LoaderFakeProgress";
import {Merchants}               from "../../app/merchants";
import ContentArea               from "../../components/layout/ContentArea";
import {ROUTE}                   from "../../routes";
import {Zones}                   from "../../app/zones";
import util                      from "util";
import {Items}                   from "../../app/items";
import ItemSelector              from "../../components/selectors/ItemSelector";
import MerchantlistEntryEdit     from "./components/MerchantlistEntryEdit";
import ExpansionIcon             from "../../components/preview/ExpansionIcon";
import ContentFlagPills          from "../../components/preview/ContentFlagPills";
import ContentFilterDisplayPills from "../../components/preview/ContentFilterDisplayPills";
import InfoErrorBanner           from "../../components/InfoErrorBanner";

export default {
  name: "MerchantEdit",
  components: {
    InfoErrorBanner,
    ContentFilterDisplayPills,
    ContentFlagPills,
    ExpansionIcon,
    MerchantlistEntryEdit,
    ItemSelector,
    ContentArea,
    LoaderFakeProgress,
    ItemPopover,
    EqCashDisplay,
    EqWindow
  },
  data() {
    return {
      // editing
      editMerchantId: 0,
      editMerchantEntrySlot: 0,

      // api notify / error
      notification: "",
      error: "",

      // state
      addItem: false,
      loading: false,
      applyingChanges: false,

      returnQuery: {},

      resetTime: 0,
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

  async created() {
    this.loadQueryState()
    this.ml             = []
    this.editList       = []
    this.itemData       = {}
    this.merchantLists  = []
    this.associatedNpcs = {}

    // editing - entry level
    this.editMerchantEntry = {}
    this.zones             = await Zones.getZones()

    this.init()
  },

  methods: {

    editItem(itemId) {
      this.$router.push(
        {
          path: util.format(ROUTE.ITEM_EDIT, itemId),
          query: {}
        }
      ).catch(() => {
      })
    },

    goBack() {
      console.log(this.$route.query.return0)

      this.$router.push(
        {
          path: ROUTE.MERCHANTS,
          query: JSON.parse(this.$route.query.return)
        }
      ).catch(() => {
      })
    },

    reset() {
      this.editMerchantId        = 0
      this.editMerchantEntrySlot = 0
      this.editMerchantEntry     = {}
      this.addItem               = false
      this.resetTime             = Date.now()
      this.returnQuery           = {}
    },

    /**
     * Init
     */
    async init() {
      this.ml      = []
      this.loading = true

      console.log(this.editMerchantId)

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

    async loadEditMerchant() {
      this.editList = await Merchants.getById(this.editMerchantId)

      let itemIds = []
      for (let e of this.editList) {
        if (e.item > 0 && !this.itemData[e.item]) {
          itemIds.push(e.item)
        }
      }
      if (itemIds.length > 0) {
        Items.loadItemsBulk(itemIds).then(async () => {
          for (let e of this.editList) {
            if (e.item > 0 && !this.itemData[e.item]) {
              this.itemData[e.item] = await Items.getItem(e.item)
            }
          }
          this.$forceUpdate()
        })
      }
    },

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.addItem) {
        queryState.addItem = this.addItem
      }
      if (this.editMerchantEntrySlot !== 0) {
        queryState.e = this.editMerchantEntrySlot
      }
      if (typeof this.$route.query.return !== "undefined") {
        queryState.return = this.$route.query.return
      }

      this.$router.push(
        {
          path: util.format(ROUTE.MERCHANT_EDIT, this.editMerchantId),
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (typeof this.$route.params.id !== 'undefined' && this.$route.params.id !== 0) {
        this.editMerchantId = parseInt(this.$route.params.id);
      }
      if (typeof this.$route.query.e !== 'undefined' && parseInt(this.$route.query.e) !== 0) {
        this.editMerchantEntrySlot = parseInt(this.$route.query.e);
      }
      if (typeof this.$route.query.addItem !== 'undefined' && this.$route.query.addItem) {
        this.addItem = true;
      }
      if (typeof this.$route.query.return !== 'undefined' && this.$route.query.return) {
        this.returnQuery = this.$route.query.return;
      }
    },

    /**
     * Merchant List entry editing functions
     */
    async deleteMerchantListFromEdit(m) {
      if (confirm(`Are you sure you want to delete this Merchant? (${m.merchantid}) with (${m.slot}) items?`)) {
        await Merchants.deleteMerchant(m.merchantid)
        this.refreshMerchantlistEntries()
      }
    },
    addItemToMerchantListQueue() {
      this.editMerchantEntry     = {}
      this.editMerchantEntrySlot = 0

      console.log("[MerchantSubEditor] Queue adding new item (sub editor)")
      this.addItem = true
      this.updateQueryState()
    },

    async reorderItems() {
      console.log("[MerchantSubEditor] Reordering items")

      this.applyingChanges = true

      try {

        let startingRewriteSlot = 0
        for (let [i, e] of this.editList.entries()) {

          // when the first slot so happens to be not starting at 1
          if (i === 0 && parseInt(e.slot) !== 1) {
            startingRewriteSlot = 1
          }

          console.log("i [%s] slot [%s]", i, e.slot)

          // gap logic
          if (this.editList[i + 1] && this.editList[i]) {
            const isGap = (this.editList[i + 1].slot - this.editList[i].slot) > 1

            // on our first gap set the first rewrite slot to renumber the rest of the entries
            if (isGap && startingRewriteSlot === 0) {
              startingRewriteSlot = this.editList[i].slot
            }
          }

          if (startingRewriteSlot > 0) {
            let desiredEntry  = JSON.parse(JSON.stringify(this.editList[i]))
            desiredEntry.slot = startingRewriteSlot

            if (this.editList[i].slot !== startingRewriteSlot) {
              await Merchants.deleteMerchantEntry(e.merchantid, e.slot)
              await Merchants.addMerchantListEntry(desiredEntry)
            }
          }

          if (startingRewriteSlot !== 0) {
            startingRewriteSlot++
          }
        }

        // this.notification = "Items reordered!"
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = "Error! " + err.response.data.error
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

      this.resetNotifications()

      try {
        await Merchants.deleteMerchantEntry(row.merchantid, row.slot)

        this.notification = "Entry deleted!"
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = "Error! " + err.response.data.error
        }
      }

      this.editList = this.editList.filter((e) => {
        return e.slot !== row.slot
      })

      this.refreshMerchantlistEntries()
    },
    async refreshMerchantlistEntries() {
      this.applyingChanges = true
      setTimeout(async () => {
        await this.reorderItems()

        // reset so we can reload after deleting entry
        this.editList              = []
        this.editMerchantEntrySlot = 0

        this.init()
      }, 1)
    },

    resetNotifications() {
      this.notification = ""
      this.error = ""
    },

    async addItemToMerchantList(e) {
      const itemId = e.id
      let newSlot  = this.editList[this.editList.length - 1] ? this.editList[this.editList.length - 1].slot + 1 : 1

      this.resetNotifications()

      // error if already added to merchant
      for (let e of this.editList) {
        if (e.item === itemId) {
          this.error = "Error: Item already added to Merchant!"
          return;
        }
      }

      try {
        await Merchants.addItemToMerchant(
          parseInt(this.editMerchantId),
          newSlot,
          itemId
        )

        this.notification = "Item added to Merchant!"
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = "Error! " + err.response.data.error
        }
      }

      this.refreshMerchantlistEntries()
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
  }
}
</script>

<style scoped>

</style>
