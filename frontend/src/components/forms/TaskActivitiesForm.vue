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
              "disabled": true,
              "featured": false,
              "inputType": "text",
              "label": "Taskid",
              "model": "taskid",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"tasks.md\" target=\"_blank\"\u003eTask Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Activityid",
              "model": "activityid",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eActivity Identifier: Starts at 0\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Step",
              "model": "step",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eStep: 0 = Always Available, \u0026gt;0 = Must Complete Previous\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Activitytype",
              "model": "activitytype",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/task-system-guide/task-activity-types\" target=\"_blank\"\u003eActivity Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Name",
              "model": "target_name",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eTarget Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Item List",
              "model": "item_list",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skill List",
              "model": "skill_list",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eSkill Identifier List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell List",
              "model": "spell_list",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Description Override",
              "model": "description_override",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eDescription Override\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Goalid",
              "model": "goalid",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eGoal Identifier or \u003ca href=\"goallists.md\" target=\"_blank\"\u003eGoal List Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Goalmethod",
              "model": "goalmethod",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eGoal Method: 0 = Single Value, 1 = List\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Goalcount",
              "model": "goalcount",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eGoal Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Delivertonpc",
              "model": "delivertonpc",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eDeliver To NPC: 0 = No Delivery NPC, \u0026gt;0 = \u003ca href=\"npc_types.md\" target=\"_blank\"\u003eNPC Type Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zones",
              "model": "zones",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZones List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Optional",
              "model": "optional",
              "styleClasses": "col-2",
              "type": "input",
              "help": "\u003cp\u003eOptional: 0 = False, 1 = True\u003c/p\u003e\n"
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
