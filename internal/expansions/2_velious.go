package expansions

var velious = Expansion{
	ExpansionNumber: 2,
	ExpansionName:   "Scars of Velious",
	ShortName:       "SoV",
	MaxLevel:        60,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "2",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "3",
			Comment: "Velious Client-Based Expansion Setting",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "3",
			Comment: "Velious Client-Based Expansion Setting",
		},
		{
			Name:    "Character:MaxExpLevel",
			Value:   "60",
			Comment: "Level 60 cap until PoP",
		},
		{
			Name:    "Character:MaxLevel",
			Value:   "60",
			Comment: "Level 60 cap until PoP",
		},
	},
}
