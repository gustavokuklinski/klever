//  Klever - Micro router framework for static web sites
package klever

import (
	"github.com/gustavokuklinski/klever/scaffold"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Render Struct
// Used with Layout and Page function
type Render struct {
  tplD string // Store the directory
  tplF string // Store the file to render
  tplR string // Store the URL
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
  tpl := Render{tplD: tplDir, tplF: tplFile}

  // Use the Struct data to build the template
  page := filepath.Join(tpl.tplD, tpl.tplF)

	tmpl, err := template.ParseFiles(layout, head, nav, footer, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Render Route and Page.
// Usage:
// In the main function use: klever.Page(route, file)
func Page(route, file string) {
  
  // Mount the struct and get the values from route and file 
  tpl := Render{tplF: file, tplR: route}

  // Use the struct data to build the route and send the file to Layout function
  // Responsible to build the template
	http.HandleFunc(tpl.tplR, func(w http.ResponseWriter, r *http.Request) {
		Layout("pages", tpl.tplF, w)
	})
}

// Generate base Dirs
// Create the main directory tree inside the root directory
func generateDirs() {
	// Generate base folders
	os.Mkdir("."+string(filepath.Separator)+"includes", 0777)
	os.Mkdir("."+string(filepath.Separator)+"pages", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets", 0777)
	os.Mkdir("."+string(filepath.Separator)+"posts", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets"+string(filepath.Separator)+"img", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets"+string(filepath.Separator)+"css", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets"+string(filepath.Separator)+"js", 0777)

}

// Start Klever in three steps:
// 1. Create the base directory tree.
// 2. Generate base Scaffold templates(Check the package: github.com/gustavokuklinski/klever/scaffold).
// 3. Load a basic HTTP Server.
func Start() {

	// Check if folders exists
	// If not, create a brand new project

	// Check [/includes] dir
	if _, err := os.Stat("includes"); os.IsNotExist(err) {
		generateDirs()

		// Check [/assets] dir
	} else if _, err := os.Stat("assets"); os.IsNotExist(err) {
		generateDirs()

		// Check [/pages] dir
	} else if _, err := os.Stat("pages"); os.IsNotExist(err) {
		generateDirs()

		// Check [/posts] dir
	} else if _, err := os.Stat("posts"); os.IsNotExist(err) {
		generateDirs()
	}

	// Start Scaffold lib to generate base files
	scaffold.GenerateScaffold()

	// File server for static files(CSS, JS and Images) in folder [/assets]
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Start webserver on port: 8080 - You can fit your need :)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
