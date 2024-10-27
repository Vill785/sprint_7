package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCafeCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, body)
}
func TestWrongCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=3&city=japan", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Contains(t, body, "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверкиgo
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	if !assert.Len(t, list, totalCount) {
		t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))
	}

}
