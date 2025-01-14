<template>
  <div>
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title">
          System Info
        </h4>
      </div>
      <div class="card-body" style="padding: 0px; overflow-y:scroll">
        <table class="table card-table sysinfo">
          <tbody>

          <tr v-for="i in items">
            <td style="vertical-align: middle">{{ i.key }}</td>
            <td class="text-right">
              <small class="text-muted">
                {{ i.value }}
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
import {SpireApi} from "@/app/api/spire-api";
import Time       from "@/app/time/time";

export default {
  name: 'DashboardSystemInfoV2',
  data() {
    return {
      host: {},

      items: []
    }
  },
  async mounted() {
    let v = []

    let r = await SpireApi.v1().get("admin/system/host")
    if (r.status === 200) {
      this.host = r.data

      const d = r.data
      v.push({ key: "Hostname", value: d.hostname })
      v.push({ key: "OS", value: d.os })
      v.push({ key: "Arch", value: d.kernelArch })
      v.push({ key: "Platform", value: `${d.platform} (${d.platformVersion})` })
      v.push({ key: "Version", value: d.kernelVersion })
      v.push({ key: "Virtualization", value: `${d.virtualizationSystem} ${d.virtualizationRole}` })
      v.push({ key: "Uptime", value: Time.fromNowUnix(d.uptime) })
      v.push({ key: "Processes", value: d.procs })
    }

    this.items = v
  },
  methods: {}
}
</script>

<style scoped>
.sysinfo td {
  padding: 0.5rem;
  vertical-align: top;
  border-top: 1px solid rgba(0, 40, 100, 0.12);
}
</style>
