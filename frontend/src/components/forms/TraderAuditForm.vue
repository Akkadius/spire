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
              "label": "Time",
              "model": "time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Seller",
              "model": "seller",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSeller\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Buyer",
              "model": "buyer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBuyer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Itemname",
              "model": "itemname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Quantity",
              "model": "quantity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eQuantity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Totalcost",
              "model": "totalcost",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTotal Cost\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Trantype",
              "model": "trantype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTransaction Type\u003c/p\u003e\n"
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
