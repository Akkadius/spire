<template>
  <div>

    <span
      style="position: absolute; right: 7%; opacity: .8"
      :class="'race-models-ctn-' + npc.race + '-' + npc.gender + '-' + npc.texture + '-' + npc.helmtexture"
    />

    <div class="row">
      <div class="col-8 pl-5">
        <h6 class="eq-header" style="margin: 0px; margin-bottom: 10px">
          {{ getCleanName() }}
        </h6>
      </div>
    </div>

    <eq-debug :data="npc"/>
<!--    {{ npc }}-->
  </div>
</template>

<script>
import EqDebug from "./EQDebug";
import {Npcs}  from "../../app/npcs";
export default {
  name: "EqNpcCardPreview",
  components: { EqDebug },
  props: {
    npc: {
      type: Object,
      default: {},
      required: true
    },
  },
  methods: {
    getCleanName() {
      return Npcs.getCleanName(this.npc.name)
    }
  }
}
</script>

<style scoped>

</style>
