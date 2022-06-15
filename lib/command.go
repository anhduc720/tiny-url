package lib

import "github.com/spf13/cobra"

type CommandRunner interface{}

// Command interface is used to implement sub-commands in the system.
type Command interface {
	Short() string

	Setup(cmd *cobra.Command)

	Run() CommandRunner
}
