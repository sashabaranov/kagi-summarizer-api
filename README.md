# kagi-summarizer-api

[![Go Reference](https://pkg.go.dev/badge/github.com/sashabaranov/kagi-summarizer-api.svg)](https://pkg.go.dev/github.com/sashabaranov/kagi-summarizer-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-gpt3)](https://goreportcard.com/report/github.com/sashabaranov/kagi-summarizer-api)


[Kagi Universal Summarizer](https://labs.kagi.com/ai/sum) API wrapper for Go

Installation:
```
go get github.com/sashabaranov/kagi-summarizer-api
```


Example usage:

```go
package main

import (
	"context"
	"fmt"
	kagi "github.com/sashabaranov/kagi-summarizer-api"
)

func main() {
	client := kagi.NewClient("your kagi token")
	response, err := client.Summarize(
		context.Background(),
		SummaryRequest{
			URL:         "your url",
			SummaryType: SummaryTypeSummary,
			Engine:      SummaryEngineAgnes,
			Cache:       true,
		},
	)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Sumarry: ", response.Data.Output)
}

```
