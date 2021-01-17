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
              "help": "\u003cp\u003eUnique Spell Identifier\u003c/p\u003e\n"
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
              "label": "Player 1",
              "model": "player_1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePlayer_1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Teleport Zone",
              "model": "teleport_zone",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe zone you are teleporting to or the \u003ca href=\"npc_types.md\" target=\"_blank\"\u003eNPC Name\u003c/a\u003e you want to spawn.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "You Cast",
              "model": "you_cast",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe message sent to others when you cast the spell.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Other Casts",
              "model": "other_casts",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe message seen when someone around you casts the spell.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cast On You",
              "model": "cast_on_you",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe message received when the spell is cast on you.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cast On Other",
              "model": "cast_on_other",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe message recieved when the spell is cast on another.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Fades",
              "model": "spell_fades",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThe message recieved when the spell fades.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Range",
              "model": "range",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRange\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aoerange",
              "model": "aoerange",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArea of Effect Range\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pushback",
              "model": "pushback",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePush Back\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pushup",
              "model": "pushup",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePush Up\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cast Time",
              "model": "cast_time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCast Time in Milliseconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recovery Time",
              "model": "recovery_time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecovery Time in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recast Time",
              "model": "recast_time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecast Time in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Buffdurationformula",
              "model": "buffdurationformula",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/buff-duration-formulas\" target=\"_blank\"\u003eBuff Duration Formula\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Buffduration",
              "model": "buffduration",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBuff Duration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "AEDuration",
              "model": "AEDuration",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArea of Effect Duration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mana",
              "model": "mana",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMana Cost\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value1",
              "model": "effect_base_value1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value2",
              "model": "effect_base_value2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value3",
              "model": "effect_base_value3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value4",
              "model": "effect_base_value4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value5",
              "model": "effect_base_value5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 5\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value6",
              "model": "effect_base_value6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 6\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value7",
              "model": "effect_base_value7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 7\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value8",
              "model": "effect_base_value8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 8\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value9",
              "model": "effect_base_value9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 9\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value10",
              "model": "effect_base_value10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 10\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value11",
              "model": "effect_base_value11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 11\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Base Value12",
              "model": "effect_base_value12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Base Value 12\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value1",
              "model": "effect_limit_value1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value2",
              "model": "effect_limit_value2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value3",
              "model": "effect_limit_value3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value4",
              "model": "effect_limit_value4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value5",
              "model": "effect_limit_value5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 5\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value6",
              "model": "effect_limit_value6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 6\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value7",
              "model": "effect_limit_value7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 7\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value8",
              "model": "effect_limit_value8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 8\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value9",
              "model": "effect_limit_value9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 9\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value10",
              "model": "effect_limit_value10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 10\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value11",
              "model": "effect_limit_value11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 11\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect Limit Value12",
              "model": "effect_limit_value12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Limit Value 12\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max1",
              "model": "max1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max2",
              "model": "max2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max3",
              "model": "max3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max4",
              "model": "max4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max5",
              "model": "max5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 5\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max6",
              "model": "max6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 6\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max7",
              "model": "max7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 7\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max8",
              "model": "max8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 8\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max9",
              "model": "max9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 9\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max10",
              "model": "max10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 10\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max11",
              "model": "max11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 11\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max12",
              "model": "max12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax 12\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Icon",
              "model": "icon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIcon\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Memicon",
              "model": "memicon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMemmed Icon\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Components1",
              "model": "components1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Components2",
              "model": "components2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Components3",
              "model": "components3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Components4",
              "model": "components4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Component Counts1",
              "model": "component_counts1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComponent Count 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Component Counts2",
              "model": "component_counts2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComponent Count 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Component Counts3",
              "model": "component_counts3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComponent Count 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Component Counts4",
              "model": "component_counts4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComponent Count 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "NoexpendReagent1",
              "model": "NoexpendReagent1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIf it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "NoexpendReagent2",
              "model": "NoexpendReagent2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIf it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "NoexpendReagent3",
              "model": "NoexpendReagent3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIf it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "NoexpendReagent4",
              "model": "NoexpendReagent4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIf it is a number between 1-4 it means component number 1-4 is a focus and not to expend it. If it is a valid item ID it means this item is a focus as well.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula1",
              "model": "formula1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 1\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula2",
              "model": "formula2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 2\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula3",
              "model": "formula3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 3\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula4",
              "model": "formula4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 4\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula5",
              "model": "formula5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 5\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula6",
              "model": "formula6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 6\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula7",
              "model": "formula7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 7\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula8",
              "model": "formula8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 8\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula9",
              "model": "formula9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 9\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula10",
              "model": "formula10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 10\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula11",
              "model": "formula11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 11\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Formula12",
              "model": "formula12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/base-value-formulas\" target=\"_blank\"\u003eFormula 12\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "LightType",
              "model": "LightType",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLight Type\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "GoodEffect",
              "model": "goodEffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGood Effect: 0 = Detrimental, 1 = Beneficial, 2 = Beneficial Group Only\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Activated",
              "model": "Activated",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eActivated\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Resisttype",
              "model": "resisttype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/resist-types\" target=\"_blank\"\u003eResist Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid1",
              "model": "effectid1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 1\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid2",
              "model": "effectid2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 2\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid3",
              "model": "effectid3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 3\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid4",
              "model": "effectid4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 4\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid5",
              "model": "effectid5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 5\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid6",
              "model": "effectid6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 6\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid7",
              "model": "effectid7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 7\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid8",
              "model": "effectid8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 8\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid9",
              "model": "effectid9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 9\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid10",
              "model": "effectid10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 10\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid11",
              "model": "effectid11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 11\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectid12",
              "model": "effectid12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eEffect Identifier 12\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Targettype",
              "model": "targettype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/target-types\" target=\"_blank\"\u003eTarget Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Basediff",
              "model": "basediff",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBase Difficult Fizzle Adjustment\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skill",
              "model": "skill",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eSkill Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zonetype",
              "model": "zonetype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-types\" target=\"_blank\"\u003eZone Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "EnvironmentType",
              "model": "EnvironmentType",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/environment-types\" target=\"_blank\"\u003eEnvironment Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "TimeOfDay",
              "model": "TimeOfDay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/time-of-day-types\" target=\"_blank\"\u003eTime of Day Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes1",
              "model": "classes1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Warrior\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes2",
              "model": "classes2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Cleric\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes3",
              "model": "classes3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Paladin\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes4",
              "model": "classes4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Shadow Knight\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes5",
              "model": "classes5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Ranger\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes6",
              "model": "classes6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Druid\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes7",
              "model": "classes7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Monk\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes8",
              "model": "classes8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Bard\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes9",
              "model": "classes9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Rogue\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes10",
              "model": "classes10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Shaman\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes11",
              "model": "classes11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Necromancer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes12",
              "model": "classes12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Wizard\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes13",
              "model": "classes13",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Magician\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes14",
              "model": "classes14",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Enchanter\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes15",
              "model": "classes15",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Beastlord\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes16",
              "model": "classes16",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level for Berserker\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "CastingAnim",
              "model": "CastingAnim",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCasting Animation\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "TargetAnim",
              "model": "TargetAnim",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Animation\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "TravelType",
              "model": "TravelType",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTravel Type (Unused)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "SpellAffectIndex",
              "model": "SpellAffectIndex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Affect Index\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Disallow Sit",
              "model": "disallow_sit",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisallow Sit: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities0",
              "model": "deities0",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities1",
              "model": "deities1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities2",
              "model": "deities2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities3",
              "model": "deities3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities4",
              "model": "deities4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities5",
              "model": "deities5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities6",
              "model": "deities6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities7",
              "model": "deities7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities8",
              "model": "deities8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities9",
              "model": "deities9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities10",
              "model": "deities10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities11",
              "model": "deities11",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities12",
              "model": "deities12",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities13",
              "model": "deities13",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities14",
              "model": "deities14",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities15",
              "model": "deities15",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deities16",
              "model": "deities16",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity List\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field142",
              "model": "field142",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field143",
              "model": "field143",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "New Icon",
              "model": "new_icon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNew Icon\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spellanim",
              "model": "spellanim",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Animation\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Uninterruptable",
              "model": "uninterruptable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUninterruptable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "ResistDiff",
              "model": "ResistDiff",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eResist Difference\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dot Stacking Exempt",
              "model": "dot_stacking_exempt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDamage Over Time Stacking Exempt: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deleteable",
              "model": "deleteable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDeleteable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "RecourseLink",
              "model": "RecourseLink",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eRecourse Spell Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "No Partial Resist",
              "model": "no_partial_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Partial Resist: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field152",
              "model": "field152",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field153",
              "model": "field153",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Short Buff Box",
              "model": "short_buff_box",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eShort Buff Box: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Descnum",
              "model": "descnum",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"db_str.md\" target=\"_blank\"\u003eDescription Number\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Typedescnum",
              "model": "typedescnum",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"db_str.md\" target=\"_blank\"\u003eType Description Number\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectdescnum",
              "model": "effectdescnum",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"db_str.md\" target=\"_blank\"\u003eEffect Description Number\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectdescnum2",
              "model": "effectdescnum2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"db_str.md\" target=\"_blank\"\u003eEffect Description Number 2\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc No Los",
              "model": "npc_no_los",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC No Line of Sight: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field160",
              "model": "field160",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Reflectable",
              "model": "reflectable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReflectable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bonushate",
              "model": "bonushate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBonus Hate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field163",
              "model": "field163",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field164",
              "model": "field164",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Trap",
              "model": "ldon_trap",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"traps.md\" target=\"_blank\"\u003eLDoN Trap Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "EndurCost",
              "model": "EndurCost",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEndurance Cost\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "EndurTimerIndex",
              "model": "EndurTimerIndex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEndurance Timer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "IsDiscipline",
              "model": "IsDiscipline",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Discipline: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field169",
              "model": "field169",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field170",
              "model": "field170",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field171",
              "model": "field171",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field172",
              "model": "field172",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "HateAdded",
              "model": "HateAdded",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHate Added\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "EndurUpkeep",
              "model": "EndurUpkeep",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEndurance Upkeep\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Numhitstype",
              "model": "numhitstype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/numhit-types\" target=\"_blank\"\u003eNumber of Hits Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Numhits",
              "model": "numhits",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNumber of Hits\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvpresistbase",
              "model": "pvpresistbase",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Resist Base\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvpresistcalc",
              "model": "pvpresistcalc",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Resist Calc\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvpresistcap",
              "model": "pvpresistcap",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Resist Cap\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Category",
              "model": "spell_category",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-groups\" target=\"_blank\"\u003eSpell Category\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field181",
              "model": "field181",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field182",
              "model": "field182",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pcnpc Only Flag",
              "model": "pcnpc_only_flag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePC/NPC Only Flag: 0 = Not Applicable, 1 = PCs and Mercs, 2 = NPCs\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cast Not Standing",
              "model": "cast_not_standing",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCast Not Standing: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Can Mgb",
              "model": "can_mgb",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Mass Group Buff: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nodispell",
              "model": "nodispell",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Dispell: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Category",
              "model": "npc_category",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/npc-spell-categories\" target=\"_blank\"\u003eNPC Spell Category Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Usefulness",
              "model": "npc_usefulness",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC Usefulness\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "MinResist",
              "model": "MinResist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "MaxResist",
              "model": "MaxResist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Viral Targets",
              "model": "viral_targets",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eViral Targets\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Viral Timer",
              "model": "viral_timer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eViral Timer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nimbuseffect",
              "model": "nimbuseffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNimbus Effect\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "ConeStartAngle",
              "model": "ConeStartAngle",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCone Start Angle\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "ConeStopAngle",
              "model": "ConeStopAngle",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCone Stop Angle\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Sneaking",
              "model": "sneaking",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSneaking: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Not Extendable",
              "model": "not_extendable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNot Extendable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field198",
              "model": "field198",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field199",
              "model": "field199",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Suspendable",
              "model": "suspendable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSuspendable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Viral Range",
              "model": "viral_range",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eViral Range\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Songcap",
              "model": "songcap",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSong Cap\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field203",
              "model": "field203",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field204",
              "model": "field204",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "No Block",
              "model": "no_block",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Block: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field206",
              "model": "field206",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spellgroup",
              "model": "spellgroup",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-groups\" target=\"_blank\"\u003eSpell Group\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rank",
              "model": "rank",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRank\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field209",
              "model": "field209",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field210",
              "model": "field210",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "CastRestriction",
              "model": "CastRestriction",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-target-restrictions\" target=\"_blank\"\u003eCast Restrictions\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Allowrest",
              "model": "allowrest",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAllow Rest: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "InCombat",
              "model": "InCombat",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIn Combat: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "OutofCombat",
              "model": "OutofCombat",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eOut Of Combat: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field215",
              "model": "field215",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field216",
              "model": "field216",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field217",
              "model": "field217",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aemaxtargets",
              "model": "aemaxtargets",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArea of Effect Max Targets\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxtargets",
              "model": "maxtargets",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax Targets\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field220",
              "model": "field220",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field221",
              "model": "field221",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field222",
              "model": "field222",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field223",
              "model": "field223",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Persistdeath",
              "model": "persistdeath",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePersist Death: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field225",
              "model": "field225",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field226",
              "model": "field226",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Dist",
              "model": "min_dist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Dist Mod",
              "model": "min_dist_mod",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Distance Modifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Dist",
              "model": "max_dist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Dist Mod",
              "model": "max_dist_mod",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Distance Modifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Range",
              "model": "min_range",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Range\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field232",
              "model": "field232",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field233",
              "model": "field233",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field234",
              "model": "field234",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field235",
              "model": "field235",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Field236",
              "model": "field236",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
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
