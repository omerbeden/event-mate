package db

import "github.com/omerbeden/event-mate/backend/notifierservice/internal/app/domain/model"

type Database interface {
	SaveToken(model.Token) (bool, error)
	GetToken(userId string) (*model.Token, error)
}
