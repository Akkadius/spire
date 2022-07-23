<template>
  <div v-if="loot">
    <div :id="loot.id + '-' + popoverId + '-popover'" style="display:inline-block; position: relative">
      <span
        v-if="showLabel"
        class="d-inline-block"
      >
        {{ loot.name }}
      </span>
      <slot></slot>
    </div>

    <b-popover
      :target="loot.id + '-' + popoverId + '-popover'"
      custom-class="no-bg"
      placement="right"
      delay="0"
      boundary="viewport"
      :no-fade="true"
      triggers="hover focus"
    >
      <eq-loot-card-preview
        style="width: 650px; height: 100%"
        :loot="loot"
      />
    </b-popover>

  </div>
</template>

<script>
import EqWindow          from "./eq-ui/EQWindow";
import EqLootCardPreview from "./preview/EQLootCardPreview";

export default {
  name: "lootPopover",
  props: {
    loot: Object,
    size: { // options: regular,sm
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
    }
  },
  watch: {
    loot: {
      deep: true,
      handler() {
      }
    },
  },
  methods: {},
  data() {
    return {
      popoverId: Math.random().toString(16).slice(2),
    }
  },
  components: { EqLootCardPreview, EqWindow }
}
</script>

<style scoped>

</style>
