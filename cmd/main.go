package main

import "basic-gin/di"

func main() {
	app := di.NewApp()

	app.Run("localhost:8000") // listen and serve on 0.0.0.0:8080
}
