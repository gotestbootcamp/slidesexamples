package hello

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gorilla/mux"
)

func serve(addr string) {
	r := newRouter()
	http.Handle("/", r)

	log.Printf("serving http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/version", version)
	return r
}
func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "hello, %s!\n", html.EscapeString(name))
}
