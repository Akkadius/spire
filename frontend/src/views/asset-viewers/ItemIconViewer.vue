<template>
  <div>
    <eq-window-simple title="Icons" style="margin-bottom: 1px">
      <div class="row">

        <!-- Item Search -->
        <div class="col-5">

          <input
            type="text"
            class="form-control ml-2"
            v-model="search"
            v-on:keyup="triggerStateSearch"
            @enter="triggerState"
            placeholder="Search for item names to find associated icons"
          >

        </div>

        <!-- Item Slot -->
        <div class="col-3">
          <select
            class="form-control list-search"
            v-model.lazy="iconSlotSearch"
            @change="iconItemTypeSearch = 0; search = ''; triggerState()"
          >
            <option value="0">Select Slot Filter</option>
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
            @change="iconSlotSearch = 0; search = ''; triggerState()"
          >
            <option value="0">Select Type Filter</option>
            <option v-for="option in iconItemTypeOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>

          </select>
        </div>

        <div class="col-lg-1 col-sm-12">

          <button
            class='btn btn-dark btn-sm mb-1 mr-2 mt-1'
            @click="reset"
          >
            <i class="fa fa-refresh"></i> Reset
          </button>
        </div>
      </div>
    </eq-window-simple>

    <eq-window class="mt-5 text-center">

      <!-- loader -->
      <div v-if="!loaded" class="text-center justify-content-center mt-5 mb-5">
        <div class="mb-3">
          {{ renderingImages ? 'Rendering images...' : 'Loading images...' }}
        </div>
        <loader-fake-progress v-if="!loaded && !renderingImages"/>
        <eq-progress-bar :percent="100" v-if="renderingImages"/>
      </div>

      <span v-if="filteredIcons && filteredIcons.length === 0">
        No icons found...
      </span>

      <div
        v-if="loaded"
        style="height: 75vh; overflow-y: scroll; "
        id="icon-viewer-viewport"
      >
        <span
          v-for="icon in filteredIcons"
          :key="icon"
          :id="'item-' + icon"
          style="margin-right: 5px"
        >
            <span
              :class="'fade-in item-' + icon" :title="icon"
              style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"
            />
          </span>
      </div>
    </eq-window>
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
import {ROUTE}              from "../../routes";
import LoaderFakeProgress   from "../../components/LoaderFakeProgress";
import EqProgressBar        from "../../components/eq-ui/EQProgressBar";
import EqAssets             from "../../app/eq-assets/eq-assets";
import {Items}              from "../../app/items";
import {debounce}           from "../../app/utility/debounce";
import ContentArea          from "../../components/layout/ContentArea";

const MAX_ICON_ID = 10000;
// const MAX_ICON_ID = 1000;

let iconExists = {}
let icons      = [];
let modelFiles = {};

export default {
  components: { ContentArea, EqProgressBar, LoaderFakeProgress, EqWindow, EqWindowComplex, EqWindowSimple, PageHeader },
  data() {
    return {
      search: "",

      iconSlotSearch: 0,
      iconItemTypeSearch: 0,
      filteredIcons: null,
      iconSlotOptions: null,
      iconItemTypeOptions: null,
      loaded: false,
      renderingImages: false,
    }
  },
  methods: {
    triggerStateSearch: debounce(function() {
      this.triggerState()
    }, 600),

    reset: function () {
      this.search             = ""
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
      if (this.search !== "") {
        queryState.search = this.search
      }
      if (this.iconItemTypeSearch !== 0) {
        queryState.iconItemType = this.iconItemTypeSearch
      }

      this.$router.push(
        {
          path: ROUTE.ITEM_ICON_VIEWER,
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
      if (this.$route.query.search) {
        this.search = this.$route.query.search;
      }
    },

    triggerState() {
      this.updateQueryState();
      this.loadIcons()
    },

    async loadIcons() {
      this.loaded = false;

      let searchIcons = []
      if (this.search.length > 0) {
        searchIcons = await Items.getItemIconsByName(this.search)
      }

      // we let the browser download the image first before trying to render the content
      let curImg    = new Image();
      curImg.src    = '/eq-asset-preview-master/assets/sprites/item-icons.png';
      curImg.onload = () => {

        // inform the user we are rendering
        this.renderingImages = true
        setTimeout(() => {
          this.renderingImages = false

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

        }, 100);
      }

    },

    // load meta data
    loadModelMeta: async function () {
      console.time("files");

      // Preload model files
      modelFiles  = {};
      const files = await EqAssets.getItemIcons()
      files.forEach((file) => {
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
    this.loadQueryState()

    this.loaded = false;
    await this.loadModelMeta()

    setTimeout(() => {
      this.loadIcons()
    }, 50);
  }
}
</script>


