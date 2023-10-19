package eqtraders

type ObjectType struct {
	Type         int      `json:"type"`
	Name         string   `json:"name"`
	Skill        int      `json:"skill"`
	TradersNames []string `json:"traders_names"`
}

var objectTypes = []ObjectType{
	{Type: 10, Name: "Tinkering", Skill: 57, TradersNames: []string{"Tinkering Table"}},
	{Type: 12, Name: "Mortar and Pestle", Skill: 56, TradersNames: []string{"Poisoncrafting Table"}},
	{Type: 15, Name: "Baking", Skill: 60, TradersNames: []string{"Oven", "Tanaan Oven", "Spit", "Ice Cream Churn", "Mixing Bowl"}},
	{Type: 16, Name: "Tailoring", Skill: 61, TradersNames: []string{"Loom", "Tanaan Loom"}},
	{Type: 17, Name: "Blacksmithing", Skill: 63, TradersNames: []string{"Tanaan Forge", "Half Elf Forge", "Crystalwing Forge", "Shar Vahl Forge"}},
	{Type: 18, Name: "Fletching", Skill: 64},
	{Type: 19, Name: "Brewing", Skill: 65, TradersNames: []string{"Brew Barrel", "Tanaan Brew Barrel"}},
	{Type: 20, Name: "Jewelry Making", Skill: 68},
	{Type: 21, Name: "Pottery", Skill: 69},
	{Type: 22, Name: "Kiln", Skill: 69},
	{Type: 24, Name: "Lexicon Wiz", Skill: 58},
	{Type: 25, Name: "Lexicon Mage", Skill: 58, TradersNames: []string{"Spell Research Table"}}, // observed in PoK
	{Type: 26, Name: "Lexicon Nec", Skill: 58},
	{Type: 27, Name: "Lexicon Enc", Skill: 58},
	{Type: 29, Name: "Lexicon Practice", Skill: 58},
	{Type: 30, Name: "Alchemy", Skill: 59},
	{Type: 31, Name: "High Elf Forge", Skill: 63, TradersNames: []string{"Koada&#39;dal Forge"}},
	{Type: 32, Name: "Dark Elf Forge", Skill: 63, TradersNames: []string{"Teir`Dal Forge"}},
	{Type: 33, Name: "Ogre Forge", Skill: 63, TradersNames: []string{"Ogre Forge"}},
	{Type: 34, Name: "Dwarf Forge", Skill: 63, TradersNames: []string{"Stormguard Forge"}},
	{Type: 35, Name: "Gnome Forge", Skill: 63, TradersNames: []string{"Clockwork Forge"}},
	{Type: 36, Name: "Barbarian Forge", Skill: 63, TradersNames: []string{"Northman Forge"}},
	{Type: 37, Name: "Iksar Forge", Skill: 63},
	{Type: 38, Name: "Iksar Forge", Skill: 63, TradersNames: []string{"Iksar Forge"}},
	{Type: 39, Name: "Human Forge", Skill: 63},
	{Type: 40, Name: "Halfling Forge", Skill: 63, TradersNames: []string{"Antonican Forge"}},
	//{Type: 41, Name: "Halfling Forge", Skill: 63},
	//{Type: 42, Name: "Erudite Forge", Skill: 63},
	//{Type: 43, Name: "Wood Elf Forge", Skill: 63},
	//{Type: 44, Name: "Wood Elf Forge", Skill: 63},
	{Type: 45, Name: "Iksar Pottery", Skill: 69},
	{Type: 47, Name: "Troll Forge", Skill: 63, TradersNames: []string{"Troll Forge"}},
	{Type: 48, Name: "Wood Elf Forge", Skill: 63, TradersNames: []string{"Feir`Dal Forge"}},
	{Type: 49, Name: "Halfling Forge", Skill: 63, TradersNames: []string{"Vale Forge"}},
	{Type: 50, Name: "Erudite Forge", Skill: 63, TradersNames: []string{"Erud Forge"}},
	{Type: 52, Name: "Froglok Forge", Skill: 63, TradersNames: []string{"Froglok Forge", "Guktan Forge"}},
}
