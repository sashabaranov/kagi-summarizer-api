package kagi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type TakeawaysResponse struct {
	URL       string  `json:"url"`
	Status    string  `json:"status"`
	Takeaways string  `json:"takeaways"`
	Elapsed   float32 `json:"elapsed"`
}

func GetTakeaways(ctx context.Context, url string) (response TakeawaysResponse, err error) {
	var reqBytes []byte
	reqBytes, err = json.Marshal(map[string]string{
		"url": url,
	})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", apiURL("/takeaways"), bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

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
