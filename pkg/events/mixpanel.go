package events

import (
	"log"
	"os"
	"reflect"

	"github.com/ThembinkosiThemba/go-project-starter/pkg/http"
	"github.com/mixpanel/mixpanel-go"

	domain "github.com/ThembinkosiThemba/go-project-starter/internal/entity/user"
)

var (
	MIXPANEL_TOKEN = os.Getenv("MIX_PANEL_PROJECT_ID")
	mp             = mixpanel.NewApiClient(MIXPANEL_TOKEN)
)

func TrackEvents(eventName string, disctictID string, eventProps map[string]interface{}) {
	go func() {
		var ctx, cancel = http.Context()
		defer cancel()

		event := mp.NewEvent(eventName, disctictID, eventProps)
		if err := mp.Track(ctx, []*mixpanel.Event{event}); err != nil {
			log.Printf("failed to track event %s: %v", eventName, err)
		}
	}()
}

func UpdateUserProfile(user domain.USER) {
	go func() {
		ctx, cancel := http.Context()
		defer cancel()

		userProps := mixpanel.NewPeopleProperties(user.ID, map[string]interface{}{
			"$name":    user.Name,
			"$surname": user.Surname,
			"$email":   user.Email,
		})

		if err := mp.PeopleSet(ctx, []*mixpanel.PeopleProperties{userProps}); err != nil {
			log.Printf("Failed to update user profile for %s: %v", user.ID, err)
		}
	}()
}

func CreateEventProperties(data interface{}) map[string]interface{} {
	props := make(map[string]interface{})
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		props[field.Name] = value
	}

	return props
}
