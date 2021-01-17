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
              "help": "\u003cp\u003eUnique Guild Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eName\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Leader",
              "model": "leader",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minstatus",
              "model": "minstatus",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003eMinimum Status\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Motd",
              "model": "motd",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMessage of the Day\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tribute",
              "model": "tribute",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTribute\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Motd Setter",
              "model": "motd_setter",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Channel",
              "model": "channel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eChannel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Url",
              "model": "url",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWebsite URL\u003c/p\u003e\n"
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
