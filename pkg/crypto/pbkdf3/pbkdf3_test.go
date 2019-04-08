package pbkdf3_test

import (
	"github.com/bitstored/crypto-service/pkg/crypto/pbkdf3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIdenticalHashWithSHA256(t *testing.T) {
	testSuite := []struct {
		data []byte
		salt []byte
		iter int
	}{{[]byte("password"), []byte("salt"), 1},
		{[]byte("password"), []byte("salt"), 2},
		{[]byte("password"), []byte("salt"), 4},
		{[]byte("password"), []byte("salt"), 1024},
		{[]byte("password"), []byte("salt"), 4096},
		{[]byte("password"), []byte("salt"), 8192},
		{[]byte("password"), []byte("salt"), 2 * 8192}}
	outputs := make([][]byte, 0)
	for _, tt := range testSuite {
		hash := pbkdf3.HashWithSHA256(tt.data, tt.salt, tt.iter, 32)
		outputs = append(outputs, hash)
		require.EqualValuesf(t, 32, len(hash), "length should be equal")
	}
	for i := 0; i < len(testSuite); i++ {
		for j := 0; j < len(testSuite); j++ {
			if j == i {
				continue
			}
			require.NotEqualf(t, outputs[i], outputs[j], "hashes on same data with different parameters should not be equal")
		}
	}
	output := pbkdf3.HashWithSHA256(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
	for i := 0; i < 100; i++ {
		output1 := pbkdf3.HashWithSHA256(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
		require.EqualValuesf(t, output, output1, "values should be equal")
	}
}

func TestIdenticalHashWithSHA512(t *testing.T) {
	testSuite := []struct {
		data []byte
		salt []byte
		iter int
	}{{[]byte("password"), []byte("salt"), 1},
		{[]byte("password"), []byte("salt"), 2},
		{[]byte("password"), []byte("salt"), 4},
		{[]byte("password"), []byte("salt"), 1024},
		{[]byte("password"), []byte("salt"), 4096},
		{[]byte("password"), []byte("salt"), 8192},
		{[]byte("password"), []byte("salt"), 2 * 8192}}
	outputs := make([][]byte, 0)
	for _, tt := range testSuite {
		hash := pbkdf3.HashWithSHA512(tt.data, tt.salt, tt.iter, 32)
		outputs = append(outputs, hash)
		require.EqualValuesf(t, 32, len(hash), "length should be equal")
	}
	for i := 0; i < len(testSuite); i++ {
		for j := 0; j < len(testSuite); j++ {
			if j == i {
				continue
			}
			require.NotEqualf(t, outputs[i], outputs[j], "hashes on same data with different parameters should not be equal")
		}
	}
	output := pbkdf3.HashWithSHA512(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
	for i := 0; i < 100; i++ {
		output1 := pbkdf3.HashWithSHA512(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
		require.EqualValuesf(t, output, output1, "values should be equal")
	}
}

func TestDifferentHashWithSHA256(t *testing.T) {
	testSuite := []struct {
		data []byte
		salt []byte
		iter int
	}{{[]byte("i have a password"), []byte("saltysalt"), 4096},
		{[]byte("i dont have a password"), []byte("bigsmall"), 4096},
		{[]byte("12321229j3222"), []byte("939399seh"), 4096},
		{[]byte("alorap"), []byte("salt"), 4096},
		{[]byte("parola"), []byte("tlas"), 4096},
		{[]byte("alorap"), []byte("tlas"), 4096},
		{[]byte("parola"), []byte("salt"), 4096},
		{[]byte("Passw0rd1sS3cur3"), []byte("don't have salt"), 4096},
		{[]byte("mylittlepassword"), []byte("salTySalTdffB"), 4096},
		{[]byte("AAAAAAAAAAAA"), []byte("123456789"), 4096}}
	outputs := make([][]byte, 0)
	for _, tt := range testSuite {
		hash := pbkdf3.HashWithSHA256(tt.data, tt.salt, tt.iter, 32)
		outputs = append(outputs, hash)
		require.EqualValuesf(t, 32, len(hash), "length should be equal")
	}
	for i := 0; i < len(testSuite); i++ {
		for j := 0; j < len(testSuite); j++ {
			if j == i {
				continue
			}
			require.NotEqualf(t, outputs[i], outputs[j], "hashes on different data with different salt should not be equal")
		}
	}
	output := pbkdf3.HashWithSHA256(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
	for i := 0; i < 100; i++ {
		output1 := pbkdf3.HashWithSHA256(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
		require.EqualValuesf(t, output, output1, "values should be equal")
	}
}

func TestDifferentHashWithSHA512(t *testing.T) {
	testSuite := []struct {
		data []byte
		salt []byte
		iter int
	}{{[]byte("i have a password"), []byte("saltysalt"), 4096},
		{[]byte("i dont have a password"), []byte("bigsmall"), 4096},
		{[]byte("12321229j3222"), []byte("939399seh"), 4096},
		{[]byte("alorap"), []byte("salt"), 4096},
		{[]byte("parola"), []byte("tlas"), 4096},
		{[]byte("alorap"), []byte("tlas"), 4096},
		{[]byte("parola"), []byte("salt"), 4096},
		{[]byte("Passw0rd1sS3cur3"), []byte("don't have salt"), 4096},
		{[]byte("mylittlepassword"), []byte("salTySalTdffB"), 4096},
		{[]byte("AAAAAAAAAAAA"), []byte("123456789"), 4096}}
	outputs := make([][]byte, 0)
	for _, tt := range testSuite {
		hash := pbkdf3.HashWithSHA512(tt.data, tt.salt, tt.iter, 32)
		outputs = append(outputs, hash)
		require.EqualValuesf(t, 32, len(hash), "length should be equal")
	}
	for i := 0; i < len(testSuite); i++ {
		for j := 0; j < len(testSuite); j++ {
			if j == i {
				continue
			}
			require.NotEqualf(t, outputs[i], outputs[j], "hashes on different data with different salt should not be equal")
		}
	}
	output := pbkdf3.HashWithSHA512(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
	for i := 0; i < 100; i++ {
		output1 := pbkdf3.HashWithSHA512(testSuite[0].data, testSuite[0].salt, 4096, len(testSuite[0].data))
		require.EqualValuesf(t, output, output1, "values should be equal")
	}
}
