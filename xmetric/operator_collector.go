package xmetric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type OperatorCollector struct {
	FetchTask         *prometheus.CounterVec
	SubmitTask        *prometheus.CounterVec
	LatestTask        *prometheus.GaugeVec
	ProcessTaskMs     *prometheus.GaugeVec
	GenPoeMs          *prometheus.GaugeVec
	SubmitTaskMs      *prometheus.GaugeVec
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
				Name:      "metadata_counter",
				Help:      "metadata",
			},
			[]string{
				"avs_name",
				"operator_addr",
				"version",
				"attestation_addr",
				"prover_url_hash",
				"prover_version",
				"task_fetch_with_context",
				"scroll_with_context",
				"linea_with_context",
			},
		)),
		FetchTask: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "fetch_task_counter",
				Help:      "numbers of request for fetching task from the aggregator",
			},
			[]string{"avs_name", "type", "with_context"},
		)),
		SubmitTask: collect(&metrics, prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "submit_task_counter",
				Help:      "numbers of request for submiting task to the aggregator",
			},
			[]string{"avs_name", "type"},
		)),
		LatestTask: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "latest_task_gauge",
				Help:      "trace the latest task id",
			},
			[]string{"avs_name", "type"},
		)),
		ProcessTaskMs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "process_task_ms_gauge",
				Help:      "The time it takes to process the task in milliseconds",
			},
			[]string{"avs_name", "type"},
		)),
		GenPoeMs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "gen_poe_ms_gauge",
				Help:      "The time it takes to generate poe in milliseconds",
			},
			[]string{"avs_name", "type"},
		)),
		SubmitTaskMs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "submit_task_ms_gauge",
				Help:      "The time it takes to submit task in milliseconds",
			},
			[]string{"avs_name", "type"},
		)),
		GenReportMs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "gen_report_ms_gauge",
				Help:      "milliseconds for generating attestation report",
			},
			[]string{"avs_name"},
		)),
		LivenessTs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "liveness_ts_gauge",
				Help:      "timestamp for last attestation submition",
			},
			[]string{"avs_name"},
		)),
		NextAttestationTs: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "next_attestation_ts_gauge",
				Help:      "the time for next attestation",
			},
			[]string{"avs_name"},
		)),
		LastAttestationCost: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "last_attestation_cost_gauge",
				Help:      "the cost of latest attestation",
			},
			[]string{"avs_name"},
		)),
		AttestationAccBalance: collect(&metrics, prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: app,
				Subsystem: LabelOperator,
				Name:      "attestation_acc_balance_gauge",
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
