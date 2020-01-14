package main

import (
	"fmt"
)

// Passed in during compiling
var version = "undefined"

// Passed in during compiling
var commit = "undefined"

func main() {
	if version == "undefined" {
		fmt.Println(version)
		panic(fmt.Sprintf("version is not defined"))
	}
	if commit == "undefined" {
		fmt.Println(commit)
		panic(fmt.Sprintf("commit is not defined"))
	}
	api := Handler{}
	api.Initialize()
	api.Run(":8000")
}
