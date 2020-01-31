package model

const BatchTopicName = "batch"
const CountTopicName = "count"

type BatchInstruction struct {
	BatchID   string `json:"batchID"`
	AccountID string `json:"accountID"`
}
