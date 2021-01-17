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
              "help": "\u003cp\u003eUnique World Server Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Long Name",
              "model": "long_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLong Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Short Name",
              "model": "short_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eShort Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tag Description",
              "model": "tag_description",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTag Description\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Login Server List Type Id",
              "model": "login_server_list_type_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"login_server_list_types.md\" target=\"_blank\"\u003eList Type Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Last Ip Address",
              "model": "last_ip_address",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast IP Address\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Login Server Admin Id",
              "model": "login_server_admin_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"login_server_admins.md\" target=\"_blank\"\u003eLogin Server Admin Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Server Trusted",
              "model": "is_server_trusted",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Server Trusted: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Note",
              "model": "note",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNote\u003c/p\u003e\n"
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
