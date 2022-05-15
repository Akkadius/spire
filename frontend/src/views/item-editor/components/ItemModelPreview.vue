<template>
  <div>
    <div v-if="previewId === 0" class="mb-3">
      Item model preview not found...
    </div>
    <div
      v-if="previewId > 0"
      style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"
    >
      <span :class="'mt-2 mb-2 fade-in object-ctn-' + previewId"></span>
    </div>
  </div>
</template>

<script>
import * as util from "util";
import EqAssets  from "../../../app/eq-assets/eq-assets";

export default {
  name: "ItemModelPreview",
  data() {
    return {
      previewId: 0,
    }
  },
  watch: {
    id: {
      handler: function (val, oldVal) {
        this.render()
      },
    },
  },
  methods: {
    async render() {
      this.previewId      = 0;
      const passedInModel = this.id.replace("IT", "")

      const r = await EqAssets.getItemModelFileNames()
      r.forEach((fileName) => {
        const modelId =
                util.format(
                  "%s",
                  fileName.toString()
                    .replace("CTN_", "")
                    .replace(".png", "")
                    .trim()
                );

        if (passedInModel === modelId) {
          this.previewId = modelId;
          return false
        }
      })
    }
  },
  async created() {
    await this.render()
  },
  props: {
    id: {
      required: true,
      type: String,
    },
  }
}
</script>

<style scoped>

</style>
