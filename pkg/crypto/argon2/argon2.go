package argon2

import (
	"fmt"
	"golang.org/x/crypto/argon2"
	"time"
)

type Argon2Arguments struct {
	Hash    []byte
	Salt    []byte
	Time    uint32
	Memory  uint32
	Threads uint8
}

const expectedDuration = 1 * time.Second

func EncryptNew(password, salt []byte) (*Argon2Arguments, error) {
	fsm := Init(Idle)
	fsm.CallibrateArgon2(expectedDuration)

	dur := fsm.GetDuration()
	mem := fsm.GetMemory()
	cpu := fsm.GetThreads()

	hash, err := Encrypt(password, salt, dur, mem, cpu)
	if err != nil {
		return nil, err
	}

	return &Argon2Arguments{hash, salt, dur, mem, cpu}, nil
}

func Encrypt(password, salt []byte, time, memory uint32, threads uint8) ([]byte, error) {
	if password == nil || len(password) == 0 {
		return nil, fmt.Errorf("password is nil or empty")
	}

	if salt == nil || len(salt) == 0 {
		return nil, fmt.Errorf("password is nil or empty")
	}

	keyLen := uint32(len(password))

	return argon2.IDKey(password, salt, time, memory, threads, keyLen), nil
}
