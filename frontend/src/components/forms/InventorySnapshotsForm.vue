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
              "label": "Time Index",
              "model": "time_index",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eInventory Snapshot Time Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charid",
              "model": "charid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slotid",
              "model": "slotid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/inventory/inventory-slots\" target=\"_blank\"\u003eSlot Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Color",
              "model": "color",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eColor\u003c/p\u003e\n"
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
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Instnodrop",
              "model": "instnodrop",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Drop: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Custom Data",
              "model": "custom_data",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCustom Data\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ornamenticon",
              "model": "ornamenticon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eOrnamentation Icon\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ornamentidfile",
              "model": "ornamentidfile",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eOrnamentation Texture\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ornament Hero Model",
              "model": "ornament_hero_model",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eOrnamentation Hero\u0026rsquo;s Forge Model\u003c/p\u003e\n"
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
