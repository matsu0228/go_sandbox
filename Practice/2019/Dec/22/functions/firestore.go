package fsfunc

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	translate "cloud.google.com/go/translate/apiv3"
	firebase "firebase.google.com/go"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	Fields     MyData    `json:"fields"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}

// MyData represents a value from Firestore. The type definition depends on the
// format of your database.
type MyData struct {
	Original struct {
		StringValue string `json:"stringValue"`
	} `json:"original"`
}

// GCLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GCLOUD_PROJECT")

// client is a Firestore client, reused between function invocations.
var client *firestore.Client

func init() {
	conf := &firebase.Config{ProjectID: projectID}
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

// MakeUpperCase is triggered by a change to a Firestore document. It updates
// the `original` value of the document to upper case.
func MakeUpperCase(ctx context.Context, e FirestoreEvent) error {
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	collection := pathParts[0]
	doc := strings.Join(pathParts[1:], "/")

	curValue := e.Value.Fields.Original.StringValue
	newValue := strings.ToUpper(curValue)
	if curValue == newValue {
		log.Printf("%q is already upper case: skipping", curValue)
		return nil
	}
	log.Printf("Replacing value: %q -> %q", curValue, newValue)

	data := map[string]string{"original": newValue}
	_, err := client.Collection(collection).Doc(doc).Set(ctx, data)
	if err != nil {
		return fmt.Errorf("Set: %v", err)
	}
	return nil
}

// Translate is triggerd by a change
func Translate(ctx context.Context, e FirestoreEvent) error {
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	collection := pathParts[0]
	doc := strings.Join(pathParts[1:], "/")

	srcText := e.Value.Fields.Original.StringValue

	sourceLang := "ja"
	targetLang := "en"
	transed, err := translateText(sourceLang, targetLang, []string{srcText})
	if err != nil {
		return err
	}
	data := map[string]string{"translated": strings.Join(transed, "\n")}
	_, err = client.Collection(collection).Doc(doc).Set(ctx, data, firestore.MergeAll)
	return err
}

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
