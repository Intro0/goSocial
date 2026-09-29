package main

import (
	"net/http"

	"github.com/Intro0/goSocial/internal/store"
)

// getUserFeedHandler godoc
//
// @Summary Fetch the user feed
// @Description Fetch posts from the current development user's feed with optional filtering and pagination.
// @Tags feed
// @Produce json
// @Param since query string false "Posts created at or after this time in YYYY-MM-DD HH:MM:SS format"
// @Param until query string false "Posts created at or before this time in YYYY-MM-DD HH:MM:SS format"
// @Param limit query int false "Maximum number of posts, from 1 to 20"
// @Param offset query int false "Number of posts to skip"
// @Param sort query string false "Sort order: asc or desc"
// @Param tags query string false "Comma-separated tags"
// @Param search query string false "Text to search for in post titles and content"
// @Success 200 {object} map[string][]store.PostWithMetaData
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(42), fq)
	if err != nil {
		app.internalServiceError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServiceError(w, r, err)
	}
}
