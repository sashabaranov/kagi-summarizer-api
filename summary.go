package kagi

import (
	"context"
	"errors"
	"net/http"
)

type (
	SummaryEngine string
	SummaryType   string
)

const (
	SummaryEngineAgnes  SummaryEngine = "agnes"
	SummaryEngineDaphne SummaryEngine = "daphne"
	SummaryEngineMuriel SummaryEngine = "muriel"

	SummaryTypeSummary  SummaryType = "summary"
	SummaryTypeTakeaway SummaryType = "takeaway"
)

var (
	ErrUrlOrTextNeedsToBeSet = errors.New("either URL or Text field should be set in SummaryRequest")
)

type SummaryRequest struct {
	URL  string `json:"url,omitempty"`
	Text string `json:"text,omitempty"`

	Engine         SummaryEngine `json:"engine,omitempty"`
	SummaryType    SummaryType   `json:"summary_type,omitempty"`
	TargetLanguage string        `json:"target_language,omitempty"`
	Cache          bool          `json:"cache"`
}

type ResponseMeta struct {
	ID   string `json:"id"`
	Node string `json:"node"`
	Ms   uint64 `json:"ms"`
}

type SummaryData struct {
	Output string `json:"output"`
	Tokens uint64 `json:"tokens"`
}

type SummaryResponse struct {
	Meta ResponseMeta `json:"meta"`
	Data SummaryData  `json:"data"`
}

func (r SummaryRequest) validate() error {
	if r.URL != "" && r.Text != "" {
		return ErrUrlOrTextNeedsToBeSet
	}

	if r.URL == "" && r.Text == "" {
		return ErrUrlOrTextNeedsToBeSet
	}

	return nil
}

func (c *Client) Summarize(ctx context.Context, summaryRequest SummaryRequest) (response SummaryResponse, err error) {
	err = summaryRequest.validate()
	if err != nil {
		return
	}

	req, err := c.buildRequest(ctx, http.MethodPost, c.apiURL("/summarize"), summaryRequest)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
