package obfuscateprocessor

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "obfuscate"
	// The stability level of the exporter.
	stability = component.StabilityLevelAlpha
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
	_ = oCfg
	processor := &obfuscate{}
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
