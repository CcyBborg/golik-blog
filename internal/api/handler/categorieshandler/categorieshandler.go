package categorieshandler

import (
	"encoding/json"
	"net/http"

	"github.com/CcyBborg/golik-blog/internal/api/schema"

	"github.com/CcyBborg/golik-blog/internal/api/utils"

	"github.com/CcyBborg/golik-blog/internal/models"
)

type keeper interface {
	GetCategories() ([]models.Category, error)
}

type Handler struct {
	keeper keeper
}

func New(keeper keeper) *Handler {
	return &Handler{keeper: keeper}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories, err := h.keeper.GetCategories()
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		jsonResponse, err := json.Marshal(schema.ConvertCategories(categories))
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		utils.WriteJSON(w, jsonResponse)
	}
}
