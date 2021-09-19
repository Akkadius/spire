<template>
  <eq-window :title="'Spells (' + spells.length + ')'">
    <div class='eq-window-nested-blue text-center' v-if="spells.length === 0">
      No spells were found
    </div>

    <div class='eq-window-nested-blue' v-if="spells.length > 0">
      <table id="tabbox1" class="eq-table eq-highlight-rows" style="display: table;">
        <thead>
        <tr>
          <th style="width: 240px">Spell</th>
          <th style="width: 170px;">Level</th>
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
          </td>
          <td>{{ getClasses(spell) }}</td>
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

export default {
  name: "EqSpellPreviewTableRow",
  components: {
    EqSpellDescription,
    EqSpellEffects,
    EqWindow
  },
  created() {

  },
  data() {
    return {
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
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
    }
  }
}
</script>

<style scoped>

</style>
