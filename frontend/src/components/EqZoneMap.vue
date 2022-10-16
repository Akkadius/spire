<template>
  <div>
    <eq-window class="p-2" style="height: 96vh">

      <!-- Loader -->
      <eq-window
        class="text-center"
        style="position: absolute; right: 3%; z-index: 99; padding: 15px; padding-top: 10px;"
        v-if="!isDataLoaded()"
      >
        <div class="mb-2">
          {{ isDataLoaded() ? 'Rendering map...' : 'Loading map...' }}
        </div>
        <loader-fake-progress v-if="!isDataLoaded()"/>
        <eq-progress-bar :percent="100" v-if="isDataLoaded()"/>
      </eq-window>

      <div class="card">
        <l-map
          v-if="center"
          :crs="crs"
          style="height: 94vh"
          class="map-tiles"
          :center="center"
          :bounds="bounds"
          :min-zoom="-5"
          :zoom="zoom"
          :zoom-animation="true"
          :zoom-animation-threshold="10"
          @update:zoom="zoomUpdate"
        >

          <!-- Draw map lines -->
          <l-polyline
            v-if="lines"
            :lat-lngs="lines"
            color="gray"
            :weight="1"
          />

          <!-- grid points -->
          <l-marker
            v-for="(m, index) in pathingGridMarkers"
            :key="index + '-' + m.point.lat + '-' + m.point.lng"
            :lat-lng="m.point"
            v-if="markers && markers.length > 0"
          >
            <l-tooltip :options="{ permanent: true, interactive: true }">
              {{ m.label }}
            </l-tooltip>
          </l-marker>

          <!-- Draw pathing grid lines -->
          <l-polyline
            v-if="pathingGridLines"
            :lat-lngs="pathingGridLines"
            color="blue"
            dashArray="5, 10"
            :opacity=".8"
            :weight="2"
          />

          <!--          &lt;!&ndash; markers from map &ndash;&gt;-->
          <!--          <l-marker-->
          <!--            v-for="(marker, index) in markers"-->
          <!--            :key="index"-->
          <!--            :lat-lng="marker.point"-->
          <!--            v-if="markers && markers.length > 0"-->
          <!--          >-->
          <!--            <l-tooltip>-->
          <!--              <eq-window>{{ marker.label }}-->
          <!--              </eq-window>-->
          <!--            </l-tooltip>-->
          <!--          </l-marker>-->

          <!-- zone points -->
          <l-marker
            v-for="(m, index) in zonelineMarkers"
            :key="index"
            :lat-lng="m.point"
            v-if="markers && markers.length > 0"
            @click="navigateToZone(m.zone.short_name, m.zone.version)"
          >
            <l-tooltip>
              <eq-window>{{ m.label }}
              </eq-window>
            </l-tooltip>
          </l-marker>

          <!-- door zone points -->
          <l-marker
            v-for="(m, index) in doorZonePoints"
            :key="index + '-' + m.destName + '-' + m.destInstance"
            :lat-lng="m.point"
            v-if="markers && markers.length > 0"
            @click="navigateToZone(m.destName, m.destInstance)"
          >
            <l-tooltip>
              <eq-window>{{ m.label }}
              </eq-window>
            </l-tooltip>
          </l-marker>

          <!-- NPC markers -->
          <l-marker
            v-for="(marker, index) in npcMarkers"
            :key="index + '-' + marker.npc.id"
            :lat-lng="marker.point"
            :opacity="getNpcOpacity(index + '-' + marker.npc.id, marker.npc.id)"
            @mouseover="npcMarkerHover(marker, index + '-' + marker.npc.id)"
            v-if="npcMarkers && npcMarkers.length > 0"
          >

            <l-tooltip>
              <eq-window>
                {{ getCleanName(marker.npc.name) }}
              </eq-window>
            </l-tooltip>

            <l-icon
              icon-url="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
              :class-name="(zoomLevel >= 1) ? marker.iconClass : marker.iconClass + '-sm'"
              :iconSize="(zoomLevel >= 1) ? marker.iconSize : calcSmallIcons(marker.iconSize)"
            >
            </l-icon>

          </l-marker>

          <!-- Door markers -->
          <l-marker
            v-for="(marker, index) in doorMarkers"
            :key="index + '-' + marker.name"
            :lat-lng="marker.point"
            v-if="doorMarkers && doorMarkers.length > 0"
          >

            <l-tooltip>
              <eq-window>
                {{ marker.label }}
              </eq-window>
            </l-tooltip>

            <l-icon
              icon-url="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
              :class-name="marker.iconClass"
              :iconSize="marker.iconSize"
            >
            </l-icon>
          </l-marker>

          <!-- Translocate markers -->
          <l-marker
            v-for="(m, index) in translocatePoints"
            :key="index + '-' + m.label"
            :lat-lng="m.point"
            @mouseover="spellMarkerHover(m.spell)"
            v-if="translocatePoints && translocatePoints.length > 0"
            style="border-radius: 10px"
          >
            <l-tooltip>
              <eq-window>
                {{ m.label }}
              </eq-window>
            </l-tooltip>

            <l-icon
              icon-url="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
              :class-name="m.iconClass"
              :iconSize="m.iconSize"
            >
            </l-icon>
          </l-marker>

          <!-- Safe coordinate markers -->
          <l-marker
            v-for="(m, index) in safeCoordinateMarker"
            :key="index + '-' + m.label"
            :lat-lng="m.point"
            v-if="safeCoordinateMarker && safeCoordinateMarker.length > 0"
            style="border-radius: 10px"
          >
            <l-tooltip>
              <eq-window>
                {{ m.label }}
              </eq-window>
            </l-tooltip>

            <l-icon
              icon-url="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
              :class-name="m.iconClass"
              :iconSize="m.iconSize"
            >
            </l-icon>
          </l-marker>

        </l-map>
      </div>
    </eq-window>
  </div>
