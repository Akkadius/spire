<template>
  <div>
    <eq-window>
      <div v-for="option in options" :key="option">
        {{ option }}
      </div>
    </eq-window>
  </div>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow.vue";

export default {
  name: "Test",
  components: { EqWindow },
  data() {
    return {
      options: ['something', 'something else', 'something else again']
    };
  },
}
</script>

<style scoped>

</style>
