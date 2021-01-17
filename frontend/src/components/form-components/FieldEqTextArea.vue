<template>
  <textarea
    v-model="value"
    :id="getFieldID(schema)"
    :class="schema.fieldClasses + ' form-control'"
    :disabled="disabled"
    :maxlength="schema.max"
    :minlength="schema.min"
    :placeholder="schema.placeholder"
    :readonly="schema.readonly"
    :required="schema.required"
    :rows="schema.rows || 2"
    :name="schema.inputName"
    v-attributes="'input'"/>
</template>

<script>
import {abstractField} from "vue-form-generator";

export default {
  mixins: [abstractField]
}
</script>
