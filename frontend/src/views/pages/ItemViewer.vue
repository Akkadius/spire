<template>
  <div>
    <!--    <page-header title="Item Model Viewer" pre-title="Search and view item models..."/>-->

    <!-- CONTENT -->
    <div>
      <div class="container-fluid">
        <eq-window title="Item Models" class="mt-5 text-center">
          <div class="row mb-4">

            <!-- Item Slot -->
            <div class="col-5">

              <!-- Input -->
              <select
                class="form-control form-control-prepended list-search"
                v-model.lazy="itemSlotSearch"
                @change="doItemSlotSearch()"
              >
                <option value="0">Select Slot Filter</option>

                <option v-for="option in itemSlotOptions" v-bind:value="option.value">
                  {{ option.text }}
                </option>

              </select>
            </div>

            <!-- Item Type -->
            <div class="col-6">

              <!-- Input -->
              <select
                class="form-control form-control-prepended list-search"
                v-model.lazy="itemTypeSearch"
                @change="doItemTypeSearch()"
              >
                <option value="0">Select Item Type Filter</option>

                <option v-for="option in itemTypeOptions" v-bind:value="option.value">
                  {{ option.text }}
                </option>

              </select>
            </div>
            <div class="col-auto">
              <b-button variant="primary" @click="reset">Reset</b-button>
            </div>
          </div>

          <app-loader :is-loading="!loaded" padding="8"/>

          <span v-if="filteredItemModels && filteredItemModels.length === 0">
            No models found...
          </span>

          <div class="row">
          <div v-for="item in filteredItemModels" :key="item"
               class="m-1"
               style="height: auto; min-width: 120px; display:flex; align-items: center; text-align: center; border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;">

            <span :class="'fade-in object-ctn-' + item" :title="'IT' + item"></span>
          </div>
          </div>

          <div class="mt-5">Images courtesy of Maudigan <3</div>
        </eq-window>

      </div>
    </div>

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

    // slugify
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },

    // get graphic from url
    getWeaponGraphicModelFromUrl: function (url) {
      return url.split("objects/CTN_")[1].split(".png")[0]
    },

    // item slot search
    doItemSlotSearch: function () {
      this.loaded         = false
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
    },

    // item type search
    doItemTypeSearch: function () {
      this.loaded         = false
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
    },

    // zero state loader
    loadModels: function () {
      this.loaded             = false;
      this.filteredItemModels = itemModels;
      this.loaded             = true
    }
  },
  async mounted() {
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

      this.itemSlotOptions.push(
        {
          text: slotDescription,
          value: slotNumbers
        }
      )
    }

    // Item Type
    this.itemTypeOptions = [];
    for (const [type, description] of Object.entries(itemTypes)) {
      this.itemTypeOptions.push(
        {
          text: description,
          value: type
        }
      )
    }

    setTimeout(() => {
      this.loadModels()
    }, 100);

  }
}
</script>

