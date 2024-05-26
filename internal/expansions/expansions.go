package expansions

type Expansion struct {
	ExpansionNumber int           `json:"expansion_number"`
	ExpansionName   string        `json:"expansion_name"`
	ShortName       string        `json:"short_name"`
	MaxLevel        int           `json:"max_level"`
	ContentFlags    []ContentFlag `json:"content_flags"`
	Rules           []Rule        `json:"rules"`
}

type ContentFlag struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Comment string `json:"comment"`
}

type Rule struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
}

var expansions = []Expansion{
	classic,
	kunark,
	velious,
	luclin,
	planesOfPower,
	legacyOfYkesha,
	lostDungeonsOfNorrath,
	gatesOfDiscord,
	omensOfWar,
	dragonsOfNorrath,
	depthsOfDarkhollow,
	{
		ExpansionNumber: 11,
		ExpansionName:   "Prophecy of Ro",
		ShortName:       "PoR",
		MaxLevel:        70,
	},
	{
		ExpansionNumber: 12,
		ExpansionName:   "The Serpent's Spine",
		ShortName:       "TSS",
		MaxLevel:        75,
	},
	{
		ExpansionNumber: 13,
		ExpansionName:   "The Buried Sea",
		ShortName:       "TBS",
		MaxLevel:        75,
	},
	{
		ExpansionNumber: 14,
		ExpansionName:   "Secrets of Faydwer",
		ShortName:       "SoF",
		MaxLevel:        80,
	},
	{
		ExpansionNumber: 15,
		ExpansionName:   "Seeds of Destruction",
		ShortName:       "SoD",
		MaxLevel:        85,
	},
	{
		ExpansionNumber: 16,
		ExpansionName:   "Underfoot",
		ShortName:       "UF",
		MaxLevel:        85,
	},
	{
		ExpansionNumber: 17,
		ExpansionName:   "House of Thule",
		ShortName:       "HoT",
		MaxLevel:        90,
	},
	{
		ExpansionNumber: 18,
		ExpansionName:   "Veil of Alaris",
		ShortName:       "VoA",
		MaxLevel:        95,
	},
	{
		ExpansionNumber: 19,
		ExpansionName:   "Rain of Fear",
		ShortName:       "RoF",
		MaxLevel:        100,
	},
	{
		ExpansionNumber: 20,
		ExpansionName:   "Call of the Forsaken",
		ShortName:       "CoTF",
		MaxLevel:        100,
	},
	{
		ExpansionNumber: 21,
		ExpansionName:   "The Darkened Sea",
		ShortName:       "TDS",
		MaxLevel:        105,
	},
	{
		ExpansionNumber: 22,
		ExpansionName:   "The Broken Mirror",
		ShortName:       "TBM",
		MaxLevel:        105,
	},
	{
		ExpansionNumber: 23,
		ExpansionName:   "Empires of Kunark",
		ShortName:       "EoK",
		MaxLevel:        105,
	},
	{
		ExpansionNumber: 24,
		ExpansionName:   "Ring of Scale",
		ShortName:       "RoS",
		MaxLevel:        110,
	},
	{
		ExpansionNumber: 25,
		ExpansionName:   "The Burning Lands",
		ShortName:       "TBL",
		MaxLevel:        110,
	},
	{
		ExpansionNumber: 26,
		ExpansionName:   "Torment of Velious",
		ShortName:       "ToV",
		MaxLevel:        115,
	},
	{
		ExpansionNumber: 27,
		ExpansionName:   "Claws of Veeshan",
		ShortName:       "CoV",
		MaxLevel:        115,
	},
	{
		ExpansionNumber: 28,
		ExpansionName:   "Terror of Luclin",
		ShortName:       "ToL",
		MaxLevel:        120,
	},
	{
		ExpansionNumber: 29,
		ExpansionName:   "Night of Shadows",
		ShortName:       "NoS",
		MaxLevel:        120,
	},
}
