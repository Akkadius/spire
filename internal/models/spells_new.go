package models

import (
	"github.com/volatiletech/null/v8"
)

type SpellsNew struct {
	ID                   int                `json:"id" gorm:"Column:id"`
	Name                 null.String        `json:"name" gorm:"Column:name"`
	Player1              null.String        `json:"player_1" gorm:"Column:player_1"`
	TeleportZone         null.String        `json:"teleport_zone" gorm:"Column:teleport_zone"`
	YouCast              null.String        `json:"you_cast" gorm:"Column:you_cast"`
	OtherCasts           null.String        `json:"other_casts" gorm:"Column:other_casts"`
	CastOnYou            null.String        `json:"cast_on_you" gorm:"Column:cast_on_you"`
	CastOnOther          null.String        `json:"cast_on_other" gorm:"Column:cast_on_other"`
	SpellFades           null.String        `json:"spell_fades" gorm:"Column:spell_fades"`
	Range                int                `json:"range" gorm:"Column:range"`
	Aoerange             int                `json:"aoerange" gorm:"Column:aoerange"`
	Pushback             int                `json:"pushback" gorm:"Column:pushback"`
	Pushup               int                `json:"pushup" gorm:"Column:pushup"`
	CastTime             int                `json:"cast_time" gorm:"Column:cast_time"`
	RecoveryTime         int                `json:"recovery_time" gorm:"Column:recovery_time"`
	RecastTime           int                `json:"recast_time" gorm:"Column:recast_time"`
	Buffdurationformula  int                `json:"buffdurationformula" gorm:"Column:buffdurationformula"`
	Buffduration         int                `json:"buffduration" gorm:"Column:buffduration"`
	AEDuration           int                `json:"ae_duration" gorm:"Column:AEDuration"`
	Mana                 int                `json:"mana" gorm:"Column:mana"`
	EffectBaseValue1     int                `json:"effect_base_value_1" gorm:"Column:effect_base_value1"`
	EffectBaseValue2     int                `json:"effect_base_value_2" gorm:"Column:effect_base_value2"`
	EffectBaseValue3     int                `json:"effect_base_value_3" gorm:"Column:effect_base_value3"`
	EffectBaseValue4     int                `json:"effect_base_value_4" gorm:"Column:effect_base_value4"`
	EffectBaseValue5     int                `json:"effect_base_value_5" gorm:"Column:effect_base_value5"`
	EffectBaseValue6     int                `json:"effect_base_value_6" gorm:"Column:effect_base_value6"`
	EffectBaseValue7     int                `json:"effect_base_value_7" gorm:"Column:effect_base_value7"`
	EffectBaseValue8     int                `json:"effect_base_value_8" gorm:"Column:effect_base_value8"`
	EffectBaseValue9     int                `json:"effect_base_value_9" gorm:"Column:effect_base_value9"`
	EffectBaseValue10    int                `json:"effect_base_value_10" gorm:"Column:effect_base_value10"`
	EffectBaseValue11    int                `json:"effect_base_value_11" gorm:"Column:effect_base_value11"`
	EffectBaseValue12    int                `json:"effect_base_value_12" gorm:"Column:effect_base_value12"`
	EffectLimitValue1    int                `json:"effect_limit_value_1" gorm:"Column:effect_limit_value1"`
	EffectLimitValue2    int                `json:"effect_limit_value_2" gorm:"Column:effect_limit_value2"`
	EffectLimitValue3    int                `json:"effect_limit_value_3" gorm:"Column:effect_limit_value3"`
	EffectLimitValue4    int                `json:"effect_limit_value_4" gorm:"Column:effect_limit_value4"`
	EffectLimitValue5    int                `json:"effect_limit_value_5" gorm:"Column:effect_limit_value5"`
	EffectLimitValue6    int                `json:"effect_limit_value_6" gorm:"Column:effect_limit_value6"`
	EffectLimitValue7    int                `json:"effect_limit_value_7" gorm:"Column:effect_limit_value7"`
	EffectLimitValue8    int                `json:"effect_limit_value_8" gorm:"Column:effect_limit_value8"`
	EffectLimitValue9    int                `json:"effect_limit_value_9" gorm:"Column:effect_limit_value9"`
	EffectLimitValue10   int                `json:"effect_limit_value_10" gorm:"Column:effect_limit_value10"`
	EffectLimitValue11   int                `json:"effect_limit_value_11" gorm:"Column:effect_limit_value11"`
	EffectLimitValue12   int                `json:"effect_limit_value_12" gorm:"Column:effect_limit_value12"`
	Max1                 int                `json:"max_1" gorm:"Column:max1"`
	Max2                 int                `json:"max_2" gorm:"Column:max2"`
	Max3                 int                `json:"max_3" gorm:"Column:max3"`
	Max4                 int                `json:"max_4" gorm:"Column:max4"`
	Max5                 int                `json:"max_5" gorm:"Column:max5"`
	Max6                 int                `json:"max_6" gorm:"Column:max6"`
	Max7                 int                `json:"max_7" gorm:"Column:max7"`
	Max8                 int                `json:"max_8" gorm:"Column:max8"`
	Max9                 int                `json:"max_9" gorm:"Column:max9"`
	Max10                int                `json:"max_10" gorm:"Column:max10"`
	Max11                int                `json:"max_11" gorm:"Column:max11"`
	Max12                int                `json:"max_12" gorm:"Column:max12"`
	Icon                 int                `json:"icon" gorm:"Column:icon"`
	Memicon              int                `json:"memicon" gorm:"Column:memicon"`
	Components1          int                `json:"components_1" gorm:"Column:components1"`
	Components2          int                `json:"components_2" gorm:"Column:components2"`
	Components3          int                `json:"components_3" gorm:"Column:components3"`
	Components4          int                `json:"components_4" gorm:"Column:components4"`
	ComponentCounts1     int                `json:"component_counts_1" gorm:"Column:component_counts1"`
	ComponentCounts2     int                `json:"component_counts_2" gorm:"Column:component_counts2"`
	ComponentCounts3     int                `json:"component_counts_3" gorm:"Column:component_counts3"`
	ComponentCounts4     int                `json:"component_counts_4" gorm:"Column:component_counts4"`
	NoexpendReagent1     int                `json:"noexpend_reagent_1" gorm:"Column:NoexpendReagent1"`
	NoexpendReagent2     int                `json:"noexpend_reagent_2" gorm:"Column:NoexpendReagent2"`
	NoexpendReagent3     int                `json:"noexpend_reagent_3" gorm:"Column:NoexpendReagent3"`
	NoexpendReagent4     int                `json:"noexpend_reagent_4" gorm:"Column:NoexpendReagent4"`
	Formula1             int                `json:"formula_1" gorm:"Column:formula1"`
	Formula2             int                `json:"formula_2" gorm:"Column:formula2"`
	Formula3             int                `json:"formula_3" gorm:"Column:formula3"`
	Formula4             int                `json:"formula_4" gorm:"Column:formula4"`
	Formula5             int                `json:"formula_5" gorm:"Column:formula5"`
	Formula6             int                `json:"formula_6" gorm:"Column:formula6"`
	Formula7             int                `json:"formula_7" gorm:"Column:formula7"`
	Formula8             int                `json:"formula_8" gorm:"Column:formula8"`
	Formula9             int                `json:"formula_9" gorm:"Column:formula9"`
	Formula10            int                `json:"formula_10" gorm:"Column:formula10"`
	Formula11            int                `json:"formula_11" gorm:"Column:formula11"`
	Formula12            int                `json:"formula_12" gorm:"Column:formula12"`
	LightType            int                `json:"light_type" gorm:"Column:LightType"`
	GoodEffect           int                `json:"good_effect" gorm:"Column:goodEffect"`
	Activated            int                `json:"activated" gorm:"Column:Activated"`
	Resisttype           int                `json:"resisttype" gorm:"Column:resisttype"`
	Effectid1            int                `json:"effectid_1" gorm:"Column:effectid1"`
	Effectid2            int                `json:"effectid_2" gorm:"Column:effectid2"`
	Effectid3            int                `json:"effectid_3" gorm:"Column:effectid3"`
	Effectid4            int                `json:"effectid_4" gorm:"Column:effectid4"`
	Effectid5            int                `json:"effectid_5" gorm:"Column:effectid5"`
	Effectid6            int                `json:"effectid_6" gorm:"Column:effectid6"`
	Effectid7            int                `json:"effectid_7" gorm:"Column:effectid7"`
	Effectid8            int                `json:"effectid_8" gorm:"Column:effectid8"`
	Effectid9            int                `json:"effectid_9" gorm:"Column:effectid9"`
	Effectid10           int                `json:"effectid_10" gorm:"Column:effectid10"`
	Effectid11           int                `json:"effectid_11" gorm:"Column:effectid11"`
	Effectid12           int                `json:"effectid_12" gorm:"Column:effectid12"`
	Targettype           int                `json:"targettype" gorm:"Column:targettype"`
	Basediff             int                `json:"basediff" gorm:"Column:basediff"`
	Skill                int                `json:"skill" gorm:"Column:skill"`
	Zonetype             int                `json:"zonetype" gorm:"Column:zonetype"`
	EnvironmentType      int                `json:"environment_type" gorm:"Column:EnvironmentType"`
	TimeOfDay            int                `json:"time_of_day" gorm:"Column:TimeOfDay"`
	Classes1             int                `json:"classes_1" gorm:"Column:classes1"`
	Classes2             int                `json:"classes_2" gorm:"Column:classes2"`
	Classes3             int                `json:"classes_3" gorm:"Column:classes3"`
	Classes4             int                `json:"classes_4" gorm:"Column:classes4"`
	Classes5             int                `json:"classes_5" gorm:"Column:classes5"`
	Classes6             int                `json:"classes_6" gorm:"Column:classes6"`
	Classes7             int                `json:"classes_7" gorm:"Column:classes7"`
	Classes8             int                `json:"classes_8" gorm:"Column:classes8"`
	Classes9             int                `json:"classes_9" gorm:"Column:classes9"`
	Classes10            int                `json:"classes_10" gorm:"Column:classes10"`
	Classes11            int                `json:"classes_11" gorm:"Column:classes11"`
	Classes12            int                `json:"classes_12" gorm:"Column:classes12"`
	Classes13            int                `json:"classes_13" gorm:"Column:classes13"`
	Classes14            int                `json:"classes_14" gorm:"Column:classes14"`
	Classes15            int                `json:"classes_15" gorm:"Column:classes15"`
	Classes16            int                `json:"classes_16" gorm:"Column:classes16"`
	CastingAnim          int                `json:"casting_anim" gorm:"Column:CastingAnim"`
	TargetAnim           int                `json:"target_anim" gorm:"Column:TargetAnim"`
	TravelType           int                `json:"travel_type" gorm:"Column:TravelType"`
	SpellAffectIndex     int                `json:"spell_affect_index" gorm:"Column:SpellAffectIndex"`
	DisallowSit          int                `json:"disallow_sit" gorm:"Column:disallow_sit"`
	Deities0             int                `json:"deities_0" gorm:"Column:deities0"`
	Deities1             int                `json:"deities_1" gorm:"Column:deities1"`
	Deities2             int                `json:"deities_2" gorm:"Column:deities2"`
	Deities3             int                `json:"deities_3" gorm:"Column:deities3"`
	Deities4             int                `json:"deities_4" gorm:"Column:deities4"`
	Deities5             int                `json:"deities_5" gorm:"Column:deities5"`
	Deities6             int                `json:"deities_6" gorm:"Column:deities6"`
	Deities7             int                `json:"deities_7" gorm:"Column:deities7"`
	Deities8             int                `json:"deities_8" gorm:"Column:deities8"`
	Deities9             int                `json:"deities_9" gorm:"Column:deities9"`
	Deities10            int                `json:"deities_10" gorm:"Column:deities10"`
	Deities11            int                `json:"deities_11" gorm:"Column:deities11"`
	Deities12            int                `json:"deities_12" gorm:"Column:deities12"`
	Deities13            int                `json:"deities_13" gorm:"Column:deities13"`
	Deities14            int                `json:"deities_14" gorm:"Column:deities14"`
	Deities15            int                `json:"deities_15" gorm:"Column:deities15"`
	Deities16            int                `json:"deities_16" gorm:"Column:deities16"`
	Field142             int                `json:"field_142" gorm:"Column:field142"`
	Field143             int                `json:"field_143" gorm:"Column:field143"`
	NewIcon              int                `json:"new_icon" gorm:"Column:new_icon"`
	Spellanim            int                `json:"spellanim" gorm:"Column:spellanim"`
	Uninterruptable      int                `json:"uninterruptable" gorm:"Column:uninterruptable"`
	ResistDiff           int                `json:"resist_diff" gorm:"Column:ResistDiff"`
	DotStackingExempt    int                `json:"dot_stacking_exempt" gorm:"Column:dot_stacking_exempt"`
	Deleteable           int                `json:"deleteable" gorm:"Column:deleteable"`
	RecourseLink         int                `json:"recourse_link" gorm:"Column:RecourseLink"`
	NoPartialResist      int                `json:"no_partial_resist" gorm:"Column:no_partial_resist"`
	Field152             int                `json:"field_152" gorm:"Column:field152"`
	Field153             int                `json:"field_153" gorm:"Column:field153"`
	ShortBuffBox         int                `json:"short_buff_box" gorm:"Column:short_buff_box"`
	Descnum              int                `json:"descnum" gorm:"Column:descnum"`
	Typedescnum          null.Int           `json:"typedescnum" gorm:"Column:typedescnum"`
	Effectdescnum        null.Int           `json:"effectdescnum" gorm:"Column:effectdescnum"`
	Effectdescnum2       int                `json:"effectdescnum_2" gorm:"Column:effectdescnum2"`
	NpcNoLos             int                `json:"npc_no_los" gorm:"Column:npc_no_los"`
	Field160             int                `json:"field_160" gorm:"Column:field160"`
	Reflectable          int                `json:"reflectable" gorm:"Column:reflectable"`
	Bonushate            int                `json:"bonushate" gorm:"Column:bonushate"`
	Field163             int                `json:"field_163" gorm:"Column:field163"`
	Field164             int                `json:"field_164" gorm:"Column:field164"`
	LdonTrap             int                `json:"ldon_trap" gorm:"Column:ldon_trap"`
	EndurCost            int                `json:"endur_cost" gorm:"Column:EndurCost"`
	EndurTimerIndex      int                `json:"endur_timer_index" gorm:"Column:EndurTimerIndex"`
	IsDiscipline         int                `json:"is_discipline" gorm:"Column:IsDiscipline"`
	Field169             int                `json:"field_169" gorm:"Column:field169"`
	Field170             int                `json:"field_170" gorm:"Column:field170"`
	Field171             int                `json:"field_171" gorm:"Column:field171"`
	Field172             int                `json:"field_172" gorm:"Column:field172"`
	HateAdded            int                `json:"hate_added" gorm:"Column:HateAdded"`
	EndurUpkeep          int                `json:"endur_upkeep" gorm:"Column:EndurUpkeep"`
	Numhitstype          int                `json:"numhitstype" gorm:"Column:numhitstype"`
	Numhits              int                `json:"numhits" gorm:"Column:numhits"`
	Pvpresistbase        int                `json:"pvpresistbase" gorm:"Column:pvpresistbase"`
	Pvpresistcalc        int                `json:"pvpresistcalc" gorm:"Column:pvpresistcalc"`
	Pvpresistcap         int                `json:"pvpresistcap" gorm:"Column:pvpresistcap"`
	SpellCategory        int                `json:"spell_category" gorm:"Column:spell_category"`
	PvpDuration          int                `json:"pvp_duration" gorm:"Column:pvp_duration"`
	PvpDurationCap       int                `json:"pvp_duration_cap" gorm:"Column:pvp_duration_cap"`
	PcnpcOnlyFlag        null.Int           `json:"pcnpc_only_flag" gorm:"Column:pcnpc_only_flag"`
	CastNotStanding      null.Int           `json:"cast_not_standing" gorm:"Column:cast_not_standing"`
	CanMgb               int                `json:"can_mgb" gorm:"Column:can_mgb"`
	Nodispell            int                `json:"nodispell" gorm:"Column:nodispell"`
	NpcCategory          int                `json:"npc_category" gorm:"Column:npc_category"`
	NpcUsefulness        int                `json:"npc_usefulness" gorm:"Column:npc_usefulness"`
	MinResist            int                `json:"min_resist" gorm:"Column:MinResist"`
	MaxResist            int                `json:"max_resist" gorm:"Column:MaxResist"`
	ViralTargets         int                `json:"viral_targets" gorm:"Column:viral_targets"`
	ViralTimer           int                `json:"viral_timer" gorm:"Column:viral_timer"`
	Nimbuseffect         null.Int           `json:"nimbuseffect" gorm:"Column:nimbuseffect"`
	ConeStartAngle       int                `json:"cone_start_angle" gorm:"Column:ConeStartAngle"`
	ConeStopAngle        int                `json:"cone_stop_angle" gorm:"Column:ConeStopAngle"`
	Sneaking             int                `json:"sneaking" gorm:"Column:sneaking"`
	NotExtendable        int                `json:"not_extendable" gorm:"Column:not_extendable"`
	Field198             int                `json:"field_198" gorm:"Column:field198"`
	Field199             int                `json:"field_199" gorm:"Column:field199"`
	Suspendable          null.Int           `json:"suspendable" gorm:"Column:suspendable"`
	ViralRange           int                `json:"viral_range" gorm:"Column:viral_range"`
	Songcap              null.Int           `json:"songcap" gorm:"Column:songcap"`
	Field203             null.Int           `json:"field_203" gorm:"Column:field203"`
	Field204             null.Int           `json:"field_204" gorm:"Column:field204"`
	NoBlock              int                `json:"no_block" gorm:"Column:no_block"`
	Field206             null.Int           `json:"field_206" gorm:"Column:field206"`
	Spellgroup           null.Int           `json:"spellgroup" gorm:"Column:spellgroup"`
	Rank                 int                `json:"rank" gorm:"Column:rank"`
	Field209             null.Int           `json:"field_209" gorm:"Column:field209"`
	Field210             null.Int           `json:"field_210" gorm:"Column:field210"`
	CastRestriction      int                `json:"cast_restriction" gorm:"Column:CastRestriction"`
	Allowrest            null.Int           `json:"allowrest" gorm:"Column:allowrest"`
	InCombat             int                `json:"in_combat" gorm:"Column:InCombat"`
	OutofCombat          int                `json:"outof_combat" gorm:"Column:OutofCombat"`
	Field215             null.Int           `json:"field_215" gorm:"Column:field215"`
	Field216             null.Int           `json:"field_216" gorm:"Column:field216"`
	Field217             null.Int           `json:"field_217" gorm:"Column:field217"`
	Aemaxtargets         int                `json:"aemaxtargets" gorm:"Column:aemaxtargets"`
	Maxtargets           null.Int           `json:"maxtargets" gorm:"Column:maxtargets"`
	Field220             null.Int           `json:"field_220" gorm:"Column:field220"`
	Field221             null.Int           `json:"field_221" gorm:"Column:field221"`
	Field222             null.Int           `json:"field_222" gorm:"Column:field222"`
	Field223             null.Int           `json:"field_223" gorm:"Column:field223"`
	Persistdeath         null.Int           `json:"persistdeath" gorm:"Column:persistdeath"`
	Field225             int                `json:"field_225" gorm:"Column:field225"`
	Field226             int                `json:"field_226" gorm:"Column:field226"`
	MinDist              float32            `json:"min_dist" gorm:"Column:min_dist"`
	MinDistMod           float32            `json:"min_dist_mod" gorm:"Column:min_dist_mod"`
	MaxDist              float32            `json:"max_dist" gorm:"Column:max_dist"`
	MaxDistMod           float32            `json:"max_dist_mod" gorm:"Column:max_dist_mod"`
	MinRange             int                `json:"min_range" gorm:"Column:min_range"`
	Field232             int                `json:"field_232" gorm:"Column:field232"`
	Field233             int                `json:"field_233" gorm:"Column:field233"`
	Field234             int                `json:"field_234" gorm:"Column:field234"`
	Field235             int                `json:"field_235" gorm:"Column:field235"`
	Field236             int                `json:"field_236" gorm:"Column:field236"`
	Aura                 *Aura              `json:"aura,omitempty" gorm:"foreignKey:id;references:spell_id"`
	Damageshieldtypes    []Damageshieldtype `json:"damageshieldtypes,omitempty" gorm:"foreignKey:spellid;references:id"`
	SpellBuckets         []SpellBucket      `json:"spell_buckets,omitempty" gorm:"foreignKey:spellid;references:id"`
	SpellGlobals         []SpellGlobal      `json:"spell_globals,omitempty" gorm:"foreignKey:spellid;references:id"`
	BlockedSpells        []BlockedSpell     `json:"blocked_spells,omitempty" gorm:"foreignKey:spellid;references:id"`
	Items                []Item             `json:"items,omitempty" gorm:"foreignKey:clickeffect;references:id"`
	NpcSpellsEntries     []NpcSpellsEntry   `json:"npc_spells_entries,omitempty" gorm:"foreignKey:spellid;references:id"`
	BotSpellsEntries     []BotSpellsEntry   `json:"bot_spells_entries,omitempty" gorm:"foreignKey:spell_id;references:id"`
}

