<template>
  <img
    :alt="getExpansionName(expansion)"
    v-if="!getExpansionIcon(expansion).includes('base64')"
    :title="title + getExpansionName(expansion) + ' (' + expansion + ')'"
    :src="getExpansionIcon(expansion)"
    :style="'width: 56px; opacity: .7; border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px;'"
    class="mr-2 p-1 hover-highlight-inner"
  >
</template>

<script>
import Expansions from "../../app/utility/expansions";

export default {
  name: "ExpansionIcon",
  props: {
    expansion: {
      type: Number,
      required: true
    },
    title: {
      type: String,
      required: false,
      default: "",
    },
  },
  methods: {
    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion)
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion)
    },
  }
}
</script>

<style scoped>

</style>
