<template>
  <div>
    <eq-window-simple title="Explore Proximity Selector" class="mb-3">
      <b-input-group prepend="Search" class="mt-3">
        <b-input
          v-model="exploreSearch"
          id="search-selector"
          class="form-control"
          v-on:keyup="searchExplore"
          placeholder="Search by zone name, id..."
        />
        <b-input-group-append>
          <b-button
            class="btn-dark btn-sm btn-outline-warning ml-1"
            @click="queueCreateProximity"
          >
            <i class="fa fa-plus"></i>
          </b-button>
        </b-input-group-append>
      </b-input-group>
    </eq-window-simple>

    <eq-window-simple
      id="explore-container"
      :style="'height: ' + ((queuedCreateProximity || Object.keys(selectedExploreEntity).length > 0) ? 45 : 82) + 'vh; overflow-y: scroll; overflow-x: hidden'" class="p-0 mt-0"
    >
      <table
        id="explore-table"
        class="eq-table eq-highlight-rows"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th></th>
          <th style="width: 30px">ID</th>
          <th>Zone</th>
          <th>Min X</th>
          <th>Max X</th>
          <th>Min Y</th>
          <th>Max Y</th>
          <th>Min Z</th>
          <th>Max Z</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'proximity-' + e.exploreid"
          :class="(isZoneSelected(e) ? 'pulsate-highlight-white' : '')"
          v-for="(e, index) in filteredExplore"
          :key="e.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              @click="selectProximity(e)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
            <b-button
              class="btn-dark btn-sm btn-outline-danger ml-1"
              @click="deleteProximity(e)"
            >
              <i class="fa fa-trash"></i>
            </b-button>
          </td>
          <td style="text-align: center" class="p-0">{{ e.exploreid }}</td>
          <td><span v-if="zones[e.zoneid]">{{ zones[e.zoneid].long_name }} ({{ e.zoneid }})</span></td>
          <td>{{ e.minx }}</td>
          <td>{{ e.maxx }}</td>
          <td>{{ e.miny }}</td>
          <td>{{ e.maxy }}</td>
          <td>{{ e.minz }}</td>
          <td>{{ e.maxz }}</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>

    <!-- Edit -->
    <eq-window-simple
      title="Edit"
      class="minified-inputs"
      v-if="selectedExplore > 0 && selectedExploreEntity && Object.keys(selectedExploreEntity).length > 0 && !queuedCreateProximity"
    >
      <div class="row">
        <div class="font-weight-bold mb-3 col-12 p-0 text-center">
          Edits made to this explore object are saved realtime
        </div>
      </div>

      <div class="row">
        <div
          v-for="field in
         [
           {
             description: 'Min X',
             field: 'minx',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Max X',
             field: 'maxx',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Min Y',
             field: 'miny',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Max Y',
             field: 'maxy',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Min Z',
             field: 'minz',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Max Z',
             field: 'maxz',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
         ]"
          :class="field.col + ' mb-3 pl-2 pr-2'"
          v-if="(typeof field.showIf !== 'undefined' && field.showIf) || typeof field.showIf === 'undefined'"
        >
          <div>
          <span
            v-if="field.itemIcon"
            :class="'item-' + field.itemIcon + '-sm'"
            style="display: inline-block"
          />
            {{ field.description }}
          </div>

          <!-- input number -->
          <b-form-input
            v-if="field.fieldType === 'number'"
            :id="field.field"
            v-model.number="selectedExploreEntity[field.field]"
            class="m-0 mt-1"
            @change="saveProximity()"
            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
            :style="(selectedExploreEntity[field.field] === 0 ? 'opacity: .5' : '')"
          />

          <!-- input text -->
          <b-form-input
            v-if="field.fieldType === 'text' || !field.fieldType"
            :id="field.field"
            v-model="selectedExploreEntity[field.field]"
            class="m-0 mt-1"
            @change="saveProximity()"
            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
            :style="(selectedExploreEntity[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
          />

          <!-- select -->
          <select
            v-model.number="selectedExploreEntity[field.field]"
            class="form-control m-0 mt-1"
            v-if="field.selectData"
            @change="saveProximity()"
            :style="(selectedExploreEntity[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
          >
            <option
              v-for="(description, index) in field.selectData"
              :key="index"
              :value="parseInt(index)"
            >
              {{ index }}) {{ description }}
            </option>
          </select>
        </div>
      </div>

      <div class="row">
        <div class="font-weight-bold col-12 p-0 text-center fade-in" v-if="notification">
          Saved!
        </div>
      </div>

      <eq-debug :data="selectedExploreEntity"></eq-debug>

    </eq-window-simple>

    <!-- Create -->
    <eq-window-simple
      title="Create"
      class="minified-inputs"
      v-if="queuedCreateProximity"
    >
      <div class="row">
        <div
          v-for="field in
         [
           {
             description: 'Zone ID',
             field: 'zoneid',
             fieldType: 'select',
             itemIcon: '2149',
             col: 'col-12',
             zeroValue: -1,
             selectData: zoneSelect
           },
           {
             description: 'Min X',
             field: 'minx',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Max X',
             field: 'maxx',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Min Y',
             field: 'miny',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Max Y',
             field: 'maxy',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Min Z',
             field: 'minz',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
           {
             description: 'Max Z',
             field: 'maxz',
             fieldType: 'number',
             itemIcon: '2149',
             col: 'col-6',
             zeroValue: -1
           },
         ]"
          :class="field.col + ' mb-3 pl-2 pr-2'"
          v-if="(typeof field.showIf !== 'undefined' && field.showIf) || typeof field.showIf === 'undefined'"
        >
          <div>
          <span
            v-if="field.itemIcon"
            :class="'item-' + field.itemIcon + '-sm'"
            style="display: inline-block"
          />
            {{ field.description }}
          </div>

          <!-- input number -->
          <b-form-input
            v-if="field.fieldType === 'number'"
            :id="field.field"
            v-model.number="newExploreEntity[field.field]"
            class="m-0 mt-1"
            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
            :style="(newExploreEntity[field.field] === 0 ? 'opacity: .5' : '')"
          />

          <!-- input text -->
          <b-form-input
            v-if="field.fieldType === 'text' || !field.fieldType"
            :id="field.field"
            v-model="newExploreEntity[field.field]"
            class="m-0 mt-1"
            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
            :style="(newExploreEntity[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
          />

          <!-- select -->
          <select
            v-model.number="newExploreEntity[field.field]"
            class="form-control m-0 mt-1"
            v-if="field.selectData"
            :style="(newExploreEntity[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
          >
            <option
              v-for="(description, index) in field.selectData"
              :key="index"
              :value="parseInt(index)"
            >
              {{ index }}) {{ description }}
            </option>
          </select>
        </div>
      </div>

      <div class="row">
        <div class="col-12 text-center">
          <b-button
            @click="createProximity()"
            size="sm"
            variant="outline-warning"
          >
            <i class="fa fa-save mr-1"></i>
            Save
          </b-button>
        </div>
      </div>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple      from "@/components/eq-ui/EQWindowSimple";
