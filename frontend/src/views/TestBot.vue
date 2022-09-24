<template>
  <content-area>
    <eq-window>

      <!--      <eq-debug :data="botlist"/>-->

      <table class="eq-table eq-highlight-rows">
        <thead>
          <tr>
            <th>name</th>
            <th>race</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in botlist">
            <td>{{row.name}}</td>
            <td>{{row.race}}</td>
          </tr>
        </tbody>
      </table>


    </eq-window>
  </content-area>
</template>

<script>
import ContentArea        from "../components/layout/ContentArea";
import EqWindow           from "../components/eq-ui/EQWindow";
import LoaderFakeProgress from "../components/LoaderFakeProgress";
import {SpireApiClient}   from "../app/api/spire-api-client";
import EqDebug            from "../components/eq-ui/EQDebug";

export default {
  name: "TestBot",
  components: { EqDebug, LoaderFakeProgress, EqWindow, ContentArea },
  data() {
    return {
      botlist: [],
    }
  },
  async mounted() {
    const r = await SpireApiClient.v1().get('bot_data')
    if (r.status === 200) {
      this.botlist = r.data
    }

    console.log(r)

  }
}
</script>

<style scoped>

</style>
