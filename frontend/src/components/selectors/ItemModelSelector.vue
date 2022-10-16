<template>
  <div>

    <eq-window-simple>
      <div class="row ">

        <!-- Item Search -->
        <div class="col-lg-4 p-0">
          <div class="input-group">
            <input
              type="text"
              class="form-control ml-2"
              v-model="search"
              v-on:keyup="loadModelsDebounce"
              @enter="loadModels"
              placeholder="Search item name"
            >

            <div class="input-group-append" v-if="findByIcon">
              <span
                :class="'fade-in item-' + findByIcon + ' ml-1'"
                style="border: 1px solid rgba(218, 218, 218, 0.3); border-radius: 7px;"
                title="Search models by icon"
                @click="searchModelsByIcon"
              />
            </div>

          </div>
        </div>

        <!-- Item Slot -->
        <div class="col-lg-3">

          <!-- Input -->
          <select
            class="form-control form-control-prepended list-search"
            v-model.lazy="itemSlotSearch"
            @change="itemTypeSearch = 0; searchByIcon = false; search = ''; loadModels()"
          >
            <option value="0">Slot Filter</option>
            <option v-for="option in itemSlotOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>

          </select>
        </div>

        <!-- Item Type -->
        <div class="col-lg-3">

          <!-- Input -->
          <select
            class="form-control form-control-prepended list-search"
            v-model.lazy="itemTypeSearch"
            @change="itemSlotSearch = 0; search = ''; searchByIcon = false;  loadModels()"
          >
            <option value="0">Type Filter</option>
            <option v-for="option in itemTypeOptions" v-bind:value="option.value">
              {{ option.text }}
            </option>
          </select>
        </div>

        <div class="col-lg-2 col-sm-12">
          <b-button variant="primary" class="btn-dark btn-sm btn-outline-warning" @click="reset">
            <i class="fa fa-eraser mr-1"></i>
            Reset
          </b-button>
        </div>
      </div>
    </eq-window-simple>

    <eq-window-simple
      style="height: 85vh; overflow-y: scroll"
      id="item-model-view-port"
    >

      <span v-if="filteredItemModels && filteredItemModels.length === 0">
        No models found...
      </span>

      <div class="row justify-content-center">
        <div
          v-for="modelId in filteredItemModels"
          :key="modelId"
          @mousedown="selectItemModel(modelId)"
          :class="'m-1 item-model ' +  classIsPulsating(modelId)"
        >
          <span
            :id="'item-model-' + modelId"
            style="filter: drop-shadow(10px 5px 7px #000);"
            :class="'fade-in object-ctn-' + modelId"
            :title="'IT' + modelId"
          >
          </span>
        </div>
      </div>

    </eq-window-simple>

    <app-loader :is-loading="!loaded" padding="4"/>

  </div>
</template>

<script>
import util                  from "util";
import itemSlots             from "@/constants/item-slots.json"
import itemSlotIdFileMapping from "@/constants/item-slot-idfile-mapping.json"
import itemTypes             from "@/constants/item-types.json"
import itemTypesModelMapping from "@/constants/item-type-model-mapping.json"
import PageHeader            from "@/components/layout/PageHeader";
import EqWindow              from "@/components/eq-ui/EQWindow";
import EqWindowSimple        from "@/components/eq-ui/EQWindowSimple";
import EqAssets              from "../../app/eq-assets/eq-assets";
import {debounce}            from "../../app/utility/debounce";
import {Items}               from "../../app/items";
import {ItemApi}           from "../../app/api";
import {SpireApi}          from "../../app/api/spire-api";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";

const MAX_ITEM_IDFILE = 100000;
let itemModels        = [];
let itemModelExists   = {};
let modelFiles        = {};

export default {
  name: "ItemModelSelector",
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      search: "",
      itemSlotSearch: 0,
      itemTypeSearch: 0,
      filteredItemModels: null,
      searchByIcon: false,
      itemSlotOptions: [],
      itemTypeOptions: null,
      loaded: false
    }
  },
  methods: {

    searchModelsByIcon() {
      this.searchByIcon = true
      this.itemTypeSearch = 0
      this.itemSlotSearch = 0

      this.loadModels()
    },

    loadModelsDebounce: debounce(function () {
      this.searchByIcon = false;
      this.loadModels()
    }, 600),

    scrollToSelected() {
      // bring focus to the selected model
      // we queue this on a timeout because elements haven't been rendered yet
      if (this.selectedModel && this.selectedModel.length > 0) {
        setTimeout(() => {
          const container = document.getElementById("item-model-view-port");
          const target    = document.getElementById(util.format("item-model-%s", this.getSelectedModelNoIT()))
          if (container && target) {
            container.scrollTop = target.offsetTop - 300;
          }
        }, 100)
      }
    },

    selectItemModel(modelId) {
      this.$emit("input", util.format("IT%s", modelId));
    },

    getSelectedModelNoIT() {
      return parseInt(this.selectedModel.toString().replace("IT", "").trim())
    },

    classIsPulsating(modelId) {
      return parseInt(modelId) === this.getSelectedModelNoIT() ? 'pulsate' : ''
    },

    // reset to zero state
    reset: function () {
      this.itemSlotSearch = 0;
      this.itemTypeSearch = 0;
      this.loadModels()
    },

    // zero state loader
    loadModels: async function () {

      let searchModels = []
      if (this.search.length > 0) {
        searchModels = await Items.getItemModelsByName(this.search)
      }

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

      // item icon based search
      if (this.searchByIcon) {
        const api = (new ItemApi(...SpireApi.cfg()))
        const r = await api.listItems(
          (new SpireQueryBuilder())
            .where("icon", "=", this.findByIcon)
            .groupBy(["idfile"])
            .get()
        )
        let idFiles = []
        if (r.status === 200) {
          for (let i of r.data) {
            const file = i.idfile.replace("IT", "")

            if (itemModelExists[file]) {
              idFiles.push(file)
            }
          }
        }
        this.filteredItemModels = idFiles
        this.loaded             = true
        return;
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
    this.loaded = false;

    modelFiles = {};

    const r = await EqAssets.getItemModelFileNames()
    r.forEach((fileName) => {
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

    this.scrollToSelected()
  },
  activated() {
    this.scrollToSelected()
  },
  props: {
    selectedModel: {
      type: [Number, String],
      default: "",
      required: true
    },
    findByIcon: {
      type: [Number, String],
      default: "",
      required: false
    },
  },
}
</script>

<style scoped>
.item-model {
  height: auto;
  min-width: 120px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 1px solid rgb(218 218 218 / 30%);
  border-radius: 7px;
}
</style>
