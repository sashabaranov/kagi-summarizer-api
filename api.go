package kagi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiEndpoint = "https://labs.kagi.com/v1"

func apiURL(suffix string) string {
	return fmt.Sprintf("%s%s", apiEndpoint, suffix)
}

func GetTakeways(ctx context.Context, url string) (response TakeawaysResponse, err error) {
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

func sendRequest(req *http.Request, v interface{}) (err error) {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	if v == nil {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(&v)
	return
}
