package main

import (
	"io/ioutil"
	"os"

	toml "github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog/log"
)

// loadConfiguration takes a straight TOML file and loads it into a Configuration object.
// the goal is to provide a tool to load a configuration file and then pass into environment variables
// this assumes a simple structure of TOML for more complex scenarios this would need to be modified
func loadConfiguration(configurationFile string) {

	config := make(map[string]interface{})

	if configurationFile != "" {

		content, err := ioutil.ReadFile(configurationFile)
		if err != nil {
			log.Error().Msgf("the content of configuration file %s were readable", configurationFile)
			os.Exit(13) // Exit with error code 13 as data is invalid (windows exit code)
		}

		toml.Unmarshal(content, &config)

		for key, value := range config {
			os.Setenv(key, value.(string))

		}

	} else {
		log.Error().Msgf("failed to import configuration file %s", configurationFile)
		os.Exit(2) // Exit with error code 2 for file not found (windows exit code)
	}
}
