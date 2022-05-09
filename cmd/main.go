package main

import (
	"net/http"
)

var UrlHello = []byte(`
<html>
	<body>
		<form action = "/" method = "POST">
			Enter Your Name: <input type="text" name="userName">
			<input type="submit" value="ENTER">
		</form>
	</body>
	<body>
		
	</body>
`)

func handle228(w http.ResponseWriter, r *http.Request) {
	w.Write(UrlHello)
	db()
	name := r.FormValue("userName")
	if name != "" {
		w.Write([]byte("Hello " + name))
	}
}

func main() {
	http.HandleFunc("/", handle228)

	http.ListenAndServe(":8080", nil)
}
