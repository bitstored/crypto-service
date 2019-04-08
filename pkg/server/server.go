package server

import (
	"context"
	"github.com/bitstored/crypto-service/pb"
	"github.com/bitstored/crypto-service/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cryptoService *service.CryptoService
}

func (s *Server) EncryptFile(ctx context.Context, in *pb.EncryptFileRequest) (*pb.EncryptFileResponse, error) {
	return nil, nil
}
func (s *Server) DecryptFile(ctx context.Context, in *pb.DecryptFileRequest) (*pb.DecryptFileResponse, error) {
	return nil, nil
}
func (s *Server) EncryptPassword(ctx context.Context, in *pb.EncryptPasswordRequest) (*pb.EncryptPasswordResponse, error) {
	if in.Password == nil || len(in.Password) == 0 {
		return nil, status.Error(codes.InvalidArgument, "password can't be nil or empty")
	}
	if in.Salt == nil || len(in.Salt) == 0 {
		return nil, status.Error(codes.InvalidArgument, "salt can't be nil or empty")
	}
	if in.IterationCount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "iterations count can't be negative")
	}
	hash, err := s.cryptoService.EncryptPassword(ctx, in.GetPassword(), in.GetSalt(), in.GetIterationCount())
	if err != nil {
		return nil, status.Error(codes.Internal)
	}
	return nil, nil
}
func (s *Server) Encrypt(ctx context.Context, in *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	return nil, nil
}
func (s *Server) Decrypt(ctx context.Context, in *pb.DecryptRequest) (*pb.DecryptResponse, error) {
	return nil, nil
}
