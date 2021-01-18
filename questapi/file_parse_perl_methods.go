package questapi

import (
	"fmt"
	"strings"
)

// parses perl methods
func parsePerlMethods(contents string, perlMethods map[string][]PerlMethod) {
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

			// params
			paramSplit := strings.Split(methodFull, "(")
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
					MethodFull: fmt.Sprintf("%v(%v)", method, params),
					Params:     p,
					MethodType: methodType,
					ReturnType: returnType,
					Categories: nil,
				},
			)
		}
	}
}
