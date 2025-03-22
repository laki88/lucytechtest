package models

import (
	"strings"

	"golang.org/x/net/html"
)

type AnalysisResult struct {
	Status        string         `json:"status"`
	HTMLVersion   string         `json:"html_version"`
	Title         string         `json:"title"`
	Headings      map[string]int `json:"headings"`
	InternalLinks int            `json:"internal_links"`
	ExternalLinks int            `json:"external_links"`
	BrokenLinks   int            `json:"broken_links"`
	LoginForm     bool           `json:"login_form"`
	Message       string         `json:"message,omitempty"`
}

func AnalyzeHTML(doc *html.Node, baseURL string) AnalysisResult {
	var title string
	headings := map[string]int{}
	internalLinks, externalLinks, brokenLinks := 0, 0, 0
	loginForm := false

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					title = n.FirstChild.Data
				}
			case "h1", "h2", "h3", "h4", "h5", "h6":
				headings[n.Data]++
			case "a":
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						if strings.HasPrefix(attr.Val, "http") {
							externalLinks++
						} else {
							internalLinks++
						}
					}
				}
			case "form":
				for _, attr := range n.Attr {
					if attr.Key == "action" && strings.Contains(attr.Val, "login") {
						loginForm = true
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return AnalysisResult{
		Status:        "Completed",
		HTMLVersion:   "HTML5",
		Title:         title,
		Headings:      headings,
		InternalLinks: internalLinks,
		ExternalLinks: externalLinks,
		BrokenLinks:   brokenLinks,
		LoginForm:     loginForm,
	}
}
