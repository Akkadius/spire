<template>
  <div>

    <!-- CONTENT -->
    <div class="container-fluid">
      <div class="panel-body">
        <div class="panel panel-default">


          <eq-window class="mt-5">
            Begin Range
            <input
              type="text"
              class="form-control"
              v-b-tooltip.hover
              v-model="beginRange"
              @change="listItems"
              placeholder="Begin range"
            >
            End Range
            <input
              type="text"
              class="form-control"
              v-b-tooltip.hover
              v-model="endRange"
              @change="listItems"
              placeholder="Begin range"
            >
            Limit
            <input
              type="text"
              class="form-control"
              v-b-tooltip.hover
              v-model="limit"
              @change="listItems"
              placeholder="Limit"
            >
          </eq-window>

          <div class="row pt-4">
            <div v-for="(item, index) in items"
                 :key="index"
                 style="display: inline-block; vertical-align: top"
                 class="col-4 mb-6">
              <eq-window style="width: auto; height: 100%">
                <eq-item-preview :item-data="item"/>
              </eq-window>
            </div>
          </div>

        </div>

      </div>
    </div>

  </div>
</template>

<script type="ts">
import {ItemApi}        from "@/app/api/api";
import EqWindow         from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemPreview    from "@/components/eq-ui/EQItemPreview.vue";
import * as util        from "util";

export default {
  components: {
    EqItemPreview,
    EqWindow,
    "test-form": () => import("@/components/forms/TasksForm"),
    "task-activity": () => import("@/components/forms/TaskActivitiesForm"),
    "page-header": () => import("@/views/layout/PageHeader")
  },
  data() {
    return {
      items: null,
      limit: 100,
      beginRange: 10000,
      endRange: 100000,
    }
  },

  mounted() {
    this.listItems()
  },
  methods: {
    listItems: function () {
      const api = (new ItemApi(SpireApiClient.getOpenApiConfig()))

      let filters = [
        ["id", "_lte_", this.endRange],
        ["id", "_gte_", this.beginRange],
        ["hp", "_gte_", 500],
        // ["name", "_like_", "kmra"],
        // ["augrestrict", "_gte_", 1],
        // ["augtype", "_gte_", 1],
        // ["augtype", "_lt_", 65536],
        // ["bardtype", "_gte_", 1],
        // ["proceffect", "_gte_", 1],
        // ["extradmgamt", "_gte_", 1],
        // ["skillmodvalue", "_gte_", 1],
      ]

      let wheres = [];
      filters.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheres.push(where)
      })

      api.listItems({limit: this.limit, where: wheres.join(".")}).then((result) => {
        if (result.status === 200) {
          this.items = result.data
        }
      })
    }
  }
}

</script>
