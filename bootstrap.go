package app

import (
	"context"
	"flag"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/arquivei/go-app"
	"github.com/rs/zerolog/log"
)

type OnSuccessHook func()
type OnFailureHook func(error)

var (
	onSuccess []OnSuccessHook
	onFailure []OnFailureHook

	runWithMetrics bool
)

func RegisterOnSuccessHooks(hooks ...OnSuccessHook) {
	onSuccess = append(onSuccess, hooks...)
}

func RegisterOnFailureHooks(hooks ...OnFailureHook) {
	onFailure = append(onFailure, hooks...)
}

// Bootstrap will set flags for the beam package. The values will be get from
// config struct so this function should be called only after `app.SetupConfig(&config)`.
func Bootstrap(appVersion string, config BeamConfiger) {
	// Normal app bootstrap
	app.Bootstrap(appVersion, config)

	// Set some magic default values
	c := config.BeamConfig()
	if c.Beam.WorkerBinary == "" {
		c.Beam.WorkerBinary = inferBinaryName()
	}

	mergeWithBeamFlags(config)

	flag.Parse()
	beam.Init()

	// Copy run configuration
	runWithMetrics = c.Beam.WithMetrics
}

// Run
func Run(p *beam.Pipeline) {
	app.RunAndWait(func(ctx context.Context) error {
		logger := log.With().Logger()

		results, err := runPipeline(ctx, p)

		if results != nil {
			logger = logger.With().
				Str("job_id", results.JobID()).
				Logger()
		}

		if err != nil {
			logger.Error().Err(err).Msg("[arqbeam-app] Pipeline failed.")
			for _, hook := range onFailure {
				hook(err)
			}

			return err
		}

		logger.Info().Msg("[arqbeam-app] Pipeline finished.")
		for _, hook := range onSuccess {
			hook()
		}

		return nil
	})
}

func runPipeline(ctx context.Context, p *beam.Pipeline) (beam.PipelineResult, error) {
	if runWithMetrics {
		return beamx.RunWithMetrics(ctx, p)
	}

	return nil, beamx.Run(ctx, p)
}
