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
              "help": "\u003cp\u003eUnique Respawn Time Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Start",
              "model": "start",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStart UNIX Timestamp\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Duration",
              "model": "duration",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDuration in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Instance Id",
              "model": "instance_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"instance_list.md\" target=\"_blank\"\u003eInstance Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
