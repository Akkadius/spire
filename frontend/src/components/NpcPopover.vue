<template>
  <div v-if="npc">
    <div :id="npc.id + '-' + popoverId + '-popover'" style="display:inline-block; position: relative">
      <div style="width: 60px; height: 50px" class="d-inline-block ml-3 text-center" v-if="showImage">
        <span
          style="top: 50%; filter: drop-shadow(10px 5px 5px #000);"
          :class="'race-models-ctn-' + getRaceImage(npc) + '-sm'"
        ></span>
      </div>
      <span
        v-if="showLabel"
        class="ml-3 d-inline-block" style="top: 30%; position: absolute; min-width: 300px"
      >
        {{ getCleanName(npc.name) }} {{ (npc.lastname && npc.lastname.length > 0 ? ` (${npc.lastname})` : "") }} {{additionalLabel ? additionalLabel : ""}}
      </span>
      <slot></slot>
    </div>

    <b-popover
      v-if="popoverEnabled"
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
          :limit-entries="limitEntries"
          :no-stats="noStats"
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
    additionalLabel: {
      type: String,
      required: false,
    },
    showImage: {
      type: Boolean,
      required: false,
      default: true
    },
    noStats: {
      type: Boolean,
      required: false,
      default: false
    },
    limitEntries: {
      type: Number,
      required: false,
      default: 20
    },
    showLabel: {
      type: Boolean,
      required: false,
      default: true
    },
    popoverEnabled: {
      type: Boolean,
      required: false,
      default: true
    }
  },
  watch: {
    npc: {
      deep: true,
      handler() {
      }
    },
  },
  mounted() {
    console.log("[popover] no limit", this.noLimitEntries)
  },
  methods: {
    getRaceImage(npc) {
      return Npcs.getRaceImage(npc)
    },
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
