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
              "label": "Merchantid",
              "model": "merchantid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMerchant Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slot",
              "model": "slot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSlot\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Item",
              "model": "item",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Faction Required",
              "model": "faction_required",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Required\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Level Required",
              "model": "level_required",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel Required\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alt Currency Cost",
              "model": "alt_currency_cost",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"alternate_currency.md\" target=\"_blank\"\u003eAlternate Currency Cost\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes Required",
              "model": "classes_required",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClasses Required\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Probability",
              "model": "probability",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProbability: 0 = Never, 100 = Always\u003c/p\u003e\n"
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
