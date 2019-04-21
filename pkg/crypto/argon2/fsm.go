package argon2

import (
	"errors"
	"runtime"
	"time"
)

// UnknownStateErrorMessage for error creation
const UnknownStateErrorMessage = "Unknown state"

const (
	kilobyte    = uint(1024)
	maxMemory   = uint(kilobyte * 512) // 512 KB
	maxThreads  = uint(60)
	maxTime     = uint(100)
	minMemory   = uint(kilobyte * 64) // 64 KB
	minThreads  = uint(12)
	minTime     = uint(1)
	blockLength = uint64(128) // constants used in argon 2 algorithm
	syncPoints  = uint64(8)   // constants used in argon 2 algorithm
)
const (
	// Idle is the decision making state
	Idle = iota
	// AdjustTime is the state when time needs to be adjusted
	AdjustTime
	// AdjustMemory is the state when memory needs to be adjusted
	AdjustMemory
	// AdjustCPU is the state when the number of threads needs to be changed
	AdjustCPU
	// Fin is the state when the duration is near the desired value
	Fin
)

// FSM is the structure for the FSM
type FSM struct {
	state int
	cpu   uint
	mem   uint
	dur   uint
}

func isValidState(state int) bool {
	if state < Idle || state > Fin {
		return false
	}
	return true
}

// Init is the constructor for the
func Init(state int) *FSM {
	if !isValidState(state) {
		panic(errors.New(UnknownStateErrorMessage))
	}
	return &FSM{
		state: state,
		cpu:   uint(runtime.NumCPU()),
		mem:   kilobyte * 64,
		dur:   1}
}

// SetState changes the state of the fsm
func (fsm *FSM) SetState(state int) {
	if !isValidState(state) {
		panic(errors.New(UnknownStateErrorMessage))
	}
	fsm.state = state
}

// GetMemory gets the memory dimension parameter in current state
func (fsm *FSM) GetMemory() uint32 {
	return uint32(fsm.mem)
}

// GetDuration gets the duration parameter in current state
func (fsm *FSM) GetDuration() uint32 {
	return uint32(fsm.dur)
}

// GetThreads gets the number of threads parameter in current state
func (fsm *FSM) GetThreads() uint8 {
	return uint8(fsm.cpu)
}

func inRange(in1, in2 uint64) (smaller bool, diff int) {
	epsilon := uint64((100 * time.Microsecond).Nanoseconds())

	if in1 > in2 {
		if in1-in2 < epsilon {
			return false, 0
		}
		return false, int(in1 - in2)
	}

	if in1 < in2 {
		if in2-in1 < epsilon {
			return true, 0
		}
		return true, int(in2 - in1)
	}

	return true, 0

}

// CallibrateArgon2 changes the parametters for the Argon2 algorithm to furfit the time constraint
// depending on the state sets the argon2 parameter to the value from function parametter, other values are computed
func (fsm *FSM) CallibrateArgon2(duration time.Duration) {
	// The number of virtual cores of the machine
	expectedDurationInNanoseconds := uint64(duration.Nanoseconds())
	for {
		estimatedDurationInNanoseconds := fsm.estimateTimeInNanoseconds()
		smaller, diff := inRange(estimatedDurationInNanoseconds, expectedDurationInNanoseconds/1000)
		switch fsm.state {

		case Fin:
			return

		case Idle:

			if diff == 0 {
				fsm.state = Fin
				continue
			}

			if smaller {
				if fsm.cpu <= maxThreads && fsm.cpu >= minThreads {
					fsm.state = AdjustCPU
					continue
				}

				if fsm.mem <= maxMemory && fsm.mem >= minMemory {
					fsm.state = AdjustMemory
					continue
				}

				if fsm.dur <= maxTime && fsm.dur >= minTime {
					fsm.state = AdjustTime
					continue
				}
				fsm.state = Fin
			} else if !smaller {
				if fsm.dur <= maxTime && fsm.dur >= minTime {
					fsm.state = AdjustTime
					continue
				}

				if fsm.cpu <= maxThreads && fsm.cpu >= minThreads {
					fsm.state = AdjustCPU
					continue
				}

				if fsm.mem <= maxMemory && fsm.mem >= minMemory {
					fsm.state = AdjustMemory
					continue
				}
				fsm.state = Fin
			}

		case AdjustTime:

			if smaller {
				if fsm.dur < maxTime {
					fsm.dur++
					fsm.state = Idle
					continue
				}
			} else if !smaller {
				if fsm.dur > minTime {
					fsm.dur--
					fsm.state = Idle
					continue
				}
			}
			fsm.state = AdjustMemory

		case AdjustMemory:

			if smaller {
				if fsm.mem < maxMemory {
					fsm.mem += kilobyte
					fsm.state = Idle
					continue
				}
			} else if !smaller {
				if fsm.mem > minMemory {
					fsm.mem -= kilobyte
					fsm.state = Idle
					continue
				}
			}
			fsm.state = AdjustCPU

		case AdjustCPU:

			if smaller {
				if fsm.cpu > minThreads {
					fsm.cpu--
					fsm.state = Idle
					continue
				}
			} else if !smaller {
				if fsm.cpu < maxThreads {
					fsm.cpu++
					fsm.state = Idle
					continue
				}
			}
			fsm.state = AdjustTime

		default:
			panic(UnknownStateErrorMessage)
		}
	}
}

func (fsm *FSM) estimateTimeInNanoseconds() uint64 {
	cpus := uint64(runtime.NumCPU())
	time := uint64(0)
	bitsInByte := uint64(16)
	processBlockCalls := uint64(6)
	segments := uint64(blockLength / syncPoints)
	indexAlphaCost := uint64(20)
	initBlocksCost := uint64(fsm.cpu)
	processBlockCost := uint64(blockLength / bitsInByte)
	processSegmentCost := uint64(segments * indexAlphaCost * processBlockCalls * processBlockCost)
	extractKeyCost := uint64(fsm.mem/kilobyte/fsm.cpu + fsm.mem/kilobyte)
	processBlocksCost := uint64(uint64(fsm.dur)*syncPoints*processSegmentCost/initBlocksCost) * cpus
	time = extractKeyCost + initBlocksCost + processBlocksCost
	return time
}
