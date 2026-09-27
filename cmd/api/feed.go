package main

import (
	"net/http"

	"github.com/Intro0/goSocial/internal/store"
)

// getUserFeedHandler godoc
//
// @Summary Fetch the user feed
// @Tags feed
// @Produce json
// @Param since query string false "Only posts created at or after this time"
// @Param until query string false "Only posts created at or before this time"
// @Param limit query int false "Maximum number of posts"
// @Param offset query int false "Number of posts to skip"
// @Param sort query string false "Sort order: asc or desc"
// @Param tags query string false "Comma-separated tags"
// @Param search query string false "Text to search for"
// @Success 200 {array} store.PostWithMetaData
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
