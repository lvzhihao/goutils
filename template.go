package goutils

import (
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo"
)

var (
	TemplateHelpers = template.FuncMap{
		"IntSub": func(v1 interface{}, v2 interface{}) int64 {
			return IntArithmetic(v1, v2, "-")
		},
		"IntAdd": func(v1 interface{}, v2 interface{}) int64 {
			return IntArithmetic(v1, v2, "+")
		},
		"ToLower": func(s interface{}) string {
			return strings.ToLower(ToString(s))
		},
	}
)

func NewTemplate(name, patter string) *template.Template {
	return template.Must(template.New(name).Funcs(TemplateHelpers).ParseGlob(patter))
}

type EchoTemplate struct {
	templates *template.Template
}

func (t *EchoTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewEchoRenderer(name, patter string) *EchoTemplate {
	return &EchoTemplate{
		templates: NewTemplate(name, patter),
	}
}

func IntArithmetic(v1, v2 interface{}, t string) (result int64) {
	switch t {
	case "+":
		result = ToInt(v1) + ToInt(v2)
		break
	case "-":
		result = ToInt(v1) - ToInt(v2)
		break
	}
	return
}
