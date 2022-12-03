package obfuscateprocessor

import (
	"context"
	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/hash"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "obfuscation"
	// The stability level of the exporter.
	stability = component.StabilityLevelAlpha

	defaultRound = 10
)

// NewFactory creates a factory for the redaction processor.
func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesProcessor(createTracesProcessor, stability),
	)
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
		EncryptRound:      defaultRound,
	}
}

// createTracesProcessor creates an instance of redaction for processing traces
func createTracesProcessor(
	ctx context.Context,
	set component.ProcessorCreateSettings,
	cfg config.Processor,
	next consumer.Traces,
) (component.TracesProcessor, error) {
	oCfg := cfg.(*Config)
	processor := &obfuscate{
		logger:            set.Logger,
		next:              next,
		encrypt:           feistel.NewFPECipher(hash.SHA_256, oCfg.EncryptKey, oCfg.EncryptRound),
		encryptAttributes: makeEncryptList(oCfg),
	}
	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		next,
		processor.processTraces,
		processorhelper.WithCapabilities(processor.Capabilities()),
		processorhelper.WithStart(processor.Start),
		processorhelper.WithShutdown(processor.Shutdown))
}

// makeEncryptList sets up a lookup table of  span attribute keys which need to be encrypted.
func makeEncryptList(c *Config) map[string]struct{} {
	allowList := make(map[string]struct{}, len(c.EncryptAttributes))
	for _, key := range c.EncryptAttributes {
		allowList[key] = struct{}{}
	}
	return allowList
}
