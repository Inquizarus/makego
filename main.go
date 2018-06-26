package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s PROJECT_NAME\n", os.Args[0])
		os.Exit(1)
	}

	response, _ := http.Get("https://gist.githubusercontent.com/Inquizarus/82793bed195c2df6943d303c8f663300/raw/d303034151e5d33caa6bca238f2f2aab1bc4fb59/golang-makefile")

	content, _ := ioutil.ReadAll(response.Body)
	old := []byte("{{binary_name}}")
	new := []byte(os.Args[1])
	newContent := bytes.Replace(content, old, new, 1)

	ioutil.WriteFile("makefile", newContent, 0777)
}
