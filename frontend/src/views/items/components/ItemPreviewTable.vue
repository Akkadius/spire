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
          <table id="items-table" class="eq-table eq-highlight-rows" style="display: table;">
            <thead>
            <tr>
              <th style="text-align: center; width: 180px"></th>
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
                    variant="outline-warning"
                  >
                    <i class="ra ra-wrench"></i>
                    Edit
                  </b-button>

                  <b-button
                    @click="editItem(item.id, true)"
                    size="sm"
                    variant="outline-light"
                  >
                    <i class="ra ra-double-team"></i>
                    Clone
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

                  <!--                    <img-->
                  <!--                      :src="itemCdnUrl + 'item_' + item.icon + '.png'"-->
                  <!--                      style="height:40px; border-radius: 25px; width:auto;"-->
                  <!--                      class="mr-2"-->
                  <!--                    >-->
                  <span class="ml-2" style="position:relative; top: -15px">{{ item.name }}</span>

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
                    <eq-item-card-preview
                      :item-data="item"
                      :show-related-data="true"
                    />
                  </eq-window>
                </b-popover>

              </td>

              <td>{{ item.reqlevel }}</td>
              <td>{{ item.ac }}</td>
              <td>{{ item.hp }}</td>
              <td>{{ item.mana }}</td>
              <td>{{ item.endur }}</td>
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
import {Items}                                   from "@/app/items";
import EqWindow                                  from "@/components/eq-ui/EQWindow.vue";
import {App}                                     from "@/constants/app";
import {DB_CLASSES_ICONS}                        from "@/app/constants/eq-class-icon-constants";
import {DB_CLASSES_SHORT, DB_CLASSES_WEAR_SHORT} from "@/app/constants/eq-classes-constants";
import {ROUTE}                                   from "@/routes";
import * as util                                 from "util";
import Tablesort                                 from "@/app/utility/tablesort.js";

export default {
  name: "ItemPreviewTable",
  components: {
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
    }
  },
  async created() {
    // do this once so we're not triggering vue re-renders in the loop
    this.sideLoadedItemData = Items.data
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
