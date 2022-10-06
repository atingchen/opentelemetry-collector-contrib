package obfuscateprocessor

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type obfuscate struct {
}

// processTraces implements ProcessMetricsFunc. It processes the incoming data
// and returns the data to be sent to the next component
func (s *obfuscate) processTraces(ctx context.Context, batch ptrace.Traces) (ptrace.Traces, error) {
	for i := 0; i < batch.ResourceSpans().Len(); i++ {
		rs := batch.ResourceSpans().At(i)
		_ = rs
	}
	return batch, nil
}

// Capabilities specifies what this processor does, such as whether it mutates data
func (s *obfuscate) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

// Start the redaction processor
func (s *obfuscate) Start(_ context.Context, _ component.Host) error {
	return nil
}

// Shutdown the redaction processor
func (s *obfuscate) Shutdown(context.Context) error {
	return nil
}
