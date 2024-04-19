package operator

import (
	"context"
	"math/big"
	"time"

	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogHandler interface {
	GetBlock() (uint64, error)
	SaveBlock(uint64) error
	OnNewLog(ctx context.Context, log *types.Log) error
}

type LogTracer struct {
	id               string
	wait             uint64
	max              uint64
	source           *ethclient.Client
	filter           ethereum.FilterQuery
	handler          LogHandler
	skipOnError      bool
	scanIntervalSecs int64
}

type LogTracerConfig struct {
	Id               string
	Wait             uint64
	Max              uint64
	Topics           [][]common.Hash
	Addresses        []common.Address
	Handler          LogHandler
	ScanIntervalSecs int64
	SkipOnError      bool
}

func NewLogTracer(source *ethclient.Client, cfg *LogTracerConfig) *LogTracer {
	if cfg.ScanIntervalSecs == 0 {
		cfg.ScanIntervalSecs = 4
	}
	return &LogTracer{
		id:          cfg.Id,
		skipOnError: cfg.SkipOnError,
		source:      source,
		wait:        cfg.Wait,
		max:         cfg.Max,
		filter: ethereum.FilterQuery{
			Addresses: cfg.Addresses,
			Topics:    cfg.Topics,
		},
		scanIntervalSecs: cfg.ScanIntervalSecs,
		handler:          cfg.Handler,
	}
}

func (l *LogTracer) sleepSecs(ctx context.Context, n int64) {
	timer := time.NewTimer(time.Duration(n) * time.Second)
	defer timer.Stop()
	select {
	case <-timer.C:
	case <-ctx.Done():
	}
}

func (l *LogTracer) saveOffset(off uint64) error {
	l.handler.SaveBlock(off)
	return nil
}

func (l *LogTracer) Run(ctx context.Context) error {
	logex.Info("starting log-tracer:", l.id)
	start, err := l.handler.GetBlock()
	if err != nil {
		return logex.Trace(err)
	}
	head, err := l.source.BlockNumber(ctx)
	if err != nil {
		return logex.Trace(err)
	}
	head -= l.wait

	if start > head || start == 0 {
		logex.Infof("incorrect start offset=%v, head=%v, reset to head", start, head)
		start = head
	}

scan:
	for {
		select {
		case <-ctx.Done():
			return logex.Trace(ctx.Err())
		default:
			head, err = l.source.BlockNumber(ctx)
			if err != nil {
				logex.Errorf("fetch head fail: %v, retry in 1 secs...", err)
				l.sleepSecs(ctx, 1)
				continue
			}
			head -= l.wait

			if start >= head {
				l.sleepSecs(ctx, l.scanIntervalSecs)
				continue
			}

			end := head
			if end-start > l.max {
				end = start + l.max
			}

			filter := l.filter
			filter.FromBlock = big.NewInt(int64(start))
			filter.ToBlock = big.NewInt(int64(end))

			logs, err := l.source.FilterLogs(ctx, filter)
			if err != nil {
				logex.Errorf("[%v-%v] fetch logs fail: %v => %v", start, end, filter, err)
				l.sleepSecs(ctx, 4)
				continue
			}
			for _, log := range logs {
				if err := l.handler.OnNewLog(ctx, &log); err != nil {
					logex.Errorf("[%v] process logs fail => %v", log.BlockNumber, err)
					if log.BlockNumber-1 > start {
						start = log.BlockNumber
						l.saveOffset(start)
					}
					l.sleepSecs(ctx, 4)
					if !l.skipOnError {
						continue scan
					}
				}
			}
			logex.Infof("[%v] scan %v -> %v, logs: %v", l.id, start, end, len(logs))
			start = end + 1
			l.saveOffset(start)
		}
	}
}
