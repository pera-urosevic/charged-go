package main

import (
	"charged/config"
	"charged/ui"
	"fmt"
	"log"
)

func main() {
	var err error

	fmt.Println("Charged")

	err = config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = ui.Init()
	if err != nil {
		log.Fatal(err)
	}
}
