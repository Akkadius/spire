<template>
  <div>
    <eq-window title="Changelog Generator">
      <div class="row">
        <div class="col-3">
          Days Back
          <input type="text" class="form-control" v-model="daysBack">
        </div>
        <div class="col-2">
          <b-button
            size="sm"
            variant="outline-warning"
            class="form mt-4"
            @click="generate()"
          >
            <i class="fa fa-refresh mr-1"></i>
            Generate Changelog
          </b-button>
        </div>
      </div>
    </eq-window>

    <app-loader :is-loading="loading"/>

    <eq-window
      style="height: 80vh; "
      class="fade-in text-center p-3" v-if="changelog && !loading">
      <button
        class='btn btn-sm btn-outline-warning mb-3'
        @click="copyToClip(changelog)"
      >
        <i class="fa fa-clipboard"></i>
        Copy to Clipboard
      </button>
      <textarea v-model="changelog" style="width: 100%; height: 70vh; overflow-y: scroll"></textarea>
    </eq-window>

  </div>

</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "@/app/api/spire-api";
import ClipBoard  from "@/app/clipboard/clipboard";

export default {
  name: "Changelog",
  components: { EqWindow },
  data() {
    return {
      loading: false,

      daysBack: 5,

      changelog: "",
    }
  },
  methods: {
    async generate() {
      this.loading = true;
      const r      = await SpireApi.v1().get(`changelog/${this.daysBack}`)
      if (r.status === 200) {
        this.changelog = r.data.data
        this.loading   = false;
      }
    },
    copyToClip(s) {
      ClipBoard.copyFromText(s)
      console.log(s)

      this.$bvToast.toast("Copied to clipboard!", {
        title: "Copy",
        autoHideDelay: 2000,
        solid: true
      })
    },
  }

}
</script>

<style scoped>

</style>
