import {
  DB_SPA,
  DB_SPELL_EFFECTS,
  DB_SPELL_NEGATETYPE,
  DB_SPELL_NUMHITSTYPE,
  DB_SPELL_PETCMDS,
  DB_SPELL_RESISTS,
  DB_SPELL_TARGET_RESTRICTION,
  DB_SPELL_TARGETS,
  DB_SPELL_WORN_ATTRIBUTE_CAP,
  SPELL_TARGET_TYPE_COLORS
}                                          from "@/app/constants/eq-spell-constants";
import {DB_RACE_NAMES}                     from "@/app/constants/eq-races-constants";
import {DB_BARD_SKILLS, DB_SKILLS}         from "@/app/constants/eq-skill-constants";
import {BODYTYPES}                         from "@/app/constants/eq-bodytype-constants";
import util                                from "util";
import {DB_CLASSES, DB_CLASSES_WEAR_SHORT} from "@/app/constants/eq-classes-constants";
import {DbStrApi, ItemApi, SpellsNewApi}   from "@/app/api";
import {SpireApiClient}                    from "@/app/api/spire-api-client";
import {App}                               from "@/constants/app";
import {ItemStore}                         from "@/app/store/itemStore";

export class Spells {
  public static data           = {}
  public static dbstrData      = {}
  public static dbstrPreloaded = false

  public static async getItem(itemId) {
    const api    = (new ItemApi(SpireApiClient.getOpenApiConfig()))
    const result = await api.getItem({id: itemId})
    if (result.status === 200) {
      return result.data
    }

    return {}
  };

  public static getClasses(spell) {
    let classData = []
    for (let i = 1; i <= 16; i++) {
      const classIndex = "classes_" + i
      if ((spell[classIndex] > 0) && (spell[classIndex] < 255)) {
        // @ts-ignore
        classData.push(DB_CLASSES[i] + " (" + spell[classIndex] + ")")
      }
    }

    return classData.join(", ")
  };

