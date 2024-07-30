package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type CreateCollectionRequest struct {
	Name    string `json:"name"`
	FileUrl string `json:"fileUrl"`
}

func (app *App) HandleCreateCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateCollectionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.logger.Error("json decode err", zap.Error(err))
		app.WriteJSONErr(w, http.StatusBadRequest, fmt.Errorf("invalid json"))
		return
	}

	err = app.grpcClient.CreateCollection(ctx, req.Name, req.FileUrl)
	if err != nil {
		app.logger.Error("create collection err", zap.Error(err))
		app.WriteJSONErr(w, http.StatusInternalServerError, fmt.Errorf("failed to create collection"))
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, JSONOK{})
	app.logger.Info("create collection request succeeded")
}
