package expansions

var depthsOfDarkhollow = Expansion{
	ExpansionNumber: 10,
	ExpansionName:   "Depths of Darkhollow",
	ShortName:       "DoDH",
	MaxLevel:        70,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "10",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "1023",
			Comment: "Depths of Darkhollow Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "1023",
			Comment: "Depths of Darkhollow Client-Based Expansion Setting",
		},
	},
}
