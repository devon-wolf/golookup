package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	lookupUrl := makeDictionaryUrl("purple")
	res, err := http.Get(lookupUrl)
	printErrorAndExit(err)
	resBody, bodyErr := ioutil.ReadAll(res.Body)
	printErrorAndExit(bodyErr)
	fmt.Printf("%s", resBody)
}

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}

func makeDictionaryUrl(word string) string {
	key := os.Getenv("DICTIONARY_KEY")
	dictionaryUrl := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word, key)
	return dictionaryUrl
}

func makeThesaurusUrl(word string) string {
	key := os.Getenv("THESAURUS_KEY")
	thesaurusUrl := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/thesaurus/json/%s?key=%s", word, key)
	return thesaurusUrl
}