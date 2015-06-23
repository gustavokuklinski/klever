# Klever

### Kickstart
Setup your Go Workspace, create a new file and paste the Klever Hello World :)

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

Will be generate a config.json for config the HTTP port

```json
{
 "AppPort":"8080"
}
``` 

By default will be port 8080, you can change.
To apply the changes on this file you must restart the aplication
*This project was made just for study GoLang, don't use in production yet :)*
