package main

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// func TestGetCompaniesHandler(t *testing.T) {
//     r := SetUpRouter()
//     r.GET("/companies", GetCompaniesHandler)
//     req, _ := http.NewRequest("GET", "/companies", nil)
//     w := httptest.NewRecorder()
//     r.ServeHTTP(w, req)

//     var companies []Company
//     json.Unmarshal(w.Body.Bytes(), &companies)

//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.NotEmpty(t, companies)
// }
