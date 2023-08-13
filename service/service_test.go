package service

import (
	"context"
	"testing"

	pbloc "github.com/dzwiedz90/helpdesk-proto/common"
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
)

var (
	validRequest = &pb.CreateUserRequest{
		User: &pb.User{
			Username:  "james.kirk",
			Password:  "password",
			Email:     "james.kirk@enterprise.com",
			FirstName: "James",
			LastName:  "Kirk",
			Age:       32,
			Gender:    "male",
			Address: &pbloc.Address{
				Street:     "Teststreet 1",
				City:       "Somecity",
				PostalCode: "2137",
				Country:    "Romulan Federation",
			},
		},
	}
	emptyRequest = &pb.CreateUserRequest{}
)

func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name string
		req  *pb.CreateUserRequest
		resp *pb.CreateUserResponse
	}{
		{
			name: "valid request",
			req:  validRequest,
			resp: &pb.CreateUserResponse{
				Id: 1,
			},
		},
		{
			name: "empty request",
			req:  emptyRequest,
		},
	}

	for _, tc := range testCases {
		// TODO mock server and it's responses
		s := Server{}
		resp, err := s.CreateUser(context.Background(), tc.req)
		if err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}

		if resp.GetId() != tc.resp.GetId() {
			t.Fatal("Wrong response")
		}
	}
}
