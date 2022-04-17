package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var cookie *sessions.CookieStore

func init() {
	cookie = sessions.NewCookieStore([]byte("Golang-Blogs"))
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := cookie.Get(r, "Golang-session")
	var authenticated interface{} = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "UnAuthorized to Access this Page.", http.StatusForbidden)
			return
		}
		fmt.Fprintf(w, "Authenticated User's Home Page")
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := cookie.Get(r, "Golang-session")
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintf(w, "Successfully Logged In")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := cookie.Get(r, "Golang-session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintf(w, "Successfully Logged Out")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
