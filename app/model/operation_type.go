package model

type OperationType int

// Declare related constants for each direction starting with index 1
const (
	CASH_PURCHASE OperationType = 1
	PURCHASE_PARCELED OperationType = 2
	TAKE_OUT OperationType = 3
	PAYMENT OperationType = 4
)