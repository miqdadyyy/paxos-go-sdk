package v2

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CreateProfileRequest struct {
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
}

type GetProfileItemsResponse struct {
	Items []ProfileItem `json:"items"`
}

type GetProfileBalanceItemsResponse struct {
	Items []ProfileBalanceItem `json:"items"`
}

type ProfileItem struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Type     string `json:"type"`
}

type ProfileBalanceItem struct {
	Asset     string `json:"asset"`
	Available string `json:"available"`
	Trading   string `json:"trading"`
}

func (v2 *PaxosV2) GetProfiles() ([]ProfileItem, error) {
	var result GetProfileItemsResponse
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath("profiles"))
	if err != nil {
		return result.Items, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result.Items, err
	}

	return result.Items, nil
}

func (v2 *PaxosV2) GetProfileByID(profileID string) (ProfileItem, error) {
	var result ProfileItem
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath("profiles/" + profileID))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}

	return result, nil
}

func (v2 *PaxosV2) GetProfileBalances(profileID string, assets ...string) ([]ProfileBalanceItem, error) {
	var result GetProfileBalanceItemsResponse
	var query string

	if len(assets) > 0 {
		query = fmt.Sprintf("?assets=%s", strings.Join(assets, "&assets="))
	}

	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("profiles/%s/balances%s", profileID, query)))
	if err != nil {
		return result.Items, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result.Items, err
	}

	return result.Items, nil
}

func (v2 *PaxosV2) GetProfileBalance(profileID string, asset string) (ProfileBalanceItem, error) {
	var result ProfileBalanceItem
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.Get(v2.generateUrlFromPath(fmt.Sprintf("profiles/%s/balances/%s", profileID, asset)))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}

	return result, nil
}

func (v2 *PaxosV2) CreateProfile(nickname, description string) (ProfileItem, error) {
	var result ProfileItem
	client := v2.PaxosClient.GenerateClientRequest()
	resp, err := client.
		SetFormData(map[string]string{
			"nickname":    nickname,
			"description": description,
		}).
		Post(v2.generateUrlFromPath("profiles"))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}

	return result, nil
}
