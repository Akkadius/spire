package expansions

var planesOfPower = Expansion{
	ExpansionNumber: 4,
	ExpansionName:   "Planes of Power",
	ShortName:       "PoP",
	MaxLevel:        65,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
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
		{
			Name:    "Character:StatCap",
			Value:   "305",
			Comment: "The stat cap in EverQuest increased from 255 to 305 with the release of the \"Planes of Power\" expansion on October 29, 2002",
		},
	},
}