</template>

<script>
import {LIcon, LMap, LMarker, LPolyline, LPopup, LTileLayer, LTooltip} from 'vue2-leaflet';
import ContentArea                                                     from "./layout/ContentArea";
import * as L                                                          from "leaflet";
import axios                                                           from "axios";
import {GridEntryApi, Spawn2Api, SpellsNewApi, ZonePointApi} from "../app/api";
import {SpireApi}                                            from "../app/api/spire-api";
import {SpireQueryBuilder}                                   from "../app/api/spire-query-builder";
import EqNpcCardPreview                                      from "./preview/EQNpcCardPreview";
import EqWindow                                              from "./eq-ui/EQWindow";
import LoaderFakeProgress                                    from "./LoaderFakeProgress";
import EqProgressBar                                         from "./eq-ui/EQProgressBar";
import {Npcs}                                                from "../app/npcs";
import {Zones}                                               from "../app/zones";
import {DoorApi}                                             from "../app/api/api/door-api";
import {EventBus}                                            from "../app/event-bus/event-bus";
import {Spawn}                                               from "../app/spawn";

export default {
  name: "EqZoneMap",
  props: {
    zone: {
      type: String,
      required: true
    },
    version: {
      type: String,
      required: true
    },
  },
  components: {
    EqProgressBar,
    LoaderFakeProgress,
    EqWindow,
    EqNpcCardPreview,
    ContentArea,
    LMap,
    LIcon,
    LMarker,
    LPopup,
    LTooltip,
    LTileLayer,
    LPolyline
  },
  watch: {
    zone: {
      handler(newVal) {
        this.loadMap()
      },
      deep: true
    },
  },

  methods: {

    handleNpcZoomEvent(e) {
      console.log("[EqZoneMap] Handling Zoom event")
      console.log(e)

      this.zoomedNpcId = e.id

      for (let n of this.npcMarkers) {
        if (n.npc.id === e.id) {
          console.log("found NPC marker at ", n)

          // zoom out first
          this.zoom = this.starterZoomLevel

          // center on target
          setTimeout(() => {
            this.center = n.point
          }, 600)

          // zoom in
          setTimeout(() => {
            this.zoom = 1
          }, 1000)

          break;
        }
      }

    },

    getNpcOpacity(elementKey, npcId) {
      if (this.zoomedNpcId === npcId) {
        return 1;
      }

      if (this.zoomedNpcId > 0 && this.zoomedNpcId !== npcId) {
        return .3;
      }

      if (this.hoveredNpc === "") {
        return 1;
      }

      if (this.hoveredNpc !== "" && this.hoveredNpc !== elementKey) {
        return .3;
      }

      return 1;
    },

    // this is not a computed property because the dependencies are not reactive
    isDataLoaded() {
      return this.npcMarkers
    },

    navigateToZone(shortName, version) {
      this.$router.push(
        {
          path: `/zone/${shortName}?v=${version}`
        }
      ).catch(() => {
      })
    },

    getCleanName(n) {
      return Npcs.getCleanName(n)
    },

    npcMarkerHover(e, elementKey) {
      // console.log(e)

      // reset
      this.hoveredNpc  = ""
      this.zoomedNpcId = 0
      if (this.pathingGridLines.length > 0) {
        this.pathingGridLines   = []
        this.pathingGridMarkers = []
        this.$forceUpdate()
      }

      if (e.grid > 0) {
        // transform grid entries into poly lines
        let polyLines   = []
        let gridMarkers = []
        for (const [id, g] of this.pathingGridData.entries()) {
          if (g && id === e.grid) {
            this.hoveredNpc = elementKey

            for (const [i, e] of g.entries()) {

              // make sure we have a valid entry as well as
              // a valid next point so we can draw a complete line
              if (e && e.x && g[i + 1]) {
                // console.log(i, e)
                const current = e
                const next    = g[i + 1]
                polyLines.push(
                  [
                    this.createPoint(-current.x, -current.y),
                    this.createPoint(-next.x, -next.y),
                  ]
                )
              }

              if (e && e.x) {
                // console.log(i, e)
                gridMarkers.push(
                  {
                    point: this.createPoint(-e.x, -e.y),
                    label: i,
                  }
                )
              }
            }
          }
        }

        this.pathingGridLines   = polyLines
        this.pathingGridMarkers = gridMarkers
      }

      this.$emit("npc-marker-hover", e.npc);
    },

    spellMarkerHover(s) {
      this.$emit("spell-marker-hover", s);
    },

    calcSmallIcons(xy) {
      // console.log("small icon")
      // console.log(xy)

      return [
        xy[0] / 2,
        xy[1] / 2,
      ]
    },

    getNpcIcon(npc) {
      return 'race-models-ctn-' + npc.race + '-' + npc.gender + '-' + npc.texture + '-' + npc.helmtexture;
    },

    zoomUpdate(e) {
      console.log("zoom level [%s]", e)

      if (this.starterZoomLevel === -100) {
        this.starterZoomLevel = e
      }

      this.zoomLevel = e
    },
    createPoint(x, y) {
      return L.latLng(
        (typeof (y) === "string" ? -parseFloat(y) : -y),
        (typeof (x) === "string" ? parseFloat(x) : x));
    },
    iconClass() {
      return this.zoomLevel >= 2 ? 'item-4472' : 'item-4472-sm'
    },
    iconSize() {
      return this.zoomLevel >= 2 ? [40, 40] : [12, 12]
    },
    async getMapContents() {
      const postfix = ["", "_1", "_3"]
      let contents  = ""
      for (let p of postfix) {
        try {
          const r = await axios.get(
            `/eq-asset-preview-master/assets/eq-maps/${this.zone}${p}.txt`
          )

          if (r.status === 200) {
            if (r.data.length > 0) {
              contents += r.data
            }
          }

        } catch (err) {
          // console.log("items.ts %s", err)
        }
      }
      return contents
    },
    async parseRaceIconSizes() {

      console.time("[EqZoneMap] parseRaceIconSizes");

      // parse CSS sheet to pull sizes
      let raceIconSizes = {}
      try {
        const r = await axios.get(
          `/eq-asset-preview-master/assets/sprites/race-models.css`
        )

        if (r.status === 200) {
          if (r.data.length > 0) {
            for (let line of r.data.split("\n")) {
              line               = line.replace(".", "")
              const raceClassKey = line.split(" ")[0].trim()
              const height       = line.split("height: ")[1].split(";")[0].replace("px", "").trim()
              const width        = line.split("width: ")[1].split(";")[0].replace("px", "").trim()
              // console.log(raceClassKey)
              // console.log(height)
              // console.log(width)

              raceIconSizes[raceClassKey] = [width, height]
            }
          }
        }

      } catch (err) {
        console.log("map.vue %s", err)
      }

      this.raceIconSizes = raceIconSizes

      console.timeEnd("[EqZoneMap] parseRaceIconSizes");
    },

    async loadSafeCoordinates() {
      const zone          = (await Zones.getZoneByShortName(this.zone))
      let safeCoordinates = []
      safeCoordinates.push({
          point: this.createPoint(-zone.safe_x, -zone.safe_y),
          label: `Safe Coordinates (${zone.safe_x}, ${zone.safe_y}, ${zone.safe_z}) (xyz)`,
          iconClass: 'fade-in item-6852',
          iconSize: [40, 40]
        }
      )

      this.safeCoordinateMarker = safeCoordinates
    },

    async loadTranslocatePoints() {
      const api = (new SpellsNewApi(...SpireApi.cfg()))

      try {
        const r = await api.listSpellsNews(
          (new SpireQueryBuilder())
            .where("teleport_zone", "=", this.zone)
            .get()
        )

        if (r.status === 200) {
          // used as a mechanism to stagger multiple markers on the same coordinate
          let sameCoord = {}

          let translocatePoints = []
          for (const s of r.data) {
            if (typeof sameCoord[s.effect_base_value_2 + s.effect_base_value_1] === "undefined") {
              sameCoord[s.effect_base_value_2 + s.effect_base_value_1] = 0
            }

            let sameCoordOffset = sameCoord[s.effect_base_value_2 + s.effect_base_value_1] * 1
            translocatePoints.push({
                point: this.createPoint(-s.effect_base_value_2 + sameCoordOffset, -s.effect_base_value_1 - sameCoordOffset),
                label: s.name,
                spell: s,
                iconClass: 'fade-in spell-' + s.new_icon + '-40',
                iconSize: [40, 40]
              }
            )

            sameCoord[s.effect_base_value_2 + s.effect_base_value_1]++
          }

          this.translocatePoints = translocatePoints
          this.$forceUpdate()
        }

      } catch (err) {
        console.log("map.vue %s", err)
      }
    },

    async loadMapLines() {
      console.time("[EqZoneMap] loadMapLines");

      let map        = await this.getMapContents()
      let bounds     = [0, 0, 0, 0];
      let mapLines   = []
      let mapMarkers = []
      for (let line of map.split("\n")) {
        const cols = line.replaceAll(",", "").split(/\s+/)

        // lines
        if (cols[0].trim() === "L") {
          const x  = cols[1].trim()
          const y  = cols[2].trim()
          const x2 = cols[4].trim()
          const y2 = cols[5].trim()
          const p  = [
            this.createPoint(x, y),
            this.createPoint(x2, y2),
          ]
          bounds   = [
            Math.min(bounds[0], p[0].lat, p[1].lat),
            Math.min(bounds[1], p[0].lng, p[1].lng),
            Math.max(bounds[2], p[0].lat, p[1].lat),
            Math.max(bounds[3], p[0].lng, p[1].lng),
          ];

          mapLines.push(p)
        }

        // points
        if (cols[0].trim() === "P") {
          const x     = cols[1].trim()
          const y     = cols[2].trim()
          const label = cols[8].trim()

          mapMarkers.push(
            {
              point: this.createPoint(x, y),
              label: label.replaceAll("_", " "),
            }
          )
        }
      }

      this.markers = mapMarkers
      this.lines   = mapLines
      this.bounds  = [
        [bounds[0], bounds[1]],
        [bounds[2], bounds[3]]
      ]

      this.center = [
        (bounds[0] + bounds[2]) / 2,
        (bounds[3] + bounds[1]) / 2
      ];

      this.$forceUpdate()

      console.timeEnd("[EqZoneMap] loadMapLines");
    },

    async loadDoors() {
      console.time("[EqZoneMap] loadDoors");

      const api       = (new DoorApi(...SpireApi.cfg()))
      let doorMarkers = []

      try {
        const r = await api.listDoors(
          (new SpireQueryBuilder())
            .where("zone", "=", this.zone)
            .where("version", "=", this.version)
            .get()
        )

        if (r.status === 200) {

          let doorZonePoints = []

          for (let d of r.data) {
            doorMarkers.push(
              {
                point: this.createPoint(-d.pos_x, -d.pos_y),
                label: d.name,
                iconClass: 'fade-in item-8057',
                iconSize: [40, 40]
              }
            )

            // zone teleport linked door
            if (d.dest_zone !== "NONE") {
              const z = (await Zones.getZoneLongNameByShortName(d.dest_zone))

              // console.log(z)

              if (z !== "") {
                doorZonePoints.push(
                  {
                    point: this.createPoint(-d.pos_x, -d.pos_y),
                    label: "(Door Click) Zone Point (" + z + ")",
                    destName: d.dest_zone,
                    destInstance: d.dest_instance,
                  }
                )

                // console.log(d)
              }

            }
          }

          this.doorMarkers    = doorMarkers
          this.doorZonePoints = doorZonePoints
        }

      } catch (err) {
        console.log("map.vue %s", err)
      }

      console.timeEnd("[EqZoneMap] loadDoors");
    },

    async loadMapSpawns() {
      let npcMarkers       = []
      const gridEntriesApi = (new GridEntryApi(...SpireApi.cfg()))
      try {
        console.time("[EqZoneMap] loadMapSpawns");

        // grids
        const zone = (await Zones.getZoneByShortName(this.zone))
        const r    = await gridEntriesApi.listGridEntries(
          (new SpireQueryBuilder())
            .where("zoneid", "=", zone.zoneidnumber)
            .orderBy(["gridid", "number"])
            .get()
        );
        if (r.status === 200) {
          let gridEntries = []
          for (let e of r.data) {
            if (typeof gridEntries[e.gridid] === "undefined") {
              gridEntries[e.gridid] = []
            }
            if (typeof gridEntries[e.gridid][e.number] === "undefined") {
              gridEntries[e.gridid][e.number] = []
            }

            gridEntries[e.gridid][e.number] =
              {
                x: e.x,
                y: e.y,
              }
          }
          this.pathingGridData = gridEntries
        }

        const result = await Spawn.getByZone(this.zone, this.version, true)
        if (result.length > 0) {
          for (let spawn2 of result) {
            if (spawn2.spawnentries) {
              for (let spawnentry of spawn2.spawnentries) {
                if (spawnentry.npc_type) {

                  // if (spawn.pathgrid > 0) {
                  //   console.log(spawn)
                  // }

                  // make sure we have a npc associated to spawn
                  let npcName = ""

                  const n = spawnentry.npc_type
                  npcName = n.name + (n.lastname ? ` (${n.lastname})` : '')

                  // console.log(this.raceIconSizes[this.getNpcIcon(n)])

                  npcMarkers.push(
                    {
                      point: this.createPoint(-spawn2.x, -spawn2.y),
                      label: Npcs.getCleanName(npcName),
                      npc: n,
                      grid: spawn2.pathgrid,
                      iconClass: 'fade-in ' + this.getNpcIcon(n),
                      iconSize: this.raceIconSizes[this.getNpcIcon(n)] ? this.raceIconSizes[this.getNpcIcon(n)] : [30, 100]
                    }
                  )
                }
              }
            }
          }

          this.npcMarkers = npcMarkers

          this.$forceUpdate()

          console.timeEnd("[EqZoneMap] loadMapSpawns");
        }
      } catch (err) {
        console.log("map.vue %s", err)
      }
    },

    async loadZonePoints() {
      console.time("[EqZoneMap] loadZonePoints");

      let zonePoints = []
      const zapi     = (new ZonePointApi(...SpireApi.cfg()))
      zapi.listZonePoints(
        (new SpireQueryBuilder())
          .where("zone", "=", this.zone)
          .get()
      ).then(async (r) => {
        if (r.status === 200) {
          console.log(r.data)
          for (let point of r.data) {
            const z = (await Zones.getZoneById(point.target_zone_id))

            zonePoints.push({
                point: this.createPoint(-point.x, -point.y),
                label: "Zone Point to: " + z.long_name,
                zone: z,
              }
            )
          }

          this.zonelineMarkers = zonePoints

          this.$forceUpdate()

          // console.log(this.zonelineMarkers)
          console.timeEnd("[EqZoneMap] loadZonePoints");
        }
      })
    },

    async loadMap() {
      // reset
      this.markers              = null
      this.lines                = null
      this.npcMarkers           = null
      this.doorMarkers          = null
      this.safeCoordinateMarker = null
      this.zonelineMarkers      = null
      this.translocatePoints    = null
      this.lines                = []
      this.pathingGridLines     = []
      this.pathingGridMarkers   = null
      this.pathingGridData      = []

      // load
      await this.parseRaceIconSizes()
      this.loadMapLines()
      this.loadMapSpawns()
      this.loadDoors()
      this.loadZonePoints()
      this.loadTranslocatePoints()
      this.loadSafeCoordinates()

      this.$forceUpdate()
    }

  },
  async mounted() {
    this.loadMap()
  },
  beforeDestroy() {
    EventBus.$off("NPC_ZOOM", this.handleNpcZoomEvent);
  },
  created() {
    this.zonelineMarkers      = null
    this.doorZonePoints       = null
    this.translocatePoints    = null
    this.npcMarkers           = null
    this.doorMarkers          = null
    this.safeCoordinateMarker = null
    this.pathingGridData      = []
    this.pathingGridLines     = null
    this.pathingGridMarkers   = []
    this.lines                = []

    EventBus.$on("NPC_ZOOM", this.handleNpcZoomEvent);
  },
  data() {
    return {
      zoom: 0,
      center: null,

      hoveredNpc: "",

      zoomedNpcId: 0,
      starterZoomLevel: -100,

      zoomLevel: 0,

      bounds: null,
      crs: L.CRS.Simple,

      icon: L.icon({
        iconUrl: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII=',
        // iconSize: [32, 37],
        // iconAnchor: [16, 37],
        className: this.zoomLevel >= 2 ? 'item-4472' : 'item-4472-sm'
      }),

      map: "",

      markers: null,

      raceIconSizes: {}
    };
  }
}
</script>

<style>
.leaflet-tooltip {
  background-color: transparent;
  border: none;
  -webkit-box-shadow: none;
  box-shadow: none;
}

</style>
