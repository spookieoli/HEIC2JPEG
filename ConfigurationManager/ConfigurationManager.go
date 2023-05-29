package ConfigurationManager

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigurationManager struct {
	Config Configuration
}

type Configuration struct {
	PosX      int    `json:"pos_x"`
	PosY      int    `json:"pos_y"`
	SourceDir string `json:"source_dir"`
	TargetDir string `json:"target_dir"`
	Worker    int    `json:"worker"`
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
	cm.Config = configuration
	return &ConfigurationManager{Config: configuration}
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
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Cannot close File")
		}
	}()

	// Create a new Configuration
	configuration.PosX = 0
	configuration.PosY = 0
	configuration.SourceDir = ""
	configuration.TargetDir = ""
	configuration.Worker = 1

	// Write the Configuration to the File
	encoder := json.NewEncoder(file)
	err = encoder.Encode(configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

// WriteConfiguration writes the Configuration to the config.json File
func (c *ConfigurationManager) WriteConfiguration() {
	// Open the File
	file, err := os.OpenFile("config.json", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Write the Configuration to the File
	encoder := json.NewEncoder(file)
	err = encoder.Encode(c.Config)
	if err != nil {
		panic(err)
	}
}
