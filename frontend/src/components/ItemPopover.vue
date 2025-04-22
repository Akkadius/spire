<template>
  <div v-if="item">
    <popper
      placement="right"
      :append-to-body="true"
      trigger="hover"
      :options="{
        placement: 'right',
        modifiers: {
          computeStyle: {
            gpuAcceleration: false  // ← disables translate3d()
          },
          offset: { offset: '0,200px' }
        }
      }"
    >
      <eq-window style="width: auto; height: auto; z-index: 99999">
        <eq-item-card-preview
          :item-data="item"
          :show-related-data="showRelatedData"
        />
      </eq-window>

      <div slot="reference" style="display:inline-block; ">
      <span
        :class="'item-' + item.icon + (this.size === 'regular' ? '' : '-sm')" :title="item.icon"
        style="display: inline-block"
      />
        <span
          :class="(size === 'sm' ? 'ml-2' : 'ml-3')"
          :style="'position:relative;' + (this.size === 'regular' ? 'top: -15px' : '')"
        >{{ item.name }} {{ annotation }} <slot></slot></span>

      </div>
    </popper>

  </div>
</template>

<script>
import EqWindow from "./eq-ui/EQWindow";
import EqItemCardPreview from "@/components/preview/EQItemCardPreview";
import Popper from 'vue-popperjs';

export default {
  name: "ItemPopover",
  components: {
    EqItemCardPreview,
    EqWindow,
    Popper
  },
  props: {
    item: Object,
    size: { // options: regular,sm
      type: String,
      required: false,
      default: "sm"
    },
    annotation: {
      type: String,
      required: false,
      default: ""
    },
    showRelatedData: {
      type: Boolean,
      required: false,
      default: false
    }
  },
  data() {
    return {
      popoverId: Math.random().toString(16).slice(2),
      itemEffectInfo: [],
    }
  },
}
</script>

<style scoped>

</style>
