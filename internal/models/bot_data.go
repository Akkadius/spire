package models

type BotDatum struct {
	BotId             uint    `json:"bot_id" gorm:"Column:bot_id"`
	OwnerId           uint    `json:"owner_id" gorm:"Column:owner_id"`
	SpellsId          uint    `json:"spells_id" gorm:"Column:spells_id"`
	Name              string  `json:"name" gorm:"Column:name"`
	LastName          string  `json:"last_name" gorm:"Column:last_name"`
	Title             string  `json:"title" gorm:"Column:title"`
	Suffix            string  `json:"suffix" gorm:"Column:suffix"`
	ZoneId            int16   `json:"zone_id" gorm:"Column:zone_id"`
	Gender            int8    `json:"gender" gorm:"Column:gender"`
	Race              int16   `json:"race" gorm:"Column:race"`
	Class             int8    `json:"class" gorm:"Column:class"`
	Level             uint8   `json:"level" gorm:"Column:level"`
	Deity             uint    `json:"deity" gorm:"Column:deity"`
	CreationDay       uint    `json:"creation_day" gorm:"Column:creation_day"`
	LastSpawn         uint    `json:"last_spawn" gorm:"Column:last_spawn"`
	TimeSpawned       uint    `json:"time_spawned" gorm:"Column:time_spawned"`
	Size              float32 `json:"size" gorm:"Column:size"`
	Face              int     `json:"face" gorm:"Column:face"`
	HairColor         int     `json:"hair_color" gorm:"Column:hair_color"`
	HairStyle         int     `json:"hair_style" gorm:"Column:hair_style"`
	Beard             int     `json:"beard" gorm:"Column:beard"`
	BeardColor        int     `json:"beard_color" gorm:"Column:beard_color"`
	EyeColor1         int     `json:"eye_color_1" gorm:"Column:eye_color_1"`
	EyeColor2         int     `json:"eye_color_2" gorm:"Column:eye_color_2"`
	DrakkinHeritage   int     `json:"drakkin_heritage" gorm:"Column:drakkin_heritage"`
	DrakkinTattoo     int     `json:"drakkin_tattoo" gorm:"Column:drakkin_tattoo"`
	DrakkinDetails    int     `json:"drakkin_details" gorm:"Column:drakkin_details"`
	Ac                int16   `json:"ac" gorm:"Column:ac"`
	Atk               int32   `json:"atk" gorm:"Column:atk"`
	Hp                int     `json:"hp" gorm:"Column:hp"`
	Mana              int     `json:"mana" gorm:"Column:mana"`
	Str               int32   `json:"str" gorm:"Column:str"`
	Sta               int32   `json:"sta" gorm:"Column:sta"`
	Cha               int32   `json:"cha" gorm:"Column:cha"`
	Dex               int32   `json:"dex" gorm:"Column:dex"`
	Int               int32   `json:"int" gorm:"Column:int"`
	Agi               int32   `json:"agi" gorm:"Column:agi"`
	Wis               int32   `json:"wis" gorm:"Column:wis"`
	Fire              int16   `json:"fire" gorm:"Column:fire"`
	Cold              int16   `json:"cold" gorm:"Column:cold"`
	Magic             int16   `json:"magic" gorm:"Column:magic"`
	Poison            int16   `json:"poison" gorm:"Column:poison"`
	Disease           int16   `json:"disease" gorm:"Column:disease"`
	Corruption        int16   `json:"corruption" gorm:"Column:corruption"`
	ShowHelm          uint    `json:"show_helm" gorm:"Column:show_helm"`
	FollowDistance    uint    `json:"follow_distance" gorm:"Column:follow_distance"`
	StopMeleeLevel    uint8   `json:"stop_melee_level" gorm:"Column:stop_melee_level"`
	ExpansionBitmask  int     `json:"expansion_bitmask" gorm:"Column:expansion_bitmask"`
}

func (BotDatum) TableName() string {
    return "bot_data"
}

func (BotDatum) Relationships() []string {
    return []string{}
}

func (BotDatum) Connection() string {
    return ""
}
