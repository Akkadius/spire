<template>
  <div v-if="ability && abilityParams">
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Summon</td>
        <td style="width:210px">
          <select class="ability_check form-control" v-model="ability[1]" style="width:200px"
                  @change="calculateSpecialAbilities">
            <option value="0" selected="">Off</option>
            <option value="1">Summon target to NPC</option>
            <option value="2">Summon NPC to target</option>
          </select></td>
        <td style="width:230px">
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control"
                 style="width:230px" v-model="abilityParams[1][0]" value=""
                 v-b-tooltip.hover title="Cooldown in ms (default: 6000)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[1][1]" value="" v-b-tooltip.hover title="HP % before summon (default: 97)">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Enrage</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[2]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[2][0]" value="" placeholder="0"
                 v-b-tooltip.hover title="HP % to Enrage (rule NPC:StartEnrageValue)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[2][1]" value="" placeholder="10000" v-b-tooltip.hover
                 title="Duration (ms) (10000)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[2][2]" value="" placeholder="360000" v-b-tooltip.hover
                 title="Cooldown (ms) (360000)">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Rampage</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[3]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][0]" value="" placeholder="20" v-b-tooltip.hover title="Proc chance"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][1]" value="" placeholder="Combat:MaxRampageTargets"
                 v-b-tooltip.hover title="Rampage target count (default: rule Combat:MaxRampageTargets)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][2]" value="" placeholder="0" v-b-tooltip.hover title="Flat damage to add">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][3]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore % armor for this attack (0) "></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][4]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore flat armor for this attack (0)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][5]" value="" placeholder="100" v-b-tooltip.hover
                 title="% NPC Crit against (100)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[3][6]" value="" placeholder="0"
                 v-b-tooltip.hover title="Flat crit bonus on top of npc's natual crit that can go toward this attack">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">AE Rampage</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[4]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][0]" value="" placeholder="0" v-b-tooltip.hover
                 title="Rampage target count (1)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][1]" value="" placeholder="100"
                 v-b-tooltip.hover title="% of normal attack damage (100)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][2]" value="" placeholder="0" v-b-tooltip.hover
                 title="Flat damage bonus to add (0)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][3]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore % armor for this attack (0) "></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][4]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore flat armor for this attack (0)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][5]" value="" placeholder="100" v-b-tooltip.hover
                 title="% NPC Crit against (100)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[4][6]" value="" placeholder="0"
                 v-b-tooltip.hover title="Flat crit bonus on top of npc's natual crit that can go toward this attack">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Flurry</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[5]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][0]" value="" placeholder="Combat:MaxFlurryHits"
                 v-b-tooltip.hover title="Flurry attack count"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][1]" value="" placeholder="100"
                 v-b-tooltip.hover title="Percent of a normal attack damage to deal"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][2]" value="" placeholder="0" v-b-tooltip.hover title="Flat damage bonus">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][3]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore % armor for this attack (0) "></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][4]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore flat armor for this attack (0)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][5]" value="" placeholder="100"
                 v-b-tooltip.hover title="% NPC Crit against attack (100)"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[5][6]" value="" placeholder="0"
                 v-b-tooltip.hover title="Flat crit bonus on top of npc's natual crit that can go toward this attack">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Ranged Attack</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[11]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[11][0]" value="" placeholder="0" v-b-tooltip.hover title="Number of Attacks">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[11][1]" value="" placeholder="250" v-b-tooltip.hover
                 title="Max Range (default: 250)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[11][2]" value="" placeholder="0" v-b-tooltip.hover
                 title="Percent Hit Chance Modifier">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[11][3]" value="" placeholder="0" v-b-tooltip.hover
                 title="Percent Damage Modifier">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[11][4]" value="" placeholder="25"
                 v-b-tooltip.hover title="Min Range (default: RuleI(Combat, MinRangedAttackDist) = 25)"></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Tunnel Vision</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[29]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[29][0]" value="" placeholder="75" v-b-tooltip.hover
                 title="Aggro modifier on non-tanks">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Leashed</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="abilityParams[32]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[32][0]" value="" placeholder="0" v-b-tooltip.hover title="Range"></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Tethered</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[33]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[33][0]" value="" placeholder="0" v-b-tooltip.hover title="Aggo Range"></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Flee Percent</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[37]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[37][0]" value="" placeholder="0" v-b-tooltip.hover
                 title="Percent NPC will flee at">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[37][1]" value="" placeholder="0" v-b-tooltip.hover
                 title="Percent chance to flee"></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Chase Distance</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[40]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[40][0]" value="" placeholder="0" v-b-tooltip.hover title="Max Chase Distance">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[40][1]" value="" placeholder="0" v-b-tooltip.hover title="Min Chase Distance">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[40][2]" value="" placeholder="0"
                 v-b-tooltip.hover title="Ignore line of sight check for chasing"></td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Allow Tank</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[41]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[41][0]" value="" placeholder="1"
                 v-b-tooltip.hover
                 title="Allows an NPC the opportunity to take aggro over a client if in melee range">
        </td>
      </tr>
      </tbody>
    </table>
    <table class="table-condensed flip-content" style="width:100%">
      <tbody>
      <tr>
        <td class="ability-label-top">Casting Resist Diff</td>
        <td style="width:50px; text-align:center">
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[43]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[43][0]" value="" placeholder="0"
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
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[44]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[44][0]" value="" placeholder="0"
                 v-b-tooltip.hover title="chance to avoid melee via dodge/parray/riposte/block skills "></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[44][1]" value="" placeholder="0-100" v-b-tooltip.hover
                 title="Avoidance % (0-100)">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[44][2]" value="" placeholder="0" v-b-tooltip.hover
                 title="% Reduction to Riposte"></td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[44][3]" value="" placeholder="0" v-b-tooltip.hover
                 title="% Reduction to Parry">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[44][4]" value="" placeholder="0" v-b-tooltip.hover
                 title="% Reduction to Block">
        </td>
        <td>
          <input type="text" @change="calculateSpecialAbilities" class="ability_check_sub form-control" style=""
                 v-model="abilityParams[44][5]" value="" placeholder="0" v-b-tooltip.hover
                 title="% Reduction to Dodge">
        </td>
      </tr>
      </tbody>
    </table>
    <br>
    <table class="table-condensed flip-content" style="width: 100%">
      <tbody>
      <tr>
        <td class="ability-label">Triple Attack</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[6]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Quad Attack</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[7]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Dual Wield</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[8]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Bane Attack</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[9]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Magic Attack</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[10]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Unslowable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[12]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Unmezable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[13]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Uncharmable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[14]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Unstunable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[15]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Unsnareable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[16]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Unfearable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[17]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Immune to Dispell</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[18]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Immune to Melee</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[19]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Immune to Magic</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[20]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Immune to Fleeing</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[21]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Immune to non-Bane Melee</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[22]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Immune to non-Magical Melee</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[23]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Will Not Aggro</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[24]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Immune to Aggro</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[25]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Resist Ranged Spells</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[26]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">See through Feign Death</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[27]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Immune to Taunt</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[28]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Does NOT buff/heal friends</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[30]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Unpacifiable</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[31]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Destructible Object</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[34]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">No Harm from Players</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[35]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Always Flee</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[36]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      <tr>
        <td class="ability-label">Allow Beneficial</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[38]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Disable Melee</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[39]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
        <td class="ability-label">Ignore Root Aggro</td>
        <td>
          <input class="ability_check" @change="calculateSpecialAbilities" v-model="ability[42]"
                 v-bind:true-value="'1'" v-bind:false-value="'0'" type="checkbox"></td>
      </tr>
      </tbody>
    </table>
    <br>

    <div v-if="showSpecialAbilitiesResult">

      <h4 class="eq-header">Abilities Result</h4>
      <input type="text" @change="calculateSpecialAbilities" class="form-control m-wrap span6" disabled
             v-model="specialAbilitiesResult" style="width:100% !important;">

    </div>

    <pre style="max-width: 400px" v-if="debug">
