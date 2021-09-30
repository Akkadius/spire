<template>
  <div>
<!--    <page-header title="Item Icon Viewer" pre-title="Search and view item icons..."/>-->

    <!-- CONTENT -->
    <div>
      <div class="container-fluid">

        <eq-window title="Icons" class="mt-5 text-center">
          <div class="row mb-4">

            <!-- Item Slot -->
            <div class="col-5">

              <!-- Input -->
              <select
                class="form-control list-search"
                v-model.lazy="iconSlotSearch"
                @change="doIconSlotSearch()"
              >
                <option value="0">Select Slot Filter</option>

                <option v-for="option in iconSlotOptions" v-bind:value="option.value">
                  {{ option.text }}
                </option>

              </select>

            </div>

            <!-- Item Type -->
            <div class="col-6">
              <!-- Input -->
              <select
                class="form-control list-search"
                v-model.lazy="iconItemTypeSearch"
                @change="doIconItemTypeSearch()"
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


    </div>
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

export default {
  components: { EqWindow, EqWindowComplex, EqWindowSimple, PageHeader },
  data() {
    return {
      iconSlotSearch: 0,
      iconItemTypeSearch: 0,
      filteredIcons: null,
      iconSlotOptions: null,
      iconItemTypeOptions: null,
      loaded: false,
    }
  },
  methods: {
    reset: function () {
      this.loaded = false;
      this.iconSlotSearch     = 0;
      this.iconItemTypeSearch = 0;
      this.loadModels()
    },

    // icon slot search
    doIconSlotSearch: function () {
      let itemAdded           = {};
      this.iconItemTypeSearch = 0
      this.loaded             = false

      if (!itemSlotIconMapping[this.iconSlotSearch]) {
        return
      }

      let icons = []
      itemSlotIconMapping[this.iconSlotSearch].forEach((icon) => {
        if (iconExists[icon] && !itemAdded[icon]) {
          icons.push(icon)
          itemAdded[icon] = 1
        }
      })

      this.filteredIcons = icons
      this.loaded        = true
    },

    // item type search
    doIconItemTypeSearch: function () {
      let itemAdded       = {};
      this.iconSlotSearch = 0

      this.loaded = false

      if (!itemTypesIconMapping[this.iconItemTypeSearch]) {
        return
      }

      let icons = []
      itemTypesIconMapping[this.iconItemTypeSearch].forEach((icon) => {
        if (iconExists[icon] && !itemAdded[icon]) {
          icons.push(icon)
          itemAdded[icon] = 1
        }
      })

      this.filteredIcons = icons
      this.loaded        = true
    },

    // zero state loader
    loadModels: function () {
      this.loaded        = false
      this.filteredIcons = icons;
      this.loaded        = true
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
      let iconSlotOptions = [];
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
      let iconItemTypeOptions = [];
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
    this.loaded = false;
    this.loadModelMeta()

    setTimeout(() => {
      this.loadModels()
    }, 50);
  }
}
</script>


