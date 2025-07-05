package handlers

import "github.com/rcarvalho-pb/checkin-go/internal/config"

func isAppPresent(app ...*config.App) bool {
	if len(app) < 1 {
		return false
	}
	return true
}
