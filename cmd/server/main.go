package main

import(
	"fmt"
)

// App - struct which contains things like pointers to database connections
type App struct{

}

func (app *App) Run() error {
	fmt.Println("Setting up App")
	return nil
}

func main(){
	fmt.Println("go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}

}