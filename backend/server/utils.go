package server

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func RequiredConfig(paths []string, log zerolog.Logger) {
	fail := false
	for _, path := range paths {
		if !viper.InConfig(path) {
			fail = true
			log.Error().Msgf("Missing configuration key: ", path)
		}
	}
	if fail {
		log.Fatal().Msg("Exiting due to missing configuration keys")
	}
}
