package httpservice

import (
	"github.com/flypay/go-kit/v4/pkg/log"

	"github.com/flypay/go-kit/v4/pkg/projections"
)

type HTTPHandler struct {
	Logger log.Logger

	DefaultDatabase projections.ReadWriter
}

type UserRecord struct {
	Identifier string // Needs to be sentence case for DynamoDB

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	DateCreated  string `json:"date_created,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
}
