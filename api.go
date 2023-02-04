package kagi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiEndpoint = "https://labs.kagi.com/v1"

func apiURL(suffix string) string {
	return fmt.Sprintf("%s%s", apiEndpoint, suffix)
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
