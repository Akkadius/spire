<template>
  <div>
    <div v-if="spellEffectInfo.length > 0 && loaded">
      <div v-for="effect in spellEffectInfo">
        <v-runtime-template
          :template="'<span>' + effect + '</span>'"
          v-if="typeof effect !== 'undefined'"
          class="pb-6 mt-3 doc"/>
      </div>
    </div>
    <app-loader :is-loading="!loaded" size="15"/>
  </div>
</template>

<script>
import {Spells} from "@/app/spells";
import {App}             from "@/constants/app";
import EqItemCardPreview from "@/components/preview/EQItemCardPreview.vue";
import EqSpellPreview    from "@/components/preview/EQSpellCardPreview.vue";
import EqWindow          from "@/components/eq-ui/EQWindow.vue";
import {Items} from "@/app/items";

export default {
  name: "EqSpellEffects",
  components: {
    "v-runtime-template": () => import("v-runtime-template"),
    EqItemCardPreview,
    EqSpellPreview,
    EqWindow,
  },
  async created() {
    // async each effect index if it exists
    // this is so loading spell effects and any subsequent ajax requests
    // do not block the card from loading
    for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
      if (this.spell["effectid_" + effectIndex] !== 254) {
        Spells.getSpellEffectInfo(this.spell, effectIndex).then((result) => {
          this.spellEffectInfo[result.index] = result.info;
          this.$forceUpdate();
        })
      }
    }

    // reagents
    let reagents = []
    for (let i = 0; i < 4; i++) {
      if (this.spell["components_" + i] > 0) {
        let reagent  = {}
        reagent.id   = this.spell["components_" + i]
        reagent.item = await Items.getItem(this.spell["components_" + i])
        reagents.push(reagent)
      }
    }

    this.sideLoadedSpellData = Spells.data;
    this.itemData = Items.items;
    this.loaded = true;

  },
  data() {
    return {
      spellEffectInfo: [],
      componentId: "",
      sideLoadedSpellData: {},
      loaded: false,
      itemData: {},
    }
  },
  props: {
    spell: Object
  },
}
</script>

<style scoped>

</style>
