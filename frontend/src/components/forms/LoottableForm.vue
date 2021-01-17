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
              "help": "\u003cp\u003eUnique Loottable Identifier\u003c/p\u003e\n"
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
              "label": "Mincash",
              "model": "mincash",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Cash in Copper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxcash",
              "model": "maxcash",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Cash in Copper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Avgcoin",
              "model": "avgcoin",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAverage Coin in Copper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Done",
              "model": "done",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDone: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Expansion",
              "model": "min_expansion",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Expansion",
              "model": "max_expansion",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags",
              "model": "content_flags",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags Disabled",
              "model": "content_flags_disabled",
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
