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
              "label": "Accid",
              "model": "accid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account.md\" target=\"_blank\"\u003eAccount Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ip",
              "model": "ip",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIP Address\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Count",
              "model": "count",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNumber of times logged in from this IP\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lastused",
              "model": "lastused",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTimestamp of when account was last logged in\u003c/p\u003e\n"
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
