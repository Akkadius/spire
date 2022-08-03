<template>
  <div class="d-inline-block ml-1">
    <div class="d-inline-block" v-if="flagsEnabled && flagsEnabled.length > 0">
      <b-badge
        title="Enabled when this content flag is enabled"
        pill
        variant="warning"
        v-for="f in flagsEnabled.split(',')"
      >{{ f }}
      </b-badge>
    </div>
    <div class="d-inline-block" v-if="flagsDisabled && flagsDisabled.length > 0">
      <b-badge
        title="Disabled when this content flag is enabled"
        pill
        variant="danger"
        v-for="f in flagsDisabled.split(',')"
      >{{ f }}
      </b-badge>
    </div>
  </div>
</template>

<script>
export default {
  name: "ContentFlagPills",
  props: {
    flagsEnabled: {
      type: String,
      required: false
    },
    flagsDisabled: {
      type: String,
      required: false
    },
  },
}
</script>

<style scoped>

</style>
