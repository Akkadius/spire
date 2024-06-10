package expansions

var omensOfWar = Expansion{
	ExpansionNumber: 8,
	ExpansionName:   "Omens of War",
	ShortName:       "OoW",
	MaxLevel:        70,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "8",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "255",
			Comment: "Omens of War Client-Based Expansion Setting",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "255",
			Comment: "Omens of War Client-Based Expansion Setting",
		},
		{
			Name:    "Spells:PreNerfBardAEDoT",
			Value:   "false",
			Comment: "Bard AE Nerf",
		},
		{
			Name:    "TaskSystem:EnableTaskSystem",
			Value:   "true",
			Comment: "Task system was introduced in Omens of War",
		},
		{
			Name:    "Character:MaxExpLevel",
			Value:   "70",
			Comment: "Level 70 cap until The Serpent's Spine",
		},
		{
			Name:    "Character:MaxLevel",
			Value:   "70",
			Comment: "Level 70 cap until The Serpent's Spine",
		},
	},
}
