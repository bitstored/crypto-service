package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bitstored/crypto-service/pb"
	"github.com/bitstored/crypto-service/pkg/service"
)

func TestNewCryptoServer(t *testing.T) {
	type args struct {
		s *service.CryptoService
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCryptoServer(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCryptoServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_EncryptFile(t *testing.T) {
	type fields struct {
		cryptoService *service.CryptoService
	}
	type args struct {
		ctx context.Context
		in  *pb.EncryptFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.EncryptFileResponse
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				cryptoService: service.NewCryptoService(),
			},
			args: args{
				ctx: context.Background(),
				in: &pb.EncryptFileRequest{
					OriginalFile: &pb.File{
						Content:      []byte("ana"),
						SecretPhrase: []byte("ana1"),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				cryptoService: tt.fields.cryptoService,
			}
			got, err := s.EncryptFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.EncryptFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got1, err := s.DecryptFile(tt.args.ctx, &pb.DecryptFileRequest{
				EncryptedFile: &pb.File{
					Content:      got.EncryptedData,
					SecretPhrase: tt.args.in.GetOriginalFile().GetSecretPhrase(),
				},
			})
			if !reflect.DeepEqual(tt.args.in.GetOriginalFile().GetContent(), got1.GetOriginalData()) {
				t.Errorf("Server.EncryptFile() = %v, want %v", tt.args.in.GetOriginalFile().Content, got1.GetOriginalData())
			}
		})
	}
}

func TestServer_DecryptFile(t *testing.T) {
	type fields struct {
		cryptoService *service.CryptoService
	}
	type args struct {
		ctx context.Context
		in  *pb.DecryptFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DecryptFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				cryptoService: tt.fields.cryptoService,
			}
			got, err := s.DecryptFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DecryptFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DecryptFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_EncryptPassword(t *testing.T) {
	type fields struct {
		cryptoService *service.CryptoService
	}
	type args struct {
		ctx context.Context
		in  *pb.EncryptPasswordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.EncryptPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				cryptoService: tt.fields.cryptoService,
			}
			got, err := s.EncryptPassword(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.EncryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.EncryptPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Encrypt(t *testing.T) {
	type fields struct {
		cryptoService *service.CryptoService
	}
	type args struct {
		ctx context.Context
		in  *pb.EncryptRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.EncryptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				cryptoService: tt.fields.cryptoService,
			}
			got, err := s.Encrypt(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Decrypt(t *testing.T) {
	type fields struct {
		cryptoService *service.CryptoService
	}
	type args struct {
		ctx context.Context
		in  *pb.DecryptRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DecryptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				cryptoService: tt.fields.cryptoService,
			}
			got, err := s.Decrypt(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
