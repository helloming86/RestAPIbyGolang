package test

import (
	"github.com/stretchr/testify/assert"
	"miMallDemo/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexGetRouter(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin", w.Body.String())
}
