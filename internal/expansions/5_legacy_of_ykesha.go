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
			Comment: "Shared bank enabled in Legacy of Ykesha",
		},
		{
			Name:    "Character:StatCap",
			Value:   "350",
			Comment: "After the increase to 305 in the \"Planes of Power\" expansion, the stat cap was further increased to 350 with the release of the \"Legacy of Ykesha\" expansion on February 25, 2003",
		},
	},
}
