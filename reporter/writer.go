package reporter

import (
	"fmt"
	"io"
	"time"

	"github.com/samuel/go-metrics/metrics"
)

type writerReporter struct {
	w io.Writer
}

func NewWriterReporter(registry *metrics.Registry, interval time.Duration, w io.Writer) *PeriodicReporter {
	return NewPeriodicReporter(registry, interval, false, &writerReporter{w})
}

func (r *writerReporter) Report(registry *metrics.Registry) {
	fmt.Fprintf(r.w, "%+v\n", time.Now())
	registry.Do(func(name string, metric interface{}) error {
		_, err := fmt.Fprintf(r.w, "%s: %+v\n", name, metric)
		return err
	})
}