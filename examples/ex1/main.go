package main

import (
	"fmt"
	"gophercises/link"
	"log"
	"strings"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	check(err)
	fmt.Printf("%+v\n", links)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
