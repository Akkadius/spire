package expansions

var classic = Expansion{
	ExpansionNumber: 0,
	ExpansionName:   "Original (Classic)",
	ShortName:       "",
	MaxLevel:        50,
	ContentFlags:    []ContentFlag{},
	Rules: []Rule{
		{
			Name:    "World:ExpansionSettings",
			Value:   "0",
			Comment: "Classic Client-Based Expansion Setting",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "0",
			Comment: "Classic Client-Based Expansion Setting",
		},
		{
			Name:    "Character:UseOldRaceRezEffects",
			Value:   "true",
			Comment: "this may need more testing with the spell file to ensure it cant be dispelled",
		},
		{
			Name:    "Character:LeaveCorpses",
			Value:   "true",
			Comment: "Leave corpses behind",
		},
		{
			Name:    "Character:LeaveNakedCorpses",
			Value:   "false",
			Comment: "Gear left on Corpses until Seeds of Destruction",
		},
		{
			Name:    "Spells:PreNerfBardAEDoT",
			Value:   "true",
			Comment: "Bard AE Nerf Not added until Omens of War",
		},
		{
			Name:    "Chat:ServerWideOOC",
			Value:   "false",
			Comment: "Disable Server-Wide OOC Chat",
		},
		{
			Name:    "Chat:ServerWideAuction",
			Value:   "false",
			Comment: "Disable Server-Wide Auction Chat",
		},
		{
			Name:    "Spells:WizCritLevel",
			Value:   "65",
			Comment: "Wizard non-AA Criticals Not added until Luclin",
		},
		{
			Name:    "Character:EnableXTargetting",
			Value:   "false",
			Comment: "Was introduced in the Call of the Forsaken expansion",
		},
		{
			Name:    "Character:MaxExpLevel",
			Value:   "50",
			Comment: "Level 50 cap until Kunark",
		},
		{
			Name:    "Character:MaxLevel",
			Value:   "50",
			Comment: "Level 50 cap until Kunark",
		},
		{
			Name:    "Character:RestRegenEnabled",
			Value:   "false",
			Comment: "OoC regen Not added until The Serpent's Spine",
		},
		{
			Name:    "Character:RestRegenEndurance",
			Value:   "false",
			Comment: "OoC regen Not added until The Serpent's Spine",
		},
		{
			Name:    "Character:SharedBankPlat",
			Value:   "false",
			Comment: "No shared bank until Legacy of Ykesha",
		},
		{
			Name:    "Character:StatCap",
			Value:   "255",
			Comment: "Classic cap is 255",
		},
		{
			Name:    "Character:UseOldBindWound",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Character:UseOldClassExpPenalties",
			Value:   "true",
			Comment: "Experience penalties were removed in Secrets of Faydwer",
		},
		{
			Name:    "Character:UseOldConSystem",
			Value:   "true",
			Comment: "Con system changed in Secrets of Faydwer",
		},
		{
			Name:    "Character:UseOldRaceExpPenalties",
			Value:   "true",
			Comment: "Experience penalties were removed in Secrets of Faydwer",
		},
		{
			Name:    "Character:UseRaceClassExpBonuses",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Chat:EnableVoiceMacros",
			Value:   "false",
			Comment: "Introduced in Gates of Discord",
		},
		{
			Name:    "Mail:EnableMailSystem",
			Value:   "false",
			Comment: "Mail System not added until Dragons of Norrath",
		},
		{
			Name:    "TaskSystem:EnableTaskSystem",
			Value:   "false",
			Comment: "Task system was introduced in Omens of War",
		},
		{
			Name:    "World:EnableReturnHomeButton",
			Value:   "false",
			Comment: "Return Home Button Not added until Dragons of Norrath",
		},
		{
			Name:    "World:EnableTutorialButton",
			Value:   "false",
			Comment: "Tutorial Not added until Dragons of Norrath",
		},
		{
			Name:    "Combat:ClassicNPCBackstab",
			Value:   "true",
			Comment: "Disables front backstab",
		},
	},
}
