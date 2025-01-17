<template>
  <div class="d-inline-block">

    <div v-if="e.event_type_id === PLAYER_EVENT.TASK_UPDATE">
      Progressed task update <span class="font-weight-bold">{{ event(e).task_name }}</span> ({{ event(e).task_id }})
      <span class="font-weight-bold">activity_id</span> {{ event(e).activity_id }}
      <span class="font-weight-bold">done_count</span> {{ event(e).done_count }}
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.TASK_COMPLETE">
      Completed task <span class="font-weight-bold">{{ event(e).task_name }}</span> ({{ event(e).task_id }})
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.ITEM_CREATION">
      Server created item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
      ({{ event(e).charges }})
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.TRADER_PURCHASE">
      Purchased item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
      ({{ event(e).charges }})
      for
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).price)"
      />
      from trader <span class="font-weight-bold">{{ event(e).trader_name }}</span>
      remaining player balance
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).player_money_balance)"
      />
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.TRADER_SELL">
      Sold item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
      ({{ event(e).charges }})
      for
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).price)"
      />

      as a trader to <span class="font-weight-bold">{{ event(e).buyer_name }}</span>

      from trader
      remaining player balance
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).player_money_balance)"
      />
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.SPLIT_MONEY">
      Split
      <eq-cash-display
        class="font-weight-bold"
        :price="calcMoney(event(e))"
      />
      remaining player balance
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).player_money_balance)"
      />
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.GM_COMMAND">
      Used GM command <span class="font-weight-bold">{{ event(e).message }}</span>
      <span v-if="event(e).target && event(e).target !== 'NONE'">
        using target [<span class="font-weight-bold">{{ event(e).target }}</span>]
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.SAY">
      Said <span class="font-weight-bold">{{ event(e).message }}</span>
      <span v-if="event(e).target && event(e).target !== 'NONE'">
        using target [<span class="font-weight-bold">{{ event(e).target }}</span>]</span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.WENT_OFFLINE">Went offline</div>
    <div v-else-if="e.event_type_id === PLAYER_EVENT.WENT_ONLINE">Went online</div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.ITEM_DESTROY">
      Destroyed item
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }}) ({{ event(e).reason }})
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.SKILL_UP">
      Increased Skill <span class="font-weight-bold">{{ DB_SKILLS[event(e).skill_id] }}</span>
      ({{ event(e).value }}/{{ event(e).max_skill }})
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.FISH_FAILURE">
      Failed to <span class="item-749-sm"></span> fish!
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.FISH_SUCCESS">We caught a lunker!
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 d-inline-block"/>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.FORAGE_FAILURE">
      Failed to forage!
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.FORAGE_SUCCESS">
      We found something!
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 d-inline-block"/>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.COMBINE_FAILURE">
      Failed to combine recipe <span class="font-weight-bold">{{ event(e).recipe_name }}</span>
      ({{ event(e).recipe_id }})
      <span v-if="TRADESKILLS[event(e).tradeskill_id]">
        tradeskill
        <span class="font-weight-bold">{{ TRADESKILLS[event(e).tradeskill_id] }}</span>
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.COMBINE_SUCCESS">
      Successfully combined recipe <span class="font-weight-bold">{{ event(e).recipe_name }}</span>
      ({{ event(e).recipe_id }})
      <span v-if="TRADESKILLS[event(e).tradeskill_id]">
        tradeskill
        <span class="font-weight-bold">{{ TRADESKILLS[event(e).tradeskill_id] }}</span>
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.DROPPED_ITEM">
      Dropped item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
      ({{ event(e).charges }})
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.DISCOVER_ITEM">
      Discover item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.REZ_ACCEPTED">
      Accepted resurrection
    </div>

    <div
      v-else-if="e.event_type_id === PLAYER_EVENT.MERCHANT_SELL"
    >
      Sold
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      to
      <npc-popover :npc="npcData[event(e).npc_id]" :show-image="false" class="d-inline-block font-weight-bold"/>
      for
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).cost)"
        v-if="parseInt(event(e).alternate_currency_id) === 0"
      />
      player balance
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).player_money_balance)"
        v-if="parseInt(event(e).alternate_currency_id) === 0"
      />
    </div>

    <div
      v-else-if="e.event_type_id === PLAYER_EVENT.POSSIBLE_HACK"
      class="d-inline-block"
    >
      {{ event(e).message.replace(/^\s+|\s+$/g, '') }}
    </div>

    <div
      v-else-if="[PLAYER_EVENT.KILLED_NPC, PLAYER_EVENT.KILLED_NAMED_NPC, PLAYER_EVENT.KILLED_RAID_NPC].includes(e.event_type_id)"
    >
      Killed
      <npc-popover :npc="npcData[event(e).npc_id]" :show-image="false" class="d-inline-block font-weight-bold"/>
    </div>

    <div
      v-else-if="e.event_type_id === PLAYER_EVENT.MERCHANT_PURCHASE"
    >
      Bought
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      from
      <npc-popover :npc="npcData[event(e).npc_id]" :show-image="false" class="d-inline-block font-weight-bold"/>
      for
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).cost)"
        v-if="parseInt(event(e).alternate_currency_id) === 0"
      />
      player balance
      <eq-cash-display
        class="font-weight-bold"
        :price="parseInt(event(e).player_money_balance)"
        v-if="parseInt(event(e).alternate_currency_id) === 0"
      />
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.ZONING">
      Zoned from
      <span class="font-weight-bold" v-if="zoneData[event(e).from_zone_id]">
      {{ zoneData[event(e).from_zone_id].long_name }}
      </span>
      to
      <span class="font-weight-bold" v-if="zoneData[event(e).to_zone_id]">
      {{ zoneData[event(e).to_zone_id].long_name }}
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.LEVEL_GAIN">
      Leveled up (+{{ event(e).levels_gained }}) to
      <span class="font-weight-bold">
      {{ event(e).to_level }}
      </span>
      from level
      <span class="font-weight-bold">
      {{ event(e).from_level }}
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.LEVEL_LOSS">
      Leveled down (-{{ event(e).levels_lost }}) to
      <span class="font-weight-bold">
      {{ event(e).to_level }}
      </span>
      from level
      <span class="font-weight-bold">
      {{ event(e).from_level }}
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.AA_GAIN">
      Gained <span class="font-weight-bold">{{ event(e).aa_gained }}</span> AA point(s)
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.AA_PURCHASE">
      Purchased AA <span class="font-weight-bold">{{ getAADescription(event(e).aa_id) }}</span>

      <span v-if="aaData[event(e).aa_id] && aaData[event(e).aa_id].spells_new">
      (<spell-popover
        v-if="aaData[event(e).aa_id] && aaData[event(e).aa_id].spells_new"
        class="ml-1 font-weight-bold"
        size="20"
        :spell="aaData[event(e).aa_id].spells_new"
      />)
        </span>
      at cost
      <span class="font-weight-bold">{{ event(e).aa_cost }}</span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.LOOT_ITEM">
      Looted
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }})
      from
      <npc-popover :npc="npcData[event(e).npc_id]" :show-image="false" class="d-inline-block font-weight-bold"/>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.GROUNDSPAWN_PICKUP">
      Picked up
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }})
      from the ground!
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.TASK_ACCEPT">
      Accepted task <span class="font-weight-bold">{{ event(e).task_name }}</span> ({{ event(e).task_id }})
      <span v-if="event(e).npc_id">
      from
        <npc-popover
          :npc="npcData[event(e).npc_id]"
          :show-image="false"
          class="d-inline-block font-weight-bold"
          v-if="event(e).npc_id"
        />
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.DEATH">
      Died
      <span v-if="event(e).killer_id">
      by
        <span
          class="font-weight-bold"
          v-if="event(e).killer_id && !npcData[event(e).killer_id]"
        >{{ event(e).killer_name }}
        </span>
      </span>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.NPC_HANDIN">

      <div v-if="!expandedEvent[e.id]" class="d-inline-block">
        Handed in ({{ event(e).handin_items.length }}) items
        <eq-cash-display
          class="font-weight-bold"
          :price="calcMoney(event(e).handin_money)"
          v-if="calcMoney(event(e).handin_money) > 0"
        />
        to <span class="font-weight-bold">{{ event(e).npc_name }}</span>

        Returned ({{ event(e).return_items.length }}) item(s)
      </div>

      <div v-if="expandedEvent[e.id]" style="min-width: 600px" class="eq-window-simple p-3 mb-3">
        <div class="row">
          <div class="col-3"></div>
          <div class="col-6">
            <npc-popover
              :npc="npcData[event(e).npc_id]"
              class="d-inline-block font-weight-bold"
            />


            <!--            <span class="font-weight-bold">{{ npcData[event(e).npc_id] }}</span>-->
          </div>
        </div>
        <div class="row mt-3">
          <div class="col-6 text-center">
            <span class="font-weight-bold">Handed In</span>
            <hr class="mt-3 mb-3" style="border-top-color: rgba(255, 255, 255, 0.14);">
            <div
              class="text-left"
              v-for="i in event(e).handin_items"
            >
              <item-popover
                :item="itemData[i.item_id]"
                :class="'font-weight-bold d-inline-block'"
              />
              (x{{ i.charges ? i.charges : 1 }})
            </div>

            <eq-cash-display
              class="font-weight-bold mr-0 pr-0 mt-3"
              :price="calcMoney(event(e).handin_money)"
              v-if="calcMoney(event(e).handin_money) > 0"
            />
          </div>
          <div class="col-6 text-center">
            <span class="font-weight-bold">Returned</span>
            <hr class="mt-3 mb-3" style="border-top-color: rgba(255, 255, 255, 0.14);">
            <div
              class="text-left"
              v-for="i in event(e).return_items"
            >
              <item-popover
                :item="itemData[i.item_id]"
                :class="' font-weight-bold d-inline-block'"
              />
              (x{{ i.charges ? i.charges : 1 }})
            </div>

          </div>
        </div>

      </div>

      <button
        title="Show Detail"
        @click="expandEvent(e)"
        v-if="!expandedEvent[e.id] && Object.keys(JSON.parse(e.event_data)).length > 0"
        class="ml-3 btn btn-sm btn-warning"
        style="font-size: 10px; "
      >
        <i class="fa fa-plus"></i>
      </button>
    </div>

    <div v-else-if="e.event_type_id === PLAYER_EVENT.TRADE">

      <div v-if="!expandedEvent[e.id]" class="d-inline-block">
        <span class="font-weight-bold" v-if="characterData[event(e).character_2_id]">
          {{ characterData[event(e).character_2_id].name }}
        </span> traded ({{ event(e).character_1_give_items.length }}) items
        <eq-cash-display
          class="font-weight-bold"
          :price="calcMoney(event(e).character_1_give_money)"
          v-if="calcMoney(event(e).character_1_give_money) > 0"
        />
        <>
        <span class="font-weight-bold" v-if="characterData[event(e).character_1_id]">
          {{ characterData[event(e).character_1_id].name }}
        </span>
        traded ({{ event(e).character_2_give_items.length }}) items
        <eq-cash-display
          class="font-weight-bold mr-0 pr-0"
          :price="calcMoney(event(e).character_2_give_money)"
          v-if="calcMoney(event(e).character_2_give_money) > 0"
        />
      </div>

      <div v-if="expandedEvent[e.id]" style="min-width: 600px" class="eq-window-simple mb-3">
        <div class="row mt-3">
          <div class="col-6 text-center">
            <span class="font-weight-bold">{{ characterData[event(e).character_2_id].name }}</span>
            <hr class="mt-3 mb-3" style="border-top-color: rgba(255, 255, 255, 0.14);">
            <div
              class="text-left"
              v-for="i in event(e).character_1_give_items"
            >
              <item-popover
                :item="itemData[i.item_id]"
                :class="(i.in_bag ? 'ml-3' : 'ml-0') + ' font-weight-bold d-inline-block'"
              />
              (x{{ i.charges ? i.charges : 1 }})
            </div>

            <hr class="mt-3 mb-3" style="border-top-color: rgba(255, 255, 255, 0.14);">

            <eq-cash-display
              class="font-weight-bold mr-0 pr-0"
              :price="calcMoney(event(e).character_1_give_money)"
              v-if="calcMoney(event(e).character_1_give_money) > 0"
            />
          </div>
          <div class="col-6 text-center">
            <span class="font-weight-bold">{{ characterData[event(e).character_1_id].name }}</span>
            <hr class="mt-3 mb-3" style="border-top-color: rgba(255, 255, 255, 0.14);">
            <div
              class="text-left"
              v-for="i in event(e).character_2_give_items"
            >
              <item-popover
                :item="itemData[i.item_id]"
                :class="(i.in_bag ? 'ml-3' : 'ml-0') + ' font-weight-bold d-inline-block'"
              />
              (x{{ i.charges ? i.charges : 1 }})
            </div>

            <hr class="mt-3 mb-3" style="border-top-color: rgba(255, 255, 255, 0.14);">

            <eq-cash-display
              class="font-weight-bold mr-0 pr-0"
              :price="calcMoney(event(e).character_2_give_money)"
              v-if="calcMoney(event(e).character_2_give_money) > 0"
            />
          </div>
        </div>

      </div>

      <button
        title="Show Detail"
        @click="expandEvent(e)"
        v-if="!expandedEvent[e.id] && Object.keys(JSON.parse(e.event_data)).length > 0"
        class="ml-3 btn btn-sm btn-warning"
        style="font-size: 10px; "
      >
        <i class="fa fa-plus"></i>
      </button>

    </div>

    <div v-else>
      Event {{ e.event_type_id }} unimplemented
    </div>
  </div>
