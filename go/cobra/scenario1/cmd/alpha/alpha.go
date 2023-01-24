package alpha

import (
	"github.com/6za/explore-golang/go/cobra/scenario1/internal/alphaController"
	"github.com/6za/explore-golang/go/cobra/scenario1/pkg/loggerUtil"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func executeAlpha(cmd *cobra.Command, args []string) error {

	//Note A: A command should decide to log or not
	log.Logger = loggerUtil.ZerologSetup(zerolog.InfoLevel)

	log.Info().Msgf("Alpha CMD executed: (someFlag:  %s) ", someFlag)

	//Note B: Internal actions should be part of implementation of `DoAlphaAction`
	// note in the body of the command, to allow more standarized command bodies
	// to help isolate the UI of th command from the implementation of its tasks.

	return alphaController.DoAlphaAction(someFlag)
	//return alphaController.DoAlphaActionAlternative(someFlag)

	//Note C:
	//This allows to  simple test a new behavior without big refactory

}
