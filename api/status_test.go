package api

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	mockResponse := `{"server_status":"ok"}`

	router := setUpRouter()
	req, _ := http.NewRequest("GET", "/status", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	require.Equal(t, mockResponse, string(responseData))
	require.Equal(t, http.StatusOK, w.Code)

	req, _ = http.NewRequest("POST", "/status", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ = ioutil.ReadAll(w.Body)
	require.Equal(t, http.StatusNotFound, w.Code)
}
