<template>
  <div v-if="debug" class="mt-4">
    <h3 class="eq-header">Debug</h3>
    <pre
      style="padding-top: 10px !important; width: 100%">{{ data }}</pre>
  </div>
</template>

<script>

import {App} from "@/constants/app";
import EqWindow from "@/components/eq-ui/EQWindow";
import {EventBus} from "@/app/event-bus/event-bus";
import LocalSettings from "@/app/local-settings/localsettings";

export default {
  name: "EqDebug",
  components: { EqWindow },
  props: {
    data: Object
  },
  methods: {
    debugUpdatedListener(){
      this.debug = LocalSettings.get("debug-mode") === "true"
    }
  },
  created() {
    EventBus.$on("DEBUG_UPDATED", this.debugUpdatedListener);
  },
  destroyed() {
    EventBus.$off("DEBUG_UPDATED", this.debugUpdatedListener);
  },
  data() {
    return {
      debug: App.DEBUG
    }
  },
}
</script>

<style scoped>

</style>
