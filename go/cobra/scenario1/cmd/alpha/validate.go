package alpha

import (
	"fmt"

	"github.com/spf13/cobra"
)

func validateAlpha(cmd *cobra.Command, args []string) error {

	if someFlag == "bad-value" {
		return fmt.Errorf("bad value for flag `some-flag`: %s", someFlag)
	}
	return nil
}
