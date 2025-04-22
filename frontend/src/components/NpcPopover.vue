<template>
  <div v-if="npc" class="d-inline-block">
    <popper
      :placement="placement"
      :append-to-body="true"
      trigger="hover"
      :options="{
        modifiers: {
          customOffset: {
            enabled: true,
            order: 830, // after built-in offset (800)
            fn: function(data) {
              if (placement === 'left') {
                data.offsets.popper.left -= offset;
              }
              if (placement === 'right') {
                data.offsets.popper.left += offset;
              }
              return data;
            }
          },
          computeStyle: {
            gpuAcceleration: false
          },
        }
      }"
    >
      <eq-window
        v-if="popoverEnabled"
        style="min-width: 600px; width: auto; height: auto; z-index: 99999" class="p-3">
        <eq-npc-card-preview
          :limit-entries="limitEntries"
          :no-stats="noStats"
          :npc="npc"
        />
      </eq-window>

      <div slot="reference" style="display:inline-block; position: relative">
        <div style="width: 60px; height: 50px" class="d-inline-block ml-3 text-center" v-if="showImage">
        <span
          style="top: 50%; filter: drop-shadow(10px 5px 5px #000);"
          :class="'race-models-ctn-' + getRaceImage(npc) + '-sm'"
        ></span>
        </div>
        <span
          v-if="showLabel"
          :class="(showImage ? 'ml-3' : '') + ' d-inline-block'" :style="(showImage ? 'top: 30%; position: absolute; min-width: 300px' : '')"
        >
        {{ getCleanName(npc.name) }} {{ formatLastName(npc.lastname) }} {{additionalLabel ? additionalLabel : ""}}
      </span>
        <slot></slot>
      </div>

    </popper>
  </div>
</template>

<script>
import EqWindow from "./eq-ui/EQWindow";
import EqNpcCardPreview from "./preview/EQNpcCardPreview";
import { Npcs } from "@/app/npcs";
import Popper from 'vue-popperjs';

export default {
  name: "NpcPopover",
  components: {
    EqNpcCardPreview,
    EqWindow,
    Popper
  },
  props: {
    npc: Object,
    size: {
      type: String,
      required: false,
      default: "sm"
    },
    popoverEnabled: {
      type: Boolean,
      required: false,
      default: true
    },
    additionalLabel: {
      type: String,
      required: false
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
    showLastName: {
      type: Boolean,
      required: false,
      default: true
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
    offset: {
      type: Number,
      required: false,
      default: 200
    },
    placement: {
      type: String,
      required: false,
      default: "right"
    }
  },
  data() {
    return {
      popoverId: Math.random().toString(16).slice(2),
      npcEffectInfo: [],
      sideLoadednpcData: {},
    }
  },
  methods: {
    formatLastName(name) {
      if (!this.showLastName) return '';
      return (name && name.length > 0 ? ` (${name})` : "");
    },
    getRaceImage(npc) {
      return Npcs.getRaceImage(npc);
    },
    getCleanName(name) {
      return Npcs.getCleanName(name);
    },
    getImageWrapperMargin() {
      return 'ml-3';
    },
    getLabelMargin() {
      return this.showImage ? 'ml-3 d-inline-block' : 'd-inline-block';
    },
    getLabelTopOffset() {
      return this.showImage ? 'top: 30%' : '';
    }
  }
}
</script>
