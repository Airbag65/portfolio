package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Title       string
	Tags        []string
	Description string
	Link        string
	Id          int
}

type Projects struct {
	Items []Project
}

func NewMockProject() Project {
	return Project{
		Title:       "Portfolio",
		Tags:        []string{"Golang", "HTMX"},
		Description: "Portfolio to showcase my projects and expertice",
		Link:        "https://www.github.com/Airbag65/portfolio",
		Id:          1,
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	templ, err := template.ParseGlob("./public/views/*.html")
	if err != nil {
		log.Fatalf("Could not parse: %v\n", err)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: templ,
	}

	project := NewMockProject()
	projects := Projects{
		Items: []Project{project},
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/projects", func(c echo.Context) error {
		return c.Render(http.StatusOK, "projects.html", projects)
	})

	e.POST("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", nil)
	})

	e.POST("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "start.html", nil)
	})

	e.POST("/project/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.Render(http.StatusOK, "project.html", id)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
