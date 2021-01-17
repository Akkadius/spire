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
              "label": "Lootdrop Id",
              "model": "lootdrop_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"lootdrop.md\" target=\"_blank\"\u003eLootdrop Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Item Charges",
              "model": "item_charges",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Charges\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Equip Item",
              "model": "equip_item",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEquip Item: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Chance",
              "model": "chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eChance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Disabled Chance",
              "model": "disabled_chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisabled Chance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Trivial Min Level",
              "model": "trivial_min_level",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Trivial Max Level",
              "model": "trivial_max_level",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Multiplier",
              "model": "multiplier",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMultiplier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Min Level",
              "model": "npc_min_level",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Max Level",
              "model": "npc_max_level",
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
