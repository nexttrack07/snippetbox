package main

import (
	"html/template"
	"nexttrack07/snippetbox/pkg/forms"
	"nexttrack07/snippetbox/pkg/models"
	"path/filepath"
	"time"
)

type templateData struct {
	CSRFToken       string
	CurrentYear     int
	Form            *forms.Form
	Flash           string
	IsAuthenticated bool
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// Use filepath.Glob method to get a slice of all the filepaths for the pages
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	//Loop through the pages
	for _, page := range pages {
		// Extract the file name using filepath.Base
		// and assign to a new variable
		name := filepath.Base(page)

		// Parse the page template file into a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Parse all the layout templates and insert into ts
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Parse all the partial templates and insert into ts
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// assign the template set to the cache
		cache[name] = ts
	}

	return cache, nil
}
