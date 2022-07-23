<template>
  <div>
    <eq-window title="Loot">
      <div class="row">
        <div class="col-lg-4">
          <input
            type="text"
            class="form-control ml-2"
            placeholder="Search for item names"
          >
        </div>
      </div>
    </eq-window>

    <eq-window>
      <div style="overflow-y: scroll; height: 80vh">
        <table
          id="loot-table"
          class="eq-table eq-highlight-rows"
          style="font-size: 14px; "
          v-if="loot && loot.length > 0"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th
              v-for="(header, index) in Object.keys(loot[0])"
              class="text-center"
              v-if="!doesColumnHaveObjects(loot, header)"
              :id="'column-' + header"
            >{{ header }}
            </th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="(row, index) in loot"
            :id="'npc-' + row.id"
            :key="index"
          >
            <td
              :style="' text-align: center;'"
              v-for="(key, colIndex) in Object.keys(row)"
              v-if="doesRowColumnHaveObjects(row, key)"
            >
              <span v-if="key !== 'name'">{{ row[key] }}</span>
              <loot-popover
                v-if="key === 'name'"
                :loot="row"
              />
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </eq-window>
  </div>
</template>

<script>
import EqWindow    from "../eq-ui/EQWindow";
import {Loot}      from "../../app/loot";
import LootPopover from "../LootPopover";

export default {
  name: "LootSubEditor",
  components: { LootPopover, EqWindow },
  data() {
    return {
      loot: {}
    }
  },
  async mounted() {
    this.loot = await Loot.getLoot()
  },
  methods: {
    doesColumnHaveObjects(data, column) {
      if (typeof column === 'object') {
        return true
      }

      return data.find((row) => {
        return typeof row[column] === 'object' && row[column] !== null && Object.keys(row[column])
      })
    },
    doesRowColumnHaveObjects(r, key) {
      return (typeof r[key] !== 'undefined') && !(typeof r[key] === 'object' && r[key] !== null && Object.keys(r[key]))
    },
  }
}
</script>

<style scoped>

</style>
