<template>
  <content-area>
    <div class="row">
      <div :class="(isSubEditActive() ? 'col-6' : 'col-12')">

        <eq-window-simple title="Strings Database">
          <div class="row">
            <div :class="(selectedType >= 0 ? 'col-10' : 'col-12') + ' text-center'">
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

            <div class="col-2 " v-if="selectedType >= 0">
              <b-button
                @click="createString()"
                class="mt-5"
                size="sm"
                variant="outline-warning"
              >
                <i class="fa fa-plus"></i>
                Create
              </b-button>
            </div>

          </div>

          <div class="row">
            <div
              class="col-12 text-center font-weight-bold mt-3"
              v-if="strings && strings.length > 0 && !loading && !isSubEditActive()"
            >
              Select a row to edit
            </div>

          </div>

          <loader-fake-progress v-if="loading"/>

        </eq-window-simple>

        <eq-window-simple
          style="height: 80vh; overflow-y: scroll; overflow-x: hidden"
          class="mt-3"
          id="db-strings-list"
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
                :key="row.id + '-' + row.type + '-' + index"
                style="border-radius: 10px"
                :class="isStringSelected(row) ? 'pulsate-highlight-white' : ''"
                @click="selectString(row.id, row.type)"
                :id="'string-' + row.id"
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
              v-model.number="selectedStringObject.id"
              @keydown="updateSelectedString('id')"
              id="selected_id"
            />
          </div>

          <div class="mt-3">
            Value
            <b-form-textarea
              v-model="selectedStringObject.value"
              placeholder="Enter something..."
              style="height: 300px"
              id="selected_value"
              @keydown="updateSelectedString('value')"
            />
          </div>

          <b-button
            @click="saveSelectedString()"
            size="sm"
            class="mt-3"
            variant="outline-warning"
          >
            <i class="fa fa-save"></i>
            Save
          </b-button>

          <b-button
            @click="deleteSelectedString()"
            size="sm"
            class="mt-3 ml-3"
            variant="outline-danger"
          >
            <i class="fa fa-trash"></i>
            Delete
          </b-button>

          <b-alert show variant="warning" v-if="message" class="mt-4 mb-0 fade-in">{{ message }}</b-alert>

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
import EqWindowSimple      from "../../components/eq-ui/EQWindowSimple";
import EqAutoTable         from "../../components/eq-ui/EQAutoTable";
import ContentArea         from "../../components/layout/ContentArea";
import {DbStrApi}          from "../../app/api";
import {SpireApiClient}    from "../../app/api/spire-api-client";
import LoaderFakeProgress  from "../../components/LoaderFakeProgress";
import {ROUTE}             from "../../routes";
import {DB_STR_TYPES}      from "../../app/constants/eq-db-str-constants";
import {EditFormFieldUtil} from "../../app/forms/edit-form-field-util";
import util                from "util";

