<template>
  <div v-if="item">
    <div :id="item.id + '-' + popoverId + '-popover'" style="display:inline-block; ">
      <span
        :class="'item-' + item.icon + (this.size === 'regular' ? '' : '-sm')" :title="item.icon"
        style="display: inline-block"
      />
      <span
        :class="(size === 'sm' ? 'ml-2' : 'ml-3')"
        :style="'position:relative;' + (this.size === 'regular' ? 'top: -15px' : '')"
      >{{ item.name }} {{annotation}} <slot></slot></span>

    </div>

    <b-popover
      :target="item.id + '-' + popoverId + '-popover'"
      custom-class="no-bg"
      placement="right"
      delay="0"
      boundary="viewport"
      :no-fade="true"
      triggers="hover focus"
      style="width: 500px !important"
    >
      <eq-window style="width: auto; height: 100%">
        <eq-item-card-preview
          :item-data="item"
          :show-related-data="showRelatedData"
        />
      </eq-window>
    </b-popover>

  </div>
</template>

<script>
import EqWindow          from "./eq-ui/EQWindow";
import {Items}           from "@/app/items";
import EqItemCardPreview from "@/components/preview/EQItemCardPreview";

export default {
  name: "ItemPopover",
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
  components: { EqItemCardPreview, EqWindow }
}
</script>

<style scoped>

</style>
