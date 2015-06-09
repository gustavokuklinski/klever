//  Klever - Micro router framework for static web sites
package klever

import (
	"github.com/gustavokuklinski/klever/scaffold"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Render Struct
// Used with Layout and Page function
type Render struct {
	TplD string // Store the directory (Default: Pages)
	TplF string // Store the file to render
	TplR string // Store the route URL
}

// Render templates.
// Use the directory tree and the scaffold package to set the right templates and
// directories of the project this function read two parameters:
// tplDir: the base [/pages] and tplFile: the file set on klever.Page(route, file)
func Layout(tplDir, tplFile string, w http.ResponseWriter) {

	// [layout.html]
	layout := filepath.Join("includes", "layout.html")

	// [head.html]
	head := filepath.Join("includes", "head.html")

	// [nav.html]
	nav := filepath.Join("includes", "nav.html")

	// [footer.html]
	footer := filepath.Join("includes", "footer.html")

	// Mount the struct and get the values from tplDir and tplFile
	tpl := Render{TplD: tplDir, TplF: tplFile}

	// Use the Struct data to build the template
	page := filepath.Join(tpl.TplD, tpl.TplF)

	// Parse template file
	tmpl, err := template.ParseFiles(layout, head, nav, footer, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Render Route and Page.
// Usage:
// In the main function use: klever.Page(route, file)
func Page(route, file string) {

	// Mount the struct and get the values from route and file
	tpl := Render{TplF: file, TplR: route}

	// Use the struct data to build the route and send the file to Layout function
	// Responsible to build the template
	http.HandleFunc(tpl.TplR, func(w http.ResponseWriter, r *http.Request) {

		// Layout function, use [/pages] as default folder to serve pages file
		Layout("pages", tpl.TplF, w)
	})
}

// Start Klever in two steps:
// 1. Generate base directories and scaffold templates(Check the package: github.com/gustavokuklinski/klever/scaffold).
// 2. Load a basic HTTP Server.
func Start() {

	// Start Scaffold lib to generate base files
	scaffold.GenerateScaffold()

	// File server for static files(CSS, JS and Images) in folder [/assets]
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Start webserver on port: 8080 - You can fit your need :)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}
