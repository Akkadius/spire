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
			Comment: "Gear left on Corpses until post-OoW",
		},
		{
			Name:    "Character:LeaveNakedCorpses",
			Value:   "false",
			Comment: "Gear left on Corpses until post-OoW",
		},
		{
			Name:    "Spells:PreNerfBardAEDoT",
			Value:   "true",
			Comment: "Bard AE Nerf Not added until OoW",
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
			Comment: "X Targeting is not classic",
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
			Comment: "OoC regen Not added until TSS",
		},
		{
			Name:    "Character:RestRegenEndurance",
			Value:   "false",
			Comment: "OoC regen Not added until TSS",
		},
		{
			Name:    "Character:SharedBankPlat",
			Value:   "false",
			Comment: "No shared bank until LoY",
		},
		{
			Name:    "Character:StatCap",
			Value:   "255",
			Comment: "Stat caps at 255 on classic, I believe?",
		},
		{
			Name:    "Character:UseOldBindWound",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Character:UseOldClassExpPenalties",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Character:UseOldConSystem",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Character:UseOldRaceExpPenalties",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Character:UseRaceClassExpBonuses",
			Value:   "true",
			Comment: "",
		},
		{
			Name:    "Chat:EnableVoiceMacros",
			Value:   "false",
			Comment: "not added until later",
		},
		{
			Name:    "Mail:EnableMailSystem",
			Value:   "false",
			Comment: "Mail System not added until DoN",
		},
		{
			Name:    "TaskSystem:EnableTaskSystem",
			Value:   "false",
			Comment: "Task System Not added until OoW",
		},
		{
			Name:    "World:EnableReturnHomeButton",
			Value:   "false",
			Comment: "Return Home Button Not added until DoN",
		},
		{
			Name:    "World:EnableTutorialButton",
			Value:   "false",
			Comment: "Tutorial Not added until DoN",
		},
		{
			Name:    "Combat:ClassicNPCBackstab",
			Value:   "true",
			Comment: "Disables front backstab",
		},
	},
}
