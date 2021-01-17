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
              "label": "Set Id",
              "model": "set_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Pet Equipment Set Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Setname",
              "model": "setname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet Equipment Set Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nested Set",
              "model": "nested_set",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNested Set Identifier\u003c/p\u003e\n"
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