</template>

<script>
import {PLAYER_EVENT} from "@/views/admin/player-event-logs/player-events";
import {Items}        from "@/app/items";
import ItemPopover    from "@/components/ItemPopover.vue";
import {DB_SKILLS}    from "@/app/constants/eq-skill-constants";
import {Npcs}         from "@/app/npcs";
import NpcPopover     from "@/components/NpcPopover.vue";
import EqCashDisplay  from "@/components/eq-ui/EqCashDisplay.vue";
import {Zones}        from "@/app/zones";
import {AA}           from "@/app/aa";
import SpellPopover   from "@/components/SpellPopover.vue";
import {TRADESKILLS}  from "@/app/constants/eq-tradeskill-constants";
import {Characters}   from "@/app/characters";

export default {
  name: "PlayerEventDisplayComponent",
  components: { SpellPopover, EqCashDisplay, NpcPopover, ItemPopover },
  props: {
    e: Object, // event
  },
  data() {
    return {
      PLAYER_EVENT: PLAYER_EVENT,
      DB_SKILLS: DB_SKILLS,
      TRADESKILLS: TRADESKILLS,

      itemData: {},
      npcData: {},
      zoneData: {},
      aaData: {},
      characterData: {},

      expandedEvent: {}
    }
  },
  async mounted() {
    const p = this.event(this.e)
    if (p.item_id && p.item_id > 0) {
      this.itemData[p.item_id] = await Items.getItem(p.item_id)
    }
    if (p.npc_id && p.npc_id > 0) {
      this.npcData[p.npc_id] = await Npcs.getNpc(p.npc_id)
    }
    if (p.from_zone_id && p.from_zone_id > 0) {
      this.zoneData[p.from_zone_id] = await Zones.getZoneById(p.from_zone_id)
    }
    if (p.to_zone_id && p.to_zone_id > 0) {
      this.zoneData[p.to_zone_id] = await Zones.getZoneById(p.to_zone_id)
    }
    if (p.aa_id && p.aa_id > 0) {
      this.aaData[p.aa_id] = await AA.getAARankByRankId(p.aa_id)
    }
    if (p.character_1_id && p.character_1_id > 0) {
      this.characterData[p.character_1_id] = await Characters.get(p.character_1_id)
    }
    if (p.character_2_id && p.character_2_id > 0) {
      this.characterData[p.character_2_id] = await Characters.get(p.character_2_id)
    }
    if (p && p.character_1_give_items && p.character_1_give_items.length > 0) {
      for (let i of p.character_1_give_items) {
        this.itemData[i.item_id] = await Items.getItem(i.item_id)
      }
    }
    if (p && p.character_2_give_items && p.character_2_give_items.length > 0) {
      for (let i of p.character_2_give_items) {
        this.itemData[i.item_id] = await Items.getItem(i.item_id)
      }
    }
    if (p && p.handin_items && p.handin_items.length > 0) {
      for (let i of p.handin_items) {
        this.itemData[i.item_id] = await Items.getItem(i.item_id)
      }
    }

    this.$forceUpdate()
  },
  methods: {

    expandEvent(e) {
      this.expandedEvent[e.id] = 1
      this.$forceUpdate()
    },

    calcMoney(m) {
      let copper = m.copper
      copper += (m.platinum * 1000)
      copper += (m.gold * 100)
      copper += (m.silver * 10)
      return copper
    },

    getAADescription(rankId) {
      const r = AA.getAARankByRankId(rankId)
      if (r && r.id > 0) {
        return AA.getAANameDbString(r.title_sid).value
      }

      return ""
    },
    async getItem(itemId) {
      return await Items.getItem(itemId)
    },
    event(e) {
      return JSON.parse(e.event_data)
    },
  }
}
</script>

<style scoped>

</style>
