package handlers

import "html/template"

var Funcs = []template.FuncMap{
	template.FuncMap(map[string]interface{}{
		"Pluralize": func(num int, singular, plural string) string {
			if num == 1 {
				return singular
			}
			return plural
		},
	}),
}
