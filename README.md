# Klever

### Kickstart

```go
package main

import (
	"github.com/gustavokuklinski/klever"
)

func main() {
  /*
   * Klever Pages
   * Syntax: klever.Page(route, file)
   */ 
  klever.Page("/", "index")

  /* Start the application */
  klever.Start()
}
```

and them

```bash
$go get
$go run app.go
```

Access your: localhost:8080

This project was made just for study GoLang, don't use in production yet :)
