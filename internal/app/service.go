package service

import (
	"fmt"

	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"google.golang.org/protobuf/proto"

	// "github.com/flypay/events/pkg/flyt"

	"github.com/flypay/events/pkg/bootcamp"
	"github.com/flypay/go-kit/v4/pkg/projections"
)

// only cosumes events and add adds to db
// passes db to handler
// consumes two events via a map
func RunService(
	defaultDatabase projections.ReadWriter,
	consumer eventbus.Consumer,
) error {
	eh := eventHandler{db: defaultDatabase}
	err := consumer.OnAll(map[proto.Message]eventbus.EventHandler{
		&bootcamp.BirthdayToday{}: eh.findBirthday,
		&bootcamp.UserCreated{}:   eh.createUser,
	})
	if err != nil {
		fmt.Println("Couldn't consume:", err)
	}
	defer consumer.Listen()

	// consumer.On(...)
	// defer consumer.Listen()

	// NOTE: For each SNS topic you produce or consume you will need to modify
	// the service.json file to specify the topic in the 'eventbus.consumes_topics' list.
	// E.g., for &flyt.UserCreated, add "flyt-user-created" (dashes instead of dots in "flyt.user.created" from the proto) to 'eventbus.consumes_topics' in service.json

	return nil
}
