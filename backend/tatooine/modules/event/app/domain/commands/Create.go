package commands

import (
	cacheadapter "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/caching"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
)

type CreateCommand struct {
	Event   model.Event
	Repo    repo.EventRepository
	LocRepo repo.LocationRepository
	Redis   caching.Cache
}

func (ccmd *CreateCommand) Handle() (bool, error) {
	err := cacheadapter.Push(ccmd.Event.Location.City, ccmd.Event, ccmd.Redis)
	if err != nil {
		return false, err
	}

	_, errCreate := ccmd.Repo.Create(ccmd.Event)
	if errCreate != nil {
		return false, errCreate
	}
	_, errLoc := ccmd.LocRepo.Create(&ccmd.Event.Location)
	if err != nil {
		return false, errLoc
	}
	// if isAddedToDB {
	// 	// input := &sns.PublishInput{
	// 	// 	Message:  aws.String(fmt.Sprintf("event created with id : %d", ccmd.Event.ID)),
	// 	// 	TopicArn: aws.String("topic"),
	// 	// }

	// 	//_, err := snsadapter.PublishMessage(context.Background(), &snsadapter.SNSAdapter{Topic: "topic_test"}, input)
	// 	if err != nil {
	// 		return false, err // todo db ye eklendi aslında oluyor, transaction ını rollback yapmak lazım
	// 	}

	// }

	return true, nil

}

type CreateCacheCommand struct {
	Redis caching.Cache
	Key   string
	Posts []model.Event
}

func (uc *CreateCacheCommand) Handle() (bool, error) {
	err := cacheadapter.Push(uc.Key, uc.Posts, uc.Redis)
	if err != nil {
		return false, err
	}

	return true, nil
}
