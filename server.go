package main

import (
	_ "embed"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"golang.org/x/time/rate"
)

//go:embed config/kid
var kid []byte

//go:embed config/iss
var iss []byte

//go:embed config/key.p8
var key []byte

func main() {
	authKey, err := token.AuthKeyFromBytes(key)
	if err != nil {
		log.Fatal("token error:", err)
	}
	token := &token.Token{
		AuthKey: authKey,
		KeyID:   string(kid),
		TeamID:  string(iss),
	}
	notificationsDev := make(chan *apns2.Notification, 1)
	notifications := make(chan *apns2.Notification, 500)
	topic := []string{"cn.sshterm.pro", "cn.sshterm.free", "cn.sshterm.dev"}
	clientDev := apns2.NewTokenClient(token).Development()
	client := apns2.NewTokenClient(token).Production()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.BodyLimit("64K"))
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(config))
	e.POST("/apn_push", func(c echo.Context) error {
		data := new(Body)
		if err = c.Bind(data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var valid bool
		for _, v := range topic {
			if strings.Compare(data.Topic, v) == 0 {
				valid = true
				break
			}
		}
		if !valid {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid topic")
		}
		data.Notification.APS.Sound = "default"
		notification := &apns2.Notification{
			DeviceToken: data.Token,
			Topic:       data.Topic,
			Payload:     data.Notification,
			PushType:    apns2.PushTypeAlert,
			Priority:    data.Priority,
		}

		if data.Dev {
			notificationsDev <- notification
		} else {
			notifications <- notification
		}
		return c.NoContent(http.StatusOK)
	})

	go func() {
		for n := range notifications {
			client.Push(n)
		}
	}()
	go func() {
		for n := range notificationsDev {
			clientDev.Push(n)
		}
	}()
	e.Logger.Fatal(e.Start(":8080"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type (
	Body struct {
		Dev          bool         `json:"dev"`
		Token        string       `json:"token" validate:"required"`
		Topic        string       `json:"topic" validate:"required"`
		Notification Notification `json:"notification" validate:"required"`
		Priority     int          `json:"priority" validate:"gte=1,lte=10"`
	}
	APS struct {
		Alert Alert  `json:"alert" validate:"required"`
		Sound string `json:"sound"`
	}
	Alert struct {
		Title    string `json:"title" validate:"required"`
		Subtitle string `json:"subtitle" validate:"required"`
		Body     string `json:"body" validate:"required"`
	}
	Notification struct {
		APS APS `json:"aps" validate:"required"`
	}
)
