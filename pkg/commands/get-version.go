package commands

import (
	"fmt"
	"github.com/restechnica/semverbot/pkg/core"

	"github.com/spf13/cobra"
)

// NewGetVersionCommand creates a new get version command.
// Returns the new spf13/cobra command.
func NewGetVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: GetVersionCommandRun,
	}

	return command
}

// GetVersionCommandRun runs the command.
func GetVersionCommandRun(cmd *cobra.Command, args []string) {
	fmt.Println(core.GetVersion())
}
