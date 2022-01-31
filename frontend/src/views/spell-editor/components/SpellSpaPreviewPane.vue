<template>
  <div v-if="spa >= 0" class="ml-3">

    <!--    <div class="mb-4 mt-3" style="font-size: 14px; color: rgb(123, 113, 74);">-->
    <!--      Depending on the spell effect, the base, limit, max, formula values can mean different things. Below shows what-->
    <!--      values correspond to what for the Spell effect you have highlighted.-->
    <!--    </div>-->

    <h6 class="eq-header mt-3 mb-3">Spell Effect Details</h6>

    <div>
      <div
        :class="'row ' + (field.description === 'Description' ? 'mb-3' : '') + (field.description === 'Notes' ? 'mt-3' : '')"
        v-for="field in fields" :key="field.description"
      >
        <div
          :class="'col-2 text-right m-0 p-0 pr-2 font-weight-bold'"
          v-if="field.field !== ''"
        >
          {{ field.description }}
        </div>
        <div class="col-8 m-0 p-0" v-if="field.field !== ''">
          {{ field.field }}
        </div>
      </div>
    </div>

    <h6 class="eq-header mb-3 mt-3" v-if="effectInfo.info !== ''">Effect Render Preview</h6>

    <v-runtime-template
      :template="'<span>' + effectInfo.info + '</span>'"
      v-if="typeof effectInfo !== 'undefined'"
      class="pb-6 mt-3 doc"
    />

    <div class="mb-5"></div>

    <eq-debug :data="getSpaDefinition(spa)"/>
  </div>
</template>

<script>
import {SPELL_SPA_DEFINITIONS} from "../../../app/constants/eq-spell-spa-definitions";
import EqDebug                 from "../../../components/eq-ui/EQDebug";
import {Spells}                from "../../../app/spells";
import {App}                   from "../../../constants/app";
import {Items}                 from "../../../app/items";
import EqWindow                from "@/components/eq-ui/EQWindow";

export default {
  name: "SpellSpaPreviewPane",
  components: {
    EqDebug,
    "v-runtime-template": () => import("v-runtime-template"),
    EqWindow,
  },
  watch: {
    'effectIndex'() {
      this.load()
    }
  },
  created() {
    this.load()
  },
  data() {
    return {
      fields: [],
      effectInfo: "",
      sideLoadedSpellData: {},

      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,

      componentId: "",
      reagents: [],

      itemData: null,
    }
  },
  props: {
    spa: {
      type: Number,
      required: true
    },
    spell: {
      type: Object,
      required: true,
    },
    effectIndex: {
      type: Number,
      required: true,
    }
  },
  methods: {

    toTitleCase(str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    },

    getSpaDefinition(spa) {
      return SPELL_SPA_DEFINITIONS[spa] ? SPELL_SPA_DEFINITIONS[spa] : {}
    },
    async load() {
      this.fields = [
        { field: this.spa, description: "SPA" },
        { field: this.getSpaDefinition(this.spa).effectName, description: "Effect Name" },
        { field: this.getSpaDefinition(this.spa).category, description: "Category" },
        { field: this.getSpaDefinition(this.spa).description, description: "Description" },
        { field: this.toTitleCase(this.getSpaDefinition(this.spa).base), description: "Base" },
        { field: this.toTitleCase(this.getSpaDefinition(this.spa).limit), description: "Limit" },
        { field: this.toTitleCase(this.getSpaDefinition(this.spa).value), description: "Max" },
        { field: this.getSpaDefinition(this.spa).notes, description: "Notes" },
      ]

      this.effectInfo = await this.getSpellEffectInfo(this.spell, this.effectIndex);

      // side loaded data
      this.sideLoadedSpellData = Spells.data;
      this.itemData            = Items.items;
    },
    getSpellEffectInfo: async function (spell, effectIndex) {
      return await Spells.getSpellEffectInfo(spell, effectIndex)
    }
  }
}
</script>
