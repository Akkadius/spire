<template>
  <eq-window title="Dynamic Zone Template Selector" class="p-0">
    <div style="overflow-y: scroll; max-height: 95vh" id="dynamic-zone-template-viewport">

      <div class="mt-3 p-3 text-center" style="padding-bottom: 0px !important;">
        Rows below are loaded from the `dynamic_zone_templates` table
      </div>

      <table
        class="eq-table eq-highlight-rows row-table bordered"
        style="display: table; font-size: 14px; overflow-x: scroll"
        v-if="rows && rows.length > 0"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 50px"></th>
          <th style="width: 40px">ID</th>
          <th>Name</th>

          <th
            v-for="(header, index) in Object.keys(rows[0]).filter((f) => { return !filterColumns.includes(f) })"
            :id="'column-' + header"
            style="text-align: center;"
          >{{ header }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'row-' + e.id"
          :class="(isRowSelected(e) ? 'pulsate-highlight-white' : '')"
          v-for="(e, index) in rows"
          :key="e.id"
          style="height: 50px"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              title="Select"
              @click="selectRow(e)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>

          <td>
            {{ e.id }}
          </td>

          <td>
            {{ e.name }}
          </td>

          <td
            :style="' text-align: center'"
            v-for="(key, colIndex) in Object.keys(e).filter((f) => { return !filterColumns.includes(f) })"
          >
            {{ e[key] }}
          </td>

        </tr>
        </tbody>
      </table>
    </div>
  </eq-window>
</template>

<script>
import EqWindow         from "@/components/eq-ui/EQWindow";
import {SpireApi}       from "../../app/api/spire-api";
import {scrollToTarget} from "@/app/utility/scrollToTarget";
import {DynamicZoneTemplateApi} from "@/app/api";

export default {
  name: "DynamicZoneTemplateSelector",
  components: { EqWindow },
  props: {
    selectedId: {
      type: Number,
      required: false
    }
  },
  data() {
    return {
      rows: [],
      filterColumns: ['id', 'name']
    }
  },
  async mounted() {
    const api = (new DynamicZoneTemplateApi(...SpireApi.cfg()))
    const r   = await api.listDynamicZoneTemplates()
    if (r.status === 200) {
      this.rows = r.data
    }

    if (this.selectedId > 0) {
      scrollToTarget(
        "dynamic-zone-template-viewport",
        'row-' + this.selectedId
      )
    }
  },
  methods: {
    selectRow(entry) {
      this.$emit('input', entry.id);
    },
    isRowSelected(e) {
      return this.selectedId &&
        this.selectedId > 0 &&
        e.id === this.selectedId;
    }
  }
}
</script>

<style>
.row-table td {
  vertical-align: middle !important;
}
</style>
