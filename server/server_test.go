package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var RequestTests = []RequestTest{
	{Method: "GET", Path: "/characters/test", Body: nil, Code: http.StatusOK},
	{Method: "POST", Path: "/newCommit", Body: strings.NewReader("{\"userEmail\":\"test\",\"experience\":\"123\"}"), Code: http.StatusCreated},
}

func TestRoutes(t *testing.T) {
	s := StartServer()
	for _, test := range RequestTests {
		r := httptest.NewRequest(test.Method, test.Path, test.Body)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, r)

		if w.Code != test.Code {
			t.Errorf("unexpected status code when performing a %s request to %s. Expected %d, Received %d", test.Method, test.Path, test.Code, w.Code)
		}
	}
}
