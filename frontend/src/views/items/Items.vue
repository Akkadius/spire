<template>
  <content-area>

    <eq-window>
      <div class="row">
        <div class="col-1">
          <div
            class="row" v-for="field in getCheckboxFilters()"
          >
            <div class="col-9 text-right p-0 pr-2 m-0">
              {{ field.description }}
            </div>
            <div class="col-3 text-left p-0">
              <eq-checkbox
                class="mb-2 d-inline-block"
                :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                v-model.number="checkboxFilters[field.field]"
                @input="checkboxFilters[field.field] = $event; triggerCheckboxFilter(field.field, (typeof field.false !== 'undefined' ? field.false : 0))"
              />
            </div>
          </div>
        </div>

        <div class="col-11 pl-4">
          <div class="row">
            <div class="col-12">
              <class-bitmask-calculator
                :centered-buttons="false"
                :display-all-none="true"
                :add-only-button-enabled="true"
                :add-only-state-enabled="selectOnlyClassEnabled"
                @fired="selectClass()"
                @selectOnly="selectOnlyClassEnabled = $event"
                :inputData.sync="selectedClasses"
                :mask="selectedClasses"
              />
            </div>
          </div>

          <div class="row">
            <div class="col-12">
              <race-bitmask-calculator
                :centered-buttons="false"
                :display-all-none="true"
                @fired="selectRaces()"
                :inputData.sync="selectedRaces"
                :mask="selectedRaces"
              />
            </div>
          </div>

          <div class="row">
            <div class="col-12">
              <deity-bitmask-calculator
                :centered-buttons="false"
                :display-all-none="true"
                @fired="selectDeities()"
                :inputData.sync="selectedDeities"
                :mask="selectedDeities"
              />
            </div>
          </div>

          <div class="row">
            <div class="col-12">
              <inventory-slot-calculator
                :skip-duplicate-slots="true"
                :display-all-none="true"
                @fired="selectSlots()"
                :inputData.sync="selectedSlots"
                :mask="selectedSlots"
              />
            </div>
          </div>

          <div class="row mt-3">

            <div class="col-lg-2 col-sm-12 p-0 pr-1 text-center">
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
                value=""
              >
            </div>

            <div class="col-lg-2 col-sm-12 p-0 pr-1 text-center">
              Item Type
              <select
                id="item_type"
                class="form-control"
                v-model="itemType"
                @change="triggerState()"
              >
                <option value="-1">-- Select --</option>
                <option v-for="option in itemTypeOptions" v-bind:value="option.value">
                  {{ option.text }}
                </option>
              </select>
            </div>

            <div class="col-lg-1 col-sm-12 p-0 pr-1 text-center">
              Level
              <select
                class="form-control"
                v-model="selectedLevel"
                @change="triggerState()"
              >
                <option value="0">-- Select --</option>
                <option v-for="l in 105" v-bind:value="l">
                  {{ l }}
                </option>
              </select>
            </div>

            <div class="col-lg-6 col-sm-12 mt-3 pl-0 pr-0">

              <div class="btn-group ml-3" role="group" aria-label="Basic example" v-if="selectedLevel">
                <b-button
                  @click="selectedLevelType = 0; triggerStateDelayed();"
                  size="sm"
                  :variant="(parseInt(selectedLevelType) === 0 ? 'warning' : 'outline-warning')"
                >Only
                </b-button>
                <b-button
                  @click="selectedLevelType = 1; triggerStateDelayed();"
                  size="sm"
                  :variant="(parseInt(selectedLevelType) === 1 ? 'warning' : 'outline-warning')"
                >Higher
                </b-button>
                <b-button
                  @click="selectedLevelType = 2; triggerStateDelayed();"
                  size="sm"
                  :variant="(parseInt(selectedLevelType) === 2 ? 'warning' : 'outline-warning')"
                >Lower
                </b-button>
              </div>

              <div class="btn-group ml-3" role="group" aria-label="Basic example">
                <b-button
                  alt="Display as table"
                  @click="listType = 'table'; triggerState()"
                  size="sm"
                  :variant="(listType === 'table' ? 'warning' : 'outline-warning')"
                ><i class="fa fa-table"></i></b-button>
                <b-button
                  alt="Display as grid"
                  @click="listType = 'card'; triggerState()"
                  size="sm"
                  :variant="(listType === 'card' ? 'warning' : 'outline-warning')"
                ><i class="fa fa-th"></i></b-button>
              </div>

              <div class="btn-group ml-3" role="group" aria-label="Basic example">
                <b-button
                  @click="limit = 10; triggerStateDelayed()"
                  size="sm"
                  :variant="(parseInt(limit) === 10 ? 'warning' : 'outline-warning')"
                >10
                </b-button>
                <b-button
                  @click="limit = 100; triggerStateDelayed()"
                  size="sm"
                  :variant="(parseInt(limit) === 100 ? 'warning' : 'outline-warning')"
                >100
                </b-button>
                <b-button
                  @click="limit = 1000; triggerStateDelayed()"
                  size="sm"
                  :variant="(parseInt(limit) === 1000 ? 'warning' : 'outline-warning')"
                >1000
                </b-button>
              </div>

              <div
                :class="'text-center btn-xs eq-button-fancy ml-3'"
                style="line-height: 25px;"
                @click="resetForm()"
              >
                Reset Form
              </div>
            </div>
          </div>

          <div class="row mt-3">
            <div class="col-12 p-0">
              <db-column-filter
                v-if="itemFields && filters"
                :set-filters="filters"
                @input="handleDbColumnFilters($event);"
                :columns="itemFields"
              />
            </div>
          </div>

        </div>
      </div>

    </eq-window>

    <app-loader :is-loading="!loaded" padding="4"/>

    <!-- card rendering -->
    <div class="row" style="justify-content: center" v-if="loaded && listType === 'card'">
      <div
        v-for="(item, index) in items"
        class="col-lg-4 col-sm-9 mb-3"
        :key="item.id"
        style="display: inline-block; vertical-align: top"
      >
        <eq-window style="margin-right: 10px; width: auto; height: 100%">
          <eq-item-card-preview
            :item-data="item"
            :show-edit="true"
            :show-related-data="true"
          />
        </eq-window>
      </div>
    </div>

    <!-- table -->
    <item-preview-table
      :items="items"
      v-if="loaded && listType === 'table' && items"
    />

    <!--          <eq-spell-preview-table :items="items" v-if="loaded && listType === 'table' && items"/>-->

  </content-area>
