package questapi

import (
	"strings"
)

type LuaConstants struct {
	Constant string `json:"constant"` // @eg var
}

func (c *ParseService) parseLuaConstants(files map[string]string) map[string][]LuaConstants {

	typeLabel := "" // Mob, Client, Item etc.
	isInConstants := false
	luaConstants := map[string][]LuaConstants{}
	for fileName, contents := range files {

		// regular constants
		if strings.Contains(fileName, ".cpp") {

			// @grep luabind::class_
			// @example
			// luabind::scope lua_register_special_abilities() {
			//	return luabind::class_<SpecialAbilities>("SpecialAbility")
			//		.enum_("constants")
			if strings.Contains(contents, "luabind::class_") {
				for _, l := range strings.Split(contents, "\n") {

					// stop a listening buffer when we see another function
					if strings.Contains(l, "luabind::class_") && typeLabel != "" {
						typeLabel = ""
						isInConstants = false
					}

					// luabind::class_ is not specific enough to grep for, we need to make sure we're
					// inside an enum list
					// start a listening buffer when we're inside a constants block
					if strings.Contains(l, ".enum_(\"constants\")") && typeLabel != "" {
						isInConstants = true
					}

					// @eg void LuaembParser::ExportZoneVariables(std::string &package_name)
					if strings.Contains(l, "luabind::class_") {
						labelSplit := strings.Split(l, "\"")
						if len(labelSplit) > 0 {
							typeLabel = strings.TrimSpace(labelSplit[1])
						}

						//fmt.Printf("\n[%+v]\n", typeLabel)
					}

					// Events and their exports are covered in another parse function
					// @eg luabind::value("summon", static_cast<int>(SPECATK_SUMMON)),
					if strings.Contains(l, "luabind::value") && strings.Contains(
						l,
						"\"",
					) && typeLabel != "" && isInConstants {
						//fmt.Println(l)
						//fmt.Println(typeLabel)

						variable := ""
						varSplit := strings.Split(l, "\"")
						if len(varSplit) > 1 {
							variable = strings.TrimSpace(varSplit[1])
						}

						//fmt.Printf("[%v.%v]\n", typeLabel, variable)

						luaConstants[typeLabel] = append(
							luaConstants[typeLabel], LuaConstants{
								Constant: variable,
							},
						)
					}
				}
			}
		}

		// ruletypes.h needs to be exploded / super-imposed
		if strings.Contains(fileName, "ruletypes.h") {
			for _, l := range strings.Split(contents, "\n") {
				// @example RULE_INT(Adventure, MinNumberForGroup, 2, "Minimum members for adventure group")
				if strings.Contains(l, "RULE_") && strings.Contains(l, "\"") {

					// @grep RULE_INT(Expedition, WorldExpeditionProcessRateMS, 6000, "Timer interval (ms) that world checks expedition states")
					// @result WorldExpeditionProcessRateMS
					// this could change per DM with KLS that we should also have the category with the constant
					// IE: Expedition.WorldExpeditionProcessRateMS
					variable := ""
					varSplit := strings.Split(l, ",")
					if len(varSplit) > 1 {
						variable = strings.TrimSpace(varSplit[1])
					}

					//fmt.Printf("[%v.%v]\n", "Rule", variable)

					luaConstants["Rule"] = append(
						luaConstants["Rule"], LuaConstants{
							Constant: variable,
						},
					)
				}
			}
		}
	}

	return luaConstants
}
