<template>
  <div class='eq-window-nested-blue' style="width: 100%; overflow-y: scroll">
    <div v-if="data && data.length === 0" class="p-3">
      There were no entries to be shown
    </div>

    <table
      class="eq-table eq-highlight-rows"
      style="display: table; font-size: 14px; overflow-x: scroll "
      v-if="data && data.length > 0"
    >
      <thead>
      <tr>
        <th
          v-for="header in Object.keys(data[0])"
          v-if="!doesColumnHaveObjects(header)"
        >{{ header }}
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(row, index) in data" :key="index">
        <td
          style="width: 1%; white-space: nowrap;"
          v-for="key in Object.keys(row)"
          v-if="(typeof row[key] !== 'undefined') && !(typeof row[key] === 'object' && row[key] !== null && Object.keys(row[key]))"
        >
          {{ row[key] }}
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "EqAutoTable",
  props: {
    data: {
      required: false
    }
  },
  methods: {
    doesColumnHaveObjects(column) {
      if (typeof column === 'object') {
        return true
      }

      return this.data.find((row) => {
        return typeof row[column] === 'object' && row[column] !== null && Object.keys(row[column])
      })
    }
  }
}
</script>

<style scoped>

</style>
