package obfuscateprocessor

import (
	"context"
	"github.com/cyrildever/feistel"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type obfuscate struct {
	// Logger
	logger *zap.Logger
	// Next trace consumer in line
	next consumer.Traces

	encryptAttributes map[string]struct{}
	encrypt           *feistel.FPECipher
}

// processTraces implements ProcessMetricsFunc. It processes the incoming data
// and returns the data to be sent to the next component
func (o *obfuscate) processTraces(ctx context.Context, batch ptrace.Traces) (ptrace.Traces, error) {
	for i := 0; i < batch.ResourceSpans().Len(); i++ {
		rs := batch.ResourceSpans().At(i)
		o.processResourceSpan(ctx, rs)
	}
	return batch, nil
}

// processResourceSpan processes the ResourceSpans and all of its spans
func (o *obfuscate) processResourceSpan(ctx context.Context, rs ptrace.ResourceSpans) {
	rsAttrs := rs.Resource().Attributes()

	// Attributes can be part of a resource span
	o.processAttrs(ctx, rsAttrs)

	for j := 0; j < rs.ScopeSpans().Len(); j++ {
		ils := rs.ScopeSpans().At(j)
		for k := 0; k < ils.Spans().Len(); k++ {
			span := ils.Spans().At(k)
			spanAttrs := span.Attributes()

			// Attributes can also be part of span
			o.processAttrs(ctx, spanAttrs)
		}
	}
}

// processAttrs obfuscate the attributes of a resource span or a span
func (o *obfuscate) processAttrs(_ context.Context, attributes pcommon.Map) {
	attributes.Range(func(k string, value pcommon.Value) bool {
		_, ok := o.encryptAttributes[k]
		if !ok {
			return true
		}
		switch value.Type() {
		case pcommon.ValueTypeInt:
			encryptValue := o.encryptNumber(value.Int())
			value.SetInt(encryptValue)
		default:
			encryptValue := o.encryptString(value.Str())
			value.SetStr(encryptValue)
		}
		return true
	})

}

// Capabilities specifies what this processor does, such as whether it mutates data
func (o *obfuscate) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

// Start the redaction processor
func (o *obfuscate) Start(_ context.Context, _ component.Host) error {
	return nil
}

// Shutdown the redaction processor
func (o *obfuscate) Shutdown(context.Context) error {
	return nil
}

func (o *obfuscate) encryptNumber(source int64) int64 {
	obfuscated, _ := o.encrypt.EncryptNumber(uint64(source))
	return int64(obfuscated.Uint64())
}

func (o *obfuscate) encryptString(source string) string {
	obfuscated, _ := o.encrypt.Encrypt(source)
	return obfuscated.String(true)
}
