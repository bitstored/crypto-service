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

func NewCryptoServer(s *service.CryptoService) *Server {
	return &Server{s}
}

func (s *Server) EncryptFile(ctx context.Context, in *pb.EncryptFileRequest) (*pb.EncryptFileResponse, error) {

	file := in.GetOriginalFile()
	if file == nil {
		return nil, status.Error(codes.InvalidArgument, "file is nil")
	}

	if file.GetContent() == nil || len(file.GetContent()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "file is empty")
	}

	hash, err := s.cryptoService.EncryptFile(ctx, file.GetContent(), file.GetSecretPhrase())

	if err != nil {
		return nil, err
	}

	return &pb.EncryptFileResponse{EncryptedData: hash}, nil
}

func (s *Server) DecryptFile(ctx context.Context, in *pb.DecryptFileRequest) (*pb.DecryptFileResponse, error) {

	file := in.GetEncryptedFile()
	if file == nil {
		return nil, status.Error(codes.InvalidArgument, "file is nil")
	}

	if file.GetContent() == nil || len(file.GetContent()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "file is empty")
	}

	data, err := s.cryptoService.DecryptFile(ctx, file.GetContent(), file.GetSecretPhrase())

	if err != nil {
		return nil, err
	}

	return &pb.DecryptFileResponse{OriginalData: data}, nil
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
	hash, err := s.cryptoService.EncryptPassword(ctx, in.GetPassword(), in.GetSalt(), int(in.GetIterationCount()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.EncryptPasswordResponse{Password: hash}, nil
}

func (s *Server) Encrypt(ctx context.Context, in *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	data := in.GetData()
	if data == nil {
		return nil, status.Error(codes.InvalidArgument, "data is nil")
	}

	salt := in.GetSalt()
	if salt == nil || len(salt) == 0 {
		return nil, status.Error(codes.InvalidArgument, "data is empty")
	}

	hash, err := s.cryptoService.Encrypt(ctx, data, salt)

	if err != nil {
		return nil, err
	}

	return &pb.EncryptResponse{Data: hash}, nil
}

func (s *Server) Decrypt(ctx context.Context, in *pb.DecryptRequest) (*pb.DecryptResponse, error) {
	data := in.GetData()
	if data == nil {
		return nil, status.Error(codes.InvalidArgument, "data is nil")
	}

	salt := in.GetSalt()
	if salt == nil || len(salt) == 0 {
		return nil, status.Error(codes.InvalidArgument, "data is empty")
	}

	hash, err := s.cryptoService.Decrypt(ctx, data, salt)

	if err != nil {
		return nil, err
	}

	return &pb.DecryptResponse{Data: hash}, nil
}
