package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMikrowaysDemoServer(t *testing.T) {

	t.Run("Any endpoint", func(t *testing.T) {
		request := newGetRequest("/")
		response := httptest.NewRecorder()

		MikrowaysDemoServer(response, request)

		assertResponseBody(t, response.Body.String(), "Hello World!!")

		request = newGetRequest("/any-other-word")
		response = httptest.NewRecorder()

		MikrowaysDemoServer(response, request)

		assertResponseBody(t, response.Body.String(), "Hello World!!")
	})

	t.Run("Health endpoint", func(t *testing.T) {
		request := newGetRequest("/healthz")
		response := httptest.NewRecorder()

		MikrowaysDemoServer(response, request)

		assertResponseBody(t, response.Body.String(), "Health OK!")
	})

	t.Run("Liveness endpoint", func(t *testing.T) {
		request := newGetRequest("/liveness")
		response := httptest.NewRecorder()

		MikrowaysDemoServer(response, request)

		assertResponseBody(t, response.Body.String(), "I am alive!!!")
	})

}

func newGetRequest(name string) * http.Request {
	req, _ := http.NewRequest(http.MethodGet, name, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
