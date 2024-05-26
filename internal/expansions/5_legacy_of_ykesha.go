package expansions

var legacyOfYkesha = Expansion{
	ExpansionNumber: 5,
	ExpansionName:   "Legacy of Ykesha",
	ShortName:       "LoY",
	MaxLevel:        65,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "5",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "31",
			Comment: "Legacy of Ykesha Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "31",
			Comment: "Legacy of Ykesha Client-Based Expansion Setting",
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
		{
			Name:    "Character:SharedBankPlat",
			Value:   "true",
			Comment: "Shared bank is now enabled",
		},
	},
}
