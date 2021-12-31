<template>
  <div>
    <chrome
      :value="color"
      @input="updateValue"
    />
  </div>
</template>

<script>
import {Chrome} from 'vue-color'

export default {
  name: "ItemColorSelector",
  components: {
    Chrome
  },
  methods: {
    updateValue(value) {
      let update = {
        color: parseInt("FF" + value.hex.replace("#", ""), 16),
        hexColor: value.hex,
      }

      this.$emit("input", update);

      // console.log(value)
      // console.log("update")
      // console.log(update)
    },
  },
  props: {
    color: {
      required: true,
      type: String,
    },
  }
}
</script>

<style scoped>

</style>
