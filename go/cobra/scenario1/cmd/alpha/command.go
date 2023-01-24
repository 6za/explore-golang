package alpha

import (
	"github.com/spf13/cobra"
)

var (
	someFlag          string
	someOtherFlagFlag bool
)

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

func ChildCommand() *cobra.Command {
	childActionCmd := &cobra.Command{
		Use:   "sub-action",
		Short: "something short",
		Long:  "something long",
		RunE:  executeAlphaChild,
	}

	return childActionCmd
}
