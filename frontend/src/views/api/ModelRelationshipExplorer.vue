<template>
  <content-area style="padding: 0px !important">
    <eq-window title="API Model Relationship Explorer">
      Select model with relationships

      <b-form-select
        v-model="selected"
        :options="options"
        @change="draw"
      />

      <div class="row mt-3" v-if="selectedModel && Object.keys(selectedModel).length > 0">
        <div class="col-12">
          <b>Table</b> {{selectedModel.table}}
          <b>Model</b> {{selectedModel.model_name}}
        </div>
      </div>

      <pre
        style="width: 100%; height: 75vh; overflow-y: scroll"
        v-if="selectedModel && Object.keys(selectedModel).length > 0 && selectedModel.relationships.length > 0" class="mt-3">{{ selectedModel.relationships.join("\n") }}</pre>
    </eq-window>
  </content-area>
</template>

<script>
import EqWindow     from "../../components/eq-ui/EQWindow";
import ContentArea  from "../../components/layout/ContentArea";
import {SpireApi}   from "../../app/api/spire-api";
import {HttpStatus} from "../../app/api/http-status";

export default {
  name: "ModelRelationshipExplorer",
  components: { ContentArea, EqWindow },
  data() {
    return {
      selected: "",
      options: [],

      models: [],

      selectedModel: {}
    }
  },
  methods: {
    draw() {
      for (let model of this.models) {
        if (this.selected === model.model_name) {
          this.selectedModel = model
        }
      }
    }
  },
  async mounted() {
    const r = await SpireApi.v1().get('/models')
    if (r.status === HttpStatus.OK) {
      this.models = r.data.sort((a, b) => a.model_name.localeCompare(b.model_name));

      let options = []
      for (let model of this.models) {
        if (model.relationships.length > 0) {
          options.push({
            value: model.model_name,
            text: `${model.model_name}` + (model.relationships.length > 0 ? ` (${model.relationships.length})` : '')
          })
        }
      }

      this.options = options

    }
  }
}
</script>

<style scoped>

</style>
