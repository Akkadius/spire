<template>
  <div>
    <div v-if="spellEffectInfo.length > 0 && loaded">
      <div v-for="effect in spellEffectInfo">
        <v-runtime-template :template="'<span>' + effect + '</span>'" class="pb-6 mt-3 doc"/>
      </div>
    </div>
    <app-loader :is-loading="!loaded" size="15"/>
  </div>
</template>

<script>
import {Spells} from "@/app/spells/spells";
import {App} from "@/constants/app";
import EqItemPreview                         from "@/components/eq-ui/EQItemPreview.vue";
import EqSpellPreview                        from "@/components/eq-ui/EQSpellPreview.vue";
import EqWindow                              from "@/components/eq-ui/EQWindow.vue";

export default {
  name: "EqSpellEffects",
  components: {
    "v-runtime-template": () => import("v-runtime-template"),
    EqItemPreview,
    EqSpellPreview,
    EqWindow,
  },
  created() {
    // async each effect index if it exists
    // this is so loading spell effects and any subsequent ajax requests
    // do not block the card from loading
    for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
      if (this.spell["effectid_" + effectIndex] !== 254) {
        Spells.getSpellEffectInfo(this.spell, effectIndex).then((result) => {
          let effectInfo = this.spellEffectInfo;
          effectInfo.push(result)
          effectInfo           = effectInfo.sort((a, b) => a - b)
          this.spellEffectInfo = effectInfo;
          this.loaded = true;
        })
      }
    }

    this.sideLoadedSpellData = Spells.data;
  },
  data() {
    return {
      spellEffectInfo: [],
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
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
