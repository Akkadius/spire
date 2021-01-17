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
              "label": "Slot",
              "model": "slot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSlot\u003c/p\u003e\n"
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
              "label": "Caster Level",
              "model": "caster_level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCaster Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Castername",
              "model": "castername",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCaster Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ticsremaining",
              "model": "ticsremaining",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTics Remaining\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Counters",
              "model": "counters",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCounters\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Numhits",
              "model": "numhits",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNumber of Hits\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rune",
              "model": "rune",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRune\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Instrument Mod",
              "model": "instrument_mod",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eInstrument Modifier\u003c/p\u003e\n"
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
