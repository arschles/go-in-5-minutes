package form

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags"
)

//CheckboxTag builds a checkbox from the options passed
func (f Form) CheckboxTag(opts tags.Options) *tags.Tag {
	opts["type"] = "checkbox"

	value := opts["value"]
	delete(opts, "value")

	checked := opts["checked"]
	delete(opts, "checked")
	if checked == nil {
		checked = "true"
	}
	opts["value"] = checked

	unchecked := opts["unchecked"]
	delete(opts, "unchecked")

	hl := opts["hide_label"]
	delete(opts, "hide_label")

	if opts["tag_only"] == true {
		delete(opts, "label")
		ct := f.InputTag(opts)
		ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
		return ct
	}

	tag := tags.New("label", tags.Options{})
	ct := f.InputTag(opts)
	ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
	tag.Append(ct)

	if opts["name"] != nil && unchecked != nil {
		tag.Append(tags.New("input", tags.Options{
			"type":  "hidden",
			"name":  opts["name"],
			"value": unchecked,
		}))
	}

	if opts["label"] != nil && hl == nil {
		label := fmt.Sprint(opts["label"])
		delete(opts, "label")
		tag.Append(" " + label)
	}
	return tag
}
