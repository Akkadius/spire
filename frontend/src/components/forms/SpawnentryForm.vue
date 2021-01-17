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
              "label": "SpawngroupID",
              "model": "spawngroupID",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spawngroup.md\" target=\"_blank\"\u003eUnique Spawn Group Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "NpcID",
              "model": "npcID",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_types.md\" target=\"_blank\"\u003eNPC Type Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Chance",
              "model": "chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eChance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Condition Value Filter",
              "model": "condition_value_filter",
              "styleClasses": "col-12",
              "type": "input"
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
