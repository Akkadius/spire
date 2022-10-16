<template>
  <div>
    <eq-window
      title="Bulk Editing"
      class="minified-inputs mb-0"
    >
      <!-- Select field -->
      <div class="row">
        <div class="col-4 text-right m-0 p-0 mt-3">
          Select Field
        </div>
        <div class="col-5">
          <select
            class="form-control"
            @change="selectField"
            v-model="selectedField"
          >
            <option
              v-for="field in npcTypeFields"
              :key="field"
            >{{ field }}
            </option>
          </select>
        </div>
      </div>

      <!-- Set all values to -->
      <div class="row">
        <div class="col-4 text-right m-0 p-0 mt-3">
          Set all values to
        </div>
        <div class="col-5">
          <b-input
            class="form-control"
            @keyup="setValuesToPreview"
            v-model="setValue"
          ></b-input>
        </div>
        <div class="col-2">
          <button
            class='btn btn-outline-warning btn-sm mt-2'
            @click="setValuesTo"
            v-if="(setValue !== '' && isDataTypeNumber(selectedField)) || (!isDataTypeNumber(selectedField))"
          >
            <i class="fa fa-edit"></i> Write
          </button>

        </div>
      </div>

      <!-- Min / Max -->
      <div class="row" v-if="isDataTypeNumber(selectedField)">
        <div
          class="col-4 text-right m-0 p-0 mt-3"
        >
          Random Between Min / Max
        </div>
        <div class="col-3 text-center">
          <b-input
            class="form-control"
            type="number"
            placeholder="Min"
            @keyup="setMinMaxValuesToPreview"
            v-model="setMin"
          />
        </div>
        <div class="col-3 text-center">
          <b-input
            placeholder="Max"
            type="number"
            class="form-control"
            @keyup="setMinMaxValuesToPreview"
            v-model="setMax"
          />
        </div>
        <div class="col-2">
          <button
            class='btn btn-outline-warning btn-sm mt-2'
            @click="setMinMaxValuesTo"
            v-if="(setMax !== '' && setMin !== '')"
          >
            <i class="fa fa-edit"></i> Write
          </button>

        </div>
      </div>

      <!-- Percentage -->
      <div class="row" v-if="isDataTypeNumber(selectedField)">
        <div
          class="col-4 text-right m-0 p-0 mt-3"
        >
          Percentage
        </div>
        <div class="col-3 text-center">
          <b-input
            class="form-control"
            type="number"
            step="0.1"
            @keyup="setPercentageToPreview"
            v-model="setPercentage"
          />
        </div>
        <div class="col-2 p-0">
          <div
            v-if="setPercentage"
            class="mt-3 d-inline-block mr-3"
          >{{ Math.round(setPercentage * 100) }}%
          </div>
          <button
            class='btn btn-outline-warning btn-sm'
            @click="setPercentageTo"
            v-if="setPercentage"
          >
            <i class="fa fa-edit"></i> Write
          </button>

        </div>
      </div>

      <!-- Number -->
      <!--        <div class="row" v-if="isDataTypeNumber(selectedField)">-->
      <!--          <div class="col-4 text-right m-0 p-0 mt-2">-->
      <!--            Select Field-->
      <!--          </div>-->
      <!--          <div class="col-5">-->
      <!--            <select-->
      <!--              class="form-control"-->
      <!--              @change="selectField"-->
      <!--              v-model="selectedField"-->
      <!--            >-->
      <!--              <option-->
      <!--                v-for="field in npcTypeFields"-->
      <!--                :key="field"-->
      <!--              >{{ field }}-->
      <!--              </option>-->
      <!--            </select>-->
      <!--          </div>-->
      <!--        </div>-->


    </eq-window>

    <eq-window
      title="Bulk Editing Results"
      class="mt-5 pr-1"
      v-if="editFeedbackLocal && editFeedbackLocal.length > 0"
    >
      <div
        class="text-center"
        style="height: 65vh; overflow-y: scroll; overflow-x: hidden;"
      >

        <div class="mb-3 font-weight-bold">
          Changed ({{ editFeedbackLocal.length }}) NPC(s)
        </div>

        <div
          v-for="m in editFeedbackLocal"
        >
          {{ m }}
        </div>
      </div>
    </eq-window>
  </div>
