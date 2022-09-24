<template>
  <div>
    <div class="row">
      <div :class="!isBotSelected() ? 'col-12' : 'col-6'">



        <eq-window style="height: 95vh" class="mt-0">

          <div class="row">
            <div class="col-3">
              <b-input
                class="form-control"
                @keyup="updateQueryState"
                v-model="search"
              />
            </div>
          </div>

          <table class="eq-table eq-highlight-rows">
            <thead>
            <tr>
              <th>name</th>
              <th>race</th>
            </tr>
            </thead>
            <tbody>
            <tr
              @click="selectBot(row)"
              v-for="row in botlist"
            >
              <td>{{ row.name }}</td>
              <td>{{ row.race }}</td>
            </tr>
            </tbody>
          </table>
        </eq-window>
      </div>
      <div :class="isBotSelected() ? 'col-6' : ''" v-if="isBotSelected()">
        <eq-window title="Preview Pane">
          {{ selectedBot }}
        </eq-window>
      </div>

    </div>
  </div>
</template>

<script>
import ContentArea         from "../components/layout/ContentArea";
import EqWindow            from "../components/eq-ui/EQWindow";
import LoaderFakeProgress  from "../components/LoaderFakeProgress";
import {SpireApiClient}    from "../app/api/spire-api-client";
import EqDebug             from "../components/eq-ui/EQDebug";
import {BotDatumApi}       from "../app/api";
import {SpireQueryBuilder} from "../app/api/spire-query-builder";
import {ROUTE}             from "../routes";

export default {
  name: "TestBot",
  components: { EqDebug, LoaderFakeProgress, EqWindow, ContentArea },
  data() {
    return {
      selectedBot: {},
      botlist: [],

      search: "",
    }
  },
  async mounted() {
    this.init();
  },

  watch: {
    $route(to, from) {
      console.log("TRIGGER WATCH FOR ROUTE NAV")

      this.init()
    }
  },


  methods: {

    async init() {
      this.loadQueryState()

      let builder = (new SpireQueryBuilder())

      if (this.search.length > 0) {
        builder.where("name", "like", this.search);
      }

      const api = (new BotDatumApi(SpireApiClient.getOpenApiConfig()))
      const r   = await api.listBotData(builder.get());
      if (r.status === 200) {
        console.log("hello")
        this.botlist = r.data
      }

      // const r = await SpireApiClient.v1().get('bot_data')
      // if (r.status === 200) {
      //   this.botlist = r.data
      // }

      console.log(r)

    },

    // when inputs are triggered and state is updated
    updateQueryState: function () {
      let q = {};

      if (this.search !== "") {
        q.q = this.search
      }

      this.$router.push(
        {
          path: '/testbot',
          query: q
        }
      ).catch(() => {
      })
    },

    // usually from loading initial state
    loadQueryState: function () {
      if (this.$route.query.q) {
        this.search = this.$route.query.q;
      }
    },

    isBotSelected() {
      return Object.keys(this.selectedBot).length > 0
    },
    selectBot(bot) {
      // console.log(bot)
      console.log("selected", bot.name)
      this.selectedBot = bot;
    },
    async searchBot() {

      let builder = (new SpireQueryBuilder())
      builder.where("name", "like", this.search);

      const api = (new BotDatumApi(SpireApiClient.getOpenApiConfig()))
      const r   = await api.listBotData(builder.get());
      if (r.status === 200) {
        this.botlist = r.data
      }
    }
  }
}
</script>

<style scoped>

</style>
