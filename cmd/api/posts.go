package main

import (
	"net/http"

	"github.com/Intro0/goSocial/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  1,
		Tags:    payload.Tags,
	}

	if err := app.store.Posts.Create(r.Context(), post); err != nil {
		app.internalServiceError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServiceError(w, r, err)
	}
}
