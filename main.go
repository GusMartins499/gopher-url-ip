package main

import (
	"gopher-url-ip/app"
	"log"
	"os"
)

func main() {
	application := app.Execute()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
