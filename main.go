package main

import (
	"fmt"
	service "github.com/ziyifast/yi-github-actions/service"
)

func main() {
	fmt.Println("before...")
	fmt.Println("got a user: ", service.UserService.GetUserName())
}
