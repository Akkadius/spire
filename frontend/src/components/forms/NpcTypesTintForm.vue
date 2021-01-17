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
              "help": "\u003cp\u003eUnique NPC Type Tint Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tint Set Name",
              "model": "tint_set_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTint Set Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red1h",
              "model": "red1h",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Helmet: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn1h",
              "model": "grn1h",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Helmet: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu1h",
              "model": "blu1h",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Helmet: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red2c",
              "model": "red2c",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Chest: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn2c",
              "model": "grn2c",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Chest: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu2c",
              "model": "blu2c",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Chest: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red3a",
              "model": "red3a",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Arms: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn3a",
              "model": "grn3a",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Arms: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu3a",
              "model": "blu3a",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Arms: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red4b",
              "model": "red4b",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Bracers: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn4b",
              "model": "grn4b",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Bracers: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu4b",
              "model": "blu4b",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Bracers: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red5g",
              "model": "red5g",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Hands: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn5g",
              "model": "grn5g",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Hands: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu5g",
              "model": "blu5g",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Hands: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red6l",
              "model": "red6l",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Legs: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn6l",
              "model": "grn6l",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Legs: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu6l",
              "model": "blu6l",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Legs: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red7f",
              "model": "red7f",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Feet: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn7f",
              "model": "grn7f",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Feet: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu7f",
              "model": "blu7f",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Feet: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red8x",
              "model": "red8x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Unknown: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn8x",
              "model": "grn8x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Unknown: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu8x",
              "model": "blu8x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Unknown: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red9x",
              "model": "red9x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed for Unknown: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Grn9x",
              "model": "grn9x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen for Unknown: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Blu9x",
              "model": "blu9x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue for Unknown: 0 = None, 255 = Max\u003c/p\u003e\n"
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
