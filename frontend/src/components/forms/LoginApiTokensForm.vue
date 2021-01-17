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
              "help": "\u003cp\u003eUnique Login API Token Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Token",
              "model": "token",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eToken Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Write",
              "model": "can_write",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Write: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Read",
              "model": "can_read",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Read: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Created At",
              "model": "created_at",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCreated At Date\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Updated At",
              "model": "updated_at",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUpdated At Date\u003c/p\u003e\n"
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
