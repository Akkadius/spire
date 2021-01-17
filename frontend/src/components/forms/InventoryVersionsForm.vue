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
              "label": "Version",
              "model": "version",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eInventory Version Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Step",
              "model": "step",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStep\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bot Step",
              "model": "bot_step",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBot Step\u003c/p\u003e\n"
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
