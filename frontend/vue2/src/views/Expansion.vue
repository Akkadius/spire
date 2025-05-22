<template>
  <content-area style="padding: 0px !important">
    <eq-window title="Expansion Stats">
      <div style="overflow-y: hidden; overflow-x: scroll">
        Filters by Expansion and Table
        <table
          class="eq-table eq-highlight-rows expansion-data-table"
          style="display: table; font-size: 14px;"
        >
          <thead
            class="eq-table-floating-header"
          >
          <tr>
            <th class="sticky-first-column">
              Table
            </th>
            <th v-for="(e, index) in expansions">
              <expansion-icon :expansion="index"/>
            </th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="t in tables"
            :key="t"
          >
            <td class="sticky-first-column">
              {{ t }}
            </td>
            <td v-for="(e, index) in expansions">
              {{ expansionData[t] && expansionData[t][index] ? expansionData[t][index] : 0 }}
            </td>
          </tr>
          </tbody>
        </table>
      </div>

    </eq-window>
  </content-area>
</template>

<script>
import ContentArea       from "../components/layout/ContentArea";
import EqWindow          from "../components/eq-ui/EQWindow";
import {SpireApi}        from "../app/api/spire-api";
import {EXPANSIONS_FULL} from "../app/constants/eq-expansions";
import ExpansionIcon     from "../components/preview/ExpansionIcon";

export default {
  name: "Expansion",
  components: { ExpansionIcon, EqWindow, ContentArea },
  data() {
    return {
      tables: [],
      expansionData: [],
      expansions: EXPANSIONS_FULL
    }
  },
  mounted() {
    let expansionData = {}
    let tables        = []
    SpireApi.v1().get(`query/expansion-stats`).then((r) => {
      if (r.data && r.data.data) {
        for (const [table, values] of Object.entries(r.data.data)) {
          // console.log(table, values)
          tables.push(table)
          for (let value of values) {
            if (typeof expansionData[table] === 'undefined') {
              expansionData[table] = {}
            }
            if (typeof expansionData[table][value.min_expansion] === 'undefined') {
              expansionData[table][value.min_expansion] = 0
            }
            expansionData[table][value.min_expansion] = value.count
          }
        }
        // console.log(expansionData)
      }

      this.expansionData = expansionData
      this.tables        = tables
    });
  }
}
</script>

<style>
.sticky-first-column {
  text-align: center;
  position: sticky;
  z-index: 9999;
  background-color: rgb(25, 31, 41);
  left: 0px;
  font-weight: bold;
}

.expansion-data-table th, .expansion-data-table td {
  text-align: center !important;
  padding: 2px !important;
}
</style>
