package services

import "web-analyzer/models"

var analysisResults = make(map[string]models.AnalysisResult)
var submittedUrls []string

func StoreAnalysis(url string, result models.AnalysisResult) {
	analysisResults[url] = result
}

func GetAnalysis(url string) (models.AnalysisResult, bool) {
	result, exists := analysisResults[url]
	return result, exists
}

func GetSubmittedUrls() []string {
	return submittedUrls
}

func AddSubmittedUrl(url string) {
	submittedUrls = append(submittedUrls, url)
}
