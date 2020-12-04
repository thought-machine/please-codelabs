// Package main implements a hello world http service
package main

import "net/http"

// Server implements the hello world HTTP server
type Server struct {
}

// ServeHTTP just responds with Hello, world! on the writer
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("Hello world!\n")); err != nil {
		panic(err)
	}
}

func main() {
	if err := http.ListenAndServe(":8080", new(Server)); err != nil {
		panic(err)
	}
}
