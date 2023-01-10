<template>
  <div
    v-if="id"
    :style="getStyle()"
    @mouseover="hovered = true"
    @mouseout="hovered = false"
  >
    <span v-if="label">{{label}}</span>
    <input
      type="checkbox"
      :id="id"
      class='eq-checkbox'
      v-model="inputVal"
      v-bind:true-value="trueValue"
      v-bind:false-value="falseValue"
      :checked="(inputVal > 0)"
      :disabled="disabled === 1"
      @change="change()"
    >
    <label :for="id" class="eq-checkbox-label"></label>
  </div>
</template>

<script>
export default {
  name: "EqCheckbox",
  data() {
    return {
      id: "",
      hovered: false,
    }
  },
  computed: {
    inputVal: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit('input', val);
      }
    }
  },
  props: {
    isChecked: {
      type: Boolean,
      required: false,
      default: false
    },
    disabled: {
      default: 0
    },
    fadeWhenNotTrue: {
      type: Boolean,
      required: false,
      default: false,
    },
    label: {
      type: String,
      required: false,
      default: "",
    },
    value: {
      type: [Number, String, Boolean],
      required: false,
      default: 0,
    },
    trueValue: {
      type: [Number, String, Boolean],
      required: false,
      default: 1,
    },
    falseValue: {
      type: [Number, String, Boolean],
      required: false,
      default: 0,
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    change() {
      this.$emit('change', this.value);
    },
    getStyle() {
      if (this.fadeWhenNotTrue) {
        if (this.hovered) {
          return 'opacity: 1; font-weight: bold;'
        }

        if (this.value !== this.trueValue) {
          return 'opacity: .3'
        }
        return 'opacity: 1'
      }
      return ''
    },
    init() {
      this.id = this.randId()
    },
    randId() {
      return "_" + Math.random().toString(36).substr(2, 9);
    }
  },

}
</script>

<style scoped>

</style>
