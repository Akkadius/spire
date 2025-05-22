<template>
  <div style="min-height: 250px">
    <div v-if="!loaded" class="mt-3 text-center">
      Loading...
      <loader-fake-progress/>
    </div>

    <div
      class="font-weight-bold text-center mt-3 mb-3"
      v-if="!races.includes(selectedRace) && loaded"
    >
      Facial appearance editing is not supported for this race
    </div>

    <div class="row" v-if="races.includes(selectedRace) && loaded">
      <div class="col-4">
        <div
          style="width: 128px; height: 128px; position: absolute; top: 50px; right: 30px; "
        >
          <span
            :class="`faces-race-${selectedRace}-gender-${selectedGender}-face-${selectedFace}`"
            style="position: absolute; right: 0; top: 0; border: 1px solid #666"
          />
          <span
            :class="`faces-race-${selectedRace}-gender-${selectedGender}-beard-${selectedBeard}-color-${selectedBeardColor}`"
            style="position: absolute; right: 0; top: 0; z-index: 1"
          />
          <span
            :class="`faces-race-${selectedRace}-gender-${selectedGender}-hair-${selectedHair}-color-${selectedHairColor}`"
            style="position: absolute; right: 0; top: 0; z-index: 2"
          />
          <span
            :class="`faces-race-${selectedRace}-gender-${selectedGender}-eyes-${selectedEye}`"
            style="position: absolute; right: 0; top: 0; z-index: 3"
          />
        </div>
      </div>
      <div class="col-6 text-left text">
        <!--      <div class="mb-3 ml-5 font-weight-bold">-->
        <!--        Facial Feature Controls-->
        <!--      </div>-->

        <!-- Spacer -->
        <div class="mt-3"></div>

        <!-- Race -->
        <div class="row mb-3" v-if="!inRace">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Race">
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('race', 'decrement')">
                -
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Race ({{ selectedRace }})</b>
              </b-input-group-prepend>
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('race', 'increment')">
                +
              </b-button>
            </div>

            <div class="ml-3 mt-1 d-inline-block">{{ getSelectedRaceName() }}</div>

          </div>
        </div>

        <!-- Gender -->
        <div class="row mb-3" v-if="typeof inGender === 'undefined'">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Race">
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('gender', 'decrement')"
              >
                -
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Gender ({{ selectedGender }})</b>
              </b-input-group-prepend>
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('gender', 'increment')"
              >
                +
              </b-button>
            </div>

            <div class="ml-3 mt-1 d-inline-block">{{ getSelectedGenderName() }}</div>
          </div>
        </div>

        <!-- Face -->
        <div class="row mb-3">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Face">
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('face', 'decrement')">
                -
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Face ({{ selectedFace }})</b>
              </b-input-group-prepend>
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('face', 'increment')">
                +
              </b-button>
            </div>
          </div>
        </div>

        <!-- Hair -->
        <div class="row mb-3" v-if="hairs.length > 0">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Face">
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('hair', 'decrement')">
                -
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Hair ({{ selectedHair }})</b>
              </b-input-group-prepend>
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('hair', 'increment')">
                +
              </b-button>
            </div>
          </div>
        </div>

        <!-- Hair Color -->
        <div class="row mb-3" v-if="hairColors.length > 0">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Face">
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('haircolor', 'decrement')"
              >-
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Hair Color ({{ selectedHairColor }})</b>
              </b-input-group-prepend>
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('haircolor', 'increment')"
              >+
              </b-button>
            </div>
          </div>
        </div>

        <!-- Beard -->
        <div class="row mb-3" v-if="beards.length > 0">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Beard">
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('beard', 'decrement')"
              >-
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Beard ({{ selectedBeard }})</b>
              </b-input-group-prepend>
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('beard', 'increment')"
              >+
              </b-button>
            </div>
          </div>
        </div>

        <!-- Beard Color -->
        <div class="row mb-3" v-if="beardColors.length > 0">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Beard Color">
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('beardcolor', 'decrement')"
              >-
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Beard Color ({{ selectedBeardColor }})</b>
              </b-input-group-prepend>
              <b-button
                size="sm"
                variant="outline-warning"
                class="facial-ctrl-btn"
                @click="cycle('beardcolor', 'increment')"
              >+
              </b-button>
            </div>
          </div>
        </div>

        <!-- Eyes -->
        <div class="row mb-3">
          <div class="col-12">
            <div class="btn-group ml-3 facial-btn-group" role="group" aria-label="Eyes">
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('eyes', 'decrement')">
                -
              </b-button>
              <b-input-group-prepend is-text class="facial-btn-group-text"><b>Eye ({{ selectedEye }})</b>
              </b-input-group-prepend>
              <b-button size="sm" variant="outline-warning" class="facial-ctrl-btn" @click="cycle('eyes', 'increment')">
                +
              </b-button>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import axios              from "axios";
