package questapi

import (
	"strings"
)

// parses perl methods
func (c *ParseService) parsePerlMethods(contents string, perlMethods map[string][]PerlMethod) {
	// first pass over file to build a list of definitions and map then to the actual functions themselves
	// definitions are not always 1:1 with the actual source function names
	perlDefinitionMap := map[string]string{}

	for _, l := range strings.Split(contents, "\n") {
		// Check if line is package.add for Perl function name versus source function name
		if strings.Contains(l, "package.add") {
			l = strings.TrimSpace(l)
			l = strings.ReplaceAll(l, "package.add(\"", "")

			// split to pull out definition name
			// ex: package.add("RemoveTrap", &Perl_EntityList_RemoveTrap)
			// get: RemoveTrap
			definitionSplit := strings.Split(l, "\",")
			definition := strings.TrimSpace(definitionSplit[0])

			// split to pull out method name to key off of
			// ex: RemoveTrap", &Perl_EntityList_RemoveTrap)
			// get: EntityList_RemoveTrap
			methodSplit := strings.Split(l, "&Perl_")
			if len(methodSplit) > 1 {
				if len(methodSplit) > 1 {
					methodName := strings.TrimSpace(strings.ReplaceAll(methodSplit[len(methodSplit)-1], ");", ""))

					methodName = strings.TrimPrefix(methodName, "_")

					// add to map
					perlDefinitionMap[methodName] = definition
				}
			}
		}
	}

	for _, l := range strings.Split(contents, "\n") {
		if strings.Contains(l, "const char *") {
			l = strings.ReplaceAll(l, "const char *", "string ")
		}

		if strings.Contains(l, "const char*") {
			l = strings.ReplaceAll(l, "const char*", "string")
		}

		if strings.Contains(l, "const char") {
			l = strings.ReplaceAll(l, "const char", "string")
		}

		if strings.Contains(l, "perl::") {
			l = strings.ReplaceAll(l, "perl::array", "array")
			l = strings.ReplaceAll(l, "perl::hash", "hash")
			l = strings.ReplaceAll(l, "perl::reference", "reference")
			l = strings.ReplaceAll(l, "perl::scalar", "scalar")
		}

		lineSplit := strings.Split(l, " ")
		if len(lineSplit) > 1 && len(lineSplit[1]) > 5 {
			methodIdentifier := lineSplit[1][:5]
			// Need to make sure it's a proper Method and not a "_new()", package.add, etc.
			if methodIdentifier == "Perl_" &&
				!strings.Contains(l, "package.add") &&
				!strings.Contains(l, "return ") &&
				!strings.Contains(l, "#ifdef") &&
				!strings.Contains(l, "register") &&
				!strings.Contains(l, "_new()") {
				// Remove return type from line and convert std::string to string and " const" to ""
				returnType := strings.TrimSpace(lineSplit[0])
				l = strings.ReplaceAll(l, returnType+" ", "")
				l = strings.ReplaceAll(l, "std::string ", "string ")
				l = strings.ReplaceAll(l, " const", "")

				// Make sure that we're properly reference a method here
				if strings.Contains(l, "(") {
					split := strings.Split(l, "(")
					methodSplit := strings.Split(split[0], "_")
					methodType := strings.TrimSpace(methodSplit[1])
					// If the length of method type is 0, that means it's a quest method
					if len(methodType) == 0 {
						methodType = "quest"
					}

					methodName := strings.TrimSpace(strings.Join(methodSplit[2:], "_"))

					// Method name in defintion map is type_name
					methodName = methodType + "_" + methodName

					// Remove quest_ from method names for definition map
					methodName = strings.ReplaceAll(methodName, "quest_", "")

					// check if there is a definition mapping that is different from the
					// method name parsed from the function itself, this is what is
					// actually used in the quest api
					if val, ok := perlDefinitionMap[methodName]; ok {
						methodName = strings.ReplaceAll(methodName, methodName, val)
					}

					paramString := split[1]

					categories := []string{}

					// Split based on method having categories (i.e // @categories Raid)
					if strings.Contains(paramString, "//") {
						commentSplit := strings.Split(paramString, "//")
						comments := ""
						// Set parameter string to first part of split
						paramString = commentSplit[0]
						// If we have categories handle the splitting and clean up here
						if len(commentSplit) > 1 {
							comments = strings.TrimSpace(commentSplit[1])

							// parse annotations individually
							// supports other annotations easily if more are added
							annotationSplit := strings.Split(comments, "@")
							if len(annotationSplit) > 1 {
								for _, s := range annotationSplit {

									// @categories [cat1, cat2, cat3]
									categoriesAnnotation := strings.Split(s, "categories")
									if len(categoriesAnnotation) > 1 {

										// multiple categories
										categorySplit := strings.Split(categoriesAnnotation[1], ",")
										if len(categorySplit) > 0 {

											// trim spaces
											for i := range categorySplit {
												categorySplit[i] = strings.TrimSpace(categorySplit[i])
											}

											categories = categorySplit
										}
									}
								}
							}
						}
					}

					// Check for self variables (i.e Mob* Perl_EntityList_GetMobID(EntityList* self, uint16_t mob_id)
					if strings.Contains(paramString, "self") {
						var selfSplit []string

						// Split based on if method only has self parameter or if it has self and other parameters as well
						if strings.Contains(paramString, "self,") {
							selfSplit = strings.Split(paramString, "self, ")
						} else {
							selfSplit = strings.Split(paramString, "self")
						}

						// If not parameterless, parameter string is second part of self split
						// If paramaterless, parameter string is empty
						if len(selfSplit) > 1 {
							paramString = selfSplit[1]
						} else {
							paramString = ""
						}
					}

					params := strings.ReplaceAll(paramString, ")", "")

					// split and zero if empty
					p := strings.Split(strings.TrimSpace(params), ",")
					for i := range p {
						p[i] = strings.TrimSpace(p[i])
					}

					if params == "" {
						p = []string{}
					}

					methodFull := methodName
					if len(methodType) > 0 {
						methodFull = methodType + "_" + methodFull
					}

					// Only add to functions if it exists in definition map or is a quest method
					if len(perlDefinitionMap[methodFull]) > 0 || methodType == "quest" {
						perlMethods[methodType] = append(
							perlMethods[methodType], PerlMethod{
								Method:     methodName,
								Params:     p,
								MethodType: methodType,
								ReturnType: returnType,
								Categories: categories,
							},
						)
					}
				}
			}
		}
	}
}
