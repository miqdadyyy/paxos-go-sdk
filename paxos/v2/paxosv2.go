package v2

import "github.com/miqdadyyy/paxos-go-sdk/paxos"

type PaxosV2 struct {
	PaxosClient *paxos.PaxosClient
}

func NewV2(paxosClient *paxos.PaxosClient) PaxosV2 {
	paxosClient.BaseURL = paxosClient.BaseURL + "v2/"
	return PaxosV2{
		PaxosClient: paxosClient,
	}
}

func (v2 *PaxosV2) generateUrlFromPath(path string) string {
	return v2.PaxosClient.BaseURL + path
}
