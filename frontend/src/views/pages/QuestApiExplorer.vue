<template>
  <!-- CONTENT -->
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-12">
            <eq-window title="Quest API Explorer" style="margin-top: 30px">

              <app-loader :is-loading="!loaded" padding="6"/>

              <div v-if="loaded">

                <div class="row justify-content-center">
                  <div class="col-2 text-center">
                    Language
                    <b-form-select
                      v-model="languageSelection"
                      :options="languageOptions"
                      @change="formChange(); languageSelect()"/>
                  </div>
                  <div class="col-2 text-center">
                    Types
                    <b-form-select
                      v-model="methodTypeSelection"
                      @change="formChange(); methodTypeSelect()"
                      :options="methodTypeOptions"/>
                  </div>

                </div>
                <div class="row mt-5 justify-content-center" v-if="apiMethods.length > 0">
                  <div class="col-10">
                    <h4 class="eq-header text-center" v-if="methodTypeSelection">Type: {{ methodTypeSelection }}</h4>
                    <pre
                      class="highlight html bg-dark hljs mb-4 code-display"
                      style="padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                      v-if="apiMethods.length > 0"><div
                      v-for="(method, index) in apiMethods"
                      :key="index">{{ method.methodPrefix }}{{
                        method.method
                      }}({{ method.params.join(", ") }}); {{ method.comments }}</div></pre>
                  </div>
                </div>
              </div>


            </eq-window>
          </div>
        </div>
      </div>
    </div>

    <!--        <debug-display-component :data="database"/>-->

  </div>

</template>

<script type="ts">
import EqWindow              from "@/components/eq-ui/EQWindow.vue";
import DebugDisplayComponent from "@/components/DebugDisplayComponent.vue";
import {SpireApiClient}      from "@/app/api/spire-api-client";
import axios                 from "axios";
import EqWindowSimple        from "@/components/eq-ui/EQWindowSimple.vue";


export default {
  components: {
    EqWindowSimple,
    DebugDisplayComponent,
    EqWindow,
    "page-header": () => import("@/views/layout/PageHeader")
  },
  data() {
    return {
      codeClass: "",

      // select
      languageSelection: null,
      languageOptions: [
        {value: null, text: 'Select a Language'},
        {value: 'perl', text: 'Perl'},
        {value: 'lua', text: 'Lua'},
      ],

      // select
      methodTypeSelection: null,
      methodTypeOptions: [],

      apiMethods: [],

      // actual code that gets highlighted
      methodDisplay: "",

      // route watcher
      routeWatcher: null,

      loaded: false,
    }
  },
  deactivated() {
    // remove route watcher
    this.routeWatcher()
  },
  activated() {
    this.init()
  },
  methods: {
    init() {
      this.routeWatcher = this.$watch('$route.query', () => {
        this.loadQueryParams()
      });

      this.loaded = false

      SpireApiClient.v1().get('/quest-api/methods', this.database).then((response) => {
        if (response.data && response.data.data) {
          this.methods = response.data.data

          this.languageSelect()
          this.methodTypeSelect()
          this.loaded = true
        }

      }, (error) => {
        // this.errorMessage = "Unknown error trying to contact the database"
        if (error.response && error.response.data) {
          // this.errorMessage = error.response.data.error
        }
      }).catch((error) => {
        if (!axios.isCancel(error)) {
          // console.log(error)
          // this.errorMessage = "Unknown error trying to contact the database"
        }
      });

      // reset
      this.languageSelection   = null
      this.methodTypeSelection = null

      this.loadQueryParams()
    },
    loadQueryParams() {
      if (this.$route.query.lang) {
        this.languageSelection = this.$route.query.lang
      }
      if (this.$route.query.type) {
        this.methodTypeSelection = this.$route.query.type
      }
    },
    formChange() {
      let query = {}
      if (this.languageSelection) {
        query.lang = this.languageSelection
      }
      if (this.methodTypeSelection) {
        query.type = this.methodTypeSelection
      }

      this.$router.push(
        {
          path: '/quest-api-explorer',
          query: query
        }
      ).catch(() => {
      })
    },
    getLanguageKey() {
      if (this.languageSelection == "perl") {
        return "perl_api"
      }
      if (this.languageSelection == "lua") {
        return "lua_api"
      }
      return ""
    },
    languageSelect: function () {
      if (this.methods[this.getLanguageKey()]) {
        let options = []
        options.push({value: null, text: "Select a type"})
        Object.keys(this.methods[this.getLanguageKey()]).sort().filter((item) => {
          return !item.includes("Deprecated")
        }).forEach((option) => {
          options.push({value: option, text: option})
        })

        this.methodTypeOptions = options

        if (!this.methodTypeSelection) {
          this.methodTypeSelection = this.methodTypeOptions[0]
        }
      }
      this.methodDisplay = ""
      this.methodTypeSelect()
    },
    methodTypeSelect: function () {

      this.apiMethods = []
      let apiMethods  = []

      if (this.methods[this.getLanguageKey()]) {
        let methodDisplay = ""
        this.codeClass    = this.getLanguageKey()
        const methods     = this.methods[this.getLanguageKey()][this.methodTypeSelection]
        const snakeCase   = string => {
          return string.replace(/\W+/g, " ")
                       .split(/ |\B(?=[A-Z])/)
                       .map(word => word.toLowerCase())
                       .join('_');
        };

        if (methods && methods.length > 0) {
          methods.forEach((method) => {
            let prefix = method.method_type
            prefix     = prefix.replace("NPC", "Npc")
            let comment;

            // perl
            if (this.languageSelection === "perl") {
              if (prefix == "quest") {
                prefix = prefix.replace(prefix, "quest::")
              } else {
                prefix = prefix.replace(prefix, "$" + snakeCase(prefix).toLowerCase() + "->")
              }
            }

            // lua
            if (this.languageSelection === "lua") {
              if (prefix == "eq") {
                prefix = prefix.replace(prefix, "eq.")
              } else {
                prefix = prefix.replace(prefix, snakeCase(prefix).toLowerCase() + ":")
              }

              if (method.return_type !== "") {
                comment = "-- return " + method.return_type
              }
            }

            apiMethods.push({
              method: method.method,
              methodPrefix: prefix,
              params: method.params,
              comments: comment,
            })

          })
        }

        this.apiMethods    = apiMethods
        this.methodDisplay = methodDisplay
      }
    }
  }
}

</script>

<style>
.code-display {
  font-size:  14px !important;
  max-width:  100% !important;
  width:      100%;
  min-height: 300px;
}
</style>
