package xmetric

import (
	"net/http"
	"sync"

	"github.com/chzyer/logex"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	dto "github.com/prometheus/client_model/go"
)

type MetricFamily = dto.MetricFamily

var LabelOperator = "operator"
var LabelAggregator = "aggregator"

type AggregatorCollector struct {
	NewTask   *prometheus.CounterVec
	FetchTask *prometheus.CounterVec

	operatorRegistry *OperatorRegistry
	registry         *prometheus.Registry
}

func collect[T prometheus.Collector](list *[]prometheus.Collector, m T) T {
	*list = append(*list, m)
	return m
}

func NewAggregatorCollector(app string) *AggregatorCollector {
	var metrics []prometheus.Collector
	collector := &AggregatorCollector{
		NewTask: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelAggregator,
				Name:      "new_task",
				Help:      "new task counter",
			},
			[]string{"type"},
		)),
		FetchTask: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelAggregator,
				Name:      "fetch_task",
				Help:      "fetch task counter",
			},
			[]string{"type", "with_context"},
		)),
	}
	collector.registry = prometheus.NewRegistry()
	collector.registry.MustRegister(metrics...)
	collector.operatorRegistry = NewOperatorRegistry()
	return collector
}

func (c *AggregatorCollector) AddOperatorMetrics(operator string, metrics []*MetricFamily) {
	c.operatorRegistry.AddMetric(operator, metrics)
}

func (c *AggregatorCollector) Serve(addr string) error {
	opt := promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	}
	mux := http.NewServeMux()
	mux.Handle("/", promhttp.HandlerFor(c, opt))
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	logex.Infof("Prometheus listen on %v", srv.Addr)
	return srv.ListenAndServe()
}

func (c *AggregatorCollector) Gather() ([]*dto.MetricFamily, error) {
	operators, err := c.operatorRegistry.Gather()
	if err != nil {
		return nil, logex.Trace(err)
	}
	registry, err := c.registry.Gather()
	if err != nil {
		return nil, logex.Trace(err)
	}
	return append(operators, registry...), nil
}

type OperatorRegistry struct {
	mutex  sync.Mutex
	family map[string][]*dto.MetricFamily
}

func NewOperatorRegistry() *OperatorRegistry {
	return &OperatorRegistry{
		family: make(map[string][]*dto.MetricFamily),
	}
}

func (c *OperatorRegistry) AddMetric(operator string, metrics []*dto.MetricFamily) {
	for _, item := range metrics {
		for _, metric := range item.Metric {
			metric.Label = append(metric.Label, &dto.LabelPair{
				Name:  &LabelOperator,
				Value: &operator,
			})
		}
	}

	c.mutex.Lock()
	c.family[operator] = metrics
	c.mutex.Unlock()
}

func (c *OperatorRegistry) Gather() ([]*dto.MetricFamily, error) {
	var out []*dto.MetricFamily
	for _, item := range c.family {
		out = append(out, item...)
	}
	return out, nil
}
