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
              "label": "Rank Id",
              "model": "rank_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"aa_ranks.md\" target=\"_blank\"\u003eRank Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slot",
              "model": "slot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Slot\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Id",
              "model": "effect_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eSpell Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base1",
              "model": "base1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFirst Base Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base2",
              "model": "base2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSecond Base Value\u003c/p\u003e\n"
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
