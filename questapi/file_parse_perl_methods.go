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
			l = strings.ReplaceAll(l, "Perl_croak(aTHX_ \"Usage:", "")
			l = strings.ReplaceAll(l, "\");", "")
			l = strings.ReplaceAll(l, "THIS, ", "")
			l = strings.ReplaceAll(l, "THIS,", "")
			l = strings.ReplaceAll(l, "THIS", "")

			split := strings.Split(l, "::")
			methodType := strings.TrimSpace(split[0])
			method := strings.TrimSpace(split[1])
			returnType := ""

			perlMethods[methodType] = append(
				perlMethods[methodType], PerlMethod{
					Method:     method,
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
