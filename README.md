# duckduckgolang
Single file library for the DuckDuckGo API written in Golang without extra dependencies

# Installing
`go get github.com/JoshuaDoes/duckduckgolang`

# Example
```go
package main

import "fmt"
import "github.com/JoshuaDoes/duckduckgolang"

var duckduckgoClient *duckduckgo.Client

func main() {
	duckduckgoClient = &duckduckgo.Client{AppName:"your app name here"}
	queryResult, err := duckduckgoClient.GetQueryResult("What is DuckDuckGo?")
	if err != nil {
		fmt.Println("Error: " + fmt.Sprintf("%v", err))
	} else {
		fmt.Println("Query: " + queryResult.Query)
		fmt.Println("AbstractText: " + queryResult.AbstractText)
	}
}
```
### Output

```
> go run main.go
Query: What is DuckDuckGo?
AbstractText: DuckDuckGo is an Internet search engine that emphasizes protecting searchers' privacy and avoiding the filter bubble of personalized search results. DuckDuckGo distinguishes itself from other search engines by not profiling its users and by deliberately showing all users the same search results for a given search term. DuckDuckGo emphasizes returning the best results, rather than the most results, and generates those results from over 400 individual sources, including key crowdsourced sites such as Wikipedia, and other search engines like Bing, Yahoo!, Yandex, and Yummly.
```