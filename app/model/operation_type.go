package model

type OperationType int

const (
	CASH_PURCHASE     OperationType = 1
	PURCHASE_PARCELED OperationType = 2
	TAKE_OUT          OperationType = 3
	PAYMENT           OperationType = 4
)
