package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://magnus-archive.herokuapp.com/episodes")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	resBody, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		fmt.Printf("%s", bodyErr)
		os.Exit(1)
	}
	fmt.Printf("%s", resBody)
}