package app

import (
	"github.com/arquivei/go-app"
)

type BeamConfiger interface {
	app.AppConfig
	BeamConfig() *Config
}

// Config groups both configuration or our application and apache beam.
type Config struct {
	app.Config

	// Dataflow specific configuration.
	// BE CAREFUL WHEN EDITTING THESE VALUES
	// They were made to be compatible with all
	// flags defined by datalfow.

	Beam struct {
		Async                            bool     `flag:"async"`
		AutoscalingAlgorithm             string   `flag:"autoscaling_algorithm" env:"BEAM_AUTOSCALING_ALGORITHM"`
		BeamStrict                       bool     `flag:"beam_strict" env:"BEAM_STRICT"`
		ControlEndpoint                  string   `flag:"control_endpoint" env:"BEAM_CONTROL_ENDPOINT"`
		CPUProfiling                     string   `flag:"cpu_profiling" env:"BEAM_CPU_PROFILING"`
		DataflowEndpoint                 string   `flag:"dataflow_endpoint" env:"BEAM_DATAFLOW_ENDPOINT"`
		DataflowWorkerJar                string   `flag:"dataflow_worker_jar" env:"BEAM_DATAFLOW_WORKER_JAR"`
		DataflowServiceOptions           string   `flag:"dataflow_service_options" env:"BEAM_DATAFLOW_SERVICE_OPTIONS"`
		DiskSizeGb                       int      `flag:"disk_size_gb" env:"BEAM_DISK_SIZE_GB"`
		DotFile                          string   `flag:"dot_file" env:"BEAM_DOT_FILE"`
		DryRun                           bool     `flag:"dry_run" env:"BEAM_DRY_RUN"`
		Endpoint                         string   `flag:"endpoint"`
		EnvironmentConfig                string   `flag:"environment_config" env:"BEAM_ENVIRONMENT_CONFIG"`
		EnvironmentType                  string   `flag:"environment_type" env:"BEAM_ENVIRONMENT_TYPE"`
		ExecuteAsync                     bool     `flag:"execute_async" env:"BEAM_EXECUTE_ASYNC"`
		Experiments                      string   `flag:"experiments"`
		ID                               string   `flag:"id"`
		JobName                          string   `flag:"job_name" env:"BEAM_JOB_NAME"`
		Labels                           string   `flag:"labels"`
		LoggingEndpoint                  string   `flag:"logging_endpoint" env:"BEAM_LOGGING_ENDPOINT"`
		MaxNumWorkers                    int      `flag:"max_num_workers" env:"BEAM_MAX_NUM_WORKERS"`
		MinCPUPlatform                   string   `flag:"min_cpu_platform" env:"BEAM_MIN_CPU_PLATFORM"`
		Network                          string   `flag:"network"`
		NoUsePublicIps                   bool     `flag:"no_use_public_ips" env:"BEAM_NO_USE_PUBLIC_IPS"`
		NumWorkers                       int      `flag:"num_workers" env:"BEAM_NUM_WORKERS"`
		Options                          string   `flag:"options"`
		Parallelism                      int      `flag:"parallelism"`
		Project                          string   `flag:"project"`
		Region                           string   `flag:"region"`
		RetainDockerContainers           bool     `flag:"retain_docker_containers" env:"BEAM_RETAIN_DOCKER_CONTAINERS"`
		Runner                           string   `flag:"runner"`
		SdkHarnessContainerImageOverride []string `flag:"sdk_harness_container_image_override" env:"BEAM_SDK_HARNESS_CONTAINER_IMAGE_OVERRIDE"`
		SemiPersistDir                   string   `flag:"semi_persist_dir" env:"BEAM_SEMI_PERSIST_DIR"`
		ServiceAccountEmail              string   `flag:"service_account_email" env:"BEAM_SERVICE_ACCOUNT_EMAIL"`
		SessionRecording                 string   `flag:"session_recording" env:"BEAM_SESSION_RECORDING"`
		SmallWordLength                  int      `flag:"small_word_length" env:"BEAM_SMALL_WORD_LENGTH"`
		StagingLocation                  string   `flag:"staging_location" env:"BEAM_STAGING_LOCATION"`
		StatusEndpoint                   string   `flag:"status_endpoint" env:"BEAM_STATUS_ENDPOINT"`
		Subnetwork                       string   `flag:"subnetwork"`
		TeardownPolicy                   string   `flag:"teardown_policy" env:"BEAM_TEARDOWN_POLICY"`
		TempLocation                     string   `flag:"temp_location" env:"BEAM_TEMP_LOCATION"`
		WithMetrics                      bool     `flag:"with-metrics" env:"BEAM_WITH_METRICS"`
		Worker                           bool     `flag:"worker"`
		WorkerBinary                     string   `flag:"worker_binary" env:"BEAM_WORKER_BINARY"`
		WorkerHarnessContainerImage      string   `flag:"worker_harness_container_image" env:"BEAM_WORKER_HARNESS_CONTAINER_IMAGE" default:"apache/beam_go_sdk:latest"`
		WorkerMachineType                string   `flag:"worker_machine_type" env:"BEAM_WORKER_MACHINE_TYPE"`
		WorkerRegion                     string   `flag:"worker_region" env:"BEAM_WORKER_REGION"`
		WorkerZone                       string   `flag:"worker_zone" env:"BEAM_WORKER_ZONE"`
		Zone                             string   `flag:"zone"`
	}
}

func (c *Config) BeamConfig() *Config {
	return c
}
