package lookups

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Lookup(word string, resource string) ([]byte, int) {
	var url string
	if resource == "t" {
		url = makeThesaurusUrl(word)
	} else {
		url = makeDictionaryUrl(word)
	}

	res, err := http.Get(url)
	printErrorAndExit(err)
	resBody, err := ioutil.ReadAll(res.Body)
	printErrorAndExit(err)
	return resBody, res.StatusCode
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

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
