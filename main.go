package main

import (
	"fmt"
	"os/user"
)

// this is a comment

func main() {
	usr := *user.User
	fmt.Println("Hello World " + user.Current(usr, ""))
}
