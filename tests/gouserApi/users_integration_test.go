package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/spudmashmedia/gouser/internal/api"
	"github.com/spudmashmedia/gouser/internal/users"
	"github.com/spudmashmedia/gouser/pkg/randomuser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration Test Entrypoint
func TestGetUserEndpoint(t *testing.T) {
	// Arrange: Init Router
	router := buildUserRouter()

	// Arrange: Build Test Server
	ts := httptest.NewServer(router)
	defer ts.Close()

	t.Run(
		"Should return HTTP 404 record When incorrect route /user/unexpected is called",
		usersShouldReturnNotFoundWhenRouteIsIncorrect(ts.URL))

	t.Run(
		"Should return 1 User record When No Query Parameters",
		usersShouldReturnOneRecordWhenNoQueryParams(ts.URL))

	t.Run(
		"Should return 1 User record When QueryParm Count=0",
		usersShouldReturnSingleRecordWhenQueryParamIsZero(ts.URL))

	t.Run(
		"Should return 5 User records When QueryParam Count is 5",
		usersShouldReturnFiveRecordsWhenQueryParamIsFive(ts.URL))

	t.Run(
		"Should return maximum 5000 User records When QueryParam Count is 5001",
		usersShouldReturnMaxRecordsWhenQueryParamIs5001(ts.URL))
}

func buildUserRouter() *chi.Mux {

	r := chi.NewRouter()

	host := "https://randomuser.me"
	route := "/api"

	svc := users.NewService(
		randomuser.NewService(
			host,
			route,
		),
	)
	api.RegisterUserRouter(r, svc)

	return r
}

func usersShouldReturnNotFoundWhenRouteIsIncorrect(sutBaseUrl string) func(t *testing.T) {
	return func(t *testing.T) {
		// Arange: query param setup
		testUrl := fmt.Sprintf("%s/user/unexpected", sutBaseUrl)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should return a HTTP 404 Not Found status code")
	}
}

func usersShouldReturnOneRecordWhenNoQueryParams(sutBaseUrl string) func(t *testing.T) {
	return func(t *testing.T) {
		// Arange: query param setup
		expectedCount := 1

		testUrl := fmt.Sprintf("%s/user", sutBaseUrl)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body) // this is a []byte
		assert.NotEmpty(t, body, "Response Body should not be empty")

		// // Debug: uncomment this to view Raw Response body
		// t.Logf("Raw response body: %s", string(body))

		// assert: cast back to a UserResponse  test the count
		var actualResponse users.UsersResponse

		err = json.Unmarshal(body, &actualResponse)

		require.NoError(t, err, fmt.Sprintf("json.Unmarshal should not fail: %s", err))

		assert.NotEmpty(t, actualResponse, "Actual Response should not be empty")
		assert.Equal(t, expectedCount, len(actualResponse.Results), fmt.Sprintf("Results should equal %d", expectedCount))
	}
}

func usersShouldReturnSingleRecordWhenQueryParamIsZero(sutBaseUrl string) func(t *testing.T) {

	return func(t *testing.T) {
		// Arange: query param setup
		expectedCount := 1

		testUrl := fmt.Sprintf("%s/user?count=0", sutBaseUrl)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body) // this is a []byte
		assert.NotEmpty(t, body, "Response Body should not be empty")

		// // Debug: uncomment this to view Raw Response body
		// t.Logf("Raw response body: %s", string(body))

		// assert: cast back to a UserResponse  test the count
		var actualResponse users.UsersResponse

		err = json.Unmarshal(body, &actualResponse)

		require.NoError(t, err, fmt.Sprintf("json.Unmarshal should not fail: %s", err))

		assert.NotEmpty(t, actualResponse, "Actual Response should not be empty")
		assert.Equal(t, expectedCount, len(actualResponse.Results), fmt.Sprintf("Results should equal %d", expectedCount))
	}
}

func usersShouldReturnFiveRecordsWhenQueryParamIsFive(sutBaseUrl string) func(t *testing.T) {
	return func(t *testing.T) {

		// Arange: query param setup
		testCount := 5
		expectedCount := 5

		testUrl := fmt.Sprintf("%s/user?count=%d", sutBaseUrl, testCount)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body) // this is a []byte
		assert.NotEmpty(t, body, "Response Body should not be empty")

		// // Debug: uncomment this to view Raw Response body
		// t.Logf("Raw response body: %s", string(body))

		// assert: cast back to a UserResponse  test the count
		var actualResponse users.UsersResponse

		err = json.Unmarshal(body, &actualResponse)

		require.NoError(t, err, fmt.Sprintf("json.Unmarshal should not fail: %s", err))

		assert.NotEmpty(t, actualResponse, "Actual Response should not be empty")
		assert.Equal(t, expectedCount, len(actualResponse.Results), fmt.Sprintf("Results should equal %d", expectedCount))
	}
}

func usersShouldReturnMaxRecordsWhenQueryParamIs5001(sutBaseUrl string) func(t *testing.T) {
	return func(t *testing.T) {

		// Arange: query param setup
		testCount := 5001
		expectedCount := 5000

		testUrl := fmt.Sprintf("%s/user?count=%d", sutBaseUrl, testCount)

		// Act
		t.Logf("Testing url: '%s'", testUrl)
		resp, err := http.Get(testUrl)

		// Assert
		require.NoError(t, err, fmt.Sprintf("%s should not fail", testUrl))
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body) // this is a []byte
		assert.NotEmpty(t, body, "Response Body should not be empty")

		// // Debug: uncomment this to view Raw Response body
		// t.Logf("Raw response body: %s", string(body))

		// assert: cast back to a UserResponse  test the count
		var actualResponse users.UsersResponse

		err = json.Unmarshal(body, &actualResponse)

		require.NoError(t, err, fmt.Sprintf("json.Unmarshal should not fail: %s", err))

		assert.NotEmpty(t, actualResponse, "Actual Response should not be empty")
		assert.Equal(t, expectedCount, len(actualResponse.Results), fmt.Sprintf("Results should equal %d", expectedCount))
	}
}
