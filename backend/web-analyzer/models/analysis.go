package models

import (
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type AnalysisResult struct {
	Status        string         `json:"Status"`
	HTMLVersion   string         `json:"HTML Version"`
	Title         string         `json:"Title"`
	Headings      map[string]int `json:"Headings"`
	InternalLinks int            `json:"Internal Links"`
	ExternalLinks int            `json:"External Links"`
	BrokenLinks   int            `json:"Broken Links"`
	LoginForm     string         `json:"Login Form"`
	Message       string         `json:"Message,omitempty"`
}

var htmlVersions = map[string]string{
	"-//W3C//DTD HTML 4.01//EN":              "HTML 4.01 Strict",
	"-//W3C//DTD HTML 4.01 TRANSITIONAL//EN": "HTML 4.01 Transitional",
	"-//W3C//DTD HTML 4.01 FRAMESET//EN":     "HTML 4.01 Frameset",
	"-//W3C//DTD HTML 4.0//EN":               "HTML 4.0 Strict",
	"-//W3C//DTD HTML 4.0 TRANSITIONAL//EN":  "HTML 4.0 Transitional",
	"-//W3C//DTD HTML 4.0 FRAMESET//EN":      "HTML 4.0 Frameset",
	"-//W3C//DTD HTML 3.2 FINAL//EN":         "HTML 3.2",
	"-//IETF//DTD HTML//EN":                  "HTML 2.0",
	"-//W3C//DTD XHTML 1.0 STRICT//EN":       "XHTML 1.0 Strict",
	"-//W3C//DTD XHTML 1.0 TRANSITIONAL//EN": "XHTML 1.0 Transitional",
	"-//W3C//DTD XHTML 1.0 FRAMESET//EN":     "XHTML 1.0 Frameset",
	"-//W3C//DTD XHTML 1.1//EN":              "XHTML 1.1",
}

func isBrokenLink(url string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(url)
	if err != nil || resp.StatusCode >= 400 {
		return true
	}
	return false
}

func DetectHTMLVersion(doc *html.Node) string {
	// Traverse the children of the document node to find the DOCTYPE node.
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.DoctypeNode {
			// Ensure the DOCTYPE is for an HTML document.
			if strings.ToLower(c.Data) != "html" {
				return "Unknown document type"
			}
			// HTML5 has no attributes in its DOCTYPE.
			if len(c.Attr) == 0 {
				return "HTML5"
			}
			// Check attributes for public identifier to determine older HTML or XHTML versions.
			for _, attr := range c.Attr {
				if attr.Key == "public" {
					if version, ok := htmlVersions[strings.ToUpper(attr.Val)]; ok {
						return version
					}
					return "Unknown HTML version"
				}
			}
			// If attributes exist but no public identifier is found, version is unknown.
			return "Unknown HTML version"
		}
	}
	// No DOCTYPE node found in the document.
	return "No DOCTYPE found"
}

func AnalyzeHTML(doc *html.Node, baseURL string) AnalysisResult {
	var title string
	headings := map[string]int{}
	internalLinks, externalLinks, brokenLinks := 0, 0, 0
	loginForm := "Not Present"

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
						url := attr.Val
						if strings.HasPrefix(url, "http") {
							externalLinks++
							if isBrokenLink(url) {
								brokenLinks++
							}
						} else {
							internalLinks++
							fullURL := baseURL + url
							if isBrokenLink(fullURL) {
								brokenLinks++
							}
						}
					}
				}
			case "form":
				for _, attr := range n.Attr {
					if attr.Key == "action" && strings.Contains(attr.Val, "login") {
						loginForm = "Present"
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	htmlVersion := DetectHTMLVersion(doc)

	return AnalysisResult{
		Status:        "Completed",
		HTMLVersion:   htmlVersion,
		Title:         title,
		Headings:      headings,
		InternalLinks: internalLinks,
		ExternalLinks: externalLinks,
		BrokenLinks:   brokenLinks,
		LoginForm:     loginForm,
	}
}
