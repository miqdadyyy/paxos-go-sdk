package v2

import (
	"encoding/json"
	"fmt"
	"github.com/miqdadyyy/paxos-go-sdk/util"
	"time"
)

type QuoteExecutionItem struct {
	ID          string      `json:"id"`
	ProfileID   string      `json:"profile_id"`
	QuoteID     string      `json:"quote_id"`
	Status      string      `json:"status"`
	Market      string      `json:"market"`
	Side        string      `json:"side"`
	Price       string      `json:"price"`
	BaseAmount  string      `json:"base_amount"`
	BaseAsset   string      `json:"base_asset"`
	QuoteAmount string      `json:"quote_amount"`
	QuoteAsset  string      `json:"quote_asset"`
	CreatedAt   time.Time   `json:"created_at"`
	SettledAt   *time.Time  `json:"settled_at"`
	Metadata    interface{} `json:"metadata"`
}

type GetQuoteExecutionResponse struct {
	Items      []QuoteExecutionItem `json:"items"`
	TotalCount int                  `json:"total_count"`
}

type CreateQuoteExecutionRequest struct {
	QuoteID     string      `json:"quote_id"`
	RefID       string      `json:"ref_id"`
	BaseAmount  string      `json:"base_amount"`
	QuoteAmount string      `json:"quote_amount"`
	Metadata    interface{} `json:"metadata"`
	IdentityID  string      `json:"identity_id"`
	AccountID   string      `json:"account_id"`
}

func (v2 *PaxosV2) GetListQuoteExecution(profileID string) ([]QuoteExecutionItem, error) {
	var result GetQuoteExecutionResponse
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("profiles/%s/quote-executions", profileID)))
	if err != nil {
		return result.Items, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result.Items, err
	}

	return result.Items, nil
}

func (v2 *PaxosV2) CreateQuoteExecution(profileID string, requestBody CreateQuoteExecutionRequest) (*QuoteExecutionItem, error) {
	var result QuoteExecutionItem
	requestData := util.GenerateBodyFromStruct(requestBody)
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.
		SetFormData(requestData).
		Post(v2.generateUrlFromPath(fmt.Sprintf("profiles/%s/quote-executions", profileID)))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (v2 PaxosV2) GetQuoteExecution(profileID, quoteExecutionID string) (*QuoteExecutionItem, error) {
	var result QuoteExecutionItem
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("profiles/%s/quote-executions/%s", profileID, quoteExecutionID)))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
