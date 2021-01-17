<template>
  <vue-form-generator
    v-if="model"
    :schema="schema"
    :model="model"
    :options="formOptions"/>
</template>

<script>
export default {
  props: {
    model: {
      type: Object
    }
  },
  data() {
    return {
      schema: {
        fields:
          [
            {
              "featured": false,
              "inputType": "text",
              "label": "Ruleset Id",
              "model": "ruleset_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"rule_sets.md\" target=\"_blank\"\u003eRule Set Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rule Name",
              "model": "rule_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRule Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rule Value",
              "model": "rule_value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRule Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Notes",
              "model": "notes",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNotes\u003c/p\u003e\n"
            }
          ]
      },

      formOptions: {
        // validateAfterLoad: true,
        // validateAfterChanged: true
      }
    };
  }

}
</script>
