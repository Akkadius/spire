<template>
  <div>
    <div class="row mb-4">

      <!-- Item Slot -->
      <div class="col-5">
        <select
          class="form-control list-search"
          v-model.lazy="iconSlotSearch"
          @change="iconItemTypeSearch = 0; loadIcons()"
        >
          <option value="0">Select Slot Filter</option>
          <option v-for="option in iconSlotOptions" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select>
      </div>

      <!-- Item Type -->
      <div class="col-5">
        <select
          class="form-control list-search"
          v-model.lazy="iconItemTypeSearch"
          @change="iconSlotSearch = 0; loadIcons()"
        >
          <option value="0">Select Type Filter</option>
          <option v-for="option in iconItemTypeOptions" v-bind:value="option.value">
            {{ option.text }}
          </option>

        </select>
      </div>

      <div class="col-1">
        <b-button
          class="btn-dark btn-sm btn-outline-warning"
          @click="reset">
          Reset
        </b-button>
      </div>
    </div>

    <app-loader :is-loading="!loaded" padding="8"/>

    <div
      style="height: 85vh; overflow-y: scroll"
      id="item-icon-view-port">

    <span v-if="filteredIcons && filteredIcons.length === 0">
      No icons found...
    </span>

      <span
        v-for="icon in filteredIcons" :key="icon" :id="'item-' + icon"
        @mousedown="selectIcon(icon)"
        :class="'p-1'"
      >
        <span
          :id="'item-icon-' + icon"
          :class="'fade-in item-' + icon + ' ' + classIsPulsating(icon)"
          style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"/>
      </span>
    </div>
  </div>
</template>

<script>
import ItemIcons            from "@/app/asset-maps/item-icons-map.json";
import util                 from "util";
import itemSlots            from "@/constants/item-slots.json"
import itemTypes            from "@/constants/item-types.json"
import itemSlotIconMapping  from "@/constants/item-slot-icon-mapping.json"
import itemTypesIconMapping from "@/constants/item-type-icon-mapping.json"
import PageHeader           from "@/components/layout/PageHeader";
import EqWindowSimple       from "@/components/eq-ui/EQWindowSimple";
import EqWindowComplex      from "@/components/eq-ui/EQWindowComplex";
import EqWindow             from "@/components/eq-ui/EQWindow";

const MAX_ICON_ID = 10000;
// const MAX_ICON_ID = 1000;

let iconExists = {}
let icons      = [];
let modelFiles = {};

export default {
  name: "ItemIconSelector",
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
  props: {
    selectedIcon: {
      type: Number,
      default: 0,
      required: true
    },
  },
  methods: {
    selectIcon(icon) {
      this.$emit("input", parseInt(icon));
    },

    classIsPulsating(iconId) {
      return parseInt(iconId) === parseInt(this.selectedIcon) ? 'pulsate' : ''
    },

    reset: function () {
      this.loaded             = false;
      this.iconSlotSearch     = 0;
      this.iconItemTypeSearch = 0;
      this.updateQueryState()
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
    this.loaded = false;
    this.loadModelMeta()

    setTimeout(() => {
      this.loadIcons()
    }, 50);

    // bring focus to the selected model
    // we queue this on a timeout because elements haven't been rendered yet
    console.log(this.selectedIcon)

    if (this.selectedIcon > 0) {
      setTimeout(() => {
        const container = document.getElementById("item-icon-view-port");
        const target    = document.getElementById(util.format("item-icon-%s", this.selectedIcon))
        if (container && target) {
          container.scrollTop = target.offsetTop - 300;
        }
      }, 100)
    }

  }
}
</script>


