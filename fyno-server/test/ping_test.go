package test

import (
	"fyno/server/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingHandler(t *testing.T) {
	// Create a new Gin router.
	router := gin.Default()

	// Register the Ping handler with the router.
	router.GET("/ping", handlers.Ping)

	// Create a new HTTP request.
	request, _ := http.NewRequest("GET", "/ping", nil)

	// Create a new HTTP response recorder.
	responseRecorder := httptest.NewRecorder()

	// Dispatch the request to the router.
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code.
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, responseRecorder.Code)
	}

	expected := `{"message":"pong"}`
	if responseRecorder.Body.String() != expected {
		t.Errorf("expected body %q but got %q", expected, responseRecorder.Body.String())
	}
}
