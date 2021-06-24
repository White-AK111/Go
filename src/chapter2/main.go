package main

import (
	"fmt"
	"os"
	"os/user"
)

// this is a comment

func main() {
	//usr := *user.User
	//fmt.Println("Hello World " + user.Current(usr, ""))

	fmt.Println(user.Current()) //return current username
	fmt.Println(os.Hostname())  //return the hostname(domain)

	fmt.Println("Hello World is White")
}
