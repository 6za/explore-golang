package alphaController

import (
	"testing"
)

func Test_controller(t *testing.T) {

	err := DoAlphaAction("valueA")
	if err != nil {
		t.Error(err)
	}

	//Check some condition
}
