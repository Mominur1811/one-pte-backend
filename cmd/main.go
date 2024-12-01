package main

import (
	"fmt"
	"one-pte-backend/app"
)

func main() {
	app := app.NewApplication()
	fmt.Println("Hello World")
	app.Init()
	app.Run()
	app.Wait()
	app.Stop()

}
