<template>
  <content-area>
    <div class="row">
      <div :class="(isSubEditActive() ? 'col-6' : 'col-12')">

        <eq-window-simple title="Strings Database">
          <div class="row">
            <div class="col-12 text-center">
              Type
              <b-form-select
                v-model.number="selectedType"
                @change="resetSelections(); updateQueryState()"
                class="mt-3 form-control"
              >
                <option value="-1">--- Select ---</option>
                <option
                  v-for="(description, index) in DB_STR_TYPES"
                  :key="index"
                  :value="parseInt(index)"
                >
                  {{ index }}) {{ description }} ({{ typeCounts[index] ? commify(typeCounts[index]) : 0 }})
                </option>
              </b-form-select>
            </div>

            <!--            <span-->
            <!--              class="font-weight-bold mt-5 ml-3"-->
            <!--              v-if="strings && strings.length > 0 && !loading && !isSubEditActive()"-->
            <!--            >-->
            <!--              Select a row to edit-->
            <!--            </span>-->
          </div>

          <loader-fake-progress v-if="loading"/>

        </eq-window-simple>

        <eq-window-simple
          style="height: 80vh; overflow-y: scroll; overflow-x: hidden"
          class="mt-3"
          v-if="strings && strings.length > 0 && !loading"
        >

          <div class='eq-window-nested-blue' style="width: 100%;">
            <table
              class="eq-table eq-highlight-rows"
              style="display: table; overflow-x: scroll " v-if="strings && strings.length > 0"
            >
              <thead>
              <tr>
                <th style="width: 100px">Id</th>
                <th>Value</th>
              </tr>
              </thead>
              <tbody>
              <tr
                v-for="(row, index) in strings"
                :key="row.id + row.type"
                style="border-radius: 10px"
                :class="isStringSelected(row) ? 'pulsate-highlight-white' : ''"
                @click="selectString(row.id, row.type)"
              >
                <td>{{ row.id }}</td>
                <td>{{ row.value }}</td>
              </tr>
              </tbody>
            </table>
          </div>

        </eq-window-simple>

      </div>

      <div class="col-6 fade-in" v-if="isSubEditActive()">
        <eq-window-simple title="Edit Database String">

          <div class="mt-3">
            ID
            <b-input
              v-model="selectedStringObject.id"
            />
          </div>

          <div class="mt-3">
            Value
            <b-form-textarea
              v-model="selectedStringObject.value"
              placeholder="Enter something..."
              style="height: 300px"
            />
          </div>

        </eq-window-simple>

        <eq-window-simple :title="'String Preview Type (' + selectedStringObject.type + ') ID (' + selectedStringObject.id + ')'">
          <v-runtime-template
            v-if="getSelectedStringObject()"
            :template="'<div>' + formatStringPreview(selectedStringObject.value) + '</div>'"
          />
        </eq-window-simple>
      </div>


    </div>

  </content-area>
</template>

<script>
import EqWindowSimple     from "../../components/eq-ui/EQWindowSimple";
import EqAutoTable        from "../../components/eq-ui/EQAutoTable";
import ContentArea        from "../../components/layout/ContentArea";
import {DbStrApi}         from "../../app/api";
import {SpireApiClient}   from "../../app/api/spire-api-client";
import LoaderFakeProgress from "../../components/LoaderFakeProgress";
import {ROUTE}            from "../../routes";
import {DB_STR_TYPES}     from "../../app/constants/eq-db-str-constants";

// api response cache of all strings
// this does not need to be reactive so don't put in data()
let allStrings = []

export default {
  name: "StringsDatabase",
  components: {
    LoaderFakeProgress: LoaderFakeProgress,
    ContentArea,
    EqAutoTable,
    EqWindowSimple,
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {
      strings: [], // strings to be viewed
      typeCounts: {}, // stores the counts per type (select)

      selectedType: -1, // selected state

      // for the sub selector pane on the right
      subSelectedId: -1,
      subSelectedType: -1,

      selectedStringObject: {},

      loading: false, // are we loading or not

      DB_STR_TYPES: DB_STR_TYPES
    }
  },

  watch: {
    '$route'() {
      console.log("route trigger")

      this.init()
    },
  },

  methods: {


    /**
     * Resets
     */
    reset() {
      this.selectedType         = -1
      this.subSelectedId        = -1
      this.subSelectedType      = -1
      this.selectedStringObject = {}
    },
    resetSelections() {
      this.subSelectedId   = -1;
      this.subSelectedType = -1;
    },

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.selectedType !== -1) {
        queryState.type = this.selectedType
      }
      if (this.subSelectedId !== -1) {
        queryState.selectedId = this.subSelectedId
      }

      this.$router.push(
        {
          path: ROUTE.STRINGS_DATABASE,
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      console.log("loading query state")
      if (this.$route.query.type >= 0) {
        this.selectedType    = parseInt(this.$route.query.type);
        this.subSelectedType = parseInt(this.$route.query.type);
      }
      if (this.$route.query.selectedId >= 0) {
        this.subSelectedId = parseInt(this.$route.query.selectedId);
        console.log("selected object", this.selectedStringObject)

      }

      console.log("selected type", this.selectedType)
      console.log("sub selected type", this.subSelectedType)
      console.log("sub selected id", this.subSelectedId)

    },

    /**
     * Helpers
     */
    listData() {
      this.loading = true
      let strings  = []
      allStrings.forEach((string) => {
        if (string.type === parseInt(this.selectedType)) {
          strings.push(string)
        }
      });

      this.strings = strings
      this.loading = false
    },
    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },
    replaceAll(str, find, replace) {
      return str.replace(new RegExp(find, 'g'), replace);
    },
    formatStringPreview(contents) {
      if (contents) {
        return this.replaceAll(contents, "<BR>", "<BR/>")
      }
      return ""
    },

    /**
     * Sub editor selection
     */
    isSubEditActive() {
      return this.subSelectedId >= 0 && this.subSelectedType >= 0
    },
    getSelectedStringObject() {
      return allStrings.find((s) => s.type === this.subSelectedType && s.id === this.subSelectedId)
    },

    selectString(stringId, typeId) {
      this.subSelectedId        = stringId
      this.subSelectedType      = typeId
      this.selectedStringObject = this.getSelectedStringObject()
      this.updateQueryState()
    },

    isStringSelected(string) {
      return string.id === this.subSelectedId && string.type === this.subSelectedType
    },

    /**
     * Initialize
     */
    async init() {
      this.reset()
      this.loadQueryState()
      if (allStrings && allStrings.length === 0) {
        allStrings = await this.getAllDbStrings()
      }
      this.calculateStringTypeCounts(allStrings)
      this.selectedStringObject = this.getSelectedStringObject()
      this.listData()
    },

    async getAllDbStrings() {
      const api      = (new DbStrApi(SpireApiClient.getOpenApiConfig()))
      const response = await api.listDbStrs()
      if (response.status === 200 && response.data) {
        return response.data
      }
    },

    calculateStringTypeCounts(allStrings) {
      let typeStringCount = {}
      allStrings.forEach((string) => {
        if (typeof typeStringCount[string.type] === "undefined") {
          typeStringCount[string.type] = 0
        }

        typeStringCount[string.type]++
      })

      this.typeCounts = typeStringCount
    }

  },
  mounted() {
    this.init()
  }
}
</script>

<style scoped>

</style>
