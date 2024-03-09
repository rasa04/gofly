package entities

type ResponseInterface interface {
	GetAnswer() string
}

type Response struct {
	Id                string
	Object            string
	Created           int
	Model             string
	SystemFingerprint string
	Choices           []Answer
	Usage             Usage
}

func (response Response) GetAnswer() string {
	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content
	}

	return ""
}
