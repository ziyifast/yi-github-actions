package service

type userService struct {
}

var UserService = new(userService)

func (u *userService) GetUserName() string {
	return "tom"
}