</template>

<script type="ts">
import {ItemApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemCardPreview from "@/components/eq-ui/EQItemCardPreview.vue";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT, DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_SPA} from "@/app/constants/eq-spell-constants";
import {ROUTE} from "@/routes";
import ClassBitmaskCalculator from "@/components/tools/ClassBitmaskCalculator.vue";
import RaceBitmaskCalculator from "@/components/tools/RaceBitmaskCalculator.vue";
import InventorySlotCalculator from "@/components/tools/InventorySlotCalculator.vue";
import DeityBitmaskCalculator from "@/components/tools/DeityCalculator.vue";
import itemTypes from "@/constants/item-types.json"
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import ItemPreviewTable from "@/views/items/components/ItemPreviewTable.vue";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import DbColumnFilter from "@/components/DbColumnFilter";
import {DbSchema} from "@/app/db-schema";
import {Zones} from "@/app/zones";
import {Items} from "@/app/items";
import ItemPopover from "@/components/ItemPopover.vue";
import ContentArea from "@/components/layout/ContentArea.vue";

export default {
  components: {
    ContentArea,
    ItemPopover,
    DbColumnFilter,
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

      checkboxFilters: {},
      filters: [],

      selectedLevel: 0,
      selectedLevelType: 0,

      listType: "table",

      itemFields: [],

      itemTypeOptions: [],
    }
  },

  created() {
    this.items = null;
  },

  async mounted() {

    this.resetCheckboxFilters()

    if (Object.keys(this.$route.query).length !== 0) {
      this.loadQueryState()
      this.listItems()
    }

    if (Object.keys(this.$route.query).length === 0) {
      this.loaded = true;
    }

    this.itemFields = await DbSchema.getTableColumns("items")

    this.itemTypeOptions = [];
    for (const [type, description] of Object.entries(itemTypes)) {
      this.itemTypeOptions.push(
        {
          text: type + ") " + description,
          value: type
        }
      )
    }

    Zones.getZones()
  },

  watch: {
    // reset state vars when we navigate
    '$route'() {
      this.loadQueryState()
      this.listItems()
    },
  },

  methods: {

    handleDbColumnFilters(filters) {
      this.filters = filters
      this.updateQueryState()
    },

    getCheckboxFilters() {
      return [
        {
          description: 'Is Magic',
          field: 'magic'
        },
        {
          description: 'No Drop',
          field: 'nodrop',
          true: 0,
          false: 1,
        },
        {
          description: 'FV No Drop',
          field: 'fvnodrop',
        },
        {
          description: 'No Rent',
          field: 'norent',
          true: 0,
          false: 1,
        },
        {
          description: 'Tradeskill Item',
          field: 'tradeskills'
        },
        {
          description: 'Book',
          field: 'book'
        },
        {
          description: 'No Transfer',
          field: 'notransfer'
        },
        {
          description: 'Summoned',
          field: 'summonedflag'
        },
        {
          description: 'Quest',
          field: 'questitemflag'
        },
        {
          description: 'Artifact',
          field: 'artifactflag'
        },
        {
          description: 'No Pet',
          field: 'nopet'
        },
        {
          description: 'Attuneable',
          field: 'attuneable'
        },
        {
          description: 'Stackable',
          field: 'stackable'
        },
        {
          description: 'Potion Belt',
          field: 'potionbelt'
        },
        // {
        //   description: 'Placeable',
        //   field: 'placeable'
        // },
        {
          description: 'Epic Item',
          field: 'epicitem'
        },
        // {
        //   description: 'Arrow Expend',
        //   field: 'expendablearrow'
        // },
        // {
        //   description: 'Heirloom',
        //   field: 'heirloom'
        // },
      ]
    },

    getChangedFilterCount() {
      let setValues = 0
      for (let key in this.checkboxFilters) {
        const value = this.checkboxFilters[key]

        this.getCheckboxFilters().forEach((filter) => {
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
      for (let key in this.checkboxFilters) {
        const value = this.checkboxFilters[key]

        this.getCheckboxFilters().forEach((filter) => {
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
      // console.log("hello")
      // console.log(this.checkboxFilters)

      // delete filter from checkboxFilters if false value set
      for (let key in this.checkboxFilters) {
        const value = this.checkboxFilters[key]
        if (field == key && value === falseValue) {
          delete this.checkboxFilters[key]
        }
      }

      this.updateQueryState()
      this.listItems()

      // if (Object.keys(this.checkboxFilters).length > 0) {
      //   this.listItems()
      // }

      // if no checkboxFilters set, clear items result
      // if (setValues === 0) {
      //   this.items = null
      // }

    },

    updateQueryState: function () {
      let queryState = {};

      if (this.selectedClasses !== 0) {
        queryState.classes = this.selectedClasses
      }
      if (this.selectedRaces !== 0) {
        queryState.races = this.selectedRaces
      }
      if (this.selectedDeities !== 0) {
        queryState.deities = this.selectedDeities
      }
      if (this.selectedSlots !== 0) {
        queryState.slots = this.selectedSlots
      }
      if (this.itemType !== -1) {
        queryState.itemType = this.itemType
      }
      if (this.listType !== "") {
        queryState.listType = this.listType
      }
      if (this.itemName !== "") {
        queryState.name = this.itemName
      }
      if (this.selectedLevel !== 0) {
        queryState.level = this.selectedLevel
      }
      if (this.limit !== 0) {
        queryState.limit = this.limit
      }
      if (this.selectedLevelType >= 0) {
        queryState.levelType = this.selectedLevelType
      }
      if (this.selectOnlyClassEnabled) {
        queryState.classSelectOnly = 1
      }
      if (this.getChangedFilterCount() > 0) {
        queryState.checkboxFilters = JSON.stringify(this.getFiltersNonZeroValues())
      }
      if (this.filters && this.filters.length > 0) {
        queryState.filters = JSON.stringify(this.filters)
      }

      this.$router.push(
        {
          path: ROUTE.ITEMS_LIST,
          query: queryState
        }
      ).catch(() => {
      })
    },

    resetCheckboxFilters() {
      this.checkboxFilters = {}

      // make sure that checkboxFilters that are non-zero defaults are initialized as their
      // zero-values first
      this.getCheckboxFilters().forEach((filter) => {
        if (typeof filter.true !== 'undefined' && filter.true === 0) {
          this.checkboxFilters[filter.field] = 1
        }
      })
    },

    resetForm: function () {
      this.resetCheckboxFilters()

      this.filters           = []
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
      this.updateQueryState()
    },

    loadQueryState: function () {
      if (this.$route.query.classes) {
        this.selectedClasses = parseInt(this.$route.query.classes);
      }
      if (this.$route.query.races) {
        this.selectedRaces = parseInt(this.$route.query.races);
      }
      if (this.$route.query.deities) {
        this.selectedDeities = parseInt(this.$route.query.deities);
      }
      if (this.$route.query.slots) {
        this.selectedSlots = parseInt(this.$route.query.slots);
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
      if (this.$route.query.limit) {
        this.limit = this.$route.query.limit;
      }
      if (this.$route.query.levelType) {
        this.selectedLevelType = this.$route.query.levelType;
      }
      if (this.$route.query.listType) {
        this.listType = this.$route.query.listType;
      }
      if (parseInt(this.$route.query.classSelectOnly) === 1) {
        this.selectOnlyClassEnabled = true;
      }
      if (this.$route.query.checkboxFilters) {
        this.resetCheckboxFilters()
        let checkboxFilters = JSON.parse(this.$route.query.checkboxFilters);
        for (let key in checkboxFilters) {
          this.checkboxFilters[key] = checkboxFilters[key]
        }
      }
      if (this.$route.query.filters) {
        this.filters = JSON.parse(this.$route.query.filters);
      } else {
        this.filters = [];
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
    },

    isClassSelected: function (eqClass) {
      return eqClass === this.selectedClasses;
    },

    listItems: function () {
      this.loaded   = false;
      const builder = new SpireQueryBuilder()
      const api     = (new ItemApi(SpireApiClient.getOpenApiConfig()))

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
        builder.where(key, "=", this.checkboxFilters[key])
      }

      // item type
      if (this.itemType && this.itemType > 0) {
        builder.where("itemtype", "=", this.itemType)
      }

      if (this.filters && this.filters.length > 0) {
        this.filters.forEach((f) => {
          builder.where(f.f, f.o, f.v)
        });
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
      builder.includes(Items.getRelationships())
      api.listItems(builder.get()).then(async (result) => {
        if (result.status === 200) {
          // set items to be rendered
          this.items  = result.data
          this.loaded = true;
        }
      })
    }
  }
}

</script>
