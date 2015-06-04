/*
 *  Klever
 *  Version: 0.1 build 07022015
 *  ---------------------------------------------------------
 *  Micro router framework
 *  for static web sites using
 *  Go templates
 *  ---------------------------------------------------------
 *
 *  DOCUMENTATION
 *  ---------------------------------------------------------
 *  1. Import the library: github/gustavo/klever
 *  2. Build a Web page: klever.Page(route, file)
 *  2-1. There is a directory struct:
 *       [/assets]   - Place for CSS, Images and JavaScripts
 *       [/includes] - Base Layout, Header, Footer, Navbar
 *       [/pages]    - Base pages for the Website
 *  ---------------------------------------------------------
 */
package klever

import (
	"net/http"
  "path"
  "log"
  "html/template"
  "os"
  "path/filepath"
  "github.com/gustavokuklinski/klever/scaffold"
)
/*
 *  Render templates
 *  ---------------------------------------------------------
 *  (02/07/2015)
 */
func Layout(tplDir, tplFile string, w http.ResponseWriter) {
	layout := path.Join("includes", "layout.html") // Base layout
	head := path.Join("includes", "head.html")     // Head tags (<head></head>) and contents
	nav := path.Join("includes", "nav.html")       // Menu (<nav></nav>)
	footer := path.Join("includes", "footer.html") // Footer (<footer></footer>)
  page := path.Join(tplDir, tplFile)

  tmpl, err := template.ParseFiles(layout, head, nav, footer, page) 
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

/*
 *  Render Route and Page
 *  Ex: klever.Page(route, file)
 *  ---------------------------------------------------------
 *  (02/07/2015)
 */
func Page(route, file string) {
  http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
    Layout("pages", file, w)
  }) 
}


/*
 *  Generate base Dirs
 *  ---------------------------------------------------------
 *  (03/07/2015)
 */
func generateDirs() {
   // Generate base folders
  os.Mkdir("." + string(filepath.Separator) + "includes", 0777) 
  os.Mkdir("." + string(filepath.Separator) + "pages", 0777) 
  os.Mkdir("." + string(filepath.Separator) + "assets", 0777)
  os.Mkdir("." + string(filepath.Separator) + "posts", 0777)
  os.Mkdir("." + string(filepath.Separator) + "assets" + string(filepath.Separator) + "img", 0777) 
  os.Mkdir("." + string(filepath.Separator) + "assets" + string(filepath.Separator) + "css", 0777) 
  os.Mkdir("." + string(filepath.Separator) + "assets" + string(filepath.Separator) + "js", 0777)

}

/*
 * Create base directories
 * Start HTTP Server
 */
func Start() {
  /*
   * Check if folders exists
   * If not, create a brand new project
   */
  if _, err := os.Stat("includes"); os.IsNotExist(err) {
    generateDirs()
  } else if _, err := os.Stat("assets"); os.IsNotExist(err) {
    generateDirs()
  } else if _, err := os.Stat("pages"); os.IsNotExist(err) {
    generateDirs()
  } else if _, err := os.Stat("posts"); os.IsNotExist(err) {
    generateDirs()
  }

  // Start Scaffold lib to generate base files
  scaffold.GenerateScaffold()

  // Assets folder (CSS, JS, IMGs)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Start webserver on port: 8080 - You can fit your need :)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
