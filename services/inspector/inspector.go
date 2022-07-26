package inspector

import (
	"context"
	"fmt"

	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/inspect"
	"github.com/forta-network/forta-node/clients"
	"github.com/forta-network/forta-node/clients/messaging"
	"github.com/forta-network/forta-node/config"
	log "github.com/sirupsen/logrus"
)

// Inspector runs continuous inspections.
type Inspector struct {
	ctx context.Context
	cfg InspectorConfig

	msgClient clients.MessageClient

	lastErr health.ErrorTracker

	inspectEvery int
	inspectTrace bool
	inspectCh    chan uint64
}

type InspectorConfig struct {
	Config    config.Config
	ProxyHost string
	ProxyPort string
}

func (ins *Inspector) Start() error {
	for {
		select {
		case <-ins.ctx.Done():
			return nil
		case blockNum := <-ins.inspectCh:
			ins.inspectionWorker(ins.ctx, blockNum)
		}
	}
}

func (ins *Inspector) inspectionWorker(ctx context.Context, blockNum uint64) {
	results, err := inspect.Inspect(ctx, inspect.InspectionConfig{
		ScanAPIURL:  ins.cfg.Config.Scan.JsonRpc.Url,
		ProxyAPIURL: fmt.Sprintf("http://%s:%s", ins.cfg.ProxyHost, ins.cfg.ProxyPort),
		TraceAPIURL: ins.cfg.Config.Trace.JsonRpc.Url,
		BlockNumber: blockNum,
		CheckTrace:  ins.inspectTrace,
	})
	if err != nil {
		log.WithError(err).Warn("error(s) during inspection")
	}

	ins.msgClient.PublishProto(messaging.SubjectInspectionDone, results.ToProto())
	return
}

func (ins *Inspector) registerMessageHandlers() {
	ins.msgClient.Subscribe(messaging.SubjectScannerBlock, messaging.ScannerHandler(ins.handleScannerBlock))
}

func (ins *Inspector) handleScannerBlock(payload messaging.ScannerPayload) error {
	if payload.LatestBlockInput%uint64(ins.inspectEvery) == 0 {
		ins.inspectCh <- payload.LatestBlockInput
	}
	return nil
}

func (ins *Inspector) Stop() error {
	return nil
}

func (ins *Inspector) Name() string {
	return "inspector"
}

// Health implements health.Reporter interface.
func (ins *Inspector) Health() health.Reports {
	return health.Reports{
		ins.lastErr.GetReport("last-error"),
	}
}

func NewInspector(ctx context.Context, cfg InspectorConfig) (*Inspector, error) {
	msgClient := messaging.NewClient("inspector", fmt.Sprintf("%s:%s", config.DockerNatsContainerName, config.DefaultNatsPort))
	chainSettings := config.GetChainSettings(cfg.Config.ChainID)

	return &Inspector{
		ctx:          ctx,
		msgClient:    msgClient,
		cfg:          cfg,
		inspectEvery: chainSettings.InspectionInterval,
		inspectTrace: chainSettings.EnableTrace,
		inspectCh:    make(chan uint64),
	}, nil
}