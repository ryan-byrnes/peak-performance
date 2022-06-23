package sign_up

import "hash"

type User struct {
	Username  string      `json: "username"`
	Email     string      `json: "email"`
	Password  hash.Hash64 `json: "password"`
	FirstName string      `json: "firstName"`
	LastName  string      `json: "lastName"`
}

func main() {

}
