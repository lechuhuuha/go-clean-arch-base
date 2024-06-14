package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	g *gin.RouterGroup
}

func NewRouter(g *gin.RouterGroup) *Router {
	return &Router{g: g}
}

func (r *Router) Use(middlewareFunc ...gin.HandlerFunc) {
	r.g.Use(middlewareFunc...)
}

func (r *Router) getHandlers(handlers []gin.HandlerFunc) []gin.HandlerFunc {
	return append([]gin.HandlerFunc{}, handlers...)
}

func (r *Router) Handle(method, path string, handlers ...gin.HandlerFunc) {
	handlers = r.getHandlers(handlers)
	r.g.Handle(method, path, handlers...)
}

func (r *Router) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return r.g.Group(relativePath, handlers...)
}

func (r *Router) Post(path string, handlers ...gin.HandlerFunc) {
	r.Handle(http.MethodPost, path, handlers...)
}

func (r *Router) Get(path string, handlers ...gin.HandlerFunc) {
	r.Handle(http.MethodGet, path, handlers...)
}

func (r *Router) Put(path string, handler gin.HandlerFunc) {
	r.Handle(http.MethodPut, path, handler)
}

func (r *Router) Delete(path string, handlers ...gin.HandlerFunc) {
	r.Handle(http.MethodDelete, path, handlers...)
}

func (r *Router) Patch(path string, handler gin.HandlerFunc) {
	r.Handle(http.MethodPatch, path, handler)
}
