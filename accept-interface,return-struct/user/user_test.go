package user

import "testing"

func TestCreateUser(t *testing.T) {
	s := new(inMemStore)
	service := NewUserService(s)

	//...test the CreateUser() function
}
