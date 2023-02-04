# kagi-summarizer-api

[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/sashabaranov/kagi-summarizer-api)
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
	ctx := context.Background()
	url := "your url"
	fmt.Println(kagi.GetSummary(ctx, url))
	fmt.Println(kagi.GetTakeAways(ctx, url))
}

```
