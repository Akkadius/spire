<template>
  <div v-if="spell" class="d-inline-block">
    <div
      :id="spell.id + '-' + popoverId + '-popover'"
    >
      <span
        :class="'d-inline-block spell-' + spell.new_icon + getIconClassLabelFromSize()"
        :title="spell.icon"
        :style="getBorderStyling()"
      />
      <span
        :class="getLeftMargin()"
        :style="'position:relative;' + getTextTopOffset()"
      >{{ clampSpellName(spell.name) }} {{ annotation }}</span>

    </div>

    <b-popover
      :target="spell.id + '-' + popoverId + '-popover'"
      custom-class="no-bg"
      placement="right"
      delay="0"
      boundary="viewport"
      :no-fade="true"
      triggers="hover focus"
      style="width: 500px !important"
    >
      <eq-window style="width: auto; height: 100%">
        <eq-spell-preview :spell-data="spell"/>
      </eq-window>
    </b-popover>

  </div>
</template>

<script>
import EqWindow           from "./eq-ui/EQWindow";
import EqSpellCardPreview from "./preview/EQSpellCardPreview";
import EqSpellPreview     from "./preview/EQSpellCardPreview";
import {Spells}           from "@/app/spells";

export default {
  name: "SpellPopover",
  props: {
    spell: Object,
    size: { // options: 12,20,30,40
      type: [Number, String],
      required: false,
      default: 40
    },
    spellNameLength: {
      type: Number,
      required: false,
    },
    annotation: {
      type: String,
      required: false,
      default: ""
    }
  },
  watch: {
    spell: {
      deep: true,
      handler() {
      }
    },
  },
  data() {
    return {
      popoverId: Math.random().toString(16).slice(2),
      spellEffectInfo: [],
      spellData: {},
      sideLoadedSpellData: {},
      reagents: [],
    }
  },
  methods: {
    truncate(string, length) {
      if (string.length > length) {
        return string.substring(0, length) + '...';
      }

      return string
    },

    clampSpellName(name) {
      return (this.spellNameLength ? this.truncate(name, this.spellNameLength) : name)
    },

    getLeftMargin() {
      if (this.size <= 12) {
        return 'ml-2'
      }
      else if (this.size <= 20) {
        return 'ml-2'
      }
      else if (this.size <= 30) {
        return 'ml-3'
      }

      return 'ml-3'
    },
    getIconClassLabelFromSize() {
      if (this.size <= 12) {
        return '-12'
      }
      else if (this.size <= 20) {
        return '-20'
      }
      else if (this.size <= 30) {
        return '-30'
      }

      return '-40'
    },
    getBorderStyling() {
      if (this.size <= 12) {
        return ' border: 1px solid ' + this.getTargetTypeColor()
      }
      else if (this.size <= 20) {
        return 'position: relative; top: 5px; border-radius: 5px; border: .5px solid ' + this.getTargetTypeColor()
      }
      else if (this.size <= 30) {
        return 'border-radius: 5px; border: 1px solid ' + this.getTargetTypeColor()
      }

      return 'border-radius: 5px; border: 1px solid ' + this.getTargetTypeColor()
    },
    getTextTopOffset() {
      if (this.size <= 12) {
        return 'top: -3px'
      }
      else if (this.size <= 20) {
        return 'top: 0px'
      }
      else if (this.size <= 30) {
        return 'top: -10px'
      }

      return 'top: -15px'
    },

    getTargetTypeColor() {
      return Spells.getTargetTypeColor(this.spell["targettype"]);
    },
  },
  mounted() {
    this.sideLoadedSpellData = Spells.data
  },
  components: { EqSpellPreview, EqSpellCardPreview, EqWindow }
}
</script>
