<template>
  <eq-window class="mt-4">
  <pre
    v-if="debug"
    style="padding-top: 10px !important; width: 100%">{{ data }}</pre>
  </eq-window>
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
