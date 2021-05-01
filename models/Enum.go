package models

type banks string
type investments string

const (
	HDFC banks = "HDFC Bank Private Limited"
	SBI        = "State Bank of India"
)

const (
	STOCK investments = "Stocks"
	FD                = "Fixed Deposit"
	MF                = "Mutual Funds"
	RD                = "Recurring Deposit"
)
