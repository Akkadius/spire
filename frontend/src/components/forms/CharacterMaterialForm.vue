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
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Blue",
              "model": "blue",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBlue: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Green",
              "model": "green",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGreen: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Red",
              "model": "red",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRed: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Use Tint",
              "model": "use_tint",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTint: 0 = None, 255 = MAx\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Color",
              "model": "color",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eColor\u003c/p\u003e\n"
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
