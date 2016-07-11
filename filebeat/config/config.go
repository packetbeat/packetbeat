package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/paths"
)

// Defaults for config variables which are not set
const (
	DefaultInputType = "log"
)

var (
	DefaultConfig = Config{
		FilebeatConfig{
			RegistryFile: "registry",
			SpoolSize:    2048,
			IdleTimeout:  5 * time.Second,
		},
	}
)

type Config struct {
	Filebeat FilebeatConfig
}

type FilebeatConfig struct {
	Inputs       []*common.Config `config:"inputs"`
	Prospectors  []*common.Config `config:"prospectors"`
	SpoolSize    uint64           `config:"spool_size" validate:"min=1"`
	PublishAsync bool             `config:"publish_async"`
	IdleTimeout  time.Duration    `config:"idle_timeout" validate:"nonzero,min=0s"`
	RegistryFile string           `config:"registry_file"`
	ConfigDir    string           `config:"config_dir"`
}

const (
	LogInputType   = "log"
	StdinInputType = "stdin"
)

// List of valid input types
var ValidInputType = map[string]struct{}{
	StdinInputType: {},
	LogInputType:   {},
}

// getConfigFiles returns list of config files.
// In case path is a file, it will be directly returned.
// In case it is a directory, it will fetch all .yml files inside this directory
func getConfigFiles(path string) (configFiles []string, err error) {

	// Check if path is valid file or dir
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Create empty slice for config file list
	configFiles = make([]string, 0)

	if stat.IsDir() {
		files, err := filepath.Glob(path + "/*.yml")

		if err != nil {
			return nil, err
		}

		configFiles = append(configFiles, files...)

	} else {
		// Only 1 config file
		configFiles = append(configFiles, path)
	}

	return configFiles, nil
}

// mergeConfigFiles reads in all config files given by list configFiles and merges them into config
func mergeConfigFiles(configFiles []string, config *Config) error {

	for _, file := range configFiles {
		logp.Info("Additional configs loaded from: %s", file)

		tmpConfig := &Config{}
		cfgfile.Read(tmpConfig, file)

		config.Filebeat.Inputs = append(config.Filebeat.Inputs, tmpConfig.Filebeat.Inputs...)

		// DEPRECATED: Should be removed in 6.0
		config.Filebeat.Prospectors = append(config.Filebeat.Prospectors, tmpConfig.Filebeat.Prospectors...)
	}

	return nil
}

// Fetches and merges all config files given by configDir. All are put into one config object
func (config *Config) FetchConfigs() {

	configDir := config.Filebeat.ConfigDir

	// If option not set, do nothing
	if configDir == "" {
		return
	}

	// If configDir is relative, consider it relative to the config path
	configDir = paths.Resolve(paths.Config, configDir)

	// Check if optional configDir is set to fetch additional config files
	logp.Info("Additional config files are fetched from: %s", configDir)

	configFiles, err := getConfigFiles(configDir)

	if err != nil {
		log.Fatal("Could not use config_dir of: ", configDir, err)
	}

	err = mergeConfigFiles(configFiles, config)
	if err != nil {
		log.Fatal("Error merging config files: ", err)
	}

	if len(config.Filebeat.Inputs) == 0 {
		log.Fatalf("No paths given. What files do you want me to watch?")
	}
}
