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
              "help": "\u003cp\u003eUnique Tradeskill Recipe Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recipe Id",
              "model": "recipe_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"tradeskill_recipe.md\" target=\"_blank\"\u003eUnique Tradeskill Recipe Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Successcount",
              "model": "successcount",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSuccess Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Failcount",
              "model": "failcount",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFail Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Componentcount",
              "model": "componentcount",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComponent Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Salvagecount",
              "model": "salvagecount",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSalvage Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Iscontainer",
              "model": "iscontainer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Container: 0 = False, 1 = True\u003c/p\u003e\n"
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
