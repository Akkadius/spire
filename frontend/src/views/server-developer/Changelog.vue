<template>
  <div>
    <eq-window title="Changelog Generator">
      <div class="row">
        <div class="col-2 text-right">
          <b-button
            size="sm"
            variant="outline-warning"
            class="form mt-3"
            @click="generate()"
          >
            <i class="fa fa-refresh mr-1"></i>
            Generate Changelog
          </b-button>
        </div>

        <div class="col-10">
          <div class="eq-alert">
            This will generate changelog notes of all commits formatted since last release
          </div>
        </div>

      </div>
    </eq-window>

    <app-loader :is-loading="loading"/>

    <eq-window
      style="height: 83vh; "
      class="fade-in text-center p-3" v-if="changelog && !loading">
      <button
        class='btn btn-sm btn-dark mb-3'
        @click="copyToClip(changelog)"
      >
        <i class="fa fa-clipboard"></i>
        Copy to Clipboard
      </button>
      <textarea v-model="changelog" style="width: 100%; height: 75vh; overflow-y: scroll"></textarea>
    </eq-window>

  </div>

</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "@/app/api/spire-api";
import ClipBoard  from "@/app/clipboard/clipboard";
import {Notify}   from "@/app/Notify";

export default {
  name: "Changelog",
  components: { EqWindow },
  data() {
    return {
      loading: false,

      changelog: "",
    }
  },
  methods: {
    async generate() {
      this.loading = true;
      const r      = await SpireApi.v1().get(`changelog`)
      if (r.status === 200) {
        this.changelog = r.data.data
        this.loading   = false;
      }
    },
    copyToClip(s) {
      ClipBoard.copyFromText(s)
      console.log(s)

      Notify.toast("Copied to clipboard!");
    },
  }

}
</script>

<style scoped>

</style>
