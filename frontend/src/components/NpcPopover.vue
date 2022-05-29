<template>
  <div v-if="npc">
    <div :id="npc.id + '-' + popoverId + '-popover'" style="display:inline-block; position: relative">
      <div style="width: 60px" class="d-inline-block ml-3 text-center">
      <span
        style="top: 50%"
        :class="'race-models-ctn-' + npc.race + '-' + npc.gender + '-' + npc.texture + '-' + npc.helmtexture + '-sm'"
      ></span>
      </div>
      <span class="ml-3 d-inline-block" style="top: 30%; position: absolute; min-width: 200px">
        {{ getCleanName(npc.name) }}
      </span>

    </div>

    <b-popover
      :target="npc.id + '-' + popoverId + '-popover'"
      custom-class="no-bg"
      placement="right"
      delay="0"
      boundary="viewport"
      :no-fade="true"
      triggers="hover focus"
    >
      <eq-window style="width: 650px; height: 100%">
        <eq-npc-card-preview
          :npc="npc"
        />
      </eq-window>
    </b-popover>

  </div>
</template>

<script>
import EqWindow         from "./eq-ui/EQWindow";
import EqNpcCardPreview from "./preview/EQNpcCardPreview";
import {Npcs}           from "../app/npcs";

export default {
  name: "npcPopover",
  props: {
    npc: Object,
    size: { // options: regular,sm
      type: String,
      required: false,
      default: "sm"
    },
  },
  watch: {
    npc: {
      deep: true,
      handler() {
      }
    },
  },
  methods: {
    getCleanName(name) {
      return Npcs.getCleanName(name)
    },
  },
  data() {
    return {
      popoverId: Math.random().toString(16).slice(2),
      npcEffectInfo: [],
      sideLoadednpcData: {},
    }
  },
  components: { EqNpcCardPreview, EqWindow }
}
</script>

<style scoped>

</style>
