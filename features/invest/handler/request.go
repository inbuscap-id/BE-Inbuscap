package handler

type InvestmentRequest struct {
	Proposal_id uint `json:"proposal_id"`
	Amount      int  `json:"amount"`
}
