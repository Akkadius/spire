<template>
  <div>
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title" v-if="cpu && cpu.cpu_percents">CPU (s) ({{ cpu.cpu_percents.length }})
        </h4>
      </div>
      <div class="card-body" style="padding: 15px; overflow-y:scroll">
        <ul class="list-unstyled list-separated">
          <li class="list-separated-item">

            <div class="row align-content-center">
              <div class="col-12 text-center">
                <small class="text-muted" v-if="cpu && cpu.info && cpu.info[0]">
                  {{cpu.info[0].modelName}} ({{(cpu.info[0].mhz / 1000).toFixed(1)}} Ghz)
                </small>

              </div>
            </div>

            <div class="row align-items-center mt-3">

              <div
                class="col"
                v-for="(l, index) in cpu.cpu_percents"
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
                        :style="'color: ' + getCpuLoadColor(parseInt(l))"
                      >{{ (Math.round(l * 100) / 100) }}%</small>
                    </div>
                  </div>

                  <div class="progress progress-sm">
                    <div
                      class="progress-bar bg-green"
                      role="progressbar"
                      v-bind:style="{ width: (Math.round(l * 100) / 100) + '%'}"
                      :aria-valuenow="(Math.round(l * 100) / 100)"
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
import * as util  from 'util'
import {SpireApi} from "@/app/api/spire-api";

export default {
  name: 'DashboardCpuInfo',
  data() {
    return {
      cpu: {},

      timer: null,
    }
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  mounted() {
    this.fetchStats()
    this.timer = setInterval(this.fetchStats, 1000)
  },
  methods: {
    fetchStats() {
      if (!document.hidden) {
        SpireApi.v1().get("admin/system/cpu").then((r) => {
          if (r.status === 200) {
            this.cpu = r.data
          }
        })
      }
    },

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
}
</script>
