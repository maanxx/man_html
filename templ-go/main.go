package main

import (
	"fmt"
	"net/http"
	"templ-go/man"

	"github.com/a-h/templ"
)

func main() {
	component := man.Hello("Man Minh")
	http.Handle("/d", templ.Handler(component))

	fmt.Println("Listening on :5000")
	http.ListenAndServe(":5000", nil)

}