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

                    <pre v-if="npcStatusMessage" style="width: 100%; margin: 0px"
                         class="mb-2">{{ npcStatusMessage }}</pre>
                    <ag-grid-vue
                      v-if="npcRowData"
                      style="width: 100%; height: 75vh;"
                      class="ag-theme-balham-dark"
                      :columnDefs="npcColumnDefs"
                      :gridOptions="gridOptions"
                      :rowData="npcRowData">
                    </ag-grid-vue>
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
import {DoorApi, NpcTypeApi, ObjectApi, Spawn2Api, ZoneApi} from "@/app/api/api";
import EqWindow                                             from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient}                                     from "@/app/api/spire-api-client";
import EqCheckbox                                           from "@/components/eq-ui/EQCheckbox.vue";
import ZoneForm                                             from "@/components/forms/ZoneForm.vue";
import {debounce}                                           from "@/app/utility/debounce.js";
import EqTabs                                               from "@/components/eq-ui/EQTabs.vue";
import EqTab                                                from "@/components/eq-ui/EQTab.vue";
import util                                                 from "util";
import EqAutoTable                                          from "@/components/eq-ui/EQAutoTable.vue";
import {AgGridVue}                                          from "ag-grid-vue";

export default {
  components: {
    AgGridVue,
    EqAutoTable,
    EqTab,
    EqTabs,
    ZoneForm,
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

      columnDefs: null,
      rowData: null,

      npcStatusMessage: "",
      npcColumnDefs: null,
      npcRowData: null,

      gridOptions: {
        defaultColDef: {
          resizable: true,
          sortable: true,
          searchable: true,
          editable: true,
          filter: 'agTextColumnFilter',
          floatingFilter: true,
        },


        // pagination: true,
        rowSelection: 'single',

        rowGroupPanelShow: 'always',
        pivotPanelShow: 'always',

        enablePivot: false,

        sideBar: {
          toolPanels: ['columns'],
        },

        onCellValueChanged: (params) => {
          const row           = params.data
          const columnChanged = params.column.colId
          const newValue      = params.newValue
          const npcApi        = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))

          // make sure we keep the integer types the same before we send to the
          // backend since it is strict
          if (parseInt(newValue) > 0 || parseInt(newValue) < 0) {
            row[columnChanged] = Number(newValue)
          }

          // updated object
          let update            = {}
          update.id             = row.id
          update[columnChanged] = row[columnChanged]

          npcApi.updateNpcType({
            id: row.id,
            npcType: update
          }).then((result) => {
            if (result.status === 200) {
              this.npcStatusMessage = util.format(
                "Updated NPC [%s] [%s] Column [%s] to Value [%s]",
                row.id,
                row.name,
                columnChanged,
                newValue,
              )
            }
          }).catch((err) => {
            this.npcStatusMessage = err.response.data.error
          })


          // console.log(params)

          // trigger filtering on cell edits
          params.api.onFilterChanged();
        },

        sideBar: true,
      },
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

      this.npcRowData = null

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


              let npcColumnDefs = [
                {
                  headerName: 'NPC',
                  children: [
                    {width: 100, headerName: 'id', field: 'id', pinned: true},
                    {width: 200, headerName: 'Name', field: 'name', pinned: true},
                    {headerName: 'Lastname', field: 'lastname'}
                  ]
                },
                {
                  headerName: 'Meta',
                  children: [
                    {width: 75, field: "level",},
                    {
                      width: 130,
                      field: "class"
                    },
                  ]
                },
                {
                  headerName: 'Stats',
                  children: [
                    {width: 100, headerName: "hp", field: "hp"},
                    {width: 100, headerName: "mana", field: "mana"},
                    {width: 75, headerName: "ac", field: "ac"},
                    {width: 75, headerName: "hp_regen_rate", field: "hp_regen_rate"},
                    {width: 75, headerName: "mana_regen_rate", field: "mana_regen_rate"},
                    {width: 75, headerName: "mindmg", field: "mindmg"},
                    {width: 75, headerName: "maxdmg", field: "maxdmg"},
                    {width: 75, headerName: "attack_count", field: "attack_count"},
                    {width: 75, headerName: "str", field: "str"},
                    {width: 75, headerName: "sta", field: "sta"},
                    {width: 75, headerName: "dex", field: "dex"},
                    {width: 75, headerName: "agi", field: "agi"},
                    {width: 75, headerName: "_int", field: "_int"},
                    {width: 75, headerName: "wis", field: "wis"},
                    {width: 75, headerName: "cha", field: "cha"},
                    {width: 75, headerName: "atk", field: "atk"},
                  ]
                },
                {
                  headerName: 'Relational Data',
                  children: [
                    {field: "loottable_id", headerName: "loottable_id",},
                    {field: "merchant_id", headerName: "merchant_id",},
                    {field: "alt_currency_id", headerName: "alt_currency_id",},
                    {field: "npc_spells_id", headerName: "npc_spells_id",},
                    {field: "npc_spells_effects_id", headerName: "npc_spells_effects_id",},
                    {field: "npc_faction_id", headerName: "npc_faction_id",},
                    {field: "adventure_template_id", headerName: "adventure_template_id",},
                    {field: "trap_template", headerName: "trap_template",},
                    {field: "emoteid", headerName: "emoteid",},
                  ]
                },
                {
                  headerName: 'Resists',
                  children: [
                    {width: 100, headerName: "Magic", field: "mr",},
                    {width: 100, headerName: "Cold", field: "cr",},
                    {width: 100, headerName: "Disease", field: "dr",},
                    {width: 100, headerName: "Fire", field: "fr",},
                    {width: 100, headerName: "Poison", field: "pr",},
                    {width: 120, headerName: "Corruption", field: "corrup",},
                    {width: 100, headerName: "Physical", field: "ph_r",},
                  ]
                },
                {
                  headerName: 'Aggro / Assist',
                  children: [
                    {field: "aggroradius",},
                    {field: "assistradius",},
                    {field: "npc_aggro",},
                  ]
                },
                {
                  headerName: 'Appearance',
                  children: [
                    {headerName: "race", field: "race",},
                    {headerName: "gender", field: "gender",},
                    {headerName: "texture", field: "texture",},
                    {headerName: "bodytype", field: "bodytype",},
                    {headerName: "helmtexture", field: "helmtexture",},
                    {headerName: "herosforgemodel", field: "herosforgemodel",},
                    {headerName: "size", field: "size",},
                  ]
                },
                {
                  headerName: 'Armor Appearance',
                  children: [
                    {headerName: "armortint_id", field: "armortint_id",},
                    {headerName: "armortint_red", field: "armortint_red",},
                    {headerName: "armortint_green", field: "armortint_green",},
                    {headerName: "armortint_blue", field: "armortint_blue",},
                    {headerName: "d_melee_texture_1", field: "d_melee_texture_1",},
                    {headerName: "d_melee_texture_2", field: "d_melee_texture_2",},
                    {headerName: "armtexture", field: "armtexture",},
                    {headerName: "bracertexture", field: "bracertexture",},
                    {headerName: "handtexture", field: "handtexture",},
                    {headerName: "legtexture", field: "legtexture",},
                    {headerName: "feettexture", field: "feettexture",},
                  ]
                },
                {
                  headerName: 'Facial / Tattoo',
                  children: [
                    {field: "face",},
                    {field: "luclin_hairstyle",},
                    {field: "luclin_haircolor",},
                    {field: "luclin_eyecolor",},
                    {field: "luclin_eyecolor_2",},
                    {field: "luclin_beardcolor",},
                    {field: "luclin_beard",},
                    {field: "drakkin_heritage",},
                    {field: "drakkin_tattoo",},
                    {field: "drakkin_details",},
                  ]
                },
                {
                  headerName: 'Special Abilities',
                  children: [
                    {width: 200, headerName: "Special Abilities", field: "special_abilities",},
                    {headerName: "Special Attacks (Deprecated)", field: "npcspecialattks", hide: true},
                  ]
                },
                {
                  headerName: 'Charm',
                  children: [
                    {headerName: "charm_ac", field: "charm_ac"},
                    {headerName: "charm_min_dmg", field: "charm_min_dmg"},
                    {headerName: "charm_max_dmg", field: "charm_max_dmg"},
                    {headerName: "charm_attack_delay", field: "charm_attack_delay"},
                    {headerName: "charm_accuracy_rating", field: "charm_accuracy_rating"},
                    {headerName: "charm_avoidance_rating", field: "charm_avoidance_rating"},
                    {headerName: "charm_atk", field: "charm_atk"},
                  ]
                },
                {
                  headerName: 'Attributes',
                  children: [
                    {width: 120, field: "untargetable", headerName: "untargetable", },
                    {width: 120, field: "findable", headerName: "findable", },
                    {width: 120, field: "see_hide", headerName: "see_hide", },
                    {width: 120, field: "see_improved_hide", headerName: "see_improved_hide", },
                    {width: 120, field: "trackable", headerName: "trackable", },
                    {width: 120, field: "underwater", headerName: "underwater", },
                    {width: 120, field: "isquest", headerName: "isquest", },
                    {width: 120, field: "light", headerName: "light", },
                    {width: 120, field: "fixed", headerName: "fixed", },
                    {width: 120, field: "ignore_despawn", headerName: "ignore_despawn", },
                    {width: 120, field: "show_name", headerName: "show_name", },
                    {width: 120, field: "untargetable", headerName: "untargetable", },
                    {width: 120, field: "rare_spawn", headerName: "rare_spawn", },
                    {width: 120, field: "flymode", headerName: "flymode", },
                    {width: 120, field: "always_aggro", headerName: "always_aggro" },
                    {width: 120, field: "see_invis", headerName: "see_invis" },
                    {width: 120, field: "see_invis_undead", headerName: "see-invis_undead" },
                  ]
                },
              ]

              if (this.npcs[0]) {
                Object.keys(this.npcs[0]).forEach((key) => {
                  if (typeof this.npcs[0][key] !== 'object') {

                    let hasField = false
                    npcColumnDefs.forEach((def) => {
                      if (def.children) {
                        def.children.forEach((child) => {
                          if (key === child.field) {
                            hasField = true
                          }
                        })
                      }
                    })

                    if (!hasField) {
                      npcColumnDefs.push({headerName: key, field: key})
                    }
                  }
                })

                this.npcColumnDefs = npcColumnDefs
                this.npcRowData    = npcs
              }

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
