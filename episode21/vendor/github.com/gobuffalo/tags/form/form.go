package form

import (
	"strings"

	"github.com/gobuffalo/tags"
)

//Form is the html form tag, this will hold other tags inside it.
type Form struct {
	*tags.Tag
}

//SetAuthenticityToken allows tags to work smoothly with Buffalo, it receives the auth token and creates an input hidden with it.
func (f *Form) SetAuthenticityToken(s string) {
	f.Prepend(tags.New("input", tags.Options{
		"value": s,
		"type":  "hidden",
		"name":  "authenticity_token",
	}))
}

//Label permits to create a label inside a Form
func (f Form) Label(value string, opts tags.Options) *tags.Tag {
	opts["body"] = value
	return tags.New("label", opts)
}

//New creates a new form from passed options, it sets defaults for method and also handles other methods as PUT by adding _method hidden input.
func New(opts tags.Options) *Form {
	if opts["method"] == nil {
		opts["method"] = "POST"
	}

	if opts["multipart"] != nil {
		opts["enctype"] = "multipart/form-data"
		delete(opts, "multipart")
	}

	form := &Form{
		Tag: tags.New("form", opts),
	}

	m := strings.ToUpper(form.Options["method"].(string))
	if m != "POST" && m != "GET" {
		form.Options["method"] = "POST"
		form.Prepend(tags.New("input", tags.Options{
			"value": m,
			"type":  "hidden",
			"name":  "_method",
		}))
	}

	return form
}
