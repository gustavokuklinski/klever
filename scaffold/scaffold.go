/*
 * Klever Scaffold
 * Version: 0.1 build 07032015
 * ----------------------------------------------------------
 * This small lib create the base template files
 * needed to run the main application.
 * It's used everytime the application Startup, if any base
 * template brokes, this lib will put it back in fallback
 * mode(Basic HTML, with no CSS)
 */

package scaffold

import (
  "os"
  "path/filepath"
)

/*
 *  Generate base Dirs and Scaffold template
 *  ---------------------------------------------------------
 *  (03/07/2015)
 */

func generateScaffoldFiles() {
    // Generate base files
    baseLayout, err := os.Create("includes" + string(filepath.Separator) + "layout.html")
    baseHead, err := os.Create("includes" + string(filepath.Separator) + "head.html")
    baseNav, err := os.Create("includes" + string(filepath.Separator) + "nav.html")
    baseFooter, err := os.Create("includes" + string(filepath.Separator) + "footer.html")

    indexPage, err := os.Create("pages" + string(filepath.Separator) + "index")

    if err != nil {
      panic(err)
    }

    defer baseFooter.Close()
    defer baseHead.Close()
    defer baseNav.Close()
    defer baseLayout.Close()
    defer indexPage.Close()

    // Write standard templates

    /*
     * Basic layout template
     */
    baseLayout.WriteString("{{ define \"layout\" }}\n<!DOCTYPE html>\n<html lang=\"en_US\">\n<head>\n{{ template \"head\" }}\n</head>\n<body>\n{{ template \"nav\" }}\n{{ template \"body\" }}\n{{ template \"footer\" }}\n</body>\n</html>\n{{ end }}")
    
    /*
     * <HEAD> Tag template
     */
    baseHead.WriteString("{{ define \"head\" }}\n<title> {{ template \"title\" }} </title>\n<meta charset=\"utf-8\" />\n{{ end }}")

    /*
     * <NAV> bar menu template
     */
    baseNav.WriteString("({{ define \"nav\" }}\n<nav>\n<a href=\"/\">Home</a> | <a href=\"http://www.golang.org\" target=\"_blank\"/>Go Web site</a>\n</nav>\n{{ end }}")
    
    /*
     * <FOOTER> Tag template
     */
    baseFooter.WriteString("{{ define \"footer\" }}\nKlever - Build with Go :D\n{{ end }}")
    
    /*
     * <BODY> An Hello World main page
     */
    indexPage.WriteString("{{ define \"title\" }} Klever Home {{ end }}\n{{ define \"body\" }}\n<section>\n<h1>Welcome to KLEVER!</h1>\n<p> Start writing your pages :D </p>\n</section>\n{{ end }}")
  
    // Write to files
    baseFooter.Sync()
    baseHead.Sync()
    baseNav.Sync()
    baseLayout.Sync()
    indexPage.Sync()

}

func GenerateScaffold() {
  /*
   * Checks if any of HTML files below exists, If not, create fallback template
   */
  if _, err := os.Stat("includes" + string(filepath.Separator) + "layout.html"); os.IsNotExist(err) {
    generateScaffoldFiles()

  } else if _, err := os.Stat("includes" + string(filepath.Separator) + "head.html"); os.IsNotExist(err) {
    generateScaffoldFiles()
  
  } else if _, err := os.Stat("includes" + string(filepath.Separator) + "nav.html"); os.IsNotExist(err) {
    generateScaffoldFiles()
  
  } else if _, err := os.Stat("includes" + string(filepath.Separator) + "footer.html"); os.IsNotExist(err) {
    generateScaffoldFiles()
  
  }


}
