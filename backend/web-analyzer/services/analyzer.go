package services

import (
	"log/slog"
	"net/http"

	"web-analyzer/models"

	"golang.org/x/net/html"
)

const inProgress string = "In progress"

func AnalyzePage(url string) {
	AddSubmittedUrl(url)
	if analysisResults, exists := GetAnalysis(url); exists {
		if analysisResults.Status == inProgress {
			slog.Info("Analysis already in progress", "url", url)
			return
		}
	}
	slog.Info("Starting analysis", "url", url)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to fetch URL", "url", url, "error", err)
		StoreAnalysis(url, models.AnalysisResult{Status: "Error", Message: "Failed to fetch URL"})
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		slog.Error("Failed to parse HTML", "url", url, "error", err)
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
	slog.Info("Analysis completed", "url", url, "status", analysis.Status)
}
