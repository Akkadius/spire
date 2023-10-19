package questapi

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

type ParseCommand struct {
	logger  *logrus.Logger
	command *cobra.Command
	parser  *ParseService
}

func (c *ParseCommand) Command() *cobra.Command {
	return c.command
}

func NewParseCommand(
	logger *logrus.Logger,
	parser *ParseService,
) *ParseCommand {
	i := &ParseCommand{
		parser: parser,
		logger: logger,
		command: &cobra.Command{
			Use:   "quest:parse",
			Short: "Parses EQEmu/Server Quest API source files for documentation",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *ParseCommand) Handle(_ *cobra.Command, args []string) {
	start := time.Now()

	_ = c.parser.Parse(false)

	//fmt.Println(response)

	//
	//fmt.Printf("# Perl\n")
	//for methodType, methods := range methods.PerlApi.PerlMethods {
	//	for _, method := range methods {
	//		returnType := ""
	//		if method.ReturnType != "" {
	//			returnType = fmt.Sprintf(" # %v", method.ReturnType)
	//		}
	//
	//		fmt.Printf("%v::%v;%v\n", methodType, method.Method, returnType)
	//	}
	//}
	//
	//fmt.Printf("\n\n")
	//
	//fmt.Printf("# Lua\n")
	//for methodType, methods := range methods.LuaApi.LuaMethods {
	//	for _, method := range methods {
	//		returnType := ""
	//		if method.ReturnType != "" {
	//			returnType = fmt.Sprintf(" -- %v", method.ReturnType)
	//		}
	//
	//		fmt.Printf("%v::%v%v\n", methodType, method.Method, returnType)
	//	}
	//}

	fmt.Printf("took %v\n", time.Since(start))
}

// Validate
func (c *ParseCommand) Validate(_ *cobra.Command, args []string) error {
	// Validate

	return nil
}
