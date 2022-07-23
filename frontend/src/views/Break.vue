<template>
  <content-area>
    <eq-window class="text-center mt-6">
      <h1 class="eq-header mt-8" style="font-size: 200px">
        Break... Be right back..
      </h1>

      <app-loader :is-loading="true" class="mt-3 mb-8"/>
    </eq-window>
  </content-area>
</template>

<script>
import ContentArea        from "../components/layout/ContentArea";
import EqWindow           from "../components/eq-ui/EQWindow";
import LoaderFakeProgress from "../components/LoaderFakeProgress";
export default {
  name: "Break",
  components: { LoaderFakeProgress, EqWindow, ContentArea }
}
</script>

<style scoped>

</style>
