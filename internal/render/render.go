package render

import (
	"assignment-3/internal/form"
	"fmt"
	"net/http"
	"text/template"
)

// TemplateData stores data to be used in Templates
type TemplateData struct {
	Data map[string]interface{}
	Form *form.Form
}

// Template parses and exectues template by its template name
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData) error {

	ts, err := template.ParseFiles(fmt.Sprintf("./templates/%s", tmpl), "./templates/base.layout.html", "./templates/header.layout.html")
	if err != nil {
		return fmt.Errorf("ParseTemplate: Unable to find template pages: %w", err)
	}

	if err := ts.Execute(w, td); err != nil {
		return fmt.Errorf("ParseTemplate: Unable to execute template: %w", err)
	}

	return nil
}
