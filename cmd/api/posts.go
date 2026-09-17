package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/go-chi/chi/v5"
)

type postKey string
const postCtx postKey = "post"
type CreatePostPayload struct { Title string `json:"title" validate:"required,max=100"`; Content string `json:"content" validate:"required,max=1000"`; Tags []string `json:"tags"` }
func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) { var payload CreatePostPayload; if err:=readJSON(w,r,&payload); err != nil { app.badRequestResponse(w,r,err); return }; if err:=Validate.Struct(payload); err != nil { app.badRequestResponse(w,r,err); return }; post:=&store.Post{Title:payload.Title,Content:payload.Content,UserID:1,Tags:payload.Tags}; if err:=app.store.Posts.Create(r.Context(),post); err != nil { app.internalServiceError(w,r,err); return }; if err:=writeJSON(w,http.StatusOK,post); err != nil { app.internalServiceError(w,r,err) } }
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) { post:=getPostFromCtx(r); comments,err:=app.store.Comments.GetByPostID(r.Context(),post.ID); if err != nil { app.internalServiceError(w,r,err); return }; post.Comments=comments; if err:=writeJSON(w,http.StatusOK,post); err != nil { app.internalServiceError(w,r,err) } }
func (app *application) postsContextMiddleware(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) { id,err:=strconv.ParseInt(chi.URLParam(r,"postID"),10,64); if err != nil { app.badRequestResponse(w,r,err); return }; post,err:=app.store.Posts.GetByID(r.Context(),id); if err != nil { if errors.Is(err,store.ErrNotFound) { app.notFoundResponse(w,r,err) } else { app.internalServiceError(w,r,err) }; return }; next.ServeHTTP(w,r.WithContext(context.WithValue(r.Context(),postCtx,post))) }) }
func getPostFromCtx(r *http.Request) *store.Post { post,_:=r.Context().Value(postCtx).(*store.Post); return post }
