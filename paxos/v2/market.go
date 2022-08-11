package v2

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarketDataItem struct {
	Market         string `json:"market"`
	BaseAsset      string `json:"base_asset"`
	QuoteAsset     string `json:"quote_asset"`
	BaseIncrement  string `json:"base_increment"`
	QuoteIncrement string `json:"quote_increment"`
	TickRate       string `json:"tick_rate"`
}

type MarketRecentExecution struct {
	MatchNumber string    `json:"match_number"`
	Price       string    `json:"price"`
	Amount      string    `json:"amount"`
	ExecutedAt  time.Time `json:"executed_at"`
}

type MarketPriceAmount struct {
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

type MarketRangeTime struct {
	Begin time.Time `json:"begin"`
	End   time.Time `json:"end"`
}

type MarketDailyRecap struct {
	High                       string          `json:"high"`
	Low                        string          `json:"low"`
	Open                       string          `json:"open"`
	Volume                     string          `json:"volume"`
	VolumeWeightedAveragePrice string          `json:"volume_weighted_average_price"`
	Range                      MarketRangeTime `json:"range"`
}

type MarketOrderBook struct {
	Market string              `json:"market"`
	Asks   []MarketPriceAmount `json:"asks"`
	Bids   []MarketPriceAmount `json:"bids"`
}

type MarketTicker struct {
	Market        string            `json:"market"`
	BestBid       MarketPriceAmount `json:"best_bid"`
	BestAsk       MarketPriceAmount `json:"best_ask"`
	LastExecution MarketPriceAmount `json:"last_execution"`
	LastDay       MarketDailyRecap  `json:"last_day"`
	Today         MarketDailyRecap  `json:"today"`
	SnapshotAt    time.Time         `json:"snapshot_at"`
}

type GetMarketDataResponse struct {
	Markets []MarketDataItem `json:"markets"`
}

type GetMarketRecentExecutionResponse struct {
	Items []MarketRecentExecution `json:"items"`
}

func (v2 *PaxosV2) GetMarketData() ([]MarketDataItem, error) {
	var result GetMarketDataResponse
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath("markets"))
	if err != nil {
		return result.Markets, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result.Markets, err
	}

	return result.Markets, nil
}

func (v2 *PaxosV2) GetMarketOrderBook(marketCode string) (MarketOrderBook, error) {
	var result MarketOrderBook
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("markets/%s/order-book", marketCode)))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}

	return result, nil
}

func (v2 *PaxosV2) GetMarketRecentExecution(marketCode string) ([]MarketRecentExecution, error) {
	var result GetMarketRecentExecutionResponse
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("markets/%s/recent-executions", marketCode)))
	if err != nil {
		return result.Items, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return result.Items, nil
}

func (v2 *PaxosV2) GetMarketTicker(marketCode string) (MarketTicker, error) {
	var result MarketTicker
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("markets/%s/ticker", marketCode)))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}

	return result, nil
}
