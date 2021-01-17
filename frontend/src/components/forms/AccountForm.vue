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
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eName\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charname",
              "model": "charname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharacter name last logged in on this account.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Sharedplat",
              "model": "sharedplat",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePlatinum in Shared Bank.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Password",
              "model": "password",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePrivate loginserver password.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Status",
              "model": "status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003eStatus\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ls Id",
              "model": "ls_id",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lsaccount Id",
              "model": "lsaccount_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLoginserver Account Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gmspeed",
              "model": "gmspeed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGM Speed: 0 = Disabled, 1 = Enabled\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Revoked",
              "model": "revoked",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eOOC Revoked: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Karma",
              "model": "karma",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eKarma\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minilogin Ip",
              "model": "minilogin_ip",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinilogin IP Address\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hideme",
              "model": "hideme",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHide Me: 0 = Disabled, 1 = Enabled\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rulesflag",
              "model": "rulesflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRules Flag\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Suspendeduntil",
              "model": "suspendeduntil",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Suspension of the Account ends\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time Creation",
              "model": "time_creation",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Creation UNIX Timestamp\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Expansion",
              "model": "expansion",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExpansion\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ban Reason",
              "model": "ban_reason",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBan Reason\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Suspend Reason",
              "model": "suspend_reason",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSuspension Reason\u003c/p\u003e\n"
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
