//go:generate go run main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Functions struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func main() {
	var functions []Functions
	var err error

	functionsFile, err := ioutil.ReadFile("example/functions.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(functionsFile, &functions)
	if err != nil {
		panic(err)
	}

	functionsTemplate, err := template.ParseFiles("example/functions.go.tpl")
	if err != nil {
		panic(err)
	}

	generatedFile, err := os.Create("example/generated.go")
	if err != nil {
		panic(err)
	}

	err = functionsTemplate.Execute(generatedFile, functions)
	if err != nil {
		panic(err)
	}

	fmt.Println("execute the generated file by `go run example/generated.go`")
}
