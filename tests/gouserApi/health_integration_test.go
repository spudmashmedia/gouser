package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/spudmashmedia/gouser/internal/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildHealthRouter() *chi.Mux {

	r := chi.NewRouter()
	api.RegisterHealthRouter(r)

	return r
}

func TestGetHealthEndpoint(t *testing.T) {
	// Arrange: Init Router
	router := buildHealthRouter()

	// Arrange: Build Test Server
	ts := httptest.NewServer(router)
	defer ts.Close()

	t.Run("Should return HTTP 200 record when route /health is called", func(t *testing.T) {
		// Arange: query param setup
		testUrl := fmt.Sprintf("%s/health", ts.URL)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "Should return a HTTP 200 status code")
	})

	t.Run("Should return HTTP 404 when route /health/unexpected is called", func(t *testing.T) {
		// Arange: query param setup
		testUrl := fmt.Sprintf("%s/health/unexpected", ts.URL)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should return a HTTP 404 status code")
	})
}
