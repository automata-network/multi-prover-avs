package operator

import (
	"io"
	"net/http"
	"net/url"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/xmetric"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	eigenmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
)

type Metrics struct {
	EigenMetrics eigenmetrics.Metrics

	registry *prometheus.Registry

	economicCollector *economic.Collector
	rpccallsCollector *rpccalls.Collector
}

func NewMetrics(appName string, clients *clients.Clients, logger logging.Logger, operatorAddr common.Address, socketAddr string, quorumNames map[types.QuorumNum]string) *Metrics {
	reg := clients.PrometheusRegistry
	economicCollector := economic.NewCollector(
		clients.ElChainReader,
		clients.AvsRegistryChainReader,
		appName,
		logger,
		operatorAddr,
		quorumNames,
	)
	rpccallsCollector := rpccalls.NewCollector(appName, reg)
	reg.MustRegister(economicCollector)
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	reg.MustRegister(collectors.NewGoCollector())

	return &Metrics{
		economicCollector: economicCollector,
		rpccallsCollector: rpccallsCollector,
		EigenMetrics:      clients.Metrics,
		registry:          reg,
	}
}

func (g *Metrics) Gather() ([]*xmetric.MetricFamily, error) {
	return g.registry.Gather()
}

func (c *Metrics) Serve(addr string, proverAddr string) chan error {
	opt := promhttp.HandlerOpts{
		EnableOpenMetrics: true,
		Registry:          c.registry,
	}
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(c, opt))
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	mux.Handle("/proverMetrics", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		metricUrl, err := url.JoinPath(proverAddr, "/metrics")
		if err != nil {
			panic(err)
		}
		response, err := http.Get(metricUrl)
		if err != nil {
			w.WriteHeader(502)
			return
		}
		io.Copy(w, response.Body)
		response.Body.Close()
	}))

	logex.Infof("Prometheus listen on %v", srv.Addr)
	errChan := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errChan <- logex.Trace(err)
		}
	}()
	return errChan
}
