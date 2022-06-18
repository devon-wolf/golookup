package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/devon-wolf/golookup/lookups"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	word := flag.String("word", "", "The word you want to look up")
	resource := flag.String("src", "d", "What resource you want to use, 'd' for dictionary, 't' for thesaurus.")
	flag.Parse()

	resBody, statusCode := lookups.Lookup(*word, *resource)
	if statusCode != 200 {
		fmt.Printf("Encountered an error with that lookup (status %d), sorry!", statusCode)
		os.Exit(0)
	}
	fmt.Printf("%s", resBody)
}
