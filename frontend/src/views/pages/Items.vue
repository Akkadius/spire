<template>
  <div>
    <!-- CONTENT -->
    <div class="container-fluid">
      <div class="panel-body">
        <div class="panel panel-default">

          <eq-window class="mt-5">

            <div class="row">
              <div class="col-12">
                <class-bitmask-calculator
                  :centered-buttons="false"
                  :display-all-none="true"
                  @fired="selectClass()"
                  :inputData.sync="selectedClasses"
                  :mask="selectedClasses"/>
              </div>
            </div>

            <div class="row">
              <div class="col-12">
                <race-bitmask-calculator
                  :centered-buttons="false"
                  :display-all-none="true"
                  @fired="selectRaces()"
                  :inputData.sync="selectedRaces"
                  :mask="selectedRaces"/>
              </div>
            </div>

            <div class="row">
              <div class="col-12">
                <deity-bitmask-calculator
                  :centered-buttons="false"
                  :display-all-none="true"
                  @fired="selectDeities()"
                  :inputData.sync="selectedDeities"
                  :mask="selectedDeities"/>
              </div>
            </div>

            <div class="row">
              <div class="col-12">
                <inventory-slot-calculator
                  :skip-duplicate-slots="true"
                  :display-all-none="true"
                  @fired="selectSlots()"
                  :inputData.sync="selectedSlots"
                  :mask="selectedSlots"/>
              </div>
            </div>

            <div class="row mt-4">

              <div class="col-lg-2 col-sm-12 text-center">
                Item Name or ID
                <input
                  name="item_name"
                  type="text"
                  class="form-control"
                  v-on:keyup.enter="triggerState"
                  v-model="itemName"
                  placeholder="Name or ID"
                  autofocus=""
                  id="item_name"
                  value="">
              </div>

              <div class="col-lg-2 col-sm-12 text-center">
                Item Type
                <select
                  id="item_type"
                  class="form-control"
                  v-model="itemType"
                  @change="triggerState()">
                  <option value="-1">-- Select --</option>
                  <option v-for="option in itemTypeOptions" v-bind:value="option.value">
                    {{ option.text }}
                  </option>
                </select>
              </div>

              <div class="col-lg-2 col-sm-12 text-center">
                Level
                <select
                  class="form-control"
                  v-model="selectedLevel"
                  @change="triggerState()">
                  <option value="0">-- Select --</option>
                  <option v-for="l in 105" v-bind:value="l">
                    {{ l }}
                  </option>
                </select>
              </div>

              <div class="col-lg-2 col-sm-12" v-if="selectedLevel">
                <b-form-group>
                  <b-form-radio v-model="selectedLevelType" @change="triggerStateDelayed()" value="0">Only
                  </b-form-radio>
                  <b-form-radio v-model="selectedLevelType" @change="triggerStateDelayed()" value="1">And Higher
                  </b-form-radio>
                  <b-form-radio v-model="selectedLevelType" @change="triggerStateDelayed()" value="2">And Lower
                  </b-form-radio>
                </b-form-group>
              </div>

              <div class="col-lg-2 col-sm-12 text-center">
                List Type
                <select
                  id="Class"
                  class="form-control"
                  v-model="listType"
                  @change="triggerState()">
                  <option value="table">Table</option>
                  <option value="card">Cards</option>
                </select>
              </div>

            </div>

            <div class="row mt-2" v-if="1 < 0">
              <div class="col-12">
                <h6 class="eq-header">Filters</h6>
              </div>
            </div>
            <div class="row" v-for="filter in 3" :key="filter" v-if="1 < 0">
              <div class="col-12" v-if="itemFields">
                <div class="input-group w-50">
                  <div class="input-group-prepend">
                    <select class="form-control">
                      <option v-for="field in itemFields" :key="field">{{ field }}</option>
                    </select>
                  </div>
                  <select class="form-control w-15">
                    <option v-for="field in filterOptions" :key="field">{{ field }}</option>
                  </select>
                  <div class="input-group-append">
                    <input type="text" class="form-control">
                  </div>
                </div>
              </div>
            </div>

            <div class="row">
              <div class="form-group">
                <button class='eq-button' @click="resetForm()">Reset Form</button>
              </div>
            </div>

            <app-loader :is-loading="!loaded" padding="4"/>
          </eq-window>

          <!-- card rendering -->
          <div class="row" style="justify-content: center" v-if="loaded && listType === 'card'">
            <div v-for="(item, index) in items"
                 class="col-lg-4 col-sm-9"
                 :key="item.id"
                 style="display: inline-block; vertical-align: top">
              <eq-window style="margin-right: 10px; width: auto; height: 90%">
                <eq-item-card-preview :item-data="item"/>
              </eq-window>
            </div>
          </div>

          <!--          <eq-spell-preview-table :items="items" v-if="loaded && listType === 'table' && items"/>-->

        </div>

      </div>
    </div>

  </div>
