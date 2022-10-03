<template>
  <content-area style="padding: 0 !important">
    <div class="row">
      <div class="col-7">
        <eq-window title="Loot">
          <div class="row">
            <div class="col-6">
              <input
                type="text"
                class="form-control ml-2"
                placeholder="Search for item names"
                v-model="search"
                v-on:keyup="doSearch"
              >
            </div>
            <div class="col-6">
              <b-button title="Reset" @click="reset()" size="sm" variant="outline-warning btn-dark">
                <i class="fa fa-refresh mr-1"/>
                Reset
              </b-button>
            </div>
          </div>
        </eq-window>

        <eq-window class="p-0">
          <div style="overflow-y: scroll; height: 80vh">
            <table
              id="loot-table"
              class="eq-table eq-highlight-rows bordered"
              style="font-size: 14px; "
              v-if="tableData && tableData.length > 0"
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th
                  v-for="(header, index) in Object.keys(tableData[0])"
                  class="text-center"
                  v-if="!doesColumnHaveObjects(tableData, header)"
                  :id="'column-' + header"
                >{{ header }}
                </th>
              </tr>
              </thead>
              <tbody>
              <tr
                v-for="(row, index) in tableData"
                :id="'npc-' + row.id"
                :key="index"
                @mouseover="showTable(row)"
              >
                <td
                  :style="' text-align: center;'"
                  v-for="(key, colIndex) in Object.keys(row)"
                  v-if="doesRowColumnHaveObjects(row, key)"
                >
                  <span>{{ row[key] }}</span>
                  <!--                  <loot-popover-->
                  <!--                    v-if="key === 'name'"-->
                  <!--                    :loot="row"-->
                  <!--                  />-->
                </td>
              </tr>
              </tbody>
            </table>
          </div>

          <div class="row text-center justify-content-center">
            <div class="col-12 text-center mt-3">
              <b-pagination
                class="mb-3"
                v-model="currentPage"
                :total-rows="totalRows"
                :hide-ellipsis="true"
                :per-page="100"
                @change="paginate"
              />
            </div>
          </div>

        </eq-window>
      </div>
      <div class="col-5">

        <eq-window
          v-if="Object.keys(previewedTable).length === 0"
        >
          Select a table on the left to preview
        </eq-window>

        <eq-loot-card-preview
          v-if="Object.keys(previewedTable).length > 0"
          :loot="previewedTable"
        />
      </div>
    </div>

  </content-area>
</template>

<script>
import EqWindow            from "../../components/eq-ui/EQWindow";
import ContentArea         from "../../components/layout/ContentArea";
import LootPopover         from "../../components/LootPopover";
import EqLootCardPreview   from "../../components/preview/EQLootCardPreview";
import {ROUTE}             from "../../routes";
import {LoottableApi}      from "../../app/api";
import {SpireApiClient}    from "../../app/api/spire-api-client";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import {debounce}          from "../../app/utility/debounce";

export default {
  name: "Loot",
  components: { EqLootCardPreview, LootPopover, ContentArea, EqWindow },
  data() {
    return {
      tableData: {},

      previewedTable: {},

      search: "",

      // pagination (all)
      currentPage: 1,
      totalRows: 0,
    }
  },
  async mounted() {
    this.init()
  },

  watch: {
    '$route'() {
      this.init()
    },
  },

  methods: {

    reset() {
      this.currentPage = 1
      this.search      = ""
      this.totalRows   = 0
      this.updateQueryState()
    },

    doSearch() {
      this.currentPage = 1
      this.updateQueryState()
    },

    async listLootTables() {
      let builder = (new SpireQueryBuilder())
        .page(this.currentPage)
        .includes([
          "NpcTypes",
          "LoottableEntries",
          "LoottableEntries.Lootdrop",
          "LoottableEntries.Lootdrop.LootdropEntries",
          "LoottableEntries.Lootdrop.LootdropEntries.Item",
        ])
        .limit(100)

      if (this.search.length > 0 && Number.isInteger(this.search)) {
        builder.where("id", "=", this.search);
      }
      if (this.search.length > 0 && !Number.isInteger(this.search)) {
        builder.where("name", "like", this.search);
      }

      const r = await (new LoottableApi(SpireApiClient.getOpenApiConfig()))
        .listLoottables(
          // @ts-ignore
          builder.get()
        )
      if (r.status === 200) {
        return r.data
      }
    },

    async getTotalLootTables() {
      let builder = (new SpireQueryBuilder())
        .select(["id"])
        .limit(100000000)

      if (this.search.length > 0 && Number.isInteger(this.search)) {
        builder.where("id", "=", this.search);
      }
      if (this.search.length > 0 && !Number.isInteger(this.search)) {
        builder.where("name", "like", this.search);
      }

      const r = await (new LoottableApi(SpireApiClient.getOpenApiConfig()))
        .listLoottables(
          // @ts-ignore
          builder.get()
        );

      return r.data.length
    },

    async init() {
      this.loadQueryState()
      this.tableData = await this.listLootTables()
      this.totalRows = await this.getTotalLootTables()

      console.log(this.tableData)

      console.log("init")
    },

    updateQueryState: debounce(function () {
      // query params
      let queryState = {};
      if (this.currentPage > 0) {
        queryState.page = this.currentPage
      }
      if (this.search !== "") {
        queryState.q = this.search
      }

      // navigation
      this.$router.push(
        {
          path: ROUTE.LOOT,
          query: queryState
        }
      ).catch(() => {
      })
    }, 600),

    loadQueryState() {
      if (typeof this.$route.query.page !== 'undefined' && parseInt(this.$route.query.page) !== 0) {
        this.currentPage = parseInt(this.$route.query.page);
        console.log("current page", this.currentPage)
      }
      if (typeof this.$route.query.q !== 'undefined' && this.$route.query.q !== "") {
        this.search = this.$route.query.q;
        console.log("search is ", this.search)
      }
    },

    paginate() {

      // models aren't quite updated when we trigger this so queue the pagination
      setTimeout(() => {
        console.log("We're paginating")
        console.log(this.currentPage)
        console.log(this.totalRows)
        this.updateQueryState()
      }, 100)

    },
    showTable(loottable) {
      this.previewedTable = loottable
    },
    doesColumnHaveObjects(data, column) {
      if (typeof column === 'object') {
        return true
      }

      return data.find((row) => {
        return typeof row[column] === 'object' && row[column] !== null && Object.keys(row[column])
      })
    },
    doesRowColumnHaveObjects(r, key) {
      return (typeof r[key] !== 'undefined') && !(typeof r[key] === 'object' && r[key] !== null && Object.keys(r[key]))
    },
  }
}
</script>

<style scoped>

</style>
