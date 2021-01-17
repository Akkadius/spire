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
              "label": "Char Id",
              "model": "char_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guild Id",
              "model": "guild_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"guilds.md\" target=\"_blank\"\u003eGuild Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rank",
              "model": "rank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/guild-ranks\" target=\"_blank\"\u003eRank\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tribute Enable",
              "model": "tribute_enable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTribute Enable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Total Tribute",
              "model": "total_tribute",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTotal Tribute\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Tribute",
              "model": "last_tribute",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Tribute\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Banker",
              "model": "banker",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBanked: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Public Note",
              "model": "public_note",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePublic Note\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alt",
              "model": "alt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAlt: 0 = False, 1 = True\u003c/p\u003e\n"
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