  public static async getSpellEffectInfo(spell, effectIndex) {
    let effectsInfo = []

    // TODO: Handle elsewhere
    let serverMaxLevel = 100;

    let printBuffer = "";
    let tmp         = ""
    let pertick     = spell["buffduration"] ? " per tick " : ""
    let base        = spell["effect_base_value_" + effectIndex]
    let limit       = spell["effect_limit_value_" + effectIndex]
    let max         = spell["max_" + effectIndex]


    if (spell["effectid_" + effectIndex] !== 254) {

      //TODO For some reason not getting level currently from spell
      //let maxlvl = spell["effect_base_value_" + effectIndex];
      let maxlvl = serverMaxLevel;
      let minlvl = 255; // make this 255; FIX THIS

      for (let classId = 1; classId <= 16; classId++) {
        if (spell["classes_" + classId] < minlvl) {
          minlvl = spell["classes_" + classId];
        }
      }

      let value_min = this.calcSpellEffectValue(spell["formula_" + effectIndex], base, max, 1, minlvl);
      let value_max = this.calcSpellEffectValue(spell["formula_" + effectIndex], base, max, 1, serverMaxLevel);

      if ((value_min < value_max) && (value_max < 0)) {
        let tn    = value_min;
        value_min = value_max;
        value_max = tn;
      }

      let special_range = this.CalcValueRange(spell["formula_" + effectIndex], base, max, spell["effectid_" + effectIndex], spell["buffduration"], serverMaxLevel)

      if ((spell["formula_" + effectIndex] != 100) && (minlvl < 255)) {
        maxlvl = this.getSpellMaxOutLevel(spell["formula_" + effectIndex], base, max, minlvl)
      }

      switch (spell["effectid_" + effectIndex]) {

        case 0:
          tmp += limit ? " (" + DB_SPELL_TARGET_RESTRICTION[Math.abs(limit)] + ")" : ""
          printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range + tmp
          break;

        case 1:
          printBuffer += this.getFormatStandard("AC", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 2:
          printBuffer += this.getFormatStandard("ATK", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 3:
          printBuffer += this.getFormatStandard("Movement Speed", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 4:
          printBuffer += this.getFormatStandard("STR", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 5:
          printBuffer += this.getFormatStandard("DEX", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 6:
          printBuffer += this.getFormatStandard("AGI", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 7:
          printBuffer += this.getFormatStandard("STA", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 8:
          printBuffer += this.getFormatStandard("INT", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 9:
          printBuffer += this.getFormatStandard("WIS", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 10:
          if (base === 0 && (spell["formula_" + effectIndex] === 100)) { //This is used as a placeholder do not calculate
            printBuffer = ""
          } else {
            printBuffer += this.getFormatStandard("CHA", "", value_min, value_max, minlvl, maxlvl);
          }
          break;

        case 11: // Slow 70=>-30%, Haste 130=>+30%
          if (base < 100) {
            value_max = (100 - value_max) * -1;
            value_min = (100 - value_min) * -1;
          } else {
            value_max = value_max - 100;
            value_min = value_min - 100;

          }
          printBuffer += this.getFormatStandard("Attack Speed", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 12: //note: eqemu does not support base1 "enhanced invisibility" value
          printBuffer += "Invisibility (Unstable)"
          break;

        case 13: //note: eqemu does not support base1 "enhanced see invisibility" value
          printBuffer += "See Invisible"
          break;

        case 14:
          printBuffer += "Enduring Breath"
          break;

        case 15:
          printBuffer += this.getFormatStandard("Current Mana", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range
          break;

        case 16:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 17:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 18:
          printBuffer += "Pacify"
          break;

        case 19:
          printBuffer += this.getFormatStandard("Faction", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 20:
          printBuffer += "Blind"
          break;

        case 21:
          if (base !== limit && limit !== 0) {
            tmp += " ( " + (limit / 1000) + " in PvP)"
          }

          printBuffer += "Stun for " + (base / 1000) + " sec" + tmp + this.getUpToMaxLvl(max)
          break;

        case 22:
          printBuffer += "Charm" + this.getUpToMaxLvl(max)
          break;

        case 23:
          printBuffer += "Fear" + this.getUpToMaxLvl(max)
          break;

        case 24:
          printBuffer += this.getFormatStandard("Stamina Loss", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 25:
          if (base === 2) {
            tmp += " (Secondary Bind Point)"
          } else if (base === 3) {
            tmp += " (Tertiary Bind Point)"
          }
          printBuffer += "Bind" + tmp
          break;

        case 26:
          if (limit === 2) {
            tmp += " to Secondary Bind Point "
          } else if (limit === 3) {
            tmp += " to Tertiary Bind Point "
          }

          printBuffer += "Gate" + tmp + " (" + (100 - base) + "% chance to fail)"
          break;

        case 27:
          printBuffer += "Dispel" + " with level modifier of " + " (" + base + ")"
          break;

        case 28:
          printBuffer += "Invisibility to Undead (Unstable)"
          break;

        case 29:
          printBuffer += "Invisibility to Animals (Unstable)"
          break;

        case 30:
          printBuffer += "Decrease Aggro Radius to " + base + this.getUpToMaxLvl(max)
          break;

        case 31:
          printBuffer += "Mesmerize" + this.getUpToMaxLvl(max)
          break;

        case 32:
          printBuffer += "Summon Item: "

          const item = <any>(await this.getItem(spell["effect_base_value_" + effectIndex]));

          ItemStore.setItem(item.id, item)

          if (item.name) {
            printBuffer += `
                <div :id="${effectIndex} + '-' + ${item.id} + '-' + componentId" style="display:inline-block" class="ml-2">

                  <div style="display: inline-block">
                    <img
                      :src="itemCdnUrl + 'item_' + ${item.icon} + '.png'"
                      style="height:15px; border-radius: 25px; width:auto;"
                      class="mr-1">
                    <span class="mr-1">${item.name}</span>
                  </div>

                </div>

                <b-popover
                  :target="${effectIndex} + '-' + ${item.id} + '-' + componentId"
                  placement="auto"
                  custom-class="no-bg"
                  delay="1"
                  triggers="hover focus"
                  style="width: 500px !important"
                >
                  <eq-window style="margin-right: 10px; width: auto; height: 90%">
                    <eq-item-preview :item-data="itemData[${item.id}]"/>
                  </eq-window>
                </b-popover>`
          }
          break;

        case 33:
          printBuffer += "Summon Pet: " + spell["teleport_zone"]
          break;

        case 34:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 35:
          printBuffer += this.getFormatStandard("Disease Counter", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 36:
          printBuffer += this.getFormatStandard("Poison Counter", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 37:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 38:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 39:
          printBuffer += "Can not be Twincast"
          break;

        case 40:
          printBuffer += "Invulnerability"
          break;

        case 41:
          printBuffer += "Destroy"
          break;

        case 42: //handled by client, unknown what base value represents
          printBuffer += "Shadowstep"
          break;

        case 43: //custom on eqemu, not used on live. Any client with this effect now has chance to crippling blow
          printBuffer += "Berserk: Allows chance to crippling blow"
          break;

        case 44: //TODO This is a type of buff stacker that I need to figure out
          printBuffer += "Lycanthropy: Need to implement on Eqemu"
          break;

        case 45: //custom on eqemu, not used on live. Stackable melee lifesteal effect.
          printBuffer += "Lifetap from Weapon Damage: " + base + "%"
          break;

        case 46:
          printBuffer += this.getFormatStandard("Fire Resist", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 47:
          printBuffer += this.getFormatStandard("Cold Resist", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 48:
          printBuffer += this.getFormatStandard("Poison Resist", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 49:
          printBuffer += this.getFormatStandard("Disease Resist", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 50:
          printBuffer += this.getFormatStandard("Magic Resist", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 51:
          printBuffer += "Sense Undead"
          break;

        case 52:
          printBuffer += "Sense Summoned"
          break;

        case 53:
          printBuffer += "Sense Animal"
          break;

        case 54:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 55: //rune
          printBuffer += "Absorb Damage: 100%, Total: " + base
          break;

        case 56:
          printBuffer += "True North"
          break;

        case 57:
          printBuffer += "Levitate"
          break;

        case 58: // Illusion:
          printBuffer += "Illusion: " + DB_RACE_NAMES[spell["effect_base_value_" + effectIndex]]
          break;

        case 59:
          printBuffer += this.getFormatStandard("Damage Shield", "", -value_min, -value_max, minlvl, maxlvl);
          break;

        case 60:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 61:
          printBuffer += "Identify Item"
          break;

        case 62:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 63:
          printBuffer += "Memory Blur" + " (" + base + "% chance)"
          break;

        case 63:
          printBuffer += "Memory Blur" + " (" + base + "% chance)"
          break;

        case 64:
          if (base !== limit && limit !== 0) {
            printBuffer += "Stun and Spin NPC for " + (base / 1000) + " sec (PC for " + (limit / 1000) + " sec " + this.getUpToMaxLvl(max)
          } else {
            printBuffer += "Stun and Spin for " + (base / 1000) + " sec " + this.getUpToMaxLvl(max)
          }
          break;

        case 65:
          printBuffer += "Infravision"
          break;

        case 66:
          printBuffer += "Ultravision"
          break;

        case 67:
          printBuffer += "Eye of Zomm"
          break;

        case 68:
          printBuffer += "Reclaim Pet Mana"
          break;

        case 69:
          printBuffer += this.getFormatStandard("Max HP", "", value_min, value_max, minlvl, maxlvl) + special_range
          break;

        case 70:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 71:
          printBuffer += "Summon Pet: " + spell["teleport_zone"]
          break;

        case 72:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 73:
          printBuffer += "Bind Sight"
          break;

        case 74:
          printBuffer += "Feign Death (" + base + "% chance)"
          break;

        case 75:
          printBuffer += "Project Voice"
          break;

        case 76:
          printBuffer += "Sentinel"
          break;

        case 77:
          printBuffer += "Locate Corpse"
          break;

        case 78: //spell rune
          printBuffer += "Absorb Spell Damage: 100%, Total: " + base
          break;

        case 79:
          tmp += limit ? " (" + DB_SPELL_TARGET_RESTRICTION[Math.abs(limit)] + ")" : ""
          printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl) + special_range + tmp
          break;

        case 80:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 81:
          printBuffer += "Resurrect with " + base + "% XP"
          break;

        case 82:
          printBuffer += "Summon Player"
          break;

        case 83: //TODO teleport zone enum
          printBuffer += "Teleport to " + spell["teleport_zone"]
          break;

        case 84: // base on emu does damage. Correct?
          printBuffer += "Gravity Flux"
          break;

        case 85: //TODO Spell proc LINK
          printBuffer += "Add Melee Proc: " + (await this.renderSpellMini(spell.id, base)) + (limit ? " with " + limit + " % Rate Mod" : "")
          break;

        case 86:
          printBuffer += "Decrease Social Radius to " + base + this.getUpToMaxLvl(max)
          break;

        case 87:
          printBuffer += this.getFormatStandard("Magnification", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 88: //TODO clean up, enum for zones
          if (spell["teleport_zone"] !== "same") {
            tmp += " (" + spell["effect_base_value_" + (effectIndex + 1)]
                   + ", " + spell["effect_base_value_" + effectIndex] + ", "
                   + spell["effect_base_value_" + (effectIndex + 2)] + ", "
                   + spell["effect_base_value_" + (effectIndex + 3)] + ")"
          }
          printBuffer += "Evacuate to " + spell["teleport_zone"] + tmp
          break;

        case 89:
          if (base < 100) {
            value_max = (100 - value_max) * -1;
            value_min = (100 - value_min) * -1;
          } else {
            value_max = value_max - 100;
            value_min = value_min - 100;
          }
          printBuffer += this.getFormatStandard("Player Size", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 90: // pet invisible - This is not implemented on eqemu
          printBuffer += "Ignore Pet (not implemented)"
          break;

        case 91:
          printBuffer += "Summon Corpse up to level " + base
          break;

        case 92:
          printBuffer += this.getFormatStandard("Hate", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 93:
          printBuffer += "Stop Rain"
          break;

        case 94:
          printBuffer += "Cancel if Combat Initiated"
          break;

        case 95:
          printBuffer += "Sacrifice"
          break;

        case 96:
          printBuffer += "Inhibit Spell Casting (Silence)"
          break;

        case 97:
          printBuffer += this.getFormatStandard("Max Mana", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 98:
          if (base < 100) {
            value_max = (100 - value_max) * -1;
            value_min = (100 - value_min) * -1;
          } else {
            value_max = value_max - 100;
            value_min = value_min - 100;

          }
          printBuffer += this.getFormatStandard("Attack Speed (v2 capped)", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 99:
          printBuffer += "Root"
          break;

        case 100:
          tmp += limit ? " (" + DB_SPELL_TARGET_RESTRICTION[Math.abs(limit)] + ")" : ""
          printBuffer += this.getFormatStandard("Current HP", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range + tmp
          break;

        case 101:
          printBuffer += "Increase Current HP by " + (base * 7500) + " with recast blocking buff"
          break;

        case 102:
          printBuffer += "Fear Immunity"
          break;

        case 103:
          printBuffer += "Summon Pet to Player"
          break;

        case 104: //TODO clean up, enum for zones
          if (spell["teleport_zone"] !== "") {
            tmp += spell["teleport_zone"] +
                   " (" + spell["effect_base_value_" + (effectIndex + 1)]
                   + ", " + spell["effect_base_value_" + effectIndex] + ", "
                   + spell["effect_base_value_" + (effectIndex + 2)] + ", "
                   + spell["effect_base_value_" + (effectIndex + 3)] + ")"
          } else {
            tmp += "bind"
          }
          printBuffer += "Translocate to " + tmp
          break;

        case 105:
          printBuffer += "Inhibit Gate";
          return;

        case 106:
          printBuffer += "Summon Warder: " + spell["teleport_zone"]
          break;

        case 107: //not on live
          printBuffer += this.getFormatStandard("NPC Level", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 108: //limit "Ignore Auto Leave" not supported, not in current spell file
          printBuffer += "Summon Familiar:" + spell["teleport_zone"]
          break;

        case 109: //later expansions allow stacks to put into bags using limit value.
          printBuffer += "Summon into Bag: "

          const item2 = <any>(await this.getItem(spell["effect_base_value_" + effectIndex]));

          ItemStore.setItem(item2.id, item2);

          if (item2.name) {
            printBuffer += `
                <div :id="${effectIndex} + '-' + ${item2.id} + '-' + componentId" style="display:inline-block" class="ml-2">

                  <div style="display: inline-block">
                    <img
                      :src="itemCdnUrl + 'item_' + ${item2.icon} + '.png'"
                      style="height:15px; border-radius: 25px; width:auto;"
                      class="mr-1">
                    <span class="mr-1">${item2.name}</span>
                  </div>

                </div>

                <b-popover
                  :target="${effectIndex} + '-' + ${item2.id} + '-' + componentId"
                  placement="auto"
                  custom-class="no-bg"
                  delay="1"
                  triggers="hover focus"
                  style="width: 500px !important"
                >
                  <eq-window style="margin-right: 10px; width: auto; height: 90%">
                    <eq-item-preview :item-data="itemData[${item2.id}]"/>
                  </eq-window>
                </b-popover>`
          }
          break;

        case 110:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 111:
          printBuffer += this.getFormatStandard("All Resists", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 112:
          printBuffer += this.getFormatStandard("Effective Casting Level", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 113:
          printBuffer += "Summon Mount: " + spell["teleport_zone"]
          break;

        case 114:
          printBuffer += this.getFormatStandard("Hate Generated", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 115:
          printBuffer += "Reset Hunger Counter"
          break;

        case 116:
          printBuffer += this.getFormatStandard("Curse Counter", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 117:
          printBuffer += "Make Weapon Magical"
          break;

        case 118:
          printBuffer += this.getFormatStandard("Singing Amplification", "%", (value_min * 10), (value_max * 10), minlvl, maxlvl);
          break;

        case 119:
          if (base < 100) {
            value_max = (100 - value_max) * -1;
            value_min = (100 - value_min) * -1;
          } else {
            value_max = value_max - 100;
            value_min = value_min - 100;
          }

          printBuffer += this.getFormatStandard("Attack Speed (v3 over cap)", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 120:
          printBuffer += this.getFormatStandard("Healing Taken", "%", value_min, value_max, minlvl, maxlvl);
          break;

        case 121:
          printBuffer += this.getFormatStandard("Reverse Damage Shield", "", value_min, value_max, minlvl, maxlvl);
          break;

        case 122: //TODO implement this
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not implemented"
          break;

        case 123:
          printBuffer += "Buff Blocker: Screech (" + base + ")"
          break;

        case 124:
          printBuffer += this.getFocusPercentRange("Spell Damage", base, limit, false);
          break;

        case 125:
          printBuffer += this.getFocusPercentRange("Healing", base, limit, false);
          break;

        case 126:
          printBuffer += this.getFocusPercentRange("Spell Resist Rate", base, limit, true);
          break;

        case 127:
          printBuffer += this.getFocusPercentRange("Spell Haste", base, limit, false);
          break;

        case 128:
          printBuffer += this.getFocusPercentRange("Spell Duration", base, limit, false);
          break;

        case 129:
          printBuffer += this.getFocusPercentRange("Spell Range", base, limit, false);
          break;

        case 130:
          printBuffer += this.getFocusPercentRange("Spell and Bash Hate", base, limit, false);
          break;

        case 131:
          printBuffer += this.getFocusPercentRange("Chance of Using Reagent", base, limit, true);
          break;

        case 132:
          printBuffer += this.getFocusPercentRange("Spell Mana Cost", base, limit, true);
          break;

        case 133: //Not on live
          printBuffer += this.getFocusPercentRange("Stun Time", base, limit, false);
          break;

        case 134:
          printBuffer += "Limit Max Level: " + base + " (lose " + (limit ? limit : 100) + "% per level)"
          break;

        case 135:
          printBuffer += "Limit Resist: " + (base < 0 ? "Exclude " : "") + DB_SPELL_RESISTS[Math.abs(base)]
          break;

        case 136:
          printBuffer += "Limit Target: " + (base < 0 ? "Exclude " : "") + DB_SPELL_TARGETS[Math.abs(base)]
          break;

        case 137: //Maybe include number id  + " (SPA: " + Math.abs(base) + ")"
          printBuffer += "Limit Effect: " + (base < 0 ? "Exclude " : "") + DB_SPA[Math.abs(base)]
          break;

        case 138:
          printBuffer += "Limit Type: " + (base ? "Beneficial" : "Detrimental")
          break;

        case 139://TODO need spell links
          printBuffer += "Limit Spell: " + (base < 0 ? "Exclude " : "") + (await this.renderSpellMini(spell.id, Math.abs(base)))
          break;

        case 140:
          printBuffer += "Limit Min Duration: " + (base * 6) + "s"
          break;

        case 141:
          printBuffer += "Limit Duration Type: " + (base ? "Non-Duration Spells" : "Duration Spells")
          break;

        case 142:
          printBuffer += "Limit Min Level: " + base
          break;

        case 143:
          printBuffer += "Limit Min Casting Time: " + (base / 1000) + "s"
          break;

        case 144:
          printBuffer += "Limit Max Casting Time: " + (base / 1000) + "s"
          break;

        case 145:
          printBuffer += "Teleport to " + spell["teleport_zone"]
          break;

        case 146: //todo data location for port xyz for 45 , Set position to
          break;

        case 147:
          printBuffer += this.getFormatStandard("Current HP", "%", value_min, value_max, minlvl, maxlvl) + " up to " + max
          break;

        case 148:
          tmp += limit ? limit : spell["formula_" + effectIndex] % 100
          printBuffer += "Stacking: Block new spell if slot " + tmp + " is " + DB_SPA[Math.abs(base)] + " and less than " + max
          break;

        case 149:
          tmp += limit ? limit : spell["formula_" + effectIndex] % 100
          printBuffer += "Stacking: Overwrite spell if slot " + tmp + " is " + DB_SPA[Math.abs(base)] + " and less than " + max
          break;

        case 150:
          tmp += max ? " (Increase heal by " + max + " if affected is above lv " + limit + ")" : ""
          printBuffer += (base === 1) ? "Divine Intervention with 300 Heal" + tmp : "Divine Intervention with 8000 Heal" + tmp
          break;

        case 151:
          printBuffer += "Suspend Pet" + (base ? " with Buffs" : "")
          break;

        case 152:
          printBuffer += "Summon Temp Pet: " + spell["teleport_zone"] + " x " + base + " for " + max + "s"
          break;

        case 153:
          printBuffer += "Balance Group HP with " + base + "% Penalty (Max HP taken: " + limit + ")"
          break;

        case 154: //TODO need to update emulator code to use percent based (+0.5% per level difference) and confirm duration change mechanic
          if (limit !== 0) {
            printBuffer += "Decrease Detrimental Duration by 50% " + (base / 10) + "% Chance)" + this.getUpToMaxLvl(max)
          } else {
            printBuffer += "Dispel Detrimental " + (base / 10) + "% Chance" + this.getUpToMaxLvl(max)
          }
          break;

        case 156:
          printBuffer += "Illusion: Target"
          break;

        case 157:
          printBuffer += this.getFormatStandard("Spell Damage Shield", "", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 158:
          tmp += max ? " with up to " + max + "% Base Damage" : ""
          if (limit > 0) {
            tmp += " and " + limit + " Improved Resist Mod"
          } else if (limit < 0) {
            tmp += " and " + limit + " Reduced Resist Mod"
          }
          printBuffer += this.getFormatStandard("Chance to Reflect Spell", "%", value_min, value_max, minlvl, maxlvl) + tmp
          break;

        case 159:
          printBuffer += this.getFormatStandard("Base Stats", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 160:
          printBuffer += "Intoxicate if Tolerance under " + base
          break;

        case 161:
          tmp += limit ? "Max Per Hit: " + limit : ""
          tmp += max ? ", Total: " + max : ""
          printBuffer += "Absorb Spell Damage: " + base + "%" + tmp
          break;

        case 162:
          tmp += limit ? "Max Per Hit: " + limit : ""
          tmp += max ? ", Total: " + max : ""
          printBuffer += "Absorb Melee Damage: " + base + "%" + tmp
          break;

        case 163:
          tmp += max ? ", Max Per Hit: " + max : ""
          printBuffer += "Absorb " + base + " Hits or Spells " + base + "%" + tmp
          break;

        case 164:
          printBuffer += "Appraise Chest " + value_max
          break;

        case 165:
          printBuffer += "Disarm Chest " + value_max
          break;

        case 166:
          printBuffer += "Unlock Chest " + value_max
          break;

        case 167:
          printBuffer += "Increase Pet Power " + value_max
          break;

        case 168:
          printBuffer += this.getFormatStandard("Melee Mitigation", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 169:
          tmp += limit >= 0 ? " with " + DB_SKILLS[limit] : ""
          printBuffer += this.getFormatStandard("Chance to Critical Hit", "%", value_min, value_max, minlvl, maxlvl) + tmp
          break;

        case 170:
          printBuffer += this.getFormatStandard("Critical Nuke Damage", "%", value_min, value_max, minlvl, maxlvl) + " of Base Damage"
          break;

        case 171:
          printBuffer += this.getFormatStandard("Chance to Crippling Blow", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 172:
          printBuffer += this.getFormatStandard("Chance to Avoid Melee", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 173:
          printBuffer += this.getFormatStandard("Chance to Riposte", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 174:
          printBuffer += this.getFormatStandard("Chance to Dodge", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 175:
          printBuffer += this.getFormatStandard("Chance to Parry", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 176:
          printBuffer += this.getFormatStandard("Chance to Dual Wield", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 177:
          printBuffer += this.getFormatStandard("Chance to Double Attack", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 178:
          printBuffer += "Lifetap from Weapon Damage: " + base + "%"
          break;

        case 179:
          printBuffer += "Instrument Modifier: " + DB_SKILLS[spell["skill"]] + " " + value_max
          break;

        case 180:
          printBuffer += this.getFormatStandard("Chance to Resist Spell", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 181:
          printBuffer += this.getFormatStandard("Chance to Resist Fear Spell", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 182:
          printBuffer += this.getFormatStandard("Weapon Delay", "%", value_min / 10, value_max / 10, minlvl, maxlvl)
          break;

        case 183:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 184:
          printBuffer += this.getFormatStandard("Chance to Hit", "%", value_min, value_max, minlvl, maxlvl) + (limit >= 0 ? " with " + DB_SKILLS[limit] : "")
          break;

        case 185:
          printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Damage", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 186:
          printBuffer += this.getFormatStandard("Min " + DB_SKILLS[limit] + " Damage", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 187:
          printBuffer += "Balance Group Mana with " + base + "% Penalty (Max HP taken: " + limit + ")"
          break;

        case 188:
          printBuffer += this.getFormatStandard("Chance to Block", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 189:
          printBuffer += this.getFormatStandard("Current Endurance", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range
          break;

        case 190:
          printBuffer += this.getFormatStandard("Max Endurance", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 191:
          printBuffer += "Inhibit Combat"
          break;

        case 192:
          printBuffer += this.getFormatStandard("Hate", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range
          break;

        case 193:
          printBuffer += DB_SKILLS[spell["skill"]] + " Attack for " + base + " with " + limit + " % Accuracy Mod"
          break;

        case 194:
          printBuffer += "Cancel Aggro (" + base + " % Chance)"
          break;

        case 195:
          printBuffer += this.getFormatStandard("Chance to Resist Any Stun", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 196:
          printBuffer += this.getFormatStandard("Srikethrough", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 197:
          tmp += (limit >= 0) ? DB_SKILLS[limit] : " Hit "
          printBuffer += this.getFormatStandard(tmp + " Damage Taken", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 198:
          printBuffer += this.getFormatStandard("Current Endurance", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 199:
          printBuffer += "Taunt with " + limit + " added Hate (Chance " + base + "%)"
          break;

        case 200:
          printBuffer += this.getFormatStandard("Worn Proc Rate", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 201:
          printBuffer += "Add Range Proc: " + (await this.renderSpellMini(spell.id, base)) + " with " + limit + "% Rate Mod"
          break;

        case 202:
          printBuffer += "Project Illusion on Next Spell"
          break;

        case 203:
          printBuffer += "Mass Group Buff on Next Spell"
          break;

        case 204:
          printBuffer += "Group Fear Immunity for " + (base * 10) + "s"
          break;

        case 205:
          printBuffer += "Rampage"
          break;

        case 206:
          printBuffer += "AE Taunt with " + base + " added Hate"
          break;

        case 207:
          printBuffer += "Flesh to Bone Chips"
          break;

        case 208:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 209: //TODO need to update emulator code to use percent based (+0.5% per level difference) and confirm duration change mechanic
          if (limit !== 0) {
            printBuffer += "Decrease Beneficial Duration by 50% " + (base / 10) + "% Chance)" + this.getUpToMaxLvl(max)
          } else {
            printBuffer += "Dispel Beneficial " + (base / 10) + "% Chance" + this.getUpToMaxLvl(max)
          }
          break;

        case 210:
          printBuffer += "Pet Shielding for " + base * 12 + "s"
          break;

        case 211: //eqemu uses this, Live uses different formula now, base=chance, limit=damage mod
          printBuffer += "AE Melee for " + base * 12 + "s"
          break;

        case 212: //eqemu uses this, Live uses different formula now, base=chance, limit, mana cost mod
          printBuffer += this.getFormatStandard("Frenzied Devastation: Chance to Critical Nuke", "%", value_min, value_max, minlvl, maxlvl) + " and Increase Spell Mana Cost 100%"
          break;

        case 213:
          printBuffer += this.getFormatStandard("Pet Max HP", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 214:
          printBuffer += this.getFormatStandard("Max HP", "%", value_min / 100, value_max / 100, minlvl, maxlvl)
          break;

        case 215:
          printBuffer += this.getFormatStandard("Pet Chance to Avoid Melee", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 216:
          printBuffer += this.getFormatStandard("Accuracy", "", value_min, value_max, minlvl, maxlvl) + (limit >= 0 ? " with " + DB_SKILLS[limit] : "")
          break;

        case 217:
          printBuffer += "Add Headshot Proc with up to " + limit + " Damage"
          break;

        case 218:
          printBuffer += this.getFormatStandard("Pet Chance to Critical Hit", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 219:
          printBuffer += this.getFormatStandard("Chance to Slay Undead", "%", value_min / 100, value_max / 100, minlvl, maxlvl) + " with " + limit + " Damage Mod"
          break;

        case 220:
          printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Damage Bonus", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 221:
          printBuffer += this.getFormatStandard("Weight", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 222:
          printBuffer += this.getFormatStandard("Chance to Block from Back", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 223:
          printBuffer += this.getFormatStandard("Chance to Double Riposte", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 224:
          if (limit > 0) {
            printBuffer += this.getFormatStandard("Chance of Additional Riposte", "%", value_min, value_max, minlvl, maxlvl) + " with " + DB_SKILLS[limit]
          } else {
            printBuffer += this.getFormatStandard("Chance of Additional Riposte", "%", value_min, value_max, minlvl, maxlvl)
          }
          break;

        case 225:
          printBuffer += this.getFormatStandard("Chance to Double Attack ", "%", value_min, value_max, minlvl, maxlvl) + " (Additive)";
          break;

        case 226:
          printBuffer += "Add Two-Handed Bash Ability";
          break;

        case 227:
          printBuffer += "Decrease " + DB_SKILLS[limit] + " Timer by " + this.humanTime(base) + " (Before Haste)"
          break;

        case 228:
          printBuffer += this.getFormatStandard("Falling Damage", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 229:
          printBuffer += this.getFormatStandard("Chance to Cast Through Stun", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 230:
          printBuffer += this.getFormatStandard("Shield Ability Range", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 231:
          printBuffer += this.getFormatStandard("Chance to Stun Bash", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 232:
          tmp += limit ? " and " + (await this.renderSpellMini(spell.id, limit)) : ""
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, 4789)) + tmp + " on Death (" + base + "% Chance Divine Save)"
          break;

        case 233:
          printBuffer += this.getFormatStandard("Food Consumption", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 234:
          printBuffer += "Decrease Poison Application Time by " + (10 - base / 1000) + +"s"
          break;

        case 235:
          printBuffer += this.getFormatStandard("Chance to Channel Spells", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 236:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 237:
          printBuffer += "Enable Pet Ability: Receive Group Buffs"
          break;

        case 238:
          printBuffer += (base === 3) ? "Permanent Illusion (Persist After Death)" : " Permanent Illusion"
          break;

        case 239:
          printBuffer += this.getFormatStandard("Chance to Feign Death Through Spell Hit", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 240:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 241:
          printBuffer += "Reclaim Pet Mana (Return " + base + "%)"
          break;

        case 242:
          printBuffer += this.getFormatStandard("Chance to Memory Blur", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 243:
          printBuffer += this.getFormatStandard("Chance of Charm Breaking", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 244:
          printBuffer += this.getFormatStandard("Chance of Root Breaking", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 245:
          printBuffer += this.getFormatStandard("Chance of Trap Circumvention", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 246:
          printBuffer += this.getFormatStandard("Lung Capacity", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 247:
          printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Skill Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 248:
          printBuffer += "Train Second Magic Specialization Ability (Secondary Forte)"
          break;

        case 249:
          printBuffer += this.getFormatStandard("Offhand Weapon Damage Bonus", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 250:
          printBuffer += this.getFormatStandard("Melee Proc Rate (from buffs, abilities and skills", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 251:
          printBuffer += this.getFormatStandard("Chance of Using Ammo", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 252:
          printBuffer += this.getFormatStandard("Chance to Backstab From Front", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 253:
          printBuffer += "Allow Frontal Backstab for Minimum Damage"
          break;

        case 255:
          printBuffer += "Increase Shield Ability Duration by " + base + "s"
          break;

        case 256:
          printBuffer += "Shroud of Stealth (" + base + ")"
          break;

        case 257:
          printBuffer += "Enable Pet Ability: Hold"
          break;

        case 258:
          printBuffer += this.getFormatStandard("Chance to Triple Backstab", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 259:
          printBuffer += this.getFormatStandard("AC Soft Cap", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 260:
          printBuffer += this.getFormatStandard(DB_BARD_SKILLS[limit] + " Bonus", "%", (value_min * 10), (value_max * 10), minlvl, maxlvl)
          break;

        case 261:
          printBuffer += this.getFormatStandard("Song Cap", "", (value_min * 10), (value_max * 10), minlvl, maxlvl)
          break;

        case 262:
          printBuffer += this.getFormatStandard(DB_SPELL_WORN_ATTRIBUTE_CAP[limit] + " Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 263:
          printBuffer += this.getFormatStandard("Ability to Specialize Tradeskills", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 264:
          printBuffer += "Reduce [AA " + limit + "] Timer by " + this.humanTime(base)
          break;

        case 265:
          printBuffer += "No Fizzle up to level " + base
          break;

        case 266:
          printBuffer += this.getFormatStandard("Chance of " + limit + " Additional 2H Attacks", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 267:
          printBuffer += "Enable Pet Ability: (" + DB_SPELL_PETCMDS[limit] + ")"
          break;

        case 268:
          printBuffer += this.getFormatStandard("Chance to Fail " + DB_SKILLS[limit] + " Combine", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 269:
          printBuffer += this.getFormatStandard("Bandage HP Cap", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 270:
          printBuffer += this.getFormatStandard("Beneficial Song Range", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 271:
          printBuffer += this.getFormatStandard("Innate Movement Speed", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 272:
          printBuffer += this.getFormatStandard("Song effective casting level", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 273: //Does live effect now have decay component?
          printBuffer += this.getFormatStandard("Chance to Critical DoT", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 274: //Does live effect now have decay component?
          printBuffer += this.getFormatStandard("Chance to Critical Heal", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 275:
          printBuffer += this.getFormatStandard("Chance to Critical Mend", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 276:
          printBuffer += this.getFormatStandard("Dual Wield Skill Amount", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 277:
          printBuffer += "Second Chance to Trigger Divine Intervention with a Heal for " + base + "% of baseline"
          break;

        case 278:
          printBuffer += " Add Finishing Blow Proc with up to " + (base / 10) + " Damage (" + limit + "% Chance)"
          break;

        case 279:
          printBuffer += this.getFormatStandard("Chance to Flurry", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 280:
          printBuffer += this.getFormatStandard("Pet Chance to Flurry", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 281:
          printBuffer += "Pet Chance to Feign Death (" + base + "%)"
          break;

        case 282:
          printBuffer += this.getFormatStandard("Bandage Amount", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 283:
          printBuffer += this.getFormatStandard("Chance to perform a Double Special Attack", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 284:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 285:
          printBuffer += "Chance Hide skill will succeed while moving (" + base + "%)"
          break;

        case 286:
          printBuffer += this.getFormatStandard("Spell Damage Amount", "", value_min, value_max, minlvl, maxlvl) + " (before crit)"
          break;

        case 287:
          printBuffer += this.getFormatStandard("Spell Duration", "seconds", (value_min * 6), (value_max * 6), minlvl, maxlvl)
          break;

        case 288: //TODO finish this when AA tables are added, this procs the spell associated with the AA, rank.spell is what the spell id that procs is
          printBuffer += "Add [Insert AA spell] Proc to" + DB_SPA[limit] + "(" + (base / 10) + "% Chance)"
          break;

        case 289:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " on Duration Fade"
          break;

        case 290:
          printBuffer += this.getFormatStandard("Movement Speed Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 291: //TODO Fix this in source code, not coded correct
          printBuffer += "Remove up to (" + base + ") detrimental effects"
          break;

        case 292:
          printBuffer += this.getFormatStandard("Chance of Strikethrough", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 293:
          printBuffer += this.getFormatStandard("Chance to Resist Melee Stun", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 294:
          printBuffer += (base) ? this.getFormatStandard("Chance to Critical Nuke", "%", value_min, value_max, minlvl, maxlvl) : ""
          printBuffer += (base) ? " and " : ""
          printBuffer += (limit) ? this.getFormatStandard("Critical Nuke Damage", "%", limit, limit, minlvl, maxlvl) + " of Base Damage" : ""
          break;

        case 295:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 296:
          printBuffer += this.getFocusPercentRange("Spell Damage Taken", base, limit, false);
          break;

        case 297:
          printBuffer += this.getFormatStandard("Spell Damage Taken", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 298:
          printBuffer += this.getFormatStandard("Pet Size", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 299:
          printBuffer += "Wake the Dead (" + max + ")"
          break;

        case 300:
          printBuffer += "Summon Doppelganger: " + spell["teleport_zone"]
          break;

        case 301:
          printBuffer += this.getFormatStandard("Archery Damage", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 302: //crit
          printBuffer += this.getFocusPercentRange("Spell Damage", base, limit, false);
          break;

        case 303: //crit
          printBuffer += this.getFormatStandard("Spell Damage", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 304:
          printBuffer += this.getFormatStandard("Chance to Avoid Offhand Riposte", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 305: //Note: Percent decrease damage taken is negative number from AA, but positive number on Spells, there are spells with negative values that increase damage taken.
          printBuffer += this.getFormatStandard("Offhand Damage Shield Taken", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 306:
          printBuffer += "Wake the Dead: + " + spell["teleport_zone"] + " x " + base + " for " + max + "s"
          break;

        case 307:
          printBuffer += "Appraisal"
          break;

        case 308:
          printBuffer += "Suspend Minion to remain after Zoning";
          break;

        case 309:
          printBuffer += "Teleport to Caster's Bind";
          break;

        case 310:
          printBuffer += "Reduce Timer by " + this.humanTime(base / 1000)
          break;

        case 311:
          printBuffer += "Limit Type: " + (base === 1 ? "Include" : "Exclude") + " Combat Skills"
          break;

        case 312:
          printBuffer += "Sanctuary: Place caster bottom hate list, fades if cast on other than self."
          break;

        case 313:
          printBuffer += this.getFormatStandard("Chance to Double Forage", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 314:
          printBuffer += "Invisibility" + (base > 1 ? " (Enhanced " + base + ")" : "")
          break;

        case 315:
          printBuffer += "Invisibility to Undead" + (base > 1 ? " (Enhanced " + base + ")" : "")
          break;

        case 316:
          printBuffer += "Invisibility to Animals" + (base > 1 ? " (Enhanced " + base + ")" : "")
          break;

        case 317:
          printBuffer += this.getFormatStandard("HP Regen Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 318:
          printBuffer += this.getFormatStandard("Mana Regen Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 319:
          printBuffer += this.getFormatStandard("Chance to Critical HoT", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 320:
          printBuffer += this.getFormatStandard("Shield Block Chance", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 321:
          printBuffer += this.getFormatStandard("Target's Target Hate", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 322:
          printBuffer += "Gate to Home City"
          break;

        case 323:
          printBuffer += "Add Defensive Proc: " + (await this.renderSpellMini(spell.id, spell["effect_base_value_" + effectIndex])) + (limit ? " with " + limit + " % Rate Mod" : "")
          break;

        case 324:
          printBuffer += "Cast from HP with " + base + "% Penalty"
          break;

        case 325:
          printBuffer += this.getFormatStandard("Chance to Remain Hidden When Hit By AE", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 326:
          printBuffer += this.getFormatStandard("Spell Memorization Gems", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 327:
          printBuffer += this.getFormatStandard("Buff Slots", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 328:
          printBuffer += this.getFormatStandard("Max Negative HP", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 329:
          printBuffer += "Absorb Damage using Mana: " + base + "%"
          break;

        case 330:
          printBuffer += this.getFormatStandard("Critical " + DB_SKILLS[limit] + " Damage", "%", value_min, value_max, minlvl, maxlvl) + " of Base Damage"
          break;

        case 331:
          printBuffer += this.getFormatStandard("Chance to Salvage Components", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 332:
          printBuffer += "Summon to Corpse"
          break;

        case 333:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + "on Rune Fade"
          break;

        case 334:
          printBuffer += this.getFormatStandard("Current HP", "%", value_min, value_max, minlvl, maxlvl) + pertick + special_range + " (If Target Not Moving)";
          break;

        case 335:
          printBuffer += "Block Next Spell" + (base < 100 ? " (" + base + "% Chance)" : "")
          break;

        case 336:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 337:
          printBuffer += this.getFormatStandard("Experience Gain", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 338:
          printBuffer += "Summon and Resurrect All Corpses"
          break;

        case 339:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + "on Spell Use (" + base + "% Chance)"
          break;

        case 340: //Only one effect casts if multiple 340s in spell
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + (base < 100 ? " (" + base + "% Chance)" : "")
          break;

        case 341:
          printBuffer += this.getFormatStandard("ATK Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 342:
          printBuffer += "Inhibit Low Health Fleeing"
          break;

        case 343:
          printBuffer += "Interrupt Casting" + (base < 100 ? "(" + base + "% Chance)" : "")
          break;

        case 344:
          printBuffer += this.getFormatStandard("Chance to Channel Item Click Effects", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 345:
          printBuffer += "Limit Assassinate Level: " + base + (limit ? "(" + limit + "% Chance Bonus)" : "")
          break;

        case 346:
          printBuffer += "Limit Headshot Level: " + base + (limit ? "(" + limit + "% Chance Bonus)" : "")
          break;

        case 347:
          printBuffer += this.getFormatStandard("Chance of Double Archery Attack", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 348:
          printBuffer += "Limit: Min Mana Cost: " + base
          break;

        case 349:
          printBuffer += this.getFormatStandard("Damage When Shield Equipped", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 350:
          printBuffer += "Manaburn: Consumes up to " + base + " mana to deal " + -limit + "% of that mana as direct damage"
          break;

        case 351: //TODO
          printBuffer += "Aura Effect: Need to link to aura table"
          break;

        case 352:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 353:
          printBuffer += this.getFormatStandard("Aura Count", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 354:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 355:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 356:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 357:
          printBuffer += "Inhibit Spell Casting (Focus Silence) (" + base + "% Chance)";
          break;

        case 358:
          printBuffer += this.getFormatStandard("Current Mana", "", value_min, value_max, minlvl, maxlvl) + special_range
          break;

        case 359:
          printBuffer += this.getFormatStandard("Chance to Sense Trap", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 360:
          printBuffer += "Add Killshot Proc: " + await this.renderSpellMini(spell.id, limit) + " (" + base + "% Chance)" + (max ? " Target Max Lv: " + max : "")
          break;

        case 361:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + " on Death (" + base + "% Chance)"
          break;

        case 362:
          printBuffer += this.getFormatStandard("Potion Belt Slots", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 363:
          printBuffer += this.getFormatStandard("Bandolier Slots", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 364:
          printBuffer += this.getFormatStandard("Chance to Triple Attack", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 365:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + " if spell Kills Target (" + base + "% Chance)"
          break;

        case 366:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 367:
          printBuffer += "Transform Body Type to " + BODYTYPES[base]
          break;

        case 368: //TODO: get faction name from dbase
          printBuffer += this.getFormatStandard("Faction with [Faction " + base + "]", "", limit, limit, minlvl, maxlvl)
          break;

        case 369:
          printBuffer += this.getFormatStandard("Corruption Counter", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 370:
          printBuffer += this.getFormatStandard("Corruption Resist", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 371:
          printBuffer += this.getFormatStandard("Attack Speed", "", -value_min, -value_max, minlvl, maxlvl) + "(Stackable)"
          break;

        case 372:
          printBuffer += this.getFormatStandard("Forage Skill Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 373:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " on Fade"
          break;

        case 374:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + (base < 100 ? " (" + base + "% Chance)" : "")
          break;

        case 375:
          printBuffer += this.getFormatStandard("Critical DoT Damage", "%", value_min, value_max, minlvl, maxlvl) + " of Base Damage"
          break;

        case 376:
          printBuffer += "Fling"
          break;

        case 377:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " on Duration Finished"
          break;

        case 378:
          printBuffer += this.getFormatStandard("Chance to Resist " + DB_SPA[limit] + " Effects", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 379://client handles this function
          if (limit === 0) {
            printBuffer += "Shadowstep Forward " + base
          }
          if (limit === 90) {
            printBuffer += "Shadowstep Right " + base
          }
          if (limit === 180) {
            printBuffer += "Shadowstep Back " + base
          }
          if (limit === 270) {
            printBuffer += "Shadowstep Left " + base
          } else {
            printBuffer += "Shadowstep " + base + " to " + limit + " Degrees"
          }
          break;

        case 380:
          printBuffer += "Push Back " + limit + " and Up " + base
          break;

        case 381:
          printBuffer += "Fling to Self (Velocity: " + base + ")" + (max ? "Target must be " + max + " or fewer lv higher than you" : "")
          break;

        case 382:
          printBuffer += "Inhibit Effect: " + DB_SPA[limit] + (base ? " (From: " + DB_SPELL_NEGATETYPE[base] + " Effects)" : "")
          break;

        case 383: // spell proc + " (Sympathetic Proc)"
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + " on Spell Use" + (base !== 100 ? " (Proc rate mod: " + (base - 100) + "%)" : "")
          break;

        case 384:
          printBuffer += "Fling to Target (Velocity: " + base + ")"
          break;

        case 385:
          const spellGroupId   = Math.abs(base);
          const spellGroupName = await this.getSpellGroupNameById(spellGroupId);

          printBuffer += util.format(
            "Limit Spell Group: %s %s",
            (base >= 0 ? "" : "Exclude "),
            spellGroupName
          )

          break;

        case 386:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " on Curer"
          break;

        case 387:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " on Cured"
          break;

        case 388:
          printBuffer += "Summon All Corpses (From Any Zone)";
          break;

        case 389:
          printBuffer += "Reset Spell Recast Timers";
          break;

        case 390:
          printBuffer += "Set Spell Lockout Recast Timers to " + this.humanTime(base / 1000)
          break;

        case 391:
          printBuffer += "Limit Max Mana Cost: " + base
          break;

        case 392:
          printBuffer += this.getFormatStandard("Healing Amount", "", value_min, value_max, minlvl, maxlvl) + "(After Crit)"
          break;

        case 393: //Focus version
          printBuffer += this.getFocusPercentRange("Heal Taken", base, limit, false);
          break;

        case 394:
          printBuffer += this.getFormatStandard("Healing Taken Amount", "", value_min, value_max, minlvl, maxlvl) + "(Before Crit)"
          break;

        case 395:
          printBuffer += this.getFocusPercentRange("Heal Taken", base, limit, false) + "(Before Crit)"
          break;

        case 396:
          printBuffer += this.getFormatStandard("Healing Amount", "", value_min, value_max, minlvl, maxlvl) + "(Before Crit)"
          break;

        case 397:
          printBuffer += this.getFormatStandard("Pet AC", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 398:
          printBuffer += "Increase Temp Pet Duration by " + (base / 1000) + " sec"
          break;

        case 399:
          printBuffer += this.getFormatStandard("Chance to Twincast", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 400:
          printBuffer += "Increase Groups Current HP by up to " + Math.abs(Math.floor(base * limit / 10)) + " (" + Math.abs(limit / 10) + " HP per 1 Mana Drained)"
          break;

        case 401:
          printBuffer += "Decrease Current HP by up to " + Math.abs(Math.floor(base * limit / 10)) + " and Drain up to " + base + " mana  (" + Math.abs(limit / 10) + " HP per 1 Target Mana Drained)"
          break;

        case 402:
          printBuffer += "Decrease Current HP by up to " + Math.abs(Math.floor(base * limit / 10)) + " and Drain up to " + base + " endurance  (" + Math.abs(limit / 10) + " HP per 1 Target Endurance Drained)"
          break;

        case 403: //Do not have defines for this, corresponds to spell table field data spell_class (field 221)
          printBuffer += "Limit Spell Class: " + (base >= 0 ? "" : "Exclude ") + "(ID: " + Math.abs(base) + ")"
          break;

        case 404: //Do not have defines for this, corresponds to spell table field data spell_subclass (field 222)
          printBuffer += "Limit Spell Subclass: " + (base >= 0 ? "" : "Exclude ") + "ID: " + Math.abs(base) + ")"
          break;

        case 405:
          printBuffer += this.getFormatStandard("Staff Block Chance", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 406:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " if Max Hits Used"
          break;

        case 407:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " on Focus Limit Match"
          break;

        case 408:
          printBuffer += "Cap HP at " + (limit > 0 ? "lowest of " + base + "% or " + limit : +base + "%")
          break;

        case 409:
          printBuffer += "Cap Mana at " + (limit > 0 ? "lowest of " + base + "% or " + limit : +base + "%")
          break;

        case 410:
          printBuffer += "Cap Endurance at " + (limit > 0 ? "lowest of " + base + "% or " + limit : +base + "%")
          break;

        case 411:
          printBuffer += "Limit Class: " + DB_CLASSES_WEAR_SHORT[base >> 1]
          break;

        case 412:
          printBuffer += "Limit Race:  " + DB_RACE_NAMES[base]
          break;

        case 413:
          printBuffer += this.getFormatStandard("Base Spell Effectiveness", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 414:
          printBuffer += "Limit Casting Skill: " + DB_SKILLS[base]
          break;

        case 415:
          printBuffer += "Error: (" + spell["effectid_" + effectIndex] + ") not used"
          break;

        case 416:
          printBuffer += this.getFormatStandard("AC", "", value_min, value_max, minlvl, maxlvl) + "(v2)"
          break

        case 417:
          printBuffer += this.getFormatStandard("Current Mana", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range + "(v2)"
          break

        case 418:
          printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Damage bonus", "", value_min, value_max, minlvl, maxlvl) + pertick + special_range + "(v2)"
          break

        case 419:
          printBuffer += "Add Melee Proc (v2): " + (await this.renderSpellMini(spell.id, base)) + (limit ? " with " + limit + " % Rate Mod" : "")
          break;

        case 420:
          printBuffer += this.getFormatStandard("Max Hits Count", "%", value_min, value_max, minlvl, maxlvl)
          break

        case 421:
          printBuffer += this.getFormatStandard("Max Hits Count", "", value_min, value_max, minlvl, maxlvl)
          break

        case 422:
          printBuffer += "Limit Max Hits Min: " + base
          break;

        case 423:
          printBuffer += "Limit Max Hits Type: " + DB_SPELL_NUMHITSTYPE[base]
          break;

        case 424:
          printBuffer += "Gradual " + (base > 0 ? "Push" : "Pull") + " to " + limit + "' away (Force=" + Math.abs(base) + ")" + this.getUpToMaxLvl(max)
          break;

        case 425: // not implemented on eqemu
          printBuffer += "Fly"
          break;

        case 426:
          printBuffer += this.getFormatStandard("Extended Target Window Slots", "", value_min, value_max, minlvl, maxlvl)
          break

        case 427:
          printBuffer += "Add Skill Proc: " + (await this.renderSpellMini(spell.id, base)) + (limit ? " with " + limit + " % Rate Mod" : "")
          break;

        case 428:
          printBuffer += "Limit Skill: " + DB_SKILLS[base]
          break;

        case 429:
          printBuffer += "Add Skill Proc on Successful Hit: " + (await this.renderSpellMini(spell.id, base)) + (limit ? " with " + limit + " % Rate Mod" : "")
          break;

        case 430: // not implemented on eqemu
          printBuffer += "Alter Vision: Base1=" + base + " Base2=" + limit + " Max=" + max
          break;

        case 431: // not implemented on eqemu
          if (base < 0) {
            printBuffer += "Tint Vision: Red= " + (base >> 16 & 0xff) + " Green=" + (base >> 8 & 0xff) + " Blue=" + (base & 0xff)
          } else {
            printBuffer += "Alter Vision: Base1=" + base + " Base2=" + limit + " Max=" + max
          }
          break;

        case 432:
          printBuffer += this.getFormatStandard("Trophy Slots", "", value_min, value_max, minlvl, maxlvl)
          break;

        /* EQEMU DOESN"T USE THESE YET, were changed on live after ROF2
         case 433:
         printBuffer += this.getFormatStandard(DB_SKILL[limit] + "Damage Bonus", "", value_min, value_max, minlvl, maxlvl) + " (v433, Delay Mod)"
         break;

         case 433:
         printBuffer += this.getFormatStandard(DB_SKILL[limit] + "Damage Bonus", "", value_min, value_max, minlvl, maxlvl) + " (v434, Delay Mod)"
         break;
         return Spell.FormatCount(Spell.FormatEnum((SpellSkill)base2) + " Damage Bonus", base1) + " (v433, Delay Mod)";

         case 435:
         printBuffer += "Fragile Defense (" + base + ")"
         break;
         */

        case 433:
          printBuffer += this.getFormatStandard("Crtical DoT Chance", "%", value_min, value_max, minlvl, maxlvl) + "Decay Rate of " + limit + " over level " + max
          break

        case 434:
          printBuffer += this.getFormatStandard("Crtical Heal Chance", "%", value_min, value_max, minlvl, maxlvl) + "Decay Rate of " + limit + " over level " + max
          break

        case 435:
          printBuffer += this.getFormatStandard("Crtical HoT Chance", "%", value_min, value_max, minlvl, maxlvl) + "Decay Rate of " + limit + " over level " + max
          break

        case 436:
          printBuffer += "Toggle: Freeze Buffs"
          break;

        case 437:
          if (base === 52584) {
            tmp += "Primary Anchor"
          }
          if (base === 52585) {
            tmp += "Secondary Anchor"
          }
          if (base === 50874) {
            tmp += "Guild Anchor"
          }
          printBuffer += "Teleport to your " + tmp
          break;

        case 438:
          if (base === 52584) {
            tmp += "Primary Anchor"
          }
          if (base === 52585) {
            tmp += "Secondary Anchor"
          }
          if (base === 50874) {
            tmp += "Guild Anchor"
          }
          printBuffer += "Teleport to their " + tmp
          break;

        case 439:
          printBuffer += "Add Assassinate Proc with up to " + limit + "Damage" + (base ? " Chance Mod:" + base : "")
          break;

        case 440:
          printBuffer += "Limit Finishing Blow Level to " + base + " and lower NPC targets with" + (limit / 10) + "% or less health."
          break;

        case 441:
          printBuffer += "Cancel if Moved " + base + "'"
          break;

        case 442:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " once if " + DB_SPELL_TARGET_RESTRICTION[limit]
          break;

        case 443:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " once if Caster " + DB_SPELL_TARGET_RESTRICTION[limit]
          break;

        case 444:
          printBuffer += "Lock Aggro on Caster and " + this.getFormatStandard("Other Aggro", "%", (limit - 100), (limit - 100), minlvl, maxlvl) + this.getUpToMaxLvl(base)
          break;

        case 445:
          printBuffer += "Grant " + base + "Mercenary Slots"
          break;

        case 446:
          printBuffer += "Buff Blocker A (" + base + ")"
          break;

        case 447:
          printBuffer += "Buff Blocker B (" + base + ")"
          break;

        case 448:
          printBuffer += "Buff Blocker C (" + base + ")"
          break;

        case 449:
          printBuffer += "Buff Blocker D (" + base + ")"
          break;

        case 450:
          printBuffer += "Absorb DoT Damage: " + base + "%" + (limit > 0 ? "Max Per Hit: " + limit : "") + (max > 0 ? " Total: " + max : "")
          break;

        case 451:
          printBuffer += "Absorb Melee Damage: " + base + "% over " + limit + +(max > 0 ? " Total: " + max : "")
          break;

        case 452:
          printBuffer += "Absorb Spell Damage: " + base + "% over " + limit + +(max > 0 ? " Total: " + max : "")
          break;

        case 453:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " if " + limit + " Melee Damage Taken in Single Hit"
          break;

        case 454:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, base)) + " if " + limit + " Spell Damage Taken in Single Hit"
          break;

        case 455:
          printBuffer += this.getFormatStandard("Current Hate", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 456:
          printBuffer += this.getFormatStandard("Current Hate", "%", value_min, value_max, minlvl, maxlvl) + pertick
          break;

        case 457:
          if (limit === 0) {
            tmp += "HP"
          }
          if (limit === 1) {
            tmp += "Mana"
          }
          if (limit === 2) {
            tmp += "Endurance"
          }
          printBuffer += "Return " + (base / 10) + "% of Spell Damage as" + tmp + (max > 0 ? ", Max Per Hit: " + max + ")" : "")
          break;

        case 458:
          printBuffer += this.getFormatStandard("Faction Hit", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 459:
          printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Damage", "%", value_min, value_max, minlvl, maxlvl) + " (v2)"
          break;

        case 460:
          printBuffer += "Limit Type: Include Non-Focusable"
          break;

        case 461:
          printBuffer += this.getFocusPercentRange("Spell Damage", base, limit, false) + " (Before Crit)"
          break;

        case 462:
          printBuffer += this.getFormatStandard("Spell Damage Amount", "", value_min, value_max, minlvl, maxlvl) + " (After Crit)"
          break;

        case 463:
          printBuffer += "Melee Shielding: " + base + "%"
          break;

        case 464:
          printBuffer += this.getFormatStandard("Pet Chance to Rampage", "%", value_min, value_max, minlvl, maxlvl) + (limit ? " with " + limit + "% of Damage" : "")
          break;

        case 465:
          printBuffer += this.getFormatStandard("Pet Chance to AE Rampage", "%", value_min, value_max, minlvl, maxlvl) + (limit ? " with " + limit + "% of Damage" : "")
          break;

        case 466:
          printBuffer += this.getFormatStandard("Pet Chance to Flurry on Double Attack", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 467:
          printBuffer += this.getFormatStandard("Damage Shield Taken", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 468:
          printBuffer += this.getFormatStandard("Damage Shield Taken", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 469: //TODO need spell group defines need query /Only one effect casts if multiple 340s in spell
          printBuffer += "Cast Highest Rank of [Group " + (await this.getSpellGroupNameById(Math.abs(base))) + "]" + (base < 100 ? " (" + base + "% Chance) " : "")
          break;

        case 470:
          printBuffer += "Cast Highest Rank of [Group " + limit + "]" + (base < 100 ? " (" + base + "% Chance)" : "")
          break;

        case 471:
          printBuffer += this.getFormatStandard("Chance to Repeat Primary Hand Round", "%", value_min, value_max, minlvl, maxlvl) + (limit ? " with " + limit + " % Damage Bonus" : "")
          break;

        case 472:
          printBuffer += "Buy AA Rank (" + base + ")"
          break;

        case 473:
          printBuffer += this.getFormatStandard("Chance to Double Backstab From Front", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 474:
          printBuffer += this.getFormatStandard("Pet Critical Hit Damage", "%", value_min, value_max, minlvl, maxlvl) + " of Base Damage"
          break;

        case 475:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + " if Not Cast By Item Click" + (base ? " (Chance " + base + "%)" : "")
          break;

        case 476:
          if (base === 0) {
            tmp += "2H Weapons"
          }
          if (base === 1) {
            tmp += "Shields"
          }
          if (base === 2) {
            tmp += "Dual Wield"
          }
          printBuffer += "Weapon Stance: Apply spell" + (await this.renderSpellMini(spell.id, limit)) + " when using " + tmp
          break;

        case 477:
          printBuffer += "Move to Top of Rampage List (" + base + "% Chance)"
          break;

        case 478:
          printBuffer += "Move to Bottom of Rampage List (" + base + "% Chance)"
          break;

        case 479:
          printBuffer += "Limit Effect: " + DB_SPA[limit] + " greater than " + base
          break;

        case 480:
          printBuffer += "Limit Effect: " + DB_SPA[limit] + " less than " + base
          break;

        case 481:
          printBuffer += "Cast " + (await this.renderSpellMini(spell.id, limit)) + " if Hit By Spell" + (base < 100 ? "(" + base + "% Chance)" : "")
          break;

        case 482:
          printBuffer += this.getFormatStandard("Base " + DB_SKILLS[limit] + " Damage", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 483:
          printBuffer += this.getFocusPercentRange("Spell Damage Taken", base, limit, false) + "(v2)"
          break;

        case 484:
          printBuffer += this.getFormatStandard("Spell Damage Taken Amount", "", value_min, value_max, minlvl, maxlvl) + "(After Crit)"
          break;

        case 485:
          printBuffer += "Limit Caster Class: " + DB_CLASSES_WEAR_SHORT[base >> 1] + "(Outgoing Focus Limit)"
          break;

        case 486:
          printBuffer += "Limit Caster: " + (base === 0 ? "Exclude " : "") + "Self"
          break;

        case 487:
          printBuffer += this.getFormatStandard(DB_SKILLS[limit] + " Skill Cap with Recipes", "", value_min, value_max, minlvl, maxlvl) + "(After Crit)"
          break;

        case 488:
          printBuffer += this.getFormatStandard("Push Taken", "%", -value_min, -value_max, minlvl, maxlvl)
          break;

        case 489:
          printBuffer += this.getFormatStandard("Endurance Regen Cap", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 490:
          printBuffer += "Limit Min Recast: " + (base / 1000) + "s"
          break;

        case 491:
          printBuffer += "Limit Max Recast: " + (base / 1000) + "s"
          break;

        case 492:
          printBuffer += "Limit Min Endurance Cost: " + base
          break;

        case 493:
          printBuffer += "Limit Max Endurance Cost: " + base
          break;

        case 494:
          printBuffer += this.getFormatStandard("Pet ATK", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 495:
          printBuffer += "Limit Max Duration: " + (base * 6)
          break;

        case 496:
          printBuffer += this.getFormatStandard("Critical " + DB_SKILLS[limit] + " Damage", "%", value_min, value_max, minlvl, maxlvl) + " of Base Damage (Non Stacking)"
          break;

        case 497:
          printBuffer += "Limit: No Procs or Twincast"
          break;

        case 498:
          printBuffer += this.getFormatStandard("Chance of " + limit + " Additional 1H Attacks", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 499:
          printBuffer += this.getFormatStandard("Chance of " + limit + " Secondary 1H Attack", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 500:
          printBuffer += this.getFocusPercentRange("Spell Haste", base, limit, false) + "(No max reduction limit)"
          break;

        case 501:
          printBuffer += (base < 0 ? "Increase" : "Decrease") + " Casting Times by " + Math.abs(base / 1000) + "s"
          break;

        case 502:
          if (base !== limit && limit !== 0) {
            tmp += " ( " + (limit / 1000) + " in PvP)"
          }
          printBuffer += "Stun and Fear " + (base / 1000) + " sec" + tmp + this.getUpToMaxLvl(max)
          break;

        case 503:
          printBuffer += this.getFormatStandard((limit === 0 ? "Rear" : "Frontal") + " Arc Melee Damage", "%", value_min / 10, value_max / 10, minlvl, maxlvl)
          break;

        case 503:
          printBuffer += this.getFormatStandard((limit === 0 ? "Rear" : "Frontal") + " Arc Melee Damage Amount", "", value_min / 10, value_max / 10, minlvl, maxlvl)
          break;

        case 505:
          printBuffer += this.getFormatStandard((limit === 0 ? "Rear" : "Frontal") + " Arc Melee Damage Taken", "%", value_min / 10, value_max / 10, minlvl, maxlvl)
          break;

        case 506:
          printBuffer += this.getFormatStandard((limit === 0 ? "Rear" : "Frontal") + " Arc Melee Damage Taken Amount", "", value_min / 10, value_max / 10, minlvl, maxlvl)
          break;

        case 507:
          printBuffer += this.getFocusPercentRange("Spell Power", base, limit, false) + " (Focus Spell DOT, DD and Healing)"
          break;

        case 509:
          printBuffer += (limit < 0 ? "Decrease" : "Increase") + " Current HP by " + (Math.abs(limit) / 10) + "% of Caster Current HP ( " + (Math.abs(base) / 10) + "% Life Burn)"
          break;

        case 510:
          printBuffer += this.getFormatStandard("Incoming Resist Modifier", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 511:
          printBuffer += "Limit Min Delay Between Trigger: " + (limit / 1000) + "s" + " (Max Triggers: " + base + ")"
          break;

        case 512:
          printBuffer += "Proc Timer: " + (limit / 1000) + "s" + " (Max Triggers: " + base + ")"
          break;

        case 513:
          printBuffer += this.getFormatStandard("Max Mana", "%", value_min / 100, value_max / 100, minlvl, maxlvl)
          break;

        case 514:
          printBuffer += this.getFormatStandard("Max Endurance", "%", value_min / 100, value_max / 100, minlvl, maxlvl)
          break;

        case 515:
          printBuffer += this.getFormatStandard("Avoidance AC", "%", value_min / 1000, value_max / 1000, minlvl, maxlvl)
          break;

        case 516:
          printBuffer += this.getFormatStandard("Mitigation AC", "%", value_min / 1000, value_max / 1000, minlvl, maxlvl)
          break;

        case 517:
          printBuffer += this.getFormatStandard("ATK Offense", "%", value_min / 1000, value_max / 1000, minlvl, maxlvl)
          break;

        case 518:
          printBuffer += this.getFormatStandard("ATK Accuracy", "%", value_min / 1000, value_max / 1000, minlvl, maxlvl)
          break;

        case 519:
          printBuffer += this.getFormatStandard("Luck", "", value_min, value_max, minlvl, maxlvl)
          break;

        case 520:
          printBuffer += this.getFormatStandard("Luck", "%", value_min, value_max, minlvl, maxlvl)
          break;

        case 521:
          printBuffer += "Absorb Damage using Endurance: " + (base / 100) + (limit !== 10000 ? (limit / 10000) + " End per 1 HP)" : "") + (max > 0 ? ", Max Per Hit: " + max : "")
          break;

        case 522:
          printBuffer += this.getFormatStandard("Current Mana", "%", value_min / 100, value_max / 100, minlvl, maxlvl) + " up to " + max
          break;

        case 523:
          printBuffer += this.getFormatStandard("Current Endurance", "%", value_min / 100, value_max / 100, minlvl, maxlvl) + " up to " + max
          break;

        case 524:
          printBuffer += this.getFormatStandard("Current HP", "%", value_min, value_max, minlvl, maxlvl) + " up to " + max + pertick
          break;

        case 525:
          printBuffer += this.getFormatStandard("Current Mana", "%", value_min, value_max, minlvl, maxlvl) + " up to " + max + pertick
          break;

        case 526:
          printBuffer += this.getFormatStandard("Current Endurance", "%", value_min, value_max, minlvl, maxlvl) + " up to " + max + pertick
          break;
      }

      if (printBuffer !== "") {
        // @ts-ignore
        effectsInfo.push(util.format("%s) %s", effectIndex, printBuffer))
      }

      if (App.DEBUG && printBuffer !== "") {
        const debug = util.format(
          "--- Debug: Effect ID (%s) (%s) Index (%s)",
          spell["effectid_" + effectIndex],
          DB_SPELL_EFFECTS[spell["effectid_" + effectIndex]] ? DB_SPELL_EFFECTS[spell["effectid_" + effectIndex]] : "UNKNOWN",
          effectIndex
        )
        // @ts-ignore
        effectsInfo.push(debug)
      }
    }

    return effectsInfo;
  };

  public static calcSpellEffectValue(calc, base, max, tick, level) {
    if (calc === 0) {
      return base;
    }

    if (calc === 100) {
      if (max > 0 && base > max) {
        return max;
      }
      return base;
    }

    let change = 0;

    switch (calc) {
      case 100:
        break;

      case 101:
        change = level / 2;
        break;
      case 102:
        change = level;
        break;
      case 103:
        change = level * 2;
        break;
      case 104:
        change = level * 3;
        break;
      case 105:
        change = level * 4;
        break;
      case 107:
        change = -1 * tick;
        break;
      case 108:
        change = -2 * tick;
        break;
      case 109:
        change = level / 4;
        break;
      case 110:
        change = level / 6;
        break;
      case 111:
        if (level > 16) change = (level - 16) * 6;
        break;
      case 112:
        if (level > 24) change = (level - 24) * 8;
        break;
      case 113:
        if (level > 34) change = (level - 34) * 10;
        break;
      case 114:
        if (level > 44) change = (level - 44) * 15;
        break;
      case 115:
        if (level > 15) change = (level - 15) * 7;
        break;
      case 116:
        if (level > 24) change = (level - 24) * 10;
        break;
      case 117:
        if (level > 34) change = (level - 34) * 13;
        break;
      case 118:
        if (level > 44) change = (level - 44) * 20;
        break;
      case 119:
        change = level / 8;
        break;
      case 120:
        change = -5 * tick;
        break;
      case 121:
        change = level / 3;
        break;
      case 122:
        change = -12 * tick;
        break;
      case 123:
        change = (Math.abs(max) - Math.abs(base)) / 2;
        break;
      case 124:
        if (level > 50) change = (level - 50);
        break;
      case 125:
        if (level > 50) change = (level - 50) * 2;
        break;
      case 126:
        if (level > 50) change = (level - 50) * 3;
        break;
      case 127:
        if (level > 50) change = (level - 50) * 4;
        break;
      case 128:
        if (level > 50) change = (level - 50) * 5;
        break;
      case 129:
        if (level > 50) change = (level - 50) * 10;
        break;
      case 130:
        if (level > 50) change = (level - 50) * 15;
        break;
      case 131:
        if (level > 50) change = (level - 50) * 20;
        break;
      case 132:
        if (level > 50) change = (level - 50) * 25;
        break;
      case 139:
        if (level > 30) change = (level - 30) / 2;
        break;
      case 140:
        if (level > 30) change = (level - 30);
        break;
      case 141:
        if (level > 30) change = 3 * (level - 30) / 2;
        break;
      case 142:
        if (level > 30) change = 2 * (level - 60);
        break;
      case 143:
        change = 3 * level / 4;
        break;

      case 3000:
        return base;

      default:
        if (calc > 0 && calc < 1000) {
          change = level * calc;
        }

        if (calc >= 1000 && calc < 2000) {
          change = tick * (calc - 1000) * -1;
        }

        if (calc >= 2000 && calc < 3000) {
          change = level * (calc - 2000);
        }

        if (calc >= 4000 && calc < 5000) {
          change = -tick * (calc - 4000);
        }

        break;
    }

    let value = Math.abs(base) + change;

    if (max !== 0 && value > Math.abs(max)) {
      value = Math.abs(max);
    }

    if (base < 0) {
      value = -value;
    }

    return Math.trunc(value);
  };

  public static CalcValueRange(calc, base, max, spa, duration, level) {
    let printBuffer = ""
    let start       = this.calcSpellEffectValue(calc, base, max, 1, level);
    let finish      = Math.abs(this.calcSpellEffectValue(calc, base, max, duration, level));

    let type = Math.abs(start) < Math.abs(finish) ? "Growing" : "Decaying";

    if (calc === 123) {
      if (base < 0) {
        max = max * -1;
      }
      printBuffer = " (Random: " + Math.abs(base) + " to " + Math.abs(max) + ")"
    }

    if (calc === 107) {
      printBuffer = " (" + type + " to " + finish + " @ 1/tick)"
    }

    if (calc === 108) {
      printBuffer = " (" + type + " to " + finish + " @ 2/tick)"
    }

    if (calc === 120) {
      printBuffer = " (" + type + " to " + finish + " @ 5/tick)"
    }

    if (calc === 122) {
      printBuffer = " (" + type + " to " + finish + " @ 12/tick)"
    }

    if (calc > 1000 && calc < 2000) {
      printBuffer = " (" + type + " to " + finish + " @ " + (calc - 1000) + "/tick)"
    }

    if (calc >= 3000 && calc < 4000) {
      if (calc - 3000 === spa) {
        printBuffer = " (Scales, Base Level: 100)";
      }
      if (calc - 3500 === spa) {
        printBuffer = " (Scales, Base Level: 105)";
      }
    }

    if (calc > 4000 && calc < 5000) {
      printBuffer = " (" + type + " to " + finish + " @ " + (calc - 4000) + "/tick)"
    }

    return printBuffer;
  };

  public static getSpellMaxOutLevel(calc, base, max, minLevel) {
    let MaxServerLevel = 100 //Better way to define this.
    let value          = 0;

    for (let i = minLevel; i <= 100; i++) {

      value = this.calcSpellEffectValue(calc, base, max, 1, i)

      if (Math.abs(value) >= max) {
        return i;
      }
    }
    return MaxServerLevel
  };

  public static getFormatStandard(effect_name, type, value_min, value_max, minlvl, maxlvl) {
    let modifier = ""

    if (value_max < 0) {
      modifier = "Decrease "
    } else {
      modifier = "Increase "
    }

    let printBuffer = modifier + effect_name
    if (value_min !== value_max) {
      printBuffer += " by " + (Math.abs(value_min)) + type + " (L" + minlvl + ") to " + (Math.abs(value_max)) + type + " (L" + maxlvl + ")"
    } else {
      printBuffer += " by " + (Math.abs(value_max)) + type
    }

    return printBuffer;
  };

  public static getUpToMaxLvl(max) {
    let printBuffer = ""
    if (max > 0) {
      printBuffer = " up to level " + max
    }
    return printBuffer;
  };

  public static getFocusPercentRange(effect_name, min, max, negate) {
    let printBuffer = ""
    let modifier    = ""

    if (min < 0) {

      if (min < max) {
        let temp = min;
        min      = max;
        max      = temp;
      }
    } else {
      if (min > max)
        max = min;
    }

    if (negate) {
      min = -min;
      max = -max;
    }

    if (max < 0) {
      modifier = "Decrease "
    } else {
      modifier = "Increase "
    }

    if (min === max || max === 0) {
      printBuffer += modifier + effect_name + " by " + Math.abs(min) + "%";
      return printBuffer;
    }
    printBuffer += modifier + effect_name + " by " + Math.abs(min) + "% to " + Math.abs(max) + "%";
    return printBuffer;
  };

  public static async getSpellGroupNameById(spellGroupId) {
    const api = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))

    let filters = [
      ["spellgroup", "__", spellGroupId]
    ]

    let wheres = [];
    filters.forEach((filter) => {
      // @ts-ignore
      wheres.push(util.format("%s%s%s", filter[0], filter[1], filter[2]))
    })

    const result = await (api.listSpellsNews(
        {
          limit: "1",
          where: wheres.join("."),
          orderBy: "id"
        }
      )
    );

    if (result.status === 200) {
      if (result.data.length > 0) {
        return result.data[0].name;
      }
    }

    return "Unknown Spell Group";
  };

  public static getTargetTypeColor(targetType) {
    return SPELL_TARGET_TYPE_COLORS[targetType];
  };

  public static async getSpell(spellId) {
    if (spellId === 0) {
      return {}
    }

    const api    = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
    const result = await api.getSpellsNew({id: spellId})
    if (result.status === 200) {
      return result.data
    }

    return {}
  };

  public static humanTime(sec) {
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
  };

  public static getMinLevel(spell) {
    let minLevel = 0
    for (let i = 1; i <= 16; i++) {
      const classIndex = "classes_" + i
      if ((spell[classIndex] > 0) && (spell[classIndex] < 255)) {
        if (spell[classIndex] < minLevel) {
          minLevel = spell[classIndex];
        }
      }
    }
    return minLevel
  };

  public static getBuffDuration(spell) {
    let i            = 0
    let minLevel     = this.getMinLevel(spell)
    let buffDuration = spell["buffduration"]

    switch (spell["buffdurationformula"]) {
      case 0:
        return 0;
      case 1:
        i = Math.ceil(minLevel / 2);

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 2:
        i = Math.ceil(buffDuration / 5 * 3);

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 3:
        i = minLevel * 30;

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 4:
        return buffDuration;
      case 5:
        i = buffDuration;

        return (i < 3 ? (i < 1 ? 1 : i) : 3);
      case 6:
        i = Math.ceil(minLevel / 2);

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 7:
        i = minLevel;

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 8:
        i = minLevel + 10;

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 9:
        i = minLevel * 2 + 10;

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 10:
        i = minLevel * 3 + 10;

        return (i < buffDuration ? (i < 1 ? 1 : i) : buffDuration);
      case 11:
      case 12:
        return buffDuration;
      case 50:
        return 72000;
      case 3600:
        return (buffDuration ? buffDuration : 3600);
      default:
        return "???"
    }
  };

  public static async getSpellDescription(spell) {
    if (spell["descnum"] > 0) {
      let description = await this.getDbstrData(spell["descnum"])
      if (description) {

        // #1 Base for effect id 1
        // $1 Limit for effect id 1
        // @1 Max for effect id 1
        // %z (# ticks)
        for (let i = 1; i <= 12; i++) {
          const baseEffect = util.format("#%s", i)
          if (description.includes(baseEffect)) {
            description = description.replaceAll(baseEffect, String(Math.abs(spell["effect_base_value_" + i])))
          }
          const limitEffect = util.format("$%s", i)
          if (description.includes(limitEffect)) {
            description = description.replaceAll(limitEffect, String(Math.abs(spell["effect_limit_value_" + i])))
          }
          const maxEffect = util.format("@%s", i)
          if (description.includes(maxEffect)) {
            description = description.replaceAll(maxEffect, String(Math.abs(spell["max_" + i])))
          }

          if (description.includes("%z")) {
            description = description.replaceAll("%z",
              util.format("(%s ticks)", spell["buffduration"])
            )
          }
        }

        return description;
      }
    }
  };

  public static async renderSpellMini(parentSpellId, renderSpellId) {
    let spell = <any>this.getSpellData(renderSpellId)
    if (!this.getSpellData(renderSpellId)) {
      spell = <any>(await this.getSpell(renderSpellId));
      this.setSpellData(renderSpellId, spell);
    }

    const targetTypeColor = this.getTargetTypeColor(spell["targettype"]);

    return `
          <div :id="${parentSpellId} + '-' + ${renderSpellId} + '-' + componentId" style="display:inline-block" class="ml-1">

            <div style="display: inline-block">
              <img
                :src="spellCdnUrl + '' + ${spell.new_icon} + '.gif'"
                style="height:15px; width:auto; border: .5px solid ${targetTypeColor}; border-radius: 2px;"
                >
              <span style="color: #f7ff00">${spell.name}</span>
            </div>

          </div>

          <b-popover
            :target="${parentSpellId} + '-' + ${renderSpellId} + '-' + componentId"
            placement="auto"
            custom-class="no-bg"
            delay="1"
            triggers="hover focus"
            style="width: 500px !important"
          >
            <eq-window style="margin-right: 10px; width: auto; height: 90%">
                <eq-spell-preview :spell-data="sideLoadedSpellData[${renderSpellId}]"/>
            </eq-window>
          </b-popover>`;
  };

  public static setSpellData(spellId, spell: any) {
    this.data[spellId] = spell;
  }

  public static getSpellData(spellId) {
    return this.data[spellId]
  }

  public static async preloadDbstr() {
    const api   = (new DbStrApi(SpireApiClient.getOpenApiConfig()))
    let filters = [
      ["type", "__", 6]
    ]

    if (this.dbstrPreloaded) {
      return;
    }


    let wheres = [];
    filters.forEach((filter) => {
      // @ts-ignore
      wheres.push(util.format("%s%s%s", filter[0], filter[1], filter[2]))
    })

    const result = await api.listDbStrs({where: wheres.join(".")});
    if (result.status === 200) {
      if (result.data && result.data.length > 0) {
        for (let index in result.data) {
          const row = result.data[index]
          this.setDbstrData(row.id, row.value)
        }

        this.dbstrPreloaded = true
      }
    }
  }

  public static setDbstrData(id, message: any) {
    this.dbstrData[id] = message;
  }

  public static async getDbstrData(id) {
    // return nothing if we're preloaded
    if (this.dbstrPreloaded && !this.dbstrData[id]) {
      return ""
    }

    if (this.dbstrData[id]) {
      return this.dbstrData[id]
    }

    const api   = (new DbStrApi(SpireApiClient.getOpenApiConfig()))
    let filters = [
      ["type", "__", 6],
      ["id", "__", id]
    ]

    let wheres = [];
    filters.forEach((filter) => {
      // @ts-ignore
      wheres.push(util.format("%s%s%s", filter[0], filter[1], filter[2]))
    })

    const result = await api.listDbStrs({where: wheres.join(".")});
    if (result.status === 200) {
      if (result.data && result.data.length > 0) {
        this.dbstrData[id] = <string>result.data[0].value;
      }
    }
  }


}
