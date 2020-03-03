package main

import (
	"SignUpLogin-Checkpoint/validation"
	"fmt"
	"html/template"
	"os"
	"net/http"

)


type User struct {
	UserName string
	Password string
}



func main() {
	rootDir, _ :=os.Getwd()
	mux := http.NewServeMux()

	// HOME
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl,err := template.ParseFiles("./template/home.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w,nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	// LOGIN
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl,err := template.ParseFiles("./template/login.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w,nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		r.ParseForm()
		user:= User {r.FormValue("username"),r.FormValue("password")}

		// Validate data
		validName := validation.IsValid(user.UserName)
		validPwd := validation.IsValid(user.Password)

		if validName || validPwd {
			fmt.Fprintf(w,"Input value is empty")
			return
		}

		usernameExemple := "youyou"
		passwordExemple := "12345"

		if user.UserName == usernameExemple && user.Password == passwordExemple {
			fmt.Fprintf(w,"%s Login succesful!",usernameExemple)
		}else{
			fmt.Fprintf(w, "Login Failed!!")
		}

	})

	// SIGN UP

	fileserver := http.FileServer(http.Dir(rootDir + "/assets"))
	mux.Handle("/static/",http.StripPrefix("/static/",fileserver))

http.ListenAndServe("localhost:555",mux)
}
