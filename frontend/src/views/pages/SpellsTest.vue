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
              @change="listSpells"
              placeholder="Begin range"
            >
            End Range
            <input
              type="text"
              class="form-control"
              v-b-tooltip.hover
              v-model="endRange"
              @change="listSpells"
              placeholder="Begin range"
            >
            Limit
            <input
              type="text"
              class="form-control"
              v-b-tooltip.hover
              v-model="limit"
              @change="listSpells"
              placeholder="Limit"
            >
          </eq-window>
          
          <div class="row" style="justify-content: center">
            <div v-for="(spell, index) in spells" :key="spell.id" style="display: inline-block; vertical-align: top">
              <eq-window style="margin-right: 10px; width: auto; height: 90%">
                <eq-spell-preview :spell-data="spell"/>
              </eq-window>
            </div>
          </div>

        </div>

      </div>
    </div>

  </div>
</template>

<script type="ts">
import {SpellsNewApi}   from "@/app/api/api";
import EqWindow         from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemPreview    from "@/components/eq-ui/EQItemPreview.vue";
import * as util        from "util";
import EqSpellPreview   from "@/components/eq-ui/EQSpellPreview.vue";

export default {
  components: {
    EqSpellPreview,
    EqItemPreview,
    EqWindow,
    "test-form": () => import("@/components/forms/TasksForm"),
    "task-activity": () => import("@/components/forms/TaskActivitiesForm"),
    "page-header": () => import("@/views/layout/PageHeader")
  },
  data() {
    return {
      spells: null,
      limit: 100,
      beginRange: 10000,
      endRange: 100000,
    }
  },

  mounted() {
    this.listSpells()
  },
  methods: {
    listSpells: function () {
      const api = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))

      let filters = [
        ["id", "_lte_", this.endRange],
        ["id", "_gte_", this.beginRange],

        // summoned items
        // ["effect_base_value2", "_gte_", 1],
        // ["effectid2", "_eq_", 32],


        // reagent
        // ["components1", "_gte_", 1],

        ["classes1", "gte", "1"],


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

      api.listSpellsNews({limit: this.limit, where: wheres.join(".")}).then((result) => {
        if (result.status === 200) {
          this.spells = result.data
        }
      })
    }
  }
}

</script>
