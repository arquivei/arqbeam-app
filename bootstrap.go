package app

import (
	"context"
	"flag"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/arquivei/go-app"
	"github.com/rs/zerolog/log"
)

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
}

// Run
func Run(p *beam.Pipeline) {
	app.RunAndWait(func(ctx context.Context) error {
		err := beamx.Run(ctx, p)
		if err != nil {
			log.Error().Err(err).Msg("[arqbeam] Pipeline failed.")

			return err
		}
		log.Info().Msg("[arqbeam] Pipeline finished.")

		return nil
	})
}

func RunWithMetrics(p *beam.Pipeline) {
	app.RunAndWait(func(ctx context.Context) error {
		results, err := beamx.RunWithMetrics(ctx, p)
		if err != nil {
			log.Error().Err(err).Msg("[arqbeam] Pipeline failed.")

			return err
		}
		log.Info().
			Str("job_id", results.JobID()).
			Interface("metrics", results.Metrics().AllMetrics()).
			Msg("[arqbeam] Pipeline finished.")

		return nil
	})
}

func RunWithCallback(p *beam.Pipeline, onSuccess func(), onFailure func(error)) {
	app.RunAndWait(func(ctx context.Context) error {
		err := beamx.Run(ctx, p)
		if err != nil {
			log.Error().Err(err).Msg("[arqbeam] Pipeline failed.")
			if onFailure != nil {
				log.Info().Msg("[arqbeam] Calling onFailure Callback.")
				onFailure(err)
			}
			return err
		}
		log.Info().Msg("[arqbeam] Pipeline finished.")
		if onSuccess != nil {
			log.Info().Msg("[arqbeam] Calling onSuccess Callback.")
			onSuccess()
		}
		return nil
	})
}