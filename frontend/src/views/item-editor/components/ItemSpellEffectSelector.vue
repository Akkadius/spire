<template>
  <div>
    <eq-window style="margin-top: 30px;">
      <div
        class="row text-center"
        style="margin: 0 auto;"
      >
        <div
          v-for="(icon, index) in dbClassIcons"
          class="mb-3 text-center"
        >
          <div class="text-center col-lg-12 p-0 col-sm-12">
            {{ dbClassesShort[index] }}
            <div class="text-center">
              <img
                @click="selectClass(index)"
                :src="itemCdnUrl + 'item_' + icon + '.png'"
                :style="'width:auto; height: 35px; ' + (isClassSelected(index) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px;')"
                class="mt-1 p-0"
                alt=""
              >
            </div>
          </div>
        </div>
      </div>

      <div class="row mt-1">
        <div class="col-3 text-center">
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
            value=""
          >
        </div>

        <div class="col-3 text-center">
          Spell Effect SPA
          <select
            name="class"
            id="spell_effect"
            class="form-control"
            v-model="selectedSpa"
            @change="triggerState()"
          >
            <option value="-1">-- Select --</option>
            <option v-for="(spellEffect, id) in dbSpellEffects" v-bind:value="id">
              {{ id }}) {{ spellEffect }}
            </option>

          </select>

        </div>

        <div class="col-2 text-center">
          Level
          <select
            name="class"
            id="Class"
            class="form-control"
            v-model="selectedLevel"
            @change="selectClass(selectedClass)"
          >
            <option value="0">Select</option>
            <option v-for="l in 105" v-bind:value="l">
              {{ l }}
            </option>
          </select>
        </div>

        <div class="col-2 m-0 p-0">
          <b-form-group>
            <b-form-radio v-model="selectedLevelType" @change="triggerState()" value="0">Only
            </b-form-radio>
            <b-form-radio v-model="selectedLevelType" @change="triggerState()" value="1">And Higher
            </b-form-radio>
            <b-form-radio v-model="selectedLevelType" @change="triggerState()" value="2">And Lower
            </b-form-radio>
          </b-form-group>
        </div>

        <div class="col-1 text-center">
          <b-button
            class="btn-dark btn-sm btn-outline-warning mb-3"
            @click="resetForm"
          >
            Reset
          </b-button>

          <b-button
            class="btn-dark btn-sm btn-outline-warning"
            @click="triggerState"
          >
            Search
          </b-button>
        </div>

      </div>

      <app-loader :is-loading="!loaded" padding="4"/>

      <div v-if="message">
        {{ message }}
      </div>

      <div
        style="height: 75vh; overflow-y: scroll"
        id="spell-effect-selector-view-port"
        v-if="loaded && spells"
      >
        <eq-spell-preview-table-selector
          :spells="spells"
          @input="bubbleToParent($event)"
          v-if="loaded && spells"
        />
      </div>

    </eq-window>
  </div>
</template>

