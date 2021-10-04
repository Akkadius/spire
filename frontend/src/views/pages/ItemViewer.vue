<template>
  <div class="container-fluid">
    <eq-window title="Item Models" class="mt-5 text-center" style="min-height: 500px">
      <div class="row mb-4">

        <!-- Item Slot -->
        <div class="col-lg-5 col-sm-12">

          <!-- Input -->
          <select
            class="form-control form-control-prepended list-search"
            v-model.lazy="itemSlotSearch"
            @change="itemTypeSearch = 0; triggerState()"
          >
            <option value="0">Select Slot Filter</option>
            <option v-for="option in itemSlotOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>

          </select>
        </div>

        <!-- Item Type -->
        <div class="col-lg-6 col-sm-12">

          <!-- Input -->
          <select
            class="form-control form-control-prepended list-search"
            v-model.lazy="itemTypeSearch"
            @change="itemSlotSearch = 0; triggerState()"
          >
            <option value="0">Select Item Type Filter</option>
            <option v-for="option in itemTypeOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>
          </select>
        </div>
        <div class="col-lg-1 col-sm-12">
          <b-button variant="primary" class="form-control mr-1 ml-2 mt-1" @click="reset">
            <i class="fa fa-eraser mr-1"></i>
            Reset
          </b-button>
        </div>
      </div>

      <app-loader :is-loading="!loaded" padding="8"/>

      <span v-if="filteredItemModels && filteredItemModels.length === 0">
            No models found...
          </span>

      <div class="row justify-content-center">
        <div v-for="item in filteredItemModels" :key="item" class="m-1 item-model">
          <span :class="'fade-in object-ctn-' + item" :title="'IT' + item"></span>
        </div>
      </div>

    </eq-window>
  </div>
</template>

<script>
import ItemModels from "@/app/asset-maps/objects-map.json";
import util from "util";
import itemSlots from "@/constants/item-slots.json"
import itemSlotIdFileMapping from "@/constants/item-slot-idfile-mapping.json"
import itemTypes from "@/constants/item-types.json"
import itemTypesModelMapping from "@/constants/item-type-model-mapping.json"
import slugify from "slugify";
import PageHeader from "@/views/layout/PageHeader";
import {App} from "@/constants/app";
import EqWindow from "@/components/eq-ui/EQWindow";

const baseUrl         = App.ASSET_CDN_BASE_URL + "assets/objects/";
const MAX_ITEM_IDFILE = 100000;

let itemModels      = [];
let itemModelExists = {};
let modelFiles      = {};

// move this to central constants later
const ITEM_VIEWER_ROUTE = "/item-viewer"

export default {
  components: { EqWindow, PageHeader },
  data() {
    return {
      itemSlotSearch: 0,
      itemTypeSearch: 0,
      filteredItemModels: null,
      itemSlotOptions: [],
      itemTypeOptions: null,
      loaded: false
    }
  },
  methods: {

    // reset to zero state
    reset: function () {
      this.itemSlotSearch = 0;
      this.itemTypeSearch = 0;
      this.loadModels()
    },

    // when inputs are triggered and state is updated
    updateQueryState: function () {
      let queryState = {};

      if (this.itemSlotSearch !== 0) {
        queryState.itemModelSlot = this.itemSlotSearch
      }
      if (this.itemTypeSearch !== 0) {
        queryState.itemModelType = this.itemTypeSearch
      }

      this.$router.push(
        {
          path: ITEM_VIEWER_ROUTE,
          query: queryState
        }
      ).catch(() => {
      })
    },

    // usually from loading initial state
    loadQueryState: function () {
      if (this.$route.query.itemModelSlot) {
        this.itemSlotSearch = this.$route.query.itemModelSlot;
      }
      if (this.$route.query.itemModelType) {
        this.itemTypeSearch = this.$route.query.itemModelType;
      }
    },

    triggerState() {
      this.updateQueryState();
      this.loadModels()
    },

    // slugify
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },

    // get graphic from url
    getWeaponGraphicModelFromUrl: function (url) {
      return url.split("objects/CTN_")[1].split(".png")[0]
    },

    // zero state loader
    loadModels: function () {

      // filter by item type
      if (this.itemTypeSearch > 0) {
        this.itemSlotSearch = 0;
        if (!itemTypesModelMapping[this.itemTypeSearch]) {
          return
        }

        let idFiles = []
        itemTypesModelMapping[this.itemTypeSearch].forEach((idFile) => {
          const file = idFile.replace("IT", "")

          if (itemModelExists[file]) {
            idFiles.push(file)
          }
        })

        this.filteredItemModels = idFiles
        this.loaded             = true
        return
      }

      // item slot search
      if (this.itemSlotSearch) {
        this.itemTypeSearch = 0;

        if (!itemSlotIdFileMapping[this.itemSlotSearch]) {
          return
        }

        let idFiles = []
        itemSlotIdFileMapping[this.itemSlotSearch].forEach((idFile) => {
          const file = idFile.replace("IT", "")

          if (itemModelExists[file]) {
            idFiles.push(file)
          }
        })

        this.filteredItemModels = idFiles
        this.loaded             = true
        return;
      }

      // fallback - load everything
      this.filteredItemModels = itemModels;
      this.loaded             = true;
    }
  },
  async mounted() {
    this.loadQueryState()
    this.loaded = false;

    modelFiles = {};
    ItemModels[0].contents.forEach((row) => {
      const pieces   = row.name.split(/\//);
      const fileName = pieces[pieces.length - 1];

      modelFiles[fileName] = 1
    })

    itemModels = [];

    for (let itemId = 0; itemId <= MAX_ITEM_IDFILE; itemId++) {
      const modelKey = util.format("CTN_%s.png", itemId);

      if (modelFiles[modelKey]) {
        itemModels.push(itemId)
        itemModelExists[itemId] = 1
      }
    }

    // slot
    for (let slot = 0; slot <= 19; slot++) {
      const slotDescription = itemSlots[slot][0];
      const slotNumbers     = itemSlots[slot][1];

      let modelCountDescription = "";
      if (itemSlotIdFileMapping[slotNumbers] && itemSlotIdFileMapping[slotNumbers].length > 0) {
        modelCountDescription = util.format(" (%s models)", itemSlotIdFileMapping[slotNumbers].length)
      }

      this.itemSlotOptions.push(
        {
          text: slotDescription + modelCountDescription,
          value: slotNumbers
        }
      )
    }

    // Item Type
    this.itemTypeOptions = [];
    for (const [type, description] of Object.entries(itemTypes)) {

      let modelCountDescription = "";
      if (itemTypesModelMapping[type] && itemTypesModelMapping[type].length > 0) {
        modelCountDescription = util.format(" (%s models)", itemTypesModelMapping[type].length)
      }

      if (itemTypesModelMapping[type].length > 0) {
        this.itemTypeOptions.push(
          {
            text: type + ") " + description + modelCountDescription,
            value: type
          }
        )
      }
    }

    setTimeout(() => {
      this.loadModels()
    }, 100);

  }
}
</script>

<style scoped>
.item-model {
  height:          auto;
  min-width:       120px;
  display:         flex;
  justify-content: center;
  align-items:     center;
  border:          1px solid rgb(218 218 218 / 30%);
  border-radius:   7px;
}
</style>
