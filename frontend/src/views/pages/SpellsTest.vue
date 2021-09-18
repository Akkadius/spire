<template>
  <div>

    <!-- CONTENT -->
    <div class="container-fluid">
      <div class="panel-body">
        <div class="panel panel-default">

          <eq-window class="mt-5">

            <div style="display: inline-block" v-for="(icon, index) in dbClassIcons" class="mb-3">
              <div class="text-center p-1">
                {{ dbClassesShort[index] }}
                <div class="text-center">
                  <img
                    @click="selectClass(index)"
                    :src="itemCdnUrl + 'item_' + icon + '.png'"
                    :style="'width:auto;' + (isClassSelected(index) ? 'border: 2px solid #dadada; border-radius: 7px;' : '')"
                    class="mt-1 p-1">
                </div>
              </div>

            </div>

            <div class="row mt-4">

              <div class="col-lg-3 col-sm-12 text-center">
                Spell Name or ID
                <input
                  name="spell_name"
                  type="text"
                  class="form-control"
                  v-on:keyup.enter="triggerState"
                  v-model="spellName"
                  placeholder="Name or ID"
                  autofocus=""
                  id="spell_name"
                  value="">
              </div>

              <div class="col-lg-3 col-sm-12 text-center">
                Spell Effect SPA
                <input
                  name="spell_effect"
                  class="form-control"
                  placeholder="Spell Effect SPA #"
                  v-model="spellEffect"
                  type="text"
                  title="SPA # or description e.g. Root, Stun, Mesmerize, Cure, Heal, HoT, DD, DoT, Proc, Snare, Pacify, Timer 3"
                  value=""
                >
              </div>

              <div class="col-lg-3 col-sm-12 text-center">
                Class
                <select name="class" id="Class" class="form-control" v-model="selectedClass" @change="selectClass(selectedClass)">
                  <option value="0">All</option>

                  <option v-for="(eqClass, eqClassId) in dbClasses" v-bind:value="eqClassId">
                    {{ eqClass }}
                  </option>
                </select>
              </div>


            </div>
            <app-loader :is-loading="!loaded" padding="4"/>
          </eq-window>



          <div class="row" style="justify-content: center" v-if="loaded">
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
import {SpellsNewApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemPreview from "@/components/eq-ui/EQItemPreview.vue";
import * as util from "util";
import EqSpellPreview from "@/components/eq-ui/EQSpellPreview.vue";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {App} from "@/constants/app";
import {DB_CLASSES_SHORT, DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";

const SPELLS_LIST_ROUTE = "/spells-test";

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
      loaded: false,
      spells: null,
      limit: 1000,
      beginRange: 10000,
      endRange: 100000,
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,
      dbClasses: DB_PLAYER_CLASSES,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,

      // form values
      selectedClass: 0,
      spellName: "",
      spellEffect: "",
    }
  },

  mounted() {
    if (Object.keys(this.$route.query).length !== 0) {
      this.loadQueryState()
      this.listSpells()
    }

    if (Object.keys(this.$route.query).length === 0) {
      this.loaded = true;
    }
  },
  methods: {

    updateQueryState: function () {
      let queryState = {};
      Object.assign(queryState, this.$route.query)
      queryState.class = this.selectedClass
      queryState.name = this.spellName
      queryState.effect = this.spellEffect

      this.$router.push(
        {
          path: SPELLS_LIST_ROUTE,
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState: function () {
      console.log("Loading query state");

      if (this.$route.query.class) {
        this.selectedClass = this.$route.query.class;
      }
    },

    selectClass: function (eqClass) {
      this.selectedClass = eqClass;
      this.spellName = ""
      this.spellEffect = ""
      this.updateQueryState();
      this.listSpells()
    },

    triggerState() {
      this.updateQueryState();
      this.listSpells()
    },

    isClassSelected: function (eqClass) {
      return eqClass === this.selectedClass;
    },

    listSpells: function () {
      this.loaded = false;

      const api = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
      let filters = [];

      if (this.selectedClass > 0) {
        filters.push(["classes" + this.selectedClass, "_gte_", "1"]);
        filters.push(["classes" + this.selectedClass, "_lte_", "250"]);
      }

      if (this.spellName) {
        filters.push(["name", "_like_", this.spellName]);
      }

      let wheres = [];
      filters.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheres.push(where)
      })

      api.listSpellsNews({limit: this.limit, where: wheres.join(".")}).then((result) => {
        if (result.status === 200) {
          this.spells = result.data
          this.loaded = true;
        }
      })
    }
  }
}

</script>
