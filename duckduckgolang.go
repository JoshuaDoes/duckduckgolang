package duckduckgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	AppName string
}

type QueryResult struct {
	Query string
	EscapedQuery string
	Answer string `json:"Answer"`
	Results []Result `json:"Results"`
	RelatedTopics []RelatedTopic `json:"RelatedTopics"`
	Infobox Infobox `json:"Infobox"`
	ImageIsLogo int `json:"ImageIsLogo"`
	Definition string `json:"Definition"`
	DefinitionURL string `json:"DefinitionURL"`
	DefinitionSource string `json:"DefinitionSource"`
	Heading string `json:"Heading"`
	Image string `json:"Image"`
	ImageWidth int `json:"ImageWidth"`
	ImageHeight int `json:"ImageHeight"`
	Abstract string `json:"Abstract"`
	AbstractText string `json:"AbstractText"`
	AbstractURL string `json:"AbstractURL"`
	AbstractSource string `json:"AbstractSource"`
	Type string `json:"Type"`
	AnswerType string `json:"AnswerType"`
	Entity string `json:"Entity"`
	Redirect string `json:"Redirect"`
	//Meta Meta `json:"meta"` //This will be added at a later date when all fields are known
}

type Result struct {
	Result string `json:"Result"`
	Icon Icon `json:"Icon"`
	FirstURL string `json:"FirstURL"`
	Text string `json:"Text"`
}

type RelatedTopic struct {
	Text string `json:"Text"`
	FirstURL string `json:"FirstURL"`
	Result string `json:"Result"`
	Icon Icon `json:"Icon"`
}

type Icon struct {
	URL string `json:"URL"`
	WidthStr string `json:"Width"`
	HeightStr string `json:"Height"`
	WidthInt int `json:"Width"`
	HeightInt int `json:"Height"`
}

type Infobox struct {
	Content []Content `json:"content"`
}

type Content struct {
	ValueText string `json:"value"`
	ValueObj Value `json:"value"`
	Label string `json:"label"`
	WikiOrderStr string `json:"wiki_order"`
	WikiOrderInt int `json:"wiki_order"`
}

type Value struct {
	EntityType string `json:"entity-type"`
	ID string `json:"id"`
	NumericID int `json:"numeric-id"`
}

func (c *Client) GetQueryResult(query string) (*QueryResult, error) {
	escapedQuery := url.QueryEscape(query)
	
	url := fmt.Sprintf("http://api.duckduckgo.com/?q=%s&format=json&t=%s", escapedQuery, c.AppName)
	
	data := &QueryResult{}
	data.Query = query
	data.EscapedQuery = escapedQuery
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	err = unmarshal(res, data)
	
	return data, err
}

func unmarshal(body *http.Response, target interface{}) error {
	defer body.Body.Close()
	return json.NewDecoder(body.Body).Decode(target)
}