<template>
  <div class="d-inline-block">
    <!--    {{ e.event_data }}-->

    <!--    {{ e.event_data }}-->

    <div v-if="e.event_type_id === PLAYER_EVENT.GM_COMMAND">
      Used GM command <span class="font-weight-bold">{{ event(e).message }}</span>
      <span v-if="event(e).target && event(e).target !== 'NONE'">
        using target [<span class="font-weight-bold">{{ event(e).target }}</span>]
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.SAY">
      Said <span class="font-weight-bold">{{ event(e).message }}</span>
      <span v-if="event(e).target && event(e).target !== 'NONE'">
        using target [<span class="font-weight-bold">{{ event(e).target }}</span>]</span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.WENT_OFFLINE">Went offline</div>
    <div v-if="e.event_type_id === PLAYER_EVENT.WENT_ONLINE">Went online</div>

    <div v-if="e.event_type_id === PLAYER_EVENT.ITEM_DESTROY">
      Destroyed item
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }}) ({{ event(e).reason }})
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.SKILL_UP">
      Increased Skill <span class="font-weight-bold">{{ DB_SKILLS[event(e).skill_id] }}</span>
      ({{ event(e).value }}/{{ event(e).max_skill }})
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.FISH_FAILURE">
      Failed to <span class="item-749-sm"></span> fish!
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.FISH_SUCCESS">We caught a lunker!
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 d-inline-block"/>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.FORAGE_FAILURE">
      Failed to forage!
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.FORAGE_SUCCESS">
      We found something!
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 d-inline-block"/>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.COMBINE_FAILURE">
      Failed to combine recipe <span class="font-weight-bold">{{ event(e).recipe_name }}</span>
      ({{ event(e).recipe_id }})
      <span v-if="TRADESKILLS[event(e).tradeskill_id]">
        tradeskill
        <span class="font-weight-bold">{{ TRADESKILLS[event(e).tradeskill_id] }}</span>
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.COMBINE_SUCCESS">
      Successfully combined recipe <span class="font-weight-bold">{{ event(e).recipe_name }}</span>
      ({{ event(e).recipe_id }})
      <span v-if="TRADESKILLS[event(e).tradeskill_id]">
        tradeskill
        <span class="font-weight-bold">{{ TRADESKILLS[event(e).tradeskill_id] }}</span>
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.DROPPED_ITEM">
      Dropped item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
      ({{ event(e).charges }})
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.DISCOVER_ITEM">
      Discover item
      <item-popover
        :item="itemData[event(e).item_id]"
        class="mr-1 font-weight-bold d-inline-block"
      />
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.REZ_ACCEPTED">
      Accepted resurrection
    </div>

    <div
      style="overflow-wrap: break-word !important; word-break: break-all !important; inline-size: 35vw; "
      v-if="e.event_type_id === PLAYER_EVENT.MERCHANT_SELL"
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
      v-if="e.event_type_id === PLAYER_EVENT.MERCHANT_PURCHASE"
      style="overflow-wrap: break-word !important; word-break: break-all !important; inline-size: 35vw; "
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

    <div v-if="e.event_type_id === PLAYER_EVENT.ZONING">
      Zoned from
      <span class="font-weight-bold" v-if="zoneData[event(e).from_zone_id]">
      {{ zoneData[event(e).from_zone_id].long_name }}
      </span>
      to
      <span class="font-weight-bold" v-if="zoneData[event(e).to_zone_id]">
      {{ zoneData[event(e).to_zone_id].long_name }}
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.LEVEL_GAIN">
      Leveled up (+{{ event(e).levels_gained }}) to
      <span class="font-weight-bold">
      {{ event(e).to_level }}
      </span>
      from level
      <span class="font-weight-bold">
      {{ event(e).from_level }}
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.LEVEL_LOSS">
      Leveled down (-{{ event(e).levels_lost }}) to
      <span class="font-weight-bold">
      {{ event(e).to_level }}
      </span>
      from level
      <span class="font-weight-bold">
      {{ event(e).from_level }}
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.AA_GAIN">
      Gained <span class="font-weight-bold">{{ event(e).aa_gained }}</span> AA point(s)
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.AA_PURCHASE">
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

    <div v-if="e.event_type_id === PLAYER_EVENT.LOOT_ITEM">
      Looted
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }})
      from
      <npc-popover :npc="npcData[event(e).npc_id]" :show-image="false" class="d-inline-block font-weight-bold"/>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.GROUNDSPAWN_PICKUP">
      Picked up
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }})
      from the ground!
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.GROUNDSPAWN_PICKUP">
      Picked up
      <item-popover :item="itemData[event(e).item_id]" class="ml-1 font-weight-bold d-inline-block"/>
      ({{ event(e).charges }})
      from the ground!
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.TASK_ACCEPT">
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

    <div v-if="e.event_type_id === PLAYER_EVENT.DEATH">
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


  </div>
</template>

<script>
import {PLAYER_EVENT} from "@/views/admin/player-event-logs/player-events";
import moment         from "moment";
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

    this.$forceUpdate()
  },
  methods: {
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
    fromNow(time) {
      return moment(time).fromNow()
    },
  }
}
</script>

<style scoped>

</style>
