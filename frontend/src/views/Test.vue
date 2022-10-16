<template>
  <div>
    <page-header title="Test" pre-title=""/>

    <!-- CONTENT -->
    <div>
      <div class="container-fluid">
        <div class="panel-body">
          <div class="panel panel-default">
            <div class="row">
              <div class="col-6">
                <test-form :model="model"/>
              </div>
              <div class="col-6">
                <div style="height: auto; max-height:70vh; overflow-y: scroll">
                  <li v-for="zone in zones" :key="zone.id">
                    <router-link :to="`/test/` + zone.id">{{ zone.long_name }} ({{ zone.short_name }})
                      {{ zone.zoneidnumber }}
                    </router-link>
                  </li>
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {ZoneApi}  from "@/app/api/api";
import {SpireApi} from "../app/api/spire-api";

export default {
  components: {
    "page-header": () => import("@/components/layout/PageHeader")
  },

  data() {
    return {
      model: {},
      zones: {}
    }
  },

  watch: {
    $route(to, from) {
      this.loadEntity()
    }
  },

  methods: {
    async loadEntity() {
      (new ZoneApi()).getZone({ id: this.$route.params.zone }).then((result) => {
        if (result.status === 200) {
          this.model = result.data
        }
      })
    }
  },

  async mounted() {
    (new ZoneApi()).listZones({ orderBy: "zoneidnumber" }).then((result) => {
      if (result.status === 200) {
        this.zones = result.data
      }
    })

    const result = await SpireApi.v1().get('/static-map/race-inventory-map.json')
    console.log(result)

    this.loadEntity()
  }
}
</script>
