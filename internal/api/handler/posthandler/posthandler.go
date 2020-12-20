package posthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CcyBborg/golik-blog/internal/api/schema"

	"github.com/CcyBborg/golik-blog/internal/api/utils"

	"github.com/CcyBborg/golik-blog/internal/models"
	"github.com/gorilla/mux"
)

type keeper interface {
	GetPost(postID int64) (models.Post, error)
	UpdatePost(post models.Post, publish bool) error
	DeletePost(postID int64) error
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
		post, err := h.keeper.GetPost(postID)
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		jsonResponse, err := json.Marshal(schema.ConvertPost(post))
		if err != nil {
			utils.WriteInternalError(w)
			return
		}

		utils.WriteJSON(w, jsonResponse)
	} else if r.Method == http.MethodPatch {
		userID := int64(1) // Remove after JWT

		post, err := h.keeper.GetPost(postID)
		if err != nil {
			utils.WriteInvalidParams(w)
			return
		}

		if post.Author.ID != userID {
			utils.WriteUnauthorized(w)
			return
		}

		if err := r.ParseForm(); err != nil {
			utils.WriteInvalidParams(w)
			return
		}

		title := r.Form.Get("title")
		if title != "" {
			post.Title = title
		}

		summary := r.Form.Get("summary")
		if summary != "" {
			post.Summary = summary
		}

		content := r.Form.Get("content")
		if content != "" {
			post.Content = content
		}

		publish, err := strconv.ParseBool(r.Form.Get("publish"))
		if err != nil {
			utils.WriteInvalidParams(w)
			return
		}

		categories := r.Form.Get("categoryList")
		if categories != "" {
			categoriesParsed := strings.Split(categories, ",")
			for i, category := range categoriesParsed {
				id, err := strconv.ParseInt(category, 10, 64)
				if err != nil {
					utils.WriteInvalidParams(w)
					return
				}
				post.Categories[i].ID = id
			}
		}

		if err = h.keeper.UpdatePost(post, publish); err != nil {
			fmt.Println(err)
			utils.WriteInternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodDelete {
		userID := int64(1) // Remove after JWT

		post, err := h.keeper.GetPost(postID)
		if err != nil {
			utils.WriteInvalidParams(w)
			return
		}

		if post.Author.ID != userID {
			utils.WriteUnauthorized(w)
			return
		}

		if err := h.keeper.DeletePost(postID); err != nil {
			utils.WriteInternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
