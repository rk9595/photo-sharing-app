package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		// u.Templates.New.Execute(w, nil)
	}
	// fmt.Printf("Email: %s, Password: %s\n", email, password)
	// u.Templates.New.Execute(w, nil)
}
