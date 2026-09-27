package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSwaggerDocumentIsServed(t *testing.T) {
	app := &application{}
	req := httptest.NewRequest(http.MethodGet, "/swagger/doc.json", nil)
	res := httptest.NewRecorder()

	app.mount().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	if !strings.Contains(res.Body.String(), `"basePath": "/v1"`) {
		t.Fatal("Swagger document does not describe the /v1 API")
	}
}
