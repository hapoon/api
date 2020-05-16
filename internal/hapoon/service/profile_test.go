package service

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProfilerList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	p := NewProfiler()

	if assert.NoError(t, p.List(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expect := `[{"id":"1","name":"Bob"},{"id":"2","name":"Steve"}]` + "\n"
		assert.Equal(t, expect, rec.Body.String())
	}
}

func TestProfilerGet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/3", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/profile/3")
	c.SetParamNames("id")
	c.SetParamValues("3")
	p := NewProfiler()

	if assert.NoError(t, p.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expect := `{"id":"3","name":"Bob"}` + "\n"
		assert.Equal(t, expect, rec.Body.String())
	}
}

func TestProfilerCreate(t *testing.T) {
	e := echo.New()
	p := NewProfiler()
	var scenario string

	scenario = "201 Created"
	{
		expect := `{"id":"4","name":"Hoge"}` + "\n"
		rec, c := newHTTPTestWithBody(e, http.MethodPost, expect)

		if assert.NoError(t, p.Create(c), scenario) {
			assert.Equal(t, http.StatusCreated, rec.Code, scenario)
			assert.Equal(t, expect, rec.Body.String(), scenario)
		}
	}

	scenario = "400 Bad Request"
	{
		expect := `` + "\n"
		_, c := newHTTPTestWithBody(e, http.MethodPost, expect)

		assert.Error(t, p.Create(c), scenario)
	}
}

func newHTTPTestWithBody(e *echo.Echo, method string, body string) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(method, "/", strings.NewReader(body))

	switch method {
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	default:
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

func TestProfilerReplace(t *testing.T) {
	e := echo.New()
	p := NewProfiler()
	var scenario string

	scenario = "200 OK"
	{
		expect := `{"id":"4","name":"Fuga"}` + "\n"
		rec, c := newHTTPTestWithBody(e, http.MethodPut, expect)

		if assert.NoError(t, p.Replace(c), scenario) {
			assert.Equal(t, http.StatusOK, rec.Code, scenario)
			assert.Equal(t, expect, rec.Body.String(), scenario)
		}
	}

	scenario = "400 Bad Request"
	{
		b := "\n"
		_, c := newHTTPTestWithBody(e, http.MethodPut, b)

		assert.Error(t, p.Replace(c), scenario)
	}
}
