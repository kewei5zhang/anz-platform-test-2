package main

import (
	"os"
	"testing"
)

var api Handler

func TestMain(m *testing.M) {
	api = Handler{}
	api.Initialize()
	res := m.Run()
	os.Exit(res)
}
