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
              "label": "Type",
              "model": "type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_types.md\" target=\"_blank\"\u003eNPC Type Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petpower",
              "model": "petpower",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet Power\u003c/p\u003e\n"
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
              "label": "Temp",
              "model": "temp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTemporary: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petcontrol",
              "model": "petcontrol",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet Control: 0 = No Control, 1 = No Attack Control, 2 = Full Control\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Petnaming",
              "model": "petnaming",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePet Naming: 0 = Soandsos Pet, 1 = Soandsos Familiar, 2 = Soandsos Warder, 3 = Random Naming (i.e. Gobaner), 4 = Keeps name from npc_types table\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Monsterflag",
              "model": "monsterflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMonster Flag: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Equipmentset",
              "model": "equipmentset",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"pet_equipmentset.md\" target=\"_blank\"\u003ePet Equipment Set Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
