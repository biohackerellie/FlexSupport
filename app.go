package main

import (
	"fmt"
	"flexsupport/internal/handlers"
	"flexsupport/internal/router"
	"net/http"
)

func App() error {
	h := handlers.NewHandler()
	r := router.NewRouter(h)
	
	fmt.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", r)


}
