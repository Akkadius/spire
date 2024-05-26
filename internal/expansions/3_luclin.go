package expansions

var luclin = Expansion{
	ExpansionNumber: 3,
	ExpansionName:   "Shadows of Luclin",
	ShortName:       "SoL",
	MaxLevel:        60,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "3",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "7",
			Comment: "Luclin Client-Based Expansion Setting",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "7",
			Comment: "Luclin Client-Based Expansion Setting",
		},
		{
			Name:    "Spells:WizCritLevel",
			Value:   "12",
			Comment: "Wizard non-AA Criticals Not added until Luclin",
		},
	},
}
