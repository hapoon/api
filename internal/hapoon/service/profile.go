package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Profiler is profileService interface
type Profiler interface {
	List(echo.Context) error
	Get(echo.Context) error
	Create(echo.Context) error
	Replace(echo.Context) error
}

type profileService struct {
}

type profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p *profileService) List(c echo.Context) error {
	res := []profile{
		{ID: "1", Name: "Bob"},
		{ID: "2", Name: "Steve"},
	}
	return c.JSON(http.StatusOK, res)
}

func (p *profileService) Get(c echo.Context) error {
	id := c.Param("id")
	res := profile{
		ID:   id,
		Name: "Bob",
	}
	return c.JSON(http.StatusOK, res)
}

func (p *profileService) Create(c echo.Context) error {
	pf := new(profile)
	if err := c.Bind(pf); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, pf)
}

func (p *profileService) Replace(c echo.Context) error {
	pf := new(profile)
	if err := c.Bind(pf); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pf)
}

// NewProfiler create Profiler instance
func NewProfiler() Profiler {
	return &profileService{}
}
