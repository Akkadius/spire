<template>
  <div>
    <eq-window-simple
      style="height: 95vh; overflow-y: scroll; overflow-x: hidden" class="p-0"
    >
      <div class="font-weight-bold text-center">
        Item Match List Preview
      </div>

      <div v-if="items && items.length > 0" class="text-center mt-1">
        Found ({{ items.length }}) matching item(s)
      </div>

      <div v-if="items.length === 0" class="text-center mt-3 font-weight-bold">
        No item(s) were found
      </div>

      <table
        class="eq-table eq-highlight-rows bordered"
        style="display: table; overflow-x: scroll; overflow-y: hidden"
      >
        <thead>
        <tr>
          <th style="width: 60px"></th>
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
            <div class="btn-group" role="group">
              <b-button
                class="btn-dark btn-sm btn-outline-danger"
                @click="removeItem(item)"
                title="Remove item from list"
              >
                <i class="fa fa-trash-o"></i>
              </b-button>
            </div>
          </td>
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
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";
import EqCheckbox     from "@/components/eq-ui/EQCheckbox";
import ItemPopover    from "@/components/ItemPopover";
import {Items}        from "@/app/items";

export default {
  name: "TaskItemMatchListPreviewer",
  components: { ItemPopover, EqCheckbox, EqWindowSimple },
  data() {
    return {
      // result sets
      items: {},
    }
  },
  props: {
    idList: {
      type: String,
      required: true,
    },
  },

  watch: {
    idList: {
      deep: true,
      handler() {
        this.load()
      }
    },
  },

  methods: {
    removeItem(item) {
      this.$emit('remove-item', item.id);
    },

    async load() {
      this.loadItemMatches()
    },

    async loadItemMatches() {
      let items = []
      for (let m of this.idList.split("|")) {
        m = m.toLowerCase()
        if (m.length === 0 && this.idList.length !== 0) {
          continue;
        }

        items.push(parseInt(m))
      }

      this.items = await Items.loadItemsBulk(items)
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
