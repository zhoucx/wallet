package app

import (
	"fmt"
	"net/http"
)

func StartServer(port int) {
	initRouter()
	fmt.Printf("Server is running on http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
