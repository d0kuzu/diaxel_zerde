package followup

import (
	"context"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/twilio"
)

type Listener struct {
	dbClient *db.Client
	bot      *twilio.Client
}

func NewListener(dbClient *db.Client, bot *twilio.Client) *Listener {
	return &Listener{
		dbClient: dbClient,
		bot:      bot,
	}
}

func (l *Listener) Start(ctx context.Context) {
	// Follow-up logic is disabled.
}
