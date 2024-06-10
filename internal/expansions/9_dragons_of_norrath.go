package expansions

var dragonsOfNorrath = Expansion{
	ExpansionNumber: 9,
	ExpansionName:   "Dragons of Norrath",
	ShortName:       "DoN",
	MaxLevel:        70,
	ContentFlags: []ContentFlag{
		{
			Name:    "don_nest",
			Enabled: false,
			Comment: "Dragons of Norrath - Unlocked Accursed Nest. This is naturally unlocked during the expansion.",
		},
	},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "9",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "511",
			Comment: "Dragons of Norrath Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "511",
			Comment: "Dragons of Norrath Client-Based Expansion Setting",
		},
		{
			Name:    "Mail:EnableMailSystem",
			Value:   "true",
			Comment: "Mail System not added until Dragons of Norrath",
		},
	},
}
