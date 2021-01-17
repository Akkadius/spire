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
              "label": "Player Id",
              "model": "player_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guk Wins",
              "model": "guk_wins",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDeepest Guk Wins\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mir Wins",
              "model": "mir_wins",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMiragul\u0026rsquo;s Menagerie Wins\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mmc Wins",
              "model": "mmc_wins",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMistmoore Catacombs Wins\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ruj Wins",
              "model": "ruj_wins",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRujarkian Hills Wins\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tak Wins",
              "model": "tak_wins",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTakish-Hiz Wins\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guk Losses",
              "model": "guk_losses",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDeepest Guk Losses\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mir Losses",
              "model": "mir_losses",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMiragul\u0026rsquo;s Menagerie Losses\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mmc Losses",
              "model": "mmc_losses",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMistmoore Catacombs Losses\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ruj Losses",
              "model": "ruj_losses",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRujarkian Hills Losses\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tak Losses",
              "model": "tak_losses",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTakish-Hiz Losses\u003c/p\u003e\n"
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
