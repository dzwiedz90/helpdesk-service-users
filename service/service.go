package service

import (
	"context"
	"fmt"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-users/pkg/core"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
)

type Server struct {
	pb.UsersServiceServer
	ServerConfig *serviceconfig.ServerConfig
	Core         core.CoreAdapter
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := ValidateRequest(req)
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to validate request: %v", err))
		return nil, err
	}

	creq := s.parseStructFromCreateRequest(req)

	id, err := s.Core.CreateUser(ctx, s.ServerConfig, creq)
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to create user within core: %v", err))
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.Core.GetUser(ctx, s.ServerConfig, req.GetId())
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to get user from core: %v", err))
		return nil, err
	}

	return &pb.GetUserResponse{
		User: s.parseStructFromGetResponse(user),
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := s.Core.GetAlUsers(ctx, s.ServerConfig)
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to get all users from core: %v", err))
		return nil, err
	}

	return &pb.GetAllUsersResponse{
		Users: s.parseStructFromGetAllResponse(users),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	err := ValidateRequest(req)
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to validate request: %v", err))
		return nil, err
	}

	creq := s.parseStructFromUpdateRequest(req)

	err = s.Core.UpdateUser(ctx, s.ServerConfig, creq)
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to update user: %v", err))
		return nil, err
	}

	return &pb.UpdateUserResponse{}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.Core.DeleteUser(ctx, s.ServerConfig, req.GetId())
	if err != nil {
		s.ServerConfig.Logger.ErrorLogger(fmt.Sprintf("Failed to delete user: %v", err))
		return nil, err
	}

	return &pb.DeleteUserResponse{}, nil
}
