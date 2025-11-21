package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
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

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/testing", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test.html", nil)
	})


	e.POST("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", nil)
	})

	e.POST("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "start.html", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

