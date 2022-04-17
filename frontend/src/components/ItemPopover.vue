<template>
  <div v-if="item">
    <div :id="item.id + '-popover'" style="display:inline-block; ">
      <span
        :class="'fade-in item-' + item.icon + (this.size === 'regular' ? '' : '-sm')" :title="item.icon"
        style="display: inline-block"
      />
      <span
        class="ml-3"
        :style="'position:relative;' + (this.size === 'regular' ? 'top: -15px' : '')"
      >{{ item.name }} {{annotation}}</span>

    </div>

    <b-popover
      :target="item.id + '-popover'"
      placement="auto"
      custom-class="no-bg"
      delay="1"
      triggers="hover focus"
      style="width: 500px !important"
    >
      <eq-window style="margin-right: 10px; width: auto; height: 90%">
        <eq-item-card-preview
          :item-data="item"
          :show-related-data="true"
        />
      </eq-window>
    </b-popover>

  </div>
</template>

<script>
import EqWindow          from "./eq-ui/EQWindow";
import EqItemCardPreview from "./eq-ui/EQItemCardPreview";
import {Items}           from "@/app/items";

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
    }
  },
  watch: {
    item: {
      deep: true,
      handler() {
      }
    },
  },
  data() {
    return {
      itemEffectInfo: [],
      itemData: {},
      sideLoadedItemData: {},
      componentId: "",
      reagents: [],
    }
  },
  mounted() {
    this.sideLoadedItemData = Items.data
  },
  components: { EqItemCardPreview, EqWindow }
}
</script>

<style scoped>

</style>