func (SpellsNew) TableName() string {
    return "spells_new"
}

func (SpellsNew) Relationships() []string {
    return []string{
		"Aura",
		"Aura.SpellsNew",
		"BlockedSpells",
		"BotSpellsEntries",
		"BotSpellsEntries.NpcSpell",
		"BotSpellsEntries.NpcSpell.BotSpellsEntries",
		"BotSpellsEntries.NpcSpell.NpcSpell",
		"BotSpellsEntries.NpcSpell.NpcSpellsEntries",
		"BotSpellsEntries.NpcSpell.NpcSpellsEntries.SpellsNew",
		"BotSpellsEntries.SpellsNew",
		"Damageshieldtypes",
		"Items",
		"Items.AlternateCurrencies",
		"Items.AlternateCurrencies.Item",
		"Items.CharacterCorpseItems",
		"Items.DiscoveredItems",
		"Items.Doors",
		"Items.Doors.Item",
		"Items.Fishings",
		"Items.Fishings.Item",
		"Items.Fishings.NpcType",
		"Items.Fishings.NpcType.AlternateCurrency",
		"Items.Fishings.NpcType.AlternateCurrency.Item",
		"Items.Fishings.NpcType.Loottable",
		"Items.Fishings.NpcType.Loottable.LoottableEntries",
		"Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Items.Fishings.NpcType.Loottable.NpcTypes",
		"Items.Fishings.NpcType.Merchantlists",
		"Items.Fishings.NpcType.Merchantlists.Items",
		"Items.Fishings.NpcType.Merchantlists.NpcTypes",
		"Items.Fishings.NpcType.NpcEmotes",
		"Items.Fishings.NpcType.NpcFactions",
		"Items.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"Items.Fishings.NpcType.NpcSpell",
		"Items.Fishings.NpcType.NpcSpell.BotSpellsEntries",
		"Items.Fishings.NpcType.NpcSpell.BotSpellsEntries.NpcSpell",
		"Items.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew",
		"Items.Fishings.NpcType.NpcSpell.NpcSpell",
		"Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Items.Fishings.NpcType.NpcTypesTint",
		"Items.Fishings.NpcType.Spawnentries",
		"Items.Fishings.NpcType.Spawnentries.NpcType",
		"Items.Fishings.NpcType.Spawnentries.Spawngroup",
		"Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Items.Fishings.Zone",
		"Items.Forages",
		"Items.Forages.Item",
		"Items.Forages.Zone",
		"Items.GroundSpawns",
		"Items.GroundSpawns.Zone",
		"Items.ItemTicks",
		"Items.Keyrings",
		"Items.LootdropEntries",
		"Items.LootdropEntries.Item",
		"Items.LootdropEntries.Lootdrop",
		"Items.LootdropEntries.Lootdrop.LootdropEntries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Items.Merchantlists",
		"Items.Merchantlists.Items",
		"Items.Merchantlists.NpcTypes",
		"Items.Merchantlists.NpcTypes.AlternateCurrency",
		"Items.Merchantlists.NpcTypes.AlternateCurrency.Item",
		"Items.Merchantlists.NpcTypes.Loottable",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"Items.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"Items.Merchantlists.NpcTypes.Merchantlists",
		"Items.Merchantlists.NpcTypes.NpcEmotes",
		"Items.Merchantlists.NpcTypes.NpcFactions",
		"Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"Items.Merchantlists.NpcTypes.NpcSpell",
		"Items.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries",
		"Items.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"Items.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Items.Merchantlists.NpcTypes.NpcTypesTint",
		"Items.Merchantlists.NpcTypes.Spawnentries",
		"Items.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Items.ObjectContents",
		"Items.Objects",
		"Items.Objects.Item",
		"Items.Objects.Zone",
		"Items.TradeskillRecipeEntries",
		"Items.TradeskillRecipeEntries.TradeskillRecipe",
		"Items.TradeskillRecipeEntries.TradeskillRecipe.TradeskillRecipeEntries",
		"Items.TributeLevels",
		"NpcSpellsEntries",
		"NpcSpellsEntries.SpellsNew",
		"SpellBuckets",
		"SpellGlobals",
	}
}

func (SpellsNew) Connection() string {
    return "eqemu_content"
}
