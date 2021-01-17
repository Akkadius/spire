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
              "label": "Raidid",
              "model": "raidid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Raid Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Loottype",
              "model": "loottype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLoot Type\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Locked",
              "model": "locked",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLocked: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Motd",
              "model": "motd",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMessage of the Day\u003c/p\u003e\n"
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
