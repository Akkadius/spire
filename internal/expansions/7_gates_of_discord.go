package expansions

var gatesOfDiscord = Expansion{
	ExpansionNumber: 7,
	ExpansionName:   "Gates of Discord",
	ShortName:       "GoD",
	MaxLevel:        65,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "7",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "127",
			Comment: "Gates of Discord Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "127",
			Comment: "Gates of Discord Client-Based Expansion Setting",
		},
	},
}
