<template>
  <div>

    <div class="row">

      <div class="col-6">
        <div class="card">
          <div class="card-body">
            <div class="row align-items-center">
              <div class="col-sm-12 col-lg-6">
                <h6 class="header-pretitle">
                  Port: {{ zonePort }}
                </h6>

                <h1 class="header-title">
                  {{ zoneAttributes.long_name }}
                </h1>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="col-sm-6 col-lg-3">
        <!-- Player count -->
        <div class="card p-3">
          <div class="d-flex align-items-center">
              <span class="stamp stamp-md bg-red mr-3">
                <i class="fe fe-users"></i>
              </span>
            <div>
              <h4 class="m-0">
                <span class="player-count mr-1"> {{ zoneClientListCount }}</span>
                <small>Players</small>
              </h4>
              <small class="text-muted">Currently in zone</small>
            </div>
          </div>
        </div>
      </div>

      <div class="col-sm-6 col-lg-3">
        <!-- NPC Count -->
        <div class="card p-3">
          <div class="d-flex align-items-center">
              <span class="stamp stamp-md bg-yellow mr-3">
                <i class="fe fe-gitlab"></i>
              </span>
            <div>
              <h4 class="m-0">
                <span class="npc-count mr-1"> {{ zoneNpcListCount }} </span>
                <small>NPCs</small>
              </h4>
              <small class="text-muted">Currently in zone</small>
            </div>
          </div>
        </div>
      </div>
    </div>


    <app-loader :is-loading="!loaded"></app-loader>

    <div v-show="loaded">

      <!-- Packet Types Sent -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Packet Types Sent from Server</h3>
        </div>
        <div class="card-body">
          <div id="packet-sent-types" style="height: 16rem"></div>
        </div>
      </div>

      <!-- Packet Types Received -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Packet Types Received from Server</h3>
        </div>
        <div class="card-body">
          <div id="packet-receive-types" style="height: 16rem"></div>
        </div>
      </div>

      <!-- Ping -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Ping (ms)</h3>
        </div>
        <div class="card-body">
          <div id="ping" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Packet Loss Inbound -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Packet Loss Inbound %</h3>
        </div>
        <div class="card-body">
          <div id="packet-loss-inbound" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Packet Loss Outbound -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Packet Loss Outbound %</h3>
        </div>
        <div class="card-body">
          <div id="packet-loss-outbound" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Receive bytes per second -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Server Receive (mbps)</h3>
        </div>
        <div class="card-body">
          <div id="receive-bytes" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Sent bytes per second -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Server Sent (mbps)</h3>
        </div>
        <div class="card-body">
          <div id="sent-bytes" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Resent Packets -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Resent Packets</h3>
        </div>
        <div class="card-body">
          <div id="resent-packets" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Resent Fragments -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Resent Fragments</h3>
        </div>
        <div class="card-body">
          <div id="resent-fragments" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Resent Non-Fragments -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Resent Non-Fragments</h3>
        </div>
        <div class="card-body">
          <div id="resent-non-fragments" style="height: 8rem"></div>
        </div>
      </div>

      <!-- Netstats from all clients -->
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Client Netstats</h3>
        </div>
        <div class="table-responsive">
          <table class="table card-table table-vcenter text-nowrap" id="client-netstats">
            <thead>
            <tr>
              <th>Client ID</th>
              <th>Client Name</th>
              <th>Average Ping</th>
              <th>Last Ping</th>
              <th>Min Ping</th>
              <th>Max Ping</th>
              <th>Packet Loss In %</th>
              <th>Packet Loss Out %</th>
              <th>Realtime Receive Packets</th>
              <th>Realtime Sent Packets</th>
              <th>Receive bps</th>
              <th>Sent bps</th>
              <th>Resent Packets</th>
              <th>Resent Fragments</th>
              <th>Resent Non Fragments</th>
              <th>Seconds Since Reset</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <td></td>
              <td></td>
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

