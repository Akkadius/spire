package questapi

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"log"
	"strings"
)

// parses perl methods
func parsePerlMethods(fs billy.Filesystem, fileName string, perlMethods map[string][]PerlMethod) {
	file, err := fs.Open(fmt.Sprintf("%v/%v", "./zone", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
