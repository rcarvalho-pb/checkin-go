package handler

import "github.com/rcarvalho-pb/checkin-go/internal/storage"

type ParticipantHandler struct {
	p storage.ParticipantStorage
}

func NewParticipantHandler(p storage.ParticipantStorage) *ParticipantHandler {
	return &ParticipantHandler{
		p,
	}
}