import {DB_RACE_NAMES}    from "../../app/constants/eq-races-constants";
import {GENDER}           from "../../app/constants/eq-gender-constants";
import LoaderFakeProgress from "../LoaderFakeProgress";

export default {
  name: "FacialAppearanceSelector",
  components: { LoaderFakeProgress },
  async mounted() {
    this.load()
  },
  data() {
    return {
      loaded: false,

      selectedRace: 0,
      selectedFace: 0,
      selectedHair: 0,
      selectedHairColor: 0,
      selectedGender: 0,
      selectedBeard: 0,
      selectedEye: 0,
      selectedBeardColor: 0,
    }
  },

  watch: {
    'inRace'() {
      // console.log("race changed")
      this.load()
    },
  },

  props: {
    inRace: { type: [Number, String], required: false },
    inHair: { type: [Number, String], required: false },
    inHairColor: { type: [Number, String], required: false },
    inFace: { type: [Number, String], required: false },
    inGender: { type: [Number, String], required: false },
    inBeard: { type: [Number, String], required: false },
    inEye: { type: [Number, String], required: false },
    inBeardColor: { type: [Number, String], required: false },
  },

  created() {
    // non-reactive
    this.races        = []
    this.eyes         = []
    this.beards       = []
    this.beardColors  = []
    this.faces        = []
    this.hairs        = []
    this.hairColors   = []
    this.cssSheetData = []

    // console.log("created")
    // console.log(this.inRace)

    // eyes
    // hair
    // beards
    // faces
    // genders (given 0-1)

    // collect each set of what's available and display it in picker options


  },
  methods: {

    calcOptions() {

      // reset
      this.eyes        = []
      this.beards      = []
      this.beardColors = []
      this.faces       = []
      this.hairs       = []
      this.hairColors  = []

      // loop through sheet
      for (let l of this.cssSheetData) {
        const d      = l.split("-")
        const raceId = parseInt(d[0])
        const gender = parseInt(d[2])

        if (raceId === this.selectedRace && gender === this.selectedGender) {
          // console.log(l)

          const type     = d[3]
          const value    = parseInt(d[4])
          const subType  = d[5]
          const subValue = parseInt(d[6])
          if (type === "hair") {
            if (!this.hairs.includes(value)) {
              this.hairs.push(value)
            }

            if (value === this.selectedHair && subType === "color") {
              if (!this.hairColors.includes(subValue)) {
                this.hairColors.push(subValue)
              }
            }
          }

          if (type === "beard") {
            if (!this.beards.includes(value)) {
              this.beards.push(value)
            }

            // TODO: filter color choices based off of selected
            if (subType === "color") {
              if (!this.beardColors.includes(subValue)) {
                this.beardColors.push(subValue)
              }
            }
          }

          if (type === "face") {
            if (!this.faces.includes(value)) {
              this.faces.push(value)
            }
          }

          if (type === "eyes") {
            if (!this.eyes.includes(value)) {
              this.eyes.push(value)
            }
          }

        }
      }

      // check if the selected hair is in the new set of options
      // or set to a valid option within the new set
      // (out of bounds checks)
      if (!this.hairs.includes(this.selectedHair)) {
        this.selectedHair = this.hairs[0]
      }
      if (!this.hairColors.includes(this.selectedHairColor)) {
        this.selectedHairColor = this.hairColors[0]
      }
      if (!this.faces.includes(this.selectedFace)) {
        this.selectedFace = this.faces[0]
      }
      if (!this.eyes.includes(this.selectedEye)) {
        this.selectedEye = this.eyes[0]
      }
      if (!this.beards.includes(this.selectedBeard)) {
        this.selectedBeard = this.beards[0]
      }
      if (!this.beardColors.includes(this.selectedBeardColor)) {
        this.selectedBeardColor = this.beardColors[0]
      }

      // console.log("hairs", this.hairs)
      // console.log("hairColors", this.hairColors)
      // console.log("beards", this.beards)
      // console.log("beardColors", this.beardColors)
      // console.log("faces", this.faces)

      this.$forceUpdate()

      if (this.loaded) {
        this.$emit('input',
          {
            face: this.selectedFace,
            hair: this.selectedHair,
            hairColor: this.selectedHairColor,
            gender: this.selectedGender,
            beard: this.selectedBeard,
            beardColor: this.selectedBeardColor,
            eye: this.selectedEye,
          }
        );
      }
    },

    getSelectedRaceName() {
      return DB_RACE_NAMES[this.selectedRace] ? DB_RACE_NAMES[this.selectedRace] : "N/A"
    },
    getSelectedGenderName() {
      return GENDER[this.selectedGender] ? GENDER[this.selectedGender] : "N/A"
    },

    cycle(option, direction) {
      console.log("Cycling")
      console.log("option", option)
      console.log("direction", direction)

      if (option === "race") {
        // console.log(this.races)
        for (const [i, r] of this.races.entries()) {
          if (r === this.selectedRace) {
            if (direction === "increment") {
              this.selectedRace = typeof this.races[i + 1] !== "undefined" ? this.races[i + 1] : this.races[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedRace = typeof this.races[i - 1] !== "undefined" ? this.races[i - 1] : this.races[i]
              this.calcOptions()
              break;
            }

          }
        }
      }

      if (option === "face") {
        // console.log(this.faces)
        for (const [i, r] of this.faces.entries()) {
          if (r === this.selectedFace) {
            if (direction === "increment") {
              this.selectedFace = typeof this.faces[i + 1] !== "undefined" ? this.faces[i + 1] : this.faces[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedFace = typeof this.faces[i - 1] !== "undefined" ? this.faces[i - 1] : this.faces[i]
              this.calcOptions()
              break;
            }
          }
        }
      }

      if (option === "beard") {
        for (const [i, r] of this.beards.entries()) {
          if (r === this.selectedBeard) {
            if (direction === "increment") {
              this.selectedBeard = typeof this.beards[i + 1] !== "undefined" ? this.beards[i + 1] : this.beards[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedBeard = typeof this.beards[i - 1] !== "undefined" ? this.beards[i - 1] : this.beards[i]
              this.calcOptions()
              break;
            }
          }
        }
      }

      if (option === "hair") {
        // console.log(this.faces)
        for (const [i, r] of this.hairs.entries()) {
          if (r === this.selectedHair) {
            if (direction === "increment") {
              this.selectedHair = typeof this.hairs[i + 1] !== "undefined" ? this.hairs[i + 1] : this.hairs[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedHair = typeof this.hairs[i - 1] !== "undefined" ? this.hairs[i - 1] : this.hairs[i]
              this.calcOptions()
              break;
            }
          }
        }
      }

      if (option === "eyes") {
        // console.log(this.faces)
        for (const [i, r] of this.eyes.entries()) {
          if (r === this.selectedEye) {
            if (direction === "increment") {
              this.selectedEye = typeof this.eyes[i + 1] !== "undefined" ? this.eyes[i + 1] : this.eyes[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedEye = typeof this.eyes[i - 1] !== "undefined" ? this.eyes[i - 1] : this.eyes[i]
              this.calcOptions()
              break;
            }
          }
        }
      }

      if (option === "haircolor") {
        // console.log(this.faces)
        for (const [i, r] of this.hairColors.entries()) {
          if (r === this.selectedHairColor) {
            if (direction === "increment") {
              this.selectedHairColor = typeof this.hairColors[i + 1] !== "undefined" ? this.hairColors[i + 1] : this.hairColors[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedHairColor = typeof this.hairColors[i - 1] !== "undefined" ? this.hairColors[i - 1] : this.hairColors[i]
              this.calcOptions()
              break;
            }
          }
        }
      }

      if (option === "beardcolor") {
        // console.log(this.faces)
        for (const [i, r] of this.beardColors.entries()) {
          if (r === this.selectedBeardColor) {
            if (direction === "increment") {
              this.selectedBeardColor = typeof this.beardColors[i + 1] !== "undefined" ? this.beardColors[i + 1] : this.beardColors[i]
              this.calcOptions()
              break;
            }
            if (direction === "decrement") {
              this.selectedBeardColor = typeof this.beardColors[i - 1] !== "undefined" ? this.beardColors[i - 1] : this.beardColors[i]
              this.calcOptions()
              break;
            }
          }
        }
      }

      if (option === "gender") {
        this.selectedGender = this.selectedGender === 1 ? 0 : 1
      }

      this.calcOptions()
    },
    async load() {
      try {
        const r = await axios.get(
          `/eq-asset-preview-master/assets/sprites/faces.css`
        )

        if (r.status === 200) {
          if (r.data.length > 0) {
            for (let line of r.data.split("\n")) {

              line = line.replaceAll(".faces-race-", "")
              line = line.replaceAll(" { ", "-")

              if (line.includes("background")) {
                line = line.split("-background")[0];

                this.cssSheetData.push(line)
              }
            }
          }
        }

      } catch (err) {
        console.log("map.vue %s", err)
      }

      // sort sheet alpha and numerically
      this.cssSheetData = this.cssSheetData.sort((a, b) => a.localeCompare(b, 'en', { numeric: true }))

      // loop through sheet
      for (let l of this.cssSheetData) {
        const d      = l.split("-")
        const raceId = parseInt(d[0])
        if (raceId > 0) {
          if (!this.races.includes(raceId)) {
            this.races.push(raceId)
          }
        }
        // console.log(l)
      }

      this.selectedRace = this.races[0]
      this.$forceUpdate()

      // input from props
      if (this.inRace) {
        this.selectedRace = parseInt(this.inRace)
      }
      if (this.inHair) {
        this.selectedHair = parseInt(this.inHair)
      }
      if (this.inHairColor) {
        this.selectedHairColor = parseInt(this.inHairColor)
      }
      if (this.inGender) {
        this.selectedGender = parseInt(this.inGender)
      }
      if (this.inFace) {
        this.selectedFace = parseInt(this.inFace)
      }
      if (this.inBeard) {
        this.selectedBeard = parseInt(this.inBeard)
      }
      if (this.inBeardColor) {
        this.selectedBeardColor = parseInt(this.inBeardColor)
      }

      this.calcOptions()

      this.loaded = true

      // console.log("races", this.races)
      // console.log("eyes", this.eyes)
      // console.log("hairs", this.hairs)
      // console.log("cssSheetData", this.cssSheetData)
      // console.log("hairColors", this.hairColors)
    }
  }
}
</script>

<style>
.facial-btn-group {
  width: 190px;
}

.facial-btn-group-text, .facial-btn-group-text .input-group-text, .facial-btn-group-text b {
  width: 140px;
  text-align: center;
  border-radius: 0px;
}

.facial-ctrl-btn {
  width: 25px;
  /*padding: 1px !important;*/
  flex: none !important;
  padding: 0px !important;
}
</style>
