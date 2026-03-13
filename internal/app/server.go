package app

import (
	"fmt"
	"net/http"
)

func StartServer() {
	initRouter()
	fmt.Println("Server is running on http://localhost:9091")
	http.ListenAndServe(":9091", nil)
}
