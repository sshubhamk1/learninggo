package main

import (
	"fmt"

	churl "github.com/sshubhamk1/learninggo/chapter2"
)

func main() {
	fmt.Println("Starting lessions from here")
	url, _ := churl.Parse("https://github.com/sshubhamk1/")
	fmt.Println(url)
}
