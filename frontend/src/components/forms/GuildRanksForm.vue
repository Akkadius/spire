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
              "help": "\u003cp\u003eRank Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Title",
              "model": "title",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTitle\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Hear",
              "model": "can_hear",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Hear: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Speak",
              "model": "can_speak",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Speak: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Invite",
              "model": "can_invite",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Invite: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Remove",
              "model": "can_remove",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Remove: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Promote",
              "model": "can_promote",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Promote: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Demote",
              "model": "can_demote",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Demote: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Motd",
              "model": "can_motd",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan MOTD: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Warpeace",
              "model": "can_warpeace",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan War Peace: 0 = False, 1 = True\u003c/p\u003e\n"
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
