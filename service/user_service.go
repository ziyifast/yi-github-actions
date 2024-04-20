package service

type userService struct {
}

// UserService get user name
var UserService = new(userService)

func (u *userService) GetUserName() string {
	return "tom"
}
