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
              "label": "Char Id",
              "model": "char_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eUnique Character Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Type",
              "model": "type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eType\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Start",
              "model": "start",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStart UNIX Timestamp\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Duration",
              "model": "duration",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDuration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Enable",
              "model": "enable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEnabled: 0 = False, 1 = True\u003c/p\u003e\n"
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
