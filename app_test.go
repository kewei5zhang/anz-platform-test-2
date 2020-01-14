package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testInitApplications(t *testing.T) {
	initApplications()
}

func TestGetAppVersion(t *testing.T) {
	req, _ := http.NewRequest("GET", "/version", nil)
	response := executeRequest(req)
	if http.StatusOK != response.Code {
		t.Errorf("Expected response: %d. Got: %d\n", http.StatusOK, response.Code)
	}
}

func TestHomeLink(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)
	if http.StatusOK != response.Code {
		t.Errorf("Expected response: %d. Got: %d\n", http.StatusOK, response.Code)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	api.Router.ServeHTTP(rr, req)
	return rr
}
