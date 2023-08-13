package model

type CreateUser struct {
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Age       int32
	Gender    string
	Address   *Address
}

type UpdateUser struct {
	UserId    int64
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Age       int32
	Gender    string
	Address   *Address
}
