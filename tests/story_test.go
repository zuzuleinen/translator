package tests

import (
	"testing"
	"github.com/zuzuleinen/translator/application"
	"github.com/zuzuleinen/translator/db"
	"github.com/zuzuleinen/translator/values"
)

//As a user I want to be able to store translations for multiple language-combinations.
func TestStoreRequest(t *testing.T) {
	client := application.NewClient(db.NewStore())

	storeReq := values.StoreRequest{
		First:  values.NewWord("de", "hund"),
		Second: values.NewWord("en", "dog"),
	}
	resp := client.StoreRequest(storeReq)

	if resp.Body != application.StatusOK {
		t.Errorf("Response is incorect. Got %s, Want %s", resp.Body, application.StatusOK)
	}
}

//As a user I want to be able to get a previous stored translation
func TestGetPreviousStoredTranslation(t *testing.T) {
	client := application.NewClient(db.NewStore())

	storeReq := values.StoreRequest{
		First:  values.NewWord("de", "hund"),
		Second: values.NewWord("en", "dog"),
	}

	client.StoreRequest(storeReq)

	resp := client.GetRequest(values.GetRequest{Word: values.NewWord("de", "hund"), InLanguage: "en"})

	if resp.Body != "dog" {
		t.Errorf("Response is incorect. Got %s, Want %s", resp.Body, "dog")
	}
}

//As a user I want to get a guess of a translation using other translations if possible
func TestGuessTranslation(t *testing.T) {
	client := application.NewClient(db.NewStore())

	germanToSpanish := values.StoreRequest{
		First:  values.NewWord("de", "katze"),
		Second: values.NewWord("es", "gato"),
	}

	spanishToEnglish := values.StoreRequest{
		First:  values.NewWord("es", "gato"),
		Second: values.NewWord("en", "cat"),
	}

	englishToFrench := values.StoreRequest{
		First:  values.NewWord("en", "cat"),
		Second: values.NewWord("fr", "chat"),
	}

	client.StoreRequest(germanToSpanish)
	client.StoreRequest(spanishToEnglish)
	client.StoreRequest(englishToFrench)

	resp := client.GetRequest(values.GetRequest{Word: values.NewWord("de", "katze"), InLanguage: "en"})

	if resp.Body != "cat" {
		t.Errorf("Response is incorect. Got %s, Want %s", resp.Body, "cat")
	}

	resp = client.GetRequest(values.GetRequest{Word: values.NewWord("de", "Katze"), InLanguage: "fr"})

	if resp.Body != "chat" {
		t.Errorf("Response is incorect. Got %s, Want %s", resp.Body, "chat")
	}
}