<script type="ts">
import {ItemApi, SpellsNewApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemCardPreview from "@/components/eq-ui/EQItemCardPreview.vue";
import * as util from "util";
import EqSpellPreview from "@/components/eq-ui/EQSpellCardPreview.vue";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {App} from "@/constants/app";
import {DB_CLASSES_SHORT, DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_SPA} from "@/app/constants/eq-spell-constants";
import EqSpellPreviewTable from "@/components/eq-ui/EQSpellPreviewTable.vue";
import {Spells} from "@/app/spells";
import {Items} from "@/app/items";
import EqSpellPreviewTableSelector from "@/components/eq-ui/EQSpellPreviewTableSelector.vue";

export default {
  name: "SpellEffectSelector",
  components: {
    EqSpellPreviewTableSelector,
    EqSpellPreviewTable,
    EqSpellPreview,
    EqItemCardPreview,
    EqWindow,
  },
  data() {
    return {
      loaded: false,
      spells: null,
      limit: 250,
      beginRange: 10000,
      endRange: 100000,
      dbClassIcons: DB_CLASSES_ICONS,
      dbClassesShort: DB_CLASSES_SHORT,
      dbClasses: DB_PLAYER_CLASSES,
      dbSpellEffects: DB_SPA,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,

      // form values
      selectedClass: 0,
      spellName: "",
      spellEffect: "",
      selectedSpa: -1,
      selectedLevel: 0,
      selectedLevelType: 0,

      message: "",

      listType: "table"
    }
  },

  mounted() {
    if (Object.keys(this.$route.query).length === 0) {
      Spells.preloadDbstr()
      this.loaded = true;
    }
  },
  methods: {

    bubbleToParent(event) {
      this.$emit('input', event);
    },
    resetForm: function () {
      this.selectedClass     = 0;
      this.spellName         = "";
      this.spellEffect       = "";
      this.selectedSpa       = -1;
      this.selectedLevel     = 0;
      this.selectedLevelType = 0;
      this.spells            = null;
    },

    selectClass: function (eqClass) {
      this.selectedClass = eqClass;
      this.spellName     = ""
      this.selectedSpa   = -1
      this.listSpells()
    },

    triggerState() {
      this.listSpells()
    },

    isClassSelected: function (eqClass) {
      return eqClass === this.selectedClass;
    },

    listSpells: function () {
      this.loaded  = false;
      this.message = ""

      const api   = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
      let filters = [];
      let whereOr = [];

      // filter by class and no level set
      if (this.selectedClass > 0 && this.selectedLevel === 0) {
        filters.push(["classes" + this.selectedClass, "_gte_", "1"]);
        filters.push(["classes" + this.selectedClass, "_lte_", "250"]);

        // exclude rk 2/3 for now
        filters.push(["name", "_notlike_", "Rk. I"]);
      }

      // filter by level if class set
      if (this.selectedLevel > 0 && this.selectedClass > 0) {
        let filterType = "__"; // equal
        if (parseInt(this.selectedLevelType) === 1) {
          filterType = "_gte_";
        }
        if (parseInt(this.selectedLevelType) === 2) {
          filterType = "_lte_";
        }

        filters.push(["classes" + this.selectedClass, filterType, this.selectedLevel]);
        filters.push(["classes" + this.selectedClass, "_lte_", "250"]);

        // exclude rk 2/3 for now
        filters.push(["name", "_notlike_", "Rk. I"]);
      }

      // if number, filter by id
      // else name
      if (!isNaN(this.spellName) && this.spellName) {
        filters.push(["id", "__", this.spellName]);
      } else if (this.spellName) {
        filters.push(["name", "_like_", this.spellName]);
      }

      if (this.selectedSpa > 0) {
        for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
          whereOr.push(["effectid" + effectIndex, "__", this.selectedSpa]);
        }
      }

      let wheres = [];
      filters.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheres.push(where)
      })

      let wheresOrs = [];
      whereOr.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheresOrs.push(where)
      })

      let request   = {};
      this.message  = "";
      request.limit = this.limit;
      if (this.selectedLevel && this.selectedLevel > 0) {
        request.limit = 1000
        this.message  = ""
      }


      // filter by class
      if (this.selectedClass > 0) {
        request.orderBy = util.format("classes%s", this.selectedClass)
      }

      if (Object.keys(wheres).length > 0) {
        request.where = wheres.join(".")
      }

      if (Object.keys(wheresOrs).length > 0) {
        request.whereOr = wheresOrs.join(".")
      }

      api.listSpellsNews(request).then(async (result) => {
        if (result.status === 200) {
          // set spells to be rendered
          this.spells = result.data

          // fetch spell ids that might be referenced by effects to bulk preload
          let spellsToPreload = [];
          let itemsToPreload  = [];
          result.data.forEach((spell) => {
            for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
              const spellId = Spells.getSpellIdFromEffectIfExists(spell, effectIndex);
              if (spellId > 0) {
                spellsToPreload.push(spellId);
              }
              const itemId = Spells.getItemIdFromEffectIfExists(spell, effectIndex);
              if (itemId > 0) {
                itemsToPreload.push(itemId);
              }
            }
          })

          // fetch spell ids that might be referenced by effects to bulk preload
          result.data.forEach((spell) => {
            for (let i = 0; i < 4; i++) {
              if (spell["components_" + i] > 0) {
                itemsToPreload.push(spell["components_" + i])
              }
            }
          });

          // bulk fetch preload
          const response = await (new ItemApi(SpireApiClient.getOpenApiConfig())).getItemsBulk({
            body: {
              ids: itemsToPreload
            }
          })

          if (response.status == 200 && response.data && parseInt(response.data.length) > 0) {
            response.data.forEach((item) => {
              Items.setItem(item.id, item);
            })
          }

          // spells bulk fetch preload
          api.getSpellsNewsBulk({
            body: {
              ids: spellsToPreload
            }
          }).then((response) => {
            if (response.status == 200 && response.data && parseInt(response.data.length) > 0) {
              response.data.forEach((spell) => {
                Spells.setSpell(spell.id, spell);
              })
            }
            this.loaded = true;
          });

        }
      })
    }
  }
}

</script>
