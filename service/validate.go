package service

import (
	"errors"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
)

func ValidateRequest(req interface{}) error {
	switch r := req.(type) {
	case *pb.CreateUserRequest:
		return validateCreateUserRequest(r)
	case *pb.GetUserRequest:
		validateGetUserRequest(r)
	case *pb.UpdateUserRequest:
		validateUpdateUserRequest(r)
	}

	return nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) error {
	u := req.GetUser()

	if u.GetUsername() == "" {
		return errors.New("validation error, failed to get username from the request")
	} else if u.GetPassword() == "" {
		return errors.New("validation error, failed to get password from the request")
	} else if u.GetEmail() == "" {
		return errors.New("validation error, failed to get email from the request")
	} else if u.GetFirstName() == "" {
		return errors.New("validation error, failed to get first name from the request")
	} else if u.GetLastName() == "" {
		return errors.New("validation error, failed to get last name from the request")
	} else if u.GetAge() == 0 {
		return errors.New("validation error, failed to get age from the request")
	} else if u.GetGender() == "" {
		return errors.New("validation error, failed to get gender from the request")
	} else if u.GetAddress().Street == "" {
		return errors.New("validation error, failed to get street from the request")
	} else if u.GetAddress().City == "" {
		return errors.New("validation error, failed to get city from the request")
	} else if u.GetAddress().PostalCode == "" {
		return errors.New("validation error, failed to get postal code from the request")
	} else if u.GetAddress().Country == "" {
		return errors.New("validation error, failed to get country from the request")
	}

	return nil
}

func validateGetUserRequest(req *pb.GetUserRequest) error {
	id := req.GetId()

	if id == 0 {
		return errors.New("validation error, failed to get id from the request")
	}

	return nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) error {
	u := req.GetUser()

	if u.GetEmail() == "" {
		return errors.New("validation error, failed to get email from the request")
	} else if u.GetFirstName() == "" {
		return errors.New("validation error, failed to get first name from the request")
	} else if u.GetLastName() == "" {
		return errors.New("validation error, failed to get last name from the request")
	} else if u.GetAge() == 0 {
		return errors.New("validation error, failed to get age from the request")
	} else if u.GetGender() == "" {
		return errors.New("validation error, failed to get gender from the request")
	} else if u.GetAddress().Street == "" {
		return errors.New("validation error, failed to get street from the request")
	} else if u.GetAddress().City == "" {
		return errors.New("validation error, failed to get city from the request")
	} else if u.GetAddress().PostalCode == "" {
		return errors.New("validation error, failed to get postal code from the request")
	} else if u.GetAddress().Country == "" {
		return errors.New("validation error, failed to get country from the request")
	}

	return nil
}
