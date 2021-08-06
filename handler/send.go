package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/ilhasoft/wwcs/pkg/websocket"
	log "github.com/sirupsen/logrus"
)

var validate = validator.New()

// handle errors
var (
	ErrorConnectionClosed = errors.New("unable to send: connection closed")
	ErrorInternalError    = errors.New("unable to send: internal error")
	ErrorBadRequest       = errors.New("unable to send: bad request")
)

// SendHandler is used to receive messages from external systems
func SendHandler(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Receiving message from %q", r.Host)
	payload := websocket.IncomingPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrorBadRequest.Error()))
		return
	}

	err = validate.Struct(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrorBadRequest.Error()))
		return
	}

	c, found := Pool.Clients[payload.To]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(ErrorConnectionClosed.Error()))
		return
	}

	err = c.Send(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrorInternalError.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
