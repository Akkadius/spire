<template>
  <div v-if="sysinfo && Object.keys(sysinfo).length > 0">
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title" v-if="cpuLoad.length">CPU (s) ({{ cpuLoad.length }})
        </h4>
      </div>
      <div class="card-body" style="padding: 15px; overflow-y:scroll">
        <ul class="list-unstyled list-separated">
          <li class="list-separated-item">

            <div class="row align-content-center">
              <div class="col-12 text-center">
                <small class="text-muted">
                  {{ cpuInfo }}
                </small>

              </div>
            </div>

            <div class="row align-items-center mt-3" v-if="Object.keys(sysinfo).length > 0">

              <div
                class="col"
                v-for="(cpu, index) in cpuLoad"
                :key="index"
                style="flex-basis:unset;"
              >
                <div class="d-block">

                  <div class="clearfix">
                    <div class="float-left">
                      <small class="text-muted">
                        #{{ index + 1 }}
                      </small>
                    </div>
                    <div class="float-right">
                      <small
                        :style="'color: ' + getCpuLoadColor(parseInt(cpu.load))"
                      >{{ (Math.round(cpu.load * 100) / 100) }}%</small>
                    </div>
                  </div>

                  <div class="progress progress-sm">
                    <div
                      class="progress-bar bg-green"
                      role="progressbar"
                      v-bind:style="{ width: (Math.round(cpu.load * 100) / 100) + '%'}"
                      :aria-valuenow="(Math.round(cpu.load * 100) / 100)"
                      aria-valuemin="0"
                      aria-valuemax="100"
                    >
                    </div>
                  </div>

                </div>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import * as util from 'util'

export default {
  name: 'DashboardCpuInfo',
  props: {
    sysinfo: {}
  },
  methods: {

    getCpuLoadColor(load) {
      if (load > 80) {
        return 'red'
      }
      if (load > 50) {
        return 'orange'
      }

      return '#95aac9'
    },

    /**
     * @param input
     * @param length
     */
    truncate: function (input, length) {
      if (input.length > length) {
        return input.substring(0, length) + '...'
      } else {
        return input
      }
    }
  },
  computed: {
    cpuInfo: function () {
      return util.format('%s %s (%s) %sGhz',
        this.sysinfo.cpu.info.manufacturer,
        this.sysinfo.cpu.info.brand,
        this.sysinfo.cpu.info.cores,
        this.sysinfo.cpu.info.speed
      )
    },
    cpuLoad: function () {
      return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.cpu.load.cpus : {})
    }
  }
}
</script>
