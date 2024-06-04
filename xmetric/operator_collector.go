package xmetric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type OperatorCollector struct {
	FetchTask         *prometheus.CounterVec
	SubmitTask        *prometheus.CounterVec
	LatestTask        *prometheus.GaugeVec
	ProcessTaskMs     *prometheus.GaugeVec
	GenReportMs       *prometheus.GaugeVec
	LivenessTs        *prometheus.GaugeVec
	NextAttestationTs *prometheus.GaugeVec
	Metadata          *prometheus.CounterVec

	AttestationAccBalance *prometheus.GaugeVec
	LastAttestationCost   *prometheus.GaugeVec

	registry *prometheus.Registry
}

func NewOperatorCollector(app string, registry *prometheus.Registry) *OperatorCollector {
	var metrics []prometheus.Collector
	collector := &OperatorCollector{
		Metadata: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "counter_metadata",
				Help:      "metadata",
			},
			[]string{"avs_name", "operator_addr", "version", "attestation_addr", "prover_url_hash", "prover_version", "task_fetch_with_context"},
		)),
		FetchTask: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "counter_fetch_task",
				Help:      "numbers of request for fetching task from the aggregator",
			},
			[]string{"avs_name", "type", "with_context"},
		)),
		SubmitTask: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "counter_submit_task",
				Help:      "numbers of request for submiting task to the aggregator",
			},
			[]string{"avs_name", "type"},
		)),
		LatestTask: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "gauge_latest_task",
				Help:      "trace the latest task id",
			},
			[]string{"avs_name", "type"},
		)),
		ProcessTaskMs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "gauge_process_task_ms",
				Help:      "The time it takes to process the task in milliseconds",
			},
			[]string{"avs_name", "type"},
		)),
		GenReportMs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "gauge_gen_report_ms",
				Help:      "milliseconds for generating attestation report",
			},
			[]string{"avs_name"},
		)),
		LivenessTs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "const_liveness_ts",
				Help:      "timestamp for last attestation submition",
			},
			[]string{"avs_name"},
		)),
		NextAttestationTs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "const_next_attestation_ts",
				Help:      "the time for next attestation",
			},
			[]string{"avs_name"},
		)),
		LastAttestationCost: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "gauge_last_attestation_cost",
				Help:      "the cost of latest attestation",
			},
			[]string{"avs_name"},
		)),
		AttestationAccBalance: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "attestation_acc_balance",
				Help:      "The time it takes to process the task in milliseconds",
			},
			[]string{"avs_name"},
		)),
	}
	collector.registry = prometheus.NewRegistry()
	collector.registry.MustRegister(metrics...)
	if registry != nil {
		registry.MustRegister(metrics...)
	}
	return collector
}

func (c *OperatorCollector) Gather() ([]*MetricFamily, error) {
	return c.registry.Gather()
}
