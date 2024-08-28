package tokens

type Payload struct {
	UserID string
}

func NewPayload(userId string) *Payload {
	return &Payload{
		UserID: userId,
	}
}
