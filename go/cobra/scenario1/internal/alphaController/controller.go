package alphaController

import "github.com/rs/zerolog/log"

func DoAlphaAction(flagOne string) error {

	log.Info().Msgf("Execute alpha steps: %s", flagOne)
	return nil
}

func DoAlphaActionAlternative(flagOne string) error {

	log.Info().Msgf("Execute alternative alpha steps: %s", flagOne)
	return nil
}
