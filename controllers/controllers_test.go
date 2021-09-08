package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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
		testName             string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			testName:             "success",
			path:                 "/users",
			expectBodyStartsWith: "{\"message\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
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

func TestGetUserByIdController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		id                   int
		expectStatus         int
		expectBodyStartsWith string
		// error                string
	}{
		{
			testName:             "un-success (not found - no record)",
			path:                 "/users/",
			id:                   1,
			expectStatus:         http.StatusNotFound,
			expectBodyStartsWith: "\"message\":\"record not found\"",
		},
		{
			testName:             "un-success (not found - deleted)",
			path:                 "/users/",
			id:                   2,
			expectStatus:         http.StatusNotFound,
			expectBodyStartsWith: "{\"message",
		},
		{
			testName:             "success",
			path:                 "/users/",
			id:                   4,
			expectStatus:         http.StatusOK, 
			expectBodyStartsWith: "{\"message\":\"success\",\"user\":{\"ID\":4",
		},
	}

	e := InitEcho()
	
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		fmt.Println(testCase.id)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))
		// Assertion
		if assert.NoError(t, GetUserByIdController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(body)
			assert.True(t, strings.Contains(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		name                 string
		email                string
		password             string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains   string
	}{
		{
			testName:             "success",
			path:                 "/users",
			name:                 "iron",
			email:                "m@rvel",
			password:             "man",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"message\":\"success create",
			expectBodyContains:   "iron",
		},
	}

	e := InitEcho()

	for _, testCase := range testCases {
		user := map[string]string{
			"name":     testCase.name,
			"email":    testCase.email,
			"password": testCase.password,
		}
		data, _ := json.Marshal(user)
		fmt.Println(string(data))
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains))
		}
	}
}

func TestUpdateUserByIdController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		id                   int
		name                 string
		email                string
		password             string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains   string
	}{
		{
			testName:             "success",
			path:                 "/users/",
			id:                   35,
			name:                 "setrika",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"message\":\"success update",
			expectBodyContains:   "setrika",
		},
	}

	e := InitEcho()

	for _, testCase := range testCases {
		user := map[string]string{
			"name":     testCase.name,
			"email":    testCase.email,
			"password": testCase.password,
		}
		data, _ := json.Marshal(user)
		fmt.Println(string(data))
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, UpdateUserByIdController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains))
		}
	}
}

func TestDeleteUserByIdController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		id                   int
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			testName:             "success",
			path:                 "/users",
			id:                   35,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"message\":\"success delete",
		},
	}
	e := InitEcho()
	
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, DeleteUserByIdController(c)) {
			assert.Equal(t, rec.Code, testCase.expectStatus)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
