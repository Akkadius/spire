<template>
  <eq-window :title="'Spells (' + spells.length + ')'">
    <div class='eq-window-nested-blue text-center' v-if="spells.length === 0">
      No spells were found
    </div>

    <div class='eq-window-nested-blue' v-if="spells.length > 0">
      <table id="tabbox1" class="eq-table eq-highlight-rows" style="display: table;">
        <thead>
        <tr>
          <th style="width: auto; min-width: 200px">Spell</th>
          <th style="width: auto; min-width: 130px">Level</th>
          <th>Mana</th>
          <th style="width: 80px">Cast</th>
          <th style="width: 80px">Recast</th>
          <th style="width: 120px">Duration</th>
          <th>Target</th>
          <th style="width: 400px">Effects</th>
          <th>Description</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(spell, index) in spells" :key="spell.id">
          <td>
            <img :src="spellCdnUrl + spell.new_icon + '.gif'"
                 :style="'width:20px;height:auto;border-radius: 10px; ' + 'border: 2px solid ' + getTargetTypeColor(spell['targettype']) + '; border-radius: 7px;'">
            {{ spell.name }}
<!--            <v-runtime-template :template="'<span>lol' + spellMini[spell.id] + '</span>'" />-->
          </td>
          <td>{{ getClasses(spell) }}</td>
          <td>{{ spell['mana'] > 0 ? spell['mana'] : "" }}</td>
          <td> {{ (spell["cast_time"] / 1000) }} sec</td>
          <td> {{ (spell["recast_time"] / 1000) }} sec</td>
          <td> {{ humanTime(getBuffDuration(spell) * 6) }} - {{ getBuffDuration(spell) }} tic(s)</td>
          <td> {{ getTargetTypeName(spell["targettype"]) }} </td>
          <td>
            <eq-spell-effects :spell="spell"/>
          </td>
          <td>
            <eq-spell-description :spell="spell"/>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </eq-window>


</template>

<script>
import {Spells} from "@/app/spells/spells";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import EqSpellEffects from "@/components/eq-ui/EQSpellEffects";
import {App} from "@/constants/app";
import EqSpellDescription from "@/components/eq-ui/EQSpellDescription";
import {DB_SPELL_TARGETS} from "@/app/constants/eq-spell-constants";

export default {
  name: "EqSpellPreviewTableRow",
  components: {
    EqSpellDescription,
    EqSpellEffects,
    EqWindow,
    "v-runtime-template": () => import("v-runtime-template"),
  },
  async created() {
    for (const spell of this.spells) {
      Spells.setSpell(spell['id'], spell)
      // Spells.renderSpellMini("0", spell['id']).then((result) => {
      //   this.spellMini[spell['id']] = result
      // })
    }

    // do this once so we're not triggering vue re-renders in the loop
    this.sideLoadedSpellData = Spells.data
  },
  data() {
    return {
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      spellMini: {},
      sideLoadedSpellData: {},
      componentId: ""
    }
  },
  props: {
    spells: Array
  },
  methods: {
    getClasses: function (spell) {
      return Spells.getClasses(spell)
    },
    getTargetTypeColor: function(targetType) {
      return Spells.getTargetTypeColor(targetType)
    },
    getTargetTypeName: function (targetType) {
      return DB_SPELL_TARGETS[targetType] ? DB_SPELL_TARGETS[targetType] : "Unknown Target (" + targetType + ")"
    },
    humanTime: function (sec) {
      let result = ""
      if (sec === 0) {
        result = "time";
      } else {
        let h  = Math.floor(sec / 3600);
        let m  = Math.floor((sec - h * 3600) / 60);
        let s  = sec - h * 3600 - m * 60;
        result = (h > 1 ? h + " hours " : "") + (h === 1 ? "1 hour " : "") + (m > 0 ? m + " min " : "") + (s > 0 ? s + " sec" : "");
      }

      return result;
    },
    getBuffDuration: function (spell) {
      return Spells.getBuffDuration(spell)
    },
  }
}
</script>

<style scoped>
.eq-table tr {
  border-bottom: .4px solid #ffffff1c;
}
.eq-table td {
  padding-top: 5px;
  padding-bottom: 5px;
}
</style>
