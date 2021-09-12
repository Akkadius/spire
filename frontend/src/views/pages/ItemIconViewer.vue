<template>
  <div>
<!--    <page-header title="Item Icon Viewer" pre-title="Search and view item icons..."/>-->

    <!-- CONTENT -->
    <div>
      <div class="container-fluid">




        <eq-window title="Icons" v-lazy-container="{ selector: 'img' }" class="mt-5 text-center">
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

          <span v-for="icon in filteredIcons" :key="icon">

            <img :src="initialLoad === false ? '' : icon" :data-src="icon" style="height: 40px" :id="slug(icon)"
                 class="fade-in p-1">

            <!-- Popover -->
            <b-popover
              :target="slug(icon)"
              placement="bottom"
              variant="light"
              triggers="hover focus"
            >
<!--              <template v-slot:title>Icon</template>-->

              <table>
                <tr>
                  <td class="mr-3"><b>Icon</b></td>
                  <td>{{ getIconFromUrl(icon) }}</td>
                </tr>
              </table>

            </b-popover>

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
import {IconViewerStore} from "@/app/store/iconViewerStore";
import slugify from "slugify";
import PageHeader from "@/views/layout/PageHeader";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";
import EqWindowComplex from "@/components/eq-ui/EQWindowComplex";
import EqWindow from "@/components/eq-ui/EQWindow";
import {App} from "@/constants/app";

const baseUrl     = App.ASSET_ITEM_ICON_BASE_URL;
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
      initialLoad: false
    }
  },
  methods: {
    reset: function () {
      this.iconSlotSearch     = 0;
      this.iconItemTypeSearch = 0;
      this.loadModels()
    },

    // slugify
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },

    // get icon from url
    getIconFromUrl: function (url) {
      return url.split("item_icons/item_")[1].split(".png")[0]
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
          icons.push(baseUrl + util.format("item_%s.png", icon))
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
          icons.push(baseUrl + util.format("item_%s.png", icon))
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
      this.initialLoad   = true
    },

    // load meta data
    loadModelMeta: function () {

      if (typeof IconViewerStore.raceImages !== "undefined" && Object.keys(IconViewerStore.icons).length > 0) {
        icons                    = IconViewerStore.icons;
        modelFiles               = IconViewerStore.modelFiles;
        iconExists               = IconViewerStore.iconExists;
        this.iconSlotOptions     = IconViewerStore.iconSlotOptions
        this.iconItemTypeOptions = IconViewerStore.iconItemTypeOptions

        this.loaded = true
        return
      }

      // Preload model files
      modelFiles = {};
      ItemIcons[0].contents.forEach((row) => {
        const pieces   = row.name.split(/\//);
        const fileName = pieces[pieces.length - 1];

        modelFiles[fileName] = 1
      })

      // Preload icons
      icons = [];
      for (let iconId = 0; iconId <= MAX_ICON_ID; iconId++) {
        const modelKey    = util.format("item_%s.png", iconId);
        const modelExists = modelFiles[modelKey]

        if (modelExists) {
          icons.push(baseUrl + modelKey)
          iconExists[iconId] = 1
        }
      }

      // Item Slot
      this.iconSlotOptions = [];
      for (let slot = 0; slot <= 19; slot++) {
        const slotDescription = itemSlots[slot][0];
        const slotNumbers     = itemSlots[slot][1];

        this.iconSlotOptions.push(
          {
            text: slotDescription,
            value: slotNumbers
          }
        )
      }

      // Item Type
      this.iconItemTypeOptions = [];
      for (const [type, description] of Object.entries(itemTypes)) {
        this.iconItemTypeOptions.push(
          {
            text: description,
            value: type
          }
        )
      }

      // Store
      IconViewerStore.modelFiles          = modelFiles
      IconViewerStore.icons               = icons
      IconViewerStore.iconExists          = iconExists
      IconViewerStore.iconSlotOptions     = this.iconSlotOptions
      IconViewerStore.iconItemTypeOptions = this.iconItemTypeOptions

    }
  },
  async mounted() {
    this.loadModelMeta()

    setTimeout(() => {
      this.loadModels()
    }, 100);
  }
}
</script>

