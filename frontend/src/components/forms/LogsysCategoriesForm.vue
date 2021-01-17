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
              "label": "Log Category Id",
              "model": "log_category_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Log Category Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Log Category Description",
              "model": "log_category_description",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLog Category Description\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Log To Console",
              "model": "log_to_console",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLog to Console: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Log To File",
              "model": "log_to_file",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLog to File: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Log To Gmsay",
              "model": "log_to_gmsay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLog to GMSay: 0 = False, 1 = True\u003c/p\u003e\n"
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
