package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(("/"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello go</h1>")
	})

	http.ListenAndServe(":3000", nil)
}
