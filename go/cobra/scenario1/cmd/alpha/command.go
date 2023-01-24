package alpha

import (
	"github.com/spf13/cobra"
)

var (
	someFlag          string
	someOtherFlagFlag bool
)

// NewCommand is the parent command of the folder
// The standard explored, this command is called `NewCommand` to help identify its role.
func NewCommand() *cobra.Command {

	myActionCmd := &cobra.Command{
		Use:     "alpha",
		Short:   "something short",
		Long:    "something long",
		PreRunE: validateAlpha,
		RunE:    executeAlpha,
		// PostRunE: runPostAction,
	}

	myActionCmd.Flags().StringVar(&someFlag, "some-flag", "", "....")
	myActionCmd.Flags().BoolVar(&someOtherFlagFlag, "some-other-flag", false, "")

	// wire up new commands
	myActionCmd.AddCommand(ChildCommand())

	return myActionCmd
}

//ChildCommand  is a command that is a child of `NewCommand`
//Its name can be the name of the child action, no standard that I am aware on this naming.
func ChildCommand() *cobra.Command {
	childActionCmd := &cobra.Command{
		Use:   "sub-action",
		Short: "something short",
		Long:  "something long",
		RunE:  executeAlphaChild,
	}

	return childActionCmd
}
