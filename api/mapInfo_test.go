package api

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapInfoHandler(t *testing.T) {
	mockResponse := `{"map_size":24,"rover_step":8}`

	router := setUpRouter()
	req, _ := http.NewRequest("GET", "/map-info", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	require.Equal(t, mockResponse, string(responseData))
	require.Equal(t, http.StatusOK, w.Code)

	req, _ = http.NewRequest("POST", "/map-info", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ = ioutil.ReadAll(w.Body)
	require.Equal(t, http.StatusNotFound, w.Code)
}
