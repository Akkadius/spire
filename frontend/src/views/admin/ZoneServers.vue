<template>
  <div>
    <app-loader :is-loading="!loaded"></app-loader>

    <span v-if="zoneList.length === 0 && loaded">Zoneservers are offline</span>

    <div class="card mb-3 mr-4">
      <div class="card-body pt-0 pb-0 pl-3 pr-3">
        <div class="input-group input-group-flush">
          <div class="input-group-prepend">
            <span class="input-group-text"><i class="fe fe-search"></i></span>
          </div>
          <input
            class="list-search form-control"
            type="search"
            @keyup="updateQueryState"
            v-model="search"
            placeholder="Filter zones..."
          >
        </div>
      </div>
    </div>

    <div style="max-height: 80vh; overflow-y: scroll">
      <div
        :class="['card', 'mb-3']"
        v-if="processStats[zone.zone_os_pid]"
        v-for="zone in filterZoneList(zoneList)"
      >
        <div
          :class="['card-body', 'lift', 'btn-default', 'card-slim']"
          style="box-shadow: 0 2px 4px 0 rgba(30,55,90,.1);"
        >
          <div class="row align-items-center">

            <div class="col ml-n1">
              <h4 class="mb-1">

                <span class="h2 fe text-muted mb-0 fe-layers mr-3"></span>

                <a href="#" style="color:#2364d2;font-weight:200;font-size: 18px;" v-if="zone.zone_long_name">
                  {{ formatZoneName(zone.zone_long_name ? zone.zone_long_name : 'Standby Zone') }}
                </a>

                <span style="color:#666;font-weight:200;font-size: 18px;" class="text-muted" v-if="!zone.zone_long_name">
                Ready for players...
              </span>

                <router-link
                  class="text-muted ml-2"
                  :to="`zoneservers/${zone.client_port}/logs?zone=${zone.zone_long_name}`"
                  style="font-size:12px"
                >
                  <i class="fe fe-play"></i> Stream Logs
                </router-link>

                <router-link
                  class="text-muted ml-2"
                  :to="`zoneservers/${zone.client_port}/netstats`"
                  style="font-size:12px"
                >
                  <i class="fe fe-airplay"></i> Network Stats
                </router-link>

              </h4>
            </div>

            <div class="col-auto" v-if="zone.zone_name">
              <p class="card-text small text-muted mb-1 mt-2">
              <span class="text-muted ml-2" style="font-size:12px">
                  {{ formatZoneName(zone.zone_name) }}
                  ({{ zone.zone_id }})
                  <span v-if="zone.instance_id">Instance: {{ zone.instance_id }}</span>
                </span>
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
              <span class="text-muted ml-2" style="font-size:12px">
                  {{ zone.is_static_zone === true ? 'Static' : 'Dynamic' }}
              </span>
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
                <i class="fe fe-proc"></i> PID {{ zone.zone_os_pid }}
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
                <i class="fe fe-cloud"></i> Port {{ zone.client_port }}
              </p>
            </div>


            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2">
                <i class="fe fe-users"></i> {{ zone.number_players }}
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2" v-if="processStats[zone.zone_os_pid]">
                <i class="fe fe-cpu"></i> {{ parseFloat(processStats[zone.zone_os_pid].cpu).toFixed(2) }} %
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2" v-if="processStats[zone.zone_os_pid]">
                {{ parseFloat(processStats[zone.zone_os_pid].memory / 1024 / 1024).toFixed(2) }}MB
              </p>
            </div>

            <div class="col-auto">
              <p class="card-text small text-muted mb-1 mt-2" v-if="processStats[zone.zone_os_pid]">
                <i class="fe fe-clock"></i>
                {{ parseFloat(processStats[zone.zone_os_pid].elapsed / 1000 / 60 / 60).toFixed(2) }}h
              </p>
            </div>

            <div class="col-auto">
              <a
                class="text-muted"
                href="javascript:void(0)"
                @click="killZone(zone.zone_os_pid)"
                style="font-size:12px"
              >
                <i class="fa fa-power-off"></i> Kill Zone
              </a>
            </div>

          </div>
        </div>
      </div>
    </div>


  </div>
</template>

<script>
import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";
import {ROUTE}         from "@/routes";

export default {
  name: "ZoneServers",
  data() {
    return {
      search: "",

      zoneList: [],
      processStats: {},
      loaded: false,
      zoneServerLoop: null,
      chartLoaded: false
    }
  },

  watch: {
    '$route'() {
      this.loadQueryState()
    },
  },

  async created() {
    this.loadQueryState()
    this.getZoneList()

    this.zoneServerLoop = setInterval(() => {
      if (!document.hidden) {
        this.getZoneList()
      }
    }, 1000)
  },
  methods: {

    // state
    updateQueryState() {
      let q = {};
      if (this.search !== "") {
        q.search = this.search
      }

      this.$router.push(
        {
          path: ROUTE.ADMIN_ZONE_SERVERS,
          query: q
        }
      ).catch(() => {
      })
    },
    loadQueryState() {
      if (this.$route.query.search && this.$route.query.search.length > 0) {
        this.search = this.$route.query.search
      }
    },

    formatZoneName(name) {
      return name.replaceAll("UNKNOWN", "Idle")
    },

    async killZone(pid) {
      if (confirm("Are you sure that you want to kill this process pid (" + pid + ")?")) {
        await OcculusClient.killProcessByPid(pid)
        this.zoneList = []
        this.getZoneList()
      }
    },
    async getZoneList() {
      const zoneList = await OcculusClient.getZoneList()
      this.loaded    = true
      if (zoneList && zoneList.zone_list !== null && Object.keys(zoneList).length > 0 && Object.keys(zoneList.process_stats).length > 0) {
        this.zoneList     = zoneList.zone_list
        this.processStats = zoneList.process_stats
      }
    },
    filterZoneList(list) {
      let zoneList = []
      for (let z of list) {
        if (
          z.zone_name.toLowerCase().includes(this.search) ||
          z.zone_long_name.toLowerCase().includes(this.search) ||
          z.client_port.toString().includes(this.search)
        ) {
          zoneList.push(z)
        }
      }

      return zoneList
    }
  },
  destroyed() {
    clearInterval(this.zoneServerLoop)
  },
}
</script>

<style>
.avatar-image {
  background-image: url('~@/assets/img/eqemu-avatar.png');
}

.connector {
  width: 15px;
  height: 15px;
  opacity: .2;
  border-bottom-left-radius: 8px;
  border-left: 1px solid #1e375a;
  border-bottom: 1px solid #1e375a;
  margin-left: 9px;
  display: inline-flex;
}

.image-label {
  background-color: rgba(35, 100, 210, .07);
  color: #2364d2;
  padding: 0 4px;
  border-radius: 2px;
}
</style>
