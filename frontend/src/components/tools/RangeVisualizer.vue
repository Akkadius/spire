<template>
  <div
    :style="'position: relative;'"
  >
    <img
      :src="getImageFromMax()"
      class="range-visualizer"
      style="width: 100%; "
    >
    <div class="unit-label" :style="'left: ' + unitsToPositionText(unitMarkerComputed) + '%'">{{ unitMarkerComputed }} Units</div>
    <div class="rv-vertical-line" :style="'left: ' + unitsToPositionLine(unitMarkerComputed) + '%'"></div>

    <div v-for="tick in unitTicks">
      <div class="unit-label tick" :style="'left: ' + (tick + 1) + '%'">{{ percentToUnits(tick) }}</div>
      <div class="rv-vertical-line-tick" :style="'left: ' + tick + '%'" v-if="tick > 0"></div>
    </div>

  </div>
</template>

<script>
export default {
  name: "RangeVisualizer",
  data() {
    return {
      unitTicks: [0, 10, 20, 30, 40, 50, 60, 70, 80, 90],
    }
  },
  computed: {
    unitMarkerComputed() {
      return this.unitMarker > 1000 ? 1000 : this.unitMarker
    }
  },
  props: {
    unitMarker: {
      required: true,
      type: Number,
    },
  },
  methods: {
    getImageFromMax() {
      if (this.getCurrentRangeMax() === 1000) {
        return require('@/assets/img/range-visualizer/range-1000.png')
      }
      if (this.getCurrentRangeMax() === 250) {
        return require('@/assets/img/range-visualizer/range-250.png')
      }

      return require('@/assets/img/range-visualizer/range-50.png')
    },

    getCurrentRangeMax() {
      let max = 1000
      if (parseInt(this.unitMarker) <= 250) {
        max = 250
      }
      if (parseInt(this.unitMarker) <= 50) {
        max = 50
      }

      return parseInt(max)
    },

    percentToUnits(percent) {
      return Math.round(this.getCurrentRangeMax() * (percent / 100))
    },

    unitsToPositionText(units) {
      return this.unitsToPosition(units) + 1 > 88 ? 88 : this.unitsToPosition(units) + 1
    },
    unitsToPositionLine(units) {
      return this.unitsToPosition(units)
    },

    unitsToPosition(units) {
      return parseInt(Math.round(units / this.getCurrentRangeMax() * 100))
    }
  }
}
</script>

<style>
.range-visualizer {
  border-radius: 5px;
  border: 1px solid black;
}

.unit-label {
  position: absolute;
  display: block;
  top: 50%;
  color: rgba(255, 240, 0, 1);
  left: 11%;
  font-size: 26px;
  font-weight: bold;
  text-shadow: 1px 3px 1px black;
  z-index: 9999;
}

.tick {
  top: 10%;
  font-size: 22px;
  color: white;
}

.rv-vertical-line {
  border-left: 2px solid rgba(255, 240, 0, .8);
  background-color: blue;
  transform: translateX(-50%);
  height: 99%;
  position: absolute;
  display: block;
  top: 0;
  left: 10%;
  z-index: 9999;
}

.rv-vertical-line-tick {
  border-left: 2px solid rgba(255, 255, 255, 1);
  background-color: white;
  transform: translateX(-50%);
  height: 98%;
  position: absolute;
  display: block;
  top: 1%;
  left: 10%;
}
</style>
