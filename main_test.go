package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	expectedStatus := http.StatusOK
	body := responseRecorder.Body

	require.Equal(t, expectedStatus, status)
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=warsaw", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	actualStatus := responseRecorder.Code
	expectedStatus := http.StatusBadRequest
	require.Equal(t, expectedStatus, actualStatus)

	actualAnswer := responseRecorder.Body.String()
	expectedAnswer := "wrong city value"
	require.Equal(t, expectedAnswer, actualAnswer)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	param := "/cafe?count=" + strconv.Itoa(totalCount+5) + "&city=moscow"
	req := httptest.NewRequest("GET", param, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	actualStatus := responseRecorder.Code
	expectedStatus := http.StatusOK
	require.Equal(t, expectedStatus, actualStatus)

	body := responseRecorder.Body
	citiesList := strings.Split(body.String(), ",")
	assert.Len(t, citiesList, totalCount)
}
