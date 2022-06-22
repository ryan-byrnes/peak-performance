package login

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r =: mux.newRouter()
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/sign-up", SignUpHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler)
}