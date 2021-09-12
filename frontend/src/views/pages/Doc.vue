<template>
  <div class="row justify-content-center" v-if="title">
    <div class="col-12 col-lg-10 col-xl-7">
      <!-- CONTENT -->
      <div class="container-fluid">


        <div class="header mt-md-0 mt-6">
          <div class="header-body">

            <!-- Title -->
            <h1 class="header-title mb-4 mt-4" style="font-size: 40px">
              {{ title }}
            </h1>

            <!-- Subtitle -->
            <p class="header-subtitle">
              {{ description }}
            </p>

          </div>
        </div>

        <v-runtime-template :template="doc" class="pb-6 mt-3 doc"/>

      </div>
    </div>
  </div>

</template>

<script>
import {SpireApiClient} from "@/app/api/spire-api-client";

export default {
  components: {
    "page-header": () => import("@/views/layout/PageHeader"),
    "v-runtime-template": () => import("v-runtime-template")
  },

  data() {
    return {
      doc: null,
      title: null,
      description: null,
      routeWatcher: null
    }
  },

  watch: {
    $route(to, from) {
    }
  },

  methods: {
    textBetween: function (text, begin, last) {
      return text.split(begin).pop().split(last)[0]
    },
    htmlDecode: function (encodedString) {
      var textArea       = document.createElement("textarea");
      textArea.innerHTML = encodedString;
      return textArea.value;
    },
    docInit() {
      const page = encodeURIComponent(this.$route.params.doc)

      SpireApiClient.v1().get(`/doc/${page}`).then((response) => {
        if (response.data && response.data.data) {

          const hljs = require("highlight.js");
          const md   = require("markdown-it")({
            html: true,
            xhtmlOut: false,
            breaks: true,
            typographer: false,
            highlight: function (str, lang) {
              if (lang && hljs.getLanguage(lang)) {
                try {
                  return "<div class='card'><div class='card-footer bg-dark'><pre class=\"highlight html bg-dark hljs mb-0 " + lang + "\">" +
                    hljs.highlight(lang, str, true).value +
                    "</pre></div></div>";
                } catch (__) {
                }
              }

              return "<pre class=\"hljs\"><code>" + md.utils.escapeHtml(str) + "</code></pre>";
            }
          });

          let result = response.data.data

          // All of the craziness below is to massage markdown data formatted from Gitbook
          // Its temporary / permanent throwaway logic

          this.description = ""


          // description
          let rs          = result.split("\n")
          let description = ""
          if (result.includes("description")) {
            let descriptionBlock = rs[0] + rs[1] + rs[2] + rs[3] + rs[4]
            description          = this.textBetween(descriptionBlock, "description", "---")
            description          = description.replaceAll(":", "")
            description          = description.replaceAll(":", "")
            description          = description.replaceAll(">-", "")
            description          = description.trim()

            this.description = description

            // cleanup body
            result = result.replaceAll("description: ", "")
            result = result.replaceAll(description, "")
            result = result.replace("---", "")
            result = result.replace("---", "")
            result = result.replaceAll(">-", "")
          }

          // console.log(result)

          result = md.render(result);
          result = result.replaceAll("<p>", "")
          result = result.replaceAll("<br>", "")
          result = result.replaceAll("</p>", "")
          result = result.replaceAll("<table>", "<table class=\"table table-sm table-nowrap mt-5\">")
          // result = result.replaceAll("<ul>", "<ul class=\"mt-5\">")
          result = result.replaceAll("\"../.gitbook/assets/", "\"https://ghcdn.rawgit.org/EQEmu/docs-quest-api/master/.gitbook/assets/")
          result = result.replaceAll("\".gitbook/assets/", "\"https://ghcdn.rawgit.org/EQEmu/docs-quest-api/master/.gitbook/assets/")
          // result = this.htmlDecode(result)

          // console.log(result)
          if (result.indexOf("% tabs ") !== -1) {
            result = result.replaceAll("{% tabs %}", "<b-tabs content-class=\"mt-2\">")
            result = result.replaceAll("{% endtabs %}", "</b-tabs>")
            result = result.replaceAll("{% tab title=\"", "<b-tab title=\"")
            result = result.replaceAll("{% tab title=&quot;", "<b-tab title=\"")
            result = result.replaceAll("{% tab title=\"", "<b-tab title=\"")
            result = result.replaceAll("\" %}", "\">")
            result = result.replaceAll("&quot; %}", "\">")
            result = result.replaceAll("{% endtab %}", "</b-tab>")
          }

          // title
          let title = result.split("\n")[0]

          // remote title from body
          result = result.replaceAll(title, "")

          this.title = title.replace("<h1>", "").replace("</h1>", "")

          // final trim
          result = result.trim()

          // doc
          this.doc = "<div>" + result + "</div>"

        }
      })
    }

  },

  async mounted() {
    // const page = 'methods%2Fraid-methods.md'

    // create route watcher
    this.routeWatcher = this.$watch("$route.params.doc", () => {
      this.doc = this.$route.query.q
      this.docInit();
    });

    this.docInit();
  }
}
</script>

<style>
.doc h2, .doc h1 {
  margin-top:    30px;
  margin-bottom: 20px;
}

.doc h4 {
  margin-top:    30px;
  margin-bottom: 20px;
}

.doc h3 {
  margin-top:    30px;
  margin-bottom: 20px;
}

.doc h5 {
  margin-top:    30px;
  margin-bottom: 20px;
}

.doc, .doc table {
  font-size: 16px;
}

.doc pre {
  font-size: 14px;
}

.doc h1 {
  font-size: 32px;
}

.doc h2 {
  font-size: 24px;
}

.doc h3 {
  font-size: 20px;
}

.doc h4 {
  font-size: 16px;
}

.doc h5 {
  font-size: 16px;
}

.doc h6 {
  font-size: 16px;
}

.doc img {
  padding-top: 50px;
  padding-bottom: 50px;

}

.doc pre .highlight {
  padding-top: 0px;
  padding-bottom: 10px;
  padding-left: 0px;
  padding-right: 0;
}



</style>
