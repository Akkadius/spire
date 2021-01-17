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
              "label": "Gid",
              "model": "gid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"group_id.md\" target=\"_blank\"\u003eGroup Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Leadername",
              "model": "leadername",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLeader Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Marknpc",
              "model": "marknpc",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMark NPC: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Leadershipaa",
              "model": "leadershipaa",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLeadership AA\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maintank",
              "model": "maintank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMain Tank: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assist",
              "model": "assist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssist: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Puller",
              "model": "puller",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePuller: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mentoree",
              "model": "mentoree",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMentoree: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mentor Percent",
              "model": "mentor_percent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMain Tank: 0 = None, 100 = Max\u003c/p\u003e\n"
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
