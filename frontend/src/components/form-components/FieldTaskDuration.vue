<template>
  <div class="input-group">
    <div class="input-group-prepend">
      <input
        type="text"
        class="form-control"
        v-model="value"
        :disabled="disabled"
        :maxlength="schema.max"
        :placeholder="schema.placeholder"
        :readonly="schema.readonly">

    </div>

    <select
      v-model="value"
      class="custom-select"
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

  </div>
</template>

<script>
import {abstractField} from "vue-form-generator";
import {TASK_DURATION_HUMAN} from "@/app/constants/eq-task-constants";

export default {
  mixins: [abstractField],
  data() {
    return {
      options: TASK_DURATION_HUMAN
    }
  }
}
</script>
