<template>
  <div v-if="ability && abilityParams" class="minified-inputs" id="special-abilities">
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top pr-3">Summon</td>
        <td style="width:210px">
          <select
            class="ability_check form-control" v-model="ability[1]" style="width:200px"
            @change="calculateSpecialAbilities"
          >
            <option value="0" selected="">Off</option>
            <option value="1">Summon target to NPC</option>
            <option value="2">Summon NPC to target</option>
          </select></td>
        <td style="width:230px">
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            style="width:230px" v-model="abilityParams[1][0]" value=""
            v-b-tooltip.hover title="Cooldown in ms (default: 6000)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[1][1]" v-b-tooltip.hover title="HP % before summon (default: 97)"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Enrage</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[2]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[2][0]" placeholder="0"
            v-b-tooltip.hover title="HP % to Enrage (rule NPC:StartEnrageValue)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[2][1]" placeholder="10000" v-b-tooltip.hover
            title="Duration (ms) (10000)"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[2][2]" placeholder="360000" v-b-tooltip.hover
            title="Cooldown (ms) (360000)"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Rampage</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[3]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][0]" placeholder="20" v-b-tooltip.hover title="Proc chance"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][1]" placeholder="Combat:MaxRampageTargets"
            v-b-tooltip.hover title="Rampage target count (default: rule Combat:MaxRampageTargets)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][2]" placeholder="0" v-b-tooltip.hover title="Flat damage to add"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][3]" placeholder="0"
            v-b-tooltip.hover title="Ignore % armor for this attack (0) "
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][4]" placeholder="0"
            v-b-tooltip.hover title="Ignore flat armor for this attack (0)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][5]" placeholder="100" v-b-tooltip.hover
            title="% NPC Crit against (100)"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[3][6]" placeholder="0"
            v-b-tooltip.hover title="Flat crit bonus on top of npc's natual crit that can go toward this attack"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">AE Rampage</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[4]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[4][0]" placeholder="0" v-b-tooltip.hover
            title="Rampage target count (1)"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[4][1]" placeholder="100"
            v-b-tooltip.hover title="% of normal attack damage (100)"
          ></td>
        <td>
          <input
            type="text"
            @change="calculateSpecialAbilities"
            class="ability_check_sub form-control"
            v-model="abilityParams[4][2]"
            value=""
            placeholder="0"
            v-b-tooltip.hover
            title="Flat damage bonus to add (0)"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[4][3]" placeholder="0"
            v-b-tooltip.hover title="Ignore % armor for this attack (0) "
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[4][4]" placeholder="0"
            v-b-tooltip.hover title="Ignore flat armor for this attack (0)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[4][5]" placeholder="100" v-b-tooltip.hover
            title="% NPC Crit against (100)"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[4][6]" placeholder="0"
            v-b-tooltip.hover title="Flat crit bonus on top of npc's natual crit that can go toward this attack"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Flurry</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[5]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][0]" placeholder="Combat:MaxFlurryHits"
            v-b-tooltip.hover title="Flurry attack count"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][1]" placeholder="100"
            v-b-tooltip.hover title="Percent of a normal attack damage to deal"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][2]" placeholder="0" v-b-tooltip.hover title="Flat damage bonus"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][3]" placeholder="0"
            v-b-tooltip.hover title="Ignore % armor for this attack (0) "
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][4]" placeholder="0"
            v-b-tooltip.hover title="Ignore flat armor for this attack (0)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][5]" placeholder="100"
            v-b-tooltip.hover title="% NPC Crit against attack (100)"
          ></td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[5][6]" placeholder="0"
            v-b-tooltip.hover title="Flat crit bonus on top of npc's natual crit that can go toward this attack"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Ranged Attack</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[11]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[11][0]" placeholder="0" v-b-tooltip.hover title="Number of Attacks"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[11][1]" placeholder="250" v-b-tooltip.hover
            title="Max Range (default: 250)"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[11][2]" placeholder="0" v-b-tooltip.hover
            title="Percent Hit Chance Modifier"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[11][3]" placeholder="0" v-b-tooltip.hover
            title="Percent Damage Modifier"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[11][4]" placeholder="25"
            v-b-tooltip.hover title="Min Range (default: RuleI(Combat, MinRangedAttackDist) = 25)"
          ></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Tunnel Vision</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[29]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[29][0]" placeholder="75" v-b-tooltip.hover
            title="Aggro modifier on non-tanks"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Leashed</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[32]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[32][0]" placeholder="0" v-b-tooltip.hover title="Range"
          ></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Tethered</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[33]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[33][0]" placeholder="0" v-b-tooltip.hover title="Aggo Range"
          ></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Flee Percent</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[37]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[37][0]" placeholder="0" v-b-tooltip.hover
            title="Percent NPC will flee at"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[37][1]" placeholder="0" v-b-tooltip.hover
            title="Percent chance to flee"
          ></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Chase Distance</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[40]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[40][0]" placeholder="0" v-b-tooltip.hover title="Max Chase Distance"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[40][1]" placeholder="0" v-b-tooltip.hover title="Min Chase Distance"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[40][2]" placeholder="0"
            v-b-tooltip.hover title="Ignore line of sight check for chasing"
          ></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Allow Tank</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[41]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[41][0]" placeholder="1"
            v-b-tooltip.hover
            title="Allows an NPC the opportunity to take aggro over a client if in melee range"
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Casting Resist Diff</td>
        <td style="width:50px; text-align:center">
          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[43]"
            @input="calculateSpecialAbilities"
          />
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[43][0]" placeholder="0"
            v-b-tooltip.hover
            title="Set an innate resist different to be applied to all spells cast by this NPC (stacks with a spells regular resist difference)."
          >
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Counter Avoid Damage</td>
        <td style="width:50px; text-align:center">

          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[44]"
            @input="calculateSpecialAbilities"
          />

        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[44][0]" placeholder="0"
            v-b-tooltip.hover title="% Reduction to avoid melee via Riposte, Parry, Block, and Dodge"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[44][1]" placeholder="0" v-b-tooltip.hover
            title="% Reduction to Riposte"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[44][2]" placeholder="0" v-b-tooltip.hover
            title="% Reduction to Block"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[44][3]" placeholder="0" v-b-tooltip.hover
            title="% Reduction to Parry"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[44][4]" placeholder="0" v-b-tooltip.hover
            title="% Reduction to Dodge"
          >
        </td>
      </tr>
      <tr>
        <td class="ability-label-top">Modify Avoid Damage</td>
        <td style="width:50px; text-align:center">

          <eq-checkbox
            class="d-inline-block"
            v-model.number="ability[51]"
            @input="calculateSpecialAbilities"
          />

        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[51][0]" placeholder="0"
            v-b-tooltip.hover title="% Addition to avoid melee via Riposte, Parry, Block, and Dodge"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[51][1]"placeholder="0" v-b-tooltip.hover
            title="% Addition to Riposte"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[51][2]" placeholder="0" v-b-tooltip.hover
            title="% Addition to Parry"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[51][3]" placeholder="0" v-b-tooltip.hover
            title="% Addition to Block"
          >
        </td>
        <td>
          <input
            type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
            v-model="abilityParams[51][4]" placeholder="0" v-b-tooltip.hover
            title="% Addition to Dodge"
          >
        </td>
      </tr>
      </tbody>
    </table>

    <div class="row mt-3">
      <div class="col-6" v-for="c in checkboxAbilities">
        <eq-checkbox
          class="mb-2 d-inline-block mr-3"
          v-model.number="ability[c.ability]"
          @input="calculateSpecialAbilities"
        />
        <div class="d-inline-block" style="user-select: none">{{ c.name }}</div>

      </div>

    </div>

    <div v-if="showSpecialAbilitiesResult" class="mt-3">
      <h4 class="eq-header">Abilities Result</h4>
      <input
        type="text" @change="calculateSpecialAbilities" class="form-control m-wrap span6" disabled
        v-model="specialAbilitiesResult" style="width:100% !important;"
      >

    </div>

    <pre style="max-width: 400px" v-if="debug">
