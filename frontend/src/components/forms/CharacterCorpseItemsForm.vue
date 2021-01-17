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
              "label": "Corpse Id",
              "model": "corpse_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_corpses.md\" target=\"_blank\"\u003eCorpse Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Equip Slot",
              "model": "equip_slot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEquipment Slot\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Item Id",
              "model": "item_id",
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
              "help": "\u003cp\u003eItem Charges\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aug 1",
              "model": "aug_1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Augment Slot 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aug 2",
              "model": "aug_2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Augment Slot 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aug 3",
              "model": "aug_3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Augment Slot 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aug 4",
              "model": "aug_4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Augment Slot 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aug 5",
              "model": "aug_5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Augment Slot 5\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aug 6",
              "model": "aug_6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Augment Slot 6\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attuned",
              "model": "attuned",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Attuned: 0 = False, 1 = True\u003c/p\u003e\n"
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
