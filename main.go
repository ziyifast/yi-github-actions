package main

import (
	"fmt"
	"ziyi.com/service"
)

func main() {
	fmt.Println("before...")
	fmt.Println("got a user: ", service.UserService.GetUserName())
}
