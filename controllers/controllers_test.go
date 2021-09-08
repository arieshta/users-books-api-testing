package controllers

import (
	"encoding/json"
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
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))
		// Assertion
		if assert.NoError(t, GetUserByIdController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
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
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
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
			id:                   43,
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
			id:                   36,
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

func TestGetBooksControllers(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains string
	}{
		{
			testName:             "success",
			path:                 "/books",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":[",
			expectBodyContains: "success",
		},
	}

	e := InitEcho()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetBookByIdController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		id                   int
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			testName:             "un-success (not found - deleted)",
			path:                 "/books/",
			id:                   2,
			expectStatus:         http.StatusNotFound,
			expectBodyStartsWith: "{\"message",
		},
		{
			testName:             "success",
			path:                 "/books/",
			id:                   6,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"book\":{\"ID\":6",
		},
	}

	e := InitEcho()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))
		// Assertion
		if assert.NoError(t, GetBookByIdController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.Contains(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestAddBookController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		title                string
		author               string
		year                 int
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1   string
		expectBodyContains2 string
	}{
		{
			testName:             "success",
			path:                 "/books",
			title:                "iron",
			author:               "m@rvel",
			year:                 2019,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"book\":{",
			expectBodyContains1:   "success",
			expectBodyContains2: "iron",
		},
	}

	e := InitEcho()

	for _, testCase := range testCases {
		user := map[string]interface{}{
			"title":  testCase.title,
			"author": testCase.author,
			"year":   testCase.year,
		}
		data, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, AddBookController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestUpdateBookByIdController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		id                   int
		title                string
		author               string
		year                 int
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1   string
		expectBodyContains2 string
	}{
		{
			testName:             "success",
			path:                 "/books/",
			id:                   1,
			title:                "setrika",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"book\":{",
			expectBodyContains1:   "success",
			expectBodyContains2: "setrika",
		},
	}

	e := InitEcho()

	for _, testCase := range testCases {
		user := map[string]interface{}{
			"title":  testCase.title,
			"author": testCase.author,
			"year":   testCase.year,
		}
		data, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, UpdateBookByIdController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestDeleteBookByIdController(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		id                   int
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			testName:             "success",
			path:                 "/books",
			id:                   4,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"message\":\"success delete book",
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

		if assert.NoError(t, DeleteBookByIdController(c)) {
			assert.Equal(t, rec.Code, testCase.expectStatus)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
