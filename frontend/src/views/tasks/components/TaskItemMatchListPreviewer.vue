<template>
  <div>
    <eq-window-simple
      style="height: 95vh; overflow-y: scroll; overflow-x: hidden" class="p-0"
    >
      <div class="font-weight-bold text-center">
        Goal Match List Preview ({{ TASK_ACTIVITY_TYPES[activity.activitytype] }})
      </div>

      <div v-if="items && items.length > 0" class="text-center mt-1">
        Found ({{ items.length }}) matching item(s)
      </div>

      <table
        class="eq-table eq-highlight-rows bordered"
        style="display: table; overflow-x: scroll"
      >
        <thead>
        <tr>
          <th style="text-align: center; width: 50px" class="text-center">Id</th>
          <th style="text-align: center;">Name</th>
        </tr>
        </thead>
        <tbody>
        <tr
          v-for="(item, index) in items"
          :key="item.id"
          :id="'item-selection-row-' + item.id"
        >
          <td>
            {{ item.id }}
          </td>
          <td class="text-left" style="vertical-align: middle">
            <item-popover
              :item="item"
              v-if="Object.keys(item).length > 0 && item"
              size="regular"
            />
          </td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple                            from "@/components/eq-ui/EQWindowSimple";
import EqCheckbox                                from "@/components/eq-ui/EQCheckbox";
import {TASK_ACTIVITY_TYPE, TASK_ACTIVITY_TYPES} from "@/app/constants/eq-task-constants";
import ItemPopover                               from "@/components/ItemPopover";
import {Items}                                   from "@/app/items";

export default {
  name: "TaskItemMatchListPreviewer",
  components: { ItemPopover, EqCheckbox, EqWindowSimple },
  data() {
    return {
      // result sets
      items: {},

      // constants
      TASK_ACTIVITY_TYPE: TASK_ACTIVITY_TYPE,
      TASK_ACTIVITY_TYPES: TASK_ACTIVITY_TYPES,
    }
  },
  props: {
    activity: {
      type: Object,
      required: true,
    },
  },

  watch: {
    activity: {
      deep: true,
      handler() {
        this.load()
      }
    },
  },

  methods: {
    async load() {
      this.loadItemMatches()
    },

    async loadItemMatches() {
      let items = []
      for (let m of this.activity.item_id_list.split("|")) {
        m = m.toLowerCase()
        if (m.length === 0 && this.activity.item_id_list.length !== 0) {
          continue;
        }

        items.push(parseInt(m))
      }

      console.log(items)

      this.items = await Items.loadItemsBulk(items)

      console.log(this.items)
    },

  },
  mounted() {
    this.load()
  }
}
</script>

<style scoped>
#npctable td {
  vertical-align: middle !important;
}
</style>
