// Klever Scaffold - Generates the basic template files
package scaffold

import (
	"os"
	"path/filepath"
)

// Generate base scaffold template
func GenerateScaffoldFiles() {

	// Create: [layout.html] in [/includes] : Render <html></html>
	baseLayout, err := os.Create("includes" + string(filepath.Separator) + "layout.html")

	// Create: [head.html] in [/includes] : Render <head></head>
	baseHead, err := os.Create("includes" + string(filepath.Separator) + "head.html")

	// Create: [nav.html] in [/includes] : Render <nav></nav>
	baseNav, err := os.Create("includes" + string(filepath.Separator) + "nav.html")

	// Create: [footer.html] in [/includes] : Render <footer></footer>
	baseFooter, err := os.Create("includes" + string(filepath.Separator) + "footer.html")

	// Create: [index] in [/pages]
	indexPage, err := os.Create("pages" + string(filepath.Separator) + "index")

	if err != nil {
		panic(err)
	}

	// When file creation finished, Close files
	defer baseFooter.Close()
	defer baseHead.Close()
	defer baseNav.Close()
	defer baseLayout.Close()
	defer indexPage.Close()

	// Write the standard templates
	// Base Scaffold HTML layout
	baseLayout.WriteString("{{ define \"layout\" }}\n<!DOCTYPE html>\n<html lang=\"en_US\">\n<head>\n{{ template \"head\" }}\n</head>\n<body>\n{{ template \"nav\" }}\n{{ template \"body\" }}\n{{ template \"footer\" }}\n</body>\n</html>\n{{ end }}")

	// <HEAD> Tag template
	baseHead.WriteString("{{ define \"head\" }}\n<title> {{ template \"title\" }} </title>\n<meta charset=\"utf-8\" />\n{{ end }}")

	// <NAV> bar menu template
	baseNav.WriteString("({{ define \"nav\" }}\n<nav>\n<a href=\"/\">Home</a> | <a href=\"http://www.golang.org\" target=\"_blank\"/>Go Web site</a>\n</nav>\n{{ end }}")

	// <FOOTER> Tag template
	baseFooter.WriteString("{{ define \"footer\" }}\nKlever - Build with Go :D\n{{ end }}")

	// <BODY> An Hello World main page
	indexPage.WriteString("{{ define \"title\" }} Klever Home {{ end }}\n{{ define \"body\" }}\n<section>\n<h1>Welcome to KLEVER!</h1>\n<p> Start writing your pages :D </p>\n</section>\n{{ end }}")

	// Write to files
	baseFooter.Sync()
	baseHead.Sync()
	baseNav.Sync()
	baseLayout.Sync()
	indexPage.Sync()

}

// Checks if any of HTML files below exists, If not, generate fallback templates
func GenerateScaffold() {

	// Check for: layout.html
	if _, err := os.Stat("includes" + string(filepath.Separator) + "layout.html"); os.IsNotExist(err) {
		GenerateScaffoldFiles()

		// Check for: head.html
	} else if _, err := os.Stat("includes" + string(filepath.Separator) + "head.html"); os.IsNotExist(err) {
		GenerateScaffoldFiles()

		// Check for: nav.html
	} else if _, err := os.Stat("includes" + string(filepath.Separator) + "nav.html"); os.IsNotExist(err) {
		GenerateScaffoldFiles()

		// check for: footer.html
	} else if _, err := os.Stat("includes" + string(filepath.Separator) + "footer.html"); os.IsNotExist(err) {
		GenerateScaffoldFiles()

	}

}
