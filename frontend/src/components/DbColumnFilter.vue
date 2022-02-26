<template>
  <div>
    <b-button
      size="sm"
      variant="outline-warning"
      @click="queueColumnFilter"
      v-if="!queuedAddFilter"
    >
      <i class="fa fa-plus"></i>
    </b-button>

    <div v-if="queuedAddFilter" class="d-inline-block">
      <div class="input-group d-inline-flex">
        <div class="input-group-prepend">
          <select class="form-control" v-model="fieldToAdd">
            <option
              v-for="field in columns"
              :key="field"
            >{{ field }}
            </option>
          </select>
        </div>
        <select class="form-control" v-model="operator">
          <option
            v-for="field in filterOptions"
            :key="field"
          >{{ field }}
          </option>
        </select>
        <div class="input-group-append">
          <input
            type="text"
            class="form-control"
            v-on:keyup.enter="addFilter"
            v-model="fieldValue"
          >
        </div>
      </div>
    </div>

    <b-button
      size="sm"
      variant="outline-warning"
      @click="addFilter"
      class="d-inline-block ml-3"
      v-if="queuedAddFilter"
    >
      <i class="fa fa-plus"></i> Add Filter
    </b-button>

    <b-button
      size="sm"
      variant="outline-danger"
      @click="queuedAddFilter = false"
      class="d-inline-block ml-3"
      v-if="queuedAddFilter"
    >
      <i class="fa fa-ban"></i> Cancel
    </b-button>

    <b-button
      size="sm"
      variant="warning"
      @click="removeFilter(filter)"
      class="d-inline-block ml-3"
      v-for="filter in filters"
      :key="filter.f + '-' + filter.o + '-' + filter.v"
    >
      <i class="fa fa-remove"></i> {{ formatFilterHuman(filter) }}
    </b-button>

  </div>
</template>

<script>
import util from "util";

export default {
  name: "DbColumnFilter",
  props: {
    columns: {
      type: Array,
      required: true
    },
    setFilters: {
      type: Array,
      required: false
    }
  },
  data() {
    return {

      // currently held filters
      filters: [],

      // queued filter to add fields
      fieldToAdd: "",
      operator: "",
      fieldValue: "",

      // flag to set whether we're in the state to add a filter not
      queuedAddFilter: false,

      // filter options
      filterOptions: ["=", "!=", "<=", "<", ">", ">=", "like", "notlike"],
    }
  },
  watch: {
    setFilters: function () {
      this.filters         = []
      this.queuedAddFilter = false;
      this.loadSetFilters()
    }
  },
  mounted() {
    this.reset()
    this.loadSetFilters()
  },
  methods: {
    loadSetFilters() {
      // if we have filters being passed down from parent, set on component state
      this.filters = this.setFilters
    },

    reset() {
      this.fieldValue = ""

      setTimeout(() => {
        // set defaults for selects
        if (this.columns) {
          this.fieldToAdd = this.columns[0]
          this.operator   = this.filterOptions[0]
        }
      }, 100)
    },

    updateParent() {
      // emit filters to parent
      let queryBuilderFilters = []
      this.filters.forEach((filter) => {
        let operator = filter.o

        queryBuilderFilters.push({
          f: filter.f,
          o: operator,
          v: filter.v,
        })
      })

      this.$emit('input', queryBuilderFilters);
    },

    removeFilter(filter) {
      this.filters = this.filters.filter((f) => {
        let a = util.format("%s-%s-%s", f.field, f.operator, f.value)
        let b = util.format("%s-%s-%s", filter.field, filter.operator, filter.value)

        return a !== b
      })

      this.updateParent()
    },

    formatFilterHuman(filter) {
      return util.format(
        "%s %s %s",
        filter.f,
        filter.o,
        filter.v,
      )
    },

    addFilter() {
      this.filters.push(
        {
          f: this.fieldToAdd,
          o: this.operator,
          v: this.fieldValue,
        }
      )

      this.queuedAddFilter = false
      this.reset()
      this.updateParent()
    },
    queueColumnFilter() {
      this.queuedAddFilter = true
    }
  }
}
</script>

<style scoped>

</style>
