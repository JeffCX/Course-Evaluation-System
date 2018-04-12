// sessions.go
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
	"reflect"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)


type Users struct {
	username string 
	password string 
}

var users = Users{"jeffcx","heizhenzhu"}

func create_cookies(w http.ResponseWriter, r *http.Request, users Users){
	key = []byte(users.username+users.password)
	store = sessions.NewCookieStore(key)
	sessionss, _ := store.Get(r, users.username+users.password)
	sessionss.Values["authenticated"] = true
	sessionss.Save(r,w)
	fmt.Println(reflect.TypeOf(sessionss))
	
}

func get_cookies(users) *sessions.Session {
	return store.
}


func remove_cookies(w http.ResponseWriter, r *http.Request, users Users){
	sessionss, _ := store.Get(r, users.username+users.password)
	sessionss.Values["authenticated"] = false
	sessionss.Save(r,w)
	
}


func secret(w http.ResponseWriter, r *http.Request) {
	create_cookies(w,r,users)
	session, _ := store.Get(r, "name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}