// Package html provides functionality to handle html stuff.
package html

import (
	"golang.org/x/net/html"
	"log"
	"strings"
)

// RemoveTags removes any HTML or XML tags from given string.
func RemoveTags(value string) string {
	doc, err := html.Parse(strings.NewReader(value))
	if err != nil {
		log.Fatal(err)
	}
	var values []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			text := strings.Replace(strings.TrimSpace(n.Data), "\n", " ", -1)
			if len(text) > 0 {
				values = append(values, text)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return strings.Join(values, " ")
}
