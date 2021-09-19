<template>
  <div v-if="spell['descnum'] > 0 && effectDescription !== ''">
    {{ effectDescription }}
  </div>
</template>

<script>
import {Spells} from "@/app/spells/spells";
import EqItemPreview from "@/components/eq-ui/EQItemPreview.vue";
import EqSpellPreview from "@/components/eq-ui/EQSpellPreview.vue";
import EqWindow from "@/components/eq-ui/EQWindow.vue";

export default {
  name: "EqSpellDescription",
  components: {
    "v-runtime-template": () => import("v-runtime-template"),
    EqItemPreview,
    EqSpellPreview,
    EqWindow
  },
  created() {
    Spells.getSpellDescription(this.spell).then((result) => {
      if (result && result.trim() !== "") {
        this.effectDescription = result;
      }
    })
  },
  data() {
    return {
      effectDescription: ""
    }
  },
  props: {
    spell: Object
  }
}
</script>

<style scoped>

</style>
