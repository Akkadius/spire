<template>
  <div v-if="loot" class="d-inline-block">
    <popper
      :placement="placement"
      :append-to-body="true"
      trigger="hover"
      offset="0,200"
      :options="{
        modifiers: {
          computeStyle: {
            gpuAcceleration: false
          },
          preventOverflow: {
            boundariesElement: 'viewport'
          },
          customOffset: {
            enabled: true,
            order: 830,
            fn: function(data) {
              data.offsets.popper.top += 200;
              data.offsets.popper.left -= 300;
              return data;
            }
          }
        }
      }"
    >
      <eq-window style="min-width: 600px; width: auto; height: auto; z-index: 99999" class="p-3">
        <eq-loot-card-preview :loot="loot" />
      </eq-window>

      <div slot="reference" style="display: inline-block; position: relative;">
        <span v-if="showLabel" class="d-inline-block">
          {{ loot.name }}
        </span>
        <slot></slot>
      </div>
    </popper>
  </div>
</template>

<script>
import EqWindow from "./eq-ui/EQWindow";
import EqLootCardPreview from "./preview/EQLootCardPreview";
import Popper from 'vue-popperjs';

export default {
  name: "LootPopover",
  components: {
    EqLootCardPreview,
    EqWindow,
    Popper
  },
  props: {
    loot: Object,
    size: {
      type: String,
      required: false,
      default: "sm"
    },
    showImage: {
      type: Boolean,
      required: false,
      default: true
    },
    showLabel: {
      type: Boolean,
      required: false,
      default: true
    },
    placement: {
      type: String,
      required: false,
      default: "right"
    }
  },
  data() {
    return {
      popoverId: Math.random().toString(16).slice(2)
    }
  }
}
</script>
