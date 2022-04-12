package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetDistance(t *testing.T){

	router := httprouter.New()
	router.POST("/v1/get_distance", GetDistance)

	req, _ := http.NewRequest("POST", "/v1/get_distance", strings.NewReader("from=Lausanne&to=Genève"))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}