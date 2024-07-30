package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type GetAnswerRequest struct {
	CollectionName string `json:"collectionName"`
	Question       string `json:"question"`
}

type GetAnswerReply struct {
	Text string `json:"text"`
}

func (app *App) HandleGetAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req GetAnswerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.logger.Error("json decode err", zap.Error(err))
		app.WriteJSONErr(w, http.StatusBadRequest, fmt.Errorf("invalid json"))
		return
	}

	answer, err := app.grpcClient.GetAnswer(ctx, req.CollectionName, req.Question)
	if err != nil {
		app.logger.Error("get answer err", zap.Error(err))
		app.WriteJSONErr(w, http.StatusInternalServerError, fmt.Errorf("failed to get answer"))
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, GetAnswerReply{Text: answer})
	app.logger.Info("get answer request succeeded")
}
