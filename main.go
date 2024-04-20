package main

import (
	"fmt"
	"myTest/demo_home/github_actions_demo/service"
)

func main() {
	fmt.Println("got a user: ", service.UserService.GetUserName())
}
