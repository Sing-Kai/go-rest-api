package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Sing-Kai/go-rest-api/internal/transport/http"
)

// App - struct which contains things like pointers to database connections
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
