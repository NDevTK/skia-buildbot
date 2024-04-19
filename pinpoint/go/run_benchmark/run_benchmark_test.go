package run_benchmark

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/swarming"
	"go.skia.org/infra/go/swarming/mocks"
	"go.skia.org/infra/pinpoint/go/backends"

	swarmingV1 "go.chromium.org/luci/common/api/swarming/swarming/v1"
)

var fakeBotID = map[string]string{
	"key":   "id",
	"value": "fake-botid-h7",
}
var req = RunBenchmarkRequest{
	JobID:     "id",
	Benchmark: "benchmark",
	Story:     "story",
	Build: &swarmingV1.SwarmingRpcsCASReference{
		CasInstance: "instance",
		Digest: &swarmingV1.SwarmingRpcsDigest{
			Hash:      "hash",
			SizeBytes: 0,
		},
	},
	Commit: "64893ca6294946163615dcf23b614afe0419bfa3",
}
var expectedErr = skerr.Fmt("some error")

func TestRun_TelemetryTest_ValidExecution(t *testing.T) {
	ctx := context.Background()
	mockClient := mocks.NewApiClient(t)
	sc := &backends.SwarmingClientImpl{
		ApiClient: mockClient,
	}

	buildArtifact := &swarmingV1.SwarmingRpcsCASReference{
		CasInstance: "instance",
		Digest: &swarmingV1.SwarmingRpcsDigest{
			Hash:      "hash",
			SizeBytes: 0,
		},
	}

	c, fakeID := "64893ca6294946163615dcf23b614afe0419bfa3", "fake-id"

	mockClient.On("TriggerTask", ctx, mock.Anything).
		Return(&swarmingV1.SwarmingRpcsTaskRequestMetadata{
			TaskId: "123",
		}, nil).Once()
	taskIds, err := Run(ctx, sc, c, "android-pixel2_webview-perf", "performance_browser_tests", "story", "all", fakeID, buildArtifact, 1, fakeBotID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(taskIds))
	assert.Equal(t, "123", taskIds[0].TaskId)
}

func TestIsTaskFinished_GivenCompleteStates_ReturnsTrue(t *testing.T) {
	states := []string{
		swarming.TASK_STATE_COMPLETED,
		swarming.TASK_STATE_BOT_DIED,
		swarming.TASK_STATE_TIMED_OUT,
	}
	for _, s := range states {
		state := State(s)
		out := state.IsTaskFinished()
		assert.True(t, out)
	}
}

func TestIsTaskFinished_GivenRunningStates_ReturnsFalse(t *testing.T) {
	states := []string{
		swarming.TASK_STATE_PENDING,
		swarming.TASK_STATE_RUNNING,
	}
	for _, s := range states {
		state := State(s)
		out := state.IsTaskFinished()
		assert.False(t, out)
	}
}

func TestIsTaskSuccessful_GivenCompleted_ReturnsTrue(t *testing.T) {
	states := []string{
		swarming.TASK_STATE_COMPLETED,
	}
	for _, s := range states {
		state := State(s)
		out := state.IsTaskSuccessful()
		assert.True(t, out)
	}
}

func TestIsTaskSuccessful_GivenNonCompleted_ReturnsFalse(t *testing.T) {
	states := []string{
		swarming.TASK_STATE_PENDING,
		swarming.TASK_STATE_RUNNING,
		swarming.TASK_STATE_BOT_DIED,
		swarming.TASK_STATE_CANCELED,
		swarming.TASK_STATE_TIMED_OUT,
	}
	for _, s := range states {
		state := State(s)
		out := state.IsTaskSuccessful()
		assert.False(t, out)
	}
}

func TestIsTaskTerminalFailure_GivenTerminalState_ReturnsTrue(t *testing.T) {
	states := []State{
		swarming.TASK_STATE_BOT_DIED,
		swarming.TASK_STATE_CANCELED,
		swarming.TASK_STATE_TIMED_OUT,
	}
	for _, s := range states {
		out := s.IsTaskTerminalFailure()
		assert.True(t, out)
	}
}

func TestIsTaskTerminalFailure_GivenNonTerminalState_ReturnsFalse(t *testing.T) {
	states := []State{
		swarming.TASK_STATE_RUNNING,
		backends.RunBenchmarkFailure,
		swarming.TASK_STATE_COMPLETED,
	}
	for _, s := range states {
		out := s.IsTaskTerminalFailure()
		assert.False(t, out)
	}
}