</template>

<script type="ts">
import {ItemApi}                             from "@/app/api/api";
import EqWindow                              from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient}                      from "@/app/api/spire-api-client";
import * as util                             from "util";
import EqItemCardPreview                     from "@/components/eq-ui/EQItemCardPreview.vue";
import {DB_CLASSES_ICONS}                    from "@/app/constants/eq-class-icon-constants";
import {App}                                 from "@/constants/app";
import {DB_CLASSES_SHORT, DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_SPA}                              from "@/app/constants/eq-spell-constants";
// import EqItemCardPreviewTable                   from "@/components/eq-ui/EQItemCardPreviewTable.vue";
import {Items}                               from "@/app/items";
import {ROUTE}                               from "@/routes";
import ClassBitmaskCalculator                from "@/components/tools/ClassBitmaskCalculator.vue";
import RaceBitmaskCalculator                 from "@/components/tools/RaceBitmaskCalculator.vue";
import InventorySlotCalculator               from "@/components/tools/InventorySlotCalculator.vue";
import DeityBitmaskCalculator                from "@/components/tools/DeityCalculator.vue";
import itemTypes                             from "@/constants/item-types.json"

export default {
  components: {
    DeityBitmaskCalculator,
    InventorySlotCalculator,
    RaceBitmaskCalculator,
    ClassBitmaskCalculator,
    // EqItemCardPreviewTable,
    EqItemCardPreview,
    EqItemCardPreview,
    EqWindow,
    "test-form": () => import("@/components/forms/TasksForm"),
    "task-activity": () => import("@/components/forms/TaskActivitiesForm"),
    "page-header": () => import("@/views/layout/PageHeader")
  },
  data() {
    return {
      loaded: false,
      items: null,
      limit: 100,
      beginRange: 10000,
      endRange: 100000,
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,
      dbClasses: DB_PLAYER_CLASSES,
      dbItemEffects: DB_SPA,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,

      // form values
      selectedClasses: "0",
      selectedRaces: "0",
      selectedDeities: "0",
      selectedSlots: "0",
      itemName: "",
      itemType: -1,
      spellEffect: "",

      selectedLevel: 0,
      selectedLevelType: 0,

      listType: "card",

      itemFields: [],
      filterOptions: ["=", "<=", ">="],

      itemTypeOptions: [],
    }
  },

  async mounted() {
    if (Object.keys(this.$route.query).length !== 0) {
      this.loadQueryState()
      //Items.preloadDbstr().then((res) => {
      this.listItems()
      //})
    }

    if (Object.keys(this.$route.query).length === 0) {
      // Items.preloadDbstr()
      this.loaded = true;
    }

    this.itemFields = await this.getItemFields()

    this.itemTypeOptions = [];
    for (const [type, description] of Object.entries(itemTypes)) {
      this.itemTypeOptions.push(
        {
          text: type + ") " + description,
          value: type
        }
      )
    }
  },
  methods: {

    updateQueryState: function () {
      let queryState = {};

      if (this.selectedClasses !== "0") {
        queryState.classes = this.selectedClasses
      }
      if (this.selectedRaces !== "0") {
        queryState.races = this.selectedRaces
      }
      if (this.selectedDeities !== "0") {
        queryState.deities = this.selectedDeities
      }
      if (this.selectedSlots !== "0") {
        queryState.slots = this.selectedSlots
      }
      if (this.itemType !== 0) {
        queryState.itemType = this.itemType
      }
      if (this.listType !== 0) {
        queryState.listType = this.listType
      }
      if (this.itemName !== "") {
        queryState.name = this.itemName
      }
      if (this.selectedLevel !== 0) {
        queryState.level = this.selectedLevel
      }
      if (this.selectedLevelType !== 0) {
        queryState.levelType = this.selectedLevelType
      }

      this.$router.push(
        {
          path: ROUTE.ITEMS_LIST,
          query: queryState
        }
      ).catch(() => {
      })
    },

    resetForm: function () {
      this.selectedClasses   = "0";
      this.selectedRaces     = "0";
      this.selectedSlots     = "0";
      this.selectedDeities   = "0";
      this.itemType          = -1;
      this.itemName          = "";
      this.spellEffect       = "";
      this.selectedLevel     = 0;
      this.selectedLevelType = 0;
      this.items             = null;
      this.updateQueryState()
    },

    loadQueryState: function () {
      if (this.$route.query.classes) {
        this.selectedClasses = this.$route.query.classes;
      }
      if (this.$route.query.races) {
        this.selectedRaces = this.$route.query.races;
      }
      if (this.$route.query.deities) {
        this.selectedDeities = this.$route.query.deities;
      }
      if (this.$route.query.slots) {
        this.selectedSlots = this.$route.query.slots;
      }
      if (this.$route.query.itemType) {
        this.itemType = this.$route.query.itemType;
      }
      if (this.$route.query.name) {
        this.itemName = this.$route.query.name;
      }
      if (this.$route.query.level) {
        this.selectedLevel = this.$route.query.level;
      }
      if (this.$route.query.levelType) {
        this.selectedLevelType = this.$route.query.levelType;
      }
      if (this.$route.query.listType) {
        this.listType = this.$route.query.listType;
      }
    },

    selectClass: function () {
      this.itemName = ""
      this.updateQueryState();
      this.listItems()
    },

    selectRaces: function () {
      this.itemName = ""
      this.updateQueryState();
      this.listItems()
    },

    selectDeities: function () {
      this.itemName = ""
      this.updateQueryState();
      this.listItems()
    },

    selectSlots: function () {
      this.itemName = ""
      this.updateQueryState();
      this.listItems()
    },

    triggerStateDelayed() {
      setTimeout(() => {
        this.triggerState()
      }, 100)
    },

    triggerState() {
      this.updateQueryState();
      this.listItems()
    },

    async getItemFields() {
      const api     = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      let request   = {};
      request.limit = 1;
      const result  = await api.listItems(request)
      if (result.status === 200 && result.data.length === 1) {
        let fields = []
        Object.keys(result.data[0]).forEach((key) => {
          fields.push(key)
        })
        return fields.sort()
      }

      return [];
    },

    isClassSelected: function (eqClass) {
      return eqClass === this.selectedClasses;
    },

    listItems: function () {
      this.loaded = false;

      const api   = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      let filters = [];
      let whereOr = [];

      // filter by class
      if (this.selectedClasses && parseInt(this.selectedClasses) > 0 && parseInt(this.selectedClasses) !== 65535) {
        filters.push(["classes", "_bitwiseand_", this.selectedClasses]);
        filters.push(["classes", "_ne_", 65535]);
      } else if (this.selectedClasses && parseInt(this.selectedClasses) > 0 && parseInt(this.selectedClasses) === 65535) {
        filters.push(["classes", "__", 65535]);
      }

      // filter by race
      if (this.selectedRaces && parseInt(this.selectedRaces) > 0 && parseInt(this.selectedRaces) !== 65535) {
        filters.push(["races", "_bitwiseand_", this.selectedRaces]);
        filters.push(["races", "_ne_", 65535]);
      } else if (this.selectedRaces && parseInt(this.selectedRaces) > 0 && parseInt(this.selectedRaces) === 65535) {
        filters.push(["races", "__", 65535]);
      }

      // filter by deity
      if (this.selectedDeities && parseInt(this.selectedDeities) > 0 && parseInt(this.selectedDeities) !== 65535) {
        filters.push(["deity", "_bitwiseand_", this.selectedDeities]);
        filters.push(["deity", "_ne_", 65535]);
      } else if (this.selectedDeities && parseInt(this.selectedDeities) > 0 && parseInt(this.selectedDeities) === 65535) {
        filters.push(["deity", "__", 65535]);
      }

      // filter by slots
      if (this.selectedSlots && parseInt(this.selectedSlots) > 0 && parseInt(this.selectedSlots) !== 65535) {
        filters.push(["slots", "_bitwiseand_", this.selectedSlots]);
        filters.push(["slots", "_ne_", 65535]);
      } else if (this.selectedSlots && parseInt(this.selectedSlots) > 0 && parseInt(this.selectedSlots) === 65535) {
        filters.push(["slots", "__", 65535]);
      }

      // item type
      if (this.itemType && this.itemType > 0) {
        filters.push(["itemtype", "__", this.itemType]);
      }

      // level
      if (this.selectedLevel > 0) {
        let filterType = "__"; // equal
        if (parseInt(this.selectedLevelType) === 1) {
          filterType = "_gte_";
        }
        if (parseInt(this.selectedLevelType) === 2) {
          filterType = "_lte_";
        }

        filters.push(["reqlevel", filterType, this.selectedLevel]);
      }

      // if number, filter by id
      // else name
      if (!isNaN(this.itemName) && this.itemName) {
        filters.push(["id", "__", this.itemName]);
      } else if (this.itemName) {
        filters.push(["name", "_like_", this.itemName]);
      }

      let wheres = [];
      filters.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheres.push(where)
      })

      let wheresOrs = [];
      whereOr.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheresOrs.push(where)
      })

      let request   = {};
      request.limit = this.limit;

      // filter by class
      if (this.selectedClasses > 0) {
        // request.orderBy = util.format("classes", this.selectedClasses)
      }

      if (Object.keys(wheres).length > 0) {
        request.where = wheres.join(".")
      }

      if (Object.keys(wheresOrs).length > 0) {
        request.whereOr = wheresOrs.join(".")
      }

      api.listItems(request).then(async (result) => {
        if (result.status === 200) {
          // set items to be rendered
          this.items = result.data

          // fetch spell ids that might be referenced by effects to bulk preload
          let itemsToPreload = [];
          result.data.forEach((item) => {
            // for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
            //   const spellId = Items.getItemIdFromEffectIfExists(spell, effectIndex);
            //   if (spellId > 0) {
            //     itemsToPreload.push(spellId);
            //   }
            //   const itemId = Items.getItemIdFromEffectIfExists(spell, effectIndex);
            //   if (itemId > 0) {
            //     itemsToPreload.push(itemId);
            //   }
            // }
          })

          // fetch spell ids that might be referenced by effects to bulk preload
          result.data.forEach((item) => {
            // for (let i = 0; i < 4; i++) {
            //   if (spell["components_" + i] > 0) {
            //     itemsToPreload.push(spell["components_" + i])
            //   }
            // }
          });

          // bulk fetch preload
          const response = await (new ItemApi(SpireApiClient.getOpenApiConfig())).getItemsBulk({
            body: {
              ids: itemsToPreload
            }
          })

          if (response.status == 200 && response.data && parseInt(response.data.length) > 0) {
            response.data.forEach((item) => {
              Items.setItem(item.id, item);
            })
          }

          // items bulk fetch preload
          api.getItemsBulk({
            body: {
              ids: itemsToPreload
            }
          }).then((response) => {
            if (response.status == 200 && response.data && parseInt(response.data.length) > 0) {
              response.data.forEach((item) => {
                Items.setItem(item.id, item);
              })
            }
            this.loaded = true;
          });

        }
      })
    }
  }
}

</script>
