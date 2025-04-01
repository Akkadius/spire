<template>
  <div>
    <eq-window-simple title="Item Models" style="margin-bottom: 1px">
      <div class="row">

        <!-- Item Search -->
        <div class="col-5">

          <input
            type="text"
            class="form-control ml-2"
            v-model="search"
            v-on:keyup="triggerStateSearch"
            @enter="triggerState"
            placeholder="Search for item names to find associated models"
          >

        </div>

        <!-- Item Slot -->
        <div class="col-lg-3 col-sm-12">

          <!-- Input -->
          <select
            class="form-control form-control-prepended list-search"
            v-model.lazy="itemSlotSearch"
            @change="itemTypeSearch = 0; search = ''; triggerState()"
          >
            <option value="0">Select Slot Filter</option>
            <option v-for="option in itemSlotOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>

          </select>
        </div>

        <!-- Item Type -->
        <div class="col-lg-3 col-sm-12">

          <!-- Input -->
          <select
            class="form-control form-control-prepended list-search"
            v-model.lazy="itemTypeSearch"
            @change="itemSlotSearch = 0; search = ''; triggerState()"
          >
            <option value="0">Select Item Type Filter</option>
            <option v-for="option in itemTypeOptions" v-bind:value="option.value">
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

    <eq-window class="mt-5 text-center" style="min-height: 500px">

      <!-- loader -->
      <div v-if="!loaded" class="text-center justify-content-center mt-5 mb-5">
        <div class="mb-3">
          {{ renderingImages ? 'Rendering images...' : 'Loading images...' }}
        </div>
        <loader-fake-progress v-if="!loaded && !renderingImages"/>
        <eq-progress-bar :percent="100" v-if="renderingImages"/>
      </div>

      <span v-if="filteredItemModels && filteredItemModels.length === 0">
        No models found...
      </span>

      <div
        v-if="loaded"
        style="height: 80vh; overflow-y: scroll; "
        id="item-viewer-viewport"
        class="row justify-content-center"
      >
        <div
          v-for="item in filteredItemModels"
          :key="item"
          style="min-height: 100px; max-height: 150px;"
          class="m-1 item-model"
        >
          <span
            style="filter: drop-shadow(10px 5px 5px #000);"
            :class="'fade-in object-ctn-' + item"
            :title="'IT' + item"
          ></span>
        </div>

        <div class="col-12 mt-3 text-center">Image Credits @Maudigan</div>

      </div>
    </eq-window>
  </div>
</template>

<script>
import util                  from "util";
import itemSlots             from "@/constants/item-slots.json"
import itemSlotIdFileMapping from "@/constants/item-slot-idfile-mapping.json"
import itemTypes             from "@/constants/item-types.json"
import itemTypesModelMapping from "@/constants/item-type-model-mapping.json"
import slugify               from "slugify";
import PageHeader            from "@/components/layout/PageHeader";
import EqWindow              from "@/components/eq-ui/EQWindow";
import {ROUTE}               from "../../routes";
import EqWindowSimple        from "../../components/eq-ui/EQWindowSimple";
import LoaderFakeProgress    from "../../components/LoaderFakeProgress";
import EqProgressBar         from "../../components/eq-ui/EQProgressBar";
import EqAssets              from "../../app/eq-assets/eq-assets";
import {debounce}            from "../../app/utility/debounce";
import {Items}               from "../../app/items";
import ContentArea           from "../../components/layout/ContentArea";

const MAX_ITEM_IDFILE = 100000;
let itemModels        = [];
let itemModelExists   = {};
let modelFiles        = {};

export default {
  components: { ContentArea, EqProgressBar, LoaderFakeProgress, EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      search: "",

      itemSlotSearch: 0,
      itemTypeSearch: 0,
      filteredItemModels: null,
      itemSlotOptions: [],
      itemTypeOptions: null,
      loaded: false,
      renderingImages: false,
    }
  },
  methods: {

    triggerStateSearch: debounce(function () {
      this.triggerState()
    }, 600),

    // reset to zero state
    reset: function () {
      this.reset          = "";
      this.itemSlotSearch = 0;
      this.itemTypeSearch = 0;
      this.loadModels()
    },

    // when inputs are triggered and state is updated
    updateQueryState: function () {
      let queryState = {};

      if (this.search !== "") {
        queryState.search = this.search
      }
      if (this.itemSlotSearch !== 0) {
        queryState.itemModelSlot = this.itemSlotSearch
      }
      if (this.itemTypeSearch !== 0) {
        queryState.itemModelType = this.itemTypeSearch
      }

      this.$router.push(
        {
          path: ROUTE.ITEM_VIEWER,
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
      if (this.$route.query.search) {
        this.search = this.$route.query.search;
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
    loadModels: async function () {
      let searchModels = []
      if (this.search.length > 0) {
        searchModels = await Items.getItemModelsByName(this.search)
      }

      let curImg    = new Image();
      curImg.src    = '/eq-asset-preview-master/assets/sprites/objects.png';
      curImg.onload = () => {
        this.renderingImages = true

        setTimeout(() => {
          this.renderingImages = false

          // item based idfile search
          if (this.search.length > 0) {
            let idFiles = []
            for (let model of searchModels) {
              if (itemModelExists[model]) {
                idFiles.push(model)
              }
            }

            this.filteredItemModels = idFiles
            this.loaded             = true;
            return
          }

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
        }, 100)
      }
    }
  },
  async mounted() {
    this.loadQueryState()
    this.loaded = false;

    modelFiles = {};
    const files = await EqAssets.getItemModelFileNames()
    files.forEach((fileName) => {
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
  width: 100px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 1px solid rgb(218 218 218 / 30%);
  border-radius: 7px;
}
</style>
