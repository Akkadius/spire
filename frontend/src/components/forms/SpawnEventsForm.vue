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
              "help": "\u003cp\u003eUnique Spawn Event Entry Identifier\u003c/p\u003e\n"
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
              "label": "Cond Id",
              "model": "cond_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spawn_conditions.md\" target=\"_blank\"\u003eSpawn Condition Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Period",
              "model": "period",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePeriod\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Next Minute",
              "model": "next_minute",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNext Minute\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Next Hour",
              "model": "next_hour",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNext Hour\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Next Day",
              "model": "next_day",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNext Day\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Next Month",
              "model": "next_month",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNext Month\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Next Year",
              "model": "next_year",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNext Year\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Enabled",
              "model": "enabled",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEnabled: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Action",
              "model": "action",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/spawns/action-types\" target=\"_blank\"\u003eAction Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Argument",
              "model": "argument",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArgument: (Based on Action) 0 = Argument Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Strict",
              "model": "strict",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStrict Date Criteria: 0 = False, 1 = True\u003c/p\u003e\n"
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
