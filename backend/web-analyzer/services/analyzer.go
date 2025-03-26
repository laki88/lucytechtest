package services

import (
	"log"
	"net/http"

	"web-analyzer/models"

	"golang.org/x/net/html"
)

const inProgress string = "In progress"

func AnalyzePage(url string) {
	AddSubmittedUrl(url)
	if analysisResults, exists := GetAnalysis(url); exists {
		if analysisResults.Status == inProgress {
			return
		}
	}
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

	analyseInProgress := models.AnalysisResult{
		Status:        inProgress,
		HTMLVersion:   "",
		Title:         "",
		Headings:      make(map[string]int),
		InternalLinks: 0,
		ExternalLinks: 0,
		BrokenLinks:   0,
		LoginForm:     "Not Present",
	}

	StoreAnalysis(url, analyseInProgress)
	analysis := models.AnalyzeHTML(doc, url)
	StoreAnalysis(url, analysis)
}
