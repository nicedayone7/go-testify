package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler:= http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.NotEmpty(t, responseRecorder.Body)
}


func TestMainHandlerWhenCityNotSupport(t *testing.T) {
    bodyResponse := `wrong city value`

    req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=Tula", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
    assert.Equal(t, bodyResponse, responseRecorder.Body.String())
}


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    
    req := httptest.NewRequest(http.MethodGet, "/cafe?count=9&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.Equal(t, totalCount, len(strings.Split(responseRecorder.Body.String(), ",")))
}