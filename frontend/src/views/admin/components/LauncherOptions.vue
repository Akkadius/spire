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
                  >{{ tag }}</b-form-tag>
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
                  #first>
                  <!-- This is required to prevent bugs with Safari -->
                  <option disabled value="">Choose a zone...</option>
                </template>
              </b-form-select>
            </template>
          </b-form-tags>
        </b-form-group>

      </div>
    </div>
  </div>
</template>

<script>
  import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";
  import {Zones}         from "@/app/zones";

  export default {
    name: 'LauncherOptions',
    props: ['launcherConfig'],
    data () {
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
      saveLauncherOptions () {
        setTimeout(() => {
          console.log(this.staticZones)

          if (this.staticZones && this.staticZones.length > 0) {
            this.launcher.staticZones = this.staticZones.join(",")
            console.log("lel")
          }

          console.log(this.launcher)

          OcculusClient.postLauncherConfig(this.launcher)
        }, 100)
      }
    }
  }
</script>

<style scoped>

</style>
