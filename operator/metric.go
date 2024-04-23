package operator

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	eigenmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
)

type Metrics struct {
	EigenMetrics eigenmetrics.Metrics

	registry *prometheus.Registry

	economicCollector *economic.Collector
	rpccallsCollector *rpccalls.Collector
}

func NewMetrics(clients *clients.Clients, logger logging.Logger, operatorAddr common.Address, socketAddr string, quorumNames map[types.QuorumNum]string) *Metrics {
	reg := clients.PrometheusRegistry
	economicCollector := economic.NewCollector(
		clients.ElChainReader,
		clients.AvsRegistryChainReader,
		AppName,
		logger,
		operatorAddr,
		quorumNames,
	)
	rpccallsCollector := rpccalls.NewCollector(AppName, reg)
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

func (g *Metrics) Start(ctx context.Context) <-chan error {
	errChan := g.EigenMetrics.Start(ctx, g.registry)
	return errChan
}
