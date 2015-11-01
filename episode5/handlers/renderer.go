package handlers

import (
	"html/template"
	"net/http"

	"github.com/unrolled/render"
)

// Renderer is the interface to rendering templates. This example repository
// has an unrolled/render (https://godoc.org/github.com/unrolled/render) implementation,
// but other implementations are certainly possible
type Renderer interface {
	// Render renders a template with data and passes the rendered result to w with the given code.
	// if layout is non-empty, this func renders the template with that name, and the templateName
	// is rendered under the {{yield}} block. otherwise this func renders the template called templateName
	Render(w http.ResponseWriter, code int, templateName string, data interface{}, layout string)
}

// RenderRenderer is a Renderer implementation that uses unrolled/render (https://godoc.org/github.com/unrolled/render#Render.HTML)
// to do template rendering
type RenderRenderer struct {
	r *render.Render
}

// NewRenderRenderer returns a new RenderRenderer, where the underlying render.Render
// serves templates out of dir with the given func map. it's configured in dev mode
// according to the dev boolean
func NewRenderRenderer(dir string, extensions []string, funcs []template.FuncMap, dev bool) *RenderRenderer {
	opts := render.Options{
		Directory:     dir,
		Extensions:    extensions,
		Funcs:         funcs,
		IsDevelopment: dev,
	}
	return &RenderRenderer{r: render.New(opts)}
}

// Render is the interface implementation
func (r *RenderRenderer) Render(w http.ResponseWriter, code int, templateName string, data interface{}, layout string) {
	if layout != "" {
		r.r.HTML(w, code, templateName, data, render.HTMLOptions{Layout: layout})
		return
	}
	r.r.HTML(w, code, templateName, data)
}
