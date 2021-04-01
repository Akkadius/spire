package questapi

import (
	"strings"
)

// parses perl methods
func (c *ParseService) parsePerlMethods(contents string, perlMethods map[string][]PerlMethod) {
	for _, l := range strings.Split(contents, "\n") {
		filterLine := strings.Contains(l, "Perl_croak") &&
			strings.Contains(l, "Usa") &&
			strings.Contains(l, "::") &&
			!strings.Contains(l, "::new")

		if filterLine {
			// cut
			l = strings.ReplaceAll(l, "Perl_croak(aTHX_ \"Usage:", "")
			l = strings.ReplaceAll(l, "\");", "")
			l = strings.ReplaceAll(l, "THIS, ", "")
			l = strings.ReplaceAll(l, "THIS,", "")
			l = strings.ReplaceAll(l, "THIS", "")

			// split method
			split := strings.Split(l, "::")
			methodType := strings.TrimSpace(split[0])
			methodFull := strings.TrimSpace(split[1])
			returnType := ""

			// has the entire line for the method parse at first
			methodLine := methodFull

			// comments
			commentSplit := strings.Split(methodFull, "//")
			comments := ""
			categories := []string{}
			if len(commentSplit) > 1 {
				comments = strings.TrimSpace(commentSplit[1])

				// entire line without the comments
				methodLine = strings.TrimSpace(commentSplit[0])

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

			// params
			paramSplit := strings.Split(methodLine, "(")
			params := strings.ReplaceAll(paramSplit[1], ")", "")
			p := strings.Split(strings.TrimSpace(params), ",")
			for i := range p {
				p[i] = strings.TrimSpace(p[i])
			}
			if params == "" {
				p = []string{}
			}
			method := paramSplit[0]

			// append to method list
			perlMethods[methodType] = append(
				perlMethods[methodType], PerlMethod{
					Method:     method,
					Params:     p,
					MethodType: methodType,
					ReturnType: returnType,
					Categories: categories,
				},
			)
		}
	}
}
