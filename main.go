package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func getPort() string {
	port := os.Getenv("PORT_WWW")
	if len(port) < 1 {
		port = "5000"
	}
	return port
}

func getStaticDir() string {
	staticPath := os.Getenv("STATIC_DIR")
	if len(staticPath) < 1 {
		staticPath = "static/"
	}
	return staticPath
}

func serveStatic(router *mux.Router) {
	handler := func(w http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		filepath := "/" + vars["path"]
		fmt.Println("handeling static call for: ", filepath, "( ", request.Method, " )")
		fmt.Println("returns the file at", path.Join(getStaticDir(), filepath))
		http.ServeFile(w, request, path.Join(getStaticDir(), filepath))
	}
	router.HandleFunc("/{path:.*}", handler)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/go", ServeHTTP)
	serveStatic(r.PathPrefix("/").Subrouter())

	http.ListenAndServe("0.0.0.0:"+getPort(), r)
}