Ability
{{ ability }}
Ability Params
{{ abilityParams }}</pre>
  </div>
</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow";
import EqCheckbox from "../eq-ui/EQCheckbox";

export default {
  name: "NpcSpecialAbilities",
  components: { EqCheckbox, EqWindow },
  props: {
    abilities: {
      type: String,
      required: true
    },
    showSpecialAbilitiesResult: {
      type: Boolean,
      required: false,
      default: false
    }
  },
  data() {
    return {
      specialAbilitiesResult: "",
      ability: {},
      abilityParams: null,
      debug: false,

      checkboxAbilities: [
        { name: "Triple Attack", ability: 6 },
        { name: "Quad Attack", ability: 7 },
        { name: "Dual Wield", ability: 8 },
        { name: "Bane Attack", ability: 9 },
        { name: "Magic Attack", ability: 10 },
        { name: "Unslowable", ability: 12 },
        { name: "Unmezable", ability: 13 },
        { name: "Uncharmable", ability: 14 },
        { name: "Unstunnable", ability: 15 },
        { name: "Unsnareable", ability: 16 },
        { name: "Unfearable", ability: 17 },
        { name: "Immune to Dispell", ability: 18 },
        { name: "Immune to Melee", ability: 19 },
        { name: "Immune to Magic", ability: 20 },
        { name: "Immune to Fleeing", ability: 21 },
        { name: "Immune to Non-Bane Melee", ability: 22 },
        { name: "Immune to Non-Magical Melee", ability: 23 },
        { name: "Will Not Aggro", ability: 24 },
        { name: "Immune to Aggro", ability: 25 },
        { name: "Resist Ranged Spells", ability: 26 },
        { name: "See through Feign Death", ability: 27 },
        { name: "Immune to Taunt", ability: 28 },
        { name: "Does NOT buff/heal friends", ability: 30 },
        { name: "Unpacifiable", ability: 31 },
        { name: "Destructible Object", ability: 34 },
        { name: "No Harm from Players", ability: 35 },
        { name: "Always Flee", ability: 36 },
        { name: "Allow Beneficial", ability: 38 },
        { name: "Disable Melee", ability: 39 },
        { name: "Ignore Root Aggro", ability: 42 },
        { name: "Proximity Aggro", ability: 45 },
        { name: "Immune to Ranged Attacks", ability: 46 },
        { name: "Immune to Client Damage", ability: 47 },
        { name: "Immune to NPC Damage", ability: 48 },
        { name: "Immune to Client Aggro", ability: 49 },
        { name: "Immune to NPC Aggro", ability: 50 },
        { name: "Immune to Memory Fades", ability: 52 },
        { name: "Immune to Open", ability: 53 },
      ]
    }
  },
  mounted() {
    this.drawValues()

    // run calculator to update result
    this.calculateSpecialAbilities()
  },
  watch: {
    abilities: {
      deep: true,
      handler() {
        this.drawValues()
        this.calculateSpecialAbilities()
      }
    },
  },

  activated() {
    this.calculateSpecialAbilities()
  },
  methods: {
    drawValues: function () {
      let abilities = {};
      let params    = {};

      for (let i = 0; i <= 200; i++) {
        params[i] = {};
      }

      for (let ability of this.abilities.split("^")) {
        // console.log(ability)
        // console.log(ability.length)

        if (ability.length === 0) {
          continue;
        }
        if (ability.split(",").length === 0) {
          continue;
        }

        const abilitySplit = ability.split(",")
        const abilityId    = abilitySplit[0].trim()
        const value        = abilitySplit[1].trim()

        if (value > 0) {
          abilities[abilityId] = parseInt(value)

          for (let i = 2; i < abilitySplit.length; i++) {
            const value = abilitySplit[i].trim()
            if (typeof params[abilityId] === "undefined") {
              params[abilityId] = {}
            }

            params[abilityId][i - 2] = parseInt(value)
          }
        }
      }

      // set form values
      this.ability       = abilities
      this.abilityParams = params
    },
    calculateSpecialAbilities: function () {
      this.$forceUpdate();

      // console.log("calculates")
      // console.log(this.ability)

      let specialAbilities = []
      for (let abilityId in this.ability) {
        const value = parseInt(this.ability[abilityId]);
        if (value > 0) {
          let abilityValues = []
          abilityValues.push(value)

          for (let abilityParamKey in this.abilityParams[abilityId]) {
            const abilityParamValue = this.abilityParams[abilityId][abilityParamKey]

            abilityValues.push(abilityParamValue)
          }

          specialAbilities.push(abilityId + "," + abilityValues.join(","))
        }
      }

      this.specialAbilitiesResult = specialAbilities.join("^")

      this.$emit("update:inputData", this.specialAbilitiesResult);
    }
  }
}
</script>

<style scoped>
.ability-label {
  text-align: right;
  padding-right: 10px;
}

.ability-label-top {
  width: 110px;
  text-align: right;
  padding-bottom: 5px;
  user-select: none;
}

#special-abilities input, #special-abilities select {
  margin-top: 0;
  margin-bottom: 0;
}
</style>
