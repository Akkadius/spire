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
              "label": "Dib",
              "model": "dib",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petid",
              "model": "petid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Petition Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charname",
              "model": "charname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Accountname",
              "model": "accountname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account.md\" target=\"_blank\"\u003eAccount Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lastgm",
              "model": "lastgm",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast GM\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petitiontext",
              "model": "petitiontext",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePetition Text\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gmtext",
              "model": "gmtext",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGM Text\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone",
              "model": "zone",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Short Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Urgency",
              "model": "urgency",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUrgency\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charclass",
              "model": "charclass",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eCharacter Class\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charrace",
              "model": "charrace",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/race-list\" target=\"_blank\"\u003eCharacter Race\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charlevel",
              "model": "charlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharacter Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Checkouts",
              "model": "checkouts",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCheckouts\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unavailables",
              "model": "unavailables",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnavailables\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ischeckedout",
              "model": "ischeckedout",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Checked Out: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Senttime",
              "model": "senttime",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSent Time UNIX Timestamp\u003c/p\u003e\n"
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
