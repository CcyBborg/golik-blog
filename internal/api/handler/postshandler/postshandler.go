package postshandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/CcyBborg/golik-blog/internal/api/schema"
	"github.com/CcyBborg/golik-blog/internal/api/utils"
	"github.com/CcyBborg/golik-blog/internal/models"
	"github.com/CcyBborg/golik-blog/internal/store"
)

type keeper interface {
	GetPosts(opts store.Opts) ([]models.Post, error)
	InsertPost(post models.Post) (postID int64, err error)
}

type Handler struct {
	keeper keeper
}

func New(keeper keeper) *Handler {
	return &Handler{keeper: keeper}
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

		posts, err := h.keeper.GetPosts(store.Opts{
			Pagination: store.Pagination{
				Offset: offset,
				Limit:  limit,
			},
			Sort: store.Sort{
				UpdatedAt: &store.SortDesc,
			},
		})
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		jsonResponse, err := json.Marshal(schema.ConvertPosts(posts))
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		utils.WriteJSON(w, jsonResponse)
	} else if r.Method == http.MethodPost {
		userID := int64(1) // Remove after JWT

		if err := r.ParseForm(); err != nil {
			utils.WriteInvalidParams(w)
		}

		title := r.Form.Get("title")
		if title == "" {
			utils.WriteInvalidParams(w)
			return
		}

		summary := r.Form.Get("summary")
		if summary == "" {
			utils.WriteInvalidParams(w)
			return
		}

		content := r.Form.Get("content")
		if content == "" {
			utils.WriteInvalidParams(w)
			return
		}

		categories := r.Form.Get("categoryList")
		if categories == "" {
			utils.WriteInvalidParams(w)
			return
		}

		categoriesParsed := strings.Split(categories, ",")

		newPost := models.Post{
			Author: models.User{
				ID: userID,
			},
			Title:      title,
			Summary:    summary,
			Content:    content,
			Categories: make([]models.Category, len(categoriesParsed)),
		}

		for i, category := range categoriesParsed {
			id, err := strconv.ParseInt(category, 10, 64)
			if err != nil {
				utils.WriteInvalidParams(w)
				return
			}
			newPost.Categories[i].ID = id
		}

		postID, err := h.keeper.InsertPost(newPost)
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		jsonResponse, err := json.Marshal(postID)
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		utils.WriteJSON(w, jsonResponse)
	}
}
