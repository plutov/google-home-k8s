package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/plutov/google-home-k8s/pkg/commands"
	"github.com/plutov/google-home-k8s/pkg/dialogflow"
	log "github.com/sirupsen/logrus"
)

// UnknownErrorMsg is sent to user when user
const UnknownErrorMsg = "Something went wrong with your request. Please try again."

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewRouter returns new router
func NewRouter(h *Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == os.Getenv("API_KEY"), nil
	}))

	e.GET("/", h.dialogflowHandler)
	e.POST("/", h.dialogflowHandler)

	return e
}

func (h *Handler) dialogflowHandler(c echo.Context) error {
	req := new(dialogflow.Request)
	if err := c.Bind(req); err != nil {
		log.WithError(err).Error("unable to parse webhook request")
		return c.JSON(http.StatusOK, dialogflow.GenerateWebhookResponse(false, UnknownErrorMsg))
	}

	log.WithField("req", *req).Debug("webhook request")

	userSession := h.UserSessionManager.GetUserSession(req)

	msg, err := commands.Execute(h.GKEClient, userSession, req)
	if err != nil {
		log.WithError(err).Error("unable to execute command")
		return c.JSON(http.StatusOK, dialogflow.GenerateWebhookResponse(false, UnknownErrorMsg))
	}

	h.UserSessionManager.SaveUserSession(*userSession)

	return c.JSON(http.StatusOK, dialogflow.GenerateWebhookResponse(true, msg))
}
