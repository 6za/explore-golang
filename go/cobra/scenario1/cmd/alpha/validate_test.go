package alpha

import (
	"testing"

	"github.com/spf13/cobra"
)

func Test_validate_bad_value(t *testing.T) {

	someFlag = "bad-value"
	err := validateAlpha(&cobra.Command{}, []string{})

	//It should fail by design, it is a bad value
	if err == nil {
		t.Error(err)
	}

	//Check some condition
}
