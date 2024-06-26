package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var (
	links   []string
	visited = map[string]bool{}
)

func main() {

	visitLink("https://aprendagolang.com.br")

	fmt.Println(len(links))
}

func visitLink(link string) {
	if ok := visited[link]; ok {
		return
	}
	visited[link] = true
	fmt.Println("link visitado:", link)

	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status code diferentede 200: %d", resp.StatusCode))
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	extractLinks(node)

}

func extractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}

			links = append(links, link.String())

			visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}

}