Ability
{{ ability }}
Ability Params
{{ abilityParams }}</pre>
  </div>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow";

export default {
  name: "NpcSpecialAbilities",
  components: { EqWindow },
  props: {
    abilities: {
      type: String,
      required: true
    },
    showSpecialAbilitiesResult: {
      type: String,
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
      propWatcher: null
    }
  },
  mounted() {
    this.drawValues()

    // run calculator to update result
    this.calculateSpecialAbilities()
  },
  activated() {
    this.calculateSpecialAbilities()

    this.propWatcher = this.$watch("abilities", (newVal, oldVal) => {
      this.drawValues()
      this.calculateSpecialAbilities()
    });
  },
  methods: {
    drawValues: function () {
      let abilities = {};
      let params    = {};

      for (let i = 0; i <= 200; i++) {
        params[i] = {};
      }

      for (let ability of this.abilities.split("^")) {
        if (ability.split(",").length === 0) {
          continue;
        }

        const abilitySplit = ability.split(",")
        const abilityId    = abilitySplit[0].trim()
        const value        = abilitySplit[1].trim()

        if (value > 0) {
          abilities[abilityId] = value

          for (let i = 2; i < abilitySplit.length; i++) {
            const value = abilitySplit[i].trim()
            if (typeof params[abilityId] === "undefined") {
              params[abilityId] = {}
            }

            params[abilityId][i - 2] = value
          }
        }
      }

      // set form values
      this.ability       = abilities
      this.abilityParams = params
    },
    calculateSpecialAbilities: function () {
      this.$forceUpdate();

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
  text-align: right
}

.ability-label-top {
  width:      90px;
  text-align: right;
}
</style>
