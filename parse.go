package link

import (
	"io"

	"golang.org/x/net/html"
)

// 1. find the <a> nodes
// 2. create a Link instance for each <a> node
// 3. return all the Links in a slice

// Link represents a link <a href="..."> in an HTML document
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML document and returns a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	var links []Link
	// html.Parse returns the root of the parse tree as a *Node from the given Reader
	rootNode, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(rootNode)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

// takes in root node, returns all link nodes
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodeSlice []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodeSlice = append(nodeSlice, linkNodes(c)...)
	}
	return nodeSlice
}

func buildLink(n *html.Node) Link {
	var link Link

	// the node's attributes are stored in a slice, have to loop over them
	// Namespace, Key, Val string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = "TODO: Parse the text"

	return link
}
