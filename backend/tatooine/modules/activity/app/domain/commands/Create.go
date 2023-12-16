package commands

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/caching"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

const CITY_KEY = "city"
const ACTIVITY_KEY = "activity"

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

	activityKey := fmt.Sprintf("%s:%s", ACTIVITY_KEY, activityId)
	err := ccmd.Redis.Set(activityKey, jsonActivity)
	if err != nil {
		fmt.Printf("activity could not inserted to Redis %s\n", activityId)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go ccmd.addCityToRedis(activity.Location.City, jsonActivity)
	wg.Wait()

	return true, nil

}

func (ccmd *CreateCommand) addCityToRedis(city string, valueJson []byte) error {
	cityKey := fmt.Sprintf("%s:%s", CITY_KEY, city)

	return ccmd.Redis.AddMember(cityKey, valueJson)
}

//TODO: Add addmember function to handle command. it should be done async
//add Get SMEMBERS method  to get command and add a
