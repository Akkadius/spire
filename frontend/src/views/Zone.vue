<template>
  <content-area>
    <div class="row">
      <div class="col-12" v-if="zone && loaded">
        <eq-window-simple class="p-0 mt-4" :title="zone.long_name">
          <loader-fake-progress v-if="!npcs" class="mt-5 mb-5"/>

          <eq-tabs v-if="loaded && npcs" style="height: 90vh">

            <!-- NPCS -->
            <eq-tab
              :name="'NPC(s) ' + (npcCount ? '(' + npcCount + ')' : '')"
              :selected="true"
            >
              <npc-list-grid :npcs="npcs" v-if="npcs"/>
            </eq-tab>

            <!-- Doors -->
            <eq-tab
              :name="'Doors(s) ' + (doorCount ? '(' + doorCount + ')' : '')"
            >
              <eq-auto-table :data="doors" style="height: 85vh"/>
            </eq-tab>

            <!-- Objects -->
            <eq-tab
              :name="'Objects(s) ' + (objectCount ? '(' + objectCount + ')' : '')"
            >
              <eq-auto-table :data="objects" style="height: 85vh"/>
            </eq-tab>

            <!-- Spawn Entries -->
            <eq-tab
              :name="'Spawn Entries(s) ' + (spawnEntriesCount ? '(' + spawnEntriesCount + ')' : '')"
            >
              <eq-auto-table :data="spawnEntries" style="height: 85vh"/>
            </eq-tab>

          </eq-tabs>

        </eq-window-simple>
      </div>
    </div>
  </content-area>
</template>

<script type="ts">
import {DoorApi, ObjectApi, Spawn2Api, ZoneApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import {debounce} from "@/app/utility/debounce.js";
import EqTabs from "@/components/eq-ui/EQTabs.vue";
import EqTab from "@/components/eq-ui/EQTab.vue";
import EqAutoTable from "@/components/eq-ui/EQAutoTable.vue";
import NpcListGrid from "@/components/grids/NpcListGrid.vue";
import ContentArea from "@/components/layout/ContentArea.vue";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";

export default {
  components: {
    LoaderFakeProgress,
    EqWindowSimple,
    ContentArea,
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

      this.npcs = null;

      // load zone
      (new ZoneApi(SpireApiClient.getOpenApiConfig()))
        .getZone({id: this.$route.params.zoneId})
        .then((result) => {
          if (result.status === 200) {
            this.zone          = result.data;
            this.filteredZones = result.data;

            // fetch NPCS
            (new Spawn2Api(SpireApiClient.getOpenApiConfig())).listSpawn2s(
              (new SpireQueryBuilder())
                .where("zone", "=", this.zone.short_name)
                .where("version", "=", this.zone.version)
                .includes(["2"])
                .get()
            ).then((result) => {
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
            });

            (new DoorApi(SpireApiClient.getOpenApiConfig())).listDoors(
              (new SpireQueryBuilder())
                .where("zone", "=", this.zone.short_name)
                .where("version", "=", this.zone.version)
                .includes(["2"])
                .get()
            ).then((result) => {
              if (result.status === 200) {
                this.doors     = result.data
                this.doorCount = result.data.length
              }
            })

            const objectApi = (new ObjectApi(SpireApiClient.getOpenApiConfig()))
            objectApi.listObjects(
              (new SpireQueryBuilder())
                .where("zoneid", "=", this.zone.zoneidnumber)
                .where("version", "=", this.zone.version)
                .includes(["2"])
                .get()
            ).then((result) => {
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
