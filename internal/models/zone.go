package models

import (
	"github.com/volatiletech/null/v8"
)

type Zone struct {
	ID                        int         `json:"id" gorm:"Column:id"`
	Zoneidnumber              int         `json:"zoneidnumber" gorm:"Column:zoneidnumber"`
	Version                   uint8       `json:"version" gorm:"Column:version"`
	ShortName                 null.String `json:"short_name" gorm:"Column:short_name"`
	LongName                  string      `json:"long_name" gorm:"Column:long_name"`
	MinStatus                 uint8       `json:"min_status" gorm:"Column:min_status"`
	MapFileName               null.String `json:"map_file_name" gorm:"Column:map_file_name"`
	Note                      null.String `json:"note" gorm:"Column:note"`
	MinExpansion              int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion              int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags              null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled      null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Expansion                 int8        `json:"expansion" gorm:"Column:expansion"`
	FileName                  null.String `json:"file_name" gorm:"Column:file_name"`
	SafeX                     float32     `json:"safe_x" gorm:"Column:safe_x"`
	SafeY                     float32     `json:"safe_y" gorm:"Column:safe_y"`
	SafeZ                     float32     `json:"safe_z" gorm:"Column:safe_z"`
	SafeHeading               float32     `json:"safe_heading" gorm:"Column:safe_heading"`
	GraveyardId               float32     `json:"graveyard_id" gorm:"Column:graveyard_id"`
	MinLevel                  uint8       `json:"min_level" gorm:"Column:min_level"`
	MaxLevel                  uint8       `json:"max_level" gorm:"Column:max_level"`
	Timezone                  int         `json:"timezone" gorm:"Column:timezone"`
	Maxclients                int         `json:"maxclients" gorm:"Column:maxclients"`
	Ruleset                   uint        `json:"ruleset" gorm:"Column:ruleset"`
	Underworld                float32     `json:"underworld" gorm:"Column:underworld"`
	Minclip                   float32     `json:"minclip" gorm:"Column:minclip"`
	Maxclip                   float32     `json:"maxclip" gorm:"Column:maxclip"`
	FogMinclip                float32     `json:"fog_minclip" gorm:"Column:fog_minclip"`
	FogMaxclip                float32     `json:"fog_maxclip" gorm:"Column:fog_maxclip"`
	FogBlue                   uint8       `json:"fog_blue" gorm:"Column:fog_blue"`
	FogRed                    uint8       `json:"fog_red" gorm:"Column:fog_red"`
	FogGreen                  uint8       `json:"fog_green" gorm:"Column:fog_green"`
	Sky                       uint8       `json:"sky" gorm:"Column:sky"`
	Ztype                     uint8       `json:"ztype" gorm:"Column:ztype"`
	ZoneExpMultiplier         float32     `json:"zone_exp_multiplier" gorm:"Column:zone_exp_multiplier"`
	Walkspeed                 float32     `json:"walkspeed" gorm:"Column:walkspeed"`
	TimeType                  uint8       `json:"time_type" gorm:"Column:time_type"`
	FogRed1                   uint8       `json:"fog_red_1" gorm:"Column:fog_red1"`
	FogGreen1                 uint8       `json:"fog_green_1" gorm:"Column:fog_green1"`
	FogBlue1                  uint8       `json:"fog_blue_1" gorm:"Column:fog_blue1"`
	FogMinclip1               float32     `json:"fog_minclip_1" gorm:"Column:fog_minclip1"`
	FogMaxclip1               float32     `json:"fog_maxclip_1" gorm:"Column:fog_maxclip1"`
	FogRed2                   uint8       `json:"fog_red_2" gorm:"Column:fog_red2"`
	FogGreen2                 uint8       `json:"fog_green_2" gorm:"Column:fog_green2"`
	FogBlue2                  uint8       `json:"fog_blue_2" gorm:"Column:fog_blue2"`
	FogMinclip2               float32     `json:"fog_minclip_2" gorm:"Column:fog_minclip2"`
	FogMaxclip2               float32     `json:"fog_maxclip_2" gorm:"Column:fog_maxclip2"`
	FogRed3                   uint8       `json:"fog_red_3" gorm:"Column:fog_red3"`
	FogGreen3                 uint8       `json:"fog_green_3" gorm:"Column:fog_green3"`
	FogBlue3                  uint8       `json:"fog_blue_3" gorm:"Column:fog_blue3"`
	FogMinclip3               float32     `json:"fog_minclip_3" gorm:"Column:fog_minclip3"`
	FogMaxclip3               float32     `json:"fog_maxclip_3" gorm:"Column:fog_maxclip3"`
	FogRed4                   uint8       `json:"fog_red_4" gorm:"Column:fog_red4"`
	FogGreen4                 uint8       `json:"fog_green_4" gorm:"Column:fog_green4"`
	FogBlue4                  uint8       `json:"fog_blue_4" gorm:"Column:fog_blue4"`
	FogMinclip4               float32     `json:"fog_minclip_4" gorm:"Column:fog_minclip4"`
	FogMaxclip4               float32     `json:"fog_maxclip_4" gorm:"Column:fog_maxclip4"`
	FogDensity                float32     `json:"fog_density" gorm:"Column:fog_density"`
	FlagNeeded                string      `json:"flag_needed" gorm:"Column:flag_needed"`
	Canbind                   int8        `json:"canbind" gorm:"Column:canbind"`
	Cancombat                 int8        `json:"cancombat" gorm:"Column:cancombat"`
	Canlevitate               int8        `json:"canlevitate" gorm:"Column:canlevitate"`
	Castoutdoor               int8        `json:"castoutdoor" gorm:"Column:castoutdoor"`
	Hotzone                   uint8       `json:"hotzone" gorm:"Column:hotzone"`
	Insttype                  uint8       `json:"insttype" gorm:"Column:insttype"`
	Shutdowndelay             uint64      `json:"shutdowndelay" gorm:"Column:shutdowndelay"`
	Peqzone                   int8        `json:"peqzone" gorm:"Column:peqzone"`
	BypassExpansionCheck      int8        `json:"bypass_expansion_check" gorm:"Column:bypass_expansion_check"`
	Suspendbuffs              uint8       `json:"suspendbuffs" gorm:"Column:suspendbuffs"`
	RainChance1               int         `json:"rain_chance_1" gorm:"Column:rain_chance1"`
	RainChance2               int         `json:"rain_chance_2" gorm:"Column:rain_chance2"`
	RainChance3               int         `json:"rain_chance_3" gorm:"Column:rain_chance3"`
	RainChance4               int         `json:"rain_chance_4" gorm:"Column:rain_chance4"`
	RainDuration1             int         `json:"rain_duration_1" gorm:"Column:rain_duration1"`
	RainDuration2             int         `json:"rain_duration_2" gorm:"Column:rain_duration2"`
	RainDuration3             int         `json:"rain_duration_3" gorm:"Column:rain_duration3"`
	RainDuration4             int         `json:"rain_duration_4" gorm:"Column:rain_duration4"`
	SnowChance1               int         `json:"snow_chance_1" gorm:"Column:snow_chance1"`
	SnowChance2               int         `json:"snow_chance_2" gorm:"Column:snow_chance2"`
	SnowChance3               int         `json:"snow_chance_3" gorm:"Column:snow_chance3"`
	SnowChance4               int         `json:"snow_chance_4" gorm:"Column:snow_chance4"`
	SnowDuration1             int         `json:"snow_duration_1" gorm:"Column:snow_duration1"`
	SnowDuration2             int         `json:"snow_duration_2" gorm:"Column:snow_duration2"`
	SnowDuration3             int         `json:"snow_duration_3" gorm:"Column:snow_duration3"`
	SnowDuration4             int         `json:"snow_duration_4" gorm:"Column:snow_duration4"`
	Gravity                   float32     `json:"gravity" gorm:"Column:gravity"`
	Type                      int         `json:"type" gorm:"Column:type"`
	Skylock                   int8        `json:"skylock" gorm:"Column:skylock"`
	FastRegenHp               int         `json:"fast_regen_hp" gorm:"Column:fast_regen_hp"`
	FastRegenMana             int         `json:"fast_regen_mana" gorm:"Column:fast_regen_mana"`
	FastRegenEndurance        int         `json:"fast_regen_endurance" gorm:"Column:fast_regen_endurance"`
	NpcMaxAggroDist           int         `json:"npc_max_aggro_dist" gorm:"Column:npc_max_aggro_dist"`
	ClientUpdateRange         int         `json:"client_update_range" gorm:"Column:client_update_range"`
	UnderworldTeleportIndex   int         `json:"underworld_teleport_index" gorm:"Column:underworld_teleport_index"`
	LavaDamage                null.Int    `json:"lava_damage" gorm:"Column:lava_damage"`
	MinLavaDamage             int         `json:"min_lava_damage" gorm:"Column:min_lava_damage"`
	IdleWhenEmpty             uint8       `json:"idle_when_empty" gorm:"Column:idle_when_empty"`
	SecondsBeforeIdle         uint        `json:"seconds_before_idle" gorm:"Column:seconds_before_idle"`
	ShardAtPlayerCount        null.Int    `json:"shard_at_player_count" gorm:"Column:shard_at_player_count"`
}

func (Zone) TableName() string {
    return "zone"
}

func (Zone) Relationships() []string {
    return []string{}
}

func (Zone) Connection() string {
    return "eqemu_content"
}
