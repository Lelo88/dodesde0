package users

import (
	"fmt"
	"time"

	"github.com/Lelo88/dodesde0/modelos"
)

func AltaUsuario(){
	user := new(modelos.User)
	user.AddUser(1, "Jorge", time.Now(), true)

	fmt.Println(user)
}