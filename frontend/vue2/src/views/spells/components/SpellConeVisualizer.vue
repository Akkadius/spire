<template>
  <div class="cone-container">
    <div class="inner">
      <h6
        class="eq-header mt-5"
        style="position: absolute; top: 18%; left: 50%; font-size: 300%; z-index: 9999999999; margin-left: -33px;"
      >Front</h6>

      <h6
        class="eq-header mt-5"
        style="position: absolute; bottom: 21%; left: 50%; font-size: 300%; z-index: 9999999999; margin-left: -28.5px;"
      >Back</h6>

      <svg style="position: absolute; top: 50%; left: 50%; z-index: 9999; overflow: visible; margin-left: -150px; margin-top: -150px">
        <path
          :d="arc"
          fill="none"
          stroke="limegreen"
          stroke-opacity=".8"
          stroke-width="150"
        />
      </svg>

      <svg style="position: absolute; top: 50%; left: 50%; overflow: visible; margin-left: -150px; margin-top: -150px">
        <path
          :d="arc2"
          fill="none"
          stroke="white"
          stroke-opacity="1"
          stroke-width="20"
        />
      </svg>
    </div>
  </div>
</template>

<script>
export default {
  name: "SpellConeVisualizer",
  data() {
    return {
      arc: "M 149.9982546707481 50.00000001523087 A 100 100 0 1 0 150 50",
      arc2: "M 149.9982546707481 50.00000001523087 A 100 100 0 1 0 150 50"
    }
  },
  props: {
    coneStartAngle: {
      default: 0,
      required: true
    },
    coneStopAngle: {
      default: 0,
      required: true
    }
  },
  watch: {
    coneStartAngle: function () {
      this.calc()
    },
    coneStopAngle: function () {
      this.calc()
    },
  },
  mounted() {
    setTimeout(() => {
      this.calc()
    }, 100)
  },
  methods: {
    calc() {
      let endAngle = this.coneStartAngle + (this.coneStopAngle - this.coneStartAngle);
      if (this.coneStopAngle < this.coneStartAngle) {
        endAngle = 360 + this.coneStopAngle
      }

      this.arc  = this.describeArc(150, 150, 100, this.coneStartAngle, endAngle)
      this.arc2 = this.describeArc(150, 150, 100, 0, 359.999)

      console.log("[SpellConeVisualizer] coneStopAngle [%s] coneStartAngle [%s] endAngle [%s]", this.coneStopAngle, this.coneStartAngle, endAngle)
    },

    polarToCartesian(centerX, centerY, radius, angleInDegrees) {
      const angleInRadians = (angleInDegrees - 90) * Math.PI / 180.0;

      return {
        x: centerX + (radius * Math.cos(angleInRadians)),
        y: centerY + (radius * Math.sin(angleInRadians))
      };
    },
    describeArc(x, y, radius, startAngle, endAngle) {
      const start        = this.polarToCartesian(x, y, radius, endAngle);
      const end          = this.polarToCartesian(x, y, radius, startAngle);
      const largeArcFlag = endAngle - startAngle <= 180 ? "0" : "1";

      return [
        "M", start.x, start.y,
        "A", radius, radius, 0, largeArcFlag, 0, end.x, end.y
      ].join(" ");
    }

  }
}
</script>

<style>
.cone-container {
  position: relative;
  width: 500px;
  height: 600px;
  padding: 10px;
  top: 50%;
  left: 50%;
  margin-top: -100px;
  margin-left: -250px;
}

.inner {
  display: inline-block;
  margin: 0 auto;
}
</style>


