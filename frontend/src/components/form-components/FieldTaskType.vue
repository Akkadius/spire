<template>
  <select
    class="form-control"
    v-model="value"
    :disabled="disabled"
    :maxlength="schema.max"
    :placeholder="schema.placeholder"
    :readonly="schema.readonly">
    <option
      v-for="(option, index) in options"
      :key="index"
      :value="index">
      {{ option }}
    </option>

  </select>
</template>

<script>
import {abstractField} from "vue-form-generator";
import {TASK_TYPES} from "@/app/constants/eq-task-constants";

export default {
  mixins: [abstractField],
  data() {
    return {
      options: TASK_TYPES
    }
  }
}
</script>
