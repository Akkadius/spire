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
              "help": "\u003cp\u003eType\u003c/p\u003e\n"
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
              "label": "Comment",
              "model": "comment",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComment\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fromlevel",
              "model": "fromlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFrom Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tolevel",
              "model": "tolevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTo Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes",
              "model": "classes",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClasses\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aacount",
              "model": "aacount",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Timezone",
              "model": "timezone",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Zone\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Timeposted",
              "model": "timeposted",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Posted UNIX Timestamp\u003c/p\u003e\n"
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
