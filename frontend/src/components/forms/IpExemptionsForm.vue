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
              "label": "Exemption Id",
              "model": "exemption_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExemption Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Exemption Ip",
              "model": "exemption_ip",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account_ip.md\" target=\"_blank\"\u003eExemption IP Address\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Exemption Amount",
              "model": "exemption_amount",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExemption Amount\u003c/p\u003e\n"
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
