package mypostshandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CcyBborg/golik-blog/internal/api/schema"
	"github.com/CcyBborg/golik-blog/internal/api/utils"
	"github.com/CcyBborg/golik-blog/internal/models"
)

type keeper interface {
	GetUserPosts(userID int64) ([]models.Post, error)
}

type Handler struct {
	keeper keeper
}

func New(keeper keeper) *Handler {
	return &Handler{keeper: keeper}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := int64(1)
	posts, err := h.keeper.GetUserPosts(userID)
	if err != nil {
		utils.WriteInternalError(w)
		fmt.Print(err)
		return
	}

	jsonResponse, err := json.Marshal(schema.ConvertPosts(posts))
	if err != nil {
		utils.WriteInternalError(w)
		return
	}

	utils.WriteJSON(w, jsonResponse)
}
