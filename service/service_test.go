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
	validCreateRequest = &pb.CreateUserRequest{
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
	emptyCreateRequest = &pb.CreateUserRequest{
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
		{
			name: "valid request",
			req:  validCreateRequest,
			resp: &pb.CreateUserResponse{
				Id: 0,
			},
			expectsError: false,
		},
		{
			name:                 "empty request",
			req:                  emptyCreateRequest,
			expectsError:         true,
			expectedErrorMessage: "validation error, failed to get username from the request",
		},
		{
			name: "incomplete request",
			req: &pb.CreateUserRequest{
				User: &pb.User{
					Username:  "james.kirk",
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
			},
			expectsError:         true,
			expectedErrorMessage: "validation error, failed to get password from the request",
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

func TestGetUser(t *testing.T) {
	var testCases = []struct {
		name                 string
		req                  *pb.GetUserRequest
		resp                 *pb.GetUserResponse
		expectsError         bool
		expectedErrorMessage string
	}{
		// {
		// 	name: "valid request",
		// 	req: &pb.GetUserRequest{
		// 		Id: 1,
		// 	},
		// 	resp: &pb.GetUserResponse{
		// 		User: &pb.User{
		// 			UserId:    1,
		// 			Username:  "james.kirk",
		// 			Email:     "james.kirk@enterprise.com",
		// 			FirstName: "James",
		// 			LastName:  "Kirk",
		// 			Age:       32,
		// 			Gender:    "male",
		// 			Address: &pbloc.Address{
		// 				Street:     "Kirksuckz 8",
		// 				City:       "Deep Space Nine",
		// 				PostalCode: "2137CU666",
		// 				Country:    "Federation",
		// 			},
		// 		},
		// 	},
		// 	expectsError: false,
		// },
		{
			name:                 "invalid request",
			req:                  &pb.GetUserRequest{},
			resp:                 &pb.GetUserResponse{},
			expectsError:         true,
			expectedErrorMessage: "sql: no rows in result set",
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
		resp, err := s.GetUser(context.Background(), tc.req)
		if !tc.expectsError {
			if err != nil {
				t.Fatalf("Failed to get user: %v", err)
			}

			if resp.GetUser().GetUserId() != tc.resp.GetUser().GetUserId() {
				t.Fatalf("Wrong user id, expected %v but got %v", tc.resp.GetUser().GetUserId(), resp.GetUser().GetUserId())
			}
			if resp.GetUser().GetUsername() != tc.resp.GetUser().GetUsername() {
				t.Fatalf("Wrong username, expected %v but got %v", tc.resp.GetUser().GetUsername(), resp.GetUser().GetUsername())
			}
			if resp.GetUser().GetEmail() != tc.resp.GetUser().GetEmail() {
				t.Fatalf("Expected %v but got %v", tc.resp.GetUser().GetEmail(), resp.GetUser().GetEmail())
			}
			if resp.GetUser().GetFirstName() != tc.resp.GetUser().GetFirstName() {
				t.Fatalf("Wrong first name, expected %v but got %v", tc.resp.GetUser().GetFirstName(), resp.GetUser().GetFirstName())
			}
			if resp.GetUser().GetLastName() != tc.resp.GetUser().GetLastName() {
				t.Fatalf("Wrong last name, expected %v but got %v", tc.resp.GetUser().GetLastName(), resp.GetUser().GetLastName())
			}
			if resp.GetUser().GetAge() != tc.resp.GetUser().GetAge() {
				t.Fatalf("Wrong age, expected %v but got %v", tc.resp.GetUser().GetAge(), resp.GetUser().GetAge())
			}
			if resp.GetUser().GetGender() != tc.resp.GetUser().GetGender() {
				t.Fatalf("Wrong gender, expected %v but got %v", tc.resp.GetUser().GetGender(), resp.GetUser().GetGender())
			}
			if resp.GetUser().GetAddress().GetStreet() != tc.resp.GetUser().GetAddress().GetStreet() {
				t.Fatalf("Wrong street, expected %v but got %v", tc.resp.GetUser().GetAddress().GetStreet(), resp.GetUser().GetAddress().GetStreet())
			}
			if resp.GetUser().GetAddress().GetCity() != tc.resp.GetUser().GetAddress().GetCity() {
				t.Fatalf("Wrong city, expected %v but got %v", tc.resp.GetUser().GetAddress().GetCity(), resp.GetUser().GetAddress().GetCity())
			}
			if resp.GetUser().GetAddress().GetPostalCode() != tc.resp.GetUser().GetAddress().GetPostalCode() {
				t.Fatalf("Wrong postal code, expected %v but got %v", tc.resp.GetUser().GetAddress().GetPostalCode(), resp.GetUser().GetAddress().GetPostalCode())
			}
			if resp.GetUser().GetAddress().GetCountry() != tc.resp.GetUser().GetAddress().GetCountry() {
				t.Fatalf("Wrong country, expected %v but got %v", tc.resp.GetUser().GetAddress().GetCountry(), resp.GetUser().GetAddress().GetCountry())
			}
		} else if tc.expectedErrorMessage != err.Error() {
			t.Fatalf("Expected error %s but got %s", tc.expectedErrorMessage, err.Error())
		}
	}
}
