package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.New("login").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Login</title>
</head>
<body>
	<h2>Login Page</h2>
	<form method="POST" action="/login">
		<label>Username:</label>
		<input type="text" name="username" required><br>
		<label>Password:</label>
		<input type="password" name="password" required><br>
		<input type="submit" value="Login">
	</form>
</body>
</html>
`))

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		fmt.Fprintln(w, "Login Successful")
	} else {
		fmt.Fprintln(w, "Invalid Credentials")
	}
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
