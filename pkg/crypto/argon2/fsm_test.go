package argon2

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	require.NotEmpty(t, Init(0))
	fsm := Init(1)
	require.Equal(t, uint32(1), fsm.GetDuration())
	require.Equal(t, uint8(runtime.NumCPU()), fsm.GetThreads())
	require.Equal(t, uint32(64*kilobyte), fsm.GetMemory())
}

func TestEstimation(t *testing.T) {
	fsm := Init(1)
	fsm.SetState(2)
	fsm.CallibrateArgon2(2 * time.Second)
	expectedDurationInNanoseconds := uint64((2 * time.Second).Nanoseconds()) / 1000
	computedDurationInNanoseconds := fsm.estimateTimeInNanoseconds()
	_, diff := inRange(expectedDurationInNanoseconds, computedDurationInNanoseconds)
	require.Equal(t, 0, diff)
}
