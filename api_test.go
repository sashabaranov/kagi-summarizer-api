package kagi_test

import (
	"context"
	"os"
	"testing"

	. "github.com/sashabaranov/kagi-summarizer-api"
)

func TestAPI(t *testing.T) {
	apiToken := os.Getenv("KAGI_TOKEN")
	if apiToken == "" {
		t.Skip("Skipping testing against production Kagi API. Set KAGI_TOKEN environment variable to enable it.")
	}

	client := NewClient(apiToken)
	response, err := client.Summarize(
		context.Background(),
		SummaryRequest{
			URL:         "https://www.ussc.gov/sites/default/files/pdf/training/annual-national-training-seminar/2018/Emerging_Tech_Bitcoin_Crypto.pdf",
			SummaryType: SummaryTypeTakeaway,
			Engine:      SummaryEngineAgnes,
			Cache:       true,
		},
	)

	if err != nil {
		t.Errorf("Test request failed: %v", err)
	}
	t.Log(response)
}
