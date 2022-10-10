<template>
  <div>

    <!-- Inputs -->
    <eq-window-simple>
      <div class="row">

        <!-- Item Slot -->
        <div class="col-4 p-0">
          <div class="input-group">
            <input
              type="text"
              class="form-control ml-2"
              v-model="search"
              v-on:keyup="searchDebounce"
              @enter="loadIcons"
              placeholder="Search by item name"
            >
          </div>

        </div>

        <!-- Item Slot -->
        <div class="col-3">
          <select
            class="form-control list-search"
            v-model.lazy="iconSlotSearch"
            @change="iconItemTypeSearch = 0; searchByModel = false; search = ''; loadIcons()"
          >
            <option value="0">Slot Filter</option>
            <option v-for="option in iconSlotOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>
          </select>
        </div>

        <!-- Item Type -->
        <div class="col-3">
          <select
            class="form-control list-search"
            v-model.lazy="iconItemTypeSearch"
            @change="iconSlotSearch = 0; searchByModel = false; search = ''; loadIcons()"
          >
            <option value="0">Type Filter</option>
            <option v-for="option in iconItemTypeOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>

          </select>
        </div>

        <div class="col-1">
          <b-button
            class="btn-dark btn-sm btn-outline-warning"
            @click="reset"
          >
            Reset
          </b-button>
        </div>

      </div>

    </eq-window-simple>

    <eq-window v-if="findByModel" class="text-center p-3">
      <div class="mb-3 font-weight-bold">Click to search by item model</div>
      <div style="border: 1px solid rgba(218, 218, 218, 0.3); border-radius: 7px;" @click="searchIconByModel">
          <span
            :class="'mt-2 mb-2 fade-in object-ctn-' + findByModel.replace('IT', '')"
            style="filter: drop-shadow(rgb(0, 0, 0) 10px 5px 7px);"
          />
      </div>
    </eq-window>

    <!-- Content -->
    <eq-window-simple
      :style="'height: ' + (findByModel ? 70 : 85) + 'vh; overflow-y: scroll; overflow-x: hidden'"
      class="text-center"
      id="item-icon-view-port"
      v-if="filteredIcons && filteredIcons.length > 0"
    >

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
          :title="icon"
          style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"
        />
      </span>
    </eq-window-simple>

    <app-loader :is-loading="!loaded" padding="4"/>
  </div>
</template>

<script>
import util                 from "util";
import itemSlots            from "@/constants/item-slots.json"
import itemTypes            from "@/constants/item-types.json"
import itemSlotIconMapping  from "@/constants/item-slot-icon-mapping.json"
import itemTypesIconMapping from "@/constants/item-type-icon-mapping.json"
import PageHeader           from "@/components/layout/PageHeader";
import EqWindowSimple       from "@/components/eq-ui/EQWindowSimple";
import EqWindowComplex      from "@/components/eq-ui/EQWindowComplex";
import EqWindow             from "@/components/eq-ui/EQWindow";
import EqAssets             from "../../../app/eq-assets/eq-assets";
import {debounce}           from "../../../app/utility/debounce";
import {Items}              from "../../../app/items";
import {ItemApi}           from "../../../app/api";
import {SpireApi}          from "../../../app/api/spire-api";
import {SpireQueryBuilder} from "../../../app/api/spire-query-builder";

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
      search: "",

      iconSlotSearch: 0,
      iconItemTypeSearch: 0,
      filteredIcons: null,
      iconSlotOptions: null,
      iconItemTypeOptions: null,
      loaded: false,

      searchByModel: false,
    }
  },
  props: {
    selectedIcon: {
      type: Number,
      default: 0,
      required: true
    },
    findByModel: {
      type: [Number, String],
      default: "",
      required: false
    },
  },
  methods: {
    searchIconByModel() {
      console.log("Trigger")
      this.searchByModel  = true
      this.itemTypeSearch = 0
      this.itemSlotSearch = 0
      this.loadIcons()
    },

    searchDebounce: debounce(function () {
      this.searchByModel = false
      this.loadIcons()
    }, 300),

    async init() {
      this.loaded = false;
      await this.loadModelMeta()

      setTimeout(() => {
        this.loadIcons()
      }, 50);

      this.scrollToSelected()
    },

    scrollToSelected() {
      if (this.selectedIcon > 0) {
        setTimeout(() => {
          const container = document.getElementById("item-icon-view-port");
          const target    = document.getElementById(util.format("item-icon-%s", this.selectedIcon))
          if (container && target) {
            container.scrollTop = target.offsetTop - 300;
          }
        }, 100)
      }
    },

    selectIcon(icon) {
      this.$emit("input", parseInt(icon));
    },

    classIsPulsating(iconId) {
      return parseInt(iconId) === parseInt(this.selectedIcon) ? 'pulsate' : ''
    },

    reset: function () {
      this.search             = ""
      this.loaded             = false;
      this.iconSlotSearch     = 0;
      this.iconItemTypeSearch = 0;
      this.loadIcons()
    },

    async loadIcons() {
      this.loaded = false;

      let searchIcons = []
      if (this.search.length > 0) {
        searchIcons = await Items.getItemIconsByName(this.search)
      }

      // item based icon search
      if (this.search.length > 0) {
        let filteredIcons = []
        let itemAdded     = {};
        for (let icon of searchIcons) {
          if (iconExists[icon]) {
            filteredIcons.push(icon)
            itemAdded[icon] = 1
          }
        }

        this.filteredIcons = filteredIcons
        this.loaded        = true;
        return
      }

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

      // item model based search
      if (this.searchByModel) {
        const api         = (new ItemApi(...SpireApi.cfg()))
        const r           = await api.listItems(
          (new SpireQueryBuilder())
            .where("idfile", "=", this.findByModel)
            .groupBy(["icon"])
            .get()
        )
        let filteredIcons = []
        if (r.status === 200) {
          for (let i of r.data) {
            if (iconExists[i.icon]) {
              filteredIcons.push(i.icon)
            }
          }
        }
        this.filteredIcons = filteredIcons
        this.loaded        = true
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
    loadModelMeta: async function () {
      console.time("files");

      // Preload model files
      modelFiles = {};
      const r    = await EqAssets.getItemIcons()
      r.forEach((file) => {
        modelFiles[file] = 1
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
    this.init()
  },
  activated() {
    this.scrollToSelected()
  }
}
</script>


