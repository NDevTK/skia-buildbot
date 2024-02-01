package pinpoint

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"go.skia.org/infra/go/auth"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/pinpoint/go/bot_configs"
	"go.skia.org/infra/pinpoint/go/read_values"
	"golang.org/x/oauth2/google"
)

const (
	missingRequiredParamTemplate = "Missing required param %s"
	chromiumSrcGit               = "https://chromium.googlesource.com/chromium/src.git"
)

// PinpointHandler is an interface to run Pinpoint jobs
type PinpointHandler interface {
	// Schedule schedules a Pinpoint job and validate the inputs.
	// jobID is an optional argument for local testing. Setting the same
	// jobID can reuse swarming results which can be helpful to triage
	// the workflow and not wait on tasks to finish.
	// TODO(sunxiaodi@): implement Schedule
	Schedule(ctx context.Context, req PinpointRunRequest, jobID string) (*PinpointRunResponse, error)
}

// PinpointRunRequest is the request arguments to run a Pinpoint job.
type PinpointRunRequest struct {
	// Device is the device to test Chrome on i.e. linux-perf
	Device string
	// Benchmark is the benchmark to test
	Benchmark string
	// Story is the benchmark's story to test
	Story string
	// Chart is the story's subtest to measure. Only used in bisections.
	Chart string
	// Magnitude is the expected absolute difference of a potential regression.
	// Only used in bisections. Default is 1.0.
	Magnitude float64
	// AggregationMethod is the method to aggregate the measurements after a single
	// benchmark runs. Some benchmarks will output multiple values in one
	// run. Aggregation is needed to be consistent with perf measurements.
	// Only used in bisection.
	AggregationMethod read_values.AggDataMethodEnum
	// StartCommit is the base or start commit hash to run
	StartCommit string
	// EndCommit is the experimental or end commit hash to run
	EndCommit string
}

type PinpointRunResponse struct {
	// JobID is the unique job ID.
	JobID string
	// Culprits is a list of culprits found in a bisection run.
	Culprits []string
}

// pinpointJobImpl implements the PinpointJob interface.
type pinpointHandlerImpl struct {
	client *http.Client
}

func New(ctx context.Context) (*pinpointHandlerImpl, error) {
	httpClientTokenSource, err := google.DefaultTokenSource(ctx, auth.ScopeReadOnly)
	if err != nil {
		return nil, skerr.Wrapf(err, "Problem setting up default token source")
	}
	c := httputils.DefaultClientConfig().WithTokenSource(httpClientTokenSource).With2xxOnly().Client()

	return &pinpointHandlerImpl{
		client: c,
	}, nil
}

// Schedule implements the pinpointJobImpl interface
func (pp *pinpointHandlerImpl) Schedule(ctx context.Context, req PinpointRunRequest, jobID string) (
	*PinpointRunResponse, error) {
	if jobID == "" {
		jobID = uuid.New().String()
	}
	err := pp.validateRunRequest(req)
	if err != nil {
		return nil, skerr.Wrapf(err, "Could not validate request inputs")
	}

	resp := &PinpointRunResponse{
		JobID:    jobID,
		Culprits: []string{},
	}
	return resp, nil
}

// validateRunRequest validates the request args and returns an error if there request is invalid
func (job *pinpointHandlerImpl) validateRunRequest(req PinpointRunRequest) error {
	if req.StartCommit == "" {
		return skerr.Fmt(missingRequiredParamTemplate, "start commit")
	}
	if req.EndCommit == "" {
		return skerr.Fmt(missingRequiredParamTemplate, "end commit")
	}
	_, err := bot_configs.GetBotConfig(req.Device, false)
	if err != nil {
		return skerr.Wrapf(err, "Device %s not allowed in bot configurations", req.Device)
	}
	if req.Benchmark == "" {
		return skerr.Fmt(missingRequiredParamTemplate, "benchmark")
	}
	if req.Story == "" {
		return skerr.Fmt(missingRequiredParamTemplate, "story")
	}
	if req.Chart == "" {
		return skerr.Fmt(missingRequiredParamTemplate, "chart")
	}
	return nil
}
