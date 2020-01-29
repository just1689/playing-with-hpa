package model

const BatchTopicName = "batch"

type BatchInstruction struct {
	BatchID   string `json:"batchID"`
	AccountID string `json:"accountID"`
}
