package obfuscateprocessor

import (
	"go.opentelemetry.io/collector/config"
)

type Config struct {
	config.ProcessorSettings `mapstructure:",squash"`

	EncryptKey   string `mapstructure:"encrypt_key"`
	EncryptRound int    `mapstructure:"encrypt_round"`

	EncryptAttributes []string `mapstructure:"encrypt_attributes"`
}
