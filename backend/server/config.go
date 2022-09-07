package server

import (
	"github.com/go-chi/httplog"
	"github.com/spf13/viper"
)

func Config() {
	log := httplog.NewLogger("leadster", httplog.Options{})

	defaults()
	viper.SetConfigName("config.secret")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Info().Msg("Secret config file not present.")
	}
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal().Err(err).Msg("Configuration file not found")
		} else {
			log.Fatal().Err(err).Msg("Could not load configuration file")
		}
	}
	viper.AutomaticEnv()
}

func defaults() {
	viper.SetDefault("ContentDir", "content")
}