import {ProximityApi}      from "@/app/api";
import {SpireApiClient}    from "@/app/api/spire-api-client";
import util                from "util";
import Expansions          from "@/app/utility/expansions";
import EqCheckbox          from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {Zones}             from "@/app/zones";
import EqDebug             from "@/components/eq-ui/EQDebug";

let proximities = {}

export default {
  name: "TaskExploreSelector",
  components: { EqDebug, EqCheckbox, EqWindowSimple },
  data() {
    return {
      // filtered content
      filteredExplore: {},

      // search
      exploreSearch: "",

      // model we work with after the prop is passed so we can manipulate it ourselves
      // props should not be mutated
      selectedExplore: 0,

      // currently selected entity
      selectedExploreEntity: {},

      // zone cache
      zones: {},
      zoneSelect: {},

      // notification
      notification: "",

      // create
      queuedCreateProximity: false,
      newExploreEntity: {},
    }
  },
  props: {
    selectedExploreId: {
      type: Number,
      required: true,
    },
  },
  methods: {

    queueCreateProximity() {
      this.queuedCreateProximity = true

      this.newExploreEntity = {
        "zoneid": 3,
        "exploreid": 0,
        "minx": -900,
        "maxx": -850,
        "miny": -240,
        "maxy": -200,
        "minz": 0,
        "maxz": 3
      }

      let nextExploreId = 0
      proximities.forEach((p) => {
        nextExploreId = p.exploreid + 1
      })

      this.newExploreEntity.exploreid = nextExploreId
    },

    async createProximity() {
      try {
        const api = (new ProximityApi(SpireApiClient.getOpenApiConfig()))
        const r   = await api.createProximity(
          {
            proximity: this.newExploreEntity,
          }
        )
        if (r.status === 200) {
          this.selectedExplore       = this.newExploreEntity.exploreid
          this.queuedCreateProximity = false
          this.init()
        }
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.sendNotification("Error! " + err.response.data.error)
        }
      }
    },

    async deleteProximity(e) {
      if (confirm(`Are you sure you want to delete this proximity?\n\n Explore ID: ${e.exploreid}`)) {
        try {
          const api = (new ProximityApi(SpireApiClient.getOpenApiConfig()))
          const r   = await api.deleteProximity(
            {
              id: e.zoneid,
            },
            {
              query: (new SpireQueryBuilder())
                .where("exploreid", "=", e.exploreid)
                .where("zoneid", "=", e.zoneid)
                .get()
            }
          )
          if (r.status === 200) {
            this.init()
          }
        } catch (err) {
          if (err.response && err.response.data && err.response.data.error) {
            this.sendNotification("Error! " + err.response.data.error)
          }
        }
      }
    },

    async saveProximity() {
      console.log("save proximity!")

      if (this.selectedExploreEntity) {
        try {
          const api    = (new ProximityApi(SpireApiClient.getOpenApiConfig()))
          const result = await api.updateProximity(
            {
              id: this.selectedExploreEntity.zoneid,
              proximity: this.selectedExploreEntity,
            },
            {
              query: (new SpireQueryBuilder())
                .where("exploreid", "=", this.selectedExploreEntity.exploreid)
                .where("zoneid", "=", this.selectedExploreEntity.zoneid)
                .get()
            }
          )

          if (result.status === 200) {
            this.sendNotification("Saved!")
          }

        } catch (err) {
          if (err.response && err.response.data && err.response.data.error) {
            this.sendNotification("Error! " + err.response.data.error)
          }
        }
      }

    },

    isZoneSelected(proximity) {
      return proximity.exploreid === parseInt(this.selectedExplore)
    },

    selectProximity(proximity) {
      this.$emit('input', {
        id: proximity.exploreid,
      });

      this.selectedExplore       = proximity.exploreid
      this.selectedExploreEntity = proximity

      this.queuedCreateProximity = false
    },

    sendNotification(message) {
      this.notification = message
      setTimeout(() => {
        this.notification = ""
      }, 5000)
    },

    searchExplore() {
      if (!this.exploreSearch && this.exploreSearch.length === 0) {
        return
      }

      const searchString  = this.exploreSearch.toString().toLowerCase().trim()
      let filteredExplore = []
      proximities.forEach((p) => {
        if (
          (
            p.exploreid.toString().includes(searchString) ||
            p.zoneid.toString().includes(searchString) ||
            (this.zones[p.zoneid] && this.zones[p.zoneid].long_name && this.zones[p.zoneid].long_name.toLowerCase().includes(searchString))
          )) {
          filteredExplore.push(p)
        }
      });
      this.filteredExplore = filteredExplore
      if (filteredExplore.length === 0) {
        this.filteredExplore = proximities;
      }
    },

    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion - 1) // zone table is offset by 1
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion - 1) // zone table is offset by 1
    },

    async loadProximities() {
      const api    = (new ProximityApi(SpireApiClient.getOpenApiConfig()))
      const result = await api.listProximities(
        (new SpireQueryBuilder())
          .orderBy(["exploreid"])
          .get()
      )

      if (result.status === 200) {
        proximities          = result.data
        this.filteredExplore = proximities
      }
    },

    init() {
      const t = document.getElementById("search-selector")
      if (t) {
        t.focus()
      }

      this.loadProximities().then(() => {
        this.scrollToSelected()

        proximities.forEach((p) => {
          if (p.exploreid === this.selectedExplore) {
            console.log("selected explore")
            this.selectedExploreEntity = p
          }
        })
      })
    },

    scrollToSelected() {
      setTimeout(() => {
        const container = document.getElementById("explore-container");
        const target    = document.getElementById(util.format("proximity-%s", this.selectedExplore))
        if (container && target) {
          const top           = target.getBoundingClientRect().top
          container.scrollTop = container.scrollTop + top - 300;
        }
      }, 100)
    }
  },
  mounted() {
    // model we work with after the prop is passed - we can manipulate it ourselves
    if (this.selectedExploreId > 0) {
      this.selectedExplore = this.selectedExploreId
      // this.exploreSearch   = this.selectedExploreId
    }

    Zones.getZones().then((r) => {
      let zones      = {}
      let zoneSelect = {}
      r.forEach((zone) => {
        zones[zone.zoneidnumber]      = zone
        zoneSelect[zone.zoneidnumber] = zone.long_name
      })
      this.zones      = zones
      this.zoneSelect = zoneSelect

      this.init()
    })

  }
}
</script>

<style scoped>
#explore-table td {
  vertical-align: middle !important;

}

#explore-table td, #explore-table th {
  padding: 3px;
}
</style>
