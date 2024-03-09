package entities

type Answer struct {
	Index int
	Message Message
	LogProbs string
	FinishReason string
}
