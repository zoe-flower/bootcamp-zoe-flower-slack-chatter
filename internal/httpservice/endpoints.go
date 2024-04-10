package httpservice

import (
	"net/http"
	"time"

	"github.com/flypay/go-kit/v4/pkg/log"
	"github.com/labstack/echo/v4"

	"github.com/flypay/go-kit/v4/pkg/projections"

	"github.com/google/uuid"
)

type HTTPHandler struct {
	Logger log.Logger

	DefaultDatabase projections.ReadWriter
}

func (h HTTPHandler) UserGet(ctx echo.Context, userId string) error {
	userRecord := UserRecord{}

	// Reads the data from DynamoDB
	if err := h.DefaultDatabase.Read(ctx.Request().Context(), userId, &userRecord); err != nil {
		h.Logger.WithContext(ctx.Request().Context()).Errorf("error reading from DynamoDB: %+v", userRecord)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if userRecord.Identifier != "" {
		return ctx.JSON(http.StatusOK, userRecord)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

func (h HTTPHandler) UserAdd(ctx echo.Context) error {
	user := User{}

	if err := ctx.Bind(&user); err != nil {
		h.Logger.WithContext(ctx.Request().Context()).Errorf("error binding the request body %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Check that the required fields are provided
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide the required fields")
	}

	now := time.Now().Format(time.RFC3339)

	userRecord := UserRecord{
		Identifier:   uuid.New().String(),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		DateCreated:  now,
		DateModified: now,
	}

	// Writes the data to DynamoDB
	if err := h.DefaultDatabase.Write(ctx.Request().Context(), userRecord, nil); err != nil {
		h.Logger.WithContext(ctx.Request().Context()).Errorf("error writing to DynamoDB: %+v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusCreated, userRecord)
}

type UserRecord struct {
	Identifier string // Needs to be sentence case for DynamoDB

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	DateCreated  string `json:"date_created,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
}