export default {
  name: 'NetStats',
  data() {
    return {
      zoneAttributes: {},
      zoneClientList: {},
      zoneClientListCount: 0,
      zoneNpcList: {},
      zoneNpcListCount: 0,
      zonePort: 0,
      loaded: false,
      chart_packet_types_sent: {},
      chart_packet_types_receive: {},
      chart_ping: {},
      chart_packet_loss_inbound: {},
      chart_packet_loss_outbound: {},
      chart_receive_bytes: {},
      chart_sent_bytes: {},
      chart_resent_packets: {},
      chart_resent_fragments: {},
      chart_resent_non_fragments: {}
    }
  },
  async created() {
    this.zonePort = this.$route.params.port

    OcculusClient.getZoneAttributes(this.zonePort).then(response => {
      this.zoneAttributes = response[0]
    })

    OcculusClient.getZoneClientList(this.zonePort).then(response => {
      this.zoneClientList      = response
      this.zoneClientListCount = Object.keys(response).length
    })

    OcculusClient.getZoneNpcList(this.zonePort).then(response => {
      this.zoneNpcList      = response
      this.zoneNpcListCount = Object.keys(response).length
    })

    this.getClientNetstats()

    var chart_base_config = {
      bindto: '#packet-loss-inbound',
      transition: {
        duration: 0
      },
      point: {
        show: false
      },
      line: {
        connectNull: true
      },
      data: {},
      legend: {
        show: false
      },
      axis: {
        x: {
          type: 'timeseries',
          tick: {
            format: '%H:%M:%S',
            fit: true
          }
        },
        y: {
          show: true
        }
      }
    }

    /**
     * Initialize Charts
     * @type {string}
     */
    const data = await OcculusClient.getZoneNetstatChartData(this.zonePort)

    this.$options.chart_packet_types_sent = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#packet-sent-types',
          data: data.sent_packet_types
        }
      ))

    this.$options.chart_packet_types_receive = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#packet-receive-types',
          data: data.receive_packet_types
        }
      ))

    this.$options.chart_ping = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#ping',
          data: data.ping
        }
      ))

    this.$options.chart_packet_loss_inbound = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#packet-loss-inbound',
          data: data.packet_loss_inbound
        }
      ))

    this.$options.chart_packet_loss_outbound = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#packet-loss-outbound',
          data: data.packet_loss_outbound
        }
      ))

    this.$options.chart_receive_bytes = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#receive-bytes',
          data: data.receive_bytes
        }
      ))

    this.$options.chart_sent_bytes = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#sent-bytes',
          data: data.sent_bytes
        }
      ))

    this.$options.chart_resent_packets = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#resent-packets',
          data: data.resent_packets
        }
      ))

    this.$options.chart_resent_fragments = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#resent-fragments',
          data: data.resent_fragments
        }
      ))

    this.$options.chart_resent_non_fragments = c3.generate(
      $.extend(chart_base_config, {
          bindto: '#resent-non-fragments',
          data: data.resent_non_fragments
        }
      ))

    this.loaded = true

    this.$bvToast.toast(
     'Netstat listener started...',
      {
        title: "Netstats",
        toaster: 'b-toaster-top-center',
        variant: 'info',
        autoHideDelay: 3000,
        appendToast: false
      }
    )

    this.refreshCharts()

    var self            = this
    this.$options.timer = setInterval(function () {
      self.refreshCharts()
      self.getClientNetstats()
    }, 5000)
  },
  beforeDestroy() {
    clearInterval(this.$options.timer)
  },
  methods: {
    getClientNetstats() {
      OcculusClient.getZoneClientNetstats(this.zonePort).then(data => {
        $('#client-netstats tbody').empty()

        for (var row in data) {
          var client_stats = [
            data[row].client_id,
            data[row].client_name,
            this.commify(data[row].average_ping) + ' ms',
            this.commify(data[row].last_ping) + ' ms',
            this.commify(data[row].min_ping) + ' ms',
            this.commify(data[row].max_ping) + ' ms',
            parseFloat(data[row].packet_loss_in).toFixed(2),
            parseFloat(data[row].packet_loss_out).toFixed(2),
            this.commify(data[row].realtime_receive_packets),
            this.commify(data[row].realtime_sent_packets),
            parseFloat(data[row].receive_bytes / data[row].seconds_since_reset).toFixed(1) + ' bps',
            parseFloat(data[row].sent_bytes / data[row].seconds_since_reset).toFixed(1) + ' bps',
            this.commify(data[row].resent_packets),
            this.commify(data[row].resent_fragments),
            this.commify(data[row].resent_non_fragments),
            parseFloat(data[row].seconds_since_reset).toFixed(2)
          ]

          var row_data = ''
          for (var stat in client_stats) {
            row_data = row_data + '<td>' + client_stats[stat] + '</td>'
          }

          $('#client-netstats > tbody').append('<tr>' + row_data + '</tr>')
        }
      })
    },
    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
    },
    refreshCharts() {
      OcculusClient.getZoneNetstatChartData(this.zonePort).then(data => {
        this.$options.chart_packet_types_sent.load(data.sent_packet_types)
        this.$options.chart_packet_types_receive.load(data.receive_packet_types)
        this.$options.chart_ping.load(data.ping)
        this.$options.chart_packet_loss_inbound.load(data.packet_loss_inbound)
        this.$options.chart_packet_loss_outbound.load(data.packet_loss_outbound)
        this.$options.chart_receive_bytes.load(data.receive_bytes)
        this.$options.chart_sent_bytes.load(data.sent_bytes)
        this.$options.chart_resent_packets.load(data.resent_packets)
        this.$options.chart_resent_fragments.load(data.resent_fragments)
        this.$options.chart_resent_non_fragments.load(data.resent_non_fragments)
      })
    }
  }
}
</script>

<style scoped>

</style>
