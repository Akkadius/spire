package expansions

var lostDungeonsOfNorrath = Expansion{
	ExpansionNumber: 6,
	ExpansionName:   "Lost Dungeons of Norrath",
	ShortName:       "LDoN",
	MaxLevel:        65,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "6",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "63",
			Comment: "Lost Dungeons of Norrath Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "63",
			Comment: "Lost Dungeons of Norrath Client-Based Expansion Setting",
		},
	},
}
