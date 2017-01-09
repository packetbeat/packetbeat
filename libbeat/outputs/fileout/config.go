package fileout

import (
	"fmt"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
)

type config struct {
	Path          string               `config:"path"`
	Filename      string               `config:"filename"`
	RotateEveryKb int                  `config:"rotate_every_kb" validate:"min=1"`
	NumberOfFiles int                  `config:"number_of_files"`
	WriterConfig  outputs.WriterConfig `config:"writer"`
}

var (
	defaultConfig = config{
		NumberOfFiles: 7,
		RotateEveryKb: 10 * 1024,
		WriterConfig:  outputs.WriterConfig{Type: outputs.JsonWriterType, Pretty: false},
	}
)

func (c *config) Validate() error {
	if c.NumberOfFiles < 2 || c.NumberOfFiles > logp.RotatorMaxFiles {
		return fmt.Errorf("The number_of_files to keep should be between 2 and %v",
			logp.RotatorMaxFiles)
	}

	return nil
}
