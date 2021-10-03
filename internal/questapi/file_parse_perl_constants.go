package questapi

import (
	"strings"
)

type PerlConstants struct {
	Constant string `json:"constant"` // @eg $variable
}

func (c *ParseService) parsePerlConstants(files map[string]string) map[string][]PerlConstants {

	typeLabel := "" // Mob, Client, Item etc.
	perlConstants := map[string][]PerlConstants{}
	for fileName, contents := range files {
		if strings.Contains(fileName, ".cpp") {
			if strings.Contains(contents, "ExportVar") {
				for _, l := range strings.Split(contents, "\n") {

					// stop a listening buffer when we see another function
					if strings.Contains(l, "void PerlembParser::") && typeLabel != "" {
						typeLabel = ""
					}

					// start a listening buffer when we're inside an export block
					// @eg void PerlembParser::ExportZoneVariables(std::string &package_name)
					if strings.Contains(l, "PerlembParser::Export") && strings.Contains(l, "Variables") {
						//fmt.Printf("%+v\n", l)

						labelSplit := strings.Split(l, "PerlembParser::Export")
						if len(labelSplit) > 0 {
							typeSplit := strings.Split(labelSplit[1], "Variables")
							if len(typeSplit) > 0 {
								typeLabel = strings.TrimSpace(typeSplit[0])
							}
						}

						//fmt.Printf("\n[%+v]\n", typeLabel)
					}

					// Events and their exports are covered in another parse function
					// @eg ExportVar(package_name.c_str(), "race", GetRaceIDName(mob->GetRace()));
					if strings.Contains(l, "ExportVar") && typeLabel != "" && typeLabel != "Event" {
						//fmt.Println(l)
						//fmt.Println(typeLabel)

						variable := ""
						varSplit := strings.Split(l, "\"")
						if len(varSplit) > 0 {
							variable = "$" + strings.TrimSpace(varSplit[1])
						}

						perlConstants[typeLabel] = append(
							perlConstants[typeLabel], PerlConstants{
								Constant: variable,
							},
						)
					}
				}
			}
		}
	}

	return perlConstants
}
