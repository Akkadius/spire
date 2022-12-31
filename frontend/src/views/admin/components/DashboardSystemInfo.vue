<template>
  <div class="col-md-6 col-lg-6">
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title">
          System Info
        </h4>
      </div>
      <div class="card-body" style="height: 15rem; padding: 0px; overflow-y:scroll">
        <table class="table card-table sysinfo" v-if="Object.keys(sysinfo).length > 0">
          <tbody>

          <!-- OS -->
          <tr>
            <td style="vertical-align: middle">OS</td>
            <td class="text-right">
              <small class="text-muted" style="font-size: 10px !important">{{osDisplay}}</small>
            </td>
          </tr>

          <!-- Disk -->
          <tr v-if="Object.keys(sysinfo.disk).length > 0">
            <td style="vertical-align: middle">Disk</td>
            <td class="text-right" style="font-size: 10px !important">
              {{ bytesToGbytes(diskUsageBytes) }} /
              {{ bytesToGbytes(diskTotalBytes) }} GB
              ({{ diskUtilizationPercent }}%)
              <div class="progress progress-sm">
                <div class="progress-bar bg-green"
                     role="progressbar"
                     v-bind:style="{ width: diskUtilizationPercent + '%'}"
                     :aria-valuenow="diskUtilizationPercent"
                     aria-valuemin="0"
                     aria-valuemax="100">
                </div>
              </div>
              <small class="text-muted" style="font-size: 10px !important">
                R {{bytesToMbytes(diskReadBytes)}} /
                W {{bytesToMbytes(diskWriteBytes)}} MB/s
              </small>
            </td>
          </tr>

          <!-- Memory -->
          <tr>
            <td style="vertical-align: middle">Memory</td>
            <td class="text-right" style="font-size: 10px !important">
              {{ bytesToGbytes(memUsageBytes) }} /
              {{ bytesToGbytes(memTotalBytes) }} GB
              ({{ Math.floor((memUsageBytes / memTotalBytes) * 100) }}%)
              <div class="progress progress-sm">
                <div class="progress-bar bg-green"
                     role="progressbar"
                     v-bind:style="{ width: memUsagePercent + '%'}"
                     :aria-valuenow="memUsagePercent"
                     aria-valuemin="0"
                     aria-valuemax="100">
                </div>
              </div>
            </td>
          </tr>

          <!-- Network -->
          <tr>
            <td style="vertical-align: middle">Network</td>
            <td class="text-right">
              <small class="text-muted" style="font-size: 10px !important">
                Rec {{bytesToMbytes(networkReadBytes)}} /
                Trx {{bytesToMbytes(networkWriteBytes)}} MB/s
              </small>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import * as util from 'util'

  export default {
    name: 'DashboardSystemInfo',
    props: {
      sysinfo: {}
    },
    methods: {
      bytesToGbytes: function (bytes) {
        return parseFloat(bytes / 1024 / 1024 / 1024).toFixed(2)
      },

      bytesToMbytes: function (bytes) {
        return parseFloat(bytes / 1024 / 1024).toFixed(2)
      },

      ucfirst: function (string) {
        return string.charAt(0).toUpperCase() + string.slice(1)
      }
    },
    computed: {
      cpuLoadDisplay: function () {
        return (Object.keys(this.sysinfo).length > 0 ? (Math.round(this.sysinfo.cpu.load.currentload * 100) / 100) : 0)
      },
      diskUsageBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.disk.fs.size[0].used : 0)
      },
      diskTotalBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.disk.fs.size[0].size : 0)
      },
      memUsageBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.mem.used : 0)
      },
      memTotalBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.mem.total : 0)
      },
      memUsagePercent: function () {
        return Math.floor((this.memUsageBytes / this.memTotalBytes) * 100)
      },
      diskWriteBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.disk.fs.stats.wx_sec : 0)
      },
      diskReadBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.disk.fs.stats.rx_sec : 0)
      },
      diskUtilizationPercent: function () {
        return Math.floor(this.diskUsageBytes / this.diskTotalBytes * 100)
      },
      networkWriteBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.network.stats[0].tx_sec : 0)
      },
      networkReadBytes: function () {
        return (Object.keys(this.sysinfo).length > 0 ? this.sysinfo.network.stats[0].rx_sec : 0)
      },
      osDisplay: function () {
        return util.format(
          '%s %s %s (%s) %s',
          this.ucfirst(this.sysinfo.os.platform),
          this.sysinfo.os.distro,
          this.sysinfo.os.release,
          this.sysinfo.os.codename,
          this.sysinfo.os.arch
        )
      }
    }
  }
</script>

<style scoped>
  .sysinfo td {
    padding:        0.5rem;
    vertical-align: top;
    border-top:     1px solid rgba(0, 40, 100, 0.12);
  }
</style>
