package main

import (
	"context"
	"fmt"
	"log"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

var (
	projectID = os.Getenv("GCLOUD_PROJECT")
)

// translateText translates input text and returns translated text.
func translateText(sourceLang, targetLang string, texts []string) ([]string, error) {

	translatedTexts := []string{}
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return translatedTexts, err
	}
	defer client.Close()

	// https://cloud.google.com/translate/docs/languages
	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           texts,
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return translatedTexts, err
	}

	for _, trans := range resp.GetTranslations() {
		translatedTexts = append(translatedTexts, trans.GetTranslatedText())
	}

	return translatedTexts, nil
}

func main() {

	sourceLang := "ja"
	targetLang := "en"
	text := []string{"愛媛ＦＣレディースの仲松 叶実選手が松本 苑佳選手と交代です！"}
	texts, err := translateText(sourceLang, targetLang, text)
	if err != nil {
		log.Printf("err: %v", err)
	}

	for _, trans := range texts {
		fmt.Println(trans)
	}
}