</template>

<script>
import EqWindowComplex     from "../../../components/eq-ui/EQWindowComplex";
import EqWindow            from "../../../components/eq-ui/EQWindow";
import {NpcTypeApi}        from "../../../app/api";
import {SpireApi}          from "../../../app/api/spire-api";
import {SpireQueryBuilder} from "../../../app/api/spire-query-builder";

export default {
  name: "NpcsBulkEditor",
  components: { EqWindow, EqWindowComplex },
  props: {
    editFeedback: {
      type: Array,
      required: false
    },
  },

  watch: {
    editFeedback: {
      handler(newVal) {
        this.editFeedbackLocal = newVal
      },
      deep: true
    },
    setMin: {
      handler(newVal) {
        const min = parseInt(this.setMin)
        const max = parseInt(this.setMax)

        if (min !== 0 && min > max) {
          setTimeout(() => {
            this.setMin = this.setMax
          }, 10)
        }
      },
    },
    setMax: {
      handler(newVal) {
        const min = parseInt(this.setMin)
        const max = parseInt(this.setMax)

        if (max !== 0 && max < min) {
          setTimeout(() => {
            this.setMax = this.setMin
          }, 10)
        }
      },
    },
  },

  data() {
    return {
      // v-model
      selectedField: "",
      setValue: "",

      // min - max
      setMin: "",
      setMax: "",

      // percentage
      setPercentage: 1,

      // fields
      npcTypeFields: [],

      editFeedbackLocal: [],
    }
  },
  methods: {

    // set all values to
    setValuesToPreview() {
      this.$emit(
        'set-values-preview',
        {
          field: this.selectedField,
          value: this.setValue
        }
      );
    },
    setValuesTo() {
      this.$emit(
        'set-values-commit',
        {
          field: this.selectedField,
          value: this.setValue
        }
      );
    },

    // set min / max values to
    setMinMaxValuesToPreview() {
      this.$emit(
        'set-min-max-values-preview',
        {
          field: this.selectedField,
          min: this.setMin,
          max: this.setMax,
        }
      );
    },
    setMinMaxValuesTo() {
      this.$emit(
        'set-min-max-values-commit',
        {
          field: this.selectedField,
          min: this.setMin,
          max: this.setMax,
        }
      );
    },

    // set min / max values to
    setPercentageToPreview() {
      this.$emit(
        'set-percentage-values-preview',
        {
          field: this.selectedField,
          percentage: this.setPercentage,
        }
      );
    },
    setPercentageTo() {
      this.setPercentage = 1

      this.$emit(
        'set-percentage-values-commit',
        {
          field: this.selectedField,
          percentage: this.setPercentage,
        }
      );
    },

    selectField() {
      this.setValue = ""

      this.editFeedbackLocal = []

      // reset
      this.$emit(
        'set-values-preview',
        {
          field: "",
          value: ""
        }
      );

      this.setValuesToPreview()

      this.$emit('field-selected', this.selectedField);
    },

    isDataTypeNumber(column) {
      if (typeof this.exampleNpcRecord[column] !== "undefined" && Number.isInteger(this.exampleNpcRecord[column])) {
        return true
      }

      return false
    }
  },
  mounted() {
    this.editFeedbackLocal = this.editFeedback
  },
  created() {
    // non-reactive properties
    this.exampleNpcRecord = {}

    // get example npc record for determining data types
    let fields  = []
    const api   = (new NpcTypeApi(...SpireApi.cfg()))
    let builder = (new SpireQueryBuilder())
    api.listNpcTypes(
      builder.orderBy(["id"]).limit(1).get()
    ).then((r) => {
      if (r.status === 200) {
        this.exampleNpcRecord = r.data[0]

        // get field names
        for (let f of Object.keys(this.exampleNpcRecord)) {
          if (["id"].includes(f)) {
            continue;
          }

          fields.push(f)
        }

        this.npcTypeFields = fields
      }
    })
  },
}
</script>

<style scoped>

</style>
