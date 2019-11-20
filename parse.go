package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link <a href="..."> in an HTML document
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML document and returns a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	// html.Parse returns the root of the parse tree as a *Node from the given Reader
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	dfs(doc, "")
	// 1. find the <a> nodes
	// 2. create a Link instance for each <a> node
	// 3. return all the Links in a slice

	return nil, nil
}

func dfs(n *html.Node, padding string) {
	// n.Data will give us what's inside the node
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
