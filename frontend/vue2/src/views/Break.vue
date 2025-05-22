<template>
  <content-area>
    <div class="break-parent">
      <div class="break-child">
        <h1 class="eq-header " style="font-size: 300px">
          Break... Be right back...
        </h1>

        <app-loader :is-loading="true" class="mt-3 mb-8"/>
      </div>
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
.break-parent {
  position: relative;
  height: 95vh; /* Full viewport height for demonstration */
}

.break-child {
  position: absolute;   /* Enables absolute positioning */
  top: 50%;             /* Move down 50% of the parent height */
  left: 50%;            /* Move right 50% of the parent width */
  transform: translate(-50%, -50%); /* Offset back by 50% of the element's size */
  width: 2200px;
  height: 350px;
  text-align: center;    /* Optional: Center text horizontally */
  line-height: 100px;    /* Optional: Center single-line text vertically */
}
</style>
