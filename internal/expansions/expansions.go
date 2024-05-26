package expansions

type Expansion struct {
	ExpansionNumber int           `json:"expansion_number"`
	ExpansionName   string        `json:"expansion_name"`
	ShortName       string        `json:"short_name"`
	MaxLevel        int           `json:"max_level"`
	Description     string        `json:"description"`
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
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "11",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 12,
		ExpansionName:   "The Serpent's Spine",
		ShortName:       "TSS",
		MaxLevel:        75,
		Rules: []Rule{
			{
				Name:    "Character:MaxLevel",
				Value:   "75",
				Comment: "Level 75 cap until Secrets of Faydwer",
			},
			{
				Name:    "Character:MaxExpLevel",
				Value:   "75",
				Comment: "Level 75 cap until Secrets of Faydwer",
			},
			{
				Name:    "Character:RestRegenEndurance",
				Value:   "true",
				Comment: "Rest Regen Endurance was added in The Serpent's Spine",
			},
			{
				Name:    "Character:RestRegenEnabled",
				Value:   "true",
				Comment: "Rest Regen was added in The Serpent's Spine",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "12",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 13,
		ExpansionName:   "The Buried Sea",
		ShortName:       "TBS",
		MaxLevel:        75,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "13",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 14,
		ExpansionName:   "Secrets of Faydwer",
		ShortName:       "SoF",
		MaxLevel:        80,
		Rules: []Rule{
			{
				Name:    "Character:MaxLevel",
				Value:   "80",
				Comment: "Level 80 cap until Seeds of Destruction",
			},
			{
				Name:    "Character:MaxExpLevel",
				Value:   "80",
				Comment: "Level 80 cap until Seeds of Destruction",
			},
			{
				Name:    "Character:UseOldConSystem",
				Value:   "false",
				Comment: "Con system changed in Secrets of Faydwer",
			},
			{
				Name:    "Character:UseOldClassExpPenalties",
				Value:   "false",
				Comment: "Experience penalties were removed in Secrets of Faydwer",
			},
			{
				Name:    "Character:UseOldRaceExpPenalties",
				Value:   "false",
				Comment: "Experience penalties were removed in Secrets of Faydwer",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "14",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 15,
		ExpansionName:   "Seeds of Destruction",
		ShortName:       "SoD",
		MaxLevel:        85,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "15",
				Comment: "Current Expansion",
			},
			{
				Name:    "Character:LeaveNakedCorpses",
				Value:   "true",
				Comment: "Gear is no longer left on your corpse when you die in Seeds of Destruction",
			},
		},
	},
	{
		ExpansionNumber: 16,
		ExpansionName:   "Underfoot",
		ShortName:       "UF",
		MaxLevel:        85,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "16",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 17,
		ExpansionName:   "House of Thule",
		ShortName:       "HoT",
		MaxLevel:        90,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "17",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 18,
		ExpansionName:   "Veil of Alaris",
		ShortName:       "VoA",
		MaxLevel:        95,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "18",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 19,
		ExpansionName:   "Rain of Fear",
		ShortName:       "RoF",
		MaxLevel:        100,
		Rules: []Rule{
			{
				Name:    "Character:StatCap",
				Value:   "500",
				Comment: "After the increase to 400 in the \"Lost Dungeons of Norrath\" expansion, the stat cap remained at 400 for quite some time. The next significant increase occurred much later, with the \"Rain of Fear\" expansion on November 28, 2012, which raised the stat cap to 500. ",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "19",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 20,
		ExpansionName:   "Call of the Forsaken",
		ShortName:       "CoTF",
		MaxLevel:        100,
		Rules: []Rule{
			{
				Name:    "Character:StatCap",
				Value:   "600",
				Comment: "After the increase to 500 in the \"Rain of Fear\" expansion, the stat cap was further increased to 600 with the release of the \"Call of the Forsaken\" expansion on October 8, 2013.",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "20",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 21,
		ExpansionName:   "The Darkened Sea",
		ShortName:       "TDS",
		MaxLevel:        105,
		Rules: []Rule{
			{
				Name:    "Character:StatCap",
				Value:   "700",
				Comment: "After the increase to 600 in the \"Call of the Forsaken\" expansion, the stat cap was increased to 700 with the release of the \"The Darkened Sea\" expansion on October 28, 2014.",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "21",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 22,
		ExpansionName:   "The Broken Mirror",
		ShortName:       "TBM",
		MaxLevel:        105,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "22",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 23,
		ExpansionName:   "Empires of Kunark",
		ShortName:       "EoK",
		MaxLevel:        105,
		Rules: []Rule{
			{
				Name:    "Character:StatCap",
				Value:   "800",
				Comment: "After the increase to 700 in \"The Darkened Sea\" expansion, the stat cap was further increased to 800 with the release of the \"Empires of Kunark\" expansion on November 16, 2016.",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "23",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 24,
		ExpansionName:   "Ring of Scale",
		ShortName:       "RoS",
		MaxLevel:        110,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "24",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 25,
		ExpansionName:   "The Burning Lands",
		ShortName:       "TBL",
		MaxLevel:        110,
		Rules: []Rule{
			{
				Name:    "Character:StatCap",
				Value:   "1100",
				Comment: "After the increase to 800 in the \"Empires of Kunark\" expansion, the stat cap remained at that level for quite some time. The next significant increase occurred with the \"The Burning Lands\" expansion on December 11, 2018, which raised the stat cap to 1100",
			},
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "25",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 26,
		ExpansionName:   "Torment of Velious",
		ShortName:       "ToV",
		MaxLevel:        115,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "26",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 27,
		ExpansionName:   "Claws of Veeshan",
		ShortName:       "CoV",
		MaxLevel:        115,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "27",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 28,
		ExpansionName:   "Terror of Luclin",
		ShortName:       "ToL",
		MaxLevel:        120,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "28",
				Comment: "Current Expansion",
			},
		},
	},
	{
		ExpansionNumber: 29,
		ExpansionName:   "Night of Shadows",
		ShortName:       "NoS",
		MaxLevel:        120,
		Rules: []Rule{
			{
				Name:    "Expansion:CurrentExpansion",
				Value:   "29",
				Comment: "Current Expansion",
			},
		},
	},
}
