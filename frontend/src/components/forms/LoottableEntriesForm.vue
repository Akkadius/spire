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
              "label": "Loottable Id",
              "model": "loottable_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"loottable.md\" target=\"_blank\"\u003eLoottable Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lootdrop Id",
              "model": "lootdrop_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"lootdrop.md\" target=\"_blank\"\u003eLootdrop Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Multiplier",
              "model": "multiplier",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMultiplier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Droplimit",
              "model": "droplimit",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Drops\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mindrop",
              "model": "mindrop",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Drops\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Probability",
              "model": "probability",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProbability: 0 = Never, 100 = Always\u003c/p\u003e\n"
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
