package main

import (
	"fmt"

	"github.com/dj5harma/podcast/auth"
	"github.com/dj5harma/podcast/user"
)

func main() {
	auth.LoginWithCredentials("Sunu", "123456")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{Name: "Sunu", Email: "sunu@email.com"}
	fmt.Println(user)
}
