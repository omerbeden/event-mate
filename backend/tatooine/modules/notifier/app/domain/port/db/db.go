package db

import "github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/model"

type Database interface {
	SaveToken(model.Token) (bool, error)
	GetToken(userId string) (*model.Token, error)
}
