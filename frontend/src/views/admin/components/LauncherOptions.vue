<template>
  <div>


    <table
      class="eq-table bordered mb-0"
    >
      <thead class="eq-table-floating-header">
      <tr>
        <td class="font-weight-bold p-3">Options</td>
      </tr>
      </thead>

      <tbody>
      <tr>
        <td>
          <eq-checkbox
            :fade-when-not-true="true"
            class="d-inline-block mr-3"
            :true-value="true"
            :false-value="false"
            v-model="launcher.updateOpcodesOnStart"
            @change="saveLauncherOptions()"
          />
          Update Server Patches (Opcodes) On Start)
        </td>
      </tr>
      <tr>
        <td>
          <eq-checkbox
            :fade-when-not-true="true"
            class="d-inline-block mr-3"
            :true-value="true"
            :false-value="false"
            v-model="launcher.runSharedMemory"
            @change="saveLauncherOptions()"
          />
          Run Shared Memory (Recommended)
        </td>
      </tr>
      <tr>
        <td>
          <eq-checkbox
            :fade-when-not-true="true"
            class="d-inline-block mr-3"
            :true-value="true"
            :false-value="false"
            v-model="launcher.runUcs"
            @change="saveLauncherOptions()"
          />
          Run UCS (Optional)
        </td>
      </tr>
      <tr>
        <td>
          <eq-checkbox
            :fade-when-not-true="true"
            class="d-inline-block mr-3"
            :true-value="true"
            :false-value="false"
            v-model="launcher.runLoginserver"
            @change="saveLauncherOptions()"
          />
          Run Loginserver (Optional)
        </td>
      </tr>
      <tr>
        <td>
          <eq-checkbox
            :fade-when-not-true="true"
            class="d-inline-block mr-3"
            :true-value="true"
            :false-value="false"
            v-model="launcher.runQueryServ"
            @change="saveLauncherOptions()"
          />
          Run QueryServ (Optional)
        </td>
      </tr>
      </tbody>
    </table>


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
                  variant="success"
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
              class="mr-3"
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

      <div class="mt-3">
        <div>
          Days to keep log files (7 days default)
        </div>

        <div>
          <p class="text-muted">
            Files older than this will be deleted periodically. Set to -1 to disable.
          </p>
        </div>
        <b-form-input
          type="number"
          v-model.number="launcher.deleteLogFilesOlderThanDays"
          @change="saveLauncherOptions()"
        />
      </div>

    </div>
  </div>
</template>

<script>
import {Zones}    from "@/app/zones";
import {SpireApi} from "@/app/api/spire-api";
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";

export default {
  name: 'LauncherOptions',
  components: { EqCheckbox },
  props: ['launcherConfig'],
  data() {
    return {
      launcher: {
        runSharedMemory: false,
        runLoginserver: false,
        runQueryServ: false,
        runUcs: true,
        updateOpcodesOnStart: true,
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

    if (typeof this.launcher.updateOpcodesOnStart === 'undefined') {
      this.launcher.updateOpcodesOnStart = true
    }

    if (typeof this.launcher.deleteLogFilesOlderThanDays !== 'undefined' && this.launcher.deleteLogFilesOlderThanDays === 0) {
      this.launcher.deleteLogFilesOlderThanDays = 7
    }

    if (typeof this.launcher.runUcs === 'undefined') {
      this.launcher.runUcs = true
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
