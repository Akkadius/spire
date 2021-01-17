<template>
  <div>
    <input type="checkbox" checked :id="getFieldID(schema)"
           class="eq-checkbox"
           v-model="value"
           v-bind:true-value="1"
           v-bind:false-value="0"
    >
    <label :for="getFieldID(schema)"></label>
  </div>
</template>

<script>
import {abstractField} from "vue-form-generator";

export default {
  mixins: [abstractField]
}
</script>
