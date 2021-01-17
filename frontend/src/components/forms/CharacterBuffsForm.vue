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
              "label": "Character Id",
              "model": "character_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slot Id",
              "model": "slot_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBuff Slot\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Id",
              "model": "spell_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eBuff Spell Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Caster Name",
              "model": "caster_name",
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
              "label": "Melee Rune",
              "model": "melee_rune",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMelee Rune\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Magic Rune",
              "model": "magic_rune",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMagic Rune\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Persistent",
              "model": "persistent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePersistent: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dot Rune",
              "model": "dot_rune",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDamage Over Time Rune\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Caston X",
              "model": "caston_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eX Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Caston Y",
              "model": "caston_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eY Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Caston Z",
              "model": "caston_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZ Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "ExtraDIChance",
              "model": "ExtraDIChance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExtra DI Chance\u003c/p\u003e\n"
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
