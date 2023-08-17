package service

import (
	"context"
	"log"
	"os"
	"testing"

	pbloc "github.com/dzwiedz90/helpdesk-proto/common"
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-users/config"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
	"github.com/dzwiedz90/helpdesk-service-users/pkg/core"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
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
	emptyRequest = &pb.CreateUserRequest{
		User: &pb.User{},
	}
)

func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name                 string
		req                  *pb.CreateUserRequest
		resp                 *pb.CreateUserResponse
		expectsError         bool
		expectedErrorMessage string
	}{
		// {
		// 	name: "valid request",
		// 	req:  validRequest,
		// 	resp: &pb.CreateUserResponse{
		// 		Id: 0,
		// 	},
		// 	expectsError: false,
		// },
		{
			name:                 "empty request",
			req:                  emptyRequest,
			expectsError:         true,
			expectedErrorMessage: "validation error, failed to get username from the request",
		},
	}

	for _, tc := range testCases {
		cfg := &config.Config{}
		logger := logs.NewTestLoggers(cfg)
		infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		cfg.InfoLog = infoLog
		errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		cfg.ErrorLog = errorLog

		core := core.NewTestCore()
		s := Server{
			Core: &core,
			ServerConfig: &serviceconfig.ServerConfig{
				Logger: &logger,
			},
		}
		resp, err := s.CreateUser(context.Background(), tc.req)
		if !tc.expectsError {
			if err != nil {
				t.Fatalf("Failed to create user: %v", err)
			}

			if resp.GetId() != tc.resp.GetId() {
				t.Fatal("Wrong response")
			}
		} else if tc.expectedErrorMessage != err.Error() {
			t.Fatalf("Expected error %s but got %s", tc.expectedErrorMessage, err.Error())
		}
	}
}
