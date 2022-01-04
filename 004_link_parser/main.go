package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type link struct {
	href, text string
}

var links []link

func main() {
	htmlFileName := flag.String("htmlFileName", "sample.html", "A HTML file to parse anchor tags")
	flag.Parse()
	htmlFile, err := os.Open(*htmlFileName)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	headNode, err := html.Parse(htmlFile)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	parseLinks(headNode)
	fmt.Println(links)

}

func parseLinks(node *html.Node) {
	if node == nil {
		return
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode && child.Data == "a" {
			text := getTextOfAnchorNode(child, "")
			href := ""
			for _, val := range child.Attr {
				if val.Key == "href" {
					href = val.Val
				}
			}
			links = append(links, link{href: href, text: text})
			continue
		}
		parseLinks(child)
	}
}

func getTextOfAnchorNode(node *html.Node, text string) string {
	if node == nil {
		return ""
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			text = getTextOfAnchorNode(child, text)
		} else if child.Type == html.TextNode {
			text += child.Data
		}
	}
	return strings.TrimSpace(text)
}
