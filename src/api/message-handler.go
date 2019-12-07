package api

type MessageHandler interface {
	Handle(string) string
}

type exchangeHandler struct {
	pairCharge []GetPairData
}

func NewExchange(pairs []GetPairData) exchangeHandler {
	e := exchangeHandler{}
	e.pairCharge = pairs

	return e
}

func (e *exchangeHandler) Handle(text string) string {
	var result string
	switch text {
	case "/help":
		result = "Available commands: \n /show - Show all available trade pairs."
		break
	case "/show":
		for _, item := range e.pairCharge {
			result += item.Content() + "\n"
		}
		break
	default:
		result = "Not available command."
	}

	return result
}
