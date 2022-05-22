<template>
  <content-area>

    <eq-window v-if="!dataLoaded || renderingMap" class="text-center justify-content-center">
      <div class="mb-3">
        {{ renderingMap ? 'Rendering map...' : 'Loading map...' }}
      </div>
      <loader-fake-progress v-if="!dataLoaded && !renderingMap"/>
      <eq-progress-bar :percent="100" v-if="renderingMap"/>
    </eq-window>

    <div class="card" v-if="dataLoaded && !renderingMap">
      <l-map
        v-if="center"
        :crs="crs"
        style="height: 98vh"
        :center="center"
        :bounds="bounds"
        :min-zoom="-5"
        :zoom="zoom"
        @update:zoom="zoomUpdate"
      >
        <l-polyline
          v-if="lines"
          :lat-lngs="lines"
          color="gray"
          :weight="1"
        />

        <l-marker
          v-for="(marker, index) in markers"
          :key="index"
          :lat-lng="marker.point"
          v-if="markers && markers.length > 0"
        >
          <l-tooltip>
            <eq-window>{{ marker.label }}
            </eq-window>
          </l-tooltip>
          <!--          </l-icon>-->
        </l-marker>

        <l-marker
          v-for="(marker, index) in npcMarkers"
          :key="index + '-' + marker.npc.id"
          :lat-lng="marker.point"
          v-if="npcMarkers && npcMarkers.length > 0"
        >

          <l-tooltip :options="{opacity: 1, direction: 'auto', keepView: true}">
            <eq-window style="width:600px">
              <eq-npc-card-preview
                :npc="marker.npc"
              />
            </eq-window>
          </l-tooltip>

          <!--          <l-icon-->
          <!--            icon-url="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="-->
          <!--            :class-name="marker.iconClass + '-sm'"-->
          <!--            :iconSize="calcSmallIcons(marker.iconSize)"-->
          <!--          >-->
          <!--          </l-icon>-->

          <l-icon
            icon-url="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
            :class-name="(zoomLevel >= 1) ? marker.iconClass : marker.iconClass + '-sm'"
            :iconSize="(zoomLevel >= 1) ? marker.iconSize : calcSmallIcons(marker.iconSize)"
          >
          </l-icon>

        </l-marker>
      </l-map>
    </div>
  </content-area>
</template>

<script>
import {LIcon, LMap, LMarker, LPolyline, LPopup, LTileLayer, LTooltip} from 'vue2-leaflet';
import ContentArea                                                     from "../components/layout/ContentArea";
import * as L                                                          from "leaflet";
import axios                                                           from "axios";
import {Spawn2Api}                                                     from "../app/api";
import {SpireApiClient}                                                from "../app/api/spire-api-client";
import {SpireQueryBuilder}                                             from "../app/api/spire-query-builder";
import EqNpcCardPreview                                                from "../components/eq-ui/EQNpcCardPreview";
import EqWindow                                                        from "../components/eq-ui/EQWindow";
import LoaderFakeProgress                                              from "../components/LoaderFakeProgress";
import EqProgressBar                                                   from "../components/eq-ui/EQProgressBar";

export default {
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
  methods: {
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
            `/eq-asset-preview-master/assets/eq-maps/${this.$route.query.zone}${p}.txt`
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
    }
  },
  async mounted() {
    this.dataLoaded = false
    this.renderingMap = false

    await this.parseRaceIconSizes()

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

    let npcMarkers = []
    const api      = (new Spawn2Api(SpireApiClient.getOpenApiConfig()))
    try {
      // @ts-ignore
      const result = await api.listSpawn2s(
        (new SpireQueryBuilder())
          .where("zone", "=", this.$route.query.zone)
          .includes(
            [
              "Spawnentries.NpcType",
              "Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
              "Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList",
              "Spawnentries.NpcType.NpcFactions",
              "Spawnentries.NpcType.NpcEmotes",
              "Spawnentries.NpcType.Merchantlists.Items",
              "Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item"
            ]
          )
          .get()
      )
      if (result.status === 200 && result.data) {
        setTimeout(() => {
          this.dataLoaded = true
          this.$forceUpdate()
        }, 1)

        for (let spawn of result.data) {
          // make sure we have a npc associated to spawn
          let npcName = ""
          if (
            spawn.spawnentries
            && spawn.spawnentries[0]
            && spawn.spawnentries[0].npc_type
          ) {
            const n = spawn.spawnentries[0].npc_type

            npcName = n.name + (n.lastname ? ` (${n.lastname})` : '')

            // console.log(this.raceIconSizes[this.getNpcIcon(n)])

            npcMarkers.push(
              {
                point: this.createPoint(-spawn.x, -spawn.y),
                label: npcName.replaceAll("_", " "),
                npc: n,
                iconClass: 'fade-in ' + this.getNpcIcon(n),
                iconSize: this.raceIconSizes[this.getNpcIcon(n)] ? this.raceIconSizes[this.getNpcIcon(n)] : [30, 100]
              }
            )

          }
        }

        this.npcMarkers = npcMarkers
        this.markers    = mapMarkers
        this.lines      = mapLines
        this.bounds     = [
          [bounds[0], bounds[1]],
          [bounds[2], bounds[3]]
        ]

        this.center = [
          (bounds[0] + bounds[2]) / 2,
          (bounds[3] + bounds[1]) / 2
        ];
      }
    } catch (err) {
      console.log("map.vue %s", err)
    }

    this.renderingMap = true
    setTimeout(() => {
      this.renderingMap = false
    }, 500)
  },
  created() {
    this.npcMarkers = null
  },
  data() {
    return {
      dataLoaded: false,
      renderingMap: false,

      zoom: 0,
      center: null,

      zoomLevel: 0,

      bounds: null,
      lines: [],
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
