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
              "help": "\u003cp\u003eUnique NPC Spell Set Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC Spell Set Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Parent List",
              "model": "parent_list",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eInherit all the spells from this list, and merge them with our spells. Only one level of inheritance is allowed, so your parent\u0026rsquo;s parent will not be included.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack Proc",
              "model": "attack_proc",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe combat proc that an NPC with this spell set will add to their list of procs. \u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Proc Chance",
              "model": "proc_chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProc Chance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Range Proc",
              "model": "range_proc",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe ranged proc that an NPC with this spell set will add to their list of procs.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rproc Chance",
              "model": "rproc_chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRanged Proc Chance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Defensive Proc",
              "model": "defensive_proc",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe defensive proc that an NPC with this spell set will add to their list of procs.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dproc Chance",
              "model": "dproc_chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDefensive Proc Chance: 0 = Never, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fail Recast",
              "model": "fail_recast",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFail Recast\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Engaged No Sp Recast Min",
              "model": "engaged_no_sp_recast_min",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Engaged No Sp Recast Max",
              "model": "engaged_no_sp_recast_max",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Engaged B Self Chance",
              "model": "engaged_b_self_chance",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Engaged B Other Chance",
              "model": "engaged_b_other_chance",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Engaged D Chance",
              "model": "engaged_d_chance",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pursue No Sp Recast Min",
              "model": "pursue_no_sp_recast_min",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pursue No Sp Recast Max",
              "model": "pursue_no_sp_recast_max",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pursue D Chance",
              "model": "pursue_d_chance",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Idle No Sp Recast Min",
              "model": "idle_no_sp_recast_min",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Idle No Sp Recast Max",
              "model": "idle_no_sp_recast_max",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Idle B Chance",
              "model": "idle_b_chance",
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
