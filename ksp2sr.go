package main

import (
	"log"

	// project imports
	menu "github.com/rDybing/ksp2scirep/menu"
)

func main() {
	if err := menu.MainMenu(); err != nil {
		log.Fatalf("Oops: %v\n", err)
	}
}
