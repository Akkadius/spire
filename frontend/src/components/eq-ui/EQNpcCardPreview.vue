<template>
  <div>

    <!-- Race Model -->
    <span
      style="position: absolute; right: 7%; opacity: .8"
      :class="'race-models-ctn-' + npc.race + '-' + npc.gender + '-' + npc.texture + '-' + npc.helmtexture"
    />

    <!-- Weapon Model 1 -->
    <span
      v-if="npc.d_melee_texture_1 > 0"
      :style="'position: absolute; right: ' + (npc.d_melee_texture_2 > 0 ? 25 : 20) + '%; top: 5%; opacity: 1; z-index: 9999'"
      :class="'mt-2 mb-2 object-ctn-' + npc.d_melee_texture_1"
    />

    <!-- Weapon Model 2 -->
    <span
      v-if="npc.d_melee_texture_2 > 0"
      style="position: absolute; right: 20%; top: 5%; opacity: .8"
      :class="'mt-2 mb-2 object-ctn-' + npc.d_melee_texture_2"
    />

    <div class="row">
      <div class="col-8 pl-5">
        <h6 class="eq-header" style="margin: 0px; margin-bottom: 10px">
          {{ getCleanName() }} {{ (npc.lastname && npc.lastname.length > 0 ? "(" + npc.lastname + ")" : "") }}
        </h6>
      </div>
      <div class="col-12">
        <div class="row">
          <div class="col-2 text-right font-weight-bold pr-0">
            ID
          </div>
          <div class="col-4 pl-3">
            {{ npc.id }}
          </div>
        </div>
        <div class="row">
          <div class="col-2 text-right font-weight-bold pr-0">
            Level
          </div>
          <div class="col-4 pl-3">
            {{ npc.level }}
          </div>
        </div>
        <div class="row">
          <div class="col-2 text-right font-weight-bold pr-0">
            Class
          </div>
          <div class="col-4 pl-3">
            <span
              v-if="getClassIcon().toString().length > 0"
              :class="'item-' + getClassIcon() + '-sm'"
              style="display: inline-block"
            />
            <span style="position: relative; top: -2px;">
              {{ getClassName(npc.class) }} ({{ npc.class }})
            </span>
          </div>
        </div>
        <div class="row">
          <div class="col-2 text-right font-weight-bold pr-0">
            Race
          </div>
          <div class="col-4 pl-3">
            <span
              v-if="getRaceIcon().toString().length > 0"
              :class="'item-' + getRaceIcon() + '-sm'"
              style="display: inline-block"
            />
            <span style="position: relative; top: -2px;">
              {{ getRaceName(npc.race) }} ({{ npc.race }})
            </span>
          </div>
        </div>
        <div class="row">
          <div class="col-2 text-right font-weight-bold pr-0">
            Bodytype
          </div>
          <div class="col-4 pl-3">
            {{ getBodytype() }} ({{ npc.bodytype }})
          </div>
        </div>
        <div class="row" v-if="npc.armortint_red !== 0 || npc.armortint_green !== 0 || npc.armortint_blue !== 0">
          <div class="col-2 text-right font-weight-bold pr-0">
            Armor RGB
          </div>
          <div class="col-4 pl-3">
            <div
              class="mr-3"
              :style="'width: 15px; height: 15px; margin-top: 3px; border-radius: 5px; background-color: rgb(' + npc.armortint_red + ', ' + npc.armortint_green + ', ' + npc.armortint_blue + ');'"
            ></div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-4 mb-5">
      <div
        class="row mt-3"
        v-for="(row, index) in rows"
        :key="index"
      >
        <div
          class="col-4 "
          v-for="(column, index) in row.columns"
          :key="index"
        >
          <div
            class="row"
            v-for="field in column"
            :key="field.label"
            v-if="((typeof field.showIf !== 'undefined' && field.showIf) || typeof field.showIf === 'undefined')"
          >
            <div class="col-8 text-left font-weight-bold pr-0">
              <span
                v-if="field.icon"
                :class="'item-' + field.icon + '-sm'"
                style="display: inline-block"
              />
              <span style="position: relative; top: -2px;">
            {{ field.label }}
            </span>

            </div>
            <div class="col-4 pl-1" style="position: relative; top: -2px;">
              {{ (parseInt(field.value) > 0 ? commify(field.value) : field.value) }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Special Abilities -->
    <div class="mt-3" v-if="npc.special_abilities.length > 0">
      <div class="font-weight-bold mb-3">
        This NPC has the following special abilities ({{ parseSpecialAbilities(npc.special_abilities).length }})
      </div>

      <ul style="padding-left: 20px;">
        <li
          style="display: list-item; list-style-type: circle;"
          v-for="a in parseSpecialAbilities(npc.special_abilities).sort((a, b) => a.localeCompare(b))"
        >
          {{ a }}
        </li>
      </ul>
    </div>

    <!-- Merchant -->
    <div v-if="npc.merchant_id > 0" class="mt-3">

      <!-- Show if under max -->
      <div v-if="merchantitems && merchantitems.length > 0 && (merchantitems.length < maxDataEntries || showMerchantItems)" class="fade-in">
        <div class="font-weight-bold mb-3">This NPC sells the following items ({{ commify(merchantitems.length) }})
          <a href="javascript:void(0);" @click="showMerchantItems = false" v-if="merchantitems.length > maxDataEntries">hide</a>
        </div>

        <div v-for="e in merchantitems" class="row">
          <div class="col-6">
            <item-popover
              class="d-inline-block"
              :item="e.item"
              v-if="Object.keys(e.item).length > 0 && e.item"
              size="sm"
            />
          </div>

          <div class="col-3">
            <eq-cash-display
              class="d-inline-block ml-1"
              :price="e.item.price"
            />
          </div>

        </div>
      </div>

      <!-- Prompt if over max -->
      <div
        v-if="merchantitems && merchantitems.length > 0 &&  (merchantitems.length > maxDataEntries && !showMerchantItems)"
        class="font-weight-bold"
      >
        This NPC sells ({{ commify(merchantitems.length) }}) items, click <a
        href="javascript:void(0);"
        @click="showMerchantItems = true"
      >here</a>
        to view them all
      </div>

    </div>

    <!-- Casted Spells -->
    <div v-if="npc.npc_spells_id > 0" class="mt-3">

      <!-- Show if under max -->
      <div
        v-if="castedSpells && castedSpells.length > 0 && (castedSpells.length < maxDataEntries || showCastedSpells)"
        class="fade-in"
      >
        <div class="font-weight-bold mb-3">This NPC casts the following spells ({{ commify(castedSpells.length) }})
          ({{ npc.npc_spell.name }})
          <a href="javascript:void(0);" @click="showCastedSpells = false" v-if="castedSpells.length > maxDataEntries">hide</a>
        </div>

        <div v-for="e in castedSpells">
          <spell-popover
            :spell="e.spell"
            :size="20"
            :spell-name-length="25"
            v-if="Object.keys(e.spell).length > 0 && e.spell"
            class="mt-2"
          />
        </div>
      </div>

      <!-- Prompt if over max -->
      <div
        v-if="castedSpells && castedSpells.length > 0 && (castedSpells.length > maxDataEntries && !showCastedSpells)"
        class="font-weight-bold"
      >
        This NPC casts ({{ commify(castedSpells.length) }}) spells, click <a
        href="javascript:void(0);"
        @click="showCastedSpells = true"
      >here</a>
        to view them all
      </div>

    </div>

    <!-- Faction -->
    <div v-if="npc.npc_spells_id > 0" class="mt-3">

      <!-- Show if under max -->
      <div
        v-if="factionHits && factionHits.length > 0 && (factionHits.length < maxDataEntries || showFactionHits)"
        class="fade-in"
      >
        <div class="font-weight-bold mb-3">This NPC has the following faction hits ({{ commify(factionHits.length) }})
          <a
            href="javascript:void(0);"
            @click="showFactionHits = false"
            v-if="factionHits.length > maxDataEntries"
          >hide</a>
        </div>

        <ul style="padding-left: 20px;">
          <li
            style="display: list-item; list-style-type: circle;"
            v-for="e in factionHits"
          >
            {{ e.name }} ({{ e.hitValue }})
          </li>
        </ul>

      </div>

      <!-- Prompt if over max -->
      <div
        v-if="factionHits && factionHits.length > 0 && (factionHits.length > maxDataEntries && !showFactionHits)"
        class="font-weight-bold"
      >
        This NPC has ({{ commify(factionHits.length) }}) faction hits, click <a
        href="javascript:void(0);"
        @click="showFactionHits = true"
      >here</a>to view them all
      </div>

    </div>

    <!-- Loot -->
    <div v-if="npc.loottable_id > 0" class="mt-3">

      <!-- Show if under max -->
      <div v-if="loot && loot.length > 0 && (loot.length < maxDataEntries || showLoot)" class="fade-in">
        <div class="font-weight-bold mb-3">This NPC drops the following items ({{ commify(loot.length) }})
          <a href="javascript:void(0);" @click="showLoot = false" v-if="loot.length > maxDataEntries">hide</a>
        </div>

        <div v-for="e in loot">
          <item-popover
            class="d-inline-block"
            :item="e.item"
            v-if="Object.keys(e.item).length > 0 && e.item"
            size="sm"
          />
        </div>
      </div>

      <!-- Prompt if over max -->
      <div v-if="loot && loot.length > 0 && (loot.length > maxDataEntries && !showLoot)" class="font-weight-bold">
        This NPC drops ({{ commify(loot.length) }}) items, click <a href="javascript:void(0);" @click="showLoot = true">here</a>
        to view them all
      </div>

    </div>

    <eq-debug :data="npc"/>
  </div>
</template>

<script>
import EqDebug                             from "./EQDebug";
import {Npcs}                              from "../../app/npcs";
import {DB_PLAYER_RACES, DB_RACE_NAMES}    from "../../app/constants/eq-races-constants";
import {DB_CLASSES, DB_PLAYER_CLASSES_ALL} from "../../app/constants/eq-classes-constants";
import {BODYTYPES}                         from "../../app/constants/eq-bodytype-constants";
import {FLYMODE}                           from "../../app/constants/eq-flymode-constants";
import ItemPopover                         from "../ItemPopover";
import {Items}                             from "../../app/items";
import EqCashDisplay                       from "./EqCashDisplay";
import SpellPopover                        from "../SpellPopover";

export default {
  name: "EqNpcCardPreview",
  components: { SpellPopover, EqCashDisplay, ItemPopover, EqDebug },
  data() {
    return {
      maxDataEntries: 10, // amount of results before collapsing

      // loot
      showLoot: false, // when too many results shown, toggle
      loot: {}, // data

      // merchant
      merchantitems: [], // data
      showMerchantItems: false, // when too many results shown, toggle

      // spells
      castedSpells: [], // data
      showCastedSpells: false,  // when too many results shown, toggle

      // faction hits
      factionHits: [], // data
      showFactionHits: false,  // when too many results shown, toggle

      // field display
      rows: [
        {
          columns: [
            [
              // basic
              // { icon: '4487', label: "Class", value: this.npc["class"] },
              // { icon: '4466', label: "Race", value: this.getRaceName(this.npc.race) + ` (${this.npc["race"]})` },
              // { icon: '2072', label: "Level", value: this.npc["level"] },

              // stats
              { icon: '2052', label: "HP", value: this.npc['hp'], showIf: this.npc['hp'] > 0 },
              {
                icon: '2052',
                label: "(Sec) HP Regen",
                value: this.npc['hp_regen_per_second'],
                showIf: this.npc['hp_regen_per_second'] > 0
              },
              {
                icon: '2052',
                label: "(Tic) HP Regen",
                value: this.npc['hp_regen_rate'],
                showIf: this.npc['hp_regen_rate'] > 0
              },
              { icon: '2057', label: "Mana", value: this.npc['mana'], showIf: this.npc['mana'] > 0 },
              {
                icon: '2057',
                label: "(Tic) Mana Regen",
                value: this.npc['mana_regen_rate'],
                showIf: this.npc['mana_regen_rate'] > 0
              },
            ],

            [
              // stats
              { icon: '5388', label: "Armor Class", value: this.npc['ac'], showIf: this.npc['ac'] > 0 },
              { icon: '1959', label: "Agility", value: this.npc['agi'], showIf: this.npc['agi'] > 0 },
              { icon: '1953', label: "Dexterity", value: this.npc['dex'], showIf: this.npc['dex'] > 0 },
              { icon: '1947', label: "Stamina", value: this.npc['sta'], showIf: this.npc['sta'] > 0 },
              { icon: '1941', label: "Strength", value: this.npc['str'], showIf: this.npc['str'] > 0 },
              { icon: '1965', label: "Charisma", value: this.npc['cha'], showIf: this.npc['cha'] > 0 },
              { icon: '1971', label: "Wisdom", value: this.npc['wis'], showIf: this.npc['wis'] > 0 },
              { icon: '1977', label: "Intelligence", value: this.npc['_int'], showIf: this.npc['_int'] > 0 },
            ],

            [
              // resists
              { icon: '2987', label: "Cold Resist", value: this.npc['cr'], showIf: this.npc['cr'] > 0 },
              { icon: '2994', label: "Disease Resist", value: this.npc['dr'], showIf: this.npc['dr'] > 0 },
              { icon: '2988', label: "Fire Resist", value: this.npc['fr'], showIf: this.npc['fr'] > 0 },
              { icon: '2984', label: "Magic Resist", value: this.npc['mr'], showIf: this.npc['mr'] > 0 },
              { icon: '2986', label: "Poison Resist", value: this.npc['pr'], showIf: this.npc['pr'] > 0 },
              { icon: '2995', label: "Corrup. Resist", value: this.npc['corrup'], showIf: this.npc['corrup'] > 0 },
              { icon: '2982', label: "Physical Resist", value: this.npc['ph_r'], showIf: this.npc['ph_r'] > 0 },

            ],
          ],
        },
        {
          columns: [
            [
              {
                icon: '',
                label: 'See Hide',
                value: this.npc['see_hide'] ? "Yes" : "No",
                showIf: this.npc['see_hide'] > 0
              },
              {
                icon: '',
                label: 'See Improved Hide',
                value: this.npc['see_improved_hide'] ? "Yes" : "No",
                showIf: this.npc['see_improved_hide'] > 0
              },
              {
                icon: '',
                label: 'See Invisible',
                value: this.npc['see_invis'] ? "Yes" : "No",
                showIf: this.npc['see_invis'] > 0
              },
              {
                icon: '',
                label: 'See Invis Undead',
                value: this.npc['see_invis_undead'] ? "Yes" : "No",
                showIf: this.npc['see_invis_undead'] > 0
              },
              {
                icon: '',
                label: 'Show Name',
                value: this.npc['show_name'] ? "Yes" : "No",
                showIf: this.npc['show_name'] > 0
              },
              {
                icon: '',
                label: 'Skip Global Loot',
                value: this.npc['skip_global_loot'] ? "Yes" : "No",
                showIf: this.npc['skip_global_loot'] > 0
              },
              { icon: '', label: 'Spawn Limit', value: this.npc['spawn_limit'], showIf: this.npc['spawn_limit'] > 0 },
              {
                icon: '',
                label: 'Trackable',
                value: this.npc['trackable'] ? "Yes" : "No",
                showIf: this.npc['trackable'] > 0
              },
              {
                icon: '',
                label: 'Stuck Behavior',
                value: this.npc['stuck_behavior'],
                showIf: this.npc['stuck_behavior'] > 0
              },
              {
                icon: '',
                label: 'No Target Hotkey',
                value: this.npc['no_target_hotkey'] ? "Yes" : "No",
                showIf: this.npc['no_target_hotkey'] > 0
              },
              {
                icon: '',
                label: 'Flymode',
                value: FLYMODE[this.npc['flymode']] ? FLYMODE[this.npc['flymode']] : "",
                showIf: this.npc['flymode'] > 0
              },
              {
                icon: '',
                label: 'Findable',
                value: this.npc['findable'] ? "Yes" : "No",
                showIf: this.npc['findable'] > 0
              },
              {
                icon: '',
                label: 'Untargetable',
                value: this.npc['untargetable'] ? "Yes" : "No",
                showIf: this.npc['untargetable'] > 0
              },
              {
                icon: '',
                label: 'Underwater',
                value: this.npc['underwater'] ? "Yes" : "No",
                showIf: this.npc['underwater'] > 0
              },
              {
                icon: '',
                label: 'Unique Spawn By Name',
                value: this.npc['unique_spawn_by_name'] ? "Yes" : "No",
                showIf: this.npc['unique_spawn_by_name'] > 0
              },
              {
                icon: '',
                label: 'QGlobals Enabled',
                value: this.npc['qglobal'] ? "Yes" : "No",
                showIf: this.npc['qglobal'] > 0
              },
              {
                icon: '',
                label: 'Ignore Despawn',
                value: this.npc['ignore_despawn'] ? "Yes" : "No",
                showIf: this.npc['ignore_despawn'] > 0
              },
              {
                icon: '',
                label: 'Experience Modifier',
                value: this.npc['exp_mod'],
                showIf: this.npc['exp_mod'] !== 100
              },
              { icon: '', label: 'Bot NPC', value: this.npc['isbot'] ? "Yes" : "No", showIf: this.npc['isbot'] > 0 },
              {
                icon: '',
                label: 'Quest NPC',
                value: this.npc['isquest'] ? "Yes" : "No",
                showIf: this.npc['isquest'] > 0
              },
            ],
            [
              //
            ],
            [
              // mods
              { icon: '2087', label: 'Min Damage', value: this.npc['mindmg'] },
              { icon: '2087', label: 'Max Damage', value: this.npc['maxdmg'] },
              {
                icon: '2090',
                label: 'Attack Count',
                value: this.npc['attack_count'],
                showIf: this.npc['attack_count'] !== -1
              },
              { icon: '2095', label: 'Attack Delay', value: this.npc['attack_delay'] },
              { icon: '2102', label: 'Attack Speed', value: this.npc['attack_speed'] },
              {
                icon: '778',
                label: 'Heal Scale',
                value: this.npc['healscale'] + '%',
                showIf: this.npc['healscale'] !== 100
              },
              {
                icon: '777',
                label: 'Spell Scale',
                value: this.npc['spellscale'] + '%',
                showIf: this.npc['spellscale'] !== 100
              },
              {
                icon: '778',
                label: 'Scale Rate',
                value: this.npc['scalerate'] + '%',
                showIf: this.npc['scalerate'] !== 100
              },
              { icon: '590', label: "Attack", value: this.npc['atk'], showIf: this.npc['atk'] > 0 },
              { icon: '597', label: "Accuracy", value: this.npc['accuracy'], showIf: this.npc['accuracy'] > 0 },
              { icon: '3656', label: "Avoidance", value: this.npc['avoidance'], showIf: this.npc['avoidance'] > 0 },

              {
                icon: '5388',
                label: 'Charm Armor Class',
                value: this.npc['charm_ac'],
                showIf: this.npc['charm_ac'] > 0
              },
              {
                icon: '597',
                label: 'Charm Accuracy',
                value: this.npc['charm_accuracy_rating'],
                showIf: this.npc['charm_accuracy_rating'] > 0
              },
              { icon: '590', label: 'Charm Attack', value: this.npc['charm_atk'], showIf: this.npc['charm_atk'] > 0 },
              {
                icon: '2095',
                label: 'Charm Attack Delay',
                value: this.npc['charm_attack_delay'],
                showIf: this.npc['charm_attack_delay'] > 0
              },
              {
                icon: '3656',
                label: 'Charm Avoidance',
                value: this.npc['charm_avoidance_rating'],
                showIf: this.npc['charm_avoidance_rating'] > 0
              },
              {
                icon: '2087',
                label: 'Charm Min Damage',
                value: this.npc['charm_min_dmg'],
                showIf: this.npc['charm_min_dmg'] > 0
              },
              {
                icon: '2087',
                label: 'Charm Max Damage',
                value: this.npc['charm_max_dmg'],
                showIf: this.npc['charm_max_dmg'] > 0
              },

            ]
          ]
        }
      ],

    }
  },
  async mounted() {

    // merchants
    if (this.npc.merchant_id > 0 && this.npc.merchantlists) {
      let merchantItems = []
      for (let listitem of this.npc.merchantlists) {
        merchantItems.push(
          {
            item: (await Items.getItem(listitem.item)),
            entry: listitem
          }
        )
      }
      this.merchantitems = merchantItems
    }

    // loot
    if (this.npc.loottable_id > 0 && this.npc.loottable.loottable_entries) {
      let lootItems = []
      for (let l of this.npc.loottable.loottable_entries) {
        // console.log(l)
        for (let e of l.lootdrop_entries) {

          // make sure we don't add the same item twice for now
          if (lootItems.filter(f => f.item.id === e.item.id).length === 0) {
            lootItems.push(
              {
                item: e.item,
              }
            )
          }
        }
      }

      // sort alpha by name
      this.loot = lootItems.sort((a, b) => {
        return a.item.name.localeCompare(b.item.name);
      });
    }

    // casted spells
    if (this.npc.npc_spells_id > 0 && this.npc.npc_spell) {
      let castedSpells = []
      for (let e of this.npc.npc_spell.npc_spells_entries) {

        // make sure we don't add the same spell twice for now
        if (castedSpells.filter(f => f.spell.id === e.spells_new.id).length === 0) {
          castedSpells.push(
            {
              spell: e.spells_new,
            }
          )
        }

      }

      // sort alpha by name
      this.castedSpells = castedSpells.sort((a, b) => {
        return a.spell.name.localeCompare(b.spell.name);
      });
    }

    // faction
    if (this.npc.npc_faction_id > 0 && this.npc.npc_factions) {
      let factionHits = []
      for (let n of this.npc.npc_factions) {
        if (n.npc_faction_entries) {
          for (let f of n.npc_faction_entries) {

            // make sure we don't add the same spell twice for now
            if (factionHits.filter(e => f.name === e.name).length === 0) {
              factionHits.push(
                {
                  name: f.faction_list.name,
                  hitValue: f.value,
                }
              )
            }
          }

        }
      }

      // sort alpha by name
      this.factionHits = factionHits.sort((a, b) => {
        return a.name.localeCompare(b.name);
      });

    }
  },
  props: {
    npc: {
      type: Object,
      default: {},
      required: true
    },
  },
  methods: {
    parseSpecialAbilities(abilities) {
      return Npcs.specialAbilitiesToHuman(abilities)
    },
    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },
    getCleanName() {
      return Npcs.getCleanName(this.npc.name)
    },
    getRaceName(raceId) {
      return DB_RACE_NAMES[raceId]
    },
    getClassName(classId) {
      return DB_CLASSES[classId]
    },
    getBodytype() {
      return BODYTYPES[this.npc.bodytype]
    },
    getClassIcon() {
      return DB_PLAYER_CLASSES_ALL[this.npc.class] ? DB_PLAYER_CLASSES_ALL[this.npc.class].icon : ""
    },
    getRaceIcon() {
      return DB_PLAYER_RACES[this.npc.race] ? DB_PLAYER_RACES[this.npc.race].icon : ""
    }
  }
}
</script>

<style scoped>

</style>
