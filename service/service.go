package service

import (
	"context"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
)

type Server struct {
	pb.UsersServiceServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// parse request to local request
	// validate stuff
	// get resposne from createUser function from facade
	// return response
	return &pb.CreateUserResponse{
		Id: 666,
	}, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
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
