package service

import (
	"context"
	"fmt"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-users/driver"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
	"github.com/dzwiedz90/helpdesk-service-users/pkg/core"
)

type Server struct {
	pb.UsersServiceServer
	ServerConfig *ServerConfig
}

type ServerConfig struct {
	DB *driver.DB
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	creq := s.parseStructFromRequest(req)
	id, err := core.CreateUser(ctx, s.ServerConfig.DB, creq)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to create user within core: %v", err))
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := core.GetUser(ctx, s.ServerConfig.DB, req.GetId())
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to get user from core: %v", err))
		return nil, err
	}

	return &pb.GetUserResponse{
		User: s.parseStructFromGetResponse(user),
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	return nil, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}
