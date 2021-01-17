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
              "label": "Platinum",
              "model": "platinum",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePlatinum\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gold",
              "model": "gold",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGold\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Silver",
              "model": "silver",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSilver\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Copper",
              "model": "copper",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCopper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Platinum Bank",
              "model": "platinum_bank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBanked Platinum\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gold Bank",
              "model": "gold_bank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBanked Gold\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Silver Bank",
              "model": "silver_bank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBanked Silver\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Copper Bank",
              "model": "copper_bank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBanked Copper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Platinum Cursor",
              "model": "platinum_cursor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCursor Platinum\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gold Cursor",
              "model": "gold_cursor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCursor Gold\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Silver Cursor",
              "model": "silver_cursor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCursor Silver\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Copper Cursor",
              "model": "copper_cursor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCursor Copper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Radiant Crystals",
              "model": "radiant_crystals",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRadiant Crystals\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Career Radiant Crystals",
              "model": "career_radiant_crystals",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCareer Radiant Crystals\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ebon Crystals",
              "model": "ebon_crystals",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEbon Crystals\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Career Ebon Crystals",
              "model": "career_ebon_crystals",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCareer Ebon Crystals\u003c/p\u003e\n"
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
