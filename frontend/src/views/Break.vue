<template>
  <content-area>
    <div class="text-center mt-6" style="left: 13%; top: 20%; position: fixed">
      <h1 class="eq-header" style="font-size: 200px">
        Break... Be right back..
      </h1>

      <app-loader :is-loading="true" class="mt-3 mb-8"/>
    </div>
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
