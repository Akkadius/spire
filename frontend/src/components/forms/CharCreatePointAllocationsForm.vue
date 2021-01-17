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
              "help": "\u003cp\u003eAllocation Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Str",
              "model": "base_str",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Strength\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Sta",
              "model": "base_sta",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Stamina\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Dex",
              "model": "base_dex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Dexterity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Agi",
              "model": "base_agi",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Agility\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Int",
              "model": "base_int",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Intelligence\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Wis",
              "model": "base_wis",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Wisdom\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Base Cha",
              "model": "base_cha",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Charisma\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Str",
              "model": "alloc_str",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Strength\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Sta",
              "model": "alloc_sta",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Stamina\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Dex",
              "model": "alloc_dex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Dexterity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Agi",
              "model": "alloc_agi",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Agility\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Int",
              "model": "alloc_int",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Intelligence\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Wis",
              "model": "alloc_wis",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Wisdom\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alloc Cha",
              "model": "alloc_cha",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllocated Charisma\u003c/p\u003e\n"
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
