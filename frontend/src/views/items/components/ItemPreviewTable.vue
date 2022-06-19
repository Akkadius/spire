<template>
  <div>
    <div
      class='eq-window-simple text-center'
      v-if="items.length === 0"
    >
      No items were found
    </div>

    <div
      class='eq-window'
      :style="'margin-bottom: 40px; min-height: 275px; ' + (title ? 'padding-top: 30px' : 'padding-top: 0px !important')"
      v-if="items.length > 0"
    >
      <div class='eq-window-title-bar' v-if="title">{{ title }}</div>
      <div :style="' ' + (title ? 'margin-top: 10px' : 'margin-top: 30px') ">

        <div class='item-table' v-if="items.length > 0">
          <!--                    <div class="ml-3">Items shown ({{items.length}})</div>-->

          <!--        <div class='eq-window-nested-blue' v-if="items.length > 0" style="overflow-y: scroll;">-->
          <table id="items-table" class="eq-table bordered eq-highlight-rows" style="display: table;">
            <thead>
            <tr>
              <th style="text-align: center; width: 100px"></th>
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
              <td>

                <div class="btn-group" role="group">

                  <b-button
                    @click="editItem(item.id)"
                    size="sm"
                    title="Edit"
                    variant="outline-warning"
                  >
                    <i class="ra ra-wrench"></i>
                  </b-button>

                  <b-button
                    @click="editItem(item.id, true)"
                    size="sm"
                    title="Clone"
                    variant="outline-light"
                  >
                    <i class="ra ra-double-team"></i>
                  </b-button>

                </div>

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
  </div>
</template>

<script>
import EqWindow                                  from "@/components/eq-ui/EQWindow.vue";
import {DB_CLASSES_ICONS}                        from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT, DB_CLASSES_WEAR_SHORT} from "@/app/constants/eq-classes-constants";
import {ROUTE}                                   from "@/routes";
import * as util                                 from "util";
import Tablesort                                 from "@/app/utility/tablesort.js";
import ItemPopover                               from "@/components/ItemPopover";

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
    }
  },
  mounted() {
    if (this.items.length > 0) {
      setTimeout(() => {
        if (document.getElementById('items-table')) {
          new Tablesort(document.getElementById('items-table'));
        }
      }, 1000)
    }
  },
  props: {
    items: Array
  },
  methods: {
    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },
    editItem(itemId, clone = false) {
      this.$router.push(
        {
          path: util.format(ROUTE.ITEM_EDIT + (clone ? "?clone" : ""), itemId),
          query: {}
        }
      ).catch(() => {
      })
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
