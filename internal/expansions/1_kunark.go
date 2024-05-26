package expansions

var kunark = Expansion{
	ExpansionNumber: 1,
	ExpansionName:   "Ruins of Kunark",
	ShortName:       "RoK",
	MaxLevel:        60,
	ContentFlags:    []ContentFlag{},
	Description: `### *The Ruins of Kunark*

<table>
<tbody>
<tr class="odd">
<td style="font-size:110%; text-align:center;"><p><strong><em>EverQuest:
The Ruins of Kunark</em></strong><br />
</p></td>
</tr>
<tr class="even">
<td><p><strong><a href="Video_game_developer"
title="wikilink">Developer(s)</a></strong></p></td>
</tr>
<tr class="odd">
<td><p><strong><a href="Video_game_publisher"
title="wikilink">Publisher(s)</a></strong></p></td>
</tr>
<tr class="even">
<td><p><strong>Release</strong></p></td>
</tr>
<tr class="odd">
<td></td>
</tr>
</tbody>
</table>

The first expansion pack for *[EverQuest](EverQuest "wikilink")* was
*[The Ruins of Kunark](EverQuest:_The_Ruins_of_Kunark "wikilink")*,
released on April 24, 2000. It introduced the continent of Kunark to the
game, which had been previously unexplored. The storyline of the
discovery of Kunark was established through in-game events and fiction
published on the web by [Verant
Interactive](Verant_Interactive "wikilink").

In the United States, *The Ruins of Kunark* sold 92,172 units between
February 2001 through the first week of November. Desslock of
[GameSpot](GameSpot "wikilink") reported that the game and *The Scars of
Velious* "sold well early in the year, but sales evaporated during the
course of the summer, especially after the release of
*[Camelot](Dark_Age_of_Camelot "wikilink")*".[^1]

During the [4th Annual Interactive Achievement
Awards](4th_Annual_Interactive_Achievement_Awards "wikilink"), the
[Academy of Interactive Arts &
Sciences](Academy_of_Interactive_Arts_&_Sciences "wikilink") awarded
*The Ruins of Kunark* with the "Massive Multiplayer/Persistent World"
award, and received nominations for "PC Game of the Year" and "[Game of
the Year](D.I.C.E._Award_for_Game_of_the_Year "wikilink")".[^2]`,
	Rules: []Rule{
		{
			Name:    "Expansion:CurrentExpansion",
			Value:   "1",
			Comment: "Current Expansion",
		},
		{
			Name:    "World:ExpansionSettings",
			Value:   "1",
			Comment: "Kunark Client-Based Expansion Setting",
		},
		{
			Name:    "World:CharacterSelectExpansionSettings",
			Value:   "1",
			Comment: "Kunark Client-Based Expansion Setting",
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