// api response cache of all strings
// this does not need to be reactive so don't put in data()
let allStrings       = []
const DbStrApiClient = (new DbStrApi(SpireApiClient.getOpenApiConfig()))

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

      message: "",

      originalSelectedStringObject: {},
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
      this.selectedType                 = -1
      this.subSelectedId                = -1
      this.subSelectedType              = -1
      this.originalSelectedStringObject = {}
      this.selectedStringObject         = {}
    },
    resetSelections() {
      this.subSelectedId   = -1;
      this.subSelectedType = -1;
      this.message         = ""
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

    async createString() {
      console.log("create")

      // filter list by type
      let r = allStrings.filter((s) => s.type === parseInt(this.subSelectedType))
        .sort((a, b) => (a.id > b.id) ? 1 : -1)

      // grab last id + 1 from list
      const newId = r[r.length - 1].id + 1

      // create
      try {
        const response = await DbStrApiClient.createDbStr(
          {
            dbStr: {
              id: newId,
              type: parseInt(this.subSelectedType),
              value: ""
            }
          }
        )

        // success
        if (response.status === 200 && response.data) {
          this.resetSelections()
          this.updateQueryState()
          this.selectString(newId, this.subSelectedType)
          this.init(true)
        }
      } catch (err) {
        if (err.response !== 200 && err.response.data.error) {
          this.message = err.response.data.error
        }
      }
    },

    async deleteSelectedString() {
      if (confirm("Are you sure you want to delete this string?")) {
        try {
          const response = await DbStrApiClient.deleteDbStr(
            {
              id: parseInt(this.subSelectedId)
            },
            {
              params: {
                where: "type__" + this.selectedType
              }
            }
          )

          // success
          if (response.status === 200 && response.data) {

            // get last element in current list and select it after deletion
            let r = allStrings.filter((s) => s.type === parseInt(this.subSelectedType))
              .sort((a, b) => (a.id > b.id) ? 1 : -1)

            let lastElement = {}

            // grab last element and select
            if (r && r.length > 0) {
              lastElement = r[r.length - 1]
              // if we deleted the last element, let's fallback to next in line...
              if (parseInt(lastElement.id) === parseInt(this.subSelectedId)) {
                if (r[r.length - 2]) {
                  lastElement = r[r.length - 2]
                }
              }
            }

            this.resetSelections()
            this.updateQueryState()

            if (lastElement) {
              this.selectString(lastElement.id, this.subSelectedType)
            }

            this.init(true)
          }
        } catch (err) {
          if (err.response !== 200 && err.response.data.error) {
            this.message = err.response.data.error
          }
        }
      }
    },

    async updateSelectedString(field) {
      EditFormFieldUtil.setFieldModifiedById("selected_" + field)
    },
    async saveSelectedString() {

      try {
        const response = await DbStrApiClient.updateDbStr(
          {
            id: parseInt(this.originalSelectedStringObject.id),
            dbStr: this.selectedStringObject
          },
          {
            params: {
              where: "type__" + this.selectedType
            }
          }
        )

        // success
        if (response.status === 200 && response.data) {
          EditFormFieldUtil.resetFieldEditedStatus()

          this.updateQueryState()
          this.init(true)
          this.message = "Saved successfully!"
          setTimeout(() => {
            this.message = ""
          }, 6000)
        }

      } catch (err) {
        if (err.response !== 200 && err.response.data.error) {
          this.message = err.response.data.error
        }
      }

    },

    /**
     * Sub editor selection
     */
    isSubEditActive() {
      return this.subSelectedId >= 0 && this.subSelectedType >= 0
    },
    getSelectedStringObject() {
      let r = allStrings.find((s) => s.type === this.subSelectedType && s.id === this.subSelectedId)

      return typeof r === 'undefined' ? {} : r
    },

    selectString(stringId, typeId) {
      this.subSelectedId                = stringId
      this.subSelectedType              = typeId
      this.originalSelectedStringObject = JSON.parse(JSON.stringify(this.getSelectedStringObject()))
      this.selectedStringObject         = this.getSelectedStringObject()
      this.updateQueryState()
    },

    isStringSelected(string) {
      return string.id === this.subSelectedId && string.type === this.subSelectedType
    },

    /**
     * Initialize
     */
    async init(reset = false) {
      this.reset()
      this.loadQueryState()
      if (allStrings && allStrings.length === 0 || reset) {
        allStrings = await this.getAllDbStrings()
      }
      this.calculateStringTypeCounts(allStrings)
      this.originalSelectedStringObject = JSON.parse(JSON.stringify(this.getSelectedStringObject()))
      this.selectedStringObject         = this.getSelectedStringObject()
      this.listData()
    },

    async getAllDbStrings() {
      // {orderBy: "id.type"}
      const response = await DbStrApiClient.listDbStrs()
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

    requestAnimationFrame(() => {
      setTimeout(() => {
        const container = document.getElementById("db-strings-list");
        const target    = document.getElementById(util.format("string-%s", this.subSelectedId))

        if (container && target) {
          const top           = target.getBoundingClientRect().top
          container.scrollTop = top - 200;
        }
      }, 500)
    })
  }
}
</script>

<style scoped>

</style>
