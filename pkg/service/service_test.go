package service

import (
	"context"
	"reflect"
	"testing"
)

func TestCryptoService_EncryptFile(t *testing.T) {
	type args struct {
		ctx          context.Context
		content      []byte
		secretPhrase []byte
	}
	tests := []struct {
		name    string
		cs      *CryptoService
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CryptoService{}
			got, err := cs.EncryptFile(tt.args.ctx, tt.args.content, tt.args.secretPhrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoService.EncryptFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptoService.EncryptFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptoService_DecryptFile(t *testing.T) {
	type args struct {
		ctx          context.Context
		content      []byte
		secretPhrase []byte
	}
	tests := []struct {
		name    string
		cs      *CryptoService
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CryptoService{}
			got, err := cs.DecryptFile(tt.args.ctx, tt.args.content, tt.args.secretPhrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoService.DecryptFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptoService.DecryptFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptoService_EncryptPassword(t *testing.T) {
	type args struct {
		ctx      context.Context
		password []byte
		salt     []byte
		iter     int
	}
	tests := []struct {
		name     string
		cs       *CryptoService
		args     args
		wantHash []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CryptoService{}
			gotHash, err := cs.EncryptPassword(tt.args.ctx, tt.args.password, tt.args.salt, tt.args.iter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoService.EncryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHash, tt.wantHash) {
				t.Errorf("CryptoService.EncryptPassword() = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}

func TestCryptoService_Encrypt(t *testing.T) {
	type args struct {
		ctx  context.Context
		data []byte
		salt []byte
	}
	tests := []struct {
		name    string
		cs      *CryptoService
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CryptoService{}
			got, err := cs.Encrypt(tt.args.ctx, tt.args.data, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoService.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptoService.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptoService_Decrypt(t *testing.T) {
	type args struct {
		ctx  context.Context
		data []byte
		salt []byte
	}
	tests := []struct {
		name    string
		cs      *CryptoService
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CryptoService{}
			got, err := cs.Decrypt(tt.args.ctx, tt.args.data, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoService.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptoService.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCryptoService(t *testing.T) {
	tests := []struct {
		name string
		want *CryptoService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCryptoService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCryptoService() = %v, want %v", got, tt.want)
			}
		})
	}
}
