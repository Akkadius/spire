<template>
  <div>

    <eq-window title="EverQuest Client Asset Exports (Publicly Downloadable Links)">
      <div class="eq-alert">
        <i class="fe fe-info"></i>
        All links below can be used publicly for download without login. If you link from an SSL site, the file downloads will be blocked by browser - so you will need to proxy the request
      </div>

      <table
        class="eq-table eq-highlight-rows bordered log-settings minified-inputs mt-3"
      >
        <thead class="eq-table-floating-header">
        <tr>
          <th class="text-center" style="width: 120px">Asset</th>
          <th class="">Description</th>
        </tr>

        </thead>
        <tbody>
        <tr
          v-for="r in assetDownloads"
          :key="r.path"
        >
          <td class="text-right pr-3 pt-3 pb-3">
            <a class="eq-button d-block" :href="download(r)" :target="r.path">{{ r.asset }}</a>
          </td>
          <td>
            <span class="font-weight-bold">{{ r.description }}</span>
          </td>
        </tr>
        </tbody>
      </table>
    </eq-window>

  </div>
</template>

<script>
import util       from "util";
import {SpireApi} from "@/app/api/spire-api";
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
//
// t := []ExportType{
//   {arg: "spells", file: "spells_us.txt"},
//   {arg: "skills", file: "SkillCaps.txt"},
//   {arg: "basedata", file: "BaseData.txt"},
//   {arg: "dbstring", file: "dbstr_us.txt"},


export default {
  components: { EqWindow },
  data() {
    return {
      assetDownloads: [
        {asset: "spells_us.txt", path: "spells", description: "Generates a Spells file for use in the EverQuest Client"},
        {asset: "SkillCaps.txt", path: "skills", description: "Generates a Skills file for use in the EverQuest Client"},
        {asset: "BaseData.txt", path: "basedata", description: "Generates a Base Data file for use in the EverQuest Client"},
        {asset: "dbstr_us.txt", path: "dbstring", description: "Generates a strings database file for use in the EverQuest Client"},
      ]
    }
  },
  methods: {
    download(a) {
      return util.format(
        "%s/eqemuserver/export-client-file/%s",
        SpireApi.getBaseV1Path(),
        a.path,
      )
    }
  },
}
</script>
