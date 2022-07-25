package v2

type GetAccountListQueryRequest struct {
	PageCursor string `json:"page_cursor"`
	Order      string `json:"order"`
	OrderBy    string `json:"order_by"`
	Limit      int    `json:"limit"`
	IdentityID string `json:"identity_id"`
}

func (v2 *PaxosV2) GetAccountList(query *GetAccountListQueryRequest) {

}
