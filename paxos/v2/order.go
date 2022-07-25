package v2

import (
	"encoding/json"
	"github.com/miqdadyyy/paxos-go-sdk/constant"
	"github.com/miqdadyyy/paxos-go-sdk/util"
	"time"
)

type GetListExecutionQueryRequest struct {
	ProfileID  string `json:"profile_id"`
	AccountID  string `json:"account_id"`
	OrderID    string `json:"order_id"`
	PageCursor string `json:"page_cursor"`
	Limit      int    `json:"limit"`
}

type GetListOrderQueryRequest struct {
	ProfileID  string `json:"profile_id"`
	AccountID  string `json:"account_id"`
	Market     string `json:"market"`
	Status     string `json:"status"`
	RefIds     string `json:"ref_ids"`
	PageCursor string `json:"page_cursor"`
	Limit      int    `json:"limit"`
}

type ListExecutionItem struct {
	ExecutionID      string    `json:"execution_id"`
	OrderID          string    `json:"order_id"`
	ExecutedAt       time.Time `json:"executed_at"`
	Market           string    `json:"market"`
	Side             string    `json:"side"`
	Amount           string    `json:"amount"`
	Price            string    `json:"price"`
	Commission       string    `json:"commission"`
	CommissionAsset  string    `json:"commission_asset"`
	Rebate           string    `json:"rebate"`
	RebateAsset      string    `json:"rebate_asset"`
	GrossTradeAmount string    `json:"gross_trade_amount"`
}

type GetListExecutionResponse struct {
	Items          []ListExecutionItem `json:"items"`
	NextPageCursor string              `json:"next_page_cursor"`
}

type ListOrderItem struct {
	ID                         string `json:"id"`
	ProfileID                  string `json:"profile_id"`
	RefID                      string `json:"ref_id"`
	Status                     string `json:"status"`
	Market                     string `json:"market"`
	Side                       string `json:"side"`
	Type                       string `json:"type"`
	Price                      string `json:"price"`
	BaseAmount                 string `json:"base_amount"`
	QuoteAmount                string `json:"quote_amount"`
	Metadata                   string `json:"metadata"`
	AmountFilled               string `json:"amount_filled"`
	VolumeWeightedAveragePrice string `json:"volume_weighted_average_price"`
	TimeInForce                string `json:"time_in_force"`
}

type GetListOrderResponse struct {
	Items          []ListOrderItem `json:"items"`
	NextPageCursor string          `json:"next_page_cursor"`
}

type CreateOrderRequestData struct {
	Side        string      `json:"side"`
	Market      string      `json:"market"`
	Type        string      `json:"type"`
	Price       string      `json:"price"`
	BaseAmount  string      `json:"base_amount"`
	QuoteAmount string      `json:"quote_amount"`
	Metadata    interface{} `json:"metadata"`
}

type CreateOrderResponseData struct {
	ID                         string      `json:"id"`
	ProfileID                  string      `json:"profile_id"`
	RefID                      string      `json:"ref_id"`
	Status                     string      `json:"status"`
	Market                     string      `json:"market"`
	Side                       string      `json:"side"`
	Type                       string      `json:"type"`
	Price                      string      `json:"price"`
	BaseAmount                 string      `json:"base_amount"`
	QuoteAmount                string      `json:"quote_amount"`
	Metadata                   interface{} `json:"metadata"`
	AmountFilled               string      `json:"amount_filled"`
	VolumeWeightedAveragePrice string      `json:"volume_weighted_average_price"`
	TimeInForce                string      `json:"time_in_force"`
}

func (v2 *PaxosV2) GetListExecutions(requestData *GetListExecutionQueryRequest) ([]ListExecutionItem, error) {
	var result GetListExecutionResponse
	var query string

	if requestData != nil {
		query = util.GenerateQueryFromStruct(requestData)
	}

	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath("executions") + query)
	if err != nil {
		return result.Items, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result.Items, err
	}

	return result.Items, nil
}

func (v2 *PaxosV2) GetListOrders(requestData *GetListExecutionQueryRequest) ([]ListOrderItem, error) {
	var query string
	var result GetListOrderResponse

	if requestData != nil {
		query = util.GenerateQueryFromStruct(requestData)
	}

	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath("orders") + query)
	if err != nil {
		return result.Items, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result.Items, err
	}

	return result.Items, nil
}

func (v2 PaxosV2) CreateOrder(requestData CreateOrderRequestData) {
	requestBody := make(map[string]interface{})

	switch requestData.Type {
	case constant.OrderTypeLimit:
		requestBody["side"] = requestData.Side
		requestBody["market"] = requestData.Market
		requestBody["type"] = requestData.Type
		requestBody["price"] = requestData.Price
		requestBody["base_amount"] = requestData.BaseAmount
	case constant.OrderTypeMarket:
		requestBody["side"] = requestData.Side
		requestBody["market"] = requestData.Market
		requestBody["type"] = requestData.Type
		if requestData.Side == constant.OrderSideBuy {
			requestBody["quote_amount"] = requestData.QuoteAmount
		} else {
			requestBody["base_amount"] = requestData.BaseAmount
		}
	}

	if requestData.Metadata != nil {
		requestBody["metadata"] = requestData.Metadata
	}

}
