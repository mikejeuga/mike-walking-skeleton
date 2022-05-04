package server_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/mike-walking-skeleton/src/adapters/web/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	is := is.New(t)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	walkingServer := server.NewServer()
	walkingServer.Handler.ServeHTTP(recorder, request)

	is.Equal(recorder.Code, http.StatusOK)
}
