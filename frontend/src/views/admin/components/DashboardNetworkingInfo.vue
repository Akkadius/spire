<template>
  <eq-window
    title="Server Addresses"
    v-if="addresses && addresses.length > 0"
    class="p-0 pt-3 mb-4"
  >
      <table class="eq-table bordered eq-highlight-rows mb-0">
        <tbody>

        <tr v-for="i in addresses">
          <td class="text-right font-weight-bold">{{ i.key }}</td>
          <td class="text-left" style="min-width: 160px">
            <a
              href="javascript:"
              @click="copyToClip(i.value)"
              class="text-muted d-inline-block"
            >
              {{ i.value }}
              <i class="ml-2 fe fe-copy d-inline-block"></i>
            </a>
          </td>
        </tr>
        </tbody>
      </table>
  </eq-window>
</template>

<script>
import {SpireApi} from "@/app/api/spire-api";
import ClipBoard  from "@/app/clipboard/clipboard";
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {Notify}   from "@/app/Notify";

export default {
  name: 'DashboardNetworkingInfo',
  components: { EqWindow },
  data() {
    return {
      addresses: [],
    }
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  async mounted() {
    this.timer = null;
    let v      = []

    try {
      const r = await SpireApi.v1().get('admin/serverconfig')
      if (r.status === 200) {
        if (typeof r.data.server.world.address !== 'undefined') {
          v.push({ key: "World Public Address", value: r.data.server.world.address })
        }
        if (typeof r.data.server.world.localaddress !== 'undefined') {
          v.push({ key: "World LAN Address", value: r.data.server.world.localaddress })
        }
      }
    } catch (e) {
    }

    this.addresses = v
  },
  methods: {
    copyToClip(s) {
      ClipBoard.copyFromText(s)
      Notify.toast("Copied to clipboard!");
    },
  }
}
</script>

<style scoped>
.sysinfo td {
  padding: 0.5rem;
  vertical-align: top;
  border-top: 1px solid rgba(0, 40, 100, 0.12);
}
</style>
