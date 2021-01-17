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
              "help": "\u003cp\u003eUnique Starting Items Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Race",
              "model": "race",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/race-list\" target=\"_blank\"\u003eRace\u003c/a\u003e: 0 = All\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Class",
              "model": "class",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClass\u003c/a\u003e: 0 = All\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deityid",
              "model": "deityid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity\u003c/a\u003e: 0 = Alll\u003c/p\u003e\n"
            },
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
              "label": "Itemid",
              "model": "itemid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Item Charges",
              "model": "item_charges",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Charges\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gm",
              "model": "gm",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGM: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slot",
              "model": "slot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/inventory/inventory-slots\" target=\"_blank\"\u003eSlot\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Expansion",
              "model": "min_expansion",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Expansion",
              "model": "max_expansion",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags",
              "model": "content_flags",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags Disabled",
              "model": "content_flags_disabled",
              "styleClasses": "col-12",
              "type": "input"
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
