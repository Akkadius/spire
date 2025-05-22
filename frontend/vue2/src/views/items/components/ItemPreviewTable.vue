<template>
  <div>
    <div
      class='eq-window-simple text-center mb-0'
      v-if="items.length === 0"
    >
      No items were found
    </div>

    <div
      class='eq-window-simple p-0'
      v-if="items && items.length > 0"
      style="transition: all 1s ease;"
    >
      <div
        class='item-table fill-screen'
        style="overflow-y: scroll; overflow-x: hidden;"
        v-if="items.length > 0"
      >
        <table id="items-table" class="eq-table bordered eq-highlight-rows">
          <thead class="eq-table-floating-header">
          <tr>
            <th style="text-align: center; width: 120px"></th>
            <th style="text-align: center; width: 100px" class="text-center">Id</th>
            <th style="text-align: center; width: auto;">Name</th>
            <th style="text-align: center; width: auto;">ReqLvl</th>
            <th style="text-align: center; width: auto;">AC</th>
            <th style="text-align: center; width: auto;">HP</th>
            <th style="text-align: center; width: auto;">Mana</th>
            <th style="text-align: center; width: auto;">End</th>
            <th style="width: auto;">Classes</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(item, index) in items" :key="item.id">
            <td class="p-0 text-center">

              <b-button
                variant="primary"
                size="sm"
                style="width: 28px; height: 28px"
                class="btn-outline-danger mr-2"
                title="Delete"
                @click="deleteItem(item)"
              >
                <i class="fa fa-trash"></i>
              </b-button>

              <router-link
                :to="ROUTE.ITEM_EDIT.replace('%s', item.id)"
                size="sm"
                tag="button"
                style="width: 28px; height: 28px"
                title="Edit"
                class="btn btn-sm btn-outline-success mr-2"
              >
                <i class="fa fa-pencil-square"></i>
              </router-link>

              <router-link
                :to="ROUTE.ITEM_EDIT.replace('%s', item.id) + '?clone=true'"
                size="sm"
                tag="button"
                style="width: 30px; height: 28px"
                title="Clone"
                class="btn btn-sm btn-outline-light mr-2"
              >
                <i class="ra ra-double-team"></i>
              </router-link>

            </td>
            <td>
              {{ item.id }}
            </td>
            <td class="text-left" style="vertical-align: middle">
              <item-popover
                :item="item"
                v-if="Object.keys(item).length > 0 && item"
                size="regular"
              />
            </td>

            <td>{{ item.reqlevel }}</td>
            <td>{{ item.ac }}</td>
            <td>{{ commify(item.hp) }}</td>
            <td>{{ commify(item.mana) }}</td>
            <td>{{ commify(item.endur) }}</td>
            <td class="text-left">{{ getClasses(item) }}</td>

          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT, DB_CLASSES_WEAR_SHORT} from "@/app/constants/eq-classes-constants";
import {ROUTE} from "@/routes";
import Tablesort from "@/app/utility/tablesort.js";
import ItemPopover from "@/components/ItemPopover";
import {Items} from "@/app/items";
import {WindowManager} from "@/app/window";

export default {
  name: "ItemPreviewTable",
  components: {
    ItemPopover,
    EqWindow,
  },
  data() {
    return {
      title: "",
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,
      ROUTE: ROUTE,
    }
  },
  mounted() {
    if (this.items.length > 0) {
      setTimeout(() => {
        if (document.getElementById('items-table')) {
          new Tablesort(document.getElementById('items-table'));
        }
      }, 1000)

      setTimeout(() => {
        WindowManager.resizeFillScreenElements()
      }, 100);
    }
  },
  props: {
    items: Array
  },
  methods: {

    async deleteItem(item) {
      if (confirm(`Are you sure you want to permanently delete this item? [${item.name}] (${item.id})`)) {
        await Items.deleteItem(item.id)
        this.$emit("reload-list", true);
      }
    },

    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },
    getClasses(item) {
      let classes = []
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
