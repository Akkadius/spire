<template>
  <div class="card">
    <div class="p-0 card-body">
      <div class="mb-3">
        <b-form-checkbox
          v-model="launcher.runSharedMemory"
          name="check-button"
          switch
          @change="saveLauncherOptions()"
        >
          Run Shared Memory
        </b-form-checkbox>
      </div>
      <div class="mb-3">
        <b-form-checkbox
          v-model="launcher.runLoginserver"
          name="check-button"
          switch
          class="custom-control custom-switch"
          @change="saveLauncherOptions()"
        >
          Run Loginserver
        </b-form-checkbox>
      </div>
      <div class="mb-3">
        <b-form-checkbox
          v-model="launcher.runQueryServ"
          name="check-button"
          switch
          @change="saveLauncherOptions()"
        >
          Run QueryServ
        </b-form-checkbox>
      </div>
      <div class="mb-3 mt-4">
        Static Zones

        <b-form-group class="mt-3">
          <!-- Prop `add-on-change` is needed to enable adding tags vie the `change` event -->
          <b-form-tags
            id="tags-component-select"
            v-model="staticZones"
            size="lg"
            tag-pills
            variant="success"
            class="mb-2"
            add-on-change
            no-outer-focus
          >
            <template v-slot="{ tags, inputAttrs, inputHandlers, disabled, removeTag }">
              <ul v-if="tags.length > 0" class="list-inline d-inline-block mb-2">
                <li v-for="tag in tags" :key="tag" class="list-inline-item">
                  <b-form-tag
                    @remove="removeTag(tag); saveLauncherOptions()"
                    :title="tag"
                    :disabled="disabled"
                    variant="info"
                  >{{ tag }}
                  </b-form-tag>
                </li>
              </ul>
              <b-form-select
                v-bind="inputAttrs"
                v-on="inputHandlers"
                @change="saveLauncherOptions()"
                :disabled="disabled || availableOptions.length === 0"
                :options="availableOptions"
              >
                <template
                  #first
                >
                  <!-- This is required to prevent bugs with Safari -->
                  <option disabled value="">Choose a zone...</option>
                </template>
              </b-form-select>
            </template>
          </b-form-tags>
        </b-form-group>

      </div>

      <div class="mt-3">
        <div>
          Min Zone Processes (Ready)
        </div>

        <div>
          <p class="text-muted">
            This is the number of zones that Spire will attempt to keep running <b>without</b> players. For example: if
            you have 10 zones with players in it and your minZoneProcesses is set to 10, you will have 20 total zones
            booted.
          </p>
        </div>
        <b-form-input
          type="number"
          v-model.number="launcher.minZoneProcesses"
          @change="saveLauncherOptions()"
        />
      </div>

    </div>
  </div>
</template>

<script>
import {Zones}    from "@/app/zones";
import {SpireApi} from "@/app/api/spire-api";

export default {
  name: 'LauncherOptions',
  props: ['launcherConfig'],
  data() {
    return {
      launcher: {
        runSharedMemory: false,
        runLoginserver: false,
        runQueryServ: false,
        staticZones: ""
      },

      staticZones: [],
      availableZoneOptions: []
    }
  },
  async created() {
    this.launcher = this.launcherConfig

    if (this.launcher.staticZones && this.launcher.staticZones.length > 0) {
      this.staticZones = this.launcher.staticZones.split(",")
    }

    // zone options
    let options = []
    const zones = await Zones.getZones()
    for (let z of zones) {
      options.push(z.short_name)
    }
    this.availableZoneOptions = options
  },
  watch: {
    launcherConfig: function (newValue) {
      this.launcher = newValue
    }
  },
  computed: {
    availableOptions() {
      return this.availableZoneOptions.filter(opt => this.staticZones.indexOf(opt) === -1)
    }
  },
  methods: {
    saveLauncherOptions() {
      setTimeout(async () => {
        if (this.staticZones && this.staticZones.length > 0) {
          this.launcher.staticZones = this.staticZones.join(",")
        }

        try {
          await SpireApi.v1().post('admin/launcherconfig', this.launcher)
        } catch (e) {
          console.log(e)
        }

      }, 100)
    }
  }
}
</script>

<style scoped>

</style>
