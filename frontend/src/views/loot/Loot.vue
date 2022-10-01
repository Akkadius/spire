<template>
  <content-area style="padding: 0 !important">
    <div class="row">
      <div class="col-7">
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

        <eq-window class="p-0">
          <div style="overflow-y: scroll; height: 84vh">
            <table
              id="loot-table"
              class="eq-table eq-highlight-rows bordered"
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
                @mouseover="showTable(row)"
              >
                <td
                  :style="' text-align: center;'"
                  v-for="(key, colIndex) in Object.keys(row)"
                  v-if="doesRowColumnHaveObjects(row, key)"
                >
                  <span >{{ row[key] }}</span>
<!--                  <loot-popover-->
<!--                    v-if="key === 'name'"-->
<!--                    :loot="row"-->
<!--                  />-->
                </td>
              </tr>
              </tbody>
            </table>
          </div>
        </eq-window>
      </div>
      <div class="col-5">
        <eq-loot-card-preview
          v-if="Object.keys(previewedTable).length > 0"
          :loot="previewedTable"
        />
      </div>
    </div>

  </content-area>
</template>

<script>
import EqWindow          from "../../components/eq-ui/EQWindow";
import ContentArea       from "../../components/layout/ContentArea";
import {Loot}            from "../../app/loot";
import LootPopover       from "../../components/LootPopover";
import EqLootCardPreview from "../../components/preview/EQLootCardPreview";

export default {
  name: "Loot",
  components: { EqLootCardPreview, LootPopover, ContentArea, EqWindow },
  data() {
    return {
      loot: {},

      previewedTable: {},
    }
  },
  async mounted() {
    this.loot = await Loot.getLoot()
  },
  methods: {
    showTable(loottable) {
      this.previewedTable = loottable
    },
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
