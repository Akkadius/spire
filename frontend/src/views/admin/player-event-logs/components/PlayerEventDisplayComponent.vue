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

    <div v-if="e.event_type_id === PLAYER_EVENT.MERCHANT_SELL">
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

    <div v-if="e.event_type_id === PLAYER_EVENT.MERCHANT_PURCHASE">
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
      Leveled up (+{{event(e).levels_gained}}) to
      <span class="font-weight-bold">
      {{ event(e).to_level }}
      </span>
      from level
      <span class="font-weight-bold">
      {{ event(e).from_level }}
      </span>
    </div>

    <div v-if="e.event_type_id === PLAYER_EVENT.LEVEL_LOSS">
      Leveled down (-{{event(e).levels_lost}}) to
      <span class="font-weight-bold">
      {{ event(e).to_level }}
      </span>
      from level
      <span class="font-weight-bold">
      {{ event(e).from_level }}
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

export default {
  name: "PlayerEventDisplayComponent",
  components: { EqCashDisplay, NpcPopover, ItemPopover },
  props: {
    e: Object, // event
  },
  data() {
    return {
      PLAYER_EVENT: PLAYER_EVENT,
      DB_SKILLS: DB_SKILLS,

      itemData: {},
      npcData: {},
      zoneData: {},
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

    this.$forceUpdate()
  },
  methods: {
    async getItem(itemId) {
      console.log(await Items.getItem(itemId))
      return await Items.getItem(itemId)
    },
    event(e) {
      console.log(JSON.parse(e.event_data))
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
