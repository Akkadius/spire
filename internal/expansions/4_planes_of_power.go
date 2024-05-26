package expansions

var planesOfPower = Expansion{
	ExpansionNumber: 4,
	ExpansionName:   "Planes of Power",
	ShortName:       "PoP",
	MaxLevel:        65,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		// translate above SQL to Go
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "4",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "15",
			Comment: "Planes of Power Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "15",
			Comment: "Planes of Power Client-Based Expansion Setting",
		},
		{
			Name:    "Character:MaxExpLevel",
			Value:   "65",
			Comment: "Level 65 cap until Omens of War",
		},
		{
			Name:    "Character:MaxLevel",
			Value:   "65",
			Comment: "Level 65 cap until Omens of War",
		},
	},
}
