package api

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoverStatusHandler(t *testing.T) {

	returnDataStepZero := "{\"status\":{\"direction\":\"N\",\"position_x\":0,\"position_y\":0},\"current_step\":0,\"status_text_format\":\"N:0,0\"}"
	router := setUpRouter()
	req, _ := http.NewRequest("GET", "/rover-status?step=0", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	fmt.Println(string(responseData))
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, string(responseData), returnDataStepZero)

	req, _ = http.NewRequest("GET", "/rover-status?step=-1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)

	req, _ = http.NewRequest("GET", "/rover-status?step=999", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)

	req, _ = http.NewRequest("POST", "/rover-status", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ = ioutil.ReadAll(w.Body)
	require.Equal(t, http.StatusNotFound, w.Code)
}
