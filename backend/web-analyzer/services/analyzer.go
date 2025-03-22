package services

import (
	"log"
	"net/http"

	"web-analyzer/models"

	"golang.org/x/net/html"
)

func AnalyzePage(url string) {
	AddSubmittedUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v", url, err)
		StoreAnalysis(url, models.AnalysisResult{Status: "Error", Message: "Failed to fetch URL"})
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		StoreAnalysis(url, models.AnalysisResult{Status: "Error", Message: "Failed to parse HTML"})
		return
	}

	analysis := models.AnalyzeHTML(doc, url)
	StoreAnalysis(url, analysis)
}
