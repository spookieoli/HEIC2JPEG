package ConfigurationManager

import (
	"encoding/json"
	"os"
)

type ConfigurationManager struct {
	config Configuration
}

type Configuration struct {
	PosX      int    `json:"pos_x"`
	PosY      int    `json:"pos_y"`
	SourceDir string `json:"source_dir"`
	TargetDir string `json:"target_dir"`
}

// Create a new ConfigurationManager
func New() *ConfigurationManager {
	// Vars
	cm := ConfigurationManager{}
	var configuration Configuration
	// First Check if the Configuration File exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		// If the File does not exist, create it
		configuration = cm.createNewConfiguration()
	} else {
		// If the File exists, read it
		configuration = cm.ReadConfiguration()
	}
	cm.config = configuration
	return &ConfigurationManager{config: configuration}
}

func (c *ConfigurationManager) ReadConfiguration() Configuration {
	// The configuration
	configuration := Configuration{}
	// Read the config
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Decode the Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

func (c *ConfigurationManager) createNewConfiguration() Configuration {
	// Vars
	configuration := Configuration{}
	// If the File does not exist, create it
	file, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new Configuration
	configuration.PosX = 0
	configuration.PosY = 0
	configuration.SourceDir = ""
	configuration.TargetDir = ""

	// Write the Configuration to the File
	encoder := json.NewEncoder(file)
	err = encoder.Encode(configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
