package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/caching"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type CreateCommand struct {
	Activity     model.Activity
	ActivityRepo repo.ActivityRepository
	LocRepo      repo.LocationRepository
	Redis        caching.Cache
}

func (ccmd *CreateCommand) Handle() (bool, error) {

	activity, errCreate := ccmd.ActivityRepo.Create(ccmd.Activity)
	if errCreate != nil {
		return false, errCreate
	}

	_, errLoc := ccmd.LocRepo.Create(&activity.Location)
	if errLoc != nil {
		return false, errLoc
	}

	activityId := strconv.FormatInt(activity.ID, 10)
	jsonActivity, errMarshall := json.Marshal(ccmd.Activity)
	if errMarshall != nil {
		return false, errMarshall
	}

	err := ccmd.Redis.Set(activityId, jsonActivity)
	if err != nil {
		fmt.Printf("activity could not inserted to Redis %s\n", activityId)
	}

	return true, nil

}
