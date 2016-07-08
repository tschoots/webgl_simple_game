package main 

import (
	"net/http"

)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./webGL")))
	http.ListenAndServe(":8080", nil)

}

