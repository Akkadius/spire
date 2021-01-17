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
              "help": "\u003cp\u003eUnique Spawn Group Identifier\u003c/p\u003e\n"
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
              "label": "Spawn Limit",
              "model": "spawn_limit",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpawn Limit\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dist",
              "model": "dist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max X",
              "model": "max_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min X",
              "model": "min_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMin X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Y",
              "model": "max_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Y",
              "model": "min_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMin Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Delay",
              "model": "delay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRoaming Delay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mindelay",
              "model": "mindelay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Delay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Despawn",
              "model": "despawn",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/spawns/npc-despawn-types\" target=\"_blank\"\u003eDespawn Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Despawn Timer",
              "model": "despawn_timer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDespawn Timer in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wp Spawns",
              "model": "wp_spawns",
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
