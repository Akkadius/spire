<template>
  <eq-window title="Release Crash Reporting" class="p-3">
    <div style="max-height:95vh; overflow-y: scroll">
      <table
        class="eq-table bordered eq-highlight-rows"
      >
        <thead class="eq-table-floating-header">
        <tr>
          <th style="width: 50px;"></th>
          <th style="width: 140px;">Release</th>
          <th style="width: 140px;">Compile Date</th>
          <th style="width: auto;">Crash Count</th>
        </tr>
        </thead>
        <tbody>
        <tr
          class="fade-in"
          v-for="(c, index) in counts"
          :key="c.server_version"
        >
          <td class="text-center">
            <b-button
              variant="primary"
              class="btn-dark btn-sm btn-outline-white"
              style="padding: 0px 6px;"
              title="View Release"
              @click="goToRelease(c)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>
          <td>{{ c.server_version }}</td>
          <td>{{ c.compile_date }}</td>
          <td>{{ c.crash_count }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </eq-window>
</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "@/app/api/spire-api";
import util       from "util";
import {ROUTE}    from "@/routes";

export default {
  name: "Releases",
  components: { EqWindow },
  data() {
    return {
      counts: []
    }
  },
  methods: {
    goToRelease(r) {
      this.$router.push(
        {
          path: util.format(ROUTE.RELEASE, r.server_version),
        }
      ).catch(() => {
      })
    }
  },
  async mounted() {
    const r = await SpireApi.v1().get(`server-crash-report/counts`)
    if (r.status === 200) {
      this.counts = r.data
    }
  }
}
</script>

<style scoped>

</style>
