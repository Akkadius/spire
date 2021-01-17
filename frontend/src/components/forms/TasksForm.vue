<template>
  <vue-form-generator
    v-if="model"
    :schema="schema"
    :model="model"
    :options="formOptions"/>
</template>

<script>
import {TASK_REWARD_METHOD} from "@/app/constants/eq-task-constants";

export default {
  props: {
    model: {
      type: Object
    }
  },
  data() {
    return {
      schema: {
        groups: [
          {
            legend: "Completion",
            fields: [
              {
                "featured": false,
                "inputType": "text",
                "label": "Reward",
                "model": "reward",
                "styleClasses": "col-8",
                "type": "eq-text-input",
                "help": "Reward Description"
              },
              {
                "featured": false,
                "inputType": "text",
                "label": "Reward ID",
                "model": "rewardid",
                "styleClasses": "col-4",
                "type": "eq-text-input",
                "help": "\u003ca href=\"items.md\" target=\"_blank\"\u003eReward Item Identifier\u003c/a\u003e"
              },
              {
                "featured": false,
                "inputType": "text",
                "label": "Cashreward",
                "model": "cashreward",
                "styleClasses": "col-12",
                "type": "eq-cash",
              },
              {
                "featured": false,
                "inputType": "text",
                "label": "EXP Reward",
                "model": "xpreward",
                "styleClasses": "col-4",
                "type": "eq-text-input",
                "help": "Experience Reward"
              },
              {
                "featured": false,
                "label": "Reward Method",
                "model": "rewardmethod",
                "styleClasses": "col-4",
                "type": "select",
                "help": "Reward Method: 0 = Single Item ID, 1 = List of Items, 2 = Quest Controlled",
                selectOptions: { hideNoneSelectedText: true },
                values: function () {
                  let rows = [];
                  for (let key in TASK_REWARD_METHOD) {
                    rows.push({ id: key, name: `(${key}) ` + TASK_REWARD_METHOD[key] })
                  }
                  return rows
                }
              },
              {
                "featured": false,
                "inputType": "text",
                "label": "Faction Reward",
                "model": "faction_reward",
                "styleClasses": "col-4",
                "type": "eq-text-input",
                "help": "Faction Reward"
              },
              {
                "featured": false,
                "inputType": "text",
                "label": "Completion Emote",
                "model": "completion_emote",
                "styleClasses": "col-12",
                "type": "eq-text-input",
                "help": "Completion Emote"
              }
            ]
          }
        ],
        fields:
          [
            {
              "featured": false,
              "inputType": "text",
              "label": "Title",
              "model": "title",
              "styleClasses": "col-6",
              "type": "eq-text-input",
              "help": "Title"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Id",
              "model": "id",
              "styleClasses": "col-3",
              "type": "eq-text-input",
              "help": "Unique Task Identifier"
            },
            {
              "featured": false,
              "label": "Task Type",
              "model": "type",
              "styleClasses": "col-3",
              "type": "task-type",
              "help": "\u003ca href=\"https://eqemu.gitbook.io/server/categories/task-system-guide/task-types\" target=\"_blank\"\u003eTask Type\u003c/a\u003e"
            },


            {
              "featured": false,
              "inputType": "",
              "label": "Description",
              "model": "description",
              "rows": "5",
              "styleClasses": "col-12",
              "type": "eq-text-area",
              "help": "Description"
            },

            {
              "featured": false,
              "inputType": "text",
              "label": "Duration",
              "model": "duration",
              "styleClasses": "col-9",
              "type": "task-duration",
              "help": "Duration"
            },
            {
              "featured": false,
              "label": "Duration Code",
              "model": "duration_code",
              "styleClasses": "col-3",
              "type": "task-duration-code"
            },


            {
              "featured": false,
              "inputType": "text",
              "label": "Min Level",
              "model": "minlevel",
              "styleClasses": "col-3",
              "type": "eq-text-input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Level",
              "model": "maxlevel",
              "styleClasses": "col-3",
              "type": "eq-text-input"
            },
            {
              "featured": false,
              "label": "Repeatable",
              "model": "repeatable",
              "styleClasses": "col-4",
              "type": "eq-yes-no",
              "help": "Repeatable: 0 = False, 1 = True"
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
