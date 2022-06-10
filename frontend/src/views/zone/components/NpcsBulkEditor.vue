<template>
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
          v-if="(setValue !== '' && isDataTypeNumber(selectedField)) || (setValue === '' && !isDataTypeNumber(selectedField))"
        >
          <i class="fa fa-edit"></i> Write
        </button>

      </div>
    </div>

    <!-- Number -->
<!--    <div class="row" v-if="isDataTypeNumber(selectedField)">-->
<!--      <div class="col-4 text-right m-0 p-0 mt-2">-->
<!--        Select Field-->
<!--      </div>-->
<!--      <div class="col-5">-->
<!--        <select-->
<!--          class="form-control"-->
<!--          @change="selectField"-->
<!--          v-model="selectedField"-->
<!--        >-->
<!--          <option-->
<!--            v-for="field in npcTypeFields"-->
<!--            :key="field"-->
<!--          >{{ field }}-->
<!--          </option>-->
<!--        </select>-->
<!--      </div>-->
<!--    </div>-->




  </eq-window>
</template>

<script>
import EqWindowComplex     from "../../../components/eq-ui/EQWindowComplex";
import EqWindow            from "../../../components/eq-ui/EQWindow";
import {NpcTypeApi}        from "../../../app/api";
import {SpireApiClient}    from "../../../app/api/spire-api-client";
import {SpireQueryBuilder} from "../../../app/api/spire-query-builder";

export default {
  name: "NpcsBulkEditor",
  components: { EqWindow, EqWindowComplex },
  data() {
    return {
      // v-model
      selectedField: "",
      setValue: "",

      // fields
      npcTypeFields: [],
    }
  },
  methods: {
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

    selectField() {
      this.setValue = ""

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
  created() {
    // non-reactive properties
    this.exampleNpcRecord = {}

    // get example npc record for determining data types
    let fields            = []
    const api             = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))
    let builder           = (new SpireQueryBuilder())
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
