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
              "label": "Id",
              "model": "id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Data Bucket Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Key",
              "model": "key",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eKey\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Value",
              "model": "value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eValue\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Expires",
              "model": "expires",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExpiration UNIX Timestamp\u003c/p\u003e\n"
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
