package model

type User struct {
	UserId    int64
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Age       int32
	Gender    string
	Address   Address
}

func (u *User) GetId() int64 {
	if u != nil {
		return u.UserId
	}

	return 0
}

func (u *User) GetUsername() string {
	if u != nil {
		return u.Username
	}

	return ""
}

func (u *User) GetEmail() string {
	if u != nil {
		return u.Email
	}

	return ""
}

func (u *User) GetFirstName() string {
	if u != nil {
		return u.FirstName
	}

	return ""
}

func (u *User) GetLastName() string {
	if u != nil {
		return u.LastName
	}

	return ""
}

func (u *User) GetAge() int32 {
	if u != nil {
		return u.Age
	}

	return 0
}

func (u *User) GetGender() string {
	if u != nil {
		return u.Gender
	}

	return ""
}

func (u *User) GetAddress() *Address {
	if u != nil {
		return &u.Address
	}

	return nil
}
