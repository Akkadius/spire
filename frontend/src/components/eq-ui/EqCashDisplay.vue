<template>
  <div class="d-inline-block">
    <div v-if="platinum > 0" class="d-inline-block cash-line">
      {{ platinum.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") }}
      <div :class="'item-644-sm cash-icon'" title="Platinum"/>
    </div>

    <div v-if="gold > 0" class="d-inline-block cash-line">
      {{ gold }}
      <div :class="'item-645-sm cash-icon'" title="Gold"/>
    </div>

    <div v-if="silver > 0" class="d-inline-block cash-line">
      {{ silver }}
      <div :class="'item-646-sm cash-icon'" title="Silver"/>
    </div>

    <div v-if="copper > 0" class="d-inline-block cash-line">
      {{ copper }}

      <div :class="'item-647-sm cash-icon'" title="Copper"/>
    </div>

    <div v-if="hasNoCost()" class="d-inline-block cash-line">
      0
      <div :class="'item-644-sm cash-icon'" title="Platinum"/>
    </div>
  </div>
</template>

<script>

export default {
  name: "EqCashDisplay",
  data() {
    return {
      platinum: 0,
      gold: 0,
      silver: 0,
      copper: 0
    }
  },
  props: {
    price: {
      type: Number,
      required: true
    },
    size: {
      type: Number,
      required: false,
      default: 15,
    },
  },
  methods: {
    hasNoCost() {
      return this.price === 0
    },
    calc: function () {
      this.value = parseInt(this.platinum + this.gold + this.silver + this.copper)
    },
    init: function () {
      let initialValue = this.price.toString()
      this.copper      = 0;
      this.silver      = 0;
      this.gold        = 0;
      this.platinum    = 0;

      const copper   = initialValue.charAt(initialValue.length - 1)
      const silver   = initialValue.charAt(initialValue.length - 2)
      const gold     = initialValue.charAt(initialValue.length - 3)
      const platinum = initialValue.substr(0, initialValue.length - 3)

      if (copper > 0) {
        this.copper = copper
      }
      if (silver > 0) {
        this.silver = silver
      }
      if (gold > 0) {
        this.gold = gold
      }
      if (platinum > 0) {
        this.platinum = platinum
      }
    }
  },
  watch: {
    value: function () {
      this.init()
    },
    price: {
      deep: true,
      handler() {
        this.init()
      }
    }
  },

  mounted() {
    this.init()
  }
}
</script>

<style>
.cash-line {
  margin-right: 5px;
  margin-left: 0px;
}
.cash-icon {
  position: relative;
  top: 2px;
}
</style>
