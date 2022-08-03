<template>
  <div class="d-inline-block">
    <div class="d-inline-block ml-3" v-if="filterData.min_expansion > 0">
      <expansion-icon
        :expansion="filterData.min_expansion"
        title="Minimum Expansion Required "
      />
    </div>

    <div class="d-inline-block ml-3" v-if="filterData.max_expansion > 0">
      <expansion-icon
        :expansion="filterData.max_expansion"
        title="Max Expansion Required "
      />
    </div>

    <content-flag-pills :flags-enabled="filterData.content_flags"/>
    <content-flag-pills :flags-disabled="filterData.content_flags_disabled"/>
  </div>
</template>

<script>
import ContentFlagPills from "./ContentFlagPills";
import ExpansionIcon    from "./ExpansionIcon";

export default {
  name: "ContentFilterDisplayPills",
  components: { ExpansionIcon, ContentFlagPills },
  props: {
    filterData: {
      type: Object,
      required: true
    },
  },
}
</script>

<style scoped>

</style>
