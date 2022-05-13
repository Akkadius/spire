package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcType struct {
	ID                     int                `json:"id" gorm:"Column:id"`
	Name                   string             `json:"name" gorm:"Column:name"`
	Lastname               null.String        `json:"lastname" gorm:"Column:lastname"`
	Level                  uint8              `json:"level" gorm:"Column:level"`
	Race                   uint16             `json:"race" gorm:"Column:race"`
	Class                  uint8              `json:"class" gorm:"Column:class"`
	Bodytype               int                `json:"bodytype" gorm:"Column:bodytype"`
	Hp                     null.Int64         `json:"hp" gorm:"Column:hp"`
	Mana                   null.Int64         `json:"mana" gorm:"Column:mana"`
	Gender                 uint8              `json:"gender" gorm:"Column:gender"`
	Texture                uint8              `json:"texture" gorm:"Column:texture"`
	Helmtexture            uint8              `json:"helmtexture" gorm:"Column:helmtexture"`
	Herosforgemodel        int                `json:"herosforgemodel" gorm:"Column:herosforgemodel"`
	Size                   float32            `json:"size" gorm:"Column:size"`
	HpRegenRate            null.Int64         `json:"hp_regen_rate" gorm:"Column:hp_regen_rate"`
	HpRegenPerSecond       null.Int64         `json:"hp_regen_per_second" gorm:"Column:hp_regen_per_second"`
	ManaRegenRate          null.Int64         `json:"mana_regen_rate" gorm:"Column:mana_regen_rate"`
	LoottableId            uint               `json:"loottable_id" gorm:"Column:loottable_id"`
	MerchantId             uint               `json:"merchant_id" gorm:"Column:merchant_id"`
	AltCurrencyId          uint               `json:"alt_currency_id" gorm:"Column:alt_currency_id"`
	NpcSpellsId            uint               `json:"npc_spells_id" gorm:"Column:npc_spells_id"`
	NpcSpellsEffectsId     uint               `json:"npc_spells_effects_id" gorm:"Column:npc_spells_effects_id"`
	NpcFactionId           int                `json:"npc_faction_id" gorm:"Column:npc_faction_id"`
	AdventureTemplateId    uint               `json:"adventure_template_id" gorm:"Column:adventure_template_id"`
	TrapTemplate           null.Uint          `json:"trap_template" gorm:"Column:trap_template"`
	Mindmg                 uint               `json:"mindmg" gorm:"Column:mindmg"`
	Maxdmg                 uint               `json:"maxdmg" gorm:"Column:maxdmg"`
	AttackCount            int16              `json:"attack_count" gorm:"Column:attack_count"`
	Npcspecialattks        string             `json:"npcspecialattks" gorm:"Column:npcspecialattks"`
	SpecialAbilities       null.String        `json:"special_abilities" gorm:"Column:special_abilities"`
	Aggroradius            uint               `json:"aggroradius" gorm:"Column:aggroradius"`
	Assistradius           uint               `json:"assistradius" gorm:"Column:assistradius"`
	Face                   uint               `json:"face" gorm:"Column:face"`
	LuclinHairstyle        uint               `json:"luclin_hairstyle" gorm:"Column:luclin_hairstyle"`
	LuclinHaircolor        uint               `json:"luclin_haircolor" gorm:"Column:luclin_haircolor"`
	LuclinEyecolor         uint               `json:"luclin_eyecolor" gorm:"Column:luclin_eyecolor"`
	LuclinEyecolor2        uint               `json:"luclin_eyecolor_2" gorm:"Column:luclin_eyecolor2"`
	LuclinBeardcolor       uint               `json:"luclin_beardcolor" gorm:"Column:luclin_beardcolor"`
	LuclinBeard            uint               `json:"luclin_beard" gorm:"Column:luclin_beard"`
	DrakkinHeritage        int                `json:"drakkin_heritage" gorm:"Column:drakkin_heritage"`
	DrakkinTattoo          int                `json:"drakkin_tattoo" gorm:"Column:drakkin_tattoo"`
	DrakkinDetails         int                `json:"drakkin_details" gorm:"Column:drakkin_details"`
	ArmortintId            uint               `json:"armortint_id" gorm:"Column:armortint_id"`
	ArmortintRed           uint8              `json:"armortint_red" gorm:"Column:armortint_red"`
	ArmortintGreen         uint8              `json:"armortint_green" gorm:"Column:armortint_green"`
	ArmortintBlue          uint8              `json:"armortint_blue" gorm:"Column:armortint_blue"`
	DMeleeTexture1         int                `json:"d_melee_texture_1" gorm:"Column:d_melee_texture1"`
	DMeleeTexture2         int                `json:"d_melee_texture_2" gorm:"Column:d_melee_texture2"`
	AmmoIdfile             string             `json:"ammo_idfile" gorm:"Column:ammo_idfile"`
	PrimMeleeType          uint8              `json:"prim_melee_type" gorm:"Column:prim_melee_type"`
	SecMeleeType           uint8              `json:"sec_melee_type" gorm:"Column:sec_melee_type"`
	RangedType             uint8              `json:"ranged_type" gorm:"Column:ranged_type"`
	Runspeed               float32            `json:"runspeed" gorm:"Column:runspeed"`
	MR                     int16              `json:"mr" gorm:"Column:MR"`
	CR                     int16              `json:"cr" gorm:"Column:CR"`
	DR                     int16              `json:"dr" gorm:"Column:DR"`
	FR                     int16              `json:"fr" gorm:"Column:FR"`
	PR                     int16              `json:"pr" gorm:"Column:PR"`
	Corrup                 int16              `json:"corrup" gorm:"Column:Corrup"`
	PhR                    uint16             `json:"ph_r" gorm:"Column:PhR"`
	SeeInvis               int16              `json:"see_invis" gorm:"Column:see_invis"`
	SeeInvisUndead         int16              `json:"see_invis_undead" gorm:"Column:see_invis_undead"`
	Qglobal                uint               `json:"qglobal" gorm:"Column:qglobal"`
	AC                     int16              `json:"ac" gorm:"Column:AC"`
	NpcAggro               int8               `json:"npc_aggro" gorm:"Column:npc_aggro"`
	SpawnLimit             int8               `json:"spawn_limit" gorm:"Column:spawn_limit"`
	AttackSpeed            float32            `json:"attack_speed" gorm:"Column:attack_speed"`
	AttackDelay            uint8              `json:"attack_delay" gorm:"Column:attack_delay"`
	Findable               int8               `json:"findable" gorm:"Column:findable"`
	STR                    uint32             `json:"str" gorm:"Column:STR"`
	STA                    uint32             `json:"sta" gorm:"Column:STA"`
	DEX                    uint32             `json:"dex" gorm:"Column:DEX"`
	AGI                    uint32             `json:"agi" gorm:"Column:AGI"`
	INT                    uint32             `json:"_int" gorm:"Column:_INT"`
	WIS                    uint32             `json:"wis" gorm:"Column:WIS"`
	CHA                    uint32             `json:"cha" gorm:"Column:CHA"`
	SeeHide                int8               `json:"see_hide" gorm:"Column:see_hide"`
	SeeImprovedHide        int8               `json:"see_improved_hide" gorm:"Column:see_improved_hide"`
	Trackable              int8               `json:"trackable" gorm:"Column:trackable"`
	Isbot                  int8               `json:"isbot" gorm:"Column:isbot"`
	Exclude                int8               `json:"exclude" gorm:"Column:exclude"`
	ATK                    int32              `json:"atk" gorm:"Column:ATK"`
	Accuracy               int32              `json:"accuracy" gorm:"Column:Accuracy"`
	Avoidance              uint32             `json:"avoidance" gorm:"Column:Avoidance"`
	SlowMitigation         int16              `json:"slow_mitigation" gorm:"Column:slow_mitigation"`
	Version                uint16             `json:"version" gorm:"Column:version"`
	Maxlevel               int8               `json:"maxlevel" gorm:"Column:maxlevel"`
	Scalerate              int                `json:"scalerate" gorm:"Column:scalerate"`
	PrivateCorpse          uint8              `json:"private_corpse" gorm:"Column:private_corpse"`
	UniqueSpawnByName      uint8              `json:"unique_spawn_by_name" gorm:"Column:unique_spawn_by_name"`
	Underwater             uint8              `json:"underwater" gorm:"Column:underwater"`
	Isquest                int8               `json:"isquest" gorm:"Column:isquest"`
	Emoteid                uint               `json:"emoteid" gorm:"Column:emoteid"`
	Spellscale             float32            `json:"spellscale" gorm:"Column:spellscale"`
	Healscale              float32            `json:"healscale" gorm:"Column:healscale"`
	NoTargetHotkey         uint8              `json:"no_target_hotkey" gorm:"Column:no_target_hotkey"`
	RaidTarget             uint8              `json:"raid_target" gorm:"Column:raid_target"`
	Armtexture             int8               `json:"armtexture" gorm:"Column:armtexture"`
	Bracertexture          int8               `json:"bracertexture" gorm:"Column:bracertexture"`
	Handtexture            int8               `json:"handtexture" gorm:"Column:handtexture"`
	Legtexture             int8               `json:"legtexture" gorm:"Column:legtexture"`
	Feettexture            int8               `json:"feettexture" gorm:"Column:feettexture"`
	Light                  int8               `json:"light" gorm:"Column:light"`
	Walkspeed              int8               `json:"walkspeed" gorm:"Column:walkspeed"`
	Peqid                  int                `json:"peqid" gorm:"Column:peqid"`
	Unique2                int8               `json:"unique_" gorm:"Column:unique_"`
	Fixed                  int8               `json:"fixed" gorm:"Column:fixed"`
	IgnoreDespawn          int8               `json:"ignore_despawn" gorm:"Column:ignore_despawn"`
	ShowName               int8               `json:"show_name" gorm:"Column:show_name"`
	Untargetable           int8               `json:"untargetable" gorm:"Column:untargetable"`
	CharmAc                null.Int16         `json:"charm_ac" gorm:"Column:charm_ac"`
	CharmMinDmg            null.Int           `json:"charm_min_dmg" gorm:"Column:charm_min_dmg"`
	CharmMaxDmg            null.Int           `json:"charm_max_dmg" gorm:"Column:charm_max_dmg"`
	CharmAttackDelay       null.Int8          `json:"charm_attack_delay" gorm:"Column:charm_attack_delay"`
	CharmAccuracyRating    null.Int32         `json:"charm_accuracy_rating" gorm:"Column:charm_accuracy_rating"`
	CharmAvoidanceRating   null.Int32         `json:"charm_avoidance_rating" gorm:"Column:charm_avoidance_rating"`
	CharmAtk               null.Int32         `json:"charm_atk" gorm:"Column:charm_atk"`
	SkipGlobalLoot         null.Int8          `json:"skip_global_loot" gorm:"Column:skip_global_loot"`
	RareSpawn              null.Int8          `json:"rare_spawn" gorm:"Column:rare_spawn"`
	StuckBehavior          int8               `json:"stuck_behavior" gorm:"Column:stuck_behavior"`
	Model                  int16              `json:"model" gorm:"Column:model"`
	Flymode                int8               `json:"flymode" gorm:"Column:flymode"`
	AlwaysAggro            int8               `json:"always_aggro" gorm:"Column:always_aggro"`
	ExpMod                 int                `json:"exp_mod" gorm:"Column:exp_mod"`
	AlternateCurrency      *AlternateCurrency `json:"alternate_currency,omitempty" gorm:"foreignKey:alt_currency_id;references:id"`
	Merchantlists          []Merchantlist     `json:"merchantlists,omitempty" gorm:"foreignKey:merchantid;references:merchant_id"`
	NpcFactions            []NpcFaction       `json:"npc_factions,omitempty" gorm:"foreignKey:id;references:npc_faction_id"`
	NpcSpells              []NpcSpell         `json:"npc_spells,omitempty" gorm:"foreignKey:id;references:npc_spells_id"`
	Spawnentries           []Spawnentry       `json:"spawnentries,omitempty" gorm:"foreignKey:npcID;references:id"`
	NpcEmotes              []NpcEmote         `json:"npc_emotes,omitempty" gorm:"foreignKey:emoteid;references:emoteid"`
	NpcTypesTint           *NpcTypesTint      `json:"npc_types_tint,omitempty" gorm:"foreignKey:armortint_id;references:id"`
	Loottable              *Loottable         `json:"loottable,omitempty" gorm:"foreignKey:loottable_id;references:id"`
}

func (NpcType) TableName() string {
    return "npc_types"
}

func (NpcType) Relationships() []string {
    return []string{
		"AlternateCurrency",
		"Loottable",
		"Loottable.LoottableEntries",
		"Loottable.LoottableEntries.LootdropEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
		"Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.Tasks",
		"Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities",
		"Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists",
		"Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType",
		"Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets",
		"Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Loottable.LoottableEntries.Loottable",
		"Loottable.NpcTypes",
		"Merchantlists",
		"Merchantlists.NpcType",
		"NpcEmotes",
		"NpcFactions",
		"NpcFactions.NpcFactionEntries",
		"NpcSpells",
		"NpcSpells.NpcSpellsEntries",
		"NpcTypesTint",
		"Spawnentries",
		"Spawnentries.NpcType",
		"Spawnentries.Spawngroup",
		"Spawnentries.Spawngroup.Spawn2",
		"Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (NpcType) Connection() string {
    return "eqemu_content"
}
