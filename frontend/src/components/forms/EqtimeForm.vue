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
              "label": "Minute",
              "model": "minute",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinute\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hour",
              "model": "hour",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHour\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Day",
              "model": "day",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Month",
              "model": "month",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMonth\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Year",
              "model": "year",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eYear\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Realtime",
              "model": "realtime",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReal Time\u003c/p\u003e\n"
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
