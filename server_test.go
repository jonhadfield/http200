package http200

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGET(t *testing.T) {
	t.Run("check GET / returns ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Server(response, request)
		assert.Equal(t, "ok", response.Body.String())
	})
}

func TestPOST(t *testing.T) {
	t.Run("check POST / returns ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()

		Server(response, request)
		assert.Equal(t, "ok", response.Body.String())
	})
}

func TestHEAD(t *testing.T) {
	t.Run("check HEAD / returns ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodHead, "/", nil)
		response := httptest.NewRecorder()

		Server(response, request)
		assert.Equal(t, "ok", response.Body.String())
	})
}
