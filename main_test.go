package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	cityError = "wrong city value"
)

// TestMainHandlerWhenCorrect проверяем условия корректности выполнения запроса
func TestMainHandlerWhenCorrect(t *testing.T) {
	t.Run("count=2&city=moscow", func(t *testing.T) {
		count := 2
		req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusOK, responseRecorder.Code)

		body := responseRecorder.Body.String()
		list := strings.Split(body, ",")

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, count, len(list))
		}
	})

	t.Run("count=4&city=moscow", func(t *testing.T) {
		count := 4
		req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusOK, responseRecorder.Code)

		body := responseRecorder.Body.String()
		list := strings.Split(body, ",")

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, count, len(list))
		}
	})

}

// TestMainHandlerWhenCountMoreThanTotal проверяем выполнение при count > 4
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	t.Run("count=10&city=moscow", func(t *testing.T) {
		count := 4
		req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusOK, responseRecorder.Code)

		body := responseRecorder.Body.String()
		list := strings.Split(body, ",")

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, count, len(list))
		}

	})

	t.Run("count=5&city=moscow", func(t *testing.T) {
		count := 4
		req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusOK, responseRecorder.Code)

		body := responseRecorder.Body.String()
		list := strings.Split(body, ",")

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, count, len(list))
		}

	})

}

// TestMainHandlerWhenSityIncorrect проверяем выполнение при некорректо переданном городе
func TestMainHandlerWhenSityIncorrect(t *testing.T) {
	t.Run("count=5&city=", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/cafe?count=5&city=", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		body := responseRecorder.Body.String()

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, cityError, body)
		}

	})

	t.Run("count=5&city=saratov", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/cafe?count=5&city=", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		body := responseRecorder.Body.String()

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, cityError, body)
		}

	})

	t.Run("count=5", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/cafe?count=5", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		body := responseRecorder.Body.String()

		if assert.NotEmpty(t, len(body)) {
			assert.Equal(t, cityError, body)
		}

	})
}
