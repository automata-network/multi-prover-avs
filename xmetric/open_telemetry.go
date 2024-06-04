package xmetric

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/chzyer/logex"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

type Config struct {
}

type OpenTelemetryConfig struct {
	Endpoint   string
	UserName   string
	Password   string
	SubmitTime int
}

func (cfg *OpenTelemetryConfig) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := cfg.UserName + ":" + cfg.Password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"Authorization": "Basic " + enc,
	}, nil
}

func (cfg *OpenTelemetryConfig) RequireTransportSecurity() bool {
	return true
}

type OpenTelemetryConverter struct {
	prevSubmit time.Time
}

func (o *OpenTelemetryConverter) LabelAttributes(labels []*io_prometheus_client.LabelPair) attribute.Set {
	kvs := make([]attribute.KeyValue, len(labels))
	for idx, label := range labels {
		kvs[idx] = attribute.Key(*label.Name).String(*label.Value)
	}

	return attribute.NewSet(kvs...)
}

func (o *OpenTelemetryConverter) ToGaugeAggregation(counters []*io_prometheus_client.Metric) metricdata.Aggregation {
	dps := make([]metricdata.DataPoint[float64], len(counters))
	for idx, metric := range counters {
		gauge := metric.Gauge
		dps[idx] = metricdata.DataPoint[float64]{
			Attributes: o.LabelAttributes(metric.Label),
			Value:      *gauge.Value,
			Exemplars:  nil,
		}
	}
	return metricdata.Gauge[float64]{
		DataPoints: dps,
	}
}

func (o *OpenTelemetryConverter) Exemplars(exemplar *io_prometheus_client.Exemplar) []metricdata.Exemplar[float64] {
	return nil
}

func (o *OpenTelemetryConverter) ToCounterAggregation(counters []*io_prometheus_client.Metric) metricdata.Aggregation {
	dps := make([]metricdata.DataPoint[float64], len(counters))
	for idx, metric := range counters {
		counter := metric.Counter
		dps[idx] = metricdata.DataPoint[float64]{
			Attributes: o.LabelAttributes(metric.Label),
			StartTime:  counter.CreatedTimestamp.AsTime(),
			Value:      *counter.Value,
			Exemplars:  o.Exemplars(counter.Exemplar), // TODO
		}
	}
	return metricdata.Sum[float64]{
		DataPoints:  dps,
		Temporality: metricdata.CumulativeTemporality,
		IsMonotonic: true,
	}
}

func (o *OpenTelemetryConverter) ToAggregation(ty io_prometheus_client.MetricType, metrics []*io_prometheus_client.Metric) metricdata.Aggregation {
	switch ty {
	case io_prometheus_client.MetricType_COUNTER:
		return o.ToCounterAggregation(metrics)
	case io_prometheus_client.MetricType_GAUGE:
		return o.ToGaugeAggregation(metrics)
	default:
		panic(fmt.Sprintf("unsupport type: %v", ty.String()))
	}
}

func (o *OpenTelemetryConverter) ToMetric(item *MetricFamily) metricdata.Metrics {
	return metricdata.Metrics{
		Name:        *item.Name,
		Description: *item.Help,
		Unit:        "",
		Data:        o.ToAggregation(*item.Type, item.Metric),
	}
}

func (o *OpenTelemetryConverter) ToResourceMetrics(res *resource.Resource, prometheusMetrics []*MetricFamily) *metricdata.ResourceMetrics {
	return &metricdata.ResourceMetrics{
		Resource: res,
		ScopeMetrics: []metricdata.ScopeMetrics{
			{
				Scope:   instrumentation.Scope{Name: "meter", Version: "", SchemaURL: ""},
				Metrics: o.ToMetrics(prometheusMetrics),
			},
		},
	}
}

func (o *OpenTelemetryConverter) ToMetrics(m []*MetricFamily) []metricdata.Metrics {
	out := make([]metricdata.Metrics, len(m))
	for idx, item := range m {
		out[idx] = o.ToMetric(item)
	}
	return out
}

func ExportMetricToOpenTelemetry(cfg *OpenTelemetryConfig, registry prometheus.Gatherer) error {
	if cfg.SubmitTime <= 0 {
		cfg.SubmitTime = 1
	}
	conn, err := grpc.NewClient(cfg.Endpoint,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(cfg),
	)
	if err != nil {
		return logex.NewErrorf("failed to create gRPC connection to collector: %w", err)
	}
	defer conn.Close()

	ctx := context.Background()

	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName("avs"),
	))
	if err != nil {
		return logex.NewErrorf("failed to create resource: %w", err)
	}

	metricExporter, err := otlpmetricgrpc.New(context.Background(), otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return logex.NewErrorf("failed to create metric exporter: %w", err)
	}
	defer metricExporter.ForceFlush(ctx)

	prevSubmit := time.Now()

	export := func() {
		prometheusMetrics, err := registry.Gather()
		if err != nil {
			logex.Error(err)
			return
		}

		converter := OpenTelemetryConverter{
			prevSubmit,
		}
		resourceMetrics := converter.ToResourceMetrics(res, prometheusMetrics)
		logex.Infof("submit %v metrics to opentelemetry", len(prometheusMetrics))
		for i := 0; i < cfg.SubmitTime; i++ {
			if err := metricExporter.Export(ctx, resourceMetrics); err != nil {
				logex.Error(err)
				return
			}
		}

		prevSubmit = time.Now()
	}
	ticker := time.NewTicker(300 * time.Second)
	defer ticker.Stop()

	export()
	for range ticker.C {
		export()
	}

	return nil
}
