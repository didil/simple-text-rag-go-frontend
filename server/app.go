package server

import (
	"encoding/json"
	"net/http"

	"github.com/didil/simple-text-rag-go-frontend/services"
	"go.uber.org/zap"
)

type App struct {
	logger     *zap.Logger
	grpcClient services.GrpcClient
}

type AppOpt func(app *App)

func NewApp(opts ...AppOpt) *App {
	app := &App{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func WithLogger(logger *zap.Logger) AppOpt {
	return func(app *App) {
		app.logger = logger
	}
}

func WithGrpcClient(client services.GrpcClient) AppOpt {
	return func(app *App) {
		app.grpcClient = client
	}
}

type JSONErr struct {
	Error string `json:"error"`
}

type JSONOK struct {
}

func (app *App) WriteJSONErr(w http.ResponseWriter, statusCode int, err error) {
	jsonErr := &JSONErr{
		Error: err.Error(),
	}
	app.WriteJSONResponse(w, statusCode, jsonErr)
}

func (app *App) WriteJSONResponse(w http.ResponseWriter, statusCode int, resp any) {
	w.WriteHeader(statusCode)
	writeErr := json.NewEncoder(w).Encode(resp)
	if writeErr != nil {
		app.logger.Error("json write err", zap.Error(writeErr))
	}
}
