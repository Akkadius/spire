<template>
  <div>
    <div class="card" v-if="addresses && addresses.length > 0">
      <div class="card-header">
        <h4 class="card-header-title">
          Server Addresses
        </h4>
      </div>
      <div class="card-body" style="padding: 0px; overflow-y:scroll">
        <table class="table card-table sysinfo">
          <tbody>

          <tr v-for="i in addresses">
            <td style="vertical-align: middle">{{ i.key }}</td>
            <td class="text-right" style="min-width: 150px">
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
      </div>
    </div>

  </div>
</template>

<script>
import {SpireApi} from "@/app/api/spire-api";
import ClipBoard  from "@/app/clipboard/clipboard";

export default {
  name: 'DashboardNetworkingInfo',
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

      this.$bvToast.toast("Copied to clipboard!", {
        title: "Copy",
        autoHideDelay: 2000,
        solid: true
      })
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
