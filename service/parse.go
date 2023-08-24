package service

import (
	pbloc "github.com/dzwiedz90/helpdesk-proto/common"
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-users/model"
)

func (s *Server) parseStructFromCreateRequest(req *pb.CreateUserRequest) *model.CreateUser {
	u := req.GetUser()
	if u != nil {
		return &model.CreateUser{
			Username:  u.GetUsername(),
			Password:  u.GetPassword(),
			Email:     u.GetEmail(),
			FirstName: u.GetFirstName(),
			LastName:  u.GetLastName(),
			Age:       u.GetAge(),
			Gender:    u.GetGender(),
			Address: &model.Address{
				Street:     u.Address.GetStreet(),
				City:       u.Address.GetCity(),
				PostalCode: u.Address.GetPostalCode(),
				Country:    u.Address.GetCountry(),
			},
		}
	} else {
		return nil
	}
}

func (s *Server) parseStructFromUpdateRequest(req *pb.UpdateUserRequest) *model.UpdateUser {
	u := req.GetUser()
	return &model.UpdateUser{
		UserId:    u.GetUserId(),
		Username:  u.GetUsername(),
		Email:     u.GetEmail(),
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Age:       u.GetAge(),
		Gender:    u.GetGender(),
		Address: &model.Address{
			Street:     u.Address.GetStreet(),
			City:       u.Address.GetCity(),
			PostalCode: u.Address.GetPostalCode(),
			Country:    u.Address.GetCountry(),
		},
	}
}

func (s *Server) parseStructFromGetResponse(u *model.User) *pb.User {
	return &pb.User{
		UserId:    u.GetId(),
		Username:  u.GetUsername(),
		Email:     u.GetEmail(),
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Age:       u.GetAge(),
		Gender:    u.GetGender(),
		Address: &pbloc.Address{
			Street:     u.GetAddress().Street,
			City:       u.GetAddress().City,
			PostalCode: u.GetAddress().PostalCode,
			Country:    u.GetAddress().Country,
		},
	}
}

func (s *Server) parseStructFromGetAllResponse(u []*model.User) []*pb.User {
	var users []*pb.User
	for _, u := range u {
		users = append(users, &pb.User{
			UserId:    u.GetId(),
			Username:  u.GetUsername(),
			Email:     u.GetEmail(),
			FirstName: u.GetFirstName(),
			LastName:  u.GetLastName(),
			Age:       u.GetAge(),
			Gender:    u.GetGender(),
			Address: &pbloc.Address{
				Street:     u.GetAddress().Street,
				City:       u.GetAddress().City,
				PostalCode: u.GetAddress().PostalCode,
				Country:    u.GetAddress().Country,
			},
		})
	}
	return users
}
