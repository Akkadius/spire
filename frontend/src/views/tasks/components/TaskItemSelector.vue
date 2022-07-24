<template>
  <div>
    <eq-window-simple
      v-if="searchPerformed"
      class="text-center"
    >
      <b-button
        class="btn-dark btn-sm btn-outline-warning"
        @click="searchPerformed = false"
      >
        Return to Search
      </b-button>
    </eq-window-simple>

    <eq-window-simple class="mt-3" v-if="!searchPerformed">
      <div class="row">
        <div class="col-12 pl-4">
          <div class="row">
            <div class="col-12 text-center">
              <class-bitmask-calculator
                :centered-buttons="false"
                :display-all-none="true"
                :add-only-button-enabled="true"
                :add-only-state-enabled="selectOnlyClassEnabled"
                @selectOnly="selectOnlyClassEnabled = $event"
                :inputData.sync="selectedClasses"
                :mask="selectedClasses"
              />
            </div>
          </div>

          <div class="row">
            <div class="col-12 text-center mt-3">
              <race-bitmask-calculator
                :centered-buttons="false"
                :display-all-none="true"
                :inputData.sync="selectedRaces"
                :mask="selectedRaces"
              />
            </div>
          </div>

          <div class="row">
            <div class="col-12 text-center mt-3">
              <inventory-slot-calculator
                :image-size="40"
                :skip-duplicate-slots="true"
                :display-all-none="false"
                :centered-buttons="true"
                :inputData.sync="selectedSlots"
                :mask="selectedSlots"
              />
            </div>
          </div>

          <div class="row mt-3 text-center">
            <div class="col-lg-4 col-sm-12 p-0 pr-1 text-center">
              Item Name or ID
              <input
                name="item_name"
                type="text"
                class="form-control"
                v-model="itemName"
                v-on:keyup.enter="triggerState"
                placeholder="Name or ID"
                autofocus=""
                id="item_name"
                value=""
              >
            </div>

            <div class="col-lg-4 col-sm-12 p-0 pr-1 text-center">
              Item Type
              <select
                id="item_type"
                class="form-control"
                v-model="itemType"
              >
                <option value="-1">-- Select --</option>
                <option v-for="option in itemTypeOptions" v-bind:value="option.value">
                  {{ option.text }}
                </option>
              </select>
            </div>

            <div class="col-lg-4 col-sm-12 p-0 pr-1 text-center">
              Level
              <select
                class="form-control"
                v-model="selectedLevel"
              >
                <option value="0">-- Select --</option>
                <option v-for="l in 105" v-bind:value="l">
                  {{ l }}
                </option>
              </select>
            </div>

          </div>

          <div class="row">
            <div class="col-lg-12 col-sm-12 mt-3 pl-0 pr-0 text-center">
              <div class="btn-group ml-3" role="group" aria-label="Basic example" v-if="selectedLevel">
                <b-button
                  @click="selectedLevelType = 0;"
                  size="sm"
                  :variant="(parseInt(selectedLevelType) === 0 ? 'warning' : 'outline-warning')"
                >Only
                </b-button>
                <b-button
                  @click="selectedLevelType = 1;"
                  size="sm"
                  :variant="(parseInt(selectedLevelType) === 1 ? 'warning' : 'outline-warning')"
                >Higher
                </b-button>
                <b-button
                  @click="selectedLevelType = 2;"
                  size="sm"
                  :variant="(parseInt(selectedLevelType) === 2 ? 'warning' : 'outline-warning')"
                >Lower
                </b-button>
              </div>

              <div class="btn-group ml-3" role="group" aria-label="Basic example">
                <b-button
                  @click="limit = 10"
                  size="sm"
                  :variant="(parseInt(limit) === 10 ? 'warning' : 'outline-warning')"
                >10
                </b-button>
                <b-button
                  @click="limit = 100"
                  size="sm"
                  :variant="(parseInt(limit) === 100 ? 'warning' : 'outline-warning')"
                >100
                </b-button>
                <b-button
                  @click="limit = 1000"
                  size="sm"
                  :variant="(parseInt(limit) === 1000 ? 'warning' : 'outline-warning')"
                >1000
                </b-button>
              </div>

              <b-button
                class="btn-dark btn-sm btn-outline-warning ml-3"
                @click="search()"
              >
                <i class="fa fa-search"></i> Search
              </b-button>

              <div
                :class="'text-center btn-xs eq-button-fancy ml-3'"
                style="line-height: 25px;"
                @click="resetForm()"
              >
                Reset Form
              </div>
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
        </div>
      </div>

    </eq-window-simple>

    <app-loader :is-loading="!loaded" padding="4"/>

    <task-item-preview-table
      :items="items"
      :selected-item-id="selectedItemId"
      @input="bubbleToParent($event)"
      v-if="loaded && items && searchPerformed"
    />

  </div>
