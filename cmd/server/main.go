package main

import (
	"fmt"
	"net/http"

	"github.com/Sing-Kai/go-rest-api/internal/comment"
	"github.com/Sing-Kai/go-rest-api/internal/database"
	transportHTTP "github.com/Sing-Kai/go-rest-api/internal/transport/http"
)

// App - struct which contains things like pointers to database connections
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up App")

	fmt.Println("Creating new database")
	var err error
	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	fmt.Println("Database migration")
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	fmt.Println("Initialising comment service")
	commentService := comment.NewService(db)

	fmt.Println("Setting up route handlers")
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}

/*
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
*/
