<template>
  <div>
    <div class="row">
      <div class="col-4">
        <eq-window>
          <h1 class="eq-header">Websocket Proof of Concept</h1>

          <div class="eq-alert">
            We're gonna play with some websockets. We're going to see if we can real-time stream command output from the
            backend to the frontend<br><br>

            The purpose of this will be to run pre-flight checks on the server to ensure that there are no errors before
            doing a full launch.<br><br>

            For example, do we need database updates? Did we miss an update? Was there an error on initial launch?
            <br><br>
            We also will need websockets for real-time notifications between the server and Spire
          </div>

          <div class="mt-3">
            <button class="btn btn-dark btn-sm" @click="runProcess('world')">
              <i class="fa fa-crosshairs"></i>
              Run World
            </button>

            <button class="btn btn-dark btn-sm ml-3" @click="runProcess('zone')">
              <i class="fa fa-crosshairs"></i>
              Run Zone
            </button>

            <button class="btn btn-dark btn-sm ml-3" @click="runProcess('loginserver')">
              <i class="fa fa-crosshairs"></i>
              Run Loginsever
            </button>

            <button class="btn btn-dark btn-sm ml-3" @click="runProcess('ucs')">
              <i class="fa fa-crosshairs"></i>
              Run UCS
            </button>
          </div>

        </eq-window>

      </div>
      <div class="col-8">
        <eq-window title="Output">
          <!--          <v-runtime-template :template='`<pre class="mt-3" id="output" style="width: 100%; height: 75vh">${output}</pre>`'/>-->
          <pre
            class="mt-3 fade-in"
            id="output"
            style="width: 100%; height: 75vh; word-wrap: break-word; white-space: pre-wrap; overflow-x: hidden"
            v-html="output"
          ></pre>
        </eq-window>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindow               from "@/components/eq-ui/EQWindow.vue";
import {debounce}       from "@/app/utility/debounce";
import {SpireWebsocket} from "@/app/api/spire-websocket";


const Convert = require('ansi-to-html');
const convert = new Convert();

export default {
  name: "WebsocketPoc",
  components: {
    EqWindow
  },
  data() {
    return {
      // output: "",
      outputContainer: null,
      ansiRegex: null,
    }
  },
  created() {
    this.output = ""
  },
  mounted() {
    this.init()

    setTimeout(() => {
      this.outputContainer = document.getElementById("output");
    }, 1000)

    const pattern = [
      '[\\u001B\\u009B][[\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]+)*|[a-zA-Z\\d]+(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\\u0007)',
      '(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-nq-uy=><~]))'
    ].join('|');

    this.ansiRegex = new RegExp(pattern);
  },
  methods: {

    hello() {
      SpireWebsocket.connect().send(JSON.stringify({
        "action": "hello"
      }));
    },

    exec(command, args = []) {
      SpireWebsocket.connect().send(JSON.stringify({
        "action": "exec_server_bin",
        "command": command,
        "args": args
      }));
    },

    runProcess(process) {
      this.exec(process)
    },

    renderOutput: debounce(function () {
      this.$forceUpdate()

      setTimeout(() => {
        if (this.outputContainer) {
          this.outputContainer.scrollTop = this.outputContainer.scrollHeight + 100;
        }
      }, 1)

    }, 10),

    init() {
      SpireWebsocket.connect().onmessage = (evt) => {
        if (this.ansiRegex.test(evt.data)) {
          this.output += convert.toHtml(evt.data) + "\n"

        } else {
          this.output += evt.data + "\n"
        }

        if (this.outputContainer) {
          this.outputContainer.scrollTop = this.outputContainer.scrollHeight;
        }
        this.renderOutput()
      }

    }
  }
}
</script>

<style scoped>

</style>
