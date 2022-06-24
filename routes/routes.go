package routes

import (
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/sign-up", SignUpHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler)
}
