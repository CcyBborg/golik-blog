package commentshandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CcyBborg/golik-blog/internal/api/schema"

	"github.com/CcyBborg/golik-blog/internal/api/utils"

	"github.com/CcyBborg/golik-blog/internal/models"
	"github.com/gorilla/mux"
)

type keeper interface {
	GetComments(postID int64) ([]models.Comment, error)
	InsertComment(comment *models.Comment) error
}

type Handler struct {
	keeper keeper
}

func New(keeper keeper) *Handler {
	return &Handler{keeper: keeper}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseInt(vars["postID"], 10, 64)
	if err != nil {
		utils.WriteInvalidParams(w)
		return
	}

	if r.Method == http.MethodGet {
		comments, err := h.keeper.GetComments(postID)
		if err != nil {
			fmt.Println(err)
			utils.WriteInternalError(w)
			return
		}

		schemaComments := make([]schema.Comment, len(comments))
		for i, comment := range comments {
			schemaComments[i] = schema.ConvertComment(comment)
		}

		jsonResponse, err := json.Marshal(schemaComments)
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		utils.WriteJSON(w, jsonResponse)
	} else if r.Method == http.MethodPost {
		userID := int64(1) // Remove after JWT

		if err := r.ParseForm(); err != nil {
			utils.WriteInvalidParams(w)
			return
		}

		content := r.Form.Get("content")
		if content == "" {
			utils.WriteInvalidParams(w)
			return
		}

		comment := models.Comment{
			Author: models.User{
				ID: userID,
			},
			PostID:  postID,
			Content: content,
		}

		if err = h.keeper.InsertComment(&comment); err != nil {
			utils.WriteInternalError(w)
			return
		}

		jsonResponse, err := json.Marshal(schema.ConvertComment(comment))
		if err != nil {
			utils.WriteInternalError(w)
		}

		utils.WriteJSON(w, jsonResponse)
	}
}
