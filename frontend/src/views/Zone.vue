<template>
  <div>

    <!-- CONTENT -->
    <div class="container-fluid">
      <div class="panel-body">
        <div class="panel panel-default">


          <div class="row">
            <div class="col-12">
              <eq-window class="mt-5" title="Zone">

                <app-loader :is-loading="!loaded" padding="8"/>


                <h3 class="eq-header" v-if="zone && loaded">{{ zone.long_name }}</h3>

                <!--                <pre v-if="zone">{{ zone }}</pre>-->

                <eq-tabs v-if="loaded && npcs">

                  <!-- NPCS -->
                  <eq-tab :name="'NPC(s) ' + (npcCount ? '(' + npcCount + ')' : '')" :selected="true">
                    <npc-list-grid :npcs="npcs" v-if="npcs"/>
                  </eq-tab>

                  <!-- Doors -->
                  <eq-tab :name="'Doors(s) ' + (doorCount ? '(' + doorCount + ')' : '')">
                    <eq-auto-table :data="doors"/>
                  </eq-tab>

                  <!-- Objects -->
                  <eq-tab :name="'Objects(s) ' + (objectCount ? '(' + objectCount + ')' : '')">
                    <eq-auto-table :data="objects"/>
                  </eq-tab>

                  <!-- Spawn Entries -->
                  <eq-tab :name="'Spawn Entries(s) ' + (spawnEntriesCount ? '(' + spawnEntriesCount + ')' : '')">
                    <eq-auto-table :data="spawnEntries"/>
                  </eq-tab>

                </eq-tabs>

              </eq-window>
            </div>
          </div>

        </div>

      </div>
    </div>

  </div>
</template>

<script type="ts">
import {DoorApi, ObjectApi, Spawn2Api, ZoneApi} from "@/app/api/api";
import EqWindow                                 from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient}                         from "@/app/api/spire-api-client";
import EqCheckbox                               from "@/components/eq-ui/EQCheckbox.vue";
import {debounce}                               from "@/app/utility/debounce.js";
import EqTabs                                   from "@/components/eq-ui/EQTabs.vue";
import EqTab                                    from "@/components/eq-ui/EQTab.vue";
import util                                     from "util";
import EqAutoTable                              from "@/components/eq-ui/EQAutoTable.vue";
import NpcListGrid                              from "@/components/grids/NpcListGrid.vue";

export default {
  components: {
    NpcListGrid,
    EqAutoTable,
    EqTab,
    EqTabs,
    EqCheckbox,
    EqWindow,
  },
  data() {
    return {
      loaded: false,

      zone: null,
      routeWatcher: null,

      spawn2s: null,
      spawn2Count: 0,

      spawnEntries: null,
      spawnEntriesCount: 0,

      npcs: null,
      npcCount: 0,

      doors: null,
      doorCount: 0,

      objects: null,
      objectCount: 0,


    }
  },
  async activated() {
    await this.init()

    // route watcher
    this.routeWatcher = this.$watch('$route.query', () => {
      this.loadState()
    });

  },
  deactivated() {
    // remove route watcher
    this.routeWatcher()
  },
  methods: {
    async init() {
      this.loaded = false
      this.loadState()
    },


    setStateDebounce: debounce(function () {
      this.setState()
    }, 300),

    // let querystring updates drive state change
    setState() {
      const zoneId = this.$router.param.zoneId

      let query = {}
      this.$router.push(
        {
          path: '/zone/' + zoneId,
          query: query
        }
      ).catch(() => {
      })
    },

    // load state from query params
    loadState() {
      this.loadZone()
    },
    loadZone: async function () {

      this.npcs = null

      // load zone
      const zoneApi = (new ZoneApi(SpireApiClient.getOpenApiConfig()))
      zoneApi.getZone({id: this.$route.params.zoneId}).then((result) => {
        if (result.status === 200) {
          this.zone          = result.data
          this.filteredZones = result.data

          // fetch NPCS

          let filters = [
            ["zone", "__", this.zone.short_name],
            ["version", "__", this.zone.version],
          ]

          let wheres = [];
          filters.forEach((filter) => {
            wheres.push(util.format("%s%s%s", filter[0], filter[1], filter[2]))
          })

          const spawn2Api = (new Spawn2Api(SpireApiClient.getOpenApiConfig()))
          spawn2Api.listSpawn2s({where: wheres.join("."), includes: "2"}).then((result) => {
            if (result.status === 200) {
              // this.npcs     = result.data
              // this.npcCount = result.data.length

              // data
              let npcs         = []
              let spawnEntries = []

              // counts
              let spawnEntriesCount = 0
              let npcCount          = 0
              result.data.forEach((row) => {
                // console.log(row)

                // spawnentry
                if (row.spawnentries) {
                  row.spawnentries.forEach((spawnentry) => {

                    // npc_type
                    if (spawnentry.npc_type) {
                      npcCount++

                      // only add the npc to the list if it doesn't exist already
                      let npcExists = false
                      npcs.forEach((npc) => {
                        if (npc.id === spawnentry.npc_type.id) {
                          npcExists = true
                        }
                      })

                      if (!npcExists) {
                        npcs.push(spawnentry.npc_type)
                      }
                    }

                    spawnEntriesCount++
                    spawnEntries.push(row)
                  })
                }

              })

              // data
              this.npcs         = npcs
              this.spawnEntries = spawnEntries


              // counts
              this.npcCount          = npcCount
              this.spawnEntriesCount = spawnEntriesCount


              // console.log(JSON.stringify(Object.keys(npcs[0])))

              // this.gridOptions.gridReady(() => {
              //   console.log("GRID IS READY!")
              // })

              setTimeout(() => {
                // this.gridOptions.columnApi.autoSizeAllColumns();
              }, 300)


            }
          })

          const doorApi = (new DoorApi(SpireApiClient.getOpenApiConfig()))
          filters       = [
            ["zone", "__", this.zone.short_name],
            ["version", "__", this.zone.version],
          ]

          wheres = [];
          filters.forEach((filter) => {
            wheres.push(util.format("%s%s%s", filter[0], filter[1], filter[2]))
          })
          doorApi.listDoors({where: wheres.join(".")}).then((result) => {
            if (result.status === 200) {
              this.doors     = result.data
              this.doorCount = result.data.length
            }
          })

          const objectApi = (new ObjectApi(SpireApiClient.getOpenApiConfig()))
          filters         = [
            ["zoneid", "__", this.zone.zoneidnumber],
            ["version", "__", this.zone.version],
          ]

          wheres = [];
          filters.forEach((filter) => {
            wheres.push(util.format("%s%s%s", filter[0], filter[1], filter[2]))
          })
          objectApi.listObjects({where: wheres.join(".")}).then((result) => {
            if (result.status === 200) {
              this.objects     = result.data
              this.objectCount = result.data ? result.data.length : 0
            }
          })

          this.loaded = true
        }
      })


    }
  }
}

</script>

<style>
#zonetable TBODY TR TD {
  padding: 2px 4px;
}
</style>
