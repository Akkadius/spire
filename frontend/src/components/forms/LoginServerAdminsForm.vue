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
              "help": "\u003cp\u003eUnique Account Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Account Name",
              "model": "account_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAccount Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Account Password",
              "model": "account_password",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAccount Password\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "First Name",
              "model": "first_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFirst Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Name",
              "model": "last_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Email",
              "model": "email",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEmail\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Registration Date",
              "model": "registration_date",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRegistration Date\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Registration Ip Address",
              "model": "registration_ip_address",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account_ip.md\" target=\"_blank\"\u003eRegistration IP Address\u003c/a\u003e\u003c/p\u003e\n"
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
