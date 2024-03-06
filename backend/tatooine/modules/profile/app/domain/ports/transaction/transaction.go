package transaction

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type TransactionManager interface {
	Begin(ctx context.Context) (db.Tx, error)
}
