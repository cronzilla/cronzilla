package main

import "net/http"

func ListOfServersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("List of Servers"))
}
