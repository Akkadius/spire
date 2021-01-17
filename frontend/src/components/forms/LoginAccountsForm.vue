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
              "help": "\u003cp\u003e\u003ca href=\"account.md\" target=\"_blank\"\u003eAccount Name\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Account Email",
              "model": "account_email",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAccount Email\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Source Loginserver",
              "model": "source_loginserver",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSource Loginserver\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Ip Address",
              "model": "last_ip_address",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account_ip.md\" target=\"_blank\"\u003eLast IP Address\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Login Date",
              "model": "last_login_date",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Login Date\u003c/p\u003e\n"
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
