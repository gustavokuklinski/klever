// Package scaffold - Generates the basic template files
package scaffold

import (
	"os"
	"path/filepath"
)

// GenerateScaffoldDirs - Create the main directory tree inside the root directory
func GenerateScaffoldDirs() {
	// Generate base folders
	os.Mkdir("."+string(filepath.Separator)+"includes", 0777)
	os.Mkdir("."+string(filepath.Separator)+"pages", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets", 0777)
	os.Mkdir("."+string(filepath.Separator)+"posts", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets"+string(filepath.Separator)+"img", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets"+string(filepath.Separator)+"css", 0777)
	os.Mkdir("."+string(filepath.Separator)+"assets"+string(filepath.Separator)+"js", 0777)

}

// GenerateScaffoldFiles create base scaffold template
func GenerateScaffoldFiles() {
	// Base Config files
	// Create: config.json on Root dir, Default on port: 8080
	baseConfig, err := os.Create("." + string(filepath.Separator) + "config.json")

	// Base Template files
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

	// Close files after end writting files
	defer baseConfig.Close()
	defer baseFooter.Close()
	defer baseHead.Close()
	defer baseNav.Close()
	defer baseLayout.Close()
	defer indexPage.Close()

	// Write the base config
	baseConfig.WriteString("{\n \"AppPort\":\"8080\"\n}")

	// Write the standard templates
	// Base Scaffold HTML layout
	baseLayout.WriteString("{{ define \"layout\" }}\n<!DOCTYPE html>\n<html lang=\"en_US\">\n<head>\n{{ template \"head\" }}\n</head>\n<body>\n{{ template \"nav\" }}\n{{ template \"body\" . }}\n{{ template \"footer\" }}\n</body>\n</html>\n{{ end }}")

	// <HEAD> Tag template
	baseHead.WriteString("{{ define \"head\" }}\n<title> {{ template \"title\" }} </title>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">*{margin:0; padding: 0;}body {background: #34495e;} a {font-family: Arial; color: #ffffff;} h1 { font-family: Arial; text-align: center; padding-bottom: 10px; } nav{width: 100%; background: #000000; padding-top: 20px; padding-bottom: 20px} nav a { padding-left: 10px; font-family: Arial; color: #ffffff; } footer{width: 100%; background: #34495e; padding-top: 25px; padding-bottom: 25px; font-family: Arial;} section{background: #cccccc; padding-top: 10%; padding-bottom: 10%;} p { font-family: Arial; text-align: center; }</style>{{ end }}")

	// <NAV> bar menu template
	baseNav.WriteString("({{ define \"nav\" }}\n<nav>\n<a href=\"/\">Home</a> <a href=\"http://www.golang.org\" target=\"_blank\"/>Golang Web site</a>\n</nav>\n{{ end }}")

	// <FOOTER> Tag template
	baseFooter.WriteString("{{ define \"footer\" }}\n<footer><center><p>Klever</p>\n<a href=\"http://github.com/gustavokuklinski/klever\" target=\"_blank\">Github</a></center></footer>{{ end }}")

	// <BODY> An Hello World main page
	indexPage.WriteString("{{ define \"title\" }} Klever Home {{ end }}\n{{ define \"body\" }}\n<section>\n<h1>Welcome to KLEVER!</h1>\n<p> Start writing your pages :) <br /><br /> This is a Dumb <b>hello world</b>.<br /> If you see this, you've done it Right!...<br /><br /> Why don't you start crafting a real nice website? </p>\n</section>\n{{ end }}")

	// Write to files
	baseConfig.Sync()
	baseFooter.Sync()
	baseHead.Sync()
	baseNav.Sync()
	baseLayout.Sync()
	indexPage.Sync()

}

// GenerateScaffold checks if any of HTML files below exists or base directories, If not, generate fallback templates
func GenerateScaffold() {

	// Check if folders exists
	// If not, create a brand new project
	// Check [/includes] dir
	if _, err := os.Stat("includes"); os.IsNotExist(err) {
		GenerateScaffoldDirs()

		// Check [/assets] dir
	} else if _, err := os.Stat("assets"); os.IsNotExist(err) {
		GenerateScaffoldDirs()

		// Check [/pages] dir
	} else if _, err := os.Stat("pages"); os.IsNotExist(err) {
		GenerateScaffoldDirs()

		// Check [/posts] dir
	} else if _, err := os.Stat("posts"); os.IsNotExist(err) {
		GenerateScaffoldDirs()

	}

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

	} else if _, err := os.Stat("." + string(filepath.Separator) + "config.json"); os.IsNotExist(err) {
		GenerateScaffoldFiles()

	}

}
