package hello

import (
	"encoding/json"
	"fmt"
	"html"
	"httptest/users"
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
	r.HandleFunc("/", greetHandler)
	r.HandleFunc("/version", versionHandler)
	r.HandleFunc("/users", usersHandler)
	return r
}
func versionHandler(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	uu, err := users.Get()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	json.NewEncoder(w).Encode(uu)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "hello, %s!\n", html.EscapeString(name))
}
