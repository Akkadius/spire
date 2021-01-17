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
              "label": "Level",
              "model": "level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Exp Mod",
              "model": "exp_mod",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExperience Modifier: 0.5 = 50%, 1 = 100%, 1.5 = 150%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aa Exp Mod",
              "model": "aa_exp_mod",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Experience Modifier: 0.5 = 50%, 1 = 100%, 1.5 = 150%\u003c/p\u003e\n"
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
