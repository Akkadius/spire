package console

import (
	"fmt"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"io"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

var templateFuncs = template.FuncMap{
	"trim":                    strings.TrimSpace,
	"trimRightSpace":          trimRightSpace,
	"trimTrailingWhitespaces": trimRightSpace,
	"appendIfNotPresent":      appendIfNotPresent,
	"rpad":                    rpad,
	"gt":                      cobra.Gt,
	"eq":                      cobra.Eq,
	"formattedCommands":       formattedCommands,
}

const usageMenu = `{{Foreground "#ffff00" "Usage"}}{{if .Runnable}}
 {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
 {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

{{Foreground "#ffff00" "Aliases"}}
 {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{Foreground "#ffff00" "Examples"}}
{{.Example}}{{end}}{{if .HasAvailableLocalFlags}}

{{Foreground "#ffff00" "Options"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableSubCommands}}

{{Foreground "#ffff00" "Available Commands"}}
{{formattedCommands .Commands}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{Foreground "#ffff00" "Additional Help Topics"}}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
 {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// help func override
func helpFunc(c *cobra.Command, a []string) {
	// The help should be sent to stdout
	// See https://github.com/spf13/cobra/issues/1002
	err := tmpl(c.OutOrStdout(), usageMenu, c)
	if err != nil {
		c.Println(err)
	}
}

// template function to return formatted and grouped commands
func formattedCommands(c []*cobra.Command) string {
	usage := ""
	p := termenv.ColorProfile()

	// categorized commands
	categorizedCommands := make(map[string][]*cobra.Command, 0)

	// uncategorized commands
	uncategorizedCommands := make([]*cobra.Command, 0)

	// max command length for padding output
	maxCommandLength := 0
	for _, command := range c {

		// Cobra has a built in padding but it does not calculate the length + terminal colors
		length := len(termenv.String(command.Name()).Foreground(p.Color("#00ff00")).String())
		if length > maxCommandLength {
			maxCommandLength = length
		}

		// Bucket category into categorizedCommands
		if strings.Contains(command.Name(), ":") {
			split := strings.Split(command.Name(), ":")
			category := split[0]
			categorizedCommands[category] = append(categorizedCommands[category], command)
			continue
		}

		// Uncategorized commands
		uncategorizedCommands = append(uncategorizedCommands, command)
	}

	// uncategorized commands
	for _, command := range uncategorizedCommands {
		cmd := termenv.String(command.Name()).Foreground(p.Color("#00ff00")).String()
		usage = fmt.Sprintf("%v%v", usage, fmt.Sprintf(" %v\n", cmd))
	}

	// get the list of sortedCategories and sort them
	sortedCategories := make([]string, 0)
	for key := range categorizedCommands {
		sortedCategories = append(sortedCategories, key)
	}
	sort.Strings(sortedCategories)

	// categorized commands
	for _, category := range sortedCategories {

		// print category
		categoryStr := termenv.String(category).Foreground(p.Color("#ffff00")).String()
		usage = fmt.Sprintf("%v%v", usage, fmt.Sprintf("%v\n", categoryStr))

		commands := categorizedCommands[category]
		for _, command := range commands {
			// print command
			line := termenv.String(command.Name()).Foreground(p.Color("#00ff00")).String()
			usage = fmt.Sprintf(
				"%v%v",
				usage,
				fmt.Sprintf(" %-*s %v\n", maxCommandLength, line, command.Short),
			)
		}
	}

	return strings.TrimSuffix(usage, "\n")
}

func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// FIXME appendIfNotPresent is unused by cobra and should be removed in a version 2. It exists only for compatibility with users of cobra.

// appendIfNotPresent will append stringToAppend to the end of s, but only if it's not yet present in s.
func appendIfNotPresent(s, stringToAppend string) string {
	if strings.Contains(s, stringToAppend) {
		return s
	}
	return s + " " + stringToAppend
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) error {
	t := template.New("top")

	// combine color template funcs with join func map
	for k, v := range termenv.TemplateFuncs(termenv.ColorProfile()) {
		templateFuncs[k] = v
	}

	t.Funcs(templateFuncs)
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}
