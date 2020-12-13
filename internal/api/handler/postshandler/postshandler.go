package postshandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CcyBborg/golik-blog/internal/api/utils"
	"github.com/CcyBborg/golik-blog/internal/models"
	"github.com/CcyBborg/golik-blog/internal/services/posts"
)

type getter interface {
	GetPosts(opts posts.Opts) ([]models.Post, error)
}

type Handler struct {
	getter getter
}

func New(getter getter) *Handler {
	return &Handler{getter: getter}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		limitParam, ok := r.URL.Query()["limit"]
		if !ok || len(limitParam[0]) < 1 {
			utils.WriteInvalidParams(w)
			return
		}
		offsetParam, ok := r.URL.Query()["offset"]
		if !ok || len(offsetParam[0]) < 1 {
			utils.WriteInvalidParams(w)
			return
		}

		limit, err := strconv.ParseInt(limitParam[0], 10, 64)
		if err != nil {
			utils.WriteInvalidParams(w)
			return
		}
		offset, err := strconv.ParseInt(offsetParam[0], 10, 64)
		if err != nil {
			utils.WriteInvalidParams(w)
			return
		}

		posts, err := h.getter.GetPosts(posts.Opts{
			Pagination: posts.Pagination{
				Offset: offset,
				Limit:  limit,
			},
			Sort: posts.Sort{
				UpdatedAt: &posts.SortDesc,
			},
		})
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		schemaPosts := convertPosts(posts)

		json, err := json.Marshal(schemaPosts)
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		utils.WriteJSON(w, json)
	} else if r.Method == http.MethodPost {

	}
}
