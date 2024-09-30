package models

type Transaction struct {
	Id              int64  `json:"id"`
	UserDocument    string `json:"user_document"`
	CreditCradToken string `json:"credit_card_token"`
	Value           int64  `json:"value"`
}
