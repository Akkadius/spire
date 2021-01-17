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
              "label": "Zoneid",
              "model": "zoneid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Parentid",
              "model": "parentid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"object.md\" target=\"_blank\"\u003eObject Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bagidx",
              "model": "bagidx",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBag Index\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Itemid",
              "model": "itemid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charges",
              "model": "charges",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharges\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Droptime",
              "model": "droptime",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDrop Time\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot1",
              "model": "augslot1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot2",
              "model": "augslot2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot3",
              "model": "augslot3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot4",
              "model": "augslot4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot5",
              "model": "augslot5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 5\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot6",
              "model": "augslot6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 6\u003c/p\u003e\n"
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