</template>

<script type="ts">
import {ItemApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemCardPreview from "@/components/preview/EQItemCardPreview.vue";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT, DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_SPA} from "@/app/constants/eq-spell-constants";
import {Items} from "@/app/items";
import ClassBitmaskCalculator from "@/components/tools/ClassBitmaskCalculator.vue";
import RaceBitmaskCalculator from "@/components/tools/RaceBitmaskCalculator.vue";
import InventorySlotCalculator from "@/components/tools/InventorySlotCalculator.vue";
import DeityBitmaskCalculator from "@/components/tools/DeityCalculator.vue";
import itemTypes from "@/constants/item-types.json"
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import ItemPreviewTable from "@/views/items/components/ItemPreviewTable.vue";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import TaskItemPreviewTable from "@/views/tasks/components/TaskItemPreviewTable";

export default {
  name: "TaskItemSelector",
  components: {
    TaskItemPreviewTable,
    EqWindowSimple,
    ItemPreviewTable,
    EqCheckbox,
    DeityBitmaskCalculator,
    InventorySlotCalculator,
    RaceBitmaskCalculator,
    ClassBitmaskCalculator,
    // EqItemCardPreviewTable,
    EqItemCardPreview,
    EqItemCardPreview,
    EqWindow,
    "page-header": () => import("@/components/layout/PageHeader.vue")
  },
  data() {
    return {
      loaded: false,
      items: [],
      limit: 100,
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,
      dbClasses: DB_PLAYER_CLASSES,
      dbItemEffects: DB_SPA,

      // form values
      selectedClasses: 0,
      selectedRaces: 0,
      selectedDeities: 0,
      selectedSlots: 0,
      itemName: "",
      itemType: -1,
      spellEffect: "",

      // when "only" option is set
      selectOnlyClassEnabled: false,

      filters: {},

      selectedLevel: 0,
      selectedLevelType: 0,

      listType: "table",

      itemFields: [],
      filterOptions: ["=", "<=", ">="],

      itemTypeOptions: [],

      searchPerformed: false,
    }
  },

  async mounted() {
    this.resetFilters()

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

    // if item is passed in, set it as the search context
    if (this.selectedItemId > 0) {
      this.itemName = this.selectedItemId
      this.search()
    }

    this.loaded = true;
  },

  props: {
    selectedItemId: {
      type: Number,
      required: false,
    },
  },

  methods: {

    bubbleToParent(event) {
      this.$emit('input', event);
    },

    search() {
      this.triggerState()
    },

    getChangedFilterCount() {
      let setValues = 0
      for (let key in this.filters) {
        const value = this.filters[key]

        this.formFilters().forEach((filter) => {
          if (filter.field == key && typeof filter.true !== 'undefined' && filter.true == 0 && value === 0) {
            setValues++
          } else if (filter.field == key && typeof filter.true === 'undefined') {
            setValues++
          }
        })
      }

      console.log("[getChangedFilterCount] set values [%s]", setValues)

      return setValues
    },

    getFiltersNonZeroValues() {
      let values = {}
      for (let key in this.filters) {
        const value = this.filters[key]

        this.formFilters().forEach((filter) => {
          if (filter.field == key && typeof filter.true !== 'undefined' && filter.true == 0 && value === 0) {
            values[filter.field] = value
          } else if (filter.field == key && typeof filter.true === 'undefined') {
            values[filter.field] = value
          }
        })
      }

      return values
    },

    triggerCheckboxFilter(field, falseValue) {

      // delete filter from filters if false value set
      for (let key in this.filters) {
        const value = this.filters[key]
        if (field == key && value === falseValue) {
          delete this.filters[key]
        }
      }

      this.listItems()
    },

    resetFilters() {
      this.filters = {}
    },

    resetForm: function () {
      this.resetFilters()

      this.selectedClasses   = 0;
      this.selectedRaces     = 0;
      this.selectedSlots     = 0;
      this.selectedDeities   = 0;
      this.itemType          = -1;
      this.itemName          = "";
      this.spellEffect       = "";
      this.selectedLevel     = 0;
      this.limit             = 100;
      this.selectedLevelType = 0;
      this.items             = null;
      this.listType          = "table"
    },

    triggerStateDelayed() {
      setTimeout(() => {
        this.triggerState()
      }, 100)
    },

    triggerState() {
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

      const api     = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      const builder = new SpireQueryBuilder()

      // filter by class
      if (this.selectedClasses && parseInt(this.selectedClasses) > 0 && parseInt(this.selectedClasses) !== 65535) {
        if (this.selectOnlyClassEnabled) {
          builder.where("classes", "=", this.selectedClasses)
        } else {
          builder.where("classes", "&", this.selectedClasses)
        }

        builder.where("classes", "!=", 65535)
      } else if (this.selectedClasses && parseInt(this.selectedClasses) > 0 && parseInt(this.selectedClasses) === 65535) {
        builder.where("classes", "=", 65535)
      }

      // filter by race
      if (this.selectedRaces && parseInt(this.selectedRaces) > 0 && parseInt(this.selectedRaces) !== 65535) {
        builder.where("classes", "&", this.selectedRaces)
        builder.where("classes", "!=", 65535)
      } else if (this.selectedRaces && parseInt(this.selectedRaces) > 0 && parseInt(this.selectedRaces) === 65535) {
        builder.where("races", "=", 65535)
      }

      // filter by deity
      if (this.selectedDeities && parseInt(this.selectedDeities) > 0 && parseInt(this.selectedDeities) !== 65535) {
        builder.where("deity", "&", this.selectedDeities)
        builder.where("deity", "!=", 65535)
      } else if (this.selectedDeities && parseInt(this.selectedDeities) > 0 && parseInt(this.selectedDeities) === 65535) {
        builder.where("deity", "=", 65535)
      }

      // filter by slots
      if (this.selectedSlots && parseInt(this.selectedSlots) > 0 && parseInt(this.selectedSlots) !== 65535) {
        builder.where("slots", "&", this.selectedSlots)
        builder.where("slots", "!=", 65535)

      } else if (this.selectedSlots && parseInt(this.selectedSlots) > 0 && parseInt(this.selectedSlots) === 65535) {
        builder.where("slots", "=", 65535)
      }

      for (let key in this.getFiltersNonZeroValues()) {
        builder.where(key, "=", this.filters[key])
      }

      // item type
      if (this.itemType && this.itemType > 0) {
        builder.where("itemtype", "=", this.itemType)
      }

      // level
      if (this.selectedLevel > 0) {
        let filterType = "="
        if (parseInt(this.selectedLevelType) === 1) {
          filterType = ">=";
        }
        if (parseInt(this.selectedLevelType) === 2) {
          filterType = "<=";
        }

        builder.where("reqlevel", filterType, this.selectedLevel)
      }

      // if number, filter by id
      // else name
      if (!isNaN(this.itemName) && this.itemName) {
        builder.where("id", "=", this.itemName)
      } else if (this.itemName) {
        builder.where("name", "like", this.itemName)
      }

      if (builder.getFilterCount() === 0) {
        this.items  = null
        this.loaded = true
        return;
      }

      builder.groupBy(["id"])
      builder.limit(this.limit)

      api.listItems(builder.get()).then(async (result) => {
        if (result.status === 200) {
          // set items to be rendered
          this.items           = result.data
          this.searchPerformed = true

          let itemsToPreload = [];

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

            if (this.items.length === 0) {
              this.searchPerformed = false
              this.resetForm()
            }
          });

        }
      })
    }
  }
}

</script>
