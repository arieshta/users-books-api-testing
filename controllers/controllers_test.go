package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"users-books-api-testing/config"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetUsersControllers(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			// assert.Equal(t, userJSON, rec.Body.String())
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}