<template>
  <div>
    <eq-window-simple style="">
      <div
        class="row text-center justify-content-center mb-3"
        style="margin: 0 auto;"
      >
        <div
          v-for="(icon, index) in dbClassIcons"
          class="mb-3 text-center mr-4"
        >
          <div class="text-center col-lg-12 p-0 col-sm-12">
            {{ dbClassesShort[index] }}
            <div class="text-center">
              <div
                style="display: block"
                @click="selectClass(index)"
                :class="'mt-1 hover-highlight-inner item-' + icon + ' ' + (isClassSelected(index) ? 'highlight-selected-inner' : '')"
              />

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

        <div class="col-3 text-center">
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

        <div class="col-3 pl-0">
          <div
            class="btn-group"
            role="group"
            v-if="selectedLevel"
            style="margin-top: 23px"
          >
            <b-button
              @click="selectedLevelType = 0; listSpells();"
              size="sm"
              :variant="(parseInt(selectedLevelType) === 0 ? 'warning' : 'outline-warning')"
            >Only
            </b-button>
            <b-button
              @click="selectedLevelType = 1; listSpells();"
              size="sm"
              :variant="(parseInt(selectedLevelType) === 1 ? 'warning' : 'outline-warning')"
            >Higher
            </b-button>
            <b-button
              @click="selectedLevelType = 2; listSpells();"
              size="sm"
              :variant="(parseInt(selectedLevelType) === 2 ? 'warning' : 'outline-warning')"
            >Lower
            </b-button>
          </div>
        </div>

      </div>

      <div class="row mt-3 text-center">
        <div class="col-12">
          <div class="btn-group ml-3" role="group" aria-label="Basic example">
            <b-button
              @click="limit = 10; listSpells()"
              size="sm"
              :variant="(parseInt(limit) === 10 ? 'warning' : 'outline-warning')"
            >10
            </b-button>
            <b-button
              @click="limit = 100; listSpells()"
              size="sm"
              :variant="(parseInt(limit) === 100 ? 'warning' : 'outline-warning')"
            >100
            </b-button>
            <b-button
              @click="limit = 250; listSpells()"
              size="sm"
              :variant="(parseInt(limit) === 250 ? 'warning' : 'outline-warning')"
            >250
            </b-button>
            <b-button
              @click="limit = 1000; listSpells()"
              size="sm"
              :variant="(parseInt(limit) === 1000 ? 'warning' : 'outline-warning')"
            >1000
            </b-button>
          </div>

          <b-button
            class="btn-dark btn-sm btn-outline-warning ml-3"
            @click="resetForm"
          >
            <i class="fa fa-refresh"></i> Reset
          </b-button>

          <b-button
            class="btn-dark btn-sm btn-outline-warning ml-3"
            @click="triggerState"
          >
            <i class="fa fa-search"></i> Search
          </b-button>
        </div>
      </div>

    </eq-window-simple>

    <app-loader :is-loading="!loaded" padding="4"/>

    <eq-window-simple
      style="overflow-y: scroll; overflow-x: hidden; height: 60vh"
      id="spell-effect-selector-view-port"
      v-if="loaded && spells"
    >
      <div v-if="message">
        {{ message }}
      </div>

      <item-spell-preview-table-selector
        :spells="spells"
        @input="bubbleToParent($event)"
        v-if="loaded && spells"
      />

    </eq-window-simple>
  </div>
</template>

<script type="ts">
import {ItemApi, SpellsNewApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqItemCardPreview from "@/components/preview/EQItemCardPreview.vue";
import * as util from "util";
import EqSpellPreview from "@/components/preview/EQSpellCardPreview.vue";
import {DB_CLASSES_ICONS} from "@/app/constants/eq-class-icon-constants";
import {App} from "@/constants/app";
import {DB_CLASSES_SHORT, DB_PLAYER_CLASSES} from "@/app/constants/eq-classes-constants";
import {DB_SPA} from "@/app/constants/eq-spell-constants";
import EqSpellPreviewTable from "@/components/preview/EQSpellPreviewTable.vue";
import {Spells} from "@/app/spells";
import {Items} from "@/app/items";
import ItemSpellPreviewTableSelector from "@/views/items/components/ItemSpellPreviewTableSelector.vue";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export default {
  name: "SpellEffectSelector",
  components: {
    ItemSpellPreviewTableSelector,
    EqWindowSimple,
    EqSpellPreviewTable,
    EqSpellPreview,
    EqItemCardPreview,
    EqWindow,
  },
  data() {
    return {
      loaded: false,
      spells: null,
      limit: 100,
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

      const api     = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
      const builder = (new SpireQueryBuilder())

      // filter by class and no level set
      if (this.selectedClass > 0 && this.selectedLevel === 0) {
        builder.where("classes" + this.selectedClass, ">=", "1")
        builder.where("classes" + this.selectedClass, "<=", "250")
      }

      let filterType = "="
      if (parseInt(this.selectedLevelType) === 1) {
        filterType = ">=";
      }
      if (parseInt(this.selectedLevelType) === 2) {
        filterType = "<=";
      }

      // filter by level if class set
      if (this.selectedLevel > 0 && this.selectedClass > 0) {
        builder.where("classes" + this.selectedClass, filterType, this.selectedLevel)
        builder.where("classes" + this.selectedClass, "<=", "250")
      }

      // when no class is set but level is greater than 0
      if (this.selectedClass === 0 && this.selectedLevel > 0) {
        for (let i = 1; i < 16; i++) {
          builder.whereOr("classes" + i, filterType, this.selectedLevel)
        }
      }

      // if number, filter by id
      // else name
      if (!isNaN(this.spellName) && this.spellName) {
        builder.whereOr("id", "=", this.spellName)
      } else if (this.spellName) {
        builder.whereOr("name", "like", this.spellName)
      }

      if (this.selectedSpa > 0) {
        for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
          builder.whereOr("effectid" + effectIndex, "=", this.selectedSpa)
        }
      }

      builder.limit(this.limit);

      // filter by class
      if (this.selectedClass > 0) {
        builder.orderBy([util.format("classes%s", this.selectedClass)])
      }

      api.listSpellsNews(builder.get()).then(async (result) => {
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
