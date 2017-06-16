package goutils

import (
	"html/template"
	"strings"
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

/*
 * template helpers
 */
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
