package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

/*	If this was a "package main" file, then we could do the below and have a
	local webserver up-n-going at localhost:5000 serving "Hello, world"
*/
// func main() {
// 	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
// }
