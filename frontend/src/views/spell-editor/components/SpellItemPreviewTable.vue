<template>
  <div>
    <div
      class='eq-window-simple text-center'
      v-if="items && items.length === 0"
    >
      No items were found
    </div>

    <eq-window-simple
      class="mt-3 p-0"
      v-if="items.length > 0"
      style="overflow-x: scroll; height: 85vh; overflow-y: scroll"
    >
      <div
        class='item-table'
        v-if="items.length > 0"
      >
        <table
          id="spell-items-table-selector"
          class="eq-table eq-highlight-rows"
          style="display: table; overflow-x: scroll"
        >
          <thead>
          <tr>
            <th style="text-align: center; width: 50px"></th>
            <th style="text-align: center; width: 50px" class="text-center">Id</th>
            <th style="text-align: center;">Name</th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="(item, index) in items"
            :class="(isItemSelected(item) ? 'pulsate-highlight-white' : '')"
            :key="item.id"
            :id="'item-selection-row-' + item.id">
            <td>
              <div class="btn-group" role="group">
                <b-button
                  class="btn-dark btn-sm btn-outline-warning"
                  @click="selectItem(item)"
                >
                  Select
                </b-button>
              </div>
            </td>
            <td>
              {{ item.id }}
            </td>
            <td class="text-left" style="vertical-align: middle">
              <div :id="item.id + '-popover'" style="display:inline-block; ">
                <span
                  :class="'fade-in item-' + item.icon" :title="item.icon"
                  style="height: 40px; width: 40px; display: inline-block"
                />
                <span
                  class="ml-2"
                  style="position:relative; top: -15px"
                >{{ item.name }}</span>
              </div>

              <b-popover
                :target="item.id + '-popover'"
                placement="auto"
                custom-class="no-bg"
                delay="1"
                triggers="hover focus"
                style="width: 500px !important"
              >
                <eq-window style="margin-right: 10px; width: auto; height: 90%">
                  <eq-item-card-preview :item-data="item"/>
                </eq-window>
              </b-popover>

            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </eq-window-simple>
  </div>
</template>

<script>
import {Items}                                   from "@/app/items";
import EqWindow                                  from "@/components/eq-ui/EQWindow.vue";
import {App}                                     from "@/constants/app";
import {DB_CLASSES_ICONS}                        from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT, DB_CLASSES_WEAR_SHORT} from "@/app/constants/eq-classes-constants";
import Tablesort                                 from "@/app/utility/tablesort.js";
import EqWindowSimple                            from "@/components/eq-ui/EQWindowSimple";

export default {
  name: "SpellItemPreviewTable",
  components: {
    EqWindowSimple,
    EqWindow,
    "eq-item-card-preview": () => import("@/components/eq-ui/EQItemCardPreview.vue"),
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {
      debug: App.DEBUG,
      debugItemEffects: false,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      itemEffectInfo: [],
      itemData: {},
      sideLoadedItemData: {},
      componentId: "",
      reagents: [],
      effectDescription: "",
      recourseLink: "",
      title: "",
      itemMinis: {},
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,

      highlightedItem: 0,
    }
  },
  async created() {
    // do this once so we're not triggering vue re-renders in the loop
    this.sideLoadedItemData = Items.data
  },
  mounted() {
    console.log("mounted")

    if (this.items.length > 0) {
      const target = document.getElementById('spell-items-table-selector')
      if (target) {
        setTimeout(() => {
          new Tablesort(target);
        }, 100)
      }
    }
  },
  props: {
    items: Array
  },
  methods: {
    selectItem(item) {
      this.$emit('input', item);
      this.highlightedItem = item.id
    },

    isItemSelected(item) {
      return item.id === this.highlightedItem
    },

    getClasses(item) {
      let classes      = []
      let classesValue = item.classes
      for (const [key, value] of Object.entries(DB_CLASSES_WEAR_SHORT).reverse()) {
        if (key <= classesValue) {
          classesValue -= key;
          classes.push(value)
        }
      }

      return item.classes >= 65535 ? 'ALL' : classes.join(", ").trim()
    },
  }
}
</script>

<style scoped>
.eq-table tr {
  border-bottom: .4px solid #ffffff1c;
}

.eq-table td {
  padding-top: 5px;
  padding-bottom: 5px;
  border-right: .1px solid #ffffff1c;
  border-left: .1px solid #ffffff1c;
}

.item-table td {
  vertical-align: middle;
  text-align: center;
}

/* For Mobile */
@media screen and (max-width: 540px) {
  .item-table {
    overflow-x: visible;
    overflow-y: scroll !important
  }
}

/* For Tablets */
@media screen and (min-width: 540px) and (max-width: 780px) {
  .item-table {
    overflow-x: visible;
    overflow-y: scroll !important
  }
}
</style>
