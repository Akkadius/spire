<template>
  <div class="container-fluid">
    <eq-window title="Icons" class="mt-5 text-center">
      <div class="row mb-4">

        <!-- Item Slot -->
        <div class="col-5">
          <select
            class="form-control list-search"
            v-model.lazy="iconSlotSearch"
            @change="iconItemTypeSearch = 0; triggerState()"
          >
            <option value="0">Select Slot Filter</option>
            <option v-for="option in iconSlotOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>
          </select>
        </div>

        <!-- Item Type -->
        <div class="col-6">
          <select
            class="form-control list-search"
            v-model.lazy="iconItemTypeSearch"
            @change="iconSlotSearch = 0; triggerState()"
          >
            <option value="0">Select Type Filter</option>
            <option v-for="option in iconItemTypeOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>

          </select>
        </div>

        <div class="col-auto">
          <b-button variant="primary" @click="reset">Reset</b-button>
        </div>
      </div>

      <app-loader :is-loading="!loaded" padding="8"/>

      <span v-if="filteredIcons && filteredIcons.length === 0">
            No icons found...
          </span>

      <span v-for="icon in filteredIcons" :key="icon" :id="'item-' + icon" class="p-1">
            <span
              :class="'fade-in item-' + icon" :title="icon"
              style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"/>
          </span>
    </eq-window>
  </div>
</template>

<script>
import ItemIcons from "@/app/asset-maps/item-icons-map.json";
import util from "util";
import itemSlots from "@/constants/item-slots.json"
import itemTypes from "@/constants/item-types.json"
import itemSlotIconMapping from "@/constants/item-slot-icon-mapping.json"
import itemTypesIconMapping from "@/constants/item-type-icon-mapping.json"
import PageHeader from "@/views/layout/PageHeader";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";
import EqWindowComplex from "@/components/eq-ui/EQWindowComplex";
import EqWindow from "@/components/eq-ui/EQWindow";

const MAX_ICON_ID = 10000;
// const MAX_ICON_ID = 1000;

let iconExists = {}
let icons      = [];
let modelFiles = {};

// move this to central constants later
const ITEM_ICON_VIEWER_ROUTE = "/item-icon-viewer"

export default {
  components: { EqWindow, EqWindowComplex, EqWindowSimple, PageHeader },
  data() {
    return {
      iconSlotSearch: 0,
      iconItemTypeSearch: 0,
      filteredIcons: null,
      iconSlotOptions: null,
      iconItemTypeOptions: null,
      loaded: false
    }
  },
  methods: {
    reset: function () {
      this.loaded             = false;
      this.iconSlotSearch     = 0;
      this.iconItemTypeSearch = 0;
      this.updateQueryState()
      this.loadIcons()
    },

    // when inputs are triggered and state is updated
    updateQueryState: function () {
      let queryState = {};

      if (this.iconSlotSearch !== 0) {
        queryState.iconSlot = this.iconSlotSearch
      }
      if (this.iconItemTypeSearch !== 0) {
        queryState.iconItemType = this.iconItemTypeSearch
      }

      this.$router.push(
        {
          path: ITEM_ICON_VIEWER_ROUTE,
          query: queryState
        }
      ).catch(() => {
      })
    },

    // usually from loading initial state
    loadQueryState: function () {
      if (this.$route.query.iconSlot) {
        this.iconSlotSearch = this.$route.query.iconSlot;
      }
      if (this.$route.query.iconItemType) {
        this.iconItemTypeSearch = this.$route.query.iconItemType;
      }
    },

    triggerState() {
      this.updateQueryState();
      this.loadIcons()
    },

    loadIcons() {
      this.loaded = false;

      // icon slot based search
      if (this.iconSlotSearch > 0) {
        let itemAdded           = {};
        this.iconItemTypeSearch = 0

        if (!itemSlotIconMapping[this.iconSlotSearch]) {
          return
        }

        let filteredIcons = []
        itemSlotIconMapping[this.iconSlotSearch].forEach((icon) => {
          if (iconExists[icon] && !itemAdded[icon]) {
            filteredIcons.push(icon)
            itemAdded[icon] = 1
          }
        })

        this.filteredIcons = filteredIcons
        this.loaded        = true;
        return;
      }

      // icon item type search
      if (this.iconItemTypeSearch > 0) {
        let itemAdded       = {};
        this.iconSlotSearch = 0

        if (!itemTypesIconMapping[this.iconItemTypeSearch]) {
          return
        }

        let filteredIcons = []
        itemTypesIconMapping[this.iconItemTypeSearch].forEach((icon) => {
          if (iconExists[icon] && !itemAdded[icon]) {
            filteredIcons.push(icon)
            itemAdded[icon] = 1
          }
        })

        this.filteredIcons = filteredIcons
        this.loaded        = true;
        return;
      }

      // if no filters or searches were hit, we load them all
      this.filteredIcons = icons
      this.loaded        = true;
    },

    // load meta data
    loadModelMeta: function () {
      console.time("files");

      // Preload model files
      modelFiles = {};
      ItemIcons[0].contents.forEach((row) => {
        const pieces   = row.name.split(/\//);
        const fileName = pieces[pieces.length - 1];

        modelFiles[fileName] = 1
      })

      console.timeEnd("files");
      console.time("icons");

      // Preload icons
      icons = [];
      for (let iconId = 0; iconId <= MAX_ICON_ID; iconId++) {
        const modelKey    = util.format("item_%s.png", iconId);
        const modelExists = modelFiles[modelKey]

        if (modelExists) {
          icons.push(iconId)
          iconExists[iconId] = 1
        }
      }

      console.timeEnd("icons");

      // Item Slot
      this.iconSlotOptions = [];
      let iconSlotOptions  = [];
      for (let slot = 0; slot <= 19; slot++) {
        const slotDescription = itemSlots[slot][0];
        const slotNumbers     = itemSlots[slot][1];

        let itemCountDescription = "";
        if (itemSlotIconMapping[slotNumbers] && itemSlotIconMapping[slotNumbers].length > 0) {
          itemCountDescription = util.format(" (%s icons)", itemSlotIconMapping[slotNumbers].length)
        }

        iconSlotOptions.push(
          {
            text: slotDescription + itemCountDescription,
            value: slotNumbers
          }
        )
      }

      this.iconSlotOptions = iconSlotOptions

      // Item Type
      this.iconItemTypeOptions = [];
      let iconItemTypeOptions  = [];
      for (const [type, description] of Object.entries(itemTypes)) {

        let itemCountDescription = "";
        if (itemTypesIconMapping[type] && itemTypesIconMapping[type].length > 0) {
          itemCountDescription = util.format(" (%s icons)", itemTypesIconMapping[type].length)
        }

        if (itemTypesIconMapping[type].length > 0) {
          iconItemTypeOptions.push(
            {
              text: type + ") " + description + itemCountDescription,
              value: type
            }
          )
        }
      }
      this.iconItemTypeOptions = iconItemTypeOptions
    }
  },
  async mounted() {
    this.loadQueryState()

    this.loaded = false;
    this.loadModelMeta()

    setTimeout(() => {
      this.loadIcons()
    }, 50);
  }
}
</script>


