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
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Book Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Txtfile",
              "model": "txtfile",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe text in the book. \u003ccode\u003eRepresents line spaces,\u003c/code\u003e\u003ccode\u003eis two line spaces,\u003c/code\u003e\u0026rdquo; is three line spaces, etc. (13 lines per book page)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Language",
              "model": "language",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/languages\" target=\"_blank\"\u003eLanguage\u003c/a\u003e\u003c/p\u003e\n"
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
