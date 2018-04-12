package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVictorHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	var jsonStr = []byte(`{"name":"Victor"}`)
	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VictorHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the response body is what we expect.
	expected := `{"respond":"Hello, Victor"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
