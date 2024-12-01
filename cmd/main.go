package main

import (
	"one-pte-backend/app"
)

func main() {
	app := app.NewApplication()
	app.Init()
	app.Run()
	app.Wait()
	app.Stop()

}
