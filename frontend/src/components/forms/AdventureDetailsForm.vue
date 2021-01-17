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
              "help": "\u003cp\u003eUnique Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Adventure Id",
              "model": "adventure_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Adventure Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Instance Id",
              "model": "instance_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"instance_list.md\" target=\"_blank\"\u003eInstance Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Count",
              "model": "count",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCount\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assassinate Count",
              "model": "assassinate_count",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssassinate Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Status",
              "model": "status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStatus\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time Created",
              "model": "time_created",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Created UNIX Timestamp\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time Zoned",
              "model": "time_zoned",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Zoned UNIX Timestamp\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time Completed",
              "model": "time_completed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Completed UNIX Timestamp\u003c/p\u003e\n"
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
