<template>
  <div>
    <eq-window title="Message of the Day">
      <div class="eq-alert">
        <div>
          <i class="fa fa-info-circle"></i>
          Message of the day is what your players see when they first log in
        </div>
      </div>

      <div class="mt-3">
        <textarea class="form-control" rows="7" v-model="motd.value"></textarea>
      </div>

      <button type="submit" class="btn btn-dark btn-sm ml-auto mt-3" @click="submit()">
        <i class="fe fe-save"></i>
        Save
      </button>


      <div
        class="row justify-content-center"
        style="position: absolute; bottom: 5%; z-index: 9999999; width: 100%"
      >
        <div class="col-4">
          <info-error-banner
            style="width: 100%"
            :slim="true"
            :notification="notification"
            :error="error"
            @dismiss-error="error = ''"
            @dismiss-notification="notification = ''"
            class="mt-3"
          />
        </div>
      </div>
    </eq-window>
  </div>
</template>

<script>
import {SpireApi}      from "@/app/api/spire-api";
import {VariableApi}   from "@/app/api/api/variable-api";
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";

export default {
  components: { InfoErrorBanner, EqWindow },
  data() {
    return {
      motd: {},
      loaded: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async created() {
    this.loaded = true

    let r = await (new VariableApi(...SpireApi.cfg())).listVariables()
    if (r.status === 200) {
      this.motd = r.data.find((e) => {
        return e.varname === 'MOTD'
      })
    }
  },
  methods: {
    submit: async function () {
      try {
        let r = await (new VariableApi(...SpireApi.cfg())).updateVariable(
          {
            id: this.motd.id,
            variable: this.motd
          }
        )
        if (r.status === 200) {
          this.notification = "Message of the day updated!"
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
