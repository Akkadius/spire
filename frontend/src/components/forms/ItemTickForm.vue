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
              "label": "It Itemid",
              "model": "it_itemid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "It Chance",
              "model": "it_chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eChance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "It Level",
              "model": "it_level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "It Id",
              "model": "it_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "It Qglobal",
              "model": "it_qglobal",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"quest_globals.md\" target=\"_blank\"\u003eQuest Global Identifier\u003c/a\u003e (Deprecated)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "It Bagslot",
              "model": "it_bagslot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/inventory/inventory-slots\" target=\"_blank\"\u003eBag Slot\u003c/a\u003e\u003c/p\u003e\n"
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
