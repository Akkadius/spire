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

export default {
  name: "EqDebug",
  components: { EqWindow },
  props: {
    data: Object
  },
  data() {
    return {
      debug: App.DEBUG
    }
  }
}
</script>

<style scoped>

</style>
