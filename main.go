package main

import (
	"log"
	"github.com/PyMarcus/go_webscraping/script"
)

func main() {
	log.Println("Starting...")
	script.Run("quotes.toscrape.com")
	log.Println("OK!")
}
