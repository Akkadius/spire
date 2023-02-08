<template>
  <div class="card">

    <div class="card-header">
      <h4 class="card-header-title">Client Assets</h4>
    </div>

    <div class="card-body p-1">

      <div class="row">
        <div class="col-lg-12">

          <div class="ml-5 mt-3">
            <i class="fe fe-info"></i>
            All links below can be used publicly for download without login
          </div>


          <table class="table card-table table-vcenter mt-3">
            <tbody>
            <tr>
              <td class="text-right" style="width: 180px">
                <a
                  :href="buildLink('/download/spells')" target="spells_download"
                  class="btn btn-outline-primary mb-2"
                ><i
                  class="fe fe-download mr-2"
                ></i>spells_us.txt</a>
              </td>
              <td>Generates and zips a spells file for use with an EverQuest client</td>
            </tr>
            <tr>
              <td class="text-right">
                <a :href="buildLink('/download/dbstring')" target="dbstr_download" class="btn btn-outline-primary mb-2"><i
                  class="fe fe-download mr-2"
                ></i>dbstr_us.txt</a>
              </td>
              <td>Generates and zips a strings file for use with an EverQuest client
              </td>
            </tr>
            <tr>
              <td class="text-right">
                <a
                  :href="buildLink('/download/skills')" target="skills_download"
                  class="btn btn-outline-primary mb-2"
                ><i
                  class="fe fe-download mr-2"
                ></i>SkillCaps.txt</a>
              </td>
              <td>Generates and zips a skills file for use with an EverQuest client</td>
            </tr>
            <tr>
              <td class="text-right">
                <a :href="buildLink('/download/basedata')" target="base_data" class="btn btn-outline-primary mb-2"><i
                  class="fe fe-download mr-2"
                ></i>BaseData.txt</a>
              </td>
              <td>Generates and zips a base data file for use with an EverQuest client
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";
import util            from "util";

export default {
  data() {
    return {}
  },
  methods: {
    buildLink(link) {
      return util.format(
        "%s/api/v1/admin/occulus%s",
        OcculusClient.getBaseUrl(),
        link,
      )
    }
  },
}
</script>

<style scoped>
.td-center {
  text-align: center
}
</style>
