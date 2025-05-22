<template>
  <eq-window
    title="System Info"
    class="p-0 pt-3 mb-3"
  >
    <table class="eq-table bordered eq-highlight-rows mb-0">
      <tbody>
      <tr v-for="i in items">
        <td class="text-right font-weight-bold">{{ i.key }}</td>
        <td class="text-left">
          <small class="text-muted">
            {{ i.value }}
          </small>
        </td>
      </tr>

      </tbody>
    </table>
  </eq-window>
</template>

<script>
import {SpireApi} from "@/app/api/spire-api";
import Time       from "@/app/time/time";
import EqWindow   from "@/components/eq-ui/EQWindow.vue";

export default {
  name: 'DashboardSystemInfoV2',
  components: { EqWindow },
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
      v.push({ key: "Uptime", value: Time.humanizeUnix((Date.now() / 1000) + d.uptime) })
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
