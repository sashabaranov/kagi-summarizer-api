package kagi

import (
	"context"
	"net/http"
	"time"
)

type SummaryResponse struct {
	URL       string  `json:"url"`
	Status    string  `json:"status"`
	Summary   string  `json:"summary"`
	WordStats string  `json:"word_stats"`
	Elapsed   float32 `json:"elapsed"`
}

func GetSummary(ctx context.Context, url string) (response SummaryResponse, err error) {
	req, err := http.NewRequest("GET", apiURL("/summary_status"), nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("url", url)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	for i := 0; i < 20; i++ {
		err = sendRequest(req, &response)
		if err != nil {
			return
		}

		if response.Status == "completed" {
			return
		}

		time.Sleep(2 * time.Second)
	}
	return
}
