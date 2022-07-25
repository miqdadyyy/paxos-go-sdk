package v2

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type QuoteItem struct {
	ID         string    `json:"id"`
	Market     string    `json:"market"`
	Side       string    `json:"side"`
	Price      string    `json:"price"`
	BaseAsset  string    `json:"base_asset"`
	QuoteAsset string    `json:"quote_asset"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

type GetQuoteResponse struct {
	Items []QuoteItem `json:"items"`
}

func (v2 *PaxosV2) GetQuotes(markets ...string) ([]QuoteItem, error) {
	var result GetQuoteResponse
	var query string

	if len(markets) > 0 {
		query = fmt.Sprintf("?markets=%s", strings.Join(markets, "&markets="))
	}

	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath("quotes" + query))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return result.Items, nil
}
