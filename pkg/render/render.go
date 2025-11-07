package render

import (
	"bytes"
	"html/template"
	"log"
	"github.com/SiddhantSShende/bookings-app/pkg/config"
	"github.com/SiddhantSShende/bookings-app/pkg/models"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Best approach for production use
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Use the template cache
		tc = app.TemplateCache
	} else {
		// Create a new template cache
		tc, _ = CreateTemplateCache()
	}

	/* // Use this code if you want to always create a new template cache (without config approach)
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatalln("Error creating template cache:", err)
		return
	}
	*/

	// Parse the template files
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("Could not get template from template cache")
		return
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	// Execute the template

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatalln("Error executing template:", err)
		return
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalln("Error writing template to browser:", err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Look for layout templates
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

/*
var templateCache = make(map[string]*template.Template)

// RenderTemplate renders templates using a cache to improve performance
// Production level approach (but need to add pages manually to the cache)
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	// check to see if we already have the template in our cache
	_, inMap := templateCache[t]
	if !inMap {
		// need to create the template
		log.Println("(Creating template and adding it to cache) Loading template:", t)
		err = CreateTemplateCache(t)
		if err != nil {
			log.Println("error creating template cache:", err)
			return
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}
	tmpl = templateCache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error executing template:", err)
		return
	}
}

func CreateTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add it to the cache map
	templateCache[t] = tmpl
	return nil
}
*/

/*
// Bad approach not recommended for production use(but the simplest one)
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl") // Parsing the template file
	err := parsedTemplate.Execute(w, nil)                                                         // Executing the template
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

*/
