package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/tech-hunters/handlers"
)

func main() {
	var router = handlers.RunServer()

	var err = router.Run("localhost:6969")

	if err != nil {
		fmt.Println("server spwan Error")
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
