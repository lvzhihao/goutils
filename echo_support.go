package goutils

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

/*
 * echo templater support
 */
type EchoTemplate struct {
	templates *template.Template
}

func NewEchoRenderer(name, patter string) *EchoTemplate {
	return &EchoTemplate{
		templates: NewTemplate(name, patter),
	}
}

func (t *EchoTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
