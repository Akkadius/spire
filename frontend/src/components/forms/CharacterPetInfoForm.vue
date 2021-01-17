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
              "label": "Char Id",
              "model": "char_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pet",
              "model": "pet",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petname",
              "model": "petname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petpower",
              "model": "petpower",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet Power\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Id",
              "model": "spell_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hp",
              "model": "hp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHealth\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mana",
              "model": "mana",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMana\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Size",
              "model": "size",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSize\u003c/p\u003e\n"
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
