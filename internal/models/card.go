package models

type Card struct {
	ID        string
	Owner     string
	Bank      string
	Limit     int
	Balance   int
	InDispute bool
}
