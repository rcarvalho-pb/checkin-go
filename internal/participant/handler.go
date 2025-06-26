package participant

type ParticipantHandler struct {
	participantStorage
}

func NewParticipantHandler(p participantStorage) *ParticipantHandler {
	return &ParticipantHandler{
		p,
	}
}
