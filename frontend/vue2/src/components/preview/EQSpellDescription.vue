<template>
  <div>
    <div v-if="spell['descnum'] > 0 && effectDescription !== '' && loaded">
      {{ effectDescription }}
    </div>
    <app-loader :is-loading="!loaded" size="15"/>
  </div>
</template>

<script>
import {Spells}          from "@/app/spells";
import EqItemCardPreview from "@/components/preview/EQItemCardPreview.vue";
import EqSpellPreview    from "@/components/preview/EQSpellCardPreview.vue";
import EqWindow          from "@/components/eq-ui/EQWindow.vue";

export default {
  name: "EqSpellDescription",
  components: {
    "v-runtime-template": () => import("v-runtime-template"),
    EqItemCardPreview,
    EqSpellPreview,
    EqWindow
  },
  created() {
    Spells.getSpellDescription(this.spell).then((result) => {
      if (result && result.trim() !== "") {
        this.effectDescription = result;
      }
      this.loaded = true;
    })
  },
  data() {
    return {
      effectDescription: "",
      loaded: false,
    }
  },
  props: {
    spell: Object
  }
}
</script>

<style scoped>

</style>
