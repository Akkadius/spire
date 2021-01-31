<template>
  <div v-if="id">
    <input type="checkbox" :id="id" class='eq-checkbox' :checked="isChecked" :disabled="disabled === 1">
    <label :for="id"></label>
  </div>
</template>

<script>
export default {
  name: "EqCheckbox",
  data() {
    return {
      id: ""
    }
  },
  props: {
    isChecked: {
      type: Boolean,
      required: false,
      default: 0
    },
    disabled: {
      default: 0
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.id = this.randId()
    },
    randId() {
      return "_" + Math.random().toString(36).substr(2, 9);
    }
  }

}
</script>

<style scoped>

</style>
