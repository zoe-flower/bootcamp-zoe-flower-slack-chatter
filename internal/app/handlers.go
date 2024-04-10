package service

import (
	"context"

	"github.com/flypay/events/pkg/bootcamp"
	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"github.com/flypay/go-kit/v4/pkg/log"
	"github.com/flypay/go-kit/v4/pkg/projections"
)

// contains properties requires by handler function
type eventHandler struct {
	db projections.ReadWriter
}

// structure to take data in correct format
type SlackRecordProjection struct {
	Identifier string // id
}

// recieves the handler/db
// instantiates a projection in the correct format to hold db data
// adds incoming data to fb object
// reads db with fb and adds required data to projection object
func (e eventHandler) findBirthday(ctx context.Context, event any, headers ...eventbus.Header) error {
	logger := log.WithContext(ctx)
	result := BirthdayRecordProjection{}
	fb := event.(*bootcamp.BirthdayToday)
	logger.Debugf("findBirthday recieved %v", event)
	err := e.db.Read(ctx, fb.User, &result)
	if err != nil {
		return err
	}
	logger.Debugf("HAPPY BIRTHDAY %v", result.SlackHandle)
	return nil
}

type BirthdayRecordProjection struct {
	Identifier  string // Id
	SlackHandle string // Slack
}

func (e eventHandler) createUser(ctx context.Context, event any, headers ...eventbus.Header) error {
	logger := log.WithContext(ctx)
	uc := event.(*bootcamp.UserCreated)
	up := BirthdayRecordProjection{
		Identifier:  uc.Id,
		SlackHandle: uc.SlackHandle,
	}
	logger.Debugf("createUser recieved %v", event)
	err := e.db.Write(ctx, up, nil)
	if err != nil {
		return err
	}
	return nil
}
